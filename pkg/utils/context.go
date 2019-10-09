package utils

import (
	"context"
	"time"

	"github.com/cheetah-fun-gs/goso/pkg/so"
)

var (
	contextDeadline = "__deadline"
)

// DumpContext 导出上下文
func DumpContext(ctx context.Context, keys []so.ContextKey) map[string]interface{} {
	r := map[string]interface{}{}
	d, ok := ctx.Deadline()
	if ok {
		r[contextDeadline] = d.Unix()
	}

	for _, k := range keys {
		if v := ctx.Value(k); v != nil {
			r[k.String()] = v
		}
	}
	return r
}

// LoadContext 加载上下文
func LoadContext(data map[string]interface{}, revert func(k string) so.ContextKey) context.Context {
	ctx := context.Background()
	var cancel func()

	for k, v := range data {
		if k == contextDeadline {
			d := time.Unix(v.(int64), 0)
			ctx, cancel = context.WithDeadline(ctx, d)
			defer cancel()
		} else if kk := revert(k); kk != nil {
			ctx = context.WithValue(ctx, kk, v)
		}
	}
	return ctx
}
