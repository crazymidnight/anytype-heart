package history

import (
	"context"
	"time"

	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/any-sync/commonspace"
	"github.com/anytypeio/any-sync/commonspace/object/tree/objecttree"
	"github.com/gogo/protobuf/proto"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/block/source"
	"github.com/anytypeio/go-anytype-middleware/core/relation"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/logging"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/space"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/anytypeio/go-anytype-middleware/util/slice"
)

const CName = "history"

const versionGroupInterval = time.Minute * 5

var log = logging.Logger("anytype-mw-history")

func New() History {
	return new(history)
}

type History interface {
	Show(pageId, versionId string) (bs *model.ObjectView, ver *pb.RpcHistoryVersion, err error)
	Versions(pageId, lastVersionId string, limit int) (resp []*pb.RpcHistoryVersion, err error)
	SetVersion(pageId, versionId string) (err error)
	app.Component
}

type BlockService interface {
	ResetToState(pageId string, s *state.State) (err error)
}

type history struct {
	a               core.Service
	blockService    BlockService
	objectStore     objectstore.ObjectStore
	relationService relation.Service
	spaceService    space.Service
}

func (h *history) Init(a *app.App) (err error) {
	h.a = a.MustComponent(core.CName).(core.Service)
	h.blockService = a.MustComponent(block.CName).(BlockService)
	h.objectStore = a.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	h.relationService = a.MustComponent(relation.CName).(relation.Service)
	h.spaceService = a.MustComponent(space.CName).(space.Service)
	return
}

func (h *history) Name() (name string) {
	return CName
}

func (h *history) Show(pageId, versionId string) (bs *model.ObjectView, ver *pb.RpcHistoryVersion, err error) {
	s, sbType, ver, err := h.buildState(pageId, versionId)
	if err != nil {
		return
	}
	metaD, _ := h.objectStore.QueryById(s.DepSmartIds(true, true, false, true, false))
	details := make([]*model.ObjectViewDetailsSet, 0, len(metaD))
	var uniqueObjTypes []string

	metaD = append(metaD, database.Record{Details: s.CombinedDetails()})
	uniqueObjTypes = s.ObjectTypes()
	for _, m := range metaD {
		details = append(details, &model.ObjectViewDetailsSet{
			Id:      pbtypes.GetString(m.Details, bundle.RelationKeyId.String()),
			Details: m.Details,
		})

		if ot := pbtypes.GetString(m.Details, bundle.RelationKeyType.String()); ot != "" {
			if slice.FindPos(uniqueObjTypes, ot) == -1 {
				uniqueObjTypes = append(uniqueObjTypes, ot)
			}
		}
	}

	rels, _ := h.relationService.FetchLinks(s.PickRelationLinks())
	return &model.ObjectView{
		RootId:        pageId,
		Type:          model.SmartBlockType(sbType),
		Blocks:        s.Blocks(),
		Details:       details,
		RelationLinks: rels.RelationLinks(),
	}, ver, nil
}

func (h *history) Versions(pageId, lastVersionId string, limit int) (resp []*pb.RpcHistoryVersion, err error) {
	if limit <= 0 {
		limit = 100
	}
	profileId, profileName, err := h.getProfileInfo()
	if err != nil {
		return
	}
	var includeLastId = true

	reverse := func(vers []*pb.RpcHistoryVersion) []*pb.RpcHistoryVersion {
		for i, j := 0, len(vers)-1; i < j; i, j = i+1, j-1 {
			vers[i], vers[j] = vers[j], vers[i]
		}
		return vers
	}
	convert := func(decrypted []byte) (any, error) {
		ch := &pb.Change{}
		err = proto.Unmarshal(decrypted, ch)
		if err != nil {
			return nil, err
		}
		return ch, nil
	}

	for len(resp) < limit {
		tree, _, e := h.treeWithId(pageId, lastVersionId, includeLastId)
		if e != nil {
			return nil, e
		}
		var data []*pb.RpcHistoryVersion

		e = tree.IterateFrom(tree.Root().Id, convert, func(c *objecttree.Change) (isContinue bool) {
			data = append(data, &pb.RpcHistoryVersion{
				Id:          c.Id,
				PreviousIds: c.PreviousIds,
				AuthorId:    profileId,
				AuthorName:  profileName,
				Time:        c.Timestamp,
			})
			return true
		})
		if e != nil {
			return nil, e
		}
		if len(data[0].PreviousIds) == 0 {
			if data[0].Id == tree.Id() {
				data = data[1:]
			}
			resp = append(data, resp...)
			break
		} else {
			resp = append(data, resp...)
			lastVersionId = tree.Root().Id
			includeLastId = false
		}

		if len(data) == 0 {
			break
		}
	}

	resp = reverse(resp)

	var groupId int64
	var nextVersionTimestamp int64

	for i := 0; i < len(resp); i++ {
		if nextVersionTimestamp-resp[i].Time > int64(versionGroupInterval.Seconds()) {
			groupId++
		}
		nextVersionTimestamp = resp[i].Time
		resp[i].GroupId = groupId
	}

	return
}

func (h *history) SetVersion(pageId, versionId string) (err error) {
	s, _, _, err := h.buildState(pageId, versionId)
	if err != nil {
		return
	}
	return h.blockService.ResetToState(pageId, s)
}

func (h *history) treeWithId(id, beforeId string, includeBeforeId bool) (ht objecttree.HistoryTree, sbt smartblock.SmartBlockType, err error) {
	spc, err := h.spaceService.AccountSpace(context.Background())
	if err != nil {
		return
	}
	ht, err = spc.BuildHistoryTree(context.Background(), id, commonspace.HistoryTreeOpts{
		BeforeId: beforeId,
		Include:  includeBeforeId,
	})
	if err != nil {
		return
	}

	payload := &model.ObjectChangePayload{}
	err = proto.Unmarshal(ht.ChangeInfo().ChangePayload, payload)
	if err != nil {
		return
	}

	sbt = smartblock.SmartBlockType(payload.SmartBlockType)
	return
}

func (h *history) buildState(pageId, versionId string) (st *state.State, sbType smartblock.SmartBlockType, ver *pb.RpcHistoryVersion, err error) {
	tree, sbType, err := h.treeWithId(pageId, versionId, true)
	if err != nil {
		return
	}

	st, _, _, err = source.BuildState(nil, tree, h.a.PredefinedBlocks().Profile)
	if _, _, err = state.ApplyStateFast(st); err != nil {
		return
	}

	st.BlocksInit(st)
	if ch, e := tree.GetChange(versionId); e == nil {
		profileId, profileName, e := h.getProfileInfo()
		if e != nil {
			err = e
			return
		}
		ver = &pb.RpcHistoryVersion{
			Id:          ch.Id,
			PreviousIds: ch.PreviousIds,
			AuthorId:    profileId,
			AuthorName:  profileName,
			Time:        ch.Timestamp,
		}
	}
	return
}

func (h *history) getProfileInfo() (profileId, profileName string, err error) {
	profileId = h.a.ProfileID()
	lp, err := h.a.LocalProfile()
	if err != nil {
		return
	}
	profileName = lp.Name
	return
}
