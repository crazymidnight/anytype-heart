package state

import (
	"fmt"
	"github.com/globalsign/mgo/bson"

	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

var (
	maxChildrenThreshold = 40
)

func (s *State) Normalize(withLayouts bool) (err error) {
	s.removeDuplicates()
	return s.normalize(withLayouts)
}

type Normalizable interface {
	Normalize(s *State) error
}

func (s *State) normalize(withLayouts bool) (err error) {
	if err = s.Iterate(func(b simple.Block) (isContinue bool) {
		return true
	}); err != nil {
		return err
	}
	// remove invalid children
	for _, b := range s.blocks {
		s.normalizeChildren(b)
	}

	for _, b := range s.blocks {
		if n, ok := b.(Normalizable); ok {
			if err := n.Normalize(s); err != nil {
				return fmt.Errorf("custom normalization for block %s: %w", b.Model().Id, err)
			}
		}
		if b.Model().Id == s.RootId() {
			s.normalizeSmartBlock(b)
		}
	}

	// remove empty layouts
	s.removeEmptyLayoutBlocks(s.blocks)
	if s.parent != nil {
		s.removeEmptyLayoutBlocks(s.parent.blocks)
	}
	// normalize rows
	for _, b := range s.blocks {
		if layout := b.Model().GetLayout(); layout != nil {
			s.normalizeLayoutRow(b)
		}
	}
	if withLayouts {
		return s.normalizeTree()
	}
	return
}

func (s *State) removeEmptyLayoutBlocks(blocks map[string]simple.Block) {
	for _, b := range blocks {
		if layout := b.Model().GetLayout(); layout != nil {
			if len(b.Model().ChildrenIds) == 0 {
				s.Unlink(b.Model().Id)
			}
			// load parent for checking
			s.GetParentOf(b.Model().Id)
		}
	}
}

func (s *State) normalizeChildren(b simple.Block) {
	m := b.Model()
	for _, cid := range m.ChildrenIds {
		if !s.Exists(cid) {
			m.ChildrenIds = slice.Remove(m.ChildrenIds, cid)
			s.normalizeChildren(b)
			return
		}
	}
}

func (s *State) normalizeLayoutRow(b simple.Block) {
	// remove empty layout
	if len(b.Model().ChildrenIds) == 0 {
		s.Unlink(b.Model().Id)
		return
	}
	if b.Model().GetLayout().Style != model.BlockContentLayout_Row {
		return
	}
	// one column - remove row
	if len(b.Model().ChildrenIds) == 1 {
		var (
			contentIds   []string
			removeColumn bool
		)
		column := s.Get(b.Model().ChildrenIds[0])
		if layout := column.Model().GetLayout(); layout != nil && layout.Style == model.BlockContentLayout_Column {
			contentIds = column.Model().ChildrenIds
			removeColumn = true
		} else {
			contentIds = append(contentIds, column.Model().Id)
		}
		s.InsertTo(b.Model().Id, model.Block_Replace, contentIds...)
		if removeColumn {
			s.Unlink(column.Model().Id)
		}
		s.Unlink(b.Model().Id)
		return
	}

	// reset columns width when count of row children was changed
	orig := s.PickOrigin(b.Model().Id)
	if orig != nil && len(orig.Model().ChildrenIds) != len(b.Model().ChildrenIds) {
		for _, chId := range b.Model().ChildrenIds {
			fields := s.Get(chId).Model().Fields
			if fields != nil && fields.Fields != nil && fields.Fields["width"] != nil {
				fields.Fields["width"] = pbtypes.Float64(0)
			}
		}
	}
}

func isDivLayout(m *model.Block) bool {
	if layout := m.GetLayout(); layout != nil && layout.Style == model.BlockContentLayout_Div {
		return true
	}
	return false
}

func (s *State) normalizeTree() (err error) {
	s.normalizeTreeBranch(s.RootId())
	const headerId = "header"
	if s.Pick(headerId) != nil && slice.FindPos(s.Pick(s.RootId()).Model().ChildrenIds, headerId) != 0 {
		s.Unlink(headerId)
		root := s.Get(s.RootId()).Model()
		root.ChildrenIds = append([]string{headerId}, root.ChildrenIds...)
	}
	return nil
}

func (s *State) normalizeTreeBranch(id string) {
	parentB := s.Pick(id)
	if parentB == nil {
		return
	}
	// corner cases for tables
	switch parentB.Model().GetLayout().GetStyle() {
	case model.BlockContentLayout_TableRows:
		return
	case model.BlockContentLayout_TableColumns:
		return
	}
	if parentB.Model().GetTableRow() != nil {
		return
	}

	parent := parentB.Model()
	if len(parent.ChildrenIds) > maxChildrenThreshold {
		if nextId := s.wrapChildrenToDiv(id); nextId != "" {
			s.normalizeTreeBranch(nextId)
			return
		}
	}
	for _, chId := range parent.ChildrenIds {
		s.normalizeTreeBranch(chId)
	}
}

func (s *State) wrapChildrenToDiv(id string) (nextId string) {
	parent := s.Get(id).Model()
	overflow := maxChildrenThreshold - len(parent.ChildrenIds)
	if isDivLayout(parent) {
		changes := false
		nextDiv := s.getNextDiv(id)
		if nextDiv == nil || len(nextDiv.Model().ChildrenIds)+overflow > maxChildrenThreshold {
			nextDiv = s.newDiv()
			s.Add(nextDiv)
			s.InsertTo(parent.Id, model.Block_Bottom, nextDiv.Model().Id)
			changes = true
		}
		for s.divBalance(parent, nextDiv.Model()) {
			parent = nextDiv.Model()
			nextDiv = s.getNextDiv(parent.Id)
			overflow = maxChildrenThreshold - len(parent.ChildrenIds)
			if nextDiv == nil || len(nextDiv.Model().ChildrenIds)+overflow > maxChildrenThreshold {
				nextDiv = s.newDiv()
				s.Add(nextDiv)
				s.InsertTo(parent.Id, model.Block_Bottom, nextDiv.Model().Id)
				changes = true
			}
		}
		if changes {
			return s.PickParentOf(parent.Id).Model().Id
		}
		return ""
	}

	div := s.newDiv()
	div.Model().ChildrenIds = parent.ChildrenIds
	s.Add(div)
	parent.ChildrenIds = []string{div.Model().Id}
	return parent.Id
}

func (s *State) divBalance(d1, d2 *model.Block) (overflow bool) {
	d1.ChildrenIds = append(d1.ChildrenIds, d2.ChildrenIds...)
	sum := len(d1.ChildrenIds)
	div := sum / 2
	if sum > maxChildrenThreshold*2 {
		div = maxChildrenThreshold / 2
	}

	d2.ChildrenIds = make([]string, len(d1.ChildrenIds[div:]))
	copy(d2.ChildrenIds, d1.ChildrenIds[div:])
	d1.ChildrenIds = d1.ChildrenIds[:div]
	return len(d2.ChildrenIds) > maxChildrenThreshold
}

func (s *State) newDiv() simple.Block {
	divId := fmt.Sprintf("div-%s", bson.NewObjectId().Hex())
	return simple.New(&model.Block{
		Id: divId,
		Content: &model.BlockContentOfLayout{
			Layout: &model.BlockContentLayout{Style: model.BlockContentLayout_Div},
		},
	})
}

func (s *State) getNextDiv(id string) simple.Block {
	if b := s.pickNextDiv(id); b != nil {
		return s.Get(b.Model().Id)
	}
	return nil
}

func (s *State) pickNextDiv(id string) simple.Block {
	parent := s.PickParentOf(id)
	if parent != nil {
		pm := parent.Model()
		pos := slice.FindPos(pm.ChildrenIds, id)
		if pos != -1 && pos < len(pm.ChildrenIds)-1 {
			b := s.Pick(pm.ChildrenIds[pos+1])
			if isDivLayout(b.Model()) {
				return b
			}
		}
	}
	return nil
}

func (s *State) removeDuplicates() {
	childrenIds := make(map[string]struct{})
	handledBlocks := make(map[string]struct{})
	s.Iterate(func(b simple.Block) (isContinue bool) {
		if _, ok := handledBlocks[b.Model().Id]; ok {
			return true
		}
		var delIdx []int
		for i, cid := range b.Model().ChildrenIds {
			if _, ok := childrenIds[cid]; ok {
				delIdx = append(delIdx, i)
			} else {
				childrenIds[cid] = struct{}{}
			}
		}
		if len(delIdx) > 0 {
			b = s.Get(b.Model().Id)
			chIds := b.Model().ChildrenIds
			for i, idx := range delIdx {
				idx = idx - i
				chIds = append(chIds[:idx], chIds[idx+1:]...)
			}
			b.Model().ChildrenIds = chIds
		}
		handledBlocks[b.Model().Id] = struct{}{}
		return true
	})
}

func (s *State) normalizeSmartBlock(b simple.Block) {
	if isBlockEmpty(b) {
		b.Model().Content = &model.BlockContentOfSmartblock{
			Smartblock: &model.BlockContentSmartblock{},
		}
	}
}

func isBlockEmpty(b simple.Block) bool {
	if b.Model().Content == nil {
		return true
	}
	smartBlock := b.Model().Content.(*model.BlockContentOfSmartblock)
	return smartBlock == nil || smartBlock.Smartblock == nil
}

func CleanupLayouts(s *State) (removedCount int) {
	var cleanup func(id string) []string
	cleanup = func(id string) (result []string) {
		b := s.Get(id)
		if b == nil {
			return
		}
		for _, chId := range b.Model().ChildrenIds {
			if chB := s.Pick(chId); chB != nil {
				if isDivLayout(chB.Model()) {
					removedCount++
					result = append(result, cleanup(chId)...)
				} else {
					result = append(result, chId)
					cleanup(chId)
				}
			}
		}
		b.Model().ChildrenIds = result
		return
	}
	cleanup(s.RootId())
	return
}
