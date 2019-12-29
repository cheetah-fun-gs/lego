// +build !windows

package gin

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func ginStart(engine *gin.Engine, addrs ...string) error {
	for _, addr := range addrs {
		if err := endless.ListenAndServe(addr, engine); err != nil {
			return err
		}
	}
	return nil
}
