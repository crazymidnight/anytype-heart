package source

import (
	"github.com/anytypeio/go-anytype-middleware/core/anytype"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/meta"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/google/uuid"
)

func NewVirtual(a anytype.Service, m meta.Service, t pb.SmartBlockType) (s Source) {
	return &virtual{
		id:     uuid.New().String(),
		a:      a,
		meta:   m,
		sbType: t,
	}
}

type virtual struct {
	id     string
	a      anytype.Service
	meta   meta.Service
	sbType pb.SmartBlockType
}

func (v *virtual) Id() string {
	return v.id
}

func (v *virtual) Anytype() anytype.Service {
	return v.a
}

func (v *virtual) Meta() meta.Service {
	return v.meta
}

func (v *virtual) Type() pb.SmartBlockType {
	return v.sbType
}
func (v *virtual) ReadDoc() (doc state.Doc, err error) {
	return state.NewDoc(v.id, nil), nil
}

func (v *virtual) PushChange(st *state.State, changes ...*pb.ChangeContent) (id string, err error) {
	return "", nil
}

func (v *virtual) Close() (err error) {
	return
}
