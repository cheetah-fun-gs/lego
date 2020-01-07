package scfgw

import (
	"context"
	"encoding/json"
	"fmt"

	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
)

// Server ...
type Server struct {
	handlers     map[string]legocore.Handler
	beforeHandle func(ctx context.Context, event Event, v interface{}) error
	behindHandle func(ctx context.Context, event Event, v interface{}) error
}

// New ...
func New(beforeHandle, behindHandle func(ctx context.Context, event Event, v interface{}) error, handlers ...legocore.Handler) *Server {
	server := &Server{
		handlers:     map[string]legocore.Handler{},
		beforeHandle: beforeHandle,
		behindHandle: behindHandle,
	}
	for _, h := range handlers {
		server.handlers[h.GetName()] = h
	}
	return server
}

type action struct {
	Action string `json:"action,omitempty"`
}

// Handle ...
func (server *Server) Handle(ctx context.Context, event Event) (resp interface{}, err error) {
	act := &action{}
	if err := json.Unmarshal(event.Body, act); err != nil {
		return nil, fmt.Errorf("act Unmarshal error: %v", err)
	}
	if act.Action == "" {
		return nil, fmt.Errorf("act is blank")
	}

	handler, ok := server.handlers[act.Action]
	if !ok {
		return nil, fmt.Errorf("act is not found: %v", act.Action)
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("act crash: %v, err: %v", act.Action, r)
			return
		}
	}()

	ctx = ContextWithAction(ctx, act.Action)

	req := handler.CloneReq()
	resp = handler.CloneResp()

	// 前置处理
	if server.beforeHandle != nil {
		if err := server.beforeHandle(ctx, event, req); err != nil {
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
	if server.behindHandle != nil {
		if err := server.behindHandle(ctx, event, resp); err != nil {
			return nil, fmt.Errorf("behindHandle err: %v", err)
		}
	}

	return resp, nil
}
