package so

import "context"

// Caller 远程调用handler的封装
type Caller interface {
	Call(ctx context.Context, req, resp interface{}) error
	GetPrivateData() interface{} // 私有数据 扩展用
}
