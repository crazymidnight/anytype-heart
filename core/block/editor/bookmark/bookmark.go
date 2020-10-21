package bookmark

import (
	"fmt"
	"sync"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/bookmark"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/linkpreview"
	"github.com/anytypeio/go-anytype-middleware/util/uri"
	"github.com/globalsign/mgo/bson"
)

func NewBookmark(sb smartblock.SmartBlock, lp linkpreview.LinkPreview, ctrl DoBookmark) Bookmark {
	return &sbookmark{SmartBlock: sb, lp: lp, ctrl: ctrl}
}

type Bookmark interface {
	Fetch(ctx *state.Context, id string, url string, isSync bool) (err error)
	CreateAndFetch(ctx *state.Context, req pb.RpcBlockBookmarkCreateAndFetchRequest) (newId string, err error)
	UpdateBookmark(id, groupId string, apply func(b bookmark.Block) error) (err error)
}

type DoBookmark interface {
	DoBookmark(id string, apply func(b Bookmark) error) error
}

type sbookmark struct {
	smartblock.SmartBlock
	lp   linkpreview.LinkPreview
	ctrl DoBookmark
}

func (b *sbookmark) Fetch(ctx *state.Context, id string, url string, isSync bool) (err error) {
	s := b.NewStateCtx(ctx).SetGroupId(bson.NewObjectId().Hex())
	if err = b.fetch(s, id, url, isSync); err != nil {
		return
	}
	return b.Apply(s)
}

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
	if bm, ok := bb.(bookmark.Block); ok {
		return bm.Fetch(bookmark.FetchParams{
			Url:     url,
			Anytype: b.Anytype(),
			Updater: func(id string, apply func(b bookmark.Block) error) (err error) {
				if isSync {
					updMu.Lock()
					defer updMu.Unlock()
					return apply(bm)
				}
				return b.ctrl.DoBookmark(b.Id(), func(b Bookmark) error {
					return b.UpdateBookmark(id, groupId, apply)
				})
			},
			LinkPreview: b.lp,
			Sync:        isSync,
		})
	}
	return fmt.Errorf("unexpected simple bock type: %T (want Bookmark)", bb)
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

func (b *sbookmark) UpdateBookmark(id, groupId string, apply func(b bookmark.Block) error) (err error) {
	s := b.NewState().SetGroupId(groupId)
	if bb := s.Get(id); bb != nil {
		if bm, ok := bb.(bookmark.Block); ok {
			if err = apply(bm); err != nil {
				return
			}
		} else {
			return fmt.Errorf("unexpected simple bock type: %T (want Bookmark)", bb)
		}
	} else {
		return smartblock.ErrSimpleBlockNotFound
	}
	return b.Apply(s)
}
