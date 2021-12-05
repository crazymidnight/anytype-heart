package core

import (
	"fmt"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/core/block/export"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/bundle"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/core/smartblock"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/database"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/localstore/objectstore"
	"github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"
	"github.com/anytypeio/go-anytype-middleware/util/pbtypes"
)

func (mw *Middleware) MakeTemplate(req *pb.RpcMakeTemplateRequest) *pb.RpcMakeTemplateResponse {
	response := func(templateId string, err error) *pb.RpcMakeTemplateResponse {
		m := &pb.RpcMakeTemplateResponse{
			Error: &pb.RpcMakeTemplateResponseError{Code: pb.RpcMakeTemplateResponseError_NULL},
			Id:    templateId,
		}
		if err != nil {
			m.Error.Code = pb.RpcMakeTemplateResponseError_UNKNOWN_ERROR
			m.Error.Description = err.Error()
		}
		return m
	}
	var templateId string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		templateId, err = bs.MakeTemplate(req.ContextId)
		return
	})
	return response(templateId, err)
}

func (mw *Middleware) CloneTemplate(req *pb.RpcCloneTemplateRequest) *pb.RpcCloneTemplateResponse {
	response := func(templateId string, err error) *pb.RpcCloneTemplateResponse {
		m := &pb.RpcCloneTemplateResponse{
			Error: &pb.RpcCloneTemplateResponseError{Code: pb.RpcCloneTemplateResponseError_NULL},
			Id:    templateId,
		}
		if err != nil {
			m.Error.Code = pb.RpcCloneTemplateResponseError_UNKNOWN_ERROR
			m.Error.Description = err.Error()
		}
		return m
	}
	var templateId string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		templateId, err = bs.CloneTemplate(req.ContextId)
		return
	})
	return response(templateId, err)
}

func (mw *Middleware) ApplyTemplate(req *pb.RpcApplyTemplateRequest) *pb.RpcApplyTemplateResponse {
	response := func(err error) *pb.RpcApplyTemplateResponse {
		m := &pb.RpcApplyTemplateResponse{
			Error: &pb.RpcApplyTemplateResponseError{Code: pb.RpcApplyTemplateResponseError_NULL},
		}
		if err != nil {
			m.Error.Code = pb.RpcApplyTemplateResponseError_UNKNOWN_ERROR
			m.Error.Description = err.Error()
		}
		return m
	}
	err := mw.doBlockService(func(bs block.Service) (err error) {
		return bs.ApplyTemplate(req.ContextId, req.TemplateId)
	})
	return response(err)
}

func (mw *Middleware) MakeTemplateByObjectType(req *pb.RpcMakeTemplateByObjectTypeRequest) *pb.RpcMakeTemplateByObjectTypeResponse {
	response := func(templateId string, err error) *pb.RpcMakeTemplateByObjectTypeResponse {
		m := &pb.RpcMakeTemplateByObjectTypeResponse{
			Error: &pb.RpcMakeTemplateByObjectTypeResponseError{Code: pb.RpcMakeTemplateByObjectTypeResponseError_NULL},
			Id:    templateId,
		}
		if err != nil {
			m.Error.Code = pb.RpcMakeTemplateByObjectTypeResponseError_UNKNOWN_ERROR
			m.Error.Description = err.Error()
		}
		return m
	}
	var templateId string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		templateId, err = bs.MakeTemplateByObjectType(req.ObjectType)
		return
	})
	return response(templateId, err)
}

func (mw *Middleware) ExportTemplates(req *pb.RpcExportTemplatesRequest) *pb.RpcExportTemplatesResponse {
	response := func(path string, err error) (res *pb.RpcExportTemplatesResponse) {
		res = &pb.RpcExportTemplatesResponse{
			Error: &pb.RpcExportTemplatesResponseError{
				Code: pb.RpcExportTemplatesResponseError_NULL,
			},
		}
		if err != nil {
			res.Error.Code = pb.RpcExportTemplatesResponseError_UNKNOWN_ERROR
			res.Error.Description = err.Error()
			return
		} else {
			res.Path = path
		}
		return res
	}
	var (
		path string
		err  error
	)
	err = mw.doBlockService(func(_ block.Service) error {
		es := mw.app.MustComponent(export.CName).(export.Export)
		ds := mw.app.MustComponent(objectstore.CName).(objectstore.ObjectStore)
		res, _, err := ds.QueryObjectInfo(database.Query{
			Filters: []*model.BlockContentDataviewFilter{
				{
					RelationKey: bundle.RelationKeyIsArchived.String(),
					Condition:   model.BlockContentDataviewFilter_Equal,
					Value:       pbtypes.Bool(false),
				},
			},
		}, []smartblock.SmartBlockType{smartblock.SmartBlockTypeTemplate})
		if err != nil {
			return err
		}
		var docIds []string
		for _, r := range res {
			docIds = append(docIds, r.Id)
		}
		if len(docIds) == 0 {
			return fmt.Errorf("no templates")
		}
		path, _, err = es.Export(pb.RpcExportRequest{
			Path:   req.Path,
			DocIds: docIds,
			Format: pb.RpcExport_Protobuf,
			Zip:    true,
		})
		return err
	})
	return response(path, err)
}

func (mw *Middleware) ExportWorkspace(req *pb.RpcExportWorkspaceRequest) *pb.RpcExportWorkspaceResponse {
	response := func(path string, err error) (res *pb.RpcExportWorkspaceResponse) {
		res = &pb.RpcExportWorkspaceResponse{
			Error: &pb.RpcExportWorkspaceResponseError{
				Code: pb.RpcExportWorkspaceResponseError_NULL,
			},
		}
		if err != nil {
			res.Error.Code = pb.RpcExportWorkspaceResponseError_UNKNOWN_ERROR
			res.Error.Description = err.Error()
			return
		} else {
			res.Path = path
		}
		return res
	}
	var (
		path string
		err  error
	)
	err = mw.doBlockService(func(_ block.Service) error {
		es := mw.app.MustComponent(export.CName).(export.Export)
		ds := mw.app.MustComponent(objectstore.CName).(objectstore.ObjectStore)
		res, _, err := ds.QueryObjectInfo(database.Query{
			Filters: []*model.BlockContentDataviewFilter{
				{
					RelationKey: bundle.RelationKeyIsArchived.String(),
					Condition:   model.BlockContentDataviewFilter_Equal,
					Value:       pbtypes.Bool(false),
				},
				{
					RelationKey: bundle.RelationKeyWorkspaceId.String(),
					Condition:   model.BlockContentDataviewFilter_Equal,
					Value:       pbtypes.String(req.WorkspaceId),
				},
			},
		}, []smartblock.SmartBlockType{})
		if err != nil {
			return err
		}
		var docIds []string
		for _, r := range res {
			docIds = append(docIds, r.Id)
		}
		if len(docIds) == 0 {
			return fmt.Errorf("no objects in workspace")
		}
		path, _, err = es.Export(pb.RpcExportRequest{
			Path:          req.Path,
			DocIds:        docIds,
			Format:        pb.RpcExport_Protobuf,
			Zip:           true,
			IncludeNested: false,
		})
		return err
	})
	return response(path, err)
}
