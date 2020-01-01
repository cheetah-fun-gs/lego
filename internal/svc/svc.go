package svc

import (
	"fmt"
	"sync"

	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	"github.com/cheetah-fun-gs/goplus/structure"
	svcgin "github.com/cheetah-fun-gs/lego/internal/svc/gin"
)

func init() {
	ok, services, err := mconfiger.Get("services")
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("services not configature")
	}

	var wg sync.WaitGroup
	for _, v := range services.([]interface{}) {
		svcConfig := &Config{}
		vv := map[string]interface{}{}
		for key, val := range v.(map[interface{}]interface{}) {
			vv[key.(string)] = val
		}
		if err := structure.MapToStruct(vv, svcConfig); err != nil {
			panic(err)
		}
		switch svcConfig.Type {
		case TypeGin:
			wg.Add(1)
			go func() {
				defer wg.Done()
				svcgin.Start(svcConfig.Name, svcConfig.Addrs...)
			}()
		default:
			panic(fmt.Sprintf("type %v is not support", svcConfig.Type))
		}
	}
	wg.Wait()
}

// Type 服务类型
type Type string

// 常量
const (
	TypeGin Type = "gin"
)

// Config 配置
type Config struct {
	Name  string   `yml:"name,omitempty" json:"name,omitempty"`
	Type  Type     `yml:"type,omitempty" json:"type,omitempty"` // 服务类型
	Addrs []string `yml:"addrs,omitempty" json:"addrs,omitempty"`
}
