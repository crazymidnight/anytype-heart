package editor

import (
	"github.com/gogo/protobuf/types"

	"github.com/anytypeio/go-anytype-middleware/core/block/editor/basic"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/collection"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/converter"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/smartblock"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/template"
	"github.com/anytypeio/go-anytype-middleware/core/block/migration"
	"github.com/anytypeio/go-anytype-middleware/core/relation"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

type Dashboard struct {
	smartblock.SmartBlock
	basic.AllOperations
	collection.Collection

	DetailsModifier DetailsModifier
	objectStore     objectstore.ObjectStore
	anytype         core.Service
}

func NewDashboard(
	detailsModifier DetailsModifier,
	objectStore objectstore.ObjectStore,
	relationService relation.Service,
	anytype core.Service,
	layoutConverter converter.LayoutConverter,
) *Dashboard {
	sb := smartblock.New()
	return &Dashboard{
		SmartBlock:      sb,
		AllOperations:   basic.NewBasic(sb, objectStore, relationService, layoutConverter),
		Collection:      collection.NewCollection(sb),
		DetailsModifier: detailsModifier,
		objectStore:     objectStore,
		anytype:         anytype,
	}
}

func (p *Dashboard) Init(ctx *smartblock.InitContext) (err error) {
	if err = p.SmartBlock.Init(ctx); err != nil {
		return
	}
	p.DisableLayouts()
	p.AddHook(p.updateObjects, smartblock.HookAfterApply)
	return nil

}

func (p *Dashboard) CreationStateMigration(ctx *smartblock.InitContext) migration.Migration {
	return migration.Migration{
		Version: 1,
		Proc: func(st *state.State) {
			template.InitTemplate(st,
				template.WithObjectTypesAndLayout([]string{bundle.TypeKeyDashboard.URL()}, model.ObjectType_dashboard),
				template.WithEmpty,
				template.WithDetailName("Home"),
				template.WithDetailIconEmoji("🏠"),
				template.WithNoRootLink(p.anytype.PredefinedBlocks().Archive),
				template.WithRequiredRelations(),
				template.WithNoDuplicateLinks(),
			)
		},
	}
}

func (p *Dashboard) StateMigrations() migration.Migrations {
	return migration.MakeMigrations(nil)
}

func (p *Dashboard) updateObjects(_ smartblock.ApplyInfo) (err error) {
	favoritedIds, err := p.GetIds()
	if err != nil {
		return
	}

	records, _, err := p.objectStore.Query(nil, database.Query{
		Filters: []*model.BlockContentDataviewFilter{
			{
				RelationKey: bundle.RelationKeyIsFavorite.String(),
				Condition:   model.BlockContentDataviewFilter_Equal,
				Value:       pbtypes.Bool(true),
			},
		},
	})
	if err != nil {
		return
	}
	var storeFavoritedIds = make([]string, 0, len(records))
	for _, rec := range records {
		storeFavoritedIds = append(storeFavoritedIds, pbtypes.GetString(rec.Details, bundle.RelationKeyId.String()))
	}

	removedIds, addedIds := slice.DifferenceRemovedAdded(storeFavoritedIds, favoritedIds)
	for _, removedId := range removedIds {
		go func(id string) {
			if err := p.DetailsModifier.ModifyLocalDetails(id, func(current *types.Struct) (*types.Struct, error) {
				if current == nil || current.Fields == nil {
					current = &types.Struct{
						Fields: map[string]*types.Value{},
					}
				}
				current.Fields[bundle.RelationKeyIsFavorite.String()] = pbtypes.Bool(false)
				return current, nil
			}); err != nil {
				log.Errorf("favorite: can't set detail to object: %v", err)
			}
		}(removedId)
	}
	for _, addedId := range addedIds {
		go func(id string) {
			if err := p.DetailsModifier.ModifyLocalDetails(id, func(current *types.Struct) (*types.Struct, error) {
				if current == nil || current.Fields == nil {
					current = &types.Struct{
						Fields: map[string]*types.Value{},
					}
				}
				current.Fields[bundle.RelationKeyIsFavorite.String()] = pbtypes.Bool(true)
				return current, nil
			}); err != nil {
				log.Errorf("favorite: can't set detail to object: %v", err)
			}
		}(addedId)
	}
	return
}
