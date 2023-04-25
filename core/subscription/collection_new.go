package subscription

import (
	"fmt"
	"sync"

	"github.com/cheggaaa/mb"
	"github.com/gogo/protobuf/types"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database/filter"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

type collectionObserver struct {
	lock   *sync.RWMutex
	ids    []string
	idsMap map[string]struct{}

	cache       *cache
	objectStore objectstore.ObjectStore
	recBatch    *mb.MB
}

func (s *service) newCollectionObserver(collectionID string) (*collectionObserver, error) {
	initialObjectIDs, changesCh, err := s.collections.SubscribeForCollection(collectionID)
	if err != nil {
		return nil, fmt.Errorf("subscribe for collection: %w", err)
	}

	idsMap := map[string]struct{}{}
	for _, id := range initialObjectIDs {
		idsMap[id] = struct{}{}
	}

	obs := &collectionObserver{
		lock:   &sync.RWMutex{},
		ids:    initialObjectIDs,
		idsMap: idsMap,

		cache:       s.cache,
		objectStore: s.objectStore,
		recBatch:    s.recBatch,
	}

	go func() {
		for chs := range changesCh {
			obs.applyChanges(chs)
		}
	}()

	return obs, nil
}

func (c *collectionObserver) listEntries() []*entry {
	c.lock.RLock()
	defer c.lock.RUnlock()
	entries := fetchEntries(c.cache, c.objectStore, c.ids)
	res := make([]*entry, len(entries))
	copy(res, entries)
	return res
}

func (c *collectionObserver) applyChanges(changes []slice.Change[string]) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.ids = slice.ApplyChanges(c.ids, changes, slice.StringIdentity[string])

	changedIDs := map[string]struct{}{}

	var isMoved bool
	for _, ch := range changes {
		if add := ch.Add(); add != nil {
			for _, id := range add.Items {
				c.idsMap[id] = struct{}{}
				changedIDs[id] = struct{}{}
			}
		}

		if rm := ch.Remove(); rm != nil {
			for _, id := range rm.IDs {
				delete(c.idsMap, id)
				changedIDs[id] = struct{}{}
			}
		}

		if mv := ch.Move(); mv != nil {
			isMoved = true
		}
	}

	var reqIDs []string
	if isMoved {
		reqIDs = c.ids
	} else {
		reqIDs = make([]string, 0, len(changedIDs))
		for id := range changedIDs {
			reqIDs = append(reqIDs, id)
		}
	}

	entries := fetchEntries(c.cache, c.objectStore, reqIDs)
	for _, e := range entries {
		c.recBatch.Add(database.Record{
			Details: e.data,
		})
	}

}

func (c *collectionObserver) Compare(a, b filter.Getter) int {
	c.lock.RLock()
	defer c.lock.RUnlock()

	ae, be := a.(*entry), b.(*entry)
	ap, bp := slice.FindPos(c.ids, ae.id), slice.FindPos(c.ids, be.id)
	if ap == bp {
		return 0
	}
	if ap < bp {
		return -1
	}
	return 1
}

func (c *collectionObserver) FilterObject(g filter.Getter) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.idsMap[g.(*entry).id]
	return ok
}

func (c *collectionObserver) String() string {
	return "collectionObserver"
}

type collectionSubscription struct {
	sortedSub   *sortedSub
	cache       *cache
	objectStore objectstore.ObjectStore
}

func (c *collectionSubscription) init(entries []*entry) (err error) {
	return nil
}

func (c *collectionSubscription) counters() (prev, next int) {
	return c.sortedSub.counters()
}

func (c *collectionSubscription) onChange(ctx *opCtx) {
	c.sortedSub.onChange(ctx)
}

func (c *collectionSubscription) getActiveRecords() (res []*types.Struct) {
	return c.sortedSub.getActiveRecords()
}

func (c *collectionSubscription) hasDep() bool {
	return c.sortedSub.hasDep()
}

func (c *collectionSubscription) close() {
	// TODO close observer
	c.sortedSub.close()
}

func (s *service) newCollectionSubscription(id string, collectionID string, keys []string, flt filter.Filter, order filter.Order, limit, offset int) (*collectionSubscription, error) {
	obs, err := s.newCollectionObserver(collectionID)
	if err != nil {
		return nil, err
	}
	flt = filter.AndFilters{flt, obs}

	ssub := s.newSortedSub(id, keys, flt, obs, limit, offset)
	// TODO set to true only if it's no user orders
	ssub.batchUpdate = true
	sub := &collectionSubscription{
		sortedSub:   ssub,
		cache:       s.cache,
		objectStore: s.objectStore,
	}

	if err := ssub.init(obs.listEntries()); err != nil {
		return nil, err
	}
	return sub, nil
}

func fetchEntries(cache *cache, objectStore objectstore.ObjectStore, ids []string) []*entry {
	res := make([]*entry, 0, len(ids))
	for _, id := range ids {
		if e := cache.Get(id); e != nil {
			res = append(res, e)
			continue
		}
		// TODO query in one batch
		recs, err := objectStore.QueryById([]string{id})
		if err != nil {
			// TODO proper logging
			fmt.Println("query new entry:", err)
		}
		if len(recs) > 0 {
			e := &entry{
				id:   id,
				data: recs[0].Details,
			}
			res = append(res, e)
		}
	}
	return res
}
