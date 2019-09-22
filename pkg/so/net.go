package so

// Net 网络传输服务
type Net interface {
	Register(handler Handler) error
	Start() error
	Stop() error
	GetPrivateData() interface{} // 私有数据 扩展用
}
