package indexer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/threads"

	"github.com/anytypeio/go-anytype-middleware/change"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
	"github.com/gogo/protobuf/types"
)

func newDoc(id string, a core.Service) (d *doc, err error) {
	sb, err := a.GetBlock(id)
	if err != nil {
		err = fmt.Errorf("anytype.GetBlock error: %v", err)
		return
	}

	d = &doc{
		id:        id,
		lastUsage: time.Now(),
		store:     a.ObjectStore(),
		sb:        sb,
	}

	d.tree, _, err = change.BuildMetaTree(sb)
	if err == change.ErrEmpty {
		d.tree = change.NewMetaTree()
		d.st = state.NewDoc(id, nil).(*state.State)
		err = nil
	} else if err != nil {
		return
	} else {
		if d.st, err = d.buildState(); err != nil {
			return
		}
	}
	return
}

type doc struct {
	id   string
	tree *change.Tree
	st   *state.State

	changesBuf []*change.Change
	store      detailsGetter
	lastUsage  time.Time
	mu         sync.Mutex
	sb         core.SmartBlock
}

type detailsGetter interface {
	GetDetails(id string) (*model.ObjectDetails, error)
}

func (d *doc) meta() core.SmartBlockMeta {
	d.mu.Lock()
	defer d.mu.Unlock()
	objectTypes := make([]string, len(d.st.ObjectTypes()))
	copy(objectTypes, d.st.ObjectTypes())
	return core.SmartBlockMeta{
		ObjectTypes: objectTypes,
		Relations:   pbtypes.CopyRelations(d.st.ExtraRelations()),
		Details:     pbtypes.CopyStruct(d.st.Details()),
	}
}

func (d *doc) addRecords(records ...core.SmartblockRecordEnvelope) (lastChangeTS int64, lastChangeOwner string, hasMetaChanges bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.lastUsage = time.Now()
	var changes = d.changesBuf[:0]
	for _, rec := range records {
		c, err := change.NewChangeFromRecord(rec)
		if err != nil {
			log.Warnf("indexer: can't make change from record: %v", err)
			continue
		}
		if n := time.Now().Unix(); c.Timestamp > n {
			c.Timestamp = n
		}

		if c.Timestamp > lastChangeTS {
			lastChangeTS = c.Timestamp
			lastChangeOwner = c.Account
		}

		if c.HasMeta() {
			changes = append(changes, c)
		}
	}
	if len(changes) == 0 {
		return
	}

	switch d.tree.Add(changes...) {
	case change.Nothing:
		return
	case change.Append:
		s, err := change.BuildStateSimpleCRDT(d.st, d.tree)
		if err != nil {
			log.Warnf("indexer: can't build crdt state (append): %v", err)
			return
		}
		_, _, err = state.ApplyStateFast(s)
		if err != nil {
			log.Warnf("indexer: can't apply state: %v", err)
			return
		}
		hasMetaChanges = true
		return
	case change.Rebuild:
		doc, err := d.buildState()
		if err != nil {
			log.Warnf("indexer: can't build crdt state (rebuild): %v", err)
			return
		}
		d.st = doc
		hasMetaChanges = true
		return
	}
	return
}

func (d *doc) injectLocalRelations(st *state.State) {
	if details, err := d.store.GetDetails(d.id); err == nil {
		if details != nil && details.Details != nil {
			for key, v := range details.Details.Fields {
				if slice.FindPos(bundle.LocalOnlyRelationsKeys, key) != -1 {
					st.SetDetail(key, v)
				}
			}
		}
	}
}

func (s *doc) findFirstChange(ctx context.Context) (c *change.Change, err error) {
	if s.tree.RootId() == "" {
		return nil, change.ErrEmpty
	}
	c = s.tree.Get(s.tree.RootId())
	for c.LastSnapshotId != "" {
		var rec *core.SmartblockRecordEnvelope
		if rec, err = s.sb.GetRecord(ctx, c.LastSnapshotId); err != nil {
			return
		}
		if c, err = change.NewChangeFromRecord(*rec); err != nil {
			return
		}
	}
	return
}

func (d *doc) buildState() (doc *state.State, err error) {
	root := d.tree.Root()
	if root == nil || root.GetSnapshot() == nil {
		return nil, fmt.Errorf("root missing or not a snapshot")
	}
	doc = state.NewDocFromSnapshot(d.id, root.GetSnapshot()).(*state.State)
	doc.SetChangeId(root.Id)
	st, err := change.BuildStateSimpleCRDT(doc, d.tree)
	if err != nil {
		return
	}

	d.injectLocalRelations(st)
	st.InjectDerivedDetails()

	err = d.injectCreationInfo(st)
	if err != nil {
		log.With("thread", d.id).Errorf("injectCreationInfo failed: %s", err.Error())
	}
	if _, _, err = state.ApplyStateFast(st); err != nil {
		return
	}
	return
}

func (d *doc) injectCreationInfo(st *state.State) (err error) {
	if pbtypes.HasField(st.Details(), bundle.RelationKeyCreator.String()) {
		return nil
	}

	// protect from the big documents with a large trees
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	fc, err := d.findFirstChange(ctx)
	if err == change.ErrEmpty {
		err = nil
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to find first change to derive creation info")
	}

	st.SetDetail(bundle.RelationKeyCreatedDate.String(), pbtypes.Float64(float64(fc.Timestamp)))
	if profileId, e := threads.ProfileThreadIDFromAccountAddress(fc.Account); e == nil {
		st.SetDetail(bundle.RelationKeyCreator.String(), pbtypes.String(profileId.String()))
	}
	return
}

func (d *doc) SetDetail(key string, val *types.Value) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.st.SetDetail(key, val)
}
