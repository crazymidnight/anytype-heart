package editor

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/block/database"
	"github.com/anytypeio/go-anytype-middleware/core/block/database/objects"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/meta"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	simpleDataview "github.com/anytypeio/go-anytype-middleware/core/block/simple/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/google/uuid"
)

var ErrAlreadyHasDataviewBlock = fmt.Errorf("already has the dataview block")

func NewSet(ms meta.Service, dbCtrl database.Ctrl) *Set {
	sb := &Set{
		SmartBlock: smartblock.New(ms, objects.BundledObjectTypeURLPrefix+"set"),
	}

	sb.Basic = basic.NewBasic(sb)
	sb.IHistory = basic.NewHistory(sb)
	sb.Dataview = dataview.NewDataview(sb, dbCtrl)
	sb.Router = database.New(dbCtrl)
	return sb
}

type Set struct {
	smartblock.SmartBlock
	basic.Basic
	basic.IHistory
	dataview.Dataview
	database.Router
}

func (p *Set) Init(s source.Source, allowEmpty bool, _ []string) (err error) {
	err = p.SmartBlock.Init(s, true, nil)
	if err != nil {
		return err
	}

	if err = template.ApplyTemplate(p, template.WithTitle, nil); err != nil {
		return
	}
	if p.Id() == p.Anytype().PredefinedBlocks().SetPages {
		return p.initPagesSet()
	}
	return
}

func (p *Set) initPagesSet() error {
	// init dataview
	relations := []*model.BlockContentDataviewRelation{{Key: "id", IsVisible: false, IsReadOnly: true}, {Key: "name", IsVisible: true}, {Key: "lastOpenedDate", IsVisible: true, IsReadOnly: true}, {Key: "lastModifiedDate", IsVisible: true, IsReadOnly: true}, {Key: "createdDate", IsVisible: true, IsReadOnly: true}}
	dataview := model.BlockContentOfDataview{
		Dataview: &model.BlockContentDataview{
			Source: "https://anytype.io/schemas/object/bundled/page",
			Views: []*model.BlockContentDataviewView{
				{
					Id:   uuid.New().String(),
					Type: model.BlockContentDataviewView_Table,
					Name: "All pages",
					Sorts: []*model.BlockContentDataviewSort{
						{
							RelationKey: "name",
							Type:        model.BlockContentDataviewSort_Asc,
						},
					},
					Relations: relations,
					Filters:   nil,
				},
			},
		},
	}

	err := p.InitDataview(dataview, "Pages", "📒")
	if err == ErrAlreadyHasDataviewBlock {
		return p.migrateOldSet()
	}

	return err
}

func (p *Set) migrateOldSet() error {
	return p.Iterate(func(b simple.Block) (isContinue bool) {
		if dvBlock, ok := b.(simpleDataview.Block); !ok {
			return true
		} else {
			if dvBlock.Model().GetDataview().Source == "" && dvBlock.Model().GetDataview().SchemaURL == "pages" {
				// migrate old pages set
				s := p.NewState()
				_ = dvBlock.SetSource("https://anytype.io/schemas/object/bundled/page")
				s.Set(dvBlock)
				p.Apply(s, smartblock.NoEvent)
			}
		}
		return true
	})
}

func (p *Set) InitDataview(blockContent model.BlockContentOfDataview, name string, icon string) error {
	s := p.NewState()
	err := p.SetDetails(s.Context(), []*pb.RpcBlockSetDetailsDetail{
		{Key: "name", Value: pbtypes.String(name)},
		{Key: "iconEmoji", Value: pbtypes.String(icon)},
	})
	if err != nil {
		return err
	}

	if !s.IsEmpty() {
		return ErrAlreadyHasDataviewBlock
	}

	// use fixed id, because it should be the only one block
	dw := simple.New(&model.Block{Content: &blockContent, Id: "dataview"})
	s.Add(dw)

	if err = s.InsertTo(template.HeaderLayoutId, model.Block_Bottom, dw.Model().Id); err != nil {
		return fmt.Errorf("can't insert dataview: %v", err)
	}

	log.Infof("create default structure for set: %v", s.RootId())

	return p.Apply(s, smartblock.NoEvent, smartblock.NoHistory)
}
