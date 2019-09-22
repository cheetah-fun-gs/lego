package so

import (
	"context"
)

// Proxy 获取handler或caller
type Proxy interface {
	RegisterHandler(hander Handler) error                                                   // 注册
	IsHandler(ctx context.Context, name string) (bool, error)                               // 判断是否本进程内能处理
	FetchOneCaller(ctx context.Context, name string) (caller Caller, err error)             // 按照默认策略获取一个caller
	FetchTheCaller(ctx context.Context, name string, addr *Addr) (caller Caller, err error) // 获取一个指定caller
	FetchAllCaller(ctx context.Context, name string) (callers []Caller, err error)          // 获取所有caller
	GetPrivateData() interface{}                                                            // 私有数据 扩展用
}
