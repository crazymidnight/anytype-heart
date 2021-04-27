package restriction

import (
	"testing"

	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/stretchr/testify/assert"
)

type testObj struct {
	id string
	tp pb.SmartBlockType
}

func (t testObj) Id() string {
	return t.id
}

func (t testObj) Type() pb.SmartBlockType {
	return t.tp
}

func TestService_ObjectRestrictionsById(t *testing.T) {
	rest := New()
	assert.Equal(t, ErrRestricted, rest.ObjectRestrictionsByObj(testObj{
		id: "",
		tp: pb.SmartBlockType_Breadcrumbs,
	}).Check(model.Restrictions_CreateBlock))
	assert.NoError(t, rest.ObjectRestrictionsByObj(testObj{
		id: "",
		tp: pb.SmartBlockType_Page,
	}).Check(model.Restrictions_CreateBlock))
}
