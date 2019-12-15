package gin

import (
	"context"
	"encoding/json"
)

// ContextValue ContextValue
type ContextValue struct {
	HTTPMethod string `json:"http_method,omitempty"`
	URI        string `json:"uri,omitempty"`
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

// ContextWithRouter returns a new Context that carries value u.
func ContextWithRouter(ctx context.Context, router *Router) context.Context {
	return context.WithValue(ctx, ctxKey, &ContextValue{HTTPMethod: router.HTTPMethod, URI: router.URI})
}

// ContextFetchValue returns the User value stored in ctx, if any.
func ContextFetchValue(ctx context.Context) (*ContextValue, bool) {
	u, ok := ctx.Value(ctxKey).(*ContextValue)
	return u, ok
}
