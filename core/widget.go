package core

import (
	"context"

	"github.com/anytypeio/go-anytype-middleware/core/block"
	"github.com/anytypeio/go-anytype-middleware/pb"
)

func (mw *Middleware) BlockCreateWidget(cctx context.Context, req *pb.RpcBlockCreateWidgetRequest) *pb.RpcBlockCreateWidgetResponse {
	ctx := mw.newContext(cctx)
	response := func(code pb.RpcBlockCreateWidgetResponseErrorCode, id string, err error) *pb.RpcBlockCreateWidgetResponse {
		m := &pb.RpcBlockCreateWidgetResponse{Error: &pb.RpcBlockCreateWidgetResponseError{Code: code}, BlockId: id}
		if err != nil {
			m.Error.Description = err.Error()
		} else {
			m.Event = ctx.GetResponseEvent()
		}
		return m
	}
	var id string
	err := mw.doBlockService(func(bs block.Service) (err error) {
		// TODO Check that context object has smartblock type Widget
		// TODO Create BlockWidget wrapper for requested block and insert that wrapper to specific position
		// TODO Refactor basic.Basic a little bit
		// id, err = bs.CreateBlock(ctx, *req)
		return
	})
	if err != nil {
		return response(pb.RpcBlockCreateWidgetResponseError_UNKNOWN_ERROR, "", err)
	}
	return response(pb.RpcBlockCreateWidgetResponseError_NULL, id, nil)
}
