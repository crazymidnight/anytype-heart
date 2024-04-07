package aclobjectmanager

import (
	"context"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/debugstat"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/commonspace/object/acl/list"
	"github.com/anyproto/any-sync/util/crypto"
	"go.uber.org/zap"

	"github.com/anyproto/anytype-heart/space/clientspace"
	"github.com/anyproto/anytype-heart/space/internal/components/aclnotifications"
	"github.com/anyproto/anytype-heart/space/internal/components/dependencies"
	"github.com/anyproto/anytype-heart/space/internal/components/participantwatcher"
	"github.com/anyproto/anytype-heart/space/internal/components/spaceloader"
	"github.com/anyproto/anytype-heart/space/internal/components/spacestatus"
	"github.com/anyproto/anytype-heart/space/spaceinfo"
)

const CName = "common.components.aclobjectmanager"

var log = logger.NewNamed(CName)

type AclObjectManager interface {
	app.ComponentRunnable
}

func New(ownerMetadata []byte) AclObjectManager {
	return &aclObjectManager{
		ownerMetadata:     ownerMetadata,
		addedParticipants: make(map[string]struct{}),
	}
}

type aclObjectManager struct {
	ctx                 context.Context
	cancel              context.CancelFunc
	wait                chan struct{}
	waitLoad            chan struct{}
	sp                  clientspace.Space
	loadErr             error
	spaceLoader         spaceloader.SpaceLoader
	status              spacestatus.SpaceStatus
	indexer             dependencies.SpaceIndexer
	statService         debugstat.StatService
	started             bool
	notificationService aclnotifications.AclNotification
	participantWatcher  participantwatcher.ParticipantWatcher

	ownerMetadata     []byte
	lastIndexed       string
	addedParticipants map[string]struct{}
}

func (a *aclObjectManager) ProvideStat() any {
	select {
	case <-a.waitLoad:
		if a.loadErr != nil {
			return parseAcl(nil, a.status.SpaceId())
		}
		return parseAcl(a.sp.CommonSpace().Acl(), a.status.SpaceId())
	default:
		return parseAcl(nil, a.status.SpaceId())
	}
}

func (a *aclObjectManager) StatId() string {
	return a.status.SpaceId()
}

func (a *aclObjectManager) StatType() string {
	return CName
}

func (a *aclObjectManager) UpdateAcl(aclList list.AclList) {
	err := a.processAcl()
	if err != nil {
		log.Error("error processing acl", zap.Error(err))
	}
}

func (a *aclObjectManager) Init(ap *app.App) (err error) {
	a.spaceLoader = ap.MustComponent(spaceloader.CName).(spaceloader.SpaceLoader)
	a.indexer = app.MustComponent[dependencies.SpaceIndexer](ap)
	a.status = app.MustComponent[spacestatus.SpaceStatus](ap)
	a.participantWatcher = app.MustComponent[participantwatcher.ParticipantWatcher](ap)
	a.notificationService = app.MustComponent[aclnotifications.AclNotification](ap)
	a.statService, _ = ap.Component(debugstat.CName).(debugstat.StatService)
	if a.statService == nil {
		a.statService = debugstat.NewNoOp()
	}
	a.statService.AddProvider(a)
	a.waitLoad = make(chan struct{})
	a.wait = make(chan struct{})
	return nil
}

func (a *aclObjectManager) Name() (name string) {
	return CName
}

func (a *aclObjectManager) Run(ctx context.Context) (err error) {
	err = a.indexer.RemoveAclIndexes(a.status.SpaceId())
	if err != nil {
		return
	}
	a.started = true
	a.ctx, a.cancel = context.WithCancel(context.Background())
	go a.waitSpace()
	go a.process()
	return
}

func (a *aclObjectManager) Close(ctx context.Context) (err error) {
	if !a.started {
		return
	}
	a.cancel()
	<-a.wait
	a.statService.RemoveProvider(a)
	return
}

func (a *aclObjectManager) waitSpace() {
	a.sp, a.loadErr = a.spaceLoader.WaitLoad(a.ctx)
	close(a.waitLoad)
}

func (a *aclObjectManager) process() {
	defer close(a.wait)
	select {
	case <-a.ctx.Done():
		return
	case <-a.waitLoad:
		if a.loadErr != nil {
			return
		}
		break
	}

	err := a.participantWatcher.RegisterOwnerIdentity(a.ctx, a.sp)
	if err != nil {
		log.Error("init my identity", zap.Error(err))
	}

	common := a.sp.CommonSpace()
	common.Acl().SetAclUpdater(a)
	common.Acl().RLock()
	defer common.Acl().RUnlock()
	err = a.processAcl()
	if err != nil {
		log.Error("error processing acl", zap.Error(err))
	}
}

func (a *aclObjectManager) processAcl() (err error) {
	var (
		common   = a.sp.CommonSpace()
		acl      = common.Acl()
		aclState = acl.AclState()
	)
	defer func() {
		if err == nil {
			permissions := aclState.Permissions(aclState.AccountKey().GetPublic())
			a.notificationService.AddRecords(acl, permissions, common.Id(), spaceinfo.AccountStatusActive)
		}
	}()
	lastIndexed := a.lastIndexed
	if lastIndexed == acl.Head().Id {
		return
	}
	decrypt := func(key crypto.PubKey) ([]byte, error) {
		if a.ownerMetadata != nil {
			return a.ownerMetadata, nil
		}
		return aclState.GetMetadata(key, true)
	}
	states := aclState.CurrentAccounts()
	// decrypt all metadata
	states, err = decryptAll(states, decrypt)
	if err != nil {
		return
	}
	statusAclHeadId := a.status.GetLatestAclHeadId()
	upToDate := statusAclHeadId == "" || acl.HasHead(statusAclHeadId)
	err = a.processStates(states, upToDate, aclState.Identity())
	if err != nil {
		return
	}
	err = a.status.SetAclIsEmpty(aclState.IsEmpty())
	if err != nil {
		return
	}
	a.lastIndexed = acl.Head().Id
	return
}

func (a *aclObjectManager) processStates(states []list.AccountState, upToDate bool, myIdentity crypto.PubKey) (err error) {
	for _, state := range states {
		if state.Permissions.NoPermissions() && state.PubKey.Equals(myIdentity) && upToDate {
			return a.status.SetPersistentStatus(spaceinfo.AccountStatusRemoving)
		}
		err := a.participantWatcher.UpdateParticipantFromAclState(a.ctx, a.sp, state)
		if err != nil {
			return err
		}
		err = a.participantWatcher.RegisterIdentity(a.ctx, a.sp, state)
		if err != nil {
			return err
		}
	}
	return nil
}

func decryptAll(states []list.AccountState, decrypt func(key crypto.PubKey) ([]byte, error)) (decrypted []list.AccountState, err error) {
	for _, state := range states {
		res, err := decrypt(state.PubKey)
		if err != nil {
			return nil, err
		}
		state.RequestMetadata = res
		decrypted = append(decrypted, state)
	}
	return
}
