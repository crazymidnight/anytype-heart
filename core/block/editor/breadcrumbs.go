package editor

import (
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/meta"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	pbrelation "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"
)

var log = logging.Logger("anytype-mw-editor")

func NewBreadcrumbs(m meta.Service) *Breadcrumbs {
	return &Breadcrumbs{
		SmartBlock: smartblock.New(m),
	}
}

type Breadcrumbs struct {
	smartblock.SmartBlock
}

func (p *Breadcrumbs) Init(s source.Source, allowEmpty bool, _ []string) (err error) {
	if err = p.SmartBlock.Init(s, true, nil); err != nil {
		return
	}
	p.SmartBlock.DisableLayouts()
	return template.ApplyTemplate(p, nil, template.WithEmpty)
}

func (p *Breadcrumbs) Relations() []*pbrelation.Relation {
	return nil
}

func (b *Breadcrumbs) SetCrumbs(ids []string) (err error) {
	s := b.NewState()
	var existingLinks = make(map[string]string)
	s.Iterate(func(b simple.Block) (isContinue bool) {
		if link := b.Model().GetLink(); link != nil {
			existingLinks[link.TargetBlockId] = b.Model().Id
		}
		return true
	})
	root := s.Get(s.RootId()).Model()
	root.ChildrenIds = make([]string, 0, len(ids))
	for _, id := range ids {
		linkId, ok := existingLinks[id]
		if !ok {
			link := simple.New(&model.Block{
				Content: &model.BlockContentOfLink{
					Link: &model.BlockContentLink{
						TargetBlockId: id,
						Style:         model.BlockContentLink_Page,
					},
				},
			})
			s.Add(link)
			linkId = link.Model().Id
		}
		root.ChildrenIds = append(root.ChildrenIds, linkId)
	}
	return b.Apply(s)
}
