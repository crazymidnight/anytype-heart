package basic

import (
	"errors"
	"fmt"

	"github.com/gogo/protobuf/types"

	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/restriction"
	"github.com/anyproto/anytype-heart/core/block/uniquekey"
	"github.com/anyproto/anytype-heart/core/session"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/util/internalflag"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

var log = logging.Logger("anytype-mw-editor-basic")

type detailUpdate struct {
	key   string
	value *types.Value
}

func (bs *basic) SetDetails(ctx session.Context, details []*pb.RpcObjectSetDetailsDetail, showEvent bool) (err error) {
	s := bs.NewStateCtx(ctx)

	// Collect updates handling special cases. These cases could update details themselves, so we
	// have to apply changes later
	updates := bs.collectDetailUpdates(details, s)
	newDetails := applyDetailUpdates(s.CombinedDetails(), updates)
	s.SetDetails(newDetails)

	if err = bs.Apply(s, smartblock.NoRestrictions); err != nil {
		return
	}

	bs.discardOwnSetDetailsEvent(ctx, showEvent)
	return nil
}

func (bs *basic) collectDetailUpdates(details []*pb.RpcObjectSetDetailsDetail, s *state.State) []*detailUpdate {
	updates := make([]*detailUpdate, 0, len(details))
	for _, detail := range details {
		update, err := bs.createDetailUpdate(s, detail)
		if err == nil {
			updates = append(updates, update)
		} else {
			log.Errorf("can't set detail %s: %s", detail.Key, err)
		}
	}
	return updates
}

func applyDetailUpdates(oldDetails *types.Struct, updates []*detailUpdate) *types.Struct {
	newDetails := pbtypes.CopyStruct(oldDetails)
	if newDetails == nil || newDetails.Fields == nil {
		newDetails = &types.Struct{
			Fields: make(map[string]*types.Value),
		}
	}
	for _, update := range updates {
		if update.value == nil {
			delete(newDetails.Fields, update.key)
		} else {
			newDetails.Fields[update.key] = update.value
		}
	}
	return newDetails
}

func (bs *basic) createDetailUpdate(st *state.State, detail *pb.RpcObjectSetDetailsDetail) (*detailUpdate, error) {
	if detail.Value != nil {
		if err := pbtypes.ValidateValue(detail.Value); err != nil {
			return nil, fmt.Errorf("detail %s validation error: %s", detail.Key, err.Error())
		}
		if err := bs.setDetailSpecialCases(st, detail); err != nil {
			return nil, fmt.Errorf("special case: %w", err)
		}
		if err := bs.addRelationLink(detail.Key, st); err != nil {
			return nil, err
		}
		if err := bs.relationService.ValidateFormat(st.SpaceID(), detail.Key, detail.Value); err != nil {
			return nil, fmt.Errorf("failed to validate relation: %w", err)
		}
	}
	return &detailUpdate{
		key:   detail.Key,
		value: detail.Value,
	}, nil
}

func (bs *basic) setDetailSpecialCases(st *state.State, detail *pb.RpcObjectSetDetailsDetail) error {
	// TODO Decide if we need this case
	// if detail.Key == bundle.RelationKeyType.String() {
	// 	// special case when client sets the type's detail directly instead of using setObjectType command
	// 	return bs.SetObjectTypesInState(st, pbtypes.GetStringListValue(detail.Value))
	// }
	if detail.Key == bundle.RelationKeyLayout.String() {
		// special case when client sets the layout detail directly instead of using SetLayoutInState command
		return bs.SetLayoutInState(st, model.ObjectTypeLayout(detail.Value.GetNumberValue()))
	}
	return nil
}

func (bs *basic) addRelationLink(relationKey string, st *state.State) error {
	// TODO: add relation.WithWorkspaceId(workspaceId) filter
	rel, err := bs.relationService.FetchRelationByKey(st.SpaceID(), relationKey)
	if err != nil || rel == nil {
		return fmt.Errorf("failed to get relation: %w", err)
	}
	st.AddRelationLinks(&model.RelationLink{
		Format: rel.Format,
		Key:    rel.Key,
	})
	return nil
}

func (bs *basic) discardOwnSetDetailsEvent(ctx session.Context, showEvent bool) {
	if !showEvent && ctx != nil {
		var filtered []*pb.EventMessage
		msgs := ctx.GetMessages()
		var isFiltered bool
		for i, msg := range msgs {
			if sd := msg.GetObjectDetailsSet(); sd == nil || sd.Id != bs.Id() {
				filtered = append(filtered, msgs[i])
			} else {
				isFiltered = true
			}
		}
		if isFiltered {
			ctx.SetMessages(bs.Id(), filtered)
		}
	}
}

func (bs *basic) SetLayout(ctx session.Context, layout model.ObjectTypeLayout) (err error) {
	if err = bs.Restrictions().Object.Check(model.Restrictions_LayoutChange); err != nil {
		return
	}

	s := bs.NewStateCtx(ctx)
	if err = bs.SetLayoutInState(s, layout); err != nil {
		return
	}
	return bs.Apply(s, smartblock.NoRestrictions)
}

func (bs *basic) SetObjectTypes(ctx session.Context, objectTypeKeys []bundle.TypeKey) (err error) {
	s := bs.NewStateCtx(ctx)
	if err = bs.SetObjectTypesInState(s, objectTypeKeys); err != nil {
		return
	}

	flags := internalflag.NewFromState(s)
	flags.Remove(model.InternalFlag_editorSelectType)
	flags.AddToState(s)

	// send event here to send updated details to client
	if err = bs.Apply(s, smartblock.NoRestrictions); err != nil {
		return
	}
	return
}

func (bs *basic) SetObjectTypesInState(s *state.State, objectTypeKeys []bundle.TypeKey) (err error) {
	if len(objectTypeKeys) == 0 {
		return fmt.Errorf("you must provide at least 1 object type")
	}
	if len(objectTypeKeys) > 1 {
		//nolint:govet
		log.With("objectID", s.RootId()).Warnf("set object types: more than one object type, setting layout to the first one")
	}

	if err = bs.Restrictions().Object.Check(model.Restrictions_TypeChange); errors.Is(err, restriction.ErrRestricted) {
		return fmt.Errorf("objectType change is restricted for object '%s': %v", bs.Id(), err)
	}

	prevTypeID := pbtypes.GetString(s.LocalDetails(), bundle.RelationKeyType.String())
	// nolint:errcheck
	prevType, _ := bs.objectStore.GetObjectType(prevTypeID)

	s.SetObjectTypeKeys(objectTypeKeys)

	toLayout, err := bs.getLayoutForType(objectTypeKeys[0])
	if err != nil {
		return fmt.Errorf("get layout for type %s: %w", objectTypeKeys[0], err)
	}
	if v := pbtypes.Get(s.Details(), bundle.RelationKeyLayout.String()); v == nil || // if layout is not set yet
		prevType == nil || // if we have no type set for some reason or it is missing
		float64(prevType.Layout) == v.GetNumberValue() { // or we have a objecttype recommended layout set for this object
		if err = bs.SetLayoutInState(s, toLayout); err != nil {
			return
		}
	}
	return
}

func (bs *basic) getLayoutForType(objectTypeKey bundle.TypeKey) (model.ObjectTypeLayout, error) {
	uk, err := uniquekey.New(model.SmartBlockType_STType, objectTypeKey.String())
	if err != nil {
		return 0, fmt.Errorf("create unique key: %w", err)
	}
	typeDetails, err := bs.objectStore.GetObjectByUniqueKey(bs.SpaceID(), uk)
	if err != nil {
		return 0, fmt.Errorf("get object by unique key: %w", err)
	}
	rawLayout := pbtypes.GetInt64(typeDetails.GetDetails(), bundle.RelationKeyLayout.String())
	return model.ObjectTypeLayout(rawLayout), nil
}

func (bs *basic) SetLayoutInState(s *state.State, toLayout model.ObjectTypeLayout) (err error) {
	if err = bs.Restrictions().Object.Check(model.Restrictions_LayoutChange); errors.Is(err, restriction.ErrRestricted) {
		return fmt.Errorf("layout change is restricted for object '%s': %v", bs.Id(), err)
	}

	return bs.SetLayoutInStateAndIgnoreRestriction(s, toLayout)
}

func (bs *basic) SetLayoutInStateAndIgnoreRestriction(s *state.State, toLayout model.ObjectTypeLayout) (err error) {
	fromLayout, _ := s.Layout()

	s.SetDetail(bundle.RelationKeyLayout.String(), pbtypes.Int64(int64(toLayout)))

	if err = bs.layoutConverter.Convert(s, fromLayout, toLayout); err != nil {
		return fmt.Errorf("convert layout: %w", err)
	}
	return nil
}
