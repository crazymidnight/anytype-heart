package core

import (
	"fmt"
	bookmark2 "github.com/anytypeio/go-anytype-middleware/core/block/editor/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/block/simple/bookmark"
	"github.com/anytypeio/go-anytype-middleware/core/indexer"
	"github.com/anytypeio/go-anytype-middleware/util/linkpreview"
	"strings"
	"time"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/editor/state"
	"github.com/anytypeio/go-anytype-middleware/core/subscription"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database/filter"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
	"github.com/araddon/dateparse"
	"github.com/gogo/protobuf/types"
	"github.com/tj/go-naturaldate"
)

// To be renamed to ObjectSetDetails
func (mw *Middleware) ObjectSetDetails(req *pb.RpcObjectSetDetailsRequest) *pb.RpcObjectSetDetailsResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectSetDetailsResponseErrorCode, err error) *pb.RpcObjectSetDetailsResponse {
		m := &pb.RpcObjectSetDetailsResponse{Error: &pb.RpcObjectSetDetailsResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetDetails(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectSetDetailsResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetDetailsResponseError_NULL, nil)
}

func (mw *Middleware) ObjectDuplicate(req *pb.RpcObjectDuplicateRequest) *pb.RpcObjectDuplicateResponse {
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
	err := mw.doBlockService(func(bs block.Service) (err error) {
		objectIds, err = bs.ObjectsDuplicate([]string{req.ContextId})
		return
	})
	if len(objectIds) == 0 {
		return response("", err)
	}
	return response(objectIds[0], err)
}

func (mw *Middleware) ObjectListDuplicate(req *pb.RpcObjectListDuplicateRequest) *pb.RpcObjectListDuplicateResponse {
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
	err := mw.doBlockService(func(bs block.Service) (err error) {
		objectIds, err = bs.ObjectsDuplicate(req.ObjectIds)
		return
	})
	return response(objectIds, err)
}

func handleDateSearch(req *pb.RpcObjectSearchRequest, records []database.Record) []database.Record {
	n := time.Now()
	f, _ := filter.MakeAndFilter(req.Filters)
	t, err := naturaldate.Parse(req.FullText, n)
	if err == nil {
		if t.Equal(n) && !strings.EqualFold(req.FullText, "now") {
			// naturaldate pkg returns NOW by default, but we don't need it
			t = time.Time{}
		}
	} else {
		// todo: use system locale to get preferred date format
		t, err = dateparse.ParseAny(req.FullText, dateparse.PreferMonthFirst(false))
	}

	if !t.IsZero() {
		d := &types.Struct{Fields: map[string]*types.Value{
			"id":        pbtypes.String("_date_" + t.Format("2006-01-02")),
			"name":      pbtypes.String(t.Format("Mon Jan  2 2006")),
			"type":      pbtypes.String(bundle.TypeKeyDate.URL()),
			"iconEmoji": pbtypes.String("📅"),
		}}
		if vg := pbtypes.ValueGetter(d); f.FilterObject(vg) {
			records = append([]database.Record{{Details: d}}, records...)
		}
	}

	return records
}

func (mw *Middleware) ObjectSearch(req *pb.RpcObjectSearchRequest) *pb.RpcObjectSearchResponse {
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

	at := mw.app.MustComponent(core.CName).(core.Service)
	if req.FullText != "" {
		mw.app.MustComponent(indexer.CName).(indexer.Indexer).ForceFTIndex()
	}
	records, _, err := at.ObjectStore().Query(nil, database.Query{
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

	records = handleDateSearch(req, records)
	var records2 = make([]*types.Struct, 0, len(records))
	for _, rec := range records {
		records2 = append(records2, pbtypes.Map(rec.Details, req.Keys...))
	}

	return response(pb.RpcObjectSearchResponseError_NULL, records2, nil)
}

func (mw *Middleware) ObjectSearchSubscribe(req *pb.RpcObjectSearchSubscribeRequest) *pb.RpcObjectSearchSubscribeResponse {
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

func (mw *Middleware) ObjectSubscribeIds(req *pb.RpcObjectSubscribeIdsRequest) *pb.RpcObjectSubscribeIdsResponse {
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

func (mw *Middleware) ObjectSearchUnsubscribe(req *pb.RpcObjectSearchUnsubscribeRequest) *pb.RpcObjectSearchUnsubscribeResponse {
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

func (mw *Middleware) ObjectGraph(req *pb.RpcObjectGraphRequest) *pb.RpcObjectGraphResponse {
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

	records, _, err := at.ObjectStore().Query(nil, database.Query{
		Filters:          req.Filters,
		Limit:            int(req.Limit),
		ObjectTypeFilter: req.ObjectTypeFilter,
	})
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
		records = append(records, database.Record{&types.Struct{
			Fields: map[string]*types.Value{
				"id":        pbtypes.String(homeId),
				"name":      pbtypes.String("Home"),
				"iconEmoji": pbtypes.String("🏠"),
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

				rel, err := at.ObjectStore().GetRelation(k)
				if err != nil {
					log.Errorf("ObjectGraph failed to get relation %s: %s", k, err.Error())
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
		links := pbtypes.GetStringList(rec.Details, bundle.RelationKeyLinks.String())
		for _, link := range links {
			sbType, _ := smartblock.SmartBlockTypeFromID(link)
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

func (mw *Middleware) ObjectRelationAdd(req *pb.RpcObjectRelationAddRequest) *pb.RpcObjectRelationAddResponse {
	ctx := state.NewContext(nil)
	response := func(relation *model.Relation, code pb.RpcObjectRelationAddResponseErrorCode, err error) *pb.RpcObjectRelationAddResponse {
		var relKey string
		if relation != nil {
			relKey = relation.Key
		}
		m := &pb.RpcObjectRelationAddResponse{RelationKey: relKey, Relation: relation, Error: &pb.RpcObjectRelationAddResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	if req.Relation == nil {
		return response(nil, pb.RpcObjectRelationAddResponseError_BAD_INPUT, fmt.Errorf("relation is nil"))
	}

	var relations []*model.Relation
	err := mw.doBlockService(func(bs block.Service) (err error) {
		relations, err = bs.AddExtraRelations(ctx, req.ContextId, []*model.Relation{req.Relation})
		return err
	})
	if err != nil {
		return response(nil, pb.RpcObjectRelationAddResponseError_BAD_INPUT, err)
	}

	if len(relations) == 0 {
		return response(nil, pb.RpcObjectRelationAddResponseError_BAD_INPUT, nil)
	}

	return response(relations[0], pb.RpcObjectRelationAddResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationUpdate(req *pb.RpcObjectRelationUpdateRequest) *pb.RpcObjectRelationUpdateResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectRelationUpdateResponseErrorCode, err error) *pb.RpcObjectRelationUpdateResponse {
		m := &pb.RpcObjectRelationUpdateResponse{Error: &pb.RpcObjectRelationUpdateResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.UpdateExtraRelations(nil, req.ContextId, []*model.Relation{req.Relation}, false)
	})
	if err != nil {
		return response(pb.RpcObjectRelationUpdateResponseError_BAD_INPUT, err)
	}

	return response(pb.RpcObjectRelationUpdateResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationDelete(req *pb.RpcObjectRelationDeleteRequest) *pb.RpcObjectRelationDeleteResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectRelationDeleteResponseErrorCode, err error) *pb.RpcObjectRelationDeleteResponse {
		m := &pb.RpcObjectRelationDeleteResponse{Error: &pb.RpcObjectRelationDeleteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.RemoveExtraRelations(ctx, req.ContextId, []string{req.RelationKey})
	})
	if err != nil {
		return response(pb.RpcObjectRelationDeleteResponseError_BAD_INPUT, err)
	}
	return response(pb.RpcObjectRelationDeleteResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationOptionAdd(req *pb.RpcObjectRelationOptionAddRequest) *pb.RpcObjectRelationOptionAddResponse {
	ctx := state.NewContext(nil)
	response := func(opt *model.RelationOption, code pb.RpcObjectRelationOptionAddResponseErrorCode, err error) *pb.RpcObjectRelationOptionAddResponse {
		m := &pb.RpcObjectRelationOptionAddResponse{Option: opt, Error: &pb.RpcObjectRelationOptionAddResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var opt *model.RelationOption
	err := mw.doBlockService(func(bs block.Service) (err error) {
		var err2 error
		opt, err2 = bs.AddExtraRelationOption(ctx, *req)
		return err2
	})
	if err != nil {
		return response(nil, pb.RpcObjectRelationOptionAddResponseError_BAD_INPUT, err)
	}

	return response(opt, pb.RpcObjectRelationOptionAddResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationOptionUpdate(req *pb.RpcObjectRelationOptionUpdateRequest) *pb.RpcObjectRelationOptionUpdateResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectRelationOptionUpdateResponseErrorCode, err error) *pb.RpcObjectRelationOptionUpdateResponse {
		m := &pb.RpcObjectRelationOptionUpdateResponse{Error: &pb.RpcObjectRelationOptionUpdateResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.UpdateExtraRelationOption(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectRelationOptionUpdateResponseError_BAD_INPUT, err)
	}

	return response(pb.RpcObjectRelationOptionUpdateResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationOptionDelete(req *pb.RpcObjectRelationOptionDeleteRequest) *pb.RpcObjectRelationOptionDeleteResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectRelationOptionDeleteResponseErrorCode, err error) *pb.RpcObjectRelationOptionDeleteResponse {
		m := &pb.RpcObjectRelationOptionDeleteResponse{Error: &pb.RpcObjectRelationOptionDeleteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.DeleteExtraRelationOption(ctx, *req)
	})
	if err != nil {
		return response(pb.RpcObjectRelationOptionDeleteResponseError_BAD_INPUT, err)
	}

	return response(pb.RpcObjectRelationOptionDeleteResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationListAvailable(req *pb.RpcObjectRelationListAvailableRequest) *pb.RpcObjectRelationListAvailableResponse {
	response := func(code pb.RpcObjectRelationListAvailableResponseErrorCode, relations []*model.Relation, err error) *pb.RpcObjectRelationListAvailableResponse {
		m := &pb.RpcObjectRelationListAvailableResponse{Relations: relations, Error: &pb.RpcObjectRelationListAvailableResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}
	var rels []*model.Relation
	err := mw.doBlockService(func(bs block.Service) (err error) {
		rels, err = bs.ListAvailableRelations(req.ContextId)
		return
	})

	if err != nil {
		return response(pb.RpcObjectRelationListAvailableResponseError_UNKNOWN_ERROR, nil, err)
	}

	return response(pb.RpcObjectRelationListAvailableResponseError_NULL, rels, nil)
}

func (mw *Middleware) ObjectSetLayout(req *pb.RpcObjectSetLayoutRequest) *pb.RpcObjectSetLayoutResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectSetLayoutResponseErrorCode, err error) *pb.RpcObjectSetLayoutResponse {
		m := &pb.RpcObjectSetLayoutResponse{Error: &pb.RpcObjectSetLayoutResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetLayout(ctx, req.ContextId, req.Layout)
	})
	if err != nil {
		return response(pb.RpcObjectSetLayoutResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetLayoutResponseError_NULL, nil)
}

func (mw *Middleware) ObjectSetIsArchived(req *pb.RpcObjectSetIsArchivedRequest) *pb.RpcObjectSetIsArchivedResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectSetIsArchivedResponseErrorCode, err error) *pb.RpcObjectSetIsArchivedResponse {
		m := &pb.RpcObjectSetIsArchivedResponse{Error: &pb.RpcObjectSetIsArchivedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetPageIsArchived(*req)
	})
	if err != nil {
		return response(pb.RpcObjectSetIsArchivedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetIsArchivedResponseError_NULL, nil)
}

func (mw *Middleware) ObjectSetIsFavorite(req *pb.RpcObjectSetIsFavoriteRequest) *pb.RpcObjectSetIsFavoriteResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectSetIsFavoriteResponseErrorCode, err error) *pb.RpcObjectSetIsFavoriteResponse {
		m := &pb.RpcObjectSetIsFavoriteResponse{Error: &pb.RpcObjectSetIsFavoriteResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.SetPageIsFavorite(*req)
	})
	if err != nil {
		return response(pb.RpcObjectSetIsFavoriteResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectSetIsFavoriteResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationAddFeatured(req *pb.RpcObjectRelationAddFeaturedRequest) *pb.RpcObjectRelationAddFeaturedResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectRelationAddFeaturedResponseErrorCode, err error) *pb.RpcObjectRelationAddFeaturedResponse {
		m := &pb.RpcObjectRelationAddFeaturedResponse{Error: &pb.RpcObjectRelationAddFeaturedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.FeaturedRelationAdd(ctx, req.ContextId, req.Relations...)
	})
	if err != nil {
		return response(pb.RpcObjectRelationAddFeaturedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectRelationAddFeaturedResponseError_NULL, nil)
}

func (mw *Middleware) ObjectRelationRemoveFeatured(req *pb.RpcObjectRelationRemoveFeaturedRequest) *pb.RpcObjectRelationRemoveFeaturedResponse {
	ctx := state.NewContext(nil)
	response := func(code pb.RpcObjectRelationRemoveFeaturedResponseErrorCode, err error) *pb.RpcObjectRelationRemoveFeaturedResponse {
		m := &pb.RpcObjectRelationRemoveFeaturedResponse{Error: &pb.RpcObjectRelationRemoveFeaturedResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.FeaturedRelationRemove(ctx, req.ContextId, req.Relations...)
	})
	if err != nil {
		return response(pb.RpcObjectRelationRemoveFeaturedResponseError_UNKNOWN_ERROR, err)
	}
	return response(pb.RpcObjectRelationRemoveFeaturedResponseError_NULL, nil)
}

func (mw *Middleware) ObjectToSet(req *pb.RpcObjectToSetRequest) *pb.RpcObjectToSetResponse {
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
	err = mw.doBlockService(func(bs block.Service) error {
		if setId, err = bs.ObjectToSet(req.ContextId, req.Source); err != nil {
			return err
		}
		return nil
	})
	return response(setId, err)
}

func (mw *Middleware) ObjectCreateBookmark(req *pb.RpcObjectCreateBookmarkRequest) *pb.RpcObjectCreateBookmarkResponse {
	response := func(code pb.RpcObjectCreateBookmarkResponseErrorCode, id string, err error) *pb.RpcObjectCreateBookmarkResponse {
		m := &pb.RpcObjectCreateBookmarkResponse{Error: &pb.RpcObjectCreateBookmarkResponseError{Code: code}, PageId: id}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	var id string
	err := mw.doBlockService(func(bs block.Service) error {
		// TODO: temp in-place logic
		content := &bookmark.Content{
			Url: req.Url,
		}
		lp := mw.app.MustComponent(linkpreview.CName).(linkpreview.LinkPreview)
		updaters, err := bookmark2.ContentFetcher(req.Url, lp, mw.GetAnytype())
		if err != nil {
			return err
		}
		for upd := range updaters {
			if err := upd(content); err != nil {
				return err
			}
		}

		id, err = bookmark2.CreateBookmarkObject(mw.GetAnytype().ObjectStore(), bs, (*model.BlockContentBookmark)(content))
		return err
	})

	if err != nil {
		return response(pb.RpcObjectCreateBookmarkResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcObjectCreateBookmarkResponseError_NULL, id, nil)
}
