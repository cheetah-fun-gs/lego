package sogin

import (
	"context"
)

// Handler sogin的处理器
type Handler struct {
	Name       string
	URI        string
	HTTPMethod string
	HandleFunc func(ctx context.Context, req interface{}) (resp interface{}, err error)
}

// GetName 获取处理器名称
func (h *Handler) GetName() string {
	return h.Name
}

// GetRouter 获取处理器路由
func (h *Handler) GetRouter() string {
	return h.URI
}

// Handle 处理器方法
func (h *Handler) Handle(ctx context.Context, req interface{}) (resp interface{}, err error) {
	return h.HandleFunc(ctx, req)
}
