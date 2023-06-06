package objectstore

import (
	"context"
	"testing"

	"github.com/anyproto/any-sync/app"
	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	dsbadgerv3 "github.com/textileio/go-ds-badger3"

	"github.com/anyproto/anytype-heart/core/wallet"
	"github.com/anyproto/anytype-heart/core/wallet/mock_wallet"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/database"
	"github.com/anyproto/anytype-heart/pkg/lib/datastore/noctxds"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/ftsearch"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/typeprovider/mock_typeprovider"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

func makeDetails(fields map[bundle.RelationKey]*types.Value) *types.Struct {
	f := map[string]*types.Value{}
	for k, v := range fields {
		f[string(k)] = v
	}
	return &types.Struct{Fields: f}
}

type storeFixture struct {
	*dsObjectStore
}

func newStoreFixture(t *testing.T) *storeFixture {
	ds, err := dsbadgerv3.NewDatastore(t.TempDir(), &dsbadgerv3.DefaultOptions)
	require.NoError(t, err)

	noCtxDS := noctxds.New(ds)

	typeProvider := mock_typeprovider.NewMockSmartBlockTypeProvider(t)
	typeProvider.EXPECT().Type(mock.Anything).Return(smartblock.SmartBlockTypePage, nil)

	walletService := mock_wallet.NewMockWallet(t)
	walletService.EXPECT().Name().Return(wallet.CName)
	walletService.EXPECT().RepoPath().Return(t.TempDir())

	fullText := ftsearch.New()
	testApp := &app.App{}
	testApp.Register(walletService)
	err = fullText.Init(testApp)
	require.NoError(t, err)
	err = fullText.Run(context.Background())
	require.NoError(t, err)

	return &storeFixture{&dsObjectStore{
		ds:          noCtxDS,
		sbtProvider: typeProvider,
		fts:         fullText,
	}}
}

func (fx *storeFixture) addObjects(t *testing.T, objects []map[bundle.RelationKey]*types.Value) {
	for _, obj := range objects {
		id := obj[bundle.RelationKeyId].GetStringValue()
		require.NotEmpty(t, id)
		err := fx.UpdateObjectDetails(id, makeDetails(obj), false)
		require.NoError(t, err)
	}
}

func assertRecordsEqual(t *testing.T, want []map[bundle.RelationKey]*types.Value, got []database.Record) {
	wantRaw := make([]database.Record, 0, len(want))
	for _, w := range want {
		wantRaw = append(wantRaw, database.Record{Details: makeDetails(w)})
	}
	assert.Equal(t, wantRaw, got)
}

func assertRecordsMatch(t *testing.T, want []map[bundle.RelationKey]*types.Value, got []database.Record) {
	wantRaw := make([]database.Record, 0, len(want))
	for _, w := range want {
		wantRaw = append(wantRaw, database.Record{Details: makeDetails(w)})
	}
	assert.ElementsMatch(t, wantRaw, got)
}

func TestQuery(t *testing.T) {
	t.Run("no filters", func(t *testing.T) {
		s := newStoreFixture(t)
		s.addObjects(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:   pbtypes.String("id1"),
				bundle.RelationKeyName: pbtypes.String("name1"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id2"),
				bundle.RelationKeyName: pbtypes.String("name2"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id3"),
				bundle.RelationKeyName: pbtypes.String("name3"),
			},
		})

		recs, _, err := s.Query(nil, database.Query{})
		require.NoError(t, err)

		assertRecordsEqual(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:   pbtypes.String("id1"),
				bundle.RelationKeyName: pbtypes.String("name1"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id2"),
				bundle.RelationKeyName: pbtypes.String("name2"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id3"),
				bundle.RelationKeyName: pbtypes.String("name3"),
			},
		}, recs)
	})

	t.Run("with filter", func(t *testing.T) {
		s := newStoreFixture(t)
		s.addObjects(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:   pbtypes.String("id1"),
				bundle.RelationKeyName: pbtypes.String("name1"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id2"),
				bundle.RelationKeyName: pbtypes.String("name2"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id3"),
				bundle.RelationKeyName: pbtypes.String("name3"),
			},
		})

		recs, _, err := s.Query(nil, database.Query{
			Filters: []*model.BlockContentDataviewFilter{
				{
					RelationKey: bundle.RelationKeyName.String(),
					Condition:   model.BlockContentDataviewFilter_Equal,
					Value:       pbtypes.String("name2"),
				},
			},
		})
		require.NoError(t, err)

		assertRecordsEqual(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:   pbtypes.String("id2"),
				bundle.RelationKeyName: pbtypes.String("name2"),
			},
		}, recs)
	})

	t.Run("with multiple filters", func(t *testing.T) {
		s := newStoreFixture(t)
		s.addObjects(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:   pbtypes.String("id1"),
				bundle.RelationKeyName: pbtypes.String("name"),
			},
			{
				bundle.RelationKeyId:          pbtypes.String("id2"),
				bundle.RelationKeyName:        pbtypes.String("name"),
				bundle.RelationKeyDescription: pbtypes.String("description"),
			},
			{
				bundle.RelationKeyId:          pbtypes.String("id3"),
				bundle.RelationKeyName:        pbtypes.String("name"),
				bundle.RelationKeyDescription: pbtypes.String("description"),
			},
		})

		recs, _, err := s.Query(nil, database.Query{
			Filters: []*model.BlockContentDataviewFilter{
				{
					RelationKey: bundle.RelationKeyName.String(),
					Condition:   model.BlockContentDataviewFilter_Equal,
					Value:       pbtypes.String("name"),
				},
				{
					RelationKey: bundle.RelationKeyDescription.String(),
					Condition:   model.BlockContentDataviewFilter_Equal,
					Value:       pbtypes.String("description"),
				},
			},
		})
		require.NoError(t, err)

		assertRecordsEqual(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:          pbtypes.String("id2"),
				bundle.RelationKeyName:        pbtypes.String("name"),
				bundle.RelationKeyDescription: pbtypes.String("description"),
			},
			{
				bundle.RelationKeyId:          pbtypes.String("id3"),
				bundle.RelationKeyName:        pbtypes.String("name"),
				bundle.RelationKeyDescription: pbtypes.String("description"),
			},
		}, recs)
	})

	t.Run("full text search", func(t *testing.T) {
		s := newStoreFixture(t)
		s.addObjects(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:   pbtypes.String("id1"),
				bundle.RelationKeyName: pbtypes.String("name"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id2"),
				bundle.RelationKeyName: pbtypes.String("some important note"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id3"),
				bundle.RelationKeyName: pbtypes.String(""),
			},
		})

		err := s.fts.Index(ftsearch.SearchDoc{
			Id:    "id1",
			Title: "name",
		})
		require.NoError(t, err)

		err = s.fts.Index(ftsearch.SearchDoc{
			Id:    "id2",
			Title: "some important note",
		})
		require.NoError(t, err)

		err = s.fts.Index(ftsearch.SearchDoc{
			Id:    "id3",
			Title: "",
			Text:  "very important text",
		})
		require.NoError(t, err)

		recs, _, err := s.Query(nil, database.Query{
			FullText: "important",
		})
		require.NoError(t, err)

		// Full-text engine has its own ordering, so just don't rely on it here and check only the content.
		assertRecordsMatch(t, []map[bundle.RelationKey]*types.Value{
			{
				bundle.RelationKeyId:   pbtypes.String("id2"),
				bundle.RelationKeyName: pbtypes.String("some important note"),
			},
			{
				bundle.RelationKeyId:   pbtypes.String("id3"),
				bundle.RelationKeyName: pbtypes.String(""),
			},
		}, recs)
	})

	t.Run("with system objects", func(t *testing.T) {

	})

	t.Run("with object type filter", func(t *testing.T) {

	})

	t.Run("with limit", func(t *testing.T) {

	})

	t.Run("with limit and offset", func(t *testing.T) {

	})

	t.Run("with ascending order", func(t *testing.T) {

	})

	t.Run("with descending order", func(t *testing.T) {

	})

	t.Run("with multiple orders", func(t *testing.T) {

	})
}
