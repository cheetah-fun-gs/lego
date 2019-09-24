package lnet

import (
	"context"
	"github.com/cheetah-fun-gs/goso/internal/biz/handlers"
	"github.com/cheetah-fun-gs/goso/pkg/handler"
	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
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
		Name:    "common.post",
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
	s, err := sohttp.NewLnet()
	if err != nil {
		return nil, err
	}
	// 先设置 config
	s.SetConfig(&sohttp.Config{
		Ports:        []int{8000},
		HTTPCodeFunc: handlers.HandleCommonRespSoHTTP,
	})
	// 再注册 handler
	for _, h := range hs {
		s.Register(h)
	}
	return s, nil
}
