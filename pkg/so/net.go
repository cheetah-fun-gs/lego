package so

// Net 网络传输服务
type Net interface {
	SetGatePack(gatePack GatePack) error // gnet 需要
	Register(handler Handler) error
	Start() error
	Stop() error
}
