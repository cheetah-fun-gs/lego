// +build !windows

package gin

import (
	"sync"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func ginStart(engine *gin.Engine, addrs ...string) {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(engine *gin.Engine, addr string) {
			defer wg.Done()
			if err := endless.ListenAndServe(addr, engine); err != nil {
				panic(err)
			}
		}(engine, addr)
	}
	wg.Wait()
}
