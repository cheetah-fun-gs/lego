package handler

import (
	"github.com/cheetah-fun-gs/lego/pkg/core"
)

func genRouters(uris, httpMethods []string) []core.Router {
	rs := []core.Router{}
	for _, uri := range uris {
		for _, method := range httpMethods {
			rs = append(rs, &core.DefaultRouter{
				URI:        uri,
				HTTPMethod: method,
			})
		}
	}
	return rs
}
