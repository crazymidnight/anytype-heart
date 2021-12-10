package subscription

import (
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
)

func (s *service) newSimpleSub(id string, keys []string, isDep bool) *simpleSub {
	sub := &simpleSub{
		id:    id,
		keys:  keys,
		cache: s.cache,
	}
	if !isDep {
		sub.ds = s.ds
	}
	return sub
}

type simpleSub struct {
	id   string
	set  map[string]struct{}
	keys []string

	depKeys          []string
	depSub           *simpleSub
	activeEntriesBuf []*entry

	cache *cache
	ds    *dependencyService
}

func (s *simpleSub) init(entries []*entry) (err error) {
	s.set = make(map[string]struct{})
	for _, e := range entries {
		e = s.cache.getOrSet(e)
		s.set[e.id] = struct{}{}
	}
	if s.ds != nil {
		s.depKeys = s.ds.depKeys(s.keys)
		if len(s.depKeys) > 0 {
			s.depSub = s.ds.makeSubscriptionByEntries(s.id+"/dep", s.getActiveEntries(), s.keys, s.depKeys)
		}
	}
	return
}

func (s *simpleSub) refill(ctx *opCtx, entries []*entry) {
	var newSet = make(map[string]struct{})
	for _, e := range entries {
		if _, inSet := s.set[e.id]; inSet {
			ctx.change = append(ctx.change, opChange{
				id:    e.id,
				subId: s.id,
				keys:  s.keys,
			})
		} else {
			ctx.add = append(ctx.add, opChange{
				id:    e.id,
				subId: s.id,
				keys:  s.keys,
			})
			e.refs++
		}
		newSet[e.id] = struct{}{}
	}
	for oldId := range s.set {
		if _, inSet := newSet[oldId]; !inSet {
			ctx.remove = append(ctx.remove, opRemove{
				id:    oldId,
				subId: s.id,
			})
			s.cache.release(oldId)
		}
	}
	s.set = newSet
}

func (s *simpleSub) counters() (prev, next int) {
	return 0, 0
}

func (s *simpleSub) onChangeBatch(ctx *opCtx, entries ...*entry) {
	var changed bool
	for _, e := range entries {
		if _, inSet := s.set[e.id]; inSet {
			ctx.change = append(ctx.change, opChange{
				id:    e.id,
				subId: s.id,
				keys:  s.keys,
			})
			changed = true
		}
	}
	if changed && s.depSub != nil {
		s.ds.refillSubscription(ctx, s.depSub, s.getActiveEntries(), s.depKeys)
	}
}

func (s *simpleSub) getActiveEntries() (res []*entry) {
	s.activeEntriesBuf = s.activeEntriesBuf[:0]
	for id := range s.set {
		res = append(res, s.cache.pick(id))
	}
	return s.activeEntriesBuf
}

func (s *simpleSub) getActiveRecords() (res []*types.Struct) {
	for id := range s.set {
		res = append(res, pbtypes.StructFilterKeys(s.cache.pick(id).data, s.keys))
	}
	return
}

func (s *simpleSub) close() {
	for id := range s.set {
		s.cache.release(id)
	}
	if s.depSub != nil {
		s.depSub.close()
	}
	return
}
