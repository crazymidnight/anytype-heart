package restriction

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
)

func TestService_ObjectRestrictionsById(t *testing.T) {
	rest := New(nil)
	assert.ErrorIs(t, rest.GetRestrictions(&restrictionHolder{
		id: "",
		tp: model.SmartBlockType_AnytypeProfile,
	}).Object.Check(model.Restrictions_Blocks),
		ErrRestricted,
	)

	assert.ErrorIs(t, rest.GetRestrictions(&restrictionHolder{
		id:     "",
		tp:     model.SmartBlockType_Page,
		layout: model.ObjectType_collection,
	}).Object.Check(model.Restrictions_Blocks),
		ErrRestricted,
	)

	assert.NoError(t, rest.GetRestrictions(&restrictionHolder{
		id: "",
		tp: model.SmartBlockType_Page,
	}).Object.Check(model.Restrictions_Blocks))

	assert.NoError(t, rest.GetRestrictions(&restrictionHolder{
		id:     bundle.TypeKeyDailyPlan.URL(),
		tp:     model.SmartBlockType_SubObject,
		layout: model.ObjectType_objectType,
	}).Object.Check(
		model.Restrictions_Details,
		model.Restrictions_Delete,
	))

	assert.ErrorIs(t, rest.GetRestrictions(&restrictionHolder{
		id:     bundle.TypeKeyPage.URL(),
		tp:     model.SmartBlockType_SubObject,
		layout: model.ObjectType_objectType,
	}).Object.Check(
		model.Restrictions_Details,
		model.Restrictions_Delete,
	), ErrRestricted)

	assert.NoError(t, rest.GetRestrictions(&restrictionHolder{
		id:     bundle.RelationKeyIconOption.String(),
		tp:     model.SmartBlockType_SubObject,
		layout: model.ObjectType_relation,
	}).Object.Check(
		model.Restrictions_Delete,
		model.Restrictions_Relations,
		model.Restrictions_Details,
	))

	assert.ErrorIs(t, rest.GetRestrictions(&restrictionHolder{
		id:     bundle.RelationKeyName.URL(),
		tp:     model.SmartBlockType_SubObject,
		layout: model.ObjectType_relation,
	}).Object.Check(
		model.Restrictions_Delete,
		model.Restrictions_Relations,
		model.Restrictions_Details,
	), ErrRestricted)
}
