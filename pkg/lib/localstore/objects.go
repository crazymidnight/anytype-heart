package localstore

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	pbrelation "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/schema"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/structs"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

var (
	// ObjectInfo is stored in db key pattern:
	pagesPrefix        = "pages"
	pagesDetailsBase   = ds.NewKey("/" + pagesPrefix + "/details")
	pagesRelationsBase = ds.NewKey("/" + pagesPrefix + "/relations")

	pagesSnippetBase       = ds.NewKey("/" + pagesPrefix + "/snippet")
	pagesInboundLinksBase  = ds.NewKey("/" + pagesPrefix + "/inbound")
	pagesOutboundLinksBase = ds.NewKey("/" + pagesPrefix + "/outbound")

	_ ObjectStore = (*dsObjectStore)(nil)
)

var ErrNotAPage = fmt.Errorf("not a page")

const (
	// special record fields
	fieldLastOpened   = "lastOpenedDate"
	fieldLastModified = "lastModifiedDate"

	pageSchema = "https://anytype.io/schemas/page"
)

type filterNotSystemObjects struct{}

func (m *filterNotSystemObjects) Filter(e query.Entry) bool {
	keyParts := strings.Split(e.Key, "/")
	id := keyParts[len(keyParts)-1]

	t, err := smartblock.SmartBlockTypeFromID(id)
	if err != nil {
		log.Errorf("failed to detect smartblock type for %s: %s", id, err.Error())
		return false
	}

	if t != smartblock.SmartBlockTypeArchive && t != smartblock.SmartBlockTypeHome && t != smartblock.SmartBlockTypeObjectType {
		return true
	}

	return false
}

func NewObjectStore(ds ds.TxnDatastore) ObjectStore {
	return &dsObjectStore{ds: ds}
}

type subscription struct {
	ids  []string
	quit chan struct{}
	ch   chan *types.Struct
}

type dsObjectStore struct {
	// underlying storage
	ds ds.TxnDatastore

	// serializing page updates
	l sync.Mutex

	subscriptions []*subscription
}

func (m *dsObjectStore) QueryAndSubscribeForChanges(ctx context.Context, schema *schema.Schema, q database.Query, updatedRecordsCh chan database.Record) (records []database.Record, total int, err error) {
	m.l.Lock()
	defer m.l.Unlock()

	records, total, err = m.Query(schema, q)
	ch := make(chan *types.Struct)
	quitCh := make(chan struct{})

	var ids []string
	for _, record := range records {
		ids = append(ids, pbtypes.GetString(record.Details, "id"))
	}

	sub := subscription{ids: ids, quit: quitCh, ch: ch}
	m.subscriptions = append(m.subscriptions, &sub)

	go func() {
		defer func() {
			close(sub.quit)
			close(updatedRecordsCh)
			m.l.Lock()
			for i, s := range m.subscriptions {
				if s.ch == sub.ch {
					m.subscriptions = append(m.subscriptions[:i], m.subscriptions[i+1:]...)
				}
			}
			m.l.Unlock()
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case details, ok := <-ch:
				if !ok {
					return
				}

				updatedRecordsCh <- database.Record{Details: details}
			}
		}
	}()

	return
}

func (m *dsObjectStore) Query(sch *schema.Schema, q database.Query) (records []database.Record, total int, err error) {
	txn, err := m.ds.NewTransaction(true)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	dsq := q.DSQuery(sch)
	dsq.Offset = 0
	dsq.Limit = 0
	dsq.Prefix = pagesDetailsBase.String() + "/"
	dsq.Filters = append([]query.Filter{&filterNotSystemObjects{}}, dsq.Filters...)
	res, err := txn.Query(dsq)
	if err != nil {
		return nil, 0, fmt.Errorf("error when querying ds: %w", err)
	}

	var (
		results []database.Record
		offset  = q.Offset
	)

	// We use own limit/offset implementation in order to find out
	// total number of records matching specified filters. Query
	// returns this number for handy pagination on clients.
	for rec := range res.Next() {
		total++

		if offset > 0 {
			offset--
			continue
		}

		if q.Limit > 0 && len(results) >= q.Limit {
			continue
		}

		var details model.ObjectDetails
		if err = proto.Unmarshal(rec.Value, &details); err != nil {
			log.Errorf("failed to unmarshal: %s", err.Error())
			total--
			continue
		}

		key := ds.NewKey(rec.Key)
		keyList := key.List()
		id := keyList[len(keyList)-1]

		if details.Details == nil || details.Details.Fields == nil {
			details.Details = &types.Struct{Fields: map[string]*types.Value{}}
		}

		details.Details.Fields[database.RecordIDField] = pb.ToValue(id)
		results = append(results, database.Record{Details: details.Details})
	}

	return results, total, nil
}

func (m *dsObjectStore) AggregateRelations(sch *schema.Schema) (relations []*pbrelation.Relation, err error) {
	txn, err := m.ds.NewTransaction(true)
	if err != nil {
		return nil, fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()
	q := database.Query{}
	dsq := q.DSQuery(sch)
	dsq.Offset = 0
	dsq.Limit = 0
	dsq.Prefix = pagesRelationsBase.String() + "/"
	dsq.Filters = append([]query.Filter{&filterNotSystemObjects{}}, dsq.Filters...)
	res, err := txn.Query(dsq)
	if err != nil {
		return nil, fmt.Errorf("error when querying ds: %w", err)
	}

	var relationsKeysMaps map[string]struct{}

	for rec := range res.Next() {
		var rels pbrelation.Relations
		if err = proto.Unmarshal(rec.Value, &rels); err != nil {
			log.Errorf("failed to unmarshal: %s", err.Error())
			continue
		}

		for i, rel := range rels.Relations {
			if _, exists := relationsKeysMaps[rel.Key]; exists {
				continue
			}

			relationsKeysMaps[rel.Key] = struct{}{}
			relations = append(relations, rels.Relations[i])
		}
	}

	return relations, nil
}

func (m *dsObjectStore) AddObject(page *model.ObjectInfoWithOutboundLinksIDs) error {
	txn, err := m.ds.NewTransaction(false)
	if err != nil {
		return fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	detailsKey := pagesDetailsBase.ChildString(page.Id)
	relationsKey := pagesRelationsBase.ChildString(page.Id)
	snippetKey := pagesSnippetBase.ChildString(page.Id)

	if exists, err := txn.Has(detailsKey); err != nil {
		return err
	} else if exists {
		return ErrDuplicateKey
	}

	page.Info.Details.Fields["type"] = pbtypes.StringList(page.Info.ObjectTypeUrls)
	b, err := proto.Marshal(page.Info.Details)
	if err != nil {
		return err
	}
	if err = txn.Put(detailsKey, b); err != nil {
		return err
	}

	b, err = proto.Marshal(page.Info.Relations)
	if err != nil {
		return err
	}

	if err = txn.Put(relationsKey, b); err != nil {
		return err
	}

	for _, key := range pageLinkKeys(page.Id, nil, page.OutboundLinks) {
		if err = txn.Put(key, nil); err != nil {
			return err
		}
	}

	if err = txn.Put(snippetKey, []byte(page.Info.Snippet)); err != nil {
		return err
	}

	return txn.Commit()
}

func (m *dsObjectStore) DeleteObject(id string) error {
	txn, err := m.ds.NewTransaction(false)
	if err != nil {
		return fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	for _, k := range []ds.Key{
		pagesDetailsBase.ChildString(id),
		pagesSnippetBase.ChildString(id),
	} {
		if err = txn.Delete(k); err != nil {
			return err
		}
	}

	inLinks, err := findInboundLinks(txn, id)
	if err != nil {
		return err
	}

	outLinks, err := findOutboundLinks(txn, id)
	if err != nil {
		return err
	}

	for _, k := range pageLinkKeys(id, inLinks, outLinks) {
		if err := txn.Delete(k); err != nil {
			return err
		}
	}

	return txn.Commit()
}

func (m *dsObjectStore) GetWithLinksInfoByID(id string) (*model.ObjectInfoWithLinks, error) {
	txn, err := m.ds.NewTransaction(true)
	if err != nil {
		return nil, fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	pages, err := getPagesInfo(txn, []string{id})
	if err != nil {
		return nil, err
	}

	if len(pages) == 0 {
		return nil, fmt.Errorf("page not found")
	}
	page := pages[0]

	inboundIds, err := findInboundLinks(txn, id)
	if err != nil {
		return nil, err
	}

	outboundsIds, err := findOutboundLinks(txn, id)
	if err != nil {
		return nil, err
	}

	inbound, err := getPagesInfo(txn, inboundIds)
	if err != nil {
		return nil, err
	}

	outbound, err := getPagesInfo(txn, outboundsIds)
	if err != nil {
		return nil, err
	}

	return &model.ObjectInfoWithLinks{
		Id:   id,
		Info: page,
		Links: &model.ObjectLinksInfo{
			Inbound:  inbound,
			Outbound: outbound,
		},
	}, nil
}

func (m *dsObjectStore) GetWithOutboundLinksInfoById(id string) (*model.ObjectInfoWithOutboundLinks, error) {
	txn, err := m.ds.NewTransaction(true)
	if err != nil {
		return nil, fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	pages, err := getPagesInfo(txn, []string{id})
	if err != nil {
		return nil, err
	}

	if len(pages) == 0 {
		return nil, fmt.Errorf("page not found")
	}
	page := pages[0]

	outboundsIds, err := findOutboundLinks(txn, id)
	if err != nil {
		return nil, err
	}

	outbound, err := getPagesInfo(txn, outboundsIds)
	if err != nil {
		return nil, err
	}

	return &model.ObjectInfoWithOutboundLinks{
		Info:          page,
		OutboundLinks: outbound,
	}, nil
}

func (m *dsObjectStore) GetDetails(id string) (*model.ObjectDetails, error) {
	txn, err := m.ds.NewTransaction(true)
	if err != nil {
		return nil, fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	return getDetails(txn, id)
}

func (m *dsObjectStore) List() ([]*model.ObjectInfo, error) {
	txn, err := m.ds.NewTransaction(true)
	if err != nil {
		return nil, fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	ids, err := findByPrefix(txn, pagesDetailsBase.String()+"/", 0)
	if err != nil {
		return nil, err
	}

	return getPagesInfo(txn, ids)
}

func (m *dsObjectStore) GetByIDs(ids ...string) ([]*model.ObjectInfo, error) {
	txn, err := m.ds.NewTransaction(true)
	if err != nil {
		return nil, fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	return getPagesInfo(txn, ids)
}

func diffSlices(a, b []string) (removed []string, added []string) {
	var amap = map[string]struct{}{}
	var bmap = map[string]struct{}{}

	for _, item := range a {
		amap[item] = struct{}{}
	}

	for _, item := range b {
		if _, exists := amap[item]; !exists {
			added = append(added, item)
		}
		bmap[item] = struct{}{}
	}

	for _, item := range a {
		if _, exists := bmap[item]; !exists {
			removed = append(removed, item)
		}
	}
	return
}

func (m *dsObjectStore) UpdateObject(id string, details *types.Struct, relations *pbrelation.Relations, links []string, snippet string) error {
	m.l.Lock()
	defer m.l.Unlock()

	txn, err := m.ds.NewTransaction(false)
	if err != nil {
		return fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	if details != nil || len(snippet) > 0 {
		exInfo, _ := getObjectInfo(txn, id)
		if exInfo != nil {
			if exInfo.Details.Equal(details) {
				// skip updating details
				details = nil
			}

			if exInfo.Snippet == snippet {
				// skip updating snippet
				snippet = ""
			}
		}
	}

	var addedLinks, removedLinks []string

	if links != nil {
		exLinks, _ := findOutboundLinks(txn, id)
		removedLinks, addedLinks = diffSlices(exLinks, links)
	}

	if details != nil {
		if err = m.updateDetails(txn, id, &model.ObjectDetails{Details: details}); err != nil {
			return err
		}
	}

	if relations != nil {
		if err = m.updateRelations(txn, id, relations); err != nil {
			return err
		}
	}

	if len(addedLinks) > 0 {
		for _, k := range pageLinkKeys(id, nil, addedLinks) {
			if err := txn.Put(k, nil); err != nil {
				return err
			}
		}
	}

	if len(removedLinks) > 0 {
		for _, k := range pageLinkKeys(id, nil, removedLinks) {
			if err := txn.Delete(k); err != nil {
				return err
			}
		}
	}

	if len(snippet) > 0 {
		if err = m.updateSnippet(txn, id, snippet); err != nil {
			return err
		}
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	if details != nil && details.Fields != nil {
		m.sendUpdatesToSubscriptions(id, details)
	}

	return nil
}

func (m *dsObjectStore) sendUpdatesToSubscriptions(id string, details *types.Struct) {
	details.Fields[database.RecordIDField] = pb.ToValue(id)
	for _, sub := range m.subscriptions {
		for _, subId := range sub.ids {
			if subId == id {
				go func(quit chan struct{}, ch chan *types.Struct) {
					select {
					case ch <- details:
						break
					case <-quit:
						break
					}
				}(sub.quit, sub.ch)
				break
			}
		}
	}
}

func (m *dsObjectStore) UpdateLastOpened(id string, time time.Time) error {
	txn, err := m.ds.NewTransaction(false)
	if err != nil {
		return fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	details, err := getDetails(txn, id)
	if err != nil && err != ds.ErrNotFound {
		return err
	}

	if details == nil || details.Details == nil || details.Details.Fields == nil {
		details = &model.ObjectDetails{Details: &types.Struct{Fields: make(map[string]*types.Value)}}
	}

	details.Details.Fields[fieldLastOpened] = structs.Float64(float64(time.Unix()))

	if err := m.updateDetails(txn, id, details); err != nil {
		return err
	}

	return txn.Commit()
}

func (m *dsObjectStore) UpdateLastModified(id string, time time.Time) error {
	txn, err := m.ds.NewTransaction(false)
	if err != nil {
		return fmt.Errorf("error creating txn in datastore: %w", err)
	}
	defer txn.Discard()

	details, err := getDetails(txn, id)
	if err != nil && err != ds.ErrNotFound {
		return err
	}

	if details == nil || details.Details == nil || details.Details.Fields == nil {
		details = &model.ObjectDetails{Details: &types.Struct{Fields: make(map[string]*types.Value)}}
	}

	details.Details.Fields[fieldLastModified] = structs.Float64(float64(time.Unix()))

	err = m.updateDetails(txn, id, details)
	if err != nil {
		return err
	}

	return txn.Commit()
}

func (m *dsObjectStore) updateDetails(txn ds.Txn, id string, details *model.ObjectDetails) error {
	detailsKey := pagesDetailsBase.ChildString(id)
	b, err := proto.Marshal(details)
	if err != nil {
		return err
	}

	return txn.Put(detailsKey, b)
}

func (m *dsObjectStore) updateRelations(txn ds.Txn, id string, relations *pbrelation.Relations) error {
	relationsKey := pagesRelationsBase.ChildString(id)
	b, err := proto.Marshal(relations)
	if err != nil {
		return err
	}

	return txn.Put(relationsKey, b)
}

func (m *dsObjectStore) updateSnippet(txn ds.Txn, id string, snippet string) error {
	snippetKey := pagesSnippetBase.ChildString(id)
	return txn.Put(snippetKey, []byte(snippet))
}

func (m *dsObjectStore) Prefix() string {
	return pagesPrefix
}

func (m *dsObjectStore) Indexes() []Index {
	return nil
}

/* internal */

func getDetails(txn ds.Txn, id string) (*model.ObjectDetails, error) {
	var details model.ObjectDetails
	if val, err := txn.Get(pagesDetailsBase.ChildString(id)); err != nil && err != ds.ErrNotFound {
		return nil, fmt.Errorf("failed to get details: %w", err)
	} else if err := proto.Unmarshal(val, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

func getObjectInfo(txn ds.Txn, id string) (*model.ObjectInfo, error) {
	sbt, err := smartblock.SmartBlockTypeFromID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to extract smartblock type: %w", err)
	}
	if sbt == smartblock.SmartBlockTypeArchive {
		return nil, ErrNotAPage
	}

	var details model.ObjectDetails
	if val, err := txn.Get(pagesDetailsBase.ChildString(id)); err != nil {
		return nil, fmt.Errorf("failed to get details: %w", err)
	} else if err := proto.Unmarshal(val, &details); err != nil {
		return nil, fmt.Errorf("failed to unmarshal details: %w", err)
	}

	var relations pbrelation.Relations
	if val, err := txn.Get(pagesRelationsBase.ChildString(id)); err != nil {
		if err != ds.ErrNotFound {
			return nil, fmt.Errorf("failed to get relations: %w", err)
		}
	} else if err := proto.Unmarshal(val, &relations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal relations: %w", err)
	}

	var objectTypes []string
	// remove hardcoded type
	// todo: maybe we should move it to a separate key?
	if details.Details != nil && details.Details.Fields != nil && details.Details.Fields["type"] != nil {
		vals := details.Details.Fields["type"].GetListValue()
		for _, val := range vals.Values {
			objectTypes = append(objectTypes, val.GetStringValue())
		}
	}

	var snippet string
	if val, err := txn.Get(pagesSnippetBase.ChildString(id)); err != nil && err != ds.ErrNotFound {
		return nil, fmt.Errorf("failed to get snippet: %w", err)
	} else {
		snippet = string(val)
	}

	// omit decoding page state
	hasInbound, err := hasInboundLinks(txn, id)
	if err != nil {
		return nil, err
	}

	return &model.ObjectInfo{
		Id:              id,
		ObjectType:      sbt.ToProto(),
		Details:         details.Details,
		Relations:       &relations,
		Snippet:         snippet,
		HasInboundLinks: hasInbound,
		ObjectTypeUrls:  objectTypes,
	}, nil
}

func getPagesInfo(txn ds.Txn, ids []string) ([]*model.ObjectInfo, error) {
	var pages []*model.ObjectInfo
	for _, id := range ids {
		info, err := getObjectInfo(txn, id)
		if err != nil {
			if strings.HasSuffix(err.Error(), "key not found") || err == ErrNotAPage {
				continue
			}
			return nil, err
		}
		pages = append(pages, info)
	}

	return pages, nil
}

func hasInboundLinks(txn ds.Txn, id string) (bool, error) {
	inboundResults, err := txn.Query(query.Query{
		Prefix:   pagesInboundLinksBase.String() + "/" + id + "/",
		Limit:    1, // we only need to know if there is at least 1 inbound link
		KeysOnly: true,
	})
	if err != nil {
		return false, err
	}

	// max is 1
	inboundLinks, err := CountAllKeysFromResults(inboundResults)
	return inboundLinks > 0, err
}

// Find to which IDs specified one has outbound links.
func findOutboundLinks(txn ds.Txn, id string) ([]string, error) {
	return findByPrefix(txn, pagesOutboundLinksBase.String()+"/"+id+"/", 0)
}

// Find from which IDs specified one has inbound links.
func findInboundLinks(txn ds.Txn, id string) ([]string, error) {
	return findByPrefix(txn, pagesInboundLinksBase.String()+"/"+id+"/", 0)
}

func findByPrefix(txn ds.Txn, prefix string, limit int) ([]string, error) {
	results, err := txn.Query(query.Query{
		Prefix:   prefix,
		Limit:    limit,
		KeysOnly: true,
	})
	if err != nil {
		return nil, err
	}

	return GetLeavesFromResults(results)
}

func pageLinkKeys(id string, in []string, out []string) []ds.Key {
	var keys = make([]ds.Key, 0, len(in)+len(out))

	// links incoming into specified node id
	for _, from := range in {
		keys = append(keys, inboundLinkKey(from, id), outgoingLinkKey(from, id))
	}

	// links outgoing from specified node id
	for _, to := range out {
		keys = append(keys, outgoingLinkKey(id, to), inboundLinkKey(id, to))
	}

	return keys
}

func outgoingLinkKey(from, to string) ds.Key {
	return pagesOutboundLinksBase.ChildString(from).ChildString(to)
}

func inboundLinkKey(from, to string) ds.Key {
	return pagesInboundLinksBase.ChildString(to).ChildString(from)
}
