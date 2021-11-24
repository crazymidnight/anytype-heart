package subscription

import (
	"fmt"
	"sync"

	"github.com/anytypeio/go-anytype-middleware/app"
	"github.com/anytypeio/go-anytype-middleware/core/event"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database/filter"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/cheggaaa/mb"
	"github.com/globalsign/mgo/bson"
	"github.com/gogo/protobuf/types"
	"github.com/ipfs/go-datastore/query"
)

const CName = "subscription"

var log = logging.Logger("anytype-mw-subscription")

func New() Service {
	return new(service)
}

type Service interface {
	Search(req pb.RpcObjectSearchSubscribeRequest) (resp *pb.RpcObjectSearchSubscribeResponse, err error)
	SubscribeIdsReq(req pb.RpcObjectIdsSubscribeRequest) (resp *pb.RpcObjectIdsSubscribeResponse, err error)
	SubscribeIds(subId string, ids []string) (records []*types.Struct, err error)
	Unsubscribe(subIds ...string) (err error)
	UnsubscribeAll() (err error)

	app.ComponentRunnable
}

type subscription interface {
	init(entries []*entry) (err error)
	counters() (prev, next int)
	onChangeBatch(ctx *opCtx, entries ...*entry)
	getActiveRecords() (res []*types.Struct)
	close()
}

type service struct {
	cache         *cache
	ds            *dependencyService
	subscriptions map[string]subscription
	recBatch      *mb.MB

	objectStore objectstore.ObjectStore
	sendEvent   func(e *pb.Event)

	m      sync.Mutex
	ctxBuf *opCtx
}

func (s *service) Init(a *app.App) (err error) {
	s.cache = newCache()
	s.ds = newDependencyService(s)
	s.subscriptions = make(map[string]subscription)
	s.objectStore = a.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	s.recBatch = mb.New(0)
	s.sendEvent = a.MustComponent(event.CName).(event.Sender).Send
	s.ctxBuf = &opCtx{}
	return
}

func (s *service) Name() (name string) {
	return CName
}

func (s *service) Run() (err error) {
	s.objectStore.SubscribeForAll(func(rec database.Record) {
		s.recBatch.Add(rec)
	})
	go s.recordsHandler()
	return
}

func (s *service) Search(req pb.RpcObjectSearchSubscribeRequest) (resp *pb.RpcObjectSearchSubscribeResponse, err error) {
	if req.SubId == "" {
		req.SubId = bson.NewObjectId().Hex()
	}

	q := database.Query{
		Filters:  req.Filters,
		Sorts:    req.Sorts,
		Limit:    int(req.Limit),
		FullText: req.FullText,
	}

	f, err := database.NewFilters(q, nil)
	if err != nil {
		return
	}

	if len(req.Source) > 0 {
		sourceFilter, err := s.filtersFromSource(req.Source)
		if err != nil {
			return nil, err
		}
		f.FilterObj = filter.AndFilters{f.FilterObj, sourceFilter}
	}

	records, err := s.objectStore.QueryRaw(query.Query{
		Filters: []query.Filter{f},
	})
	if err != nil {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()
	if exists, ok := s.subscriptions[req.SubId]; ok {
		exists.close()
	}
	if req.Offset < 0 {
		req.Offset = 0
	}
	if req.Limit < 0 {
		req.Limit = 0
	}
	sub := s.newSortedSub(req.SubId, req.Keys, f.FilterObj, f.Order, int(req.Limit), int(req.Offset))
	entries := make([]*entry, 0, len(records))
	for _, r := range records {
		entries = append(entries, &entry{
			id:   pbtypes.GetString(r.Details, "id"),
			data: r.Details,
		})
	}
	if err = sub.init(entries); err != nil {
		return
	}
	s.subscriptions[sub.id] = sub
	prev, next := sub.counters()

	var depRecords, subRecords []*types.Struct
	subRecords = sub.getActiveRecords()

	if sub.depSub != nil {
		depRecords = sub.depSub.getActiveRecords()
	}

	resp = &pb.RpcObjectSearchSubscribeResponse{
		Records:      subRecords,
		Dependencies: depRecords,
		SubId:        sub.id,
		Counters: &pb.EventObjectSubscriptionCounters{
			Total:     int64(sub.skl.Len()),
			NextCount: int64(prev),
			PrevCount: int64(next),
		},
	}
	return
}

func (s *service) SubscribeIdsReq(req pb.RpcObjectIdsSubscribeRequest) (resp *pb.RpcObjectIdsSubscribeResponse, err error) {
	return &pb.RpcObjectIdsSubscribeResponse{
		Error:        &pb.RpcObjectIdsSubscribeResponseError{},
		Records:      nil,
		Dependencies: nil,
		SubId:        req.SubId,
	}, nil
}

func (s *service) SubscribeIds(subId string, ids []string) (records []*types.Struct, err error) {
	return
}

func (s *service) Unsubscribe(subIds ...string) (err error) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, subId := range subIds {
		if sub, ok := s.subscriptions[subId]; ok {
			sub.close()
			delete(s.subscriptions, subId)
		}
	}
	return
}

func (s *service) UnsubscribeAll() (err error) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, sub := range s.subscriptions {
		sub.close()
	}
	s.subscriptions = make(map[string]subscription)
	return
}

func (s *service) recordsHandler() {
	var entries []*entry
	for {
		records := s.recBatch.Wait()
		if len(records) == 0 {
			return
		}
		for _, rec := range records {
			entries = append(entries, &entry{
				id:   pbtypes.GetString(rec.(database.Record).Details, "id"),
				data: rec.(database.Record).Details,
			})
		}
		s.onChange(entries)
		entries = entries[:0]
	}
}

func (s *service) onChange(entries []*entry) {
	s.m.Lock()
	defer s.m.Unlock()
	s.ctxBuf.reset()
	for _, sub := range s.subscriptions {
		sub.onChangeBatch(s.ctxBuf, entries...)
	}
	log.Debugf("handle %d etries; ctx: %#v", len(entries), s.ctxBuf)
	events := s.ctxBuf.apply(s.cache, entries)
	for _, e := range events {
		s.sendEvent(e)
	}
}

func (s *service) filtersFromSource(sources []string) (filter.Filter, error) {
	var objTypeIds, relTypeKeys []string

	for _, source := range sources {
		sbt, err := smartblock.SmartBlockTypeFromID(source)
		if err != nil {
			return nil, err
		}
		if sbt == smartblock.SmartBlockTypeObjectType || sbt == smartblock.SmartBlockTypeBundledObjectType {
			objTypeIds = append(objTypeIds, source)
		} else {
			relKey, err := pbtypes.RelationIdToKey(source)
			if err != nil {
				return nil, fmt.Errorf("failed to get relation key from id %s: %s", relKey, err.Error())
			}
			relTypeKeys = append(relTypeKeys, relKey)
		}
	}

	var relTypeFilter filter.OrFilters

	if len(objTypeIds) > 0 {
		relTypeFilter = append(relTypeFilter, filter.In{
			Key:   bundle.RelationKeyType.String(),
			Value: pbtypes.StringList(objTypeIds).GetListValue(),
		})
	}

	for _, key := range relTypeKeys {
		relTypeFilter = append(relTypeFilter, filter.Exists{
			Key: key,
		})
	}
	return relTypeFilter, nil
}

func (s *service) Close() (err error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.recBatch.Close()
	for _, sub := range s.subscriptions {
		sub.close()
	}
	return
}
