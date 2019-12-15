package handler

import (
	"context"

	"github.com/cheetah-fun-gs/lego/pkg/so"
)

// Handler 默认处理器
type Handler struct {
	Name    string
	Nets    []so.NetType
	Routers []so.Router // 路由器
	Func    func(ctx context.Context, req, resp interface{}) error
}

// IsAnyNet 是否某个网络
func (h *Handler) IsAnyNet(netType so.NetType) bool {
	for _, t := range h.Nets {
		if t == netType {
			return true
		}
	}
	return false
}

// GetName 获取处理器名称
func (h *Handler) GetName() string {
	return h.Name
}

// GetRouter 获取处理器路由
func (h *Handler) GetRouter() []so.Router {
	return h.Routers
}

// CloneReq 克隆请求结构体
func (h *Handler) CloneReq() interface{} {
	panic("Not Implement") // 需要各handle自己实现
}

// CloneResp 克隆响应结构体
func (h *Handler) CloneResp() interface{} {
	panic("Not Implement") // 需要各handle自己实现
}

// Handle 处理器方法
func (h *Handler) Handle(ctx context.Context, req, resp interface{}) error {
	return h.Func(ctx, req, resp)
}
