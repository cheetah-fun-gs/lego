package cmd

import (
	"sync"

	"github.com/cheetah-fun-gs/goso/internal/generated/net/gnet"
	"github.com/cheetah-fun-gs/goso/internal/generated/net/lnet"
)

// Execute 主函数
func Execute() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	gnetSoHTTP, err := gnet.SoHTTP()
	if err != nil {
		panic(err)
	}

	go func() {
		gnetSoHTTP.Start()
		wg.Done()
	}()

	lnetSoHTTP, err := lnet.SoHTTP()
	if err != nil {
		panic(err)
	}

	go func() {
		lnetSoHTTP.Start()
		wg.Done()
	}()

	wg.Wait()
}
