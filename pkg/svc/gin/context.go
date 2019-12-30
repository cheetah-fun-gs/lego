package gin

import (
	"context"
	"encoding/json"
	"time"
)

// ContextValue 上下文结构体
type ContextValue struct {
	RequestID string `json:"request_id,omitempty"` // 服务端ID
	Method    string `json:"method,omitempty"`
	Path      string `json:"path,omitempty"`
	RawQuery  string `json:"raw_query,omitempty"`
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

// ContextWithValue returns a new Context that carries value u.
func ContextWithValue(value *ContextValue) context.Context {
	ctx := context.Background()
	return context.WithValue(ctx, ctxKey, value)
}

// ContextWithRequestID returns a new Context that carries value u.
func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.RequestID = requestID
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{RequestID: requestID})
}

// ContextWithMethod returns a new Context that carries value u.
func ContextWithMethod(ctx context.Context, method string) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.Method = method
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{Method: method})
}

// ContextWithPath returns a new Context that carries value u.
func ContextWithPath(ctx context.Context, path string) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.Path = path
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{Path: path})
}

// ContextWithRawQuery returns a new Context that carries value u.
func ContextWithRawQuery(ctx context.Context, rawQuery string) context.Context {
	if val, ok := ContextFetchValue(ctx); ok {
		val.RawQuery = rawQuery
		return ctx
	}
	return context.WithValue(ctx, ctxKey, &ContextValue{RawQuery: rawQuery})
}

// ContextDump 导出
func ContextDump(ctx context.Context) *ContextData {
	var data = &ContextData{}

	d, ok := ctx.Deadline()
	if ok {
		data.Deadline = d.Unix()
	}

	if v := ctx.Value(&ContextValue{}); v != nil {
		data.Value = v.(*ContextValue)
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
