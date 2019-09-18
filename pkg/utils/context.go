package utils

import (
	"context"
	"time"
)

// CtxKey ctx key
type CtxKey interface{}

// ctx 所用到的key
var (
	CtxKeyDeadline CtxKey = "deadline"
)

// DumpContext 导出上下文
func DumpContext(ctx context.Context, keys []CtxKey) map[string]interface{} {
	r := map[string]interface{}{}
	d, ok := ctx.Deadline()
	if ok {
		r[CtxKeyDeadline.(string)] = d.Unix()
	}
	for _, k := range keys {
		v := ctx.Value(CtxKey(k))
		if v != nil {
			r[k.(string)] = v
		}
	}
	return r
}

// LoadContext 加载上下文
func LoadContext(data map[string]interface{}) context.Context {
	ctx := context.Background()
	var cancel func()
	for k, v := range data {
		if k == CtxKeyDeadline.(string) {
			d := time.Unix(v.(int64), 0)
			ctx, cancel = context.WithDeadline(ctx, d)
			defer cancel()
		} else {
			ctx = context.WithValue(ctx, CtxKey(k), v)
		}
	}
	return ctx
}
