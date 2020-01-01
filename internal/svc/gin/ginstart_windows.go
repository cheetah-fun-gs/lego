// +build windows

package gin

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func ginStart(engine *gin.Engine, addrs ...string) {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(engine *gin.Engine, addr string) {
			defer wg.Done()
			svc := &http.Server{
				Addr:           addr,
				Handler:        engine,
				ReadTimeout:    8 * time.Second,
				WriteTimeout:   8 * time.Second,
				MaxHeaderBytes: 1 << 16,
			}
			if err := svc.ListenAndServe(); err != nil {
				panic(err)
			}
		}(engine, addr)
	}
	wg.Wait()
}
