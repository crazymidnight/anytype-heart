package fileobject

import (
	"context"
	"fmt"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock/smarttest"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/core/files/mock_files"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/objectstore"
	"github.com/anyproto/anytype-heart/pkg/lib/mill"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/storage"
	mock_space "github.com/anyproto/anytype-heart/space/clientspace/mock_clientspace"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

type indexerFixture struct {
	*indexer
	fileService        *mock_files.MockService
	objectStoreFixture *objectstore.StoreFixture
}

func newIndexerFixture(t *testing.T) *indexerFixture {
	objectStore := objectstore.NewStoreFixture(t)
	fileService := mock_files.NewMockService(t)

	svc := &service{
		objectStore: objectStore,
		fileService: fileService,
	}
	ind := svc.newIndexer()

	return &indexerFixture{
		objectStoreFixture: objectStore,
		fileService:        fileService,
		indexer:            ind,
	}
}

func TestIndexer_buildDetails(t *testing.T) {
	t.Run("with file", func(t *testing.T) {
		for _, typeKey := range []domain.TypeKey{
			bundle.TypeKeyFile,
			bundle.TypeKeyAudio,
			bundle.TypeKeyVideo,
		} {
			t.Run(fmt.Sprintf("with type %s", typeKey), func(t *testing.T) {
				fx := newIndexerFixture(t)
				id := domain.FullFileId{
					SpaceId: "space1",
					FileId:  testFileId,
				}
				ctx := context.Background()

				file := mock_files.NewMockFile(t)
				file.EXPECT().Info().Return(&storage.FileInfo{
					Mill:  mill.BlobId,
					Media: "text",
				})
				file.EXPECT().Details(ctx).Return(&types.Struct{
					Fields: map[string]*types.Value{
						bundle.RelationKeyName.String(): pbtypes.String("name"),
					},
				}, typeKey, nil)
				fx.fileService.EXPECT().FileByHash(ctx, id).Return(file, nil)

				details, gotTypeKey, err := fx.buildDetails(ctx, id)
				require.NoError(t, err)
				assert.Equal(t, typeKey, gotTypeKey)
				assert.Equal(t, "name", pbtypes.GetString(details, bundle.RelationKeyName.String()))
				assert.Equal(t, pbtypes.Int64(int64(model.FileIndexingStatus_Indexed)), details.Fields[bundle.RelationKeyFileIndexingStatus.String()])
			})
		}
	})
	t.Run("with image", func(t *testing.T) {
		fx := newIndexerFixture(t)
		id := domain.FullFileId{
			SpaceId: "space1",
			FileId:  testFileId,
		}
		ctx := context.Background()

		file := mock_files.NewMockFile(t)
		file.EXPECT().Info().Return(&storage.FileInfo{
			Mill:  mill.ImageResizeId,
			Media: "image/jpeg",
		})

		image := mock_files.NewMockImage(t)
		image.EXPECT().Details(ctx).Return(&types.Struct{
			Fields: map[string]*types.Value{
				bundle.RelationKeyName.String(): pbtypes.String("name"),
			},
		}, nil)
		fx.fileService.EXPECT().FileByHash(ctx, id).Return(file, nil)
		fx.fileService.EXPECT().ImageByHash(ctx, id).Return(image, nil)

		details, gotTypeKey, err := fx.buildDetails(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, bundle.TypeKeyImage, gotTypeKey)
		assert.Equal(t, "name", pbtypes.GetString(details, bundle.RelationKeyName.String()))
		assert.Equal(t, pbtypes.Int64(int64(model.FileIndexingStatus_Indexed)), details.Fields[bundle.RelationKeyFileIndexingStatus.String()])
	})
	t.Run("with image fell back to file", func(t *testing.T) {
		for _, typeKey := range []domain.TypeKey{
			bundle.TypeKeyFile,
			bundle.TypeKeyAudio,
			bundle.TypeKeyVideo,
			bundle.TypeKeyImage,
		} {
			t.Run(fmt.Sprintf("with type %s", typeKey), func(t *testing.T) {
				fx := newIndexerFixture(t)
				id := domain.FullFileId{
					SpaceId: "space1",
					FileId:  testFileId,
				}
				ctx := context.Background()

				file := mock_files.NewMockFile(t)
				file.EXPECT().Info().Return(&storage.FileInfo{
					Mill:  mill.BlobId,
					Media: "image/jpeg",
				})
				file.EXPECT().Details(ctx).Return(&types.Struct{
					Fields: map[string]*types.Value{
						bundle.RelationKeyName.String(): pbtypes.String("name"),
					},
				}, typeKey, nil)
				fx.fileService.EXPECT().FileByHash(ctx, id).Return(file, nil)

				details, gotTypeKey, err := fx.buildDetails(ctx, id)
				require.NoError(t, err)
				assert.Equal(t, bundle.TypeKeyImage, gotTypeKey)
				assert.Equal(t, "name", pbtypes.GetString(details, bundle.RelationKeyName.String()))
				assert.Equal(t, pbtypes.Int64(int64(model.FileIndexingStatus_Indexed)), details.Fields[bundle.RelationKeyFileIndexingStatus.String()])
			})
		}
	})
}

func TestIndexer_addFromObjectStore(t *testing.T) {
	t.Run("no records in store", func(t *testing.T) {
		fx := newIndexerFixture(t)
		ctx := context.Background()

		err := fx.addToQueueFromObjectStore(ctx)
		require.NoError(t, err)

		got := fx.indexQueue.GetAll()
		assert.Empty(t, got)
	})

	t.Run("get records only with not indexed status", func(t *testing.T) {
		fx := newIndexerFixture(t)
		ctx := context.Background()

		//  Use same testFileId everywhere to pass domain.IsFileId check. It doesn't matter that files are same here
		fx.objectStoreFixture.AddObjects(t, []objectstore.TestObject{
			{
				bundle.RelationKeyId:                 pbtypes.String("id1"),
				bundle.RelationKeyFileId:             pbtypes.String(testFileId.String()),
				bundle.RelationKeySpaceId:            pbtypes.String("space1"),
				bundle.RelationKeyFileIndexingStatus: pbtypes.Int64(int64(model.FileIndexingStatus_NotIndexed)),
				bundle.RelationKeyLayout:             pbtypes.Int64(int64(model.ObjectType_file)),
			},
			{
				bundle.RelationKeyId:                 pbtypes.String("id2"),
				bundle.RelationKeyFileId:             pbtypes.String(testFileId.String()),
				bundle.RelationKeySpaceId:            pbtypes.String("space2"),
				bundle.RelationKeyFileIndexingStatus: pbtypes.Int64(int64(model.FileIndexingStatus_Indexed)),
				bundle.RelationKeyLayout:             pbtypes.Int64(int64(model.ObjectType_image)),
			},
			{
				bundle.RelationKeyId:                 pbtypes.String("id3"),
				bundle.RelationKeyFileId:             pbtypes.String(testFileId.String()),
				bundle.RelationKeySpaceId:            pbtypes.String("space3"),
				bundle.RelationKeyFileIndexingStatus: pbtypes.Int64(int64(model.FileIndexingStatus_NotFound)),
				bundle.RelationKeyLayout:             pbtypes.Int64(int64(model.ObjectType_video)),
			},
			{
				bundle.RelationKeyId:      pbtypes.String("id4"),
				bundle.RelationKeyFileId:  pbtypes.String(testFileId.String()),
				bundle.RelationKeySpaceId: pbtypes.String("space4"),
				bundle.RelationKeyLayout:  pbtypes.Int64(int64(model.ObjectType_audio)),
			},
			{
				bundle.RelationKeyId:      pbtypes.String("id5"),
				bundle.RelationKeySpaceId: pbtypes.String("space5"),
				bundle.RelationKeyLayout:  pbtypes.Int64(int64(model.ObjectType_basic)),
			},
		})

		err := fx.addToQueueFromObjectStore(ctx)
		require.NoError(t, err)

		got := fx.indexQueue.GetAll()

		want := []indexRequest{
			{id: domain.FullID{SpaceID: "space1", ObjectID: "id1"}, fileId: domain.FullFileId{SpaceId: "space1", FileId: testFileId}},
			{id: domain.FullID{SpaceID: "space3", ObjectID: "id3"}, fileId: domain.FullFileId{SpaceId: "space3", FileId: testFileId}},
			{id: domain.FullID{SpaceID: "space4", ObjectID: "id4"}, fileId: domain.FullFileId{SpaceId: "space4", FileId: testFileId}},
		}

		assert.ElementsMatch(t, want, got)
	})

	t.Run("don't add same records twice", func(t *testing.T) {
		fx := newIndexerFixture(t)
		ctx := context.Background()

		fx.objectStoreFixture.AddObjects(t, []objectstore.TestObject{
			{
				bundle.RelationKeyId:                 pbtypes.String("id1"),
				bundle.RelationKeyFileId:             pbtypes.String(testFileId.String()),
				bundle.RelationKeySpaceId:            pbtypes.String("space1"),
				bundle.RelationKeyLayout:             pbtypes.Int64(int64(model.ObjectType_audio)),
				bundle.RelationKeyFileIndexingStatus: pbtypes.Int64(int64(model.FileIndexingStatus_NotIndexed)),
			},
		})

		err := fx.addToQueueFromObjectStore(ctx)
		require.NoError(t, err)
		err = fx.addToQueueFromObjectStore(ctx)
		require.NoError(t, err)

		got := fx.indexQueue.GetAll()

		want := []indexRequest{
			{id: domain.FullID{SpaceID: "space1", ObjectID: "id1"}, fileId: domain.FullFileId{SpaceId: "space1", FileId: testFileId}},
		}

		assert.ElementsMatch(t, want, got)
	})
}

func TestIndexer_unhideRecommendedRelations(t *testing.T) {
	const spaceId = "spaceId"
	t.Run("no recommended relations that could be hidden - no panic", func(t *testing.T) {
		// given
		i := indexer{}

		// when
		i.unhideRecommendedRelations(nil, bundle.TypeKeyVideo, []domain.RelationKey{bundle.RelationKeyArtist, bundle.RelationKeyFileMimeType})

		// then
		// test is successful. If method call does not panic, then no Store or Space call was initiated
	})

	t.Run("all recommended relations are not hidden - no space call", func(t *testing.T) {
		// given
		store := objectstore.NewStoreFixture(t)
		store.AddObjects(t, []objectstore.TestObject{
			{
				bundle.RelationKeyId:        pbtypes.String(bundle.RelationKeyAudioLyrics.URL()),
				bundle.RelationKeyUniqueKey: pbtypes.String(bundle.RelationKeyAudioLyrics.URL()),
				bundle.RelationKeySpaceId:   pbtypes.String(spaceId),
				bundle.RelationKeyIsHidden:  pbtypes.Bool(false),
			},
			{
				bundle.RelationKeyId:        pbtypes.String(bundle.RelationKeyAudioAlbumTrackNumber.URL()),
				bundle.RelationKeyUniqueKey: pbtypes.String(bundle.RelationKeyAudioAlbumTrackNumber.URL()),
				bundle.RelationKeySpaceId:   pbtypes.String(spaceId),
				bundle.RelationKeyIsHidden:  pbtypes.Bool(false),
			},
		})
		i := indexer{objectStore: store}

		space := mock_space.NewMockSpace(t)
		space.EXPECT().Id().Return(spaceId).Maybe()
		// space.EXPECT().Do(mock.Anything, mock.Anything)

		// when
		i.unhideRecommendedRelations(space, bundle.TypeKeyAudio, []domain.RelationKey{bundle.RelationKeyAudioAlbumTrackNumber, bundle.RelationKeyAudioLyrics})

		// then
		// no space.Do call as expected
	})

	t.Run("all recommended relations are hidden - space call for each", func(t *testing.T) {
		// given
		store := objectstore.NewStoreFixture(t)
		store.AddObjects(t, []objectstore.TestObject{
			{
				bundle.RelationKeyId:        pbtypes.String(bundle.RelationKeyAudioGenre.URL()),
				bundle.RelationKeyUniqueKey: pbtypes.String(bundle.RelationKeyAudioGenre.URL()),
				bundle.RelationKeySpaceId:   pbtypes.String(spaceId),
				bundle.RelationKeyIsHidden:  pbtypes.Bool(true),
			},
			{
				bundle.RelationKeyId:        pbtypes.String(bundle.RelationKeyAudioAlbum.URL()),
				bundle.RelationKeyUniqueKey: pbtypes.String(bundle.RelationKeyAudioAlbum.URL()),
				bundle.RelationKeySpaceId:   pbtypes.String(spaceId),
				bundle.RelationKeyIsHidden:  pbtypes.Bool(true),
			},
		})
		i := indexer{objectStore: store}
		ids := []string{bundle.RelationKeyAudioGenre.URL(), bundle.RelationKeyAudioAlbum.URL()}

		space := mock_space.NewMockSpace(t)
		space.EXPECT().Id().Return(spaceId).Maybe()
		space.EXPECT().Do(mock.Anything, mock.Anything).RunAndReturn(func(id string, apply func(smartblock.SmartBlock) error) error {
			assert.Contains(t, ids, id)

			sb := smarttest.New(id)
			s := sb.NewState()
			s.SetDetail(bundle.RelationKeyIsHidden.String(), pbtypes.Bool(true))
			err := sb.Apply(s)
			require.NoError(t, err)

			err = apply(sb)
			assert.NoError(t, err)

			assert.False(t, pbtypes.GetBool(sb.CombinedDetails(), bundle.RelationKeyIsHidden.String()))
			return nil
		}).Times(2)

		// when
		i.unhideRecommendedRelations(space, bundle.TypeKeyAudio, []domain.RelationKey{bundle.RelationKeyAudioGenre, bundle.RelationKeyAudioAlbum})

		// then
		// 2 space calls are expected
	})
}
