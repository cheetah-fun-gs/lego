package interfaces

// Module 模块 需要注册的模块定义
type Module interface {
	Start() error
	Stop() error
}
