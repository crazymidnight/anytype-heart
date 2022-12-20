package editor

import (
	bookmarksvc "github.com/anytypeio/go-anytype-middleware/core/block/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/clipboard"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/file"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/stext"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/table"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
)

type Page struct {
	smartblock.SmartBlock
	basic.AllOperations
	basic.IHistory
	file.File
	stext.Text
	clipboard.Clipboard
	bookmark.Bookmark
	dataview.Dataview
	table.TableEditor
}

func NewPage(
	fileSource file.BlockService,
	pageManager bookmark.BlockService,
	bookmarkSvc bookmark.BookmarkService,
) *Page {
	sb := smartblock.New()
	f := file.NewFile(sb, fileSource)
	return &Page{
		SmartBlock:    sb,
		AllOperations: basic.NewBasic(sb),
		IHistory:      basic.NewHistory(sb),
		Text:          stext.NewText(sb),
		File:          f,
		Clipboard:     clipboard.NewClipboard(sb, f),
		Bookmark:      bookmark.NewBookmark(sb, pageManager, bookmarkSvc),
		Dataview:      dataview.NewDataview(sb),
		TableEditor:   table.NewEditor(sb),
	}
}

func (p *Page) Init(ctx *smartblock.InitContext) (err error) {
	if ctx.ObjectTypeUrls == nil {
		ctx.ObjectTypeUrls = []string{bundle.TypeKeyPage.URL()}
	}
	newDoc := ctx.State != nil
	if err = p.SmartBlock.Init(ctx); err != nil {
		return
	}
	layout, ok := ctx.State.Layout()
	if !ok {
		otypes, _ := objectstore.GetObjectTypes(p.ObjectStore(), ctx.ObjectTypeUrls)
		for _, ot := range otypes {
			layout = ot.Layout
		}
	}

	tmpls := []template.StateTransformer{
		template.WithObjectTypesAndLayout(ctx.ObjectTypeUrls, layout),
		bookmarksvc.WithFixedBookmarks(p.Bookmark),
	}

	// replace title to text block for note
	if newDoc && layout == model.ObjectType_note {
		if name := pbtypes.GetString(ctx.State.Details(), bundle.RelationKeyName.String()); name != "" {
			ctx.State.RemoveDetail(bundle.RelationKeyName.String())
			tmpls = append(tmpls, template.WithFirstTextBlockContent(name))
		}
	}

	return smartblock.ObjectApplyTemplate(p, ctx.State,
		template.ByLayout(
			layout,
			tmpls...,
		)...,
	)
}
