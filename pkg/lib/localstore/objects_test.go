package localstore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/threads"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/ftsearch"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	pbrelation "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
	"github.com/gogo/protobuf/types"
	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/ipfs/go-datastore/sync"
	badger "github.com/ipfs/go-ds-badger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/textileio/go-threads/core/thread"
)

func TestDsObjectStore_IndexQueue(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(tmpDir)

	bds, err := badger.NewDatastore(tmpDir, nil)
	require.NoError(t, err)

	ds := NewObjectStore(bds, nil)

	require.NoError(t, ds.AddToIndexQueue("one"))
	require.NoError(t, ds.AddToIndexQueue("one"))
	require.NoError(t, ds.AddToIndexQueue("two"))
	var count int
	require.NoError(t, ds.IndexForEach(func(id string, tm time.Time) error {
		assert.NotEqual(t, -1, slice.FindPos([]string{"one", "two"}, id))
		assert.NotEmpty(t, tm)
		count++
		if id == "one" {
			return nil
		} else {
			return fmt.Errorf("test err")
		}
	}))
	assert.Equal(t, 2, count)
	count = 0
	require.NoError(t, ds.IndexForEach(func(id string, tm time.Time) error {
		assert.Equal(t, "two", id)
		assert.NotEmpty(t, tm)
		count++
		return nil
	}))
	assert.Equal(t, 1, count)

	count = 0
	require.NoError(t, ds.IndexForEach(func(id string, tm time.Time) error {
		count++
		return nil
	}))

	assert.Equal(t, 0, count)

	require.NoError(t, ds.AddToIndexQueue("one"))
	require.NoError(t, ds.AddToIndexQueue("one"))
	require.NoError(t, ds.AddToIndexQueue("two"))

	count = 0
	require.NoError(t, ds.IndexForEach(func(id string, tm time.Time) error {
		count++
		return nil
	}))
	assert.Equal(t, 2, count)
}

func TestDsObjectStore_Query(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(tmpDir)

	fts, err := ftsearch.NewFTSearch(filepath.Join(tmpDir, "fts"))
	require.NoError(t, err)

	bds, err := badger.NewDatastore(tmpDir, nil)
	require.NoError(t, err)

	ds := NewObjectStore(bds, fts)
	defer ds.Close()
	newDet := func(name string) *types.Struct {
		return &types.Struct{
			Fields: map[string]*types.Value{
				"name": pbtypes.String(name),
			},
		}
	}
	tid1, _ := threads.ThreadCreateID(thread.AccessControlled, smartblock.SmartBlockTypePage)
	tid2, _ := threads.ThreadCreateID(thread.AccessControlled, smartblock.SmartBlockTypePage)
	tid3, _ := threads.ThreadCreateID(thread.AccessControlled, smartblock.SmartBlockTypePage)
	id1 := tid1.String()
	id2 := tid2.String()
	id3 := tid3.String()
	require.NoError(t, ds.UpdateObject(id1, newDet("one"), nil, nil, "s1"))
	require.NoError(t, ds.UpdateObject(id2, newDet("two"), nil, nil, "s2"))
	require.NoError(t, ds.UpdateObject(id3, newDet("three"), nil, nil, "s3"))
	require.NoError(t, fts.Index(ftsearch.SearchDoc{
		Id:    id1,
		Title: "one",
		Text:  "text twoone uniqone",
	}))
	require.NoError(t, fts.Index(ftsearch.SearchDoc{
		Id:    id2,
		Title: "two",
		Text:  "twoone text twoone uniqtwo",
	}))
	require.NoError(t, fts.Index(ftsearch.SearchDoc{
		Id:    id3,
		Title: "three",
		Text:  "text uniqthree",
	}))

	// should return all records
	rec, tot, err := ds.Query(nil, database.Query{})
	require.NoError(t, err)
	assert.Equal(t, 3, tot)
	assert.Len(t, rec, 3)

	// filter
	rec, tot, err = ds.Query(nil, database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				Operator:    model.BlockContentDataviewFilter_And,
				RelationKey: "name",
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String("two"),
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 1, tot)
	assert.Len(t, rec, 1)

	// fulltext
	rec, tot, err = ds.Query(nil, database.Query{
		FullText: "twoone",
	})
	require.NoError(t, err)
	assert.Equal(t, 2, tot)
	assert.Len(t, rec, 2)
	var names []string
	for _, r := range rec {
		names = append(names, pbtypes.GetString(r.Details, "name"))
	}
	assert.Equal(t, []string{"two", "one"}, names)

	// fulltext + filter
	rec, tot, err = ds.Query(nil, database.Query{
		FullText: "twoone",
		Filters: []*model.BlockContentDataviewFilter{
			{
				Operator:    model.BlockContentDataviewFilter_And,
				RelationKey: "name",
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.String("one"),
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 1, tot)
	assert.Len(t, rec, 1)
}
func getId() string {
	return thread.NewIDV1(thread.Raw, 32).String()
}
func TestDsObjectStore_PrefixQuery(t *testing.T) {
	bds := sync.MutexWrap(ds.NewMapDatastore())
	err := bds.Put(ds.NewKey("/p1/abc/def/1"), []byte{})

	require.NoError(t, err)

	res, err := bds.Query(query.Query{Prefix: "/p1/abc", KeysOnly: true})
	require.NoError(t, err)

	entries, err := res.Rest()
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "/p1/abc/def/1", entries[0].Key)

}
func TestDsObjectStore_RelationsIndex(t *testing.T) {
	tmpDir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(tmpDir)

	fts, err := ftsearch.NewFTSearch(filepath.Join(tmpDir, "fts"))
	require.NoError(t, err)

	bds, err := badger.NewDatastore(tmpDir, nil)
	require.NoError(t, err)

	ds := NewObjectStore(bds, fts)
	defer ds.Close()
	newDet := func(name, objtype string) *types.Struct {
		return &types.Struct{
			Fields: map[string]*types.Value{
				"name": pbtypes.String(name),
				"type": pbtypes.StringList([]string{objtype}),
			},
		}
	}
	id1 := getId()
	id2 := getId()
	id3 := getId()
	require.NoError(t, ds.UpdateObject(id1, newDet("one", "https://anytype.io/schemas/object/bundled/a1"), &pbrelation.Relations{Relations: []*pbrelation.Relation{
		{
			Key:          "rel1",
			Format:       pbrelation.RelationFormat_status,
			Name:         "rel 1",
			DefaultValue: nil,
			SelectDict: []*pbrelation.RelationOption{
				{"id1", "option1", "red", pbrelation.RelationOption_local},
				{"id2", "option2", "red", pbrelation.RelationOption_local},
				{"id3", "option3", "red", pbrelation.RelationOption_local},
			},
		},
		{
			Key:          "rel2",
			Format:       pbrelation.RelationFormat_shorttext,
			Name:         "rel 2",
			DefaultValue: nil,
		},
	}}, nil, "s1"))

	require.NoError(t, ds.UpdateObject(id2, newDet("two", "https://anytype.io/schemas/object/bundled/a2"), &pbrelation.Relations{Relations: []*pbrelation.Relation{
		{
			Key:          "rel1",
			Format:       pbrelation.RelationFormat_status,
			Name:         "rel 1",
			DefaultValue: nil,
			SelectDict: []*pbrelation.RelationOption{
				{"id3", "option3", "yellow", pbrelation.RelationOption_local},
				{"id4", "option4", "red", pbrelation.RelationOption_local},
				{"id5", "option5", "red", pbrelation.RelationOption_local},
			},
		},
		{
			Key:          "rel3",
			Format:       pbrelation.RelationFormat_status,
			Name:         "rel 3",
			DefaultValue: nil,
			SelectDict: []*pbrelation.RelationOption{
				{"id5", "option5", "red", pbrelation.RelationOption_local},
				{"id6", "option6", "red", pbrelation.RelationOption_local},
			},
		},
		{
			Key:          "rel4",
			Format:       pbrelation.RelationFormat_tag,
			Name:         "rel 4",
			DefaultValue: nil,
			SelectDict: []*pbrelation.RelationOption{
				{"id7", "option7", "red", pbrelation.RelationOption_local},
			},
		},
	}}, nil, "s2"))
	require.NoError(t, ds.UpdateObject(id3, newDet("three", "https://anytype.io/schemas/object/bundled/a2"), nil, nil, "s3"))

	restOpts, err := ds.GetAggregatedOptions("rel1", pbrelation.RelationFormat_status, "https://anytype.io/schemas/object/bundled/ffff")
	require.NoError(t, err)
	require.Len(t, restOpts, 6)

	rels, err := ds.AggregateRelationsFromObjectsOfType("https://anytype.io/schemas/object/bundled/a1")
	require.NoError(t, err)
	require.Len(t, rels, 2)

	require.Equal(t, "rel1", rels[0].Key)
	require.Equal(t, "rel2", rels[1].Key)

	rels, err = ds.ListRelations("https://anytype.io/schemas/object/bundled/a1")
	require.NoError(t, err)
	require.Len(t, rels, len(bundle.ListRelationsKeys())+4)

	require.Equal(t, "rel1", rels[0].Key)
	require.Equal(t, "rel2", rels[1].Key)
	require.Equal(t, "rel3", rels[2].Key)
	require.Equal(t, "rel4", rels[3].Key)
}
