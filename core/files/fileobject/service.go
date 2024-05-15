package fileobject

import (
	"context"
	"fmt"
	"time"

	"github.com/anyproto/any-sync/app"
	"github.com/avast/retry-go/v4"
	"github.com/gogo/protobuf/types"
	"github.com/ipfs/go-cid"

	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/editor/template"
	"github.com/anyproto/anytype-heart/core/block/object/idresolver"
	"github.com/anyproto/anytype-heart/core/block/object/objectcreator"
	"github.com/anyproto/anytype-heart/core/block/object/payloadcreator"
	"github.com/anyproto/anytype-heart/core/block/source"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/domain/objectorigin"
	"github.com/anyproto/anytype-heart/core/files"
	"github.com/anyproto/anytype-heart/core/filestorage"
	"github.com/anyproto/anytype-heart/core/filestorage/filesync"
	"github.com/anyproto/anytype-heart/core/syncstatus/filesyncstatus"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/datastore"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space"
	"github.com/anyproto/anytype-heart/space/clientspace"
	"github.com/anyproto/anytype-heart/util/pbtypes"
	"github.com/anyproto/anytype-heart/util/persistentqueue"
)

// TODO UNsugar
var log = logging.Logger("fileobject")

var (
	ErrObjectNotFound = fmt.Errorf("file object not found")
	ErrEmptyFileId    = fmt.Errorf("empty file id")
)

const CName = "fileobject"

type Service interface {
	app.ComponentRunnable

	InitEmptyFileState(st *state.State)
	DeleteFileData(objectId string) error
	Create(ctx context.Context, spaceId string, req CreateRequest) (id string, object *types.Struct, err error)
	CreateFromImport(fileId domain.FullFileId, origin objectorigin.ObjectOrigin) (string, error)
	GetFileIdFromObject(objectId string) (domain.FullFileId, error)
	GetFileIdFromObjectWaitLoad(ctx context.Context, objectId string) (domain.FullFileId, error)
	GetObjectDetailsByFileId(fileId domain.FullFileId) (string, *types.Struct, error)
	MigrateFileIdsInDetails(st *state.State, spc source.Space)
	MigrateFileIdsInBlocks(st *state.State, spc source.Space)
	MigrateFiles(st *state.State, spc source.Space, keysChanges []*pb.ChangeFileKeys)
	EnsureFileAddedToSyncQueue(id domain.FullID, details *types.Struct) error

	FileOffload(ctx context.Context, objectId string, includeNotPinned bool) (totalSize uint64, err error)
	FilesOffload(ctx context.Context, objectIds []string, includeNotPinned bool) (filesOffloaded int, totalSize uint64, err error)
	FileSpaceOffload(ctx context.Context, spaceId string, includeNotPinned bool) (filesOffloaded int, totalSize uint64, err error)
}

type objectCreatorService interface {
	CreateSmartBlockFromStateInSpaceWithOptions(ctx context.Context, space clientspace.Space, objectTypeKeys []domain.TypeKey, createState *state.State, opts ...objectcreator.CreateOption) (id string, newDetails *types.Struct, err error)
}

type service struct {
	spaceService    space.Service
	objectCreator   objectCreatorService
	fileService     files.Service
	fileSync        filesync.FileSync
	fileStore       filestore.FileStore
	fileStorage     filestorage.FileStorage
	objectStore     objectstore.ObjectStore
	spaceIdResolver idresolver.Resolver
	migrationQueue  *persistentqueue.Queue[*migrationItem]

	indexer *indexer

	resolverRetryStartDelay time.Duration
	resolverRetryMaxDelay   time.Duration
}

func New(
	resolverRetryStartDelay time.Duration,
	resolverRetryMaxDelay time.Duration,
) Service {
	return &service{
		resolverRetryStartDelay: resolverRetryStartDelay,
		resolverRetryMaxDelay:   resolverRetryMaxDelay,
	}
}

func (s *service) Name() string {
	return CName
}

func (s *service) Init(a *app.App) error {
	s.spaceService = app.MustComponent[space.Service](a)
	s.objectCreator = app.MustComponent[objectCreatorService](a)
	s.fileService = app.MustComponent[files.Service](a)
	s.fileSync = app.MustComponent[filesync.FileSync](a)
	s.objectStore = app.MustComponent[objectstore.ObjectStore](a)
	s.fileStore = app.MustComponent[filestore.FileStore](a)
	s.spaceIdResolver = app.MustComponent[idresolver.Resolver](a)
	s.fileStorage = app.MustComponent[filestorage.FileStorage](a)

	s.indexer = s.newIndexer()

	dbProvider := app.MustComponent[datastore.Datastore](a)
	db, err := dbProvider.LocalStorage()
	if err != nil {
		return fmt.Errorf("get badger: %w", err)
	}
	s.migrationQueue = persistentqueue.New(persistentqueue.NewBadgerStorage(db, []byte("queue/file_migration/"), makeMigrationItem), log.Desugar(), s.migrationQueueHandler)
	return nil
}

func (s *service) Run(_ context.Context) error {
	go func() {
		err := s.ensureNotSyncedFilesAddedToQueue()
		if err != nil {
			log.Errorf("ensure not synced files added to queue: %v", err)
		}
	}()
	s.indexer.run()
	s.migrationQueue.Run()
	return nil
}

// After migrating to new sync queue we need to ensure that all not synced files are added to the queue
func (s *service) ensureNotSyncedFilesAddedToQueue() error {
	records, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_NotEmpty,
			},
			{
				RelationKey: bundle.RelationKeyFileBackupStatus.String(),
				Condition:   model.BlockContentDataviewFilter_NotEqual,
				Value:       pbtypes.Int64(int64(filesyncstatus.Synced)),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("query file objects: %w", err)
	}

	for _, record := range records {
		fullId := extractFullFileIdFromDetails(record.Details)
		id := pbtypes.GetString(record.Details, bundle.RelationKeyId.String())
		err := s.addToSyncQueue(id, fullId, false, false)
		if err != nil {
			log.Errorf("add to sync queue: %v", err)
		}
	}

	return nil
}

func extractFullFileIdFromDetails(details *types.Struct) domain.FullFileId {
	return domain.FullFileId{
		SpaceId: pbtypes.GetString(details, bundle.RelationKeySpaceId.String()),
		FileId:  domain.FileId(pbtypes.GetString(details, bundle.RelationKeyFileId.String())),
	}
}

// EnsureFileAddedToSyncQueue adds file to sync queue if it is not synced yet, we need to do this
// after migrating to new sync queue
func (s *service) EnsureFileAddedToSyncQueue(id domain.FullID, details *types.Struct) error {
	if pbtypes.GetInt64(details, bundle.RelationKeyFileBackupStatus.String()) == int64(filesyncstatus.Synced) {
		return nil
	}
	fullId := domain.FullFileId{
		SpaceId: id.SpaceID,
		FileId:  domain.FileId(pbtypes.GetString(details, bundle.RelationKeyFileId.String())),
	}
	err := s.addToSyncQueue(id.ObjectID, fullId, false, false)
	return err
}

func (s *service) Close(ctx context.Context) error {
	return s.indexer.close()
}

type CreateRequest struct {
	FileId                domain.FileId
	EncryptionKeys        map[string]string
	ObjectOrigin          objectorigin.ObjectOrigin
	AdditionalDetails     *types.Struct
	AsyncMetadataIndexing bool
}

func (s *service) InitEmptyFileState(st *state.State) {
	template.InitTemplate(st,
		template.WithEmpty,
		template.WithTitle,
		template.WithDefaultFeaturedRelations,
		template.WithFeaturedRelations,
		template.WithAllBlocksEditsRestricted,
	)
}

func (s *service) Create(ctx context.Context, spaceId string, req CreateRequest) (id string, object *types.Struct, err error) {
	space, err := s.spaceService.Get(ctx, spaceId)
	if err != nil {
		return "", nil, fmt.Errorf("get space: %w", err)
	}

	id, object, err = s.createInSpace(ctx, space, req)
	if err != nil {
		return "", nil, fmt.Errorf("create in space: %w", err)
	}
	err = s.addToSyncQueue(id, domain.FullFileId{SpaceId: space.Id(), FileId: req.FileId}, true, req.ObjectOrigin.IsImported())
	if err != nil {
		return "", nil, fmt.Errorf("add to sync queue: %w", err)
	}

	return id, object, nil
}

func (s *service) createInSpace(ctx context.Context, space clientspace.Space, req CreateRequest) (id string, object *types.Struct, err error) {
	if req.FileId == "" {
		return "", nil, fmt.Errorf("file hash is empty")
	}

	details := s.makeInitialDetails(req.FileId, req.ObjectOrigin)

	payload, err := space.CreateTreePayload(ctx, payloadcreator.PayloadCreationParams{
		Time:           time.Now(),
		SmartblockType: coresb.SmartBlockTypeFileObject,
	})
	if err != nil {
		return "", nil, fmt.Errorf("create tree payload: %w", err)
	}

	createState := state.NewDoc(payload.RootRawChange.Id, nil).(*state.State)
	createState.SetDetails(details)
	createState.SetFileInfo(state.FileInfo{
		FileId:         req.FileId,
		EncryptionKeys: req.EncryptionKeys,
	})
	if !req.AsyncMetadataIndexing {
		s.InitEmptyFileState(createState)
		fullFileId := domain.FullFileId{SpaceId: space.Id(), FileId: req.FileId}
		fullObjectId := domain.FullID{SpaceID: space.Id(), ObjectID: payload.RootRawChange.Id}
		err := s.indexer.injectMetadataToState(ctx, createState, fullFileId, fullObjectId)
		if err != nil {
			return "", nil, fmt.Errorf("inject metadata to state: %w", err)
		}
	}

	if req.AdditionalDetails != nil {
		for k, v := range req.AdditionalDetails.GetFields() {
			createState.SetDetailAndBundledRelation(domain.RelationKey(k), v)
		}
	}

	// Type will be changed after indexing, just use general type File for now
	id, object, err = s.objectCreator.CreateSmartBlockFromStateInSpaceWithOptions(ctx, space, []domain.TypeKey{bundle.TypeKeyFile}, createState, objectcreator.WithPayload(&payload))
	if err != nil {
		return "", nil, fmt.Errorf("create object: %w", err)
	}

	if req.AsyncMetadataIndexing {
		err = s.indexer.addToQueue(ctx, domain.FullID{SpaceID: space.Id(), ObjectID: id}, domain.FullFileId{SpaceId: space.Id(), FileId: req.FileId})
		if err != nil {
			// Will be retried in background, so don't return error
			log.Errorf("add to index queue: %v", err)
		}
	}

	return id, object, nil
}

func (s *service) makeInitialDetails(fileId domain.FileId, origin objectorigin.ObjectOrigin) *types.Struct {
	details := &types.Struct{
		Fields: map[string]*types.Value{
			bundle.RelationKeyFileId.String(): pbtypes.String(fileId.String()),
			// Use general file layout. It will be changed for proper layout after indexing
			bundle.RelationKeyLayout.String():             pbtypes.Int64(int64(model.ObjectType_file)),
			bundle.RelationKeyFileIndexingStatus.String(): pbtypes.Int64(int64(model.FileIndexingStatus_NotIndexed)),
		},
	}
	origin.AddToDetails(details)
	return details
}

// CreateFromImport creates file object from imported raw IPFS file. Encryption keys for this file should exist in file store.
func (s *service) CreateFromImport(fileId domain.FullFileId, origin objectorigin.ObjectOrigin) (string, error) {
	// Check that fileId is not a file object id
	recs, _, err := s.objectStore.QueryObjectIDs(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId.FileId.String()),
			},
			{
				RelationKey: bundle.RelationKeySpaceId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId.SpaceId),
			},
		},
	})
	if err == nil && len(recs) > 0 {
		return recs[0], nil
	}

	fileObjectId, _, err := s.GetObjectDetailsByFileId(fileId)
	if err == nil {
		return fileObjectId, nil
	}
	keys, err := s.fileStore.GetFileKeys(fileId.FileId)
	if err != nil {
		return "", fmt.Errorf("get file keys: %w", err)
	}
	fileObjectId, _, err = s.Create(context.Background(), fileId.SpaceId, CreateRequest{
		FileId:                fileId.FileId,
		EncryptionKeys:        keys,
		ObjectOrigin:          origin,
		AsyncMetadataIndexing: true,
	})
	if err != nil {
		return "", fmt.Errorf("create object: %w", err)
	}
	return fileObjectId, nil
}

func (s *service) addToSyncQueue(objectId string, fileId domain.FullFileId, uploadedByUser bool, imported bool) error {
	if err := s.fileSync.AddFile(objectId, fileId, uploadedByUser, imported); err != nil {
		return fmt.Errorf("add file to sync queue: %w", err)
	}
	return nil
}

func (s *service) GetObjectIdByFileId(fileId domain.FullFileId) (string, error) {
	records, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId.FileId.String()),
			},
			{
				RelationKey: bundle.RelationKeySpaceId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId.SpaceId),
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("query objects by file hash: %w", err)
	}
	if len(records) == 0 {
		return "", ErrObjectNotFound
	}
	return pbtypes.GetString(records[0].Details, bundle.RelationKeyId.String()), nil
}

func (s *service) GetObjectDetailsByFileId(fileId domain.FullFileId) (string, *types.Struct, error) {
	records, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId.FileId.String()),
			},
			{
				RelationKey: bundle.RelationKeySpaceId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId.SpaceId),
			},
		},
	})
	if err != nil {
		return "", nil, fmt.Errorf("query objects by file hash: %w", err)
	}
	if len(records) == 0 {
		return "", nil, ErrObjectNotFound
	}
	details := records[0].Details
	return pbtypes.GetString(details, bundle.RelationKeyId.String()), details, nil
}

func (s *service) GetFileIdFromObject(objectId string) (domain.FullFileId, error) {
	details, err := s.objectStore.GetDetails(objectId)
	if err != nil {
		return domain.FullFileId{}, fmt.Errorf("get object details: %w", err)
	}
	spaceId := pbtypes.GetString(details.Details, bundle.RelationKeySpaceId.String())
	fileId := pbtypes.GetString(details.Details, bundle.RelationKeyFileId.String())
	if fileId == "" {
		return domain.FullFileId{}, ErrEmptyFileId
	}
	return domain.FullFileId{
		SpaceId: spaceId,
		FileId:  domain.FileId(fileId),
	}, nil
}

func (s *service) GetFileIdFromObjectWaitLoad(ctx context.Context, objectId string) (domain.FullFileId, error) {
	spaceId, err := s.resolveSpaceIdWithRetry(ctx, objectId)
	if err != nil {
		return domain.FullFileId{}, fmt.Errorf("resolve space id: %w", err)
	}
	spc, err := s.spaceService.Get(ctx, spaceId)
	if err != nil {
		return domain.FullFileId{}, fmt.Errorf("get space: %w", err)
	}
	id := domain.FullFileId{
		SpaceId: spaceId,
	}
	err = spc.Do(objectId, func(sb smartblock.SmartBlock) error {
		details := sb.Details()
		id.FileId = domain.FileId(pbtypes.GetString(details, bundle.RelationKeyFileId.String()))
		if id.FileId == "" {
			return ErrEmptyFileId
		}
		return nil
	})
	if err != nil {
		return domain.FullFileId{}, fmt.Errorf("get object details: %w", err)
	}
	return id, nil
}

func (s *service) resolveSpaceIdWithRetry(ctx context.Context, objectId string) (string, error) {
	_, err := cid.Decode(objectId)
	if err != nil {
		return "", fmt.Errorf("decode object id: %w", err)
	}
	if domain.IsFileId(objectId) {
		return "", fmt.Errorf("object id is file cid")
	}

	spaceId, err := retry.DoWithData(func() (string, error) {
		return s.spaceIdResolver.ResolveSpaceID(objectId)
	},
		retry.Context(ctx),
		retry.Attempts(0),
		retry.Delay(s.resolverRetryStartDelay),
		retry.MaxDelay(s.resolverRetryMaxDelay),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
	)
	return spaceId, err
}

func (s *service) FileOffload(ctx context.Context, objectId string, includeNotPinned bool) (totalSize uint64, err error) {
	details, err := s.objectStore.GetDetails(objectId)
	if err != nil {
		return 0, fmt.Errorf("get object details: %w", err)
	}
	return s.fileOffload(ctx, details.GetDetails(), includeNotPinned)
}

func (s *service) fileOffload(ctx context.Context, fileDetails *types.Struct, includeNotPinned bool) (uint64, error) {
	fileId := pbtypes.GetString(fileDetails, bundle.RelationKeyFileId.String())
	if fileId == "" {
		return 0, ErrEmptyFileId
	}
	backupStatus := filesyncstatus.Status(pbtypes.GetInt64(fileDetails, bundle.RelationKeyFileBackupStatus.String()))
	id := domain.FullFileId{
		SpaceId: pbtypes.GetString(fileDetails, bundle.RelationKeySpaceId.String()),
		FileId:  domain.FileId(fileId),
	}

	if !includeNotPinned && backupStatus != filesyncstatus.Synced {
		return 0, nil
	}

	return s.fileService.FileOffload(ctx, id)
}

func (s *service) FilesOffload(ctx context.Context, objectIds []string, includeNotPinned bool) (filesOffloaded int, totalSize uint64, err error) {
	if len(objectIds) == 0 {
		return s.offloadAllFiles(ctx, includeNotPinned)
	}

	for _, objectId := range objectIds {
		size, err := s.FileOffload(ctx, objectId, includeNotPinned)
		if err != nil {
			log.Errorf("failed to offload file %s: %v", objectId, err)
			continue
		}
		totalSize += size
		if size > 0 {
			filesOffloaded++
		}
	}
	return filesOffloaded, totalSize, nil
}

func (s *service) offloadAllFiles(ctx context.Context, includeNotPinned bool) (filesOffloaded int, totalSize uint64, err error) {
	records, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_NotEmpty,
			},
		},
	})
	if err != nil {
		return 0, 0, fmt.Errorf("query file objects by spaceId: %w", err)
	}
	for _, record := range records {
		size, err := s.fileOffload(ctx, record.Details, includeNotPinned)
		if err != nil {
			objectId := pbtypes.GetString(record.Details, bundle.RelationKeyId.String())
			log.Errorf("failed to offload file %s: %v", objectId, err)
			continue
		}
		totalSize += size
		if size > 0 {
			filesOffloaded++
		}
	}
	return filesOffloaded, totalSize, nil
}

func (s *service) FileSpaceOffload(ctx context.Context, spaceId string, includeNotPinned bool) (filesOffloaded int, totalSize uint64, err error) {
	records, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeySpaceId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(spaceId),
			},
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_NotEmpty,
			},
		},
	})
	if err != nil {
		return 0, 0, fmt.Errorf("query file objects by spaceId: %w", err)
	}
	for _, record := range records {
		fileId := pbtypes.GetString(record.Details, bundle.RelationKeyFileId.String())
		size, err := s.offloadFileSafe(ctx, spaceId, fileId, record, includeNotPinned)
		if err != nil {
			log.Errorf("failed to offload file %s: %v", fileId, err)
			return 0, 0, err
		}
		if size > 0 {
			filesOffloaded++
			err = s.fileStore.DeleteFile(domain.FileId(fileId))
			if err != nil {
				return 0, 0, fmt.Errorf("failed to delete file from store: %w", err)
			}
		}
		totalSize += size
	}
	return filesOffloaded, totalSize, nil
}

func (s *service) DeleteFileData(objectId string) error {
	fullId, err := s.GetFileIdFromObject(objectId)
	if err != nil {
		return fmt.Errorf("get file id from object: %w", err)
	}
	records, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyId.String(),
				Condition:   model.BlockContentDataviewFilter_NotEqual,
				Value:       pbtypes.String(objectId),
			},
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fullId.FileId.String()),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("list objects that use file id: %w", err)
	}
	if len(records) == 0 {
		if err := s.fileStore.DeleteFile(fullId.FileId); err != nil {
			return err
		}
		if err := s.fileSync.DeleteFile(objectId, fullId); err != nil {
			return fmt.Errorf("failed to remove file from sync: %w", err)
		}
		_, err = s.FileOffload(context.Background(), objectId, true)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *service) offloadFileSafe(ctx context.Context,
	spaceId string,
	fileId string,
	record database.Record,
	includeNotPinned bool,
) (uint64, error) {
	existingObjects, err := s.objectStore.Query(database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyFileId.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String(fileId),
			},
			{
				RelationKey: bundle.RelationKeySpaceId.String(),
				Condition:   model.BlockContentDataviewFilter_NotEqual,
				Value:       pbtypes.String(spaceId),
			},
		},
	})
	if err != nil {
		return 0, err
	}
	if len(existingObjects) > 0 {
		return s.fileOffload(ctx, record.Details, false)
	}
	return s.fileOffload(ctx, record.Details, includeNotPinned)
}
