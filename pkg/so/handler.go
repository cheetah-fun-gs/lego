package so

import (
	"context"
)

// Handler 处理器定义
type Handler interface {
	GetName() string   // 处理器名称
	GetRouter() string // 处理器路由
	// 处理器处理方法 req, resp 均为 go struct 的指针
	Handle(ctx context.Context, req, resp interface{}) error
}
