package so

import (
	"context"
)

// Handler 处理器定义
type Handler interface {
	GetName() string                                         // 名称
	GetRouter() string                                       // 路由
	GetReq() interface{}                                     // 请求结构体 空结构体 指针
	GetResp() interface{}                                    // 响应结构体 空结构体 指针
	Handle(ctx context.Context, req, resp interface{}) error // 方法 req, resp 均为 go struct 的指针
	GetPrivateData() interface{}                             // 私有数据 扩展用
}
