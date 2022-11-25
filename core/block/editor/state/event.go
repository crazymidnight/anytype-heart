package state

import (
	"fmt"

	"github.com/gogo/protobuf/types"

	"github.com/anytypeio/go-anytype-middleware/core/block/simple"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/base"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/dataview"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/file"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/latex"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/link"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/relation"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/table"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/text"
	"github.com/anytypeio/go-anytype-middleware/core/relation/relationutils"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

func (s *State) applyEvent(ev *pb.EventMessage) (err error) {
	var apply = func(id string, f func(b simple.Block) error) (err error) {
		if b := s.Get(id); b != nil {
			return f(b)
		}
		return fmt.Errorf("can't apply change: block not found")
	}
	switch o := ev.Value.(type) {
	case *pb.EventMessageValueOfBlockSetAlign:
		if err = apply(o.BlockSetAlign.Id, func(b simple.Block) error {
			b.Model().Align = o.BlockSetAlign.Align
			return nil
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetVerticalAlign:
		if err = apply(o.BlockSetVerticalAlign.Id, func(b simple.Block) error {
			b.Model().VerticalAlign = o.BlockSetVerticalAlign.VerticalAlign
			return nil
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetBackgroundColor:
		if err = apply(o.BlockSetBackgroundColor.Id, func(b simple.Block) error {
			b.Model().BackgroundColor = o.BlockSetBackgroundColor.BackgroundColor
			return nil
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetBookmark:
		if err = apply(o.BlockSetBookmark.Id, func(b simple.Block) error {
			if bm, ok := b.(bookmark.Block); ok {
				return bm.ApplyEvent(o.BlockSetBookmark)
			}
			return fmt.Errorf("not a bookmark block")
		}); err != nil {
			return
		}

	case *pb.EventMessageValueOfBlockSetTableRow:
		if err = apply(o.BlockSetTableRow.Id, func(b simple.Block) error {
			if tr, ok := b.(table.RowBlock); ok {
				return tr.ApplyEvent(o.BlockSetTableRow)
			}
			return fmt.Errorf("not a table row block")
		}); err != nil {
			return
		}

	case *pb.EventMessageValueOfBlockSetDiv:
		if err = apply(o.BlockSetDiv.Id, func(b simple.Block) error {
			if d, ok := b.(base.DivBlock); ok {
				return d.ApplyEvent(o.BlockSetDiv)
			}
			return fmt.Errorf("not a div block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetText:
		if err = apply(o.BlockSetText.Id, func(b simple.Block) error {
			if t, ok := b.(text.Block); ok {
				return t.ApplyEvent(o.BlockSetText)
			}
			return fmt.Errorf("not a text block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetFields:
		if err = apply(o.BlockSetFields.Id, func(b simple.Block) error {
			b.Model().Fields = o.BlockSetFields.Fields
			return nil
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetFile:
		if err = apply(o.BlockSetFile.Id, func(b simple.Block) error {
			if f, ok := b.(file.Block); ok {
				return f.ApplyEvent(o.BlockSetFile)
			}
			return fmt.Errorf("not a file block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetLink:
		if err = apply(o.BlockSetLink.Id, func(b simple.Block) error {
			if f, ok := b.(link.Block); ok {
				return f.ApplyEvent(o.BlockSetLink)
			}
			return fmt.Errorf("not a link block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataviewSourceSet:
		if err = apply(o.BlockDataviewSourceSet.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok {
				url, migrated := relationutils.MigrateObjectTypeIds(o.BlockDataviewSourceSet.Source)
				if len(migrated) > 0 {
					s.SetObjectTypesToMigrate(append(s.ObjectTypesToMigrate(), migrated...))
					o.BlockDataviewSourceSet.Source = url
				}
				return f.SetSource(o.BlockDataviewSourceSet.Source)
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataviewViewSet:
		if err = apply(o.BlockDataviewViewSet.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok && o.BlockDataviewViewSet.View != nil {
				if f.SetView(o.BlockDataviewViewSet.ViewId, *o.BlockDataviewViewSet.View) != nil {
					f.AddView(*o.BlockDataviewViewSet.View)
				}
				return nil
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataviewViewOrder:
		if err = apply(o.BlockDataviewViewOrder.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok {
				f.SetViewOrder(o.BlockDataviewViewOrder.ViewIds)
				return nil
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataviewViewDelete:
		if err = apply(o.BlockDataviewViewDelete.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok {
				err := f.DeleteView(o.BlockDataviewViewDelete.ViewId)
				if err != nil && err != dataview.ErrViewNotFound {
					return err
				}
				return nil
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataviewOldRelationSet:
		if err = apply(o.BlockDataviewOldRelationSet.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok && o.BlockDataviewOldRelationSet.Relation != nil {
				if er := f.UpdateRelationOld(o.BlockDataviewOldRelationSet.RelationKey, *o.BlockDataviewOldRelationSet.Relation); er == dataview.ErrRelationNotFound {
					rel := o.BlockDataviewOldRelationSet.Relation
					f.AddRelationOld(*rel)
					// MIGRATION: reinterpretation of old changes as new changes
					f.AddRelation(&model.RelationLink{
						Key:    rel.Key,
						Format: rel.Format,
					})
				} else {
					return er
				}
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}

	case *pb.EventMessageValueOfBlockDataviewOldRelationDelete:
		if err = apply(o.BlockDataviewOldRelationDelete.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok {
				err = f.DeleteRelationOld(o.BlockDataviewOldRelationDelete.RelationKey)
				if err != nil {
					return err
				}
				// MIGRATION: reinterpretation of old changes as new changes
				f.DeleteRelation(o.BlockDataviewOldRelationDelete.RelationKey)
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataviewRelationSet:
		if err = apply(o.BlockDataviewRelationSet.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok {
				for _, rel := range o.BlockDataviewRelationSet.RelationLinks {
					f.AddRelation(rel)
				}
				return nil
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataviewRelationDelete:
		if err = apply(o.BlockDataviewRelationDelete.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok {
				for _, key := range o.BlockDataviewRelationDelete.RelationKeys {
					// todo: implement DeleteRelations?
					f.DeleteRelation(key)
				}
				return nil
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetRelation:
		if err = apply(o.BlockSetRelation.Id, func(b simple.Block) error {
			if f, ok := b.(relation.Block); ok {
				return f.ApplyEvent(o.BlockSetRelation)
			}
			return fmt.Errorf("not a relation block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockSetLatex:
		if err = apply(o.BlockSetLatex.Id, func(b simple.Block) error {
			if f, ok := b.(latex.Block); ok {
				return f.ApplyEvent(o.BlockSetLatex)
			}
			return fmt.Errorf("not a latex block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataViewGroupOrderUpdate:
		if err = apply(o.BlockDataViewGroupOrderUpdate.Id, func(b simple.Block) error {
			if f, ok := b.(dataview.Block); ok {
				f.SetViewGroupOrder(o.BlockDataViewGroupOrderUpdate.GroupOrder)
				return nil
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	case *pb.EventMessageValueOfBlockDataViewObjectOrderUpdate:
		event := o.BlockDataViewObjectOrderUpdate
		if err = apply(event.Id, func(b simple.Block) error {
			if dvBlock, ok := b.(dataview.Block); ok {
				var existOrder []string
				for _, order := range dvBlock.Model().GetDataview().ObjectOrders {
					if order.ViewId == event.ViewId && order.GroupId == event.GroupId {
						existOrder = order.ObjectIds
					}
				}

				changes := o.BlockDataViewObjectOrderUpdate.GetSliceChanges()
				changedIds := slice.ApplyChanges(existOrder, pbtypes.EventsToSliceChange(changes))

				dvBlock.SetViewObjectOrder([]*model.BlockContentDataviewObjectOrder{
					{ViewId: event.ViewId, GroupId: event.GroupId, ObjectIds: changedIds},
				})

				return nil
			}
			return fmt.Errorf("not a dataview block")
		}); err != nil {
			return
		}
	}

	return nil
}

func WrapEventMessages(virtual bool, msgs []*pb.EventMessage) []simple.EventMessage {
	var wmsgs []simple.EventMessage
	for i := range msgs {
		wmsgs = append(wmsgs, simple.EventMessage{
			Virtual: virtual,
			Msg:     msgs[i],
		})
	}
	return wmsgs
}

func StructDiffIntoEvents(contextId string, diff *types.Struct) (msgs []*pb.EventMessage) {
	return StructDiffIntoEventsWithSubIds(contextId, diff, nil, nil)
}

// StructDiffIntoEvents converts map into events. nil map value converts to Remove event
func StructDiffIntoEventsWithSubIds(contextId string, diff *types.Struct, keys []string, subIds []string) (msgs []*pb.EventMessage) {
	if diff == nil || len(diff.Fields) == 0 {
		return nil
	}
	var (
		removed []string
		details []*pb.EventObjectDetailsAmendKeyValue
	)

	for k, v := range diff.Fields {
		if len(keys) > 0 && slice.FindPos(keys, k) == -1 {
			continue
		}
		if v == nil {
			removed = append(removed, k)
			continue
		}
		details = append(details, &pb.EventObjectDetailsAmendKeyValue{Key: k, Value: v})
	}
	if len(details) > 0 {
		msgs = append(msgs, &pb.EventMessage{
			Value: &pb.EventMessageValueOfObjectDetailsAmend{
				ObjectDetailsAmend: &pb.EventObjectDetailsAmend{
					Id:      contextId,
					Details: details,
					SubIds:  subIds,
				},
			},
		})
	}
	if len(removed) > 0 {
		msgs = append(msgs, &pb.EventMessage{
			Value: &pb.EventMessageValueOfObjectDetailsUnset{
				ObjectDetailsUnset: &pb.EventObjectDetailsUnset{
					Id:     contextId,
					Keys:   removed,
					SubIds: subIds,
				},
			},
		})
	}

	return
}
