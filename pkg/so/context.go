package so

// ContextKey  ContextKey
type ContextKey interface {
	String() string
}

// ContextTraceID 服务端包ID
type ContextTraceID struct{}

func (ctx *ContextTraceID) String() string {
	return "trace_id"
}

// ContextSeqID 客户端包
type ContextSeqID struct{}

func (ctx *ContextSeqID) String() string {
	return "seq_id"
}

// ContextVersion 版本
type ContextVersion struct{}

func (ctx *ContextVersion) String() string {
	return "version"
}

// ContextRouter 路由
type ContextRouter struct{}

func (ctx *ContextRouter) String() string {
	return "router"
}

// ContextRevert 还原
func ContextRevert(k string) ContextKey {
	for _, c := range []ContextKey{
		&ContextTraceID{},
		&ContextSeqID{},
		&ContextVersion{},
		&ContextRouter{},
	} {
		if c.String() == k {
			return c
		}
	}
	return nil
}
