package objectid

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"

	"github.com/anyproto/anytype-heart/core/block"
	"github.com/anyproto/anytype-heart/core/block/import/common"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/domain/objectorigin"
	"github.com/anyproto/anytype-heart/core/files/fileobject"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

// oldFile represents file in pre Files-as-Objects format
type oldFile struct {
	blockService      *block.Service
	fileStore         filestore.FileStore
	fileObjectService fileobject.Service
}

func (f *oldFile) GetIDAndPayload(ctx context.Context, spaceId string, sn *common.Snapshot, _ time.Time, _ bool, origin objectorigin.ObjectOrigin) (string, treestorage.TreeStorageCreatePayload, error) {
	filePath := pbtypes.GetString(sn.Snapshot.Data.Details, bundle.RelationKeySource.String())
	if filePath != "" {
		fileObjectId, err := uploadFile(ctx, f.blockService, spaceId, filePath, origin)
		if err != nil {
			return "", treestorage.TreeStorageCreatePayload{}, fmt.Errorf("upload file: %w", err)
		}
		return fileObjectId, treestorage.TreeStorageCreatePayload{}, nil
	}

	fileId := pbtypes.GetString(sn.Snapshot.Data.Details, bundle.RelationKeyId.String())
	filesKeys := map[string]string{}
	for _, fileKeys := range sn.Snapshot.FileKeys {
		if fileKeys.Hash == fileId {
			filesKeys = fileKeys.Keys
			break
		}
	}
	err := f.fileStore.AddFileKeys(domain.FileEncryptionKeys{
		FileId:         domain.FileId(fileId),
		EncryptionKeys: filesKeys,
	})
	if err != nil {
		return "", treestorage.TreeStorageCreatePayload{}, fmt.Errorf("add file keys: %w", err)
	}
	objectId, err := f.fileObjectService.CreateFromImport(domain.FullFileId{SpaceId: spaceId, FileId: domain.FileId(fileId)}, origin)
	if err != nil {
		return "", treestorage.TreeStorageCreatePayload{}, fmt.Errorf("create file object: %w", err)
	}
	return objectId, treestorage.TreeStorageCreatePayload{}, nil
}

func uploadFile(ctx context.Context, blockService *block.Service, spaceId string, filePath string, origin objectorigin.ObjectOrigin) (string, error) {
	params := pb.RpcFileUploadRequest{
		SpaceId:   spaceId,
		LocalPath: filePath,
	}
	if strings.HasPrefix(filePath, "http://") || strings.HasPrefix(filePath, "https://") {
		params = pb.RpcFileUploadRequest{
			SpaceId: spaceId,
			Url:     filePath,
		}
	}
	dto := block.FileUploadRequest{
		RpcFileUploadRequest: params,
		ObjectOrigin:         origin,
	}

	fileObjectId, _, err := blockService.UploadFile(ctx, spaceId, dto)
	if err != nil {
		return "", err
	}
	return fileObjectId, nil
}
