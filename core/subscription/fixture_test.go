package subscription

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/anyproto/any-sync/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/anyproto/anytype-heart/core/event"
	"github.com/anyproto/anytype-heart/core/event/mock_event"
	"github.com/anyproto/anytype-heart/core/subscription/mock_subscription"
	"github.com/anyproto/anytype-heart/core/system_object/mock_system_object"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/space/typeprovider/mock_typeprovider"
	"github.com/anyproto/anytype-heart/tests/testutil"
	"github.com/anyproto/anytype-heart/util/testMock"
)

type collectionServiceMock struct {
	*mock_subscription.MockCollectionService
}

func (c *collectionServiceMock) Name() string {
	return "collectionService"
}

func (c *collectionServiceMock) Init(a *app.App) error { return nil }

type fixture struct {
	Service
	a                   *app.App
	ctrl                *gomock.Controller
	store               *testMock.MockObjectStore
	systemObjectService *mock_system_object.MockService
	sender              *mock_event.MockSender
	events              []*pb.Event
	collectionService   *collectionServiceMock
}

func newFixture(t *testing.T) *fixture {
	ctrl := gomock.NewController(t)
	a := new(app.App)
	testMock.RegisterMockObjectStore(ctrl, a)
	testMock.RegisterMockKanban(ctrl, a)
	sbtProvider := mock_typeprovider.NewMockSmartBlockTypeProvider(t)
	sbtProvider.EXPECT().Name().Return("smartBlockTypeProvider")
	sbtProvider.EXPECT().Init(mock.Anything).Return(nil)
	a.Register(sbtProvider)

	systemObjectService := mock_system_object.NewMockService(t)
	a.Register(testutil.PrepareMock(a, systemObjectService))

	collectionService := &collectionServiceMock{MockCollectionService: mock_subscription.NewMockCollectionService(t)}
	a.Register(collectionService)

	fx := &fixture{
		Service:             New(),
		a:                   a,
		ctrl:                ctrl,
		store:               a.MustComponent(objectstore.CName).(*testMock.MockObjectStore),
		systemObjectService: systemObjectService,
		collectionService:   collectionService,
	}
	sender := mock_event.NewMockSender(t)
	sender.EXPECT().Init(mock.Anything).Return(nil)
	sender.EXPECT().Name().Return(event.CName)
	sender.EXPECT().Broadcast(mock.Anything).Run(func(e *pb.Event) {
		fx.events = append(fx.events, e)
	}).Maybe()
	fx.sender = sender
	a.Register(fx.Service)
	a.Register(fx.sender)

	fx.store.EXPECT().SubscribeForAll(gomock.Any())
	require.NoError(t, a.Start(context.Background()))
	return fx
}

type fixtureRealStore struct {
	Service
	a                       *app.App
	ctrl                    *gomock.Controller
	store                   *objectstore.StoreFixture
	sender                  *mock_event.MockSender
	eventsLock              sync.Mutex
	events                  []pb.IsEventMessageValue
	systemObjectServiceMock *mock_system_object.MockService
}

func newFixtureWithRealObjectStore(t *testing.T) *fixtureRealStore {
	ctrl := gomock.NewController(t)
	a := new(app.App)
	store := objectstore.NewStoreFixture(t)
	a.Register(store)
	testMock.RegisterMockKanban(ctrl, a)
	a.Register(&collectionServiceMock{})
	sbtProvider := mock_typeprovider.NewMockSmartBlockTypeProvider(t)
	sbtProvider.EXPECT().Name().Return("smartBlockTypeProvider")
	sbtProvider.EXPECT().Init(mock.Anything).Return(nil)
	a.Register(sbtProvider)
	systemObjectService := mock_system_object.NewMockService(t)
	a.Register(testutil.PrepareMock(a, systemObjectService))
	fx := &fixtureRealStore{
		Service:                 New(),
		a:                       a,
		ctrl:                    ctrl,
		store:                   store,
		systemObjectServiceMock: systemObjectService,
	}
	sender := mock_event.NewMockSender(t)
	sender.EXPECT().Init(mock.Anything).Return(nil)
	sender.EXPECT().Name().Return(event.CName)
	sender.EXPECT().Broadcast(mock.Anything).Run(func(e *pb.Event) {
		fx.eventsLock.Lock()
		defer fx.eventsLock.Unlock()
		for _, em := range e.Messages {
			fx.events = append(fx.events, em.Value)
		}
	}).Maybe()
	fx.sender = sender
	a.Register(fx.Service)
	a.Register(fx.sender)

	require.NoError(t, a.Start(context.Background()))
	return fx
}

func (fx *fixtureRealStore) waitEvents(t *testing.T, ev ...pb.IsEventMessageValue) {
	timeout := time.NewTimer(1 * time.Second)
	ticker := time.NewTicker(1 * time.Millisecond)
	for {
		select {
		case <-timeout.C:
			fx.eventsLock.Lock()
			assert.Equal(t, ev, fx.events)
			fx.eventsLock.Unlock()
			return
		case <-ticker.C:
		}

		fx.eventsLock.Lock()
		if reflect.DeepEqual(fx.events, ev) {
			fx.events = nil
			fx.eventsLock.Unlock()
			return
		}
		fx.eventsLock.Unlock()
	}
}
