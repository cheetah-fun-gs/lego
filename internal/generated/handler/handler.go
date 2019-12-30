package handler

import (
	"context"
	"fmt"

	handlercommon "github.com/cheetah-fun-gs/lego/internal/biz/handler/common"
	"github.com/cheetah-fun-gs/lego/pkg/core"
)

// CommonPingHandler CommonPingHandler
type CommonPingHandler struct {
	*core.DefaultHandler
	SvcNames         []string
	beforeInjectFunc []func(ctx context.Context, req *handlercommon.PingReq)
	behindInjectFunc []func(ctx context.Context, req *handlercommon.PingReq, resp *handlercommon.PingResp)
}

// InjectBeforeFunc 注入前置函数
func (h *CommonPingHandler) InjectBeforeFunc(f ...func(ctx context.Context, req *handlercommon.PingReq)) {
	h.beforeInjectFunc = append(h.beforeInjectFunc, f...)
}

// InjectBehindFunc 注入后置函数
func (h *CommonPingHandler) InjectBehindFunc(f ...func(ctx context.Context, req *handlercommon.PingReq, resp *handlercommon.PingResp)) {
	h.behindInjectFunc = append(h.behindInjectFunc, f...)
}

// CloneReq CloneReq
func (h *CommonPingHandler) CloneReq() interface{} {
	return &handlercommon.PingReq{}
}

// CloneResp CloneResp
func (h *CommonPingHandler) CloneResp() interface{} {
	return &handlercommon.PingResp{}
}

// Handle 处理函数
func (h *CommonPingHandler) Handle(ctx context.Context, req, resp interface{}) error {
	reqBody, okReq := req.(*handlercommon.PingReq)
	respBody, okResp := resp.(*handlercommon.PingResp)
	if !okReq || !okResp {
		return fmt.Errorf("req or resp type error")
	}

	for _, f := range h.beforeInjectFunc {
		go f(ctx, reqBody)
	}
	if err := handlercommon.PingHandle(ctx, reqBody, respBody); err != nil {
		return err
	}
	for _, f := range h.behindInjectFunc {
		go f(ctx, reqBody, respBody)
	}
	return nil
}

// CommonPing handler
var CommonPing = &CommonPingHandler{
	DefaultHandler: &core.DefaultHandler{
		Name:    "CommonPing",
		Routers: genRouters(handlercommon.PingURIS, handlercommon.PingHTTPMethods),
	},
	SvcNames: handlercommon.SvcNames,
}

// Handlers 所有handler
var Handlers = []core.Handler{
	CommonPing,
}