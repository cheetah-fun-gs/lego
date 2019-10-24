package so

import (
	"context"
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

// ContextWithVersion 添加上下文字段 会覆盖
func ContextWithVersion(ctx context.Context, version string) context.Context {
	val := ctx.Value(&ContextValue{})
	if val != nil {
		v := val.(*ContextValue)
		v.Version = version
		return ctx
	}
	return context.WithValue(ctx, &ContextValue{}, &ContextValue{Version: version})
}

// ContextWithTraceID 添加上下文字段 会覆盖
func ContextWithTraceID(ctx context.Context, traceID string) context.Context {
	val := ctx.Value(&ContextValue{})
	if val != nil {
		v := val.(*ContextValue)
		v.TraceID = traceID
		return ctx
	}
	return context.WithValue(ctx, &ContextValue{}, &ContextValue{TraceID: traceID})
}

// ContextWithSeqID 添加上下文字段 会覆盖
func ContextWithSeqID(ctx context.Context, seqID int) context.Context {
	val := ctx.Value(&ContextValue{})
	if val != nil {
		v := val.(*ContextValue)
		v.SeqID = seqID
		return ctx
	}
	return context.WithValue(ctx, &ContextValue{}, &ContextValue{SeqID: seqID})
}

// ContextWithRouter 添加上下文字段 会覆盖
func ContextWithRouter(ctx context.Context, router interface{}) context.Context {
	val := ctx.Value(&ContextValue{})
	if val != nil {
		v := val.(*ContextValue)
		v.Router = router
		return ctx
	}
	return context.WithValue(ctx, &ContextValue{}, &ContextValue{Router: router})
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
