package objectcreator

import (
	"context"
	"strings"
	"time"

	"github.com/anyproto/any-sync/commonspace/object/tree/treestorage"
	"github.com/gogo/protobuf/types"
	"golang.org/x/exp/slices"

	"github.com/anyproto/anytype-heart/core/block/editor/lastused"
	"github.com/anyproto/anytype-heart/core/block/editor/smartblock"
	"github.com/anyproto/anytype-heart/core/block/editor/state"
	"github.com/anyproto/anytype-heart/core/block/object/objectcache"
	"github.com/anyproto/anytype-heart/core/domain"
	"github.com/anyproto/anytype-heart/metrics"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	coresb "github.com/anyproto/anytype-heart/pkg/lib/core/smartblock"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/addr"
	"github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	"github.com/anyproto/anytype-heart/space/clientspace"
	"github.com/anyproto/anytype-heart/util/pbtypes"
)

type eventKey int

const eventCreate eventKey = 0

type CreateOptions struct {
	payload *treestorage.TreeStorageCreatePayload
}

type CreateOption func(opts *CreateOptions)

func WithPayload(payload *treestorage.TreeStorageCreatePayload) CreateOption {
	return func(opts *CreateOptions) {
		opts.payload = payload
	}
}

// CreateSmartBlockFromState create new object from the provided `createState` and `details`.
// If you pass `details` into the function, it will automatically add missing relationLinks and override the details from the `createState`
// It will return error if some of the relation keys in `details` not installed in the workspace.
func (s *service) CreateSmartBlockFromState(
	ctx context.Context, spaceID string, objectTypeKeys []domain.TypeKey, createState *state.State,
) (id string, newDetails *types.Struct, err error) {
	spc, err := s.spaceService.Get(ctx, spaceID)
	if err != nil {
		return "", nil, err
	}
	return s.CreateSmartBlockFromStateInSpace(ctx, spc, objectTypeKeys, createState)
}

func (s *service) CreateSmartBlockFromStateInSpace(
	ctx context.Context, spc clientspace.Space, objectTypeKeys []domain.TypeKey, createState *state.State,
) (id string, newDetails *types.Struct, err error) {
	return s.CreateSmartBlockFromStateInSpaceWithOptions(ctx, spc, objectTypeKeys, createState)
}

func (s *service) CreateSmartBlockFromStateInSpaceWithOptions(
	ctx context.Context, spc clientspace.Space, objectTypeKeys []domain.TypeKey, createState *state.State, opts ...CreateOption,
) (id string, newDetails *types.Struct, err error) {
	if createState == nil {
		createState = state.NewDoc("", nil).(*state.State)
	}
	startTime := time.Now()
	// priority:
	// 1. details
	// 2. createState
	// 3. createState details
	// 4. default object type by smartblock type
	if len(objectTypeKeys) == 0 {
		objectTypeKeys = []domain.TypeKey{bundle.TypeKeyPage}
	}
	sbType := objectTypeKeysToSmartBlockType(objectTypeKeys)

	createState.SetDetailAndBundledRelation(bundle.RelationKeySpaceId, pbtypes.String(spc.Id()))

	ev := &metrics.CreateObjectEvent{
		SetDetailsMs: time.Since(startTime).Milliseconds(),
	}

	ctx = context.WithValue(ctx, eventCreate, ev)
	initFunc := func(id string) *smartblock.InitContext {
		createState.SetRootId(id)
		return &smartblock.InitContext{
			Ctx:            ctx,
			ObjectTypeKeys: objectTypeKeys,
			State:          createState,
			RelationKeys:   generateRelationKeysFromState(createState),
			SpaceID:        spc.Id(),
		}
	}

	sb, err := createSmartBlock(ctx, spc, initFunc, createState, sbType, opts...)
	if err != nil {
		return "", nil, err
	}

	sb.Lock()
	newDetails = sb.CombinedDetails()
	sb.Unlock()
	id = sb.Id()

	if pbtypes.GetInt64(newDetails, bundle.RelationKeyLastUsedDate.String()) == 0 {
		uk := pbtypes.GetString(newDetails, bundle.RelationKeyUniqueKey.String())
		switch sbType {
		case coresb.SmartBlockTypeObjectType:
			lastused.UpdateLastUsedDate(spc, s.objectStore, domain.TypeKey(strings.TrimPrefix(uk, addr.ObjectTypeKeyToIdPrefix)))
		case coresb.SmartBlockTypeRelation:
			lastused.UpdateLastUsedDate(spc, s.objectStore, domain.RelationKey(strings.TrimPrefix(uk, addr.RelationKeyToIdPrefix)))
		}
	} else if pbtypes.GetInt64(newDetails, bundle.RelationKeyOrigin.String()) == int64(model.ObjectOrigin_none) {
		lastused.UpdateLastUsedDate(spc, s.objectStore, objectTypeKeys[0])
	}

	ev.SmartblockCreateMs = time.Since(startTime).Milliseconds() - ev.SetDetailsMs - ev.WorkspaceCreateMs - ev.GetWorkspaceBlockWaitMs
	ev.SmartblockType = int(sbType)
	ev.ObjectId = id
	metrics.Service.Send(ev)
	return id, newDetails, nil
}

func objectTypeKeysToSmartBlockType(typeKeys []domain.TypeKey) coresb.SmartBlockType {
	// TODO Add validation for types that user can't create

	if slices.Contains(typeKeys, bundle.TypeKeyTemplate) {
		return coresb.SmartBlockTypeTemplate
	}
	typeKey := typeKeys[0]

	switch typeKey {
	case bundle.TypeKeyObjectType:
		return coresb.SmartBlockTypeObjectType
	case bundle.TypeKeyRelation:
		return coresb.SmartBlockTypeRelation
	case bundle.TypeKeyRelationOption:
		return coresb.SmartBlockTypeRelationOption
	case bundle.TypeKeyFile, bundle.TypeKeyImage, bundle.TypeKeyAudio, bundle.TypeKeyVideo:
		return coresb.SmartBlockTypeFileObject
	default:
		return coresb.SmartBlockTypePage
	}
}

func createSmartBlock(
	ctx context.Context, spc clientspace.Space, initFunc objectcache.InitFunc, st *state.State, sbType coresb.SmartBlockType, opts ...CreateOption,
) (smartblock.SmartBlock, error) {
	if uKey := st.UniqueKeyInternal(); uKey != "" {
		uk, err := domain.NewUniqueKey(sbType, uKey)
		if err != nil {
			return nil, err
		}
		if sbType == coresb.SmartBlockTypeFileObject {
			return spc.DeriveTreeObjectWithAccountSignature(ctx, objectcache.TreeDerivationParams{
				Key:      uk,
				InitFunc: initFunc,
			})
		} else {
			return spc.DeriveTreeObject(ctx, objectcache.TreeDerivationParams{
				Key:      uk,
				InitFunc: initFunc,
			})
		}
	}

	createOpts := &CreateOptions{}
	for _, opt := range opts {
		opt(createOpts)
	}
	if createOpts.payload != nil {
		return spc.CreateTreeObjectWithPayload(ctx, *createOpts.payload, initFunc)
	}

	return spc.CreateTreeObject(ctx, objectcache.TreeCreationParams{
		Time:           time.Now(),
		SmartblockType: sbType,
		InitFunc:       initFunc,
	})
}

func generateRelationKeysFromState(st *state.State) (relationKeys []string) {
	if st == nil {
		return
	}
	details := st.Details().GetFields()
	localDetails := st.LocalDetails().GetFields()
	relationKeys = make([]string, 0, len(details)+len(localDetails))
	for k := range details {
		relationKeys = append(relationKeys, k)
	}
	for k := range localDetails {
		relationKeys = append(relationKeys, k)
	}
	return
}
