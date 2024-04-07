package spacestatus

import (
	"context"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/debugstat"

	"github.com/anyproto/anytype-heart/space/spaceinfo"
	"github.com/anyproto/anytype-heart/space/techspace"
)

const CName = "client.components.spacestatus"

type SpaceStatus interface {
	app.ComponentRunnable
	SpaceId() string
	GetLocalStatus() spaceinfo.LocalStatus
	GetPersistentStatus() spaceinfo.AccountStatus
	GetLatestAclHeadId() string
	SetPersistentStatus(status spaceinfo.AccountStatus) (err error)
	SetPersistentInfo(info spaceinfo.SpacePersistentInfo) (err error)
	SetLocalStatus(status spaceinfo.LocalStatus) error
	SetLocalInfo(info spaceinfo.SpaceLocalInfo) (err error)
	SetAccessType(status spaceinfo.AccessType) (err error)
	SetAclIsEmpty(isEmpty bool) (err error)
}

type spaceStatus struct {
	spaceId     string
	techSpace   techspace.TechSpace
	spaceView   techspace.SpaceView
	statService debugstat.StatService
}

func (s *spaceStatus) ProvideStat() any {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	localInfo := s.spaceView.GetLocalInfo()
	persistentInfo := s.spaceView.GetPersistentInfo()
	return spaceStatusStat{
		SpaceId:       s.spaceId,
		AccountStatus: persistentInfo.GetAccountStatus().String(),
		AclHeadId:     persistentInfo.GetAclHeadId(),
		LocalStatus:   localInfo.GetLocalStatus().String(),
		RemoteStatus:  localInfo.GetRemoteStatus().String(),
	}
}

func (s *spaceStatus) StatId() string {
	return s.spaceId
}

func (s *spaceStatus) StatType() string {
	return CName
}

func New(spaceId string) SpaceStatus {
	return &spaceStatus{
		spaceId: spaceId,
	}
}

func (s *spaceStatus) Init(a *app.App) (err error) {
	s.techSpace = a.MustComponent(techspace.CName).(techspace.TechSpace)
	s.spaceView, err = s.techSpace.GetSpaceView(context.Background(), s.spaceId)
	if err != nil {
		return err
	}
	s.statService, _ = a.Component(debugstat.CName).(debugstat.StatService)
	if s.statService == nil {
		s.statService = debugstat.NewNoOp()
	}
	s.statService.AddProvider(s)
	return nil
}

func (s *spaceStatus) Run(ctx context.Context) (err error) {
	return nil
}

func (s *spaceStatus) Close(ctx context.Context) (err error) {
	s.statService.RemoveProvider(s)
	return nil
}

func (s *spaceStatus) Name() (name string) {
	return CName
}

func (s *spaceStatus) SpaceId() string {
	return s.spaceId
}

func (s *spaceStatus) GetLocalStatus() spaceinfo.LocalStatus {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	info := s.spaceView.GetLocalInfo()
	return info.GetLocalStatus()
}

func (s *spaceStatus) GetPersistentStatus() spaceinfo.AccountStatus {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	info := s.spaceView.GetPersistentInfo()
	return info.GetAccountStatus()
}

func (s *spaceStatus) GetLatestAclHeadId() string {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	info := s.spaceView.GetPersistentInfo()
	return info.GetAclHeadId()
}

func (s *spaceStatus) SetLocalStatus(status spaceinfo.LocalStatus) error {
	info := spaceinfo.NewSpaceLocalInfo(s.spaceId)
	info.SetLocalStatus(status)
	return s.SetLocalInfo(info)
}

func (s *spaceStatus) SetLocalInfo(info spaceinfo.SpaceLocalInfo) (err error) {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	return s.spaceView.SetSpaceLocalInfo(info)
}

func (s *spaceStatus) SetPersistentInfo(info spaceinfo.SpacePersistentInfo) (err error) {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	return s.spaceView.SetSpacePersistentInfo(info)
}

func (s *spaceStatus) SetPersistentStatus(status spaceinfo.AccountStatus) (err error) {
	info := spaceinfo.NewSpacePersistentInfo(s.spaceId)
	info.SetAccountStatus(status)
	return s.spaceView.SetSpacePersistentInfo(info)
}

func (s *spaceStatus) SetAccessType(acc spaceinfo.AccessType) (err error) {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	return s.spaceView.SetAccessType(acc)
}

func (s *spaceStatus) SetAclIsEmpty(isEmpty bool) (err error) {
	s.spaceView.Lock()
	defer s.spaceView.Unlock()
	return s.spaceView.SetAclIsEmpty(isEmpty)
}

type spaceStatusStat struct {
	SpaceId       string `json:"spaceId"`
	AccountStatus string `json:"accountStatus"`
	LocalStatus   string `json:"localStatus"`
	RemoteStatus  string `json:"remoteStatus"`
	AclHeadId     string `json:"aclHeadId"`
}
