package handler

import (
	"context"
	"fmt"

	jsonplus "github.com/cheetah-fun-gs/goplus/encoding/json"
	mlogger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
	handlercommon "github.com/cheetah-fun-gs/lego/internal/biz/handler/common"
	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
)

// CommonPingHandler CommonPingHandler
type CommonPingHandler struct {
	*legocore.DefaultHandler
	SvcNames         []string
	HTTPMethods      []string
	HTTPPaths        []string
	beforeInjectFunc []func(ctx context.Context, req handlercommon.PingReq)                              // 注入不能使用指针
	behindInjectFunc []func(ctx context.Context, req handlercommon.PingReq, resp handlercommon.PingResp) // 注入不能使用指针
}

// InjectBeforeFunc 注入前置函数
func (h *CommonPingHandler) InjectBeforeFunc(f ...func(ctx context.Context, req handlercommon.PingReq)) {
	h.beforeInjectFunc = append(h.beforeInjectFunc, f...)
}

// InjectBehindFunc 注入后置函数
func (h *CommonPingHandler) InjectBehindFunc(f ...func(ctx context.Context, req handlercommon.PingReq, resp handlercommon.PingResp)) {
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

	// 前置注入
	if len(h.beforeInjectFunc) > 0 {
		beforeInjectReq := handlercommon.PingReq{}
		if err := jsonplus.Convert(reqBody, &beforeInjectReq); err != nil {
			mlogger.Warnc(ctx, "jsonplus.Convert err: %v", err)
		} else {
			for _, f := range h.beforeInjectFunc {
				go f(ctx, beforeInjectReq)
			}
		}
	}

	// 真正处理
	if err := handlercommon.PingHandle(ctx, reqBody, respBody); err != nil {
		return err
	}

	// 后置注入
	if len(h.behindInjectFunc) > 0 {
		behindInjectReq := handlercommon.PingReq{}
		behindInjectResp := handlercommon.PingResp{}
		errReq := jsonplus.Convert(reqBody, &behindInjectReq)
		errResp := jsonplus.Convert(respBody, &behindInjectResp)
		if errReq != nil || errResp != nil {
			mlogger.Warnc(ctx, "jsonplus.Convert errReq: %v, errResp: %v", errReq, errResp)
		} else {
			for _, f := range h.behindInjectFunc {
				go f(ctx, behindInjectReq, behindInjectResp)
			}
		}
	}
	return nil
}

// CommonPing handler
var CommonPing = &CommonPingHandler{
	DefaultHandler: &legocore.DefaultHandler{
		Name: "CommonPing",
	},
	SvcNames:    handlercommon.PingSvcNames,
	HTTPMethods: handlercommon.PingHTTPMethods,
	HTTPPaths:   handlercommon.PingHTTPPaths,
}

// Handlers 所有handler
var Handlers = map[string]legocore.Handler{
	"CommonPing": CommonPing,
}
