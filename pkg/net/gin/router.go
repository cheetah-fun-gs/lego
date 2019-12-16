package gin

import "encoding/json"

// Router 路由器
type Router struct {
	Method string `json:"method,omitempty"`
	Path   string `json:"path,omitempty"`
}

func (router *Router) String() string {
	data, err := json.Marshal(router)
	if err != nil {
		return ""
	}
	return string(data)
}
