package gnet

import (
	"context"
	"goso/internal/biz/handlers"
	"goso/pkg/handler"
	sohttp "goso/pkg/net/sohttp"
)

// handler 定义
var hs = []*handler.Handler{
	&handler.Handler{
		Name:    "common.time",
		Routers: sohttp.NewRouters(handlers.CommonTimeURIS, handlers.CommonTimeHTTPMethods),
		Req:     &handlers.CommonTimeReq{},
		Resp:    &handlers.CommonTimeResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlers.CommonTimeHandle(ctx, req.(*handlers.CommonTimeReq), resp.(*handlers.CommonTimeResp))
		},
	},
	&handler.Handler{
		Name:    "common.test",
		Routers: sohttp.NewRouters(handlers.CommonPostURIS, handlers.CommonPostHTTPMethods),
		Req:     &handlers.CommonPostReq{},
		Resp:    &handlers.CommonPostResp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlers.CommonPostHandle(ctx, req.(*handlers.CommonPostReq), resp.(*handlers.CommonPostResp))
		},
	},
}

// SoHTTP 获取 gnet http 服务
func SoHTTP() (*sohttp.SoHTTP, error) {
	s, err := sohttp.NewGnet([]int{8080})
	if err != nil {
		return nil, err
	}
	for _, h := range hs {
		s.Register(h)
	}
	return s, nil
}
