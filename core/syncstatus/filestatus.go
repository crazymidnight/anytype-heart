package syncstatus

import (
	"context"
	"fmt"

	"github.com/anyproto/any-sync/commonspace/syncstatus"

	"github.com/anyproto/anytype-heart/core/block/cache"
	"github.com/anyproto/anytype-heart/core/block/editor/basic"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/syncstatus/filesyncstatus"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func (s *service) OnFileUploadStarted(objectId string) error {
	return s.indexFileSyncStatus(objectId, filesyncstatus.Syncing)
}

func (s *service) OnFileUploaded(objectId string) error {
	return s.indexFileSyncStatus(objectId, filesyncstatus.Synced)
}

func (s *service) OnFileLimited(objectId string) error {
	return s.indexFileSyncStatus(objectId, filesyncstatus.Limited)
}

func (s *service) indexFileSyncStatus(fileObjectId string, status filesyncstatus.Status) error {
	err := cache.Do(s.objectGetter, fileObjectId, func(sb smartblock.SmartBlock) (err error) {
		prevStatus := pbtypes.GetInt64(sb.Details(), bundle.RelationKeyFileBackupStatus.String())
		newStatus := int64(status)
		if prevStatus == newStatus {
			return nil
		}
		detailsSetter, ok := sb.(basic.DetailsSettable)
		if !ok {
			return fmt.Errorf("setting of details is not supported for %T", sb)
		}
		return detailsSetter.SetDetails(nil, []*model.Detail{
			{
				Key:   bundle.RelationKeyFileBackupStatus.String(),
				Value: pbtypes.Int64(newStatus),
			},
		}, true)
	})
	if err != nil {
		return fmt.Errorf("get object: %w", err)
	}

	err = s.updateReceiver.UpdateTree(context.Background(), fileObjectId, status.ToSyncStatus())
	if err != nil {
		return fmt.Errorf("update tree: %w", err)
	}

	s.sendSpaceStatusUpdate(status)
	return nil
}

func (s *service) sendSpaceStatusUpdate(status filesyncstatus.Status) {
	var (
		spaceStatus    syncstatus.SpaceSyncStatus
		numberOfObject int
		spaceError     syncstatus.SpaceSyncError
		syncInProgress bool
	)
	switch status {
	case filesyncstatus.Synced:
		spaceStatus = syncstatus.Synced
	case filesyncstatus.Syncing:
		spaceStatus = syncstatus.Syncing
		numberOfObject++
		syncInProgress = true
	case filesyncstatus.Limited:
		spaceStatus = syncstatus.Error
		spaceError = syncstatus.StorageLimitExceed
	case filesyncstatus.Unknown:
		spaceStatus = syncstatus.Error
		spaceError = syncstatus.NetworkError
	}

	syncStatus := syncstatus.MakeSyncStatus(spaceStatus, numberOfObject, spaceError, syncInProgress, false)
	s.spaceSyncStatus.SendUpdate(syncStatus)
}
