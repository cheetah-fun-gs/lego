package so

import (
	"context"
)

// HandlerFunc 处理器方法 req, resp 均为 go struct 的指针
type HandlerFunc func(ctx context.Context, req, resp interface{}) error

// Handler 处理器定义
type Handler interface {
	GetName() string             // 名称
	GetRouter() interface{}      // 路由器对象 用来建立唯一索引
	GetReq() interface{}         // 请求结构体 空结构体 指针
	GetResp() interface{}        // 响应结构体 空结构体 指针
	GetPrivateData() interface{} // 私有数据 扩展用
	Func() HandlerFunc
}
