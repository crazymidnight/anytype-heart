package fileobject

import (
	"context"
	"fmt"

	"github.com/anyproto/any-sync/app"
	"github.com/gogo/protobuf/types"
	"github.com/ipfs/go-cid"

	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/object/idresolver"
	"github.com/anyproto/anytype-heart/core/block/object/objectcreator"
	"github.com/anyproto/anytype-heart/core/block/simple"
	"github.com/anyproto/anytype-heart/core/block/source"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/files"
	"github.com/anyproto/anytype-heart/core/filestorage/filesync"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/pkg/lib/mill"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

var log = logging.Logger("fileobject")

const CName = "fileobject"

type Service interface {
	app.Component

	Create(ctx context.Context, spaceId string, req CreateRequest) (id string, object *types.Struct, err error)
	GetFileIdFromObject(ctx context.Context, objectId string) (domain.FullFileId, error)
	GetObjectIdByFileId(fileId domain.FileId) (string, error)
	Migrate(st *state.State, spc source.Space, keys []*pb.ChangeFileKeys) error
}

type service struct {
	spaceService  space.Service
	resolver      idresolver.Resolver
	objectCreator objectcreator.Service
	fileService   files.Service
	fileSync      filesync.FileSync
	objectStore   objectstore.ObjectStore
}

func New() Service {
	return &service{}
}

func (s *service) Name() string {
	return CName
}

func (s *service) Init(a *app.App) error {
	s.spaceService = app.MustComponent[space.Service](a)
	s.resolver = app.MustComponent[idresolver.Resolver](a)
	s.objectCreator = app.MustComponent[objectcreator.Service](a)
	s.fileService = app.MustComponent[files.Service](a)
	s.fileSync = app.MustComponent[filesync.FileSync](a)
	s.objectStore = app.MustComponent[objectstore.ObjectStore](a)
	return nil
}

type CreateRequest struct {
	FileId         domain.FileId
	EncryptionKeys map[string]string
	IsImported     bool
}

func (s *service) Create(ctx context.Context, spaceId string, req CreateRequest) (id string, object *types.Struct, err error) {
	space, err := s.spaceService.Get(ctx, spaceId)
	if err != nil {
		return "", nil, fmt.Errorf("get space: %w", err)
	}
	return s.createInSpace(ctx, space, req)
}

func (s *service) createInSpace(ctx context.Context, space space.Space, req CreateRequest) (id string, object *types.Struct, err error) {
	if req.FileId == "" {
		return "", nil, fmt.Errorf("file hash is empty")
	}
	details, typeKey, err := s.getDetailsForFileOrImage(ctx, domain.FullFileId{
		SpaceId: space.Id(),
		FileId:  req.FileId,
	})
	if err != nil {
		return "", nil, fmt.Errorf("get details for file or image: %w", err)
	}
	details.Fields[bundle.RelationKeyFileId.String()] = pbtypes.String(req.FileId.String())

	createState := state.NewDoc("", nil).(*state.State)
	createState.SetDetails(details)
	createState.SetFileInfo(state.FileInfo{
		FileId:         req.FileId,
		EncryptionKeys: req.EncryptionKeys,
	})

	id, object, err = s.objectCreator.CreateSmartBlockFromStateInSpace(ctx, space, []domain.TypeKey{typeKey}, createState)
	if err != nil {
		return "", nil, fmt.Errorf("create object: %w", err)
	}

	err = s.addToSyncQueue(domain.FullFileId{SpaceId: space.Id(), FileId: req.FileId}, true, req.IsImported)
	if err != nil {
		return "", nil, fmt.Errorf("add to sync queue: %w", err)
	}
	return id, object, nil
}

func (s *service) getDetailsForFileOrImage(ctx context.Context, id domain.FullFileId) (*types.Struct, domain.TypeKey, error) {
	file, err := s.fileService.FileByHash(ctx, id)
	if err != nil {
		return nil, "", err
	}
	if mill.IsImage(file.Info().Media) {
		image, err := s.fileService.ImageByHash(ctx, id)
		if err != nil {
			return nil, "", err
		}
		details, err := image.Details(ctx)
		if err != nil {
			return nil, "", err
		}
		return details, bundle.TypeKeyImage, nil
	}

	d, typeKey, err := file.Details(ctx)
	if err != nil {
		return nil, "", err
	}
	return d, typeKey, nil
}

func (s *service) addToSyncQueue(id domain.FullFileId, uploadedByUser bool, imported bool) error {
	if err := s.fileSync.AddFile(id.SpaceId, id.FileId, uploadedByUser, imported); err != nil {
		return fmt.Errorf("add file to sync queue: %w", err)
	}
	// TODO Maybe we need a watcher here?
	return nil
}

func (s *service) GetObjectIdByFileId(fileId domain.FileId) (string, error) {
	records, _, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId.String()),
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("query objects by file hash: %w", err)
	}
	if len(records) == 0 {
		return "", fmt.Errorf("file object not found")
	}
	return pbtypes.GetString(records[0].Details, bundle.RelationKeyId.String()), nil
}

func (s *service) GetFileIdFromObject(ctx context.Context, objectId string) (domain.FullFileId, error) {
	spaceId, err := s.resolver.ResolveSpaceID(objectId)
	if err != nil {
		return domain.FullFileId{}, fmt.Errorf("resolve spaceId: %w", err)
	}

	space, err := s.spaceService.Get(ctx, spaceId)
	if err != nil {
		return domain.FullFileId{}, fmt.Errorf("get space: %w", err)
	}

	return s.getFileIdFromObjectInSpace(ctx, space, objectId)
}

func (s *service) getFileIdFromObjectInSpace(ctx context.Context, space smartblock.Space, objectId string) (domain.FullFileId, error) {
	var fileId string
	err := space.Do(objectId, func(sb smartblock.SmartBlock) error {
		fileId = pbtypes.GetString(sb.Details(), bundle.RelationKeyFileId.String())
		if fileId == "" {
			return fmt.Errorf("empty file hash")
		}
		return nil
	})
	if err != nil {
		return domain.FullFileId{}, fmt.Errorf("get file object: %w", err)
	}

	return domain.FullFileId{
		SpaceId: space.Id(),
		FileId:  domain.FileId(fileId),
	}, nil
}

func (s *service) migrate(space space.Space, keys []*pb.ChangeFileKeys, hash string) string {
	if hash == "" {
		return hash
	}
	var fileKeys map[string]string
	for _, k := range keys {
		if k.Hash == hash {
			fileKeys = k.Keys
		}
	}

	fileObjectId, err := s.GetObjectIdByFileId(domain.FileId(hash))
	if err == nil {
		fmt.Println("FILE OBJECT ID", hash, "->", fileObjectId)
		return fileObjectId
	}

	fileObjectId, _, err = s.createInSpace(context.Background(), space, CreateRequest{
		FileId:         domain.FileId(hash),
		EncryptionKeys: fileKeys,
		IsImported:     false, // TODO what to do?
	})
	if err != nil {
		log.Errorf("create file object for hash %s: %v", hash, err)
		return hash
	}

	fmt.Println("MIGRATED FILE OBJECT ID", hash, "->", fileObjectId)

	return fileObjectId
}

func (s *service) Migrate(st *state.State, spc source.Space, keys []*pb.ChangeFileKeys) error {
	st.Iterate(func(b simple.Block) (isContinue bool) {
		if fh, ok := b.(simple.FileHashes); ok {
			fh.MigrateFile(func(oldHash string) (newHash string) {
				return s.migrate(spc.(space.Space), keys, oldHash)
			})
		}
		return true
	})
	det := st.Details()
	if det == nil || det.Fields == nil {
		return nil
	}

	for _, key := range st.FileRelationKeys() {
		if key == bundle.RelationKeyCoverId.String() {
			v := pbtypes.GetString(det, key)
			_, err := cid.Decode(v)
			if err != nil {
				// this is an exception cause coverId can contains not a file hash but color
				continue
			}
		}
		if hashList := pbtypes.GetStringList(det, key); hashList != nil {
			var anyChanges bool
			for i, hash := range hashList {
				if hash == "" {
					continue
				}
				newHash := s.migrate(spc.(space.Space), keys, hash)
				if hash != newHash {
					hashList[i] = newHash
					anyChanges = true
				}
			}
			if anyChanges {
				st.SetDetail(key, pbtypes.StringList(hashList))
			}
		}
	}
	return nil
}
