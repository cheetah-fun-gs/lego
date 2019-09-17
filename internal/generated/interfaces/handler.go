package interfaces

import (
	"context"
)

// Handler 处理器定义
type Handler struct {
	Route      string                                                 // 路由
	RouteAlias string                                                 // 路由别名
	Handle     func(ctx context.Context, req, resp interface{}) error // req, resp 均为 go struct 的指针
}
