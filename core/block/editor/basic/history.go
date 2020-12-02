package basic

import (
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
)

type IHistory interface {
	Undo(*state.Context) (err error)
	Redo(*state.Context) (err error)
}

func NewHistory(sb smartblock.SmartBlock) IHistory {
	return &history{sb}
}

type history struct {
	smartblock.SmartBlock
}

func (h *history) Undo(ctx *state.Context) (err error) {
	s := h.NewStateCtx(ctx)
	action, err := h.History().Previous()
	if err != nil {
		return
	}

	for _, b := range action.Add {
		s.Unlink(b.Model().Id)
	}
	for _, b := range action.Remove {
		s.Set(b.Copy())
	}
	for _, b := range action.Change {
		s.Set(b.Before.Copy())
	}
	if action.Details != nil {
		s.SetDetails(pbtypes.CopyStruct(action.Details.Before))
	}
	return h.Apply(s, smartblock.NoHistory)
}

func (h *history) Redo(ctx *state.Context) (err error) {
	s := h.NewStateCtx(ctx)
	action, err := h.History().Next()
	if err != nil {
		return
	}

	for _, b := range action.Add {
		s.Set(b.Copy())
	}
	for _, b := range action.Remove {
		s.Unlink(b.Model().Id)
	}
	for _, b := range action.Change {
		s.Set(b.After.Copy())
	}
	if action.Details != nil {
		s.SetDetails(pbtypes.CopyStruct(action.Details.After))
	}
	return h.Apply(s, smartblock.NoHistory)
}
