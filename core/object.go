package core

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/go-naturaldate/v2"
	"github.com/araddon/dateparse"
	"github.com/gogo/protobuf/types"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	importer "github.com/anytypeio/go-anytype-middleware/core/block/import"
	"github.com/anytypeio/go-anytype-middleware/core/indexer"
	"github.com/anytypeio/go-anytype-middleware/core/relation"
	"github.com/anytypeio/go-anytype-middleware/core/subscription"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database/filter"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/space/typeprovider"
	"github.com/anytypeio/go-anytype-middleware/util/internalflag"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
)

// To be renamed to ObjectSetDetails
func (mw *Middleware) ObjectSetDetails(cctx context.Context, req *pb.RpcObjectSetDetailsRequest) *pb.RpcObjectSetDetailsResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetDetailsResponseErrorCode, err error) *pb.RpcObjectSetDetailsResponse {
		m := &pb.RpcObjectSetDetailsResponse{Error: &pb.RpcObjectSetDetailsResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetDetails(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectSetDetailsResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetDetailsResponseError_NULL, nil)
}

func (mw *Middleware) ObjectDuplicate(cctx context.Context, req *pb.RpcObjectDuplicateRequest) *pb.RpcObjectDuplicateResponse {
	response := func(templateId string, err error) *pb.RpcObjectDuplicateResponse {
		m := &pb.RpcObjectDuplicateResponse{
			Error: &pb.RpcObjectDuplicateResponseError{Code: pb.RpcObjectDuplicateResponseError_NULL},
			Id:    templateId,
		}
		if err != nil {
			m.Error.Code = pb.RpcObjectDuplicateResponseError_UNKNOWN_ERROR
			m.Error.Description = err.Error()
		}
		return m
	}
	var objectIds []string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		objectIds, err = bs.ObjectsDuplicate([]string{req.ContextId})
		return
	})
	if len(objectIds) == 0 {
		return response("", err)
	}
	return response(objectIds[0], err)
}

func (mw *Middleware) ObjectListDuplicate(cctx context.Context, req *pb.RpcObjectListDuplicateRequest) *pb.RpcObjectListDuplicateResponse {
	response := func(objectIds []string, err error) *pb.RpcObjectListDuplicateResponse {
		m := &pb.RpcObjectListDuplicateResponse{
			Error: &pb.RpcObjectListDuplicateResponseError{Code: pb.RpcObjectListDuplicateResponseError_NULL},
			Ids:   objectIds,
		}
		if err != nil {
			m.Error.Code = pb.RpcObjectListDuplicateResponseError_UNKNOWN_ERROR
			m.Error.Description = err.Error()
		}
		return m
	}
	var objectIds []string
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		objectIds, err = bs.ObjectsDuplicate(req.ObjectIds)
		return
	})
	return response(objectIds, err)
}

func (mw *Middleware) ObjectSearch(cctx context.Context, req *pb.RpcObjectSearchRequest) *pb.RpcObjectSearchResponse {
	response := func(code pb.RpcObjectSearchResponseErrorCode, records []*types.Struct, err error) *pb.RpcObjectSearchResponse {
		m := &pb.RpcObjectSearchResponse{Error: &pb.RpcObjectSearchResponseError{Code: code}, Records: records}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	if mw.app == nil {
		return response(pb.RpcObjectSearchResponseError_BAD_INPUT, nil, fmt.Errorf("account must be started"))
	}

	if req.FullText != "" {
		mw.app.MustComponent(indexer.CName).(indexer.Indexer).ForceFTIndex()
	}

	ds := mw.app.MustComponent(objectstore.CName).(objectstore.ObjectStore)
	records, _, err := ds.Query(nil, database.Query{
		Filters:          req.Filters,
		Sorts:            req.Sorts,
		Offset:           int(req.Offset),
		Limit:            int(req.Limit),
		FullText:         req.FullText,
		ObjectTypeFilter: req.ObjectTypeFilter,
	})
	if err != nil {
		return response(pb.RpcObjectSearchResponseError_UNKNOWN_ERROR, nil, err)
	}

	// Add dates only to the first page of search results
	if req.Offset == 0 {
		records, err = enrichWithDateSuggestion(records, req)
		if err != nil {
			return response(pb.RpcObjectSearchResponseError_UNKNOWN_ERROR, nil, err)
		}
	}

	var records2 = make([]*types.Struct, 0, len(records))
	for _, rec := range records {
		records2 = append(records2, pbtypes.Map(rec.Details, req.Keys...))
	}

	return response(pb.RpcObjectSearchResponseError_NULL, records2, nil)
}

func enrichWithDateSuggestion(records []database.Record, req *pb.RpcObjectSearchRequest) ([]database.Record, error) {
	dt := suggestDateForSearch(time.Now(), req.FullText)
	if dt.IsZero() {
		return records, nil
	}

	id := deriveDateId(dt)

	// Don't duplicate search suggestions
	var found bool
	for _, r := range records {
		if r.Details == nil || r.Details.Fields == nil {
			continue
		}
		if v, ok := r.Details.Fields[bundle.RelationKeyId.String()]; ok {
			if v.GetStringValue() == id {
				found = true
				break
			}
		}

	}
	if found {
		return records, nil
	}

	var rec database.Record
	var workspaceId string
	for _, f := range req.Filters {
		if f.RelationKey == bundle.RelationKeyWorkspaceId.String() && f.Condition == model.BlockContentDataviewFilter_Equal {
			workspaceId = f.Value.GetStringValue()
			break
		}
	}
	rec = makeSuggestedDateRecord(dt, workspaceId)
	f, _ := filter.MakeAndFilter(req.Filters)
	if vg := pbtypes.ValueGetter(rec.Details); f.FilterObject(vg) {
		return append([]database.Record{rec}, records...), nil
	}
	return records, nil
}

func suggestDateForSearch(now time.Time, raw string) time.Time {
	suggesters := []func() time.Time{
		func() time.Time {
			var exprType naturaldate.ExprType
			t, exprType, err := naturaldate.Parse(raw, now)
			if err != nil {
				return time.Time{}
			}
			if exprType == naturaldate.ExprTypeInvalid {
				return time.Time{}
			}

			// naturaldate parses numbers without qualifiers (m,s) as hours in 24 hours clock format. It leads to weird behavior
			// when inputs like "123" represented as "current time + 123 hours"
			if (exprType & naturaldate.ExprTypeClock24Hour) != 0 {
				t = time.Time{}
			}
			return t
		},
		func() time.Time {
			// Don't use plain numbers, because they will be represented as years
			if _, err := strconv.Atoi(strings.TrimSpace(raw)); err == nil {
				return time.Time{}
			}
			// todo: use system locale to get preferred date format
			t, err := dateparse.ParseIn(raw, now.Location(), dateparse.PreferMonthFirst(false))
			if err != nil {
				return time.Time{}
			}
			return t
		},
	}

	var t time.Time
	for _, s := range suggesters {
		if t = s(); !t.IsZero() {
			break
		}
	}
	if t.IsZero() {
		return t
	}

	// Sanitize date

	// Date without year
	if t.Year() == 0 {
		_, month, day := t.Date()
		h, m, s := t.Clock()
		t = time.Date(now.Year(), month, day, h, m, s, 0, t.Location())
	}

	return t
}

func deriveDateId(t time.Time) string {
	return "_date_" + t.Format("2006-01-02")
}

func makeSuggestedDateRecord(t time.Time, workspaceId string) database.Record {
	id := deriveDateId(t)

	d := &types.Struct{Fields: map[string]*types.Value{
		bundle.RelationKeyId.String():          pbtypes.String(id),
		bundle.RelationKeyName.String():        pbtypes.String(t.Format("Mon Jan  2 2006")),
		bundle.RelationKeyType.String():        pbtypes.String(bundle.TypeKeyDate.URL()),
		bundle.RelationKeyIconEmoji.String():   pbtypes.String("📅"),
		bundle.RelationKeyWorkspaceId.String(): pbtypes.String(workspaceId),
	}}

	return database.Record{
		Details: d,
	}
}

func (mw *Middleware) ObjectSearchSubscribe(cctx context.Context, req *pb.RpcObjectSearchSubscribeRequest) *pb.RpcObjectSearchSubscribeResponse {
	errResponse := func(err error) *pb.RpcObjectSearchSubscribeResponse {
		r := &pb.RpcObjectSearchSubscribeResponse{
			Error: &pb.RpcObjectSearchSubscribeResponseError{
				Code: pb.RpcObjectSearchSubscribeResponseError_UNKNOWN_ERROR,
			},
		}
		if err != nil {
			r.Error.Description = err.Error()
		}
		return r
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	if mw.app == nil {
		return errResponse(fmt.Errorf("account must be started"))
	}

	subService := mw.app.MustComponent(subscription.CName).(subscription.Service)

	resp, err := subService.Search(*req)
	if err != nil {
		return errResponse(err)
	}

	return resp
}

func (mw *Middleware) ObjectGroupsSubscribe(_ context.Context, req *pb.RpcObjectGroupsSubscribeRequest) *pb.RpcObjectGroupsSubscribeResponse {
	errResponse := func(err error) *pb.RpcObjectGroupsSubscribeResponse {
		r := &pb.RpcObjectGroupsSubscribeResponse{
			Error: &pb.RpcObjectGroupsSubscribeResponseError{
				Code: pb.RpcObjectGroupsSubscribeResponseError_UNKNOWN_ERROR,
			},
		}
		if err != nil {
			r.Error.Description = err.Error()
		}
		return r
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	if mw.app == nil {
		return errResponse(errors.New("app must be started"))
	}

	subService := mw.app.MustComponent(subscription.CName).(subscription.Service)

	resp, err := subService.SubscribeGroups(*req)
	if err != nil {
		return errResponse(err)
	}

	return resp
}

func (mw *Middleware) ObjectSubscribeIds(_ context.Context, req *pb.RpcObjectSubscribeIdsRequest) *pb.RpcObjectSubscribeIdsResponse {
	errResponse := func(err error) *pb.RpcObjectSubscribeIdsResponse {
		r := &pb.RpcObjectSubscribeIdsResponse{
			Error: &pb.RpcObjectSubscribeIdsResponseError{
				Code: pb.RpcObjectSubscribeIdsResponseError_UNKNOWN_ERROR,
			},
		}
		if err != nil {
			r.Error.Description = err.Error()
		}
		return r
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	if mw.app == nil {
		return errResponse(fmt.Errorf("account must be started"))
	}

	subService := mw.app.MustComponent(subscription.CName).(subscription.Service)

	resp, err := subService.SubscribeIdsReq(*req)
	if err != nil {
		return errResponse(err)
	}

	return resp
}

func (mw *Middleware) ObjectSearchUnsubscribe(cctx context.Context, req *pb.RpcObjectSearchUnsubscribeRequest) *pb.RpcObjectSearchUnsubscribeResponse {
	response := func(err error) *pb.RpcObjectSearchUnsubscribeResponse {
		r := &pb.RpcObjectSearchUnsubscribeResponse{
			Error: &pb.RpcObjectSearchUnsubscribeResponseError{
				Code: pb.RpcObjectSearchUnsubscribeResponseError_NULL,
			},
		}
		if err != nil {
			r.Error.Code = pb.RpcObjectSearchUnsubscribeResponseError_UNKNOWN_ERROR
			r.Error.Description = err.Error()
		}
		return r
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	if mw.app == nil {
		return response(fmt.Errorf("account must be started"))
	}

	subService := mw.app.MustComponent(subscription.CName).(subscription.Service)

	err := subService.Unsubscribe(req.SubIds...)
	if err != nil {
		return response(err)
	}
	return response(nil)
}

// TODO Move logic to block service
func (mw *Middleware) ObjectGraph(cctx context.Context, req *pb.RpcObjectGraphRequest) *pb.RpcObjectGraphResponse {
	response := func(code pb.RpcObjectGraphResponseErrorCode, nodes []*types.Struct, edges []*pb.RpcObjectGraphEdge, err error) *pb.RpcObjectGraphResponse {
		m := &pb.RpcObjectGraphResponse{Error: &pb.RpcObjectGraphResponseError{Code: code}, Nodes: nodes, Edges: edges}
		if err != nil {
			m.Error.Description = err.Error()
		}

		return m
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	if mw.app == nil {
		return response(pb.RpcObjectGraphResponseError_BAD_INPUT, nil, nil, fmt.Errorf("account must be started"))
	}

	at := mw.app.MustComponent(core.CName).(core.Service)
	rs := mw.app.MustComponent(relation.CName).(relation.Service)

	store := app.MustComponent[objectstore.ObjectStore](mw.app)
	records, _, err := store.Query(nil, database.Query{
		Filters:          req.Filters,
		Limit:            int(req.Limit),
		ObjectTypeFilter: req.ObjectTypeFilter,
	})

	if err != nil {
		return response(pb.RpcObjectGraphResponseError_UNKNOWN_ERROR, nil, nil, err)
	}

	relations, err := rs.ListAll(relation.WithWorkspaceId(at.PredefinedBlocks().Account))
	if err != nil {
		return response(pb.RpcObjectGraphResponseError_UNKNOWN_ERROR, nil, nil, err)
	}

	var nodes = make([]*types.Struct, 0, len(records))
	var edges = make([]*pb.RpcObjectGraphEdge, 0, len(records)*2)
	var nodeExists = make(map[string]struct{}, len(records))

	for _, rec := range records {
		id := pbtypes.GetString(rec.Details, bundle.RelationKeyId.String())
		nodeExists[id] = struct{}{}
	}

	homeId := at.PredefinedBlocks().Home
	if _, exists := nodeExists[homeId]; !exists {
		// we don't index home object, but we DO index outgoing links from it
		links, qErr := store.GetOutboundLinksById(homeId)
		if qErr != nil {
			log.Info("failed to query object links, err: ", err)
		}
		records = append(records, database.Record{&types.Struct{
			Fields: map[string]*types.Value{
				"id":        pbtypes.String(homeId),
				"name":      pbtypes.String("Home"),
				"iconEmoji": pbtypes.String("🏠"),
				"links":     pbtypes.StringList(links),
			},
		}})
	}

	for _, rec := range records {
		id := pbtypes.GetString(rec.Details, bundle.RelationKeyId.String())

		nodes = append(nodes, pbtypes.Map(rec.Details, req.Keys...))

		var outgoingRelationLink = make(map[string]struct{}, 10)
		for k, v := range rec.Details.GetFields() {
			if list := pbtypes.GetStringListValue(v); len(list) == 0 {
				continue
			} else {
				rel := relations.GetByKey(k)
				if rel == nil {
					continue
				}

				if rel.Format != model.RelationFormat_object && rel.Format != model.RelationFormat_file {
					continue
				}

				for _, l := range list {
					if _, exists := nodeExists[l]; !exists {
						continue
					}

					if rel.Hidden ||
						rel.Key == bundle.RelationKeyId.String() ||
						rel.Key == bundle.RelationKeyCreator.String() ||
						rel.Key == bundle.RelationKeyWorkspaceId.String() ||
						rel.Key == bundle.RelationKeyLastModifiedBy.String() {
						continue
					}

					edges = append(edges, &pb.RpcObjectGraphEdge{
						Source:      id,
						Target:      l,
						Name:        rel.Name,
						Type:        pb.RpcObjectGraphEdge_Relation,
						Description: rel.Description,
						Hidden:      rel.Hidden,
					})
					outgoingRelationLink[l] = struct{}{}
				}
			}
		}

		sbtProvider := app.MustComponent[typeprovider.SmartBlockTypeProvider](mw.app)
		links := pbtypes.GetStringList(rec.Details, bundle.RelationKeyLinks.String())
		for _, link := range links {
			sbType, _ := sbtProvider.Type(link)
			// ignore files because we index all file blocks as outgoing links
			if sbType == smartblock.SmartBlockTypeFile {
				continue
			}
			if _, exists := outgoingRelationLink[link]; !exists {
				if _, exists := nodeExists[link]; !exists {
					continue
				}
				edges = append(edges, &pb.RpcObjectGraphEdge{
					Source: id,
					Target: link,
					Name:   "",
					Type:   pb.RpcObjectGraphEdge_Link,
				})
			}
		}
	}

	return response(pb.RpcObjectGraphResponseError_NULL, nodes, edges, nil)
}

func (mw *Middleware) ObjectRelationAdd(cctx context.Context, req *pb.RpcObjectRelationAddRequest) *pb.RpcObjectRelationAddResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectRelationAddResponseErrorCode, err error) *pb.RpcObjectRelationAddResponse {
		m := &pb.RpcObjectRelationAddResponse{Error: &pb.RpcObjectRelationAddResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	if len(req.RelationKeys) == 0 {
		return response(pb.RpcObjectRelationAddResponseError_BAD_INPUT, fmt.Errorf("relation is nil"))
	}

	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.AddExtraRelations(ctx, req.ContextId, req.RelationKeys)
	})
	if err != nil {
		return response(pb.RpcObjectRelationAddResponseError_BAD_INPUT, err)
	}

	return response(pb.RpcObjectRelationAddResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationDelete(cctx context.Context, req *pb.RpcObjectRelationDeleteRequest) *pb.RpcObjectRelationDeleteResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectRelationDeleteResponseErrorCode, err error) *pb.RpcObjectRelationDeleteResponse {
		m := &pb.RpcObjectRelationDeleteResponse{Error: &pb.RpcObjectRelationDeleteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.RemoveExtraRelations(ctx, req.ContextId, req.RelationKeys)
	})
	if err != nil {
		return response(pb.RpcObjectRelationDeleteResponseError_BAD_INPUT, err)
	}
	return response(pb.RpcObjectRelationDeleteResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationListAvailable(cctx context.Context, req *pb.RpcObjectRelationListAvailableRequest) *pb.RpcObjectRelationListAvailableResponse {
	response := func(code pb.RpcObjectRelationListAvailableResponseErrorCode, relations []*model.Relation, err error) *pb.RpcObjectRelationListAvailableResponse {
		m := &pb.RpcObjectRelationListAvailableResponse{Relations: relations, Error: &pb.RpcObjectRelationListAvailableResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	var rels []*model.Relation
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		rels, err = bs.ListAvailableRelations(req.ContextId)
		return
	})

	if err != nil {
		return response(pb.RpcObjectRelationListAvailableResponseError_UNKNOWN_ERROR, nil, err)
	}

	return response(pb.RpcObjectRelationListAvailableResponseError_NULL, rels, nil)
}

func (mw *Middleware) ObjectSetLayout(cctx context.Context, req *pb.RpcObjectSetLayoutRequest) *pb.RpcObjectSetLayoutResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetLayoutResponseErrorCode, err error) *pb.RpcObjectSetLayoutResponse {
		m := &pb.RpcObjectSetLayoutResponse{Error: &pb.RpcObjectSetLayoutResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetLayout(ctx, req.ContextId, req.Layout)
	})
	if err != nil {
		return response(pb.RpcObjectSetLayoutResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetLayoutResponseError_NULL, nil)
}

func (mw *Middleware) ObjectSetIsArchived(cctx context.Context, req *pb.RpcObjectSetIsArchivedRequest) *pb.RpcObjectSetIsArchivedResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetIsArchivedResponseErrorCode, err error) *pb.RpcObjectSetIsArchivedResponse {
		m := &pb.RpcObjectSetIsArchivedResponse{Error: &pb.RpcObjectSetIsArchivedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetPageIsArchived(*req)
	})
	if err != nil {
		return response(pb.RpcObjectSetIsArchivedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetIsArchivedResponseError_NULL, nil)
}

func (mw *Middleware) ObjectSetSource(cctx context.Context,
	req *pb.RpcObjectSetSourceRequest) *pb.RpcObjectSetSourceResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetSourceResponseErrorCode, err error) *pb.RpcObjectSetSourceResponse {
		m := &pb.RpcObjectSetSourceResponse{Error: &pb.RpcObjectSetSourceResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetSource(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectSetSourceResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetSourceResponseError_NULL, nil)
}

func (mw *Middleware) ObjectWorkspaceSetDashboard(cctx context.Context, req *pb.RpcObjectWorkspaceSetDashboardRequest) *pb.RpcObjectWorkspaceSetDashboardResponse {
	ctx := mw.newContext(cctx)
	response := func(setId string, err error) *pb.RpcObjectWorkspaceSetDashboardResponse {
		resp := &pb.RpcObjectWorkspaceSetDashboardResponse{
			ObjectId: setId,
			Error: &pb.RpcObjectWorkspaceSetDashboardResponseError{
				Code: pb.RpcObjectWorkspaceSetDashboardResponseError_NULL,
			},
		}
		if err != nil {
			resp.Error.Code = pb.RpcObjectWorkspaceSetDashboardResponseError_UNKNOWN_ERROR
			resp.Error.Description = err.Error()
		} else {
			resp.Event = ctx.GetResponseEvent()
		}
		return resp
	}
	var (
		setId string
		err   error
	)
	err = mw.doBlockService(func(bs *block.Service) error {
		if setId, err = bs.SetWorkspaceDashboardId(ctx, req.ContextId, req.ObjectId); err != nil {
			return err
		}
		return nil
	})
	return response(setId, err)
}

func (mw *Middleware) ObjectSetIsFavorite(cctx context.Context, req *pb.RpcObjectSetIsFavoriteRequest) *pb.RpcObjectSetIsFavoriteResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetIsFavoriteResponseErrorCode, err error) *pb.RpcObjectSetIsFavoriteResponse {
		m := &pb.RpcObjectSetIsFavoriteResponse{Error: &pb.RpcObjectSetIsFavoriteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.SetPageIsFavorite(*req)
	})
	if err != nil {
		return response(pb.RpcObjectSetIsFavoriteResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetIsFavoriteResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationAddFeatured(cctx context.Context, req *pb.RpcObjectRelationAddFeaturedRequest) *pb.RpcObjectRelationAddFeaturedResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectRelationAddFeaturedResponseErrorCode, err error) *pb.RpcObjectRelationAddFeaturedResponse {
		m := &pb.RpcObjectRelationAddFeaturedResponse{Error: &pb.RpcObjectRelationAddFeaturedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.FeaturedRelationAdd(ctx, req.ContextId, req.Relations...)
	})
	if err != nil {
		return response(pb.RpcObjectRelationAddFeaturedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectRelationAddFeaturedResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationRemoveFeatured(cctx context.Context, req *pb.RpcObjectRelationRemoveFeaturedRequest) *pb.RpcObjectRelationRemoveFeaturedResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectRelationRemoveFeaturedResponseErrorCode, err error) *pb.RpcObjectRelationRemoveFeaturedResponse {
		m := &pb.RpcObjectRelationRemoveFeaturedResponse{Error: &pb.RpcObjectRelationRemoveFeaturedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.FeaturedRelationRemove(ctx, req.ContextId, req.Relations...)
	})
	if err != nil {
		return response(pb.RpcObjectRelationRemoveFeaturedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectRelationRemoveFeaturedResponseError_NULL, nil)
}

func (mw *Middleware) ObjectToSet(cctx context.Context, req *pb.RpcObjectToSetRequest) *pb.RpcObjectToSetResponse {
	response := func(setId string, err error) *pb.RpcObjectToSetResponse {
		resp := &pb.RpcObjectToSetResponse{
			SetId: setId,
			Error: &pb.RpcObjectToSetResponseError{
				Code: pb.RpcObjectToSetResponseError_NULL,
			},
		}
		if err != nil {
			resp.Error.Code = pb.RpcObjectToSetResponseError_UNKNOWN_ERROR
			resp.Error.Description = err.Error()
		}
		return resp
	}
	var (
		setId string
		err   error
	)
	err = mw.doBlockService(func(bs *block.Service) error {
		if setId, err = bs.ObjectToSet(req.ContextId, req.Source); err != nil {
			return err
		}
		return nil
	})
	return response(setId, err)
}

func (mw *Middleware) ObjectCreateBookmark(cctx context.Context, req *pb.RpcObjectCreateBookmarkRequest) *pb.RpcObjectCreateBookmarkResponse {
	response := func(code pb.RpcObjectCreateBookmarkResponseErrorCode, id string, details *types.Struct, err error) *pb.RpcObjectCreateBookmarkResponse {
		m := &pb.RpcObjectCreateBookmarkResponse{Error: &pb.RpcObjectCreateBookmarkResponseError{Code: code}, ObjectId: id, Details: details}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	var (
		id         string
		newDetails *types.Struct
	)
	err := mw.doBlockService(func(bs *block.Service) error {
		var err error
		id, newDetails, err = bs.CreateObject(req, bundle.TypeKeyBookmark)
		return err
	})
	if err != nil {
		return response(pb.RpcObjectCreateBookmarkResponseError_UNKNOWN_ERROR, "", newDetails, err)
	}
	return response(pb.RpcObjectCreateBookmarkResponseError_NULL, id, newDetails, nil)
}

func (mw *Middleware) ObjectBookmarkFetch(cctx context.Context, req *pb.RpcObjectBookmarkFetchRequest) *pb.RpcObjectBookmarkFetchResponse {
	response := func(code pb.RpcObjectBookmarkFetchResponseErrorCode, err error) *pb.RpcObjectBookmarkFetchResponse {
		m := &pb.RpcObjectBookmarkFetchResponse{Error: &pb.RpcObjectBookmarkFetchResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	err := mw.doBlockService(func(bs *block.Service) error {
		return bs.ObjectBookmarkFetch(*req)
	})

	if err != nil {
		return response(pb.RpcObjectBookmarkFetchResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectBookmarkFetchResponseError_NULL, nil)
}

func (mw *Middleware) ObjectToBookmark(cctx context.Context, req *pb.RpcObjectToBookmarkRequest) *pb.RpcObjectToBookmarkResponse {
	response := func(code pb.RpcObjectToBookmarkResponseErrorCode, id string, err error) *pb.RpcObjectToBookmarkResponse {
		m := &pb.RpcObjectToBookmarkResponse{Error: &pb.RpcObjectToBookmarkResponseError{Code: code}, ObjectId: id}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	var id string
	err := mw.doBlockService(func(bs *block.Service) error {
		var err error
		id, err = bs.ObjectToBookmark(req.ContextId, req.Url)
		return err
	})

	if err != nil {
		return response(pb.RpcObjectToBookmarkResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcObjectToBookmarkResponseError_NULL, id, nil)
}

func (mw *Middleware) ObjectSetInternalFlags(cctx context.Context, req *pb.RpcObjectSetInternalFlagsRequest) *pb.RpcObjectSetInternalFlagsResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectSetInternalFlagsResponseErrorCode, err error) *pb.RpcObjectSetInternalFlagsResponse {
		m := &pb.RpcObjectSetInternalFlagsResponse{Error: &pb.RpcObjectSetInternalFlagsResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs *block.Service) (err error) {
		return bs.ModifyDetails(req.ContextId, func(current *types.Struct) (*types.Struct, error) {
			d := pbtypes.CopyStruct(current)
			return internalflag.PutToDetails(d, req.InternalFlags), nil
		})
	})
	if err != nil {
		return response(pb.RpcObjectSetInternalFlagsResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetInternalFlagsResponseError_NULL, nil)
}

func (mw *Middleware) ObjectImport(cctx context.Context, req *pb.RpcObjectImportRequest) *pb.RpcObjectImportResponse {
	ctx := mw.newContext(cctx)

	response := func(code pb.RpcObjectImportResponseErrorCode, err error) *pb.RpcObjectImportResponse {
		m := &pb.RpcObjectImportResponse{Error: &pb.RpcObjectImportResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	importer := mw.app.MustComponent(importer.CName).(importer.Importer)
	err := importer.Import(ctx, req)

	if err != nil {
		return response(pb.RpcObjectImportResponseError_INTERNAL_ERROR, err)
	}
	return response(pb.RpcObjectImportResponseError_NULL, nil)
}

func (mw *Middleware) ObjectImportList(cctx context.Context, req *pb.RpcObjectImportListRequest) *pb.RpcObjectImportListResponse {
	ctx := mw.newContext(cctx)

	response := func(res []*pb.RpcObjectImportListImportResponse, code pb.RpcObjectImportListResponseErrorCode, err error) *pb.RpcObjectImportListResponse {
		m := &pb.RpcObjectImportListResponse{Response: res, Error: &pb.RpcObjectImportListResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	importer := mw.app.MustComponent(importer.CName).(importer.Importer)
	res, err := importer.ListImports(ctx, req)

	if err != nil {
		return response(res, pb.RpcObjectImportListResponseError_INTERNAL_ERROR, err)
	}
	return response(res, pb.RpcObjectImportListResponseError_NULL, nil)
}

func (mw *Middleware) ObjectImportNotionValidateToken(ctx context.Context,
	request *pb.RpcObjectImportNotionValidateTokenRequest) *pb.RpcObjectImportNotionValidateTokenResponse {
	// nolint: lll
	response := func(code pb.RpcObjectImportNotionValidateTokenResponseErrorCode) *pb.RpcObjectImportNotionValidateTokenResponse {
		err := &pb.RpcObjectImportNotionValidateTokenResponseError{Code: code}
		switch code {
		case pb.RpcObjectImportNotionValidateTokenResponseError_UNAUTHORIZED:
			err.Description = "Sorry, token not found. Please check Notion integrations."
		case pb.RpcObjectImportNotionValidateTokenResponseError_FORBIDDEN:
			err.Description = "Can't access user information, please fill user capabilities."
		case pb.RpcObjectImportNotionValidateTokenResponseError_SERVICE_UNAVAILABLE:
			err.Description = "Notion is currently unavailable."
		case pb.RpcObjectImportNotionValidateTokenResponseError_NULL:
			err.Description = ""
		default:
			err.Description = "Internal error"
		}
		return &pb.RpcObjectImportNotionValidateTokenResponse{Error: err}
	}

	mw.m.RLock()
	defer mw.m.RUnlock()

	importer := mw.app.MustComponent(importer.CName).(importer.Importer)
	errCode := importer.ValidateNotionToken(ctx, request)
	return response(errCode)
}
