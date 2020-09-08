package template

import (
	"github.com/anytypeio/go-anytype-library/pb/model"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/text"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
)

const (
	HeaderLayoutId = "header"
	TitleBlockId   = "title"
)

var Empty = state.NewDoc("tmpl_empty", map[string]simple.Block{
	"tmpl_empty": simple.New(&model.Block{
		Id: "tmpl_empty",
		Content: &model.BlockContentOfSmartblock{
			Smartblock: &model.BlockContentSmartblock{},
		},
	}),
})

var WithHeader = state.NewDoc("tmpl_header", map[string]simple.Block{
	"tmpl_header": simple.New(&model.Block{
		Id: "tmpl_header",
		Content: &model.BlockContentOfSmartblock{
			Smartblock: &model.BlockContentSmartblock{},
		},
		ChildrenIds: []string{HeaderLayoutId},
	}),
	HeaderLayoutId: simple.New(&model.Block{
		Id: HeaderLayoutId,
		Content: &model.BlockContentOfLayout{
			Layout: &model.BlockContentLayout{
				Style: model.BlockContentLayout_Header,
			},
		},
	}),
})

var WithTitle = state.NewDoc("tmpl_title", map[string]simple.Block{
	"tmpl_title": simple.New(&model.Block{
		Id: "tmpl_title",
		Content: &model.BlockContentOfSmartblock{
			Smartblock: &model.BlockContentSmartblock{},
		},
		ChildrenIds: []string{HeaderLayoutId},
	}),
	HeaderLayoutId: simple.New(&model.Block{
		Id: HeaderLayoutId,
		Content: &model.BlockContentOfLayout{
			Layout: &model.BlockContentLayout{
				Style: model.BlockContentLayout_Header,
			},
		},
		ChildrenIds: []string{TitleBlockId},
	}),
	TitleBlockId: simple.New(&model.Block{
		Id:      TitleBlockId,
		Content: &model.BlockContentOfText{Text: &model.BlockContentText{}},
		Fields: &types.Struct{
			Fields: map[string]*types.Value{
				text.DetailsKeyFieldName: pbtypes.String("name"),
			},
		},
	}),
})

func InitTemplate(sb smartblock.SmartBlock, tmpl state.Doc, s *state.State) (err error) {
	if !s.Exists(sb.RootId()) {
		tmpl.Iterate(func(b simple.Block) (isContinue bool) {
			b = b.Copy()
			if b.Model().Id == tmpl.RootId() {
				b.Model().Id = sb.Id()
			}
			s.Add(b)
			return true
		})
	} else {
		// migration to title
		if tmpl == WithTitle && !s.Exists(HeaderLayoutId) {
			tmpl.Iterate(func(b simple.Block) (isContinue bool) {
				b = b.Copy()
				if b.Model().Id != tmpl.RootId() {
					s.Add(b)
				}
				return true
			})
			s.Get(s.RootId()).Model().ChildrenIds = append([]string{HeaderLayoutId}, s.Get(s.RootId()).Model().ChildrenIds...)
		}
	}
	return
}

func ApplyTemplate(sb smartblock.SmartBlock, tmpl state.Doc, s *state.State) (err error) {
	if s == nil {
		s = sb.NewState()
	}
	if err = InitTemplate(sb, tmpl, s); err != nil {
		return
	}
	return sb.Apply(s, smartblock.NoHistory, smartblock.NoEvent)
}
