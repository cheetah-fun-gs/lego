package generated

import (
	"context"

	handlercommon "github.com/cheetah-fun-gs/goso/internal/biz/handler/common"
	"github.com/cheetah-fun-gs/goso/pkg/handler"
	"github.com/cheetah-fun-gs/goso/pkg/so"
)

// CommonPingHandler CommonPingHandler
type CommonPingHandler struct {
	*handler.Handler
}

// CloneReq CloneReq
func (h *CommonPingHandler) CloneReq() interface{} {
	return &handlercommon.PingReq{}
}

// CloneResp CloneResp
func (h *CommonPingHandler) CloneResp() interface{} {
	return &handlercommon.PingResp{}
}

// CommonPostHandler CommonPostHandler
type CommonPostHandler struct {
	*handler.Handler
}

// CloneReq CloneReq
func (h *CommonPostHandler) CloneReq() interface{} {
	return &handlercommon.PostReq{}
}

// CloneResp CloneResp
func (h *CommonPostHandler) CloneResp() interface{} {
	return &handlercommon.PostResp{}
}

// Handlers gnet handler
var Handlers = []so.Handler{
	&CommonPingHandler{
		Handler: &handler.Handler{
			Name:    "CommonPing",
			Nets:    handlercommon.PingNetTypes,
			Routers: handlercommon.PingRouters,
			Func: func(ctx context.Context, req, resp interface{}) error {
				return handlercommon.PingHandle(ctx, req.(*handlercommon.PingReq), resp.(*handlercommon.PingResp))
			},
		},
	},
	&CommonPostHandler{
		Handler: &handler.Handler{
			Name:    "CommonPost",
			Nets:    handlercommon.PostNetTypes,
			Routers: handlercommon.PostRouters,
			Func: func(ctx context.Context, req, resp interface{}) error {
				return handlercommon.PostHandle(ctx, req.(*handlercommon.PostReq), resp.(*handlercommon.PostResp))
			},
		},
	},
}
