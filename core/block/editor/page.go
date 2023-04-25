package editor

import (
	bookmarksvc "github.com/anytypeio/go-anytype-middleware/core/block/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/clipboard"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/file"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/stext"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/table"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/migration"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	relation2 "github.com/anytypeio/go-anytype-middleware/core/relation"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/space/typeprovider"
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

	objectStore objectstore.ObjectStore
}

func NewPage(
	objectStore objectstore.ObjectStore,
	anytype core.Service,
	fileBlockService file.BlockService,
	bookmarkBlockService bookmark.BlockService,
	bookmarkService bookmark.BookmarkService,
	relationService relation2.Service,
	tempDirProvider core.TempDirProvider,
	sbtProvider typeprovider.SmartBlockTypeProvider,
) *Page {
	sb := smartblock.New()
	f := file.NewFile(
		sb,
		fileBlockService,
		anytype,
		tempDirProvider,
	)
	return &Page{
		SmartBlock:    sb,
		AllOperations: basic.NewBasic(sb),
		IHistory:      basic.NewHistory(sb),
		Text: stext.NewText(
			sb,
			objectStore,
		),
		File: f,
		Clipboard: clipboard.NewClipboard(
			sb,
			f,
			anytype,
			tempDirProvider,
		),
		Bookmark: bookmark.NewBookmark(
			sb,
			bookmarkBlockService,
			bookmarkService,
			objectStore,
		),
		Dataview: dataview.NewDataview(
			sb,
			anytype,
			objectStore,
			relationService,
			sbtProvider,
		),
		TableEditor: table.NewEditor(sb),

		objectStore: objectStore,
	}
}

func (p *Page) Init(ctx *smartblock.InitContext) (err error) {
	if ctx.ObjectTypeUrls == nil {
		ctx.ObjectTypeUrls = []string{bundle.TypeKeyPage.URL()}
	}

	if err = p.SmartBlock.Init(ctx); err != nil {
		return
	}
	return nil
}

func (p *Page) DefaultState(ctx *smartblock.InitContext) migration.Migration {
	layout, ok := ctx.State.Layout()
	if !ok {
		// nolint:errcheck
		otypes, _ := objectstore.GetObjectTypes(p.objectStore, ctx.ObjectTypeUrls)
		for _, ot := range otypes {
			layout = ot.Layout
		}
	}

	tmpls := []template.StateTransformer{
		template.WithObjectTypesAndLayout(ctx.ObjectTypeUrls, layout),
		bookmarksvc.WithFixedBookmarks(p.Bookmark),
	}

	// TODO Introduce Converter module
	// replace title to text block for note
	if layout == model.ObjectType_note {
		if name := pbtypes.GetString(ctx.State.Details(), bundle.RelationKeyName.String()); name != "" {
			ctx.State.RemoveDetail(bundle.RelationKeyName.String())
			tmpls = append(tmpls, template.WithFirstTextBlockContent(name))
		}
	}

	return migration.Migration{
		Version: 2,
		Proc: func(s *state.State) {
			trans := template.ByLayout(
				layout,
				tmpls...,
			)
			for _, t := range trans {
				t(s)
			}
		},
	}
}

func (p *Page) StateMigrations() migration.Migrations {
	return migration.MakeMigrations(
		[]migration.Migration{
			{
				Version: 2,
				Proc: func(s *state.State) {
					b := simple.New(&model.Block{
						Content: &model.BlockContentOfText{
							Text: &model.BlockContentText{
								Text: "Test 1 " + p.Id(),
							},
						},
					})

					s.Add(b)
					s.InsertTo("", model.Block_Inner, b.Model().Id)
				},
			},
			{
				Version: 3,
				Proc: func(s *state.State) {
					b := simple.New(&model.Block{
						Content: &model.BlockContentOfText{
							Text: &model.BlockContentText{
								Text: "Test 2 " + p.Id(),
							},
						},
					})

					s.Add(b)
					s.InsertTo("", model.Block_Inner, b.Model().Id)
				},
			},
		})
}
