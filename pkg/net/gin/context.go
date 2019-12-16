package gin

import (
	"context"
)

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userKey is the key for user.User values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var ctxKey key

// ContextWithRouter returns a new Context that carries value u.
func ContextWithRouter(ctx context.Context, router *Router) context.Context {
	return context.WithValue(ctx, ctxKey, router)
}

// ContextFetchValue returns the User value stored in ctx, if any.
func ContextFetchValue(ctx context.Context) (*Router, bool) {
	u, ok := ctx.Value(ctxKey).(*Router)
	return u, ok
}
