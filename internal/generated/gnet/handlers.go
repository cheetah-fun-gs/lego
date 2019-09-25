package gnet

import (
	"context"

	handlerscommon "github.com/cheetah-fun-gs/goso/internal/biz/handlers/common"
	"github.com/cheetah-fun-gs/goso/pkg/handler"
	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
)

// Handlers gnet Handlers
var Handlers = []*handler.Handler{
	&handler.Handler{
		Name:    "CommonPing",
		Routers: sohttp.NewRouters(handlerscommon.PingURIS, handlerscommon.PingHTTPMethods),
		Req:     &handlerscommon.PingReq{},
		Resp:    &handlerscommon.PingResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlerscommon.PingHandle(ctx, req.(*handlerscommon.PingReq), resp.(*handlerscommon.PingResp))
		},
	},
	&handler.Handler{
		Name:    "CommonPost",
		Routers: sohttp.NewRouters(handlerscommon.PostURIS, handlerscommon.PostHTTPMethods),
		Req:     &handlerscommon.PostReq{},
		Resp:    &handlerscommon.PostResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlerscommon.PostHandle(ctx, req.(*handlerscommon.PostReq), resp.(*handlerscommon.PostResp))
		},
	},
}
