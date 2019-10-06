package generated

import (
	"context"

	handlercommon "github.com/cheetah-fun-gs/goso/internal/biz/handler/common"
	"github.com/cheetah-fun-gs/goso/pkg/handler"
	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
)

// Handlers gnet handler
var Handlers = []*handler.Handler{
	&handler.Handler{
		Name:    "CommonPing",
		Nets:    handlercommon.PingNetTypes,
		Routers: sohttp.NewRouters(handlercommon.PingURIS, handlercommon.PingHTTPMethods),
		Req:     &handlercommon.PingReq{},
		Resp:    &handlercommon.PingResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlercommon.PingHandle(ctx, req.(*handlercommon.PingReq), resp.(*handlercommon.PingResp))
		},
	},
	&handler.Handler{
		Name:    "CommonPost",
		Nets:    handlercommon.PostNetTypes,
		Routers: sohttp.NewRouters(handlercommon.PostURIS, handlercommon.PostHTTPMethods),
		Req:     &handlercommon.PostReq{},
		Resp:    &handlercommon.PostResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlercommon.PostHandle(ctx, req.(*handlercommon.PostReq), resp.(*handlercommon.PostResp))
		},
	},
}
