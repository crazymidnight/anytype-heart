package database

import (
	"context"

	"github.com/anytypeio/go-anytype-middleware/core/block/database/objects"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/session"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	coresb "github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/gogo/protobuf/types"
)

type Ctrl interface {
	Anytype() core.Service

	SetDetails(ctx *session.Context, req pb.RpcObjectSetDetailsRequest) error
	GetRelations(objectId string) (relations []*model.Relation, err error)

	CreateSmartBlockFromTemplate(ctx context.Context, sbType coresb.SmartBlockType, details *types.Struct, relations []*model.Relation, templateId string) (id string, newDetails *types.Struct, err error)
	UpdateExtraRelations(ctx *session.Context, id string, relations []*model.Relation, createIfMissing bool) (err error)
	AddExtraRelations(ctx *session.Context, id string, relations []*model.Relation) (relationsWithKeys []*model.Relation, err error)
	RemoveExtraRelations(ctx *session.Context, id string, relationKeys []string) (err error)
	ModifyExtraRelations(ctx *session.Context, objectId string, modifier func(current []*model.Relation) ([]*model.Relation, error)) (err error)
	UpdateExtraRelationOption(ctx *session.Context, req pb.RpcObjectRelationOptionUpdateRequest) (err error)
	AddExtraRelationOption(ctx *session.Context, req pb.RpcObjectRelationOptionAddRequest) (option *model.RelationOption, err error)
	Do(id string, apply func(b smartblock.SmartBlock) error) error

	SetObjectTypes(ctx *session.Context, objectId string, objectTypes []string) (err error)
}

type Router interface {
	Get(objectType *model.ObjectType) (database.Database, error) // objectType can be empty, means it is not explicitly set
}

func New(s Ctrl) Router {
	return &router{s: s}
}

type router struct{ s Ctrl }

func (r router) Get(objectType *model.ObjectType) (database.Database, error) {
	// todo: wrap into iface
	setDetailsNoContext := func(req pb.RpcObjectSetDetailsRequest) error {
		return r.s.SetDetails(nil, req)
	}

	updateOptionNoContext := func(req pb.RpcObjectRelationOptionUpdateRequest) (opt *model.RelationOption, err error) {
		if req.Option.Id == "" {
			return r.s.AddExtraRelationOption(nil, pb.RpcObjectRelationOptionAddRequest{ContextId: req.ContextId, RelationKey: req.RelationKey, Option: req.Option})
		}

		return req.Option, r.s.UpdateExtraRelationOption(nil, req)
	}

	deleteOptionNoContext := func(id, relKey, optionId string) error {
		return r.s.Do(id, func(b smartblock.SmartBlock) error {
			err := b.DeleteExtraRelationOption(nil, relKey, optionId, true)
			if err != nil {
				return err
			}
			return nil
		})
	}

	modifyExtraRelationsNoContext := func(objectId string, modifier func(current []*model.Relation) ([]*model.Relation, error)) (err error) {
		return r.s.ModifyExtraRelations(nil, objectId, modifier)
	}

	setOrAddRelations := func(id string, relations []*model.Relation) error {
		newRels := pbtypes.CopyRelations(relations)
		return r.s.ModifyExtraRelations(nil, id, func(current []*model.Relation) ([]*model.Relation, error) {
			newRels = pbtypes.MergeRelationsDicts(current, newRels)
			for _, newRel := range newRels {
				newRel.Scope = model.Relation_object
			}

			return newRels, nil
		})
	}

	return objects.New(
		r.s.Anytype().ObjectStore(),
		objectType,
		setDetailsNoContext,
		r.s.GetRelations,
		setOrAddRelations,
		modifyExtraRelationsNoContext,
		updateOptionNoContext,
		deleteOptionNoContext,
		r.s.CreateSmartBlockFromTemplate,
	), nil
}
