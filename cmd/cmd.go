package cmd

import (
	"goso/internal/generated/net/gnet"
	"goso/internal/generated/net/lnet"
	"sync"
)

// Run 主函数
func Run() {
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
