package scfgw

import (
	"context"
	"encoding/json"
	"fmt"

	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
)

type handlerName struct {
	HandlerName string `json:"handler_name,omitempty"`
}

// ParseHandlerName ...
func ParseHandlerName(ctx context.Context, event Event) (string, error) {
	hn := &handlerName{}
	if err := json.Unmarshal(event.Body, hn); err != nil {
		return "", fmt.Errorf("act Unmarshal error: %v", err)
	}
	if hn.HandlerName == "" {
		return "", fmt.Errorf("handler_name is blank")
	}

	ctx = ContextWithHandlerName(ctx, hn.HandlerName)
	return hn.HandlerName, nil
}

// Handle ...
func Handle(ctx context.Context, event *Event,
	beforeHandle, behindHandle func(ctx context.Context, event *Event, v interface{}) error,
	handler legocore.Handler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("crash err: %v", r)
			return
		}
	}()

	req := handler.CloneReq()
	resp = handler.CloneResp()

	// 前置处理
	if beforeHandle != nil {
		if err := beforeHandle(ctx, event, req); err != nil {
			return nil, fmt.Errorf("beforeHandle err: %v", err)
		}
	} else {
		req = event
	}

	// 处理
	if err := handler.Handle(ctx, req, resp); err != nil {
		return nil, fmt.Errorf("Handle err: %v", err)
	}

	// 后置处理
	if behindHandle != nil {
		if err := behindHandle(ctx, event, resp); err != nil {
			return nil, fmt.Errorf("behindHandle err: %v", err)
		}
	}

	return resp, nil
}
