package interfaces

// NetAttr 网络传输服务属性
type NetAttr struct {
	Ports []int
}

// Net 网络传输服务
type Net interface {
	SetAttr(attr *NetAttr)
	Register(handler Handler)
	Start() error
	Stop() error
}
