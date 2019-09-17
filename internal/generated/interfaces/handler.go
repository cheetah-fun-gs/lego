package interfaces

import (
	"context"
)

// Handler 处理器定义
type Handler interface {
	GetName() string
	GetRouter() string
	// req, resp 均为 go struct 的指针
	Handle(ctx context.Context, req, resp interface{}) error
}
