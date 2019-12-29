package core

// Router 路由器
type Router interface {
	GetURI() string
	GetHTTPMethod() string
}

// DefaultRouter 默认路由
type DefaultRouter struct {
	URI        string `json:"uri,omitempty"`
	HTTPMethod string `json:"http_method,omitempty"`
}

// GetURI GetURI
func (r *DefaultRouter) GetURI() string {
	return r.URI
}

// GetHTTPMethod GetHTTPMethod
func (r *DefaultRouter) GetHTTPMethod() string {
	return r.HTTPMethod
}
