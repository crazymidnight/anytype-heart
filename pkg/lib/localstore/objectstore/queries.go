package objectstore

import (
	"fmt"
	"sort"
	"strings"

	"github.com/blevesearch/bleve/v2/search"
	"github.com/dgraph-io/badger/v4"

	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/space/spacecore/typeprovider"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func (s *dsObjectStore) Query(q database.Query) ([]database.Record, int, error) {
	filters, err := s.buildQuery(q)
	if err != nil {
		return nil, 0, fmt.Errorf("build query: %w", err)
	}
	recs, err := s.QueryRaw(filters, q.Limit, q.Offset)
	return recs, 0, err
}

func (s *dsObjectStore) QueryRaw(filters *database.Filters, limit int, offset int) ([]database.Record, error) {
	if filters == nil || filters.FilterObj == nil {
		return nil, fmt.Errorf("filter cannot be nil or unitialized")
	}

	var records []database.Record
	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		opts.Prefix = pagesDetailsBase.Bytes()
		iterator := txn.NewIterator(opts)
		defer iterator.Close()

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			it := iterator.Item()
			details, err := s.extractDetailsFromItem(it)
			if err != nil {
				return err
			}
			id := pbtypes.GetString(details.Details, bundle.RelationKeyId.String())
			name := pbtypes.GetString(details.Details, bundle.RelationKeyName.String())
			rec := database.Record{Details: details.Details}

			if innerId, ok := filters.ObjectInnerId[id]; ok && details.Details != nil {
				detailsCopy := pbtypes.CopyStruct(details.Details)
				if strings.HasPrefix(innerId, "r_") {
					detailsCopy.Fields[bundle.RelationKeySearchTargetRelation.String()] = pbtypes.String(innerId[2:])
					detailsCopy.Fields[bundle.RelationKeyName.String()] = pbtypes.String(fmt.Sprintf("%s (rel %s)", name, innerId[2:]))
				} else {
					detailsCopy.Fields[bundle.RelationKeySearchTargetBlock.String()] = pbtypes.String(innerId)
					detailsCopy.Fields[bundle.RelationKeyName.String()] = pbtypes.String(fmt.Sprintf("%s (block %s)", name, innerId))
				}
				rec.Details = detailsCopy
			}

			if filters.FilterObj != nil && filters.FilterObj.FilterObject(rec) {
				records = append(records, rec)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if offset >= len(records) {
		return nil, nil
	}
	if filters.Order != nil {
		sort.Slice(records, func(i, j int) bool {
			return filters.Order.Compare(records[i], records[j]) == -1
		})
	}
	if limit > 0 {
		upperBound := offset + limit
		if upperBound > len(records) {
			upperBound = len(records)
		}
		return records[offset:upperBound], nil
	}
	return records[offset:], nil
}

func (s *dsObjectStore) buildQuery(q database.Query) (*database.Filters, error) {
	filters, err := database.NewFilters(q, s)
	if err != nil {
		return nil, fmt.Errorf("new filters: %w", err)
	}

	if q.FullText != "" {
		filters, err = s.makeFTSQuery(q.FullText, filters)
		if err != nil {
			return nil, fmt.Errorf("append full text search query: %w", err)
		}
	}
	return filters, nil
}

func (s *dsObjectStore) makeFTSQuery(text string, filters *database.Filters) (*database.Filters, error) {
	if s.fts == nil {
		return filters, fmt.Errorf("fullText search not configured")
	}
	results, err := s.fts.Search(getSpaceIDFromFilter(filters.FilterObj), text)
	if filters.ObjectInnerId == nil {
		filters.ObjectInnerId = make(map[string]string)
	}
	var resultsByObjectId = make(map[string][]*search.DocumentMatch)
	for _, result := range results {
		objectId, _, _ := domain.ExtractFromFullTextId(result.ID)
		if _, ok := resultsByObjectId[objectId]; !ok {
			resultsByObjectId[objectId] = make([]*search.DocumentMatch, 0, 1)
		}

		resultsByObjectId[objectId] = append(resultsByObjectId[objectId], result)
	}
	for objectId := range resultsByObjectId {
		sort.Slice(resultsByObjectId[objectId], func(i, j int) bool {
			return results[i].Score > results[j].Score
		})
	}

	var objectIds = make([]string, 0, len(resultsByObjectId))
	for objectId, results := range resultsByObjectId {
		if len(results) == 0 {
			continue
		}
		_, blockId, relationKey := domain.ExtractFromFullTextId(results[0].ID)

		if blockId != "" {
			filters.ObjectInnerId[objectId] = blockId
		} else if relationKey != "" {
			filters.ObjectInnerId[objectId] = "r_" + relationKey
		}
		objectIds = append(objectIds, objectId)
	}
	if err != nil {
		return filters, err
	}
	idsQuery := newIdsFilter(objectIds)
	filters.FilterObj = database.FiltersAnd{filters.FilterObj, idsQuery}
	filters.Order = database.SetOrder(append([]database.Order{idsQuery}, filters.Order))
	return filters, nil
}

func getSpaceIDFromFilter(fltr database.Filter) (spaceID string) {
	switch f := fltr.(type) {
	case database.FilterEq:
		if f.Key == bundle.RelationKeySpaceId.String() {
			return f.Value.GetStringValue()
		}
	case database.FiltersAnd:
		spaceID = iterateOverAndFilters(f)
	}
	return spaceID
}

func iterateOverAndFilters(fs []database.Filter) (spaceID string) {
	for _, f := range fs {
		if spaceID = getSpaceIDFromFilter(f); spaceID != "" {
			return spaceID
		}
	}
	return ""
}

// TODO: objstore: no one uses total
func (s *dsObjectStore) QueryObjectIDs(q database.Query) (ids []string, total int, err error) {
	filters, err := s.buildQuery(q)
	if err != nil {
		return nil, 0, fmt.Errorf("build query: %w", err)
	}
	recs, err := s.QueryRaw(filters, q.Limit, q.Offset)
	if err != nil {
		return nil, 0, fmt.Errorf("query raw: %w", err)
	}
	ids = make([]string, 0, len(recs))
	for _, rec := range recs {
		ids = append(ids, pbtypes.GetString(rec.Details, bundle.RelationKeyId.String()))
	}
	return ids, 0, nil
}

func (s *dsObjectStore) QueryByID(ids []string) (records []database.Record, err error) {
	err = s.db.View(func(txn *badger.Txn) error {
		for _, id := range ids {
			// Don't use spaceID because expected objects are virtual
			if sbt, err := typeprovider.SmartblockTypeFromID(id); err == nil {
				if indexDetails, _ := sbt.Indexable(); !indexDetails && s.sourceService != nil {
					details, err := s.sourceService.DetailsFromIdBasedSource(id)
					if err != nil {
						log.Errorf("QueryByIds failed to GetDetailsFromIdBasedSource id: %s", id)
						continue
					}
					details.Fields[database.RecordIDField] = pbtypes.ToValue(id)
					records = append(records, database.Record{Details: details})
					continue
				}
			}
			it, err := txn.Get(pagesDetailsBase.ChildString(id).Bytes())
			if err != nil {
				log.Infof("QueryByIds failed to find id: %s", id)
				continue
			}

			details, err := s.extractDetailsFromItem(it)
			if err != nil {
				log.Errorf("QueryByIds failed to extract details: %s", id)
				continue
			}
			records = append(records, database.Record{Details: details.Details})
		}
		return nil
	})
	return
}

func (s *dsObjectStore) QueryByIDAndSubscribeForChanges(ids []string, sub database.Subscription) (records []database.Record, closeFunc func(), err error) {
	s.Lock()
	defer s.Unlock()

	if sub == nil {
		err = fmt.Errorf("subscription func is nil")
		return
	}
	sub.Subscribe(ids)
	records, err = s.QueryByID(ids)
	if err != nil {
		// can mean only the datastore is already closed, so we can resign and return
		log.Errorf("QueryByIDAndSubscribeForChanges failed to query ids: %v", err)
		return nil, nil, err
	}

	closeFunc = func() {
		s.closeAndRemoveSubscription(sub)
	}

	s.addSubscriptionIfNotExists(sub)
	return
}
