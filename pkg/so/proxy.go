package so

// Proxy 获取handler的真正地址
type Proxy interface {
	GetPrivateData() interface{} // 私有数据 扩展用
}
