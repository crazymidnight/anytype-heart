package bookmark

import (
	"context"
	"fmt"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"sync"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/bookmark"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	coresb "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/linkpreview"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/uri"
	"github.com/globalsign/mgo/bson"
	"github.com/gogo/protobuf/types"
)

func NewBookmark(sb smartblock.SmartBlock, lp linkpreview.LinkPreview, blockService BlockService) Bookmark {
	return &sbookmark{SmartBlock: sb, lp: lp, blockService: blockService}
}

type Bookmark interface {
	Fetch(ctx *state.Context, id string, url string, isSync bool) (err error)
	CreateAndFetch(ctx *state.Context, req pb.RpcBlockBookmarkCreateAndFetchRequest) (newId string, err error)
	UpdateBookmark(id, groupId string, apply func(b bookmark.Block) error) (err error)
}

type sbookmark struct {
	smartblock.SmartBlock
	lp           linkpreview.LinkPreview
	blockService BlockService
}

type BlockService interface {
	PageManager
	DoBookmark(id string, apply func(b Bookmark) error) error
}

type PageManager interface {
	// TODO: remove
	CreateSmartBlock(ctx context.Context, sbType coresb.SmartBlockType, details *types.Struct, relations []*model.Relation) (id string, newDetails *types.Struct, err error)
	CreateSmartBlockFromState(ctx context.Context, sbType coresb.SmartBlockType, details *types.Struct, relations []*model.Relation, createState *state.State) (id string, newDetails *types.Struct, err error)
	SetDetails(ctx *state.Context, req pb.RpcObjectSetDetailsRequest) (err error)
}

func (b *sbookmark) Fetch(ctx *state.Context, id string, url string, isSync bool) (err error) {
	s := b.NewStateCtx(ctx).SetGroupId(bson.NewObjectId().Hex())
	if err = b.fetch(s, id, url, isSync); err != nil {
		return
	}
	return b.Apply(s)
}

// TODO isSync is controversial, it muds the logic
func (b *sbookmark) fetch(s *state.State, id, url string, isSync bool) (err error) {
	bb := s.Get(id)
	if b == nil {
		return smartblock.ErrSimpleBlockNotFound
	}
	url, err = uri.ProcessURI(url)
	if err != nil {
		// Do nothing
	}
	groupId := s.GroupId()
	var updMu sync.Mutex
	bm, ok := bb.(bookmark.Block)
	if !ok {
		return fmt.Errorf("unexpected simple bock type: %T (want Bookmark)", bb)
	}

	err = Fetch(id, FetchParams{
		Url:     url,
		Anytype: b.Anytype(),
		Updater: func(id string, apply func(b bookmark.Block) error) (err error) {
			if isSync {
				updMu.Lock()
				defer updMu.Unlock()
				return b.updateBlock(bm, apply)
			}
			return b.blockService.DoBookmark(b.Id(), func(b Bookmark) error {
				return b.UpdateBookmark(id, groupId, apply)
			})
		},
		LinkPreview: b.lp,
		Sync:        isSync,
	})
	return err
}

func (b *sbookmark) CreateAndFetch(ctx *state.Context, req pb.RpcBlockBookmarkCreateAndFetchRequest) (newId string, err error) {
	s := b.NewStateCtx(ctx).SetGroupId(bson.NewObjectId().Hex())
	nb := simple.New(&model.Block{
		Content: &model.BlockContentOfBookmark{
			Bookmark: &model.BlockContentBookmark{
				Url: req.Url,
			},
		},
	})
	s.Add(nb)
	newId = nb.Model().Id
	if err = s.InsertTo(req.TargetId, req.Position, newId); err != nil {
		return
	}
	if err = b.fetch(s, newId, req.Url, false); err != nil {
		return
	}
	if err = b.Apply(s); err != nil {
		return
	}
	return
}

func CreateBookmarkObject(store objectstore.ObjectStore, manager PageManager, content *model.BlockContentBookmark) (pageId string, err error) {
	records, _, err := store.Query(nil, database.Query{
		Sorts: []*model.BlockContentDataviewSort{
			{
				RelationKey: bundle.RelationKeyLastModifiedDate.String(),
				Type:        model.BlockContentDataviewSort_Desc,
			},
		},
		Filters: []*model.BlockContentDataviewFilter{
			{
				Condition:   model.BlockContentDataviewFilter_Equal,
				RelationKey: bundle.RelationKeyUrl.String(),
				Value:       pbtypes.String(content.Url),
			},
		},
		Limit: 1,
		ObjectTypeFilter: []string{
			bundle.TypeKeyBookmark.URL(),
		},
	})
	if err != nil {
		return "", fmt.Errorf("query: %w", err)
	}

	ogDetails := map[string]*types.Value{
		bundle.RelationKeyDescription.String(): pbtypes.String(content.Description),
		bundle.RelationKeyUrl.String():         pbtypes.String(content.Url),
		bundle.RelationKeyPicture.String():     pbtypes.String(content.ImageHash),
		bundle.RelationKeyIconImage.String():   pbtypes.String(content.FaviconHash),
	}

	if len(records) > 0 {
		rec := records[0]

		details := make([]*pb.RpcObjectSetDetailsDetail, 0, len(ogDetails))
		for k, v := range ogDetails {
			details = append(details, &pb.RpcObjectSetDetailsDetail{
				Key:   k,
				Value: v,
			})
		}

		pageId = rec.Details.Fields[bundle.RelationKeyId.String()].GetStringValue()
		err = manager.SetDetails(nil, pb.RpcObjectSetDetailsRequest{
			ContextId: pageId,
			Details:   details,
		})
		if err != nil {
			return "", fmt.Errorf("set details: %w", err)
		}
		return pageId, nil
	}

	details := &types.Struct{
		Fields: map[string]*types.Value{
			bundle.RelationKeyType.String(): pbtypes.String(bundle.TypeKeyBookmark.URL()),
			bundle.RelationKeyName.String(): pbtypes.String(content.Title),
		},
	}
	for k, v := range ogDetails {
		details.Fields[k] = v
	}

	// pageId, _, err = manager.CreateSmartBlock(context.TODO(), coresb.SmartBlockTypePage, details, nil)
	st := state.NewDoc("", nil).NewState()

	relationKeys := []bundle.RelationKey{
		bundle.RelationKeyUrl,
		bundle.RelationKeyPicture,
		bundle.RelationKeyCreatedDate,
		bundle.RelationKeyTag,
		bundle.RelationKeyNotes,
		bundle.RelationKeyQuote,
	}
	blocks := make([]*model.Block, 0, len(relationKeys))
	for _, k := range relationKeys {
		blocks = append(blocks, &model.Block{
			Id: k.String(),
			Content: &model.BlockContentOfRelation{
				Relation: &model.BlockContentRelation{
					Key: k.String(),
				},
			},
		})
	}

	if err = template.InitTemplate(st, template.WithRootBlocks(blocks)); err != nil {
		return "", fmt.Errorf("init template: %w", err)
	}

	pageId, _, err = manager.CreateSmartBlockFromState(context.TODO(), coresb.SmartBlockTypePage, details, nil, st)
	return
}

func (b *sbookmark) UpdateBookmark(id, groupId string, apply func(b bookmark.Block) error) error {
	s := b.NewState().SetGroupId(groupId)
	if bb := s.Get(id); bb != nil {
		if bm, ok := bb.(bookmark.Block); ok {
			if err := b.updateBlock(bm, apply); err != nil {
				return fmt.Errorf("update block: %w", err)
			}
		} else {
			return fmt.Errorf("unexpected simple bock type: %T (want Bookmark)", bb)
		}
	} else {
		return smartblock.ErrSimpleBlockNotFound
	}
	return b.Apply(s)
}

// updateBlock updates a block and creates associated Bookmark object
func (b *sbookmark) updateBlock(block bookmark.Block, apply func(bookmark.Block) error) error {
	if err := apply(block); err != nil {
		return err
	}

	store := b.Anytype().ObjectStore()
	pageId, err := CreateBookmarkObject(store, b.blockService, block.GetContent())
	if err != nil {
		return fmt.Errorf("create bookmark object: %w", err)
	}

	block.SetTargetObjectId(pageId)
	return nil
}
