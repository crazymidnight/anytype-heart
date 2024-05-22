package syncstatus

import (
	"context"
	"fmt"
	"sync"

	"github.com/anyproto/any-sync/nodeconf"

	"github.com/anyproto/anytype-heart/core/anytype/config"
	"github.com/anyproto/anytype-heart/core/event"
	"github.com/anyproto/anytype-heart/core/syncstatus/filesyncstatus"
	"github.com/anyproto/anytype-heart/core/syncstatus/helpers"
	"github.com/anyproto/anytype-heart/core/syncstatus/nodestatus"
	"github.com/anyproto/anytype-heart/core/syncstatus/objectsyncstatus"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
)

type updateReceiver struct {
	eventSender event.Sender

	nodeConfService nodeconf.Service
	sync.Mutex
	nodeConnected     bool
	lastStatus        map[string]pb.EventStatusThreadSyncStatus
	objectStore       objectstore.ObjectStore
	nodeStatus        nodestatus.NodeStatus
	syncStatusUpdater Updater
}

func newUpdateReceiver(nodeConfService nodeconf.Service, cfg *config.Config, eventSender event.Sender, objectStore objectstore.ObjectStore, nodeStatus nodestatus.NodeStatus, syncStatusUpdater Updater) *updateReceiver {
	if cfg.DisableThreadsSyncEvents {
		eventSender = nil
	}
	return &updateReceiver{
		nodeConfService:   nodeConfService,
		lastStatus:        make(map[string]pb.EventStatusThreadSyncStatus),
		eventSender:       eventSender,
		objectStore:       objectStore,
		nodeStatus:        nodeStatus,
		syncStatusUpdater: syncStatusUpdater,
	}
}

func (r *updateReceiver) UpdateTree(_ context.Context, objId string, status objectsyncstatus.SyncStatus) error {
	objStatus, syncError := r.getObjectStatus(objId, status)

	if !r.isStatusUpdated(objId, objStatus) {
		return nil
	}
	syncStatus := mapEventToSyncStatus(objStatus)
	err := r.syncStatusUpdater.UpdateDetails(objId, syncStatus, syncError)
	if err != nil {
		return err
	}
	r.notify(objId, objStatus)
	return nil
}

func (r *updateReceiver) isStatusUpdated(objectID string, objStatus pb.EventStatusThreadSyncStatus) bool {
	r.Lock()
	defer r.Unlock()
	if lastObjStatus, ok := r.lastStatus[objectID]; ok && objStatus == lastObjStatus {
		return false
	}
	r.lastStatus[objectID] = objStatus
	return true
}

func (r *updateReceiver) getFileStatus(fileId string) (filesyncstatus.Status, error) {
	details, err := r.objectStore.GetDetails(fileId)
	if err != nil {
		return filesyncstatus.Unknown, fmt.Errorf("get file details: %w", err)
	}
	if v, ok := details.GetDetails().GetFields()[bundle.RelationKeyFileBackupStatus.String()]; ok {
		return filesyncstatus.Status(v.GetNumberValue()), nil
	}
	return filesyncstatus.Unknown, fmt.Errorf("no backup status")
}

func (r *updateReceiver) getObjectStatus(objectId string, status objectsyncstatus.SyncStatus) (pb.EventStatusThreadSyncStatus, helpers.SyncError) {
	fileStatus, err := r.getFileStatus(objectId)
	var syncError helpers.SyncError
	if err == nil {
		// Prefer file backup status
		if fileStatus != filesyncstatus.Synced {
			status = fileStatus.ToSyncStatus()
		}
		if fileStatus == filesyncstatus.Limited {
			syncError = helpers.StorageLimitExceed
		}
	}

	if r.nodeConfService.NetworkCompatibilityStatus() == nodeconf.NetworkCompatibilityStatusIncompatible {
		return pb.EventStatusThread_IncompatibleVersion, helpers.IncompatibleVersion
	}

	if !r.isNodeConnected() {
		return pb.EventStatusThread_Offline, helpers.NetworkError
	}

	switch status {
	case objectsyncstatus.StatusUnknown:
		return pb.EventStatusThread_Unknown, syncError
	case objectsyncstatus.StatusSynced:
		return pb.EventStatusThread_Synced, syncError
	}
	return pb.EventStatusThread_Syncing, syncError
}

func (r *updateReceiver) ClearLastObjectStatus(objectID string) {
	r.Lock()
	defer r.Unlock()
	delete(r.lastStatus, objectID)
}

func (r *updateReceiver) isNodeConnected() bool {
	r.Lock()
	defer r.Unlock()
	return r.nodeConnected
}

func (r *updateReceiver) UpdateNodeConnection(online bool) {
	r.Lock()
	defer r.Unlock()
	r.nodeConnected = online
}

func (r *updateReceiver) UpdateNodeStatus() {
	r.Lock()
	defer r.Unlock()
	r.nodeConnected = r.nodeStatus.GetNodeStatus() == nodestatus.Online
}

func (r *updateReceiver) notify(
	objId string,
	objStatus pb.EventStatusThreadSyncStatus,
) {
	r.sendEvent(objId, &pb.EventMessageValueOfThreadStatus{ThreadStatus: &pb.EventStatusThread{
		Summary: &pb.EventStatusThreadSummary{Status: objStatus},
		Cafe: &pb.EventStatusThreadCafe{
			Status: objStatus,
			Files:  &pb.EventStatusThreadCafePinStatus{},
		},
	}})
}

func (r *updateReceiver) sendEvent(ctx string, event pb.IsEventMessageValue) {
	if r.eventSender == nil {
		return
	}
	r.eventSender.Broadcast(&pb.Event{
		Messages:  []*pb.EventMessage{{Value: event}},
		ContextId: ctx,
	})
}

func mapEventToSyncStatus(status pb.EventStatusThreadSyncStatus) helpers.SyncStatus {
	switch status {
	case pb.EventStatusThread_Syncing:
		return helpers.Syncing
	case pb.EventStatusThread_Offline, pb.EventStatusThread_Unknown, pb.EventStatusThread_IncompatibleVersion, pb.EventStatusThread_Failed:
		return helpers.Error
	default:
		return helpers.Synced
	}
}
