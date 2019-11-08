package so

import (
	"context"
)

// Handler 处理器定义
type Handler interface {
	IsAnyNet(netType NetType) bool // 是否某个网络的处理器
	GetName() string               // 名称
	GetRouter() (routers []Router) // 路由器对象 允许多个路由器指向相同处理器 router 如果是结构体必须提供 String() 方法
	CloneReq() interface{}         // 克隆一个请求结构体的新指针
	CloneResp() interface{}        // 克隆一个响应结构体的新指针
	Handle(ctx context.Context, req, resp interface{}) error
}
