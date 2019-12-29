// +build windows

package gin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ginStart(engine *gin.Engine, addrs ...string) error {
	for _, addr := range addrs {
		svc := &http.Server{
			Addr:           addr,
			Handler:        engine,
			ReadTimeout:    8 * time.Second,
			WriteTimeout:   8 * time.Second,
			MaxHeaderBytes: 1 << 16,
		}
		if err := svc.ListenAndServe(); err != nil {
			return err
		}
	}
	return nil
}
