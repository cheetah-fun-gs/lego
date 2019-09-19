package so

// Module 模块 需要注册的模块定义
type Module interface {
	Start() error
	Stop() error
	GetPrivateData() interface{} // 私有数据 扩展用
}
