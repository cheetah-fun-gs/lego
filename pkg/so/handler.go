package so

import (
	"context"
)

// Handler 处理器定义
type Handler interface {
	GetName() string                    // 名称
	GetRouter() (routers []interface{}) // 路由器对象 允许多个路由器指向相同处理器 router 如果是结构体必须提供 String() 方法
	GetReq() interface{}                // 请求结构体 空结构体 指针
	GetResp() interface{}               // 响应结构体 空结构体 指针
	GetPrivateData() interface{}        // 私有数据 扩展用
	Handle(ctx context.Context, req, resp interface{}) error
}
