package core

import (
	"context"
)

// Handler 处理器定义
type Handler interface {
	GetName() string                                         // 名称
	CloneReq() interface{}                                   // 克隆一个请求结构体的新指针
	CloneResp() interface{}                                  // 克隆一个响应结构体的新指针
	Handle(ctx context.Context, req, resp interface{}) error // 处理方法
}

// DefaultHandler 默认处理器
type DefaultHandler struct {
	Name string
	Func func(ctx context.Context, req, resp interface{}) error
}

// GetName 获取处理器名称
func (h *DefaultHandler) GetName() string {
	return h.Name
}

// CloneReq 克隆请求结构体
func (h *DefaultHandler) CloneReq() interface{} {
	panic("Not Implement") // 需要各handle自己实现
}

// CloneResp 克隆响应结构体
func (h *DefaultHandler) CloneResp() interface{} {
	panic("Not Implement") // 需要各handle自己实现
}
