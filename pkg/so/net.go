package so

// ErrorNetCode so net error
type ErrorNetCode int

// so net error
const (
	ErrorNetCodeBadGateway ErrorNetCode = iota
	ErrorNetCodeBadRequest
)

// Net 网络传输服务
type Net interface {
	SetGatePack(gatePack GatePack) error // gnet 需要
	SetErrorNetFunc(errFunc func(code ErrorNetCode, err error) interface{}) error
	Register(handler Handler) error
	Start() error
	Stop() error
	GetPrivateData() interface{} // 私有数据 扩展用
}
