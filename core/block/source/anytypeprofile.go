package source

import (
	"context"

	"github.com/anytypeio/go-anytype-middleware/change"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
)

func NewAnytypeProfile(a core.Service, id string) (s Source) {
	return &anytypeProfile{
		id: id,
		a:  a,
	}
}

type anytypeProfile struct {
	id string
	a  core.Service
}

func (v *anytypeProfile) Id() string {
	return v.id
}

func (v *anytypeProfile) Anytype() core.Service {
	return v.a
}

func (v *anytypeProfile) Type() pb.SmartBlockType {
	return pb.SmartBlockType_ProfilePage
}

func (v *anytypeProfile) Virtual() bool {
	return false
}

func (v *anytypeProfile) getDetails() (p *types.Struct) {
	return &types.Struct{Fields: map[string]*types.Value{
		bundle.RelationKeyName.String():        pbtypes.String("Anytype"),
		bundle.RelationKeyDescription.String(): pbtypes.String("Authored by Anytype team"),
		bundle.RelationKeyIconImage.String():   pbtypes.String("bafybeigpvcr42wobigjm63hmx3223gktsz4hpwp3jmiwkpcawo7wtpimv4"),
		bundle.RelationKeyId.String():          pbtypes.String(v.id),
		bundle.RelationKeyLayout.String():      pbtypes.Float64(float64(relation.ObjectType_profile)),
	}}
}

func (v *anytypeProfile) ReadDoc(receiver ChangeReceiver, empty bool) (doc state.Doc, err error) {
	s := state.NewDoc(v.id, nil).(*state.State)

	d := v.getDetails()

	s.SetDetails(d)
	s.SetObjectType(bundle.TypeKeyProfile.URL())
	return s, nil
}

func (v *anytypeProfile) ReadMeta(_ ChangeReceiver) (doc state.Doc, err error) {
	s := &state.State{}
	d := v.getDetails()

	s.SetDetails(d)
	s.SetObjectType(bundle.TypeKeyProfile.URL())
	return s, nil
}

func (v *anytypeProfile) PushChange(params PushChangeParams) (id string, err error) {
	return "", nil
}

func (v *anytypeProfile) FindFirstChange(ctx context.Context) (c *change.Change, err error) {
	return nil, change.ErrEmpty
}

func (v *anytypeProfile) Close() (err error) {
	return
}
