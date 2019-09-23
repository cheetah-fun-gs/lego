package handler

import (
	"context"
)

// Handler 默认处理器
type Handler struct {
	Name    string
	Routers []interface{} // 路由器
	Req     interface{}   // 请求结构体指针
	Resp    interface{}   // 响应结构体指针
	Func    func(ctx context.Context, req, resp interface{}) error
}

// GetName 获取处理器名称
func (h *Handler) GetName() string {
	return h.Name
}

// GetRouter 获取处理器路由
func (h *Handler) GetRouter() []interface{} {
	return h.Routers
}

// GetReq 获取请求结构体
func (h *Handler) GetReq() interface{} {
	return h.Req
}

// GetResp 获取响应结构体
func (h *Handler) GetResp() interface{} {
	return h.Resp
}

// Handle 处理器方法
func (h *Handler) Handle(ctx context.Context, req, resp interface{}) error {
	return h.Func(ctx, req, resp)
}

// GetPrivateData 获取私有数据
func (h *Handler) GetPrivateData() interface{} {
	return nil
}
