package so

import "context"

// GatePack 网关包对象
type GatePack interface {
	Verify() error               // 校验包的有效性
	GetRouter() Router           // 获取路由
	GetLogicPack() interface{}   // 获取业务对象
	GetContext() context.Context // 获取上下文
}
