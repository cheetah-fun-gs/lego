package so

import (
	"context"
	"encoding/json"
	"time"
)

// ContextValue 上下文结构体
type ContextValue struct {
	Version string      `json:"version,omitempty"`  //
	TraceID string      `json:"trace_id,omitempty"` // 服务端ID
	SeqID   int         `json:"seq_id,omitempty"`   // 客户端ID
	Router  interface{} `json:"router,omitempty"`   // 路由器
}

// ContextData 传输用
type ContextData struct {
	Value    *ContextValue `json:"value,omitempty"`
	Deadline int64         `json:"deadline,omitempty"`
}

// String 添加 stringer 方法
func (v *ContextValue) String() string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userKey is the key for user.User values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var ctxKey key

// ContextFetchValue returns the User value stored in ctx, if any.
func ContextFetchValue(ctx context.Context) (*ContextValue, bool) {
	u, ok := ctx.Value(ctxKey).(*ContextValue)
	return u, ok
}

// ContextWithVersion returns a new Context that carries value u.
func ContextWithVersion(ctx context.Context, version string) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.Version = version
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{Version: version})
}

// ContextWithTraceID returns a new Context that carries value u.
func ContextWithTraceID(ctx context.Context, traceID string) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.TraceID = traceID
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{TraceID: traceID})
}

// ContextWithSeqID returns a new Context that carries value u.
func ContextWithSeqID(ctx context.Context, seqID int) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.SeqID = seqID
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{SeqID: seqID})
}

// ContextWithRouter returns a new Context that carries value u.
func ContextWithRouter(ctx context.Context, router interface{}) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.Router = router
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{Router: router})
}

// ContextDump 导出
func ContextDump(ctx context.Context) *ContextData {
	var data = &ContextData{}

	d, ok := ctx.Deadline()
	if ok {
		data.Deadline = d.Unix()
	}

	if v := ctx.Value(&ContextValue{}); v != nil {
		val := v.(*ContextValue)
		data.Value = &ContextValue{
			Router:  val.Router,
			Version: val.Version,
			TraceID: val.TraceID,
			SeqID:   val.SeqID,
		}
	}
	return data
}

// ContextLoad 导入
func ContextLoad(data *ContextData) context.Context {
	ctx := context.Background()
	var cancel func()

	if data.Deadline != 0 {
		d := time.Unix(data.Deadline, 0)
		ctx, cancel = context.WithDeadline(ctx, d)
		defer cancel()
	}

	if data.Value != nil {
		context.WithValue(ctx, &ContextValue{}, data.Value)
	}
	return ctx
}
