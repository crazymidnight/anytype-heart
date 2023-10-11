package core

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/types"

	"github.com/anyproto/anytype-heart/core/block"
	"github.com/anyproto/anytype-heart/core/block/object/objectcreator"
	"github.com/anyproto/anytype-heart/pb"
	"github.com/anyproto/anytype-heart/pkg/lib/bundle"
	"github.com/anyproto/anytype-heart/pkg/lib/localstore/addr"
	"github.com/anyproto/anytype-heart/util/pbtypes"
	"github.com/anyproto/anytype-heart/util/slice"
)

func (mw *Middleware) ObjectTypeRelationAdd(cctx context.Context, req *pb.RpcObjectTypeRelationAddRequest) *pb.RpcObjectTypeRelationAddResponse {
	blockService := getService[*block.Service](mw)
	err := blockService.ObjectTypeRelationAdd(cctx, req)
	code := mapErrorCode(err,
		errToCode(block.ErrBundledTypeIsReadonly, pb.RpcObjectTypeRelationAddResponseError_READONLY_OBJECT_TYPE),
	)
	return &pb.RpcObjectTypeRelationAddResponse{
		Error: &pb.RpcObjectTypeRelationAddResponseError{
			Code: code,
		},
	}
}

func (mw *Middleware) ObjectTypeRelationRemove(cctx context.Context, req *pb.RpcObjectTypeRelationRemoveRequest) *pb.RpcObjectTypeRelationRemoveResponse {
	response := func(code pb.RpcObjectTypeRelationRemoveResponseErrorCode, err error) *pb.RpcObjectTypeRelationRemoveResponse {
		m := &pb.RpcObjectTypeRelationRemoveResponse{Error: &pb.RpcObjectTypeRelationRemoveResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	if strings.HasPrefix(req.ObjectTypeUrl, bundle.TypePrefix) {
		return response(pb.RpcObjectTypeRelationRemoveResponseError_READONLY_OBJECT_TYPE, fmt.Errorf("can't modify bundled object type"))
	}

	err := mw.doBlockService(func(bs *block.Service) (err error) {
		err = bs.ModifyDetails(req.ObjectTypeUrl, func(current *types.Struct) (*types.Struct, error) {
			list := pbtypes.GetStringList(current, bundle.RelationKeyRecommendedRelations.String())
			for _, relKey := range req.RelationKeys {
				relId := addr.RelationKeyToIdPrefix + relKey
				if pos := slice.FindPos(list, relId); pos != -1 {
					list = append(list[:pos], list[pos+1:]...)
				}
			}

			detCopy := pbtypes.CopyStruct(current)
			detCopy.Fields[bundle.RelationKeyRecommendedRelations.String()] = pbtypes.StringList(list)
			return detCopy, nil
		})
		if err != nil {
			return err
		}
		return
	})

	if err != nil {
		return response(pb.RpcObjectTypeRelationRemoveResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcObjectTypeRelationRemoveResponseError_NULL, nil)
}

func (mw *Middleware) ObjectCreateObjectType(cctx context.Context, req *pb.RpcObjectCreateObjectTypeRequest) *pb.RpcObjectCreateObjectTypeResponse {
	response := func(code pb.RpcObjectCreateObjectTypeResponseErrorCode, id string, details *types.Struct, err error) *pb.RpcObjectCreateObjectTypeResponse {
		m := &pb.RpcObjectCreateObjectTypeResponse{ObjectId: id, Details: details, Error: &pb.RpcObjectCreateObjectTypeResponseError{Code: code}}
		if err != nil {
			m.Error.Description = err.Error()
		}
		return m
	}

	creator := getService[objectcreator.Service](mw)
	createReq := objectcreator.CreateObjectRequest{
		ObjectTypeKey: bundle.TypeKeyObjectType,
		InternalFlags: req.InternalFlags,
		Details:       req.Details,
	}
	id, newDetails, err := creator.CreateObject(cctx, req.SpaceId, createReq)
	if err != nil {
		return response(pb.RpcObjectCreateObjectTypeResponseError_UNKNOWN_ERROR, "", nil, err)
	}

	return response(pb.RpcObjectCreateObjectTypeResponseError_NULL, id, newDetails, nil)
}

func (mw *Middleware) ObjectCreateSet(cctx context.Context, req *pb.RpcObjectCreateSetRequest) *pb.RpcObjectCreateSetResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcObjectCreateSetResponseErrorCode, id string, newDetails *types.Struct, err error) *pb.RpcObjectCreateSetResponse {
		m := &pb.RpcObjectCreateSetResponse{Error: &pb.RpcObjectCreateSetResponseError{Code: code}, ObjectId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = mw.getResponseEvent(ctx)
			m.Details = newDetails
		}
		return m
	}

	if req.Details == nil {
		req.Details = &types.Struct{}
	}
	if req.Details.Fields == nil {
		req.Details.Fields = map[string]*types.Value{}
	}
	req.Details.Fields[bundle.RelationKeySetOf.String()] = pbtypes.StringList(req.Source)

	creator := getService[objectcreator.Service](mw)
	createReq := objectcreator.CreateObjectRequest{
		ObjectTypeKey: bundle.TypeKeySet,
		InternalFlags: req.InternalFlags,
		Details:       req.Details,
	}
	id, newDetails, err := creator.CreateObject(cctx, req.SpaceId, createReq)
	if err != nil {
		if errors.Is(err, block.ErrUnknownObjectType) {
			return response(pb.RpcObjectCreateSetResponseError_UNKNOWN_OBJECT_TYPE_URL, "", nil, err)
		}
		return response(pb.RpcObjectCreateSetResponseError_UNKNOWN_ERROR, "", nil, err)
	}

	return response(pb.RpcObjectCreateSetResponseError_NULL, id, newDetails, nil)
}

func (mw *Middleware) ObjectCreateRelation(cctx context.Context, req *pb.RpcObjectCreateRelationRequest) *pb.RpcObjectCreateRelationResponse {
	response := func(id string, object *types.Struct, err error) *pb.RpcObjectCreateRelationResponse {
		if err != nil {
			return &pb.RpcObjectCreateRelationResponse{
				Error: &pb.RpcObjectCreateRelationResponseError{
					Code:        pb.RpcObjectCreateRelationResponseError_UNKNOWN_ERROR,
					Description: err.Error(),
				},
			}
		}
		key := pbtypes.GetString(object, bundle.RelationKeyRelationKey.String())
		return &pb.RpcObjectCreateRelationResponse{
			Error: &pb.RpcObjectCreateRelationResponseError{
				Code: pb.RpcObjectCreateRelationResponseError_NULL,
			},
			ObjectId: id,
			Key:      key,
			Details:  object,
		}
	}
	creator := getService[objectcreator.Service](mw)
	createReq := objectcreator.CreateObjectRequest{
		ObjectTypeKey: bundle.TypeKeyRelation,
		Details:       req.Details,
	}
	id, newDetails, err := creator.CreateObject(cctx, req.SpaceId, createReq)
	if err != nil {
		return response("", nil, err)
	}
	return response(id, newDetails, err)
}

func (mw *Middleware) ObjectCreateRelationOption(cctx context.Context, req *pb.RpcObjectCreateRelationOptionRequest) *pb.RpcObjectCreateRelationOptionResponse {
	response := func(id string, newDetails *types.Struct, err error) *pb.RpcObjectCreateRelationOptionResponse {
		if err != nil {
			return &pb.RpcObjectCreateRelationOptionResponse{
				Error: &pb.RpcObjectCreateRelationOptionResponseError{
					Code:        pb.RpcObjectCreateRelationOptionResponseError_UNKNOWN_ERROR,
					Description: err.Error(),
				},
			}
		}
		return &pb.RpcObjectCreateRelationOptionResponse{
			Error: &pb.RpcObjectCreateRelationOptionResponseError{
				Code: pb.RpcObjectCreateRelationOptionResponseError_NULL,
			},
			ObjectId: id,
			Details:  newDetails,
		}
	}

	creator := getService[objectcreator.Service](mw)
	createReq := objectcreator.CreateObjectRequest{
		ObjectTypeKey: bundle.TypeKeyRelationOption,
		Details:       req.Details,
	}
	id, newDetails, err := creator.CreateObject(cctx, req.SpaceId, createReq)
	return response(id, newDetails, err)
}

func (mw *Middleware) RelationListRemoveOption(cctx context.Context, request *pb.RpcRelationListRemoveOptionRequest) *pb.RpcRelationListRemoveOptionResponse {
	response := func(code pb.RpcRelationListRemoveOptionResponseErrorCode, err error) *pb.RpcRelationListRemoveOptionResponse {
		if err != nil {
			return &pb.RpcRelationListRemoveOptionResponse{
				Error: &pb.RpcRelationListRemoveOptionResponseError{
					Code:        code,
					Description: err.Error(),
				},
			}
		}

		return &pb.RpcRelationListRemoveOptionResponse{
			Error: &pb.RpcRelationListRemoveOptionResponseError{
				Code: code,
			},
		}
	}

	err := mw.doBlockService(func(bs *block.Service) error {
		var err error
		err = bs.RemoveListOption(request.OptionIds, request.CheckInObjects)
		return err
	})
	if err != nil {
		if errors.Is(err, block.ErrOptionUsedByOtherObjects) {
			return response(pb.RpcRelationListRemoveOptionResponseError_OPTION_USED_BY_OBJECTS, err)
		}
		return response(pb.RpcRelationListRemoveOptionResponseError_UNKNOWN_ERROR, err)
	}

	return response(pb.RpcRelationListRemoveOptionResponseError_NULL, nil)
}

func (mw *Middleware) RelationOptions(cctx context.Context, request *pb.RpcRelationOptionsRequest) *pb.RpcRelationOptionsResponse {
	// TODO implement me
	panic("implement me")
}
