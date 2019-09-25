package so

// GNet 网络传输服务
type GNet interface {
	SetGatePack(gatePack GatePack) error
	Register(handler Handler) error
	Start() error
	Stop() error
	GetPrivateData() interface{} // 私有数据 扩展用
}

// LNet 网络传输服务
type LNet interface {
	Register(handler Handler) error
	Start() error
	Stop() error
	GetPrivateData() interface{} // 私有数据 扩展用
}
