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

	gnetSoGin, err := gnet.SoGin()
	if err != nil {
		panic(err)
	}

	go func() {
		gnetSoGin.Start()
		wg.Done()
	}()

	lnetSoGin, err := lnet.SoGin()
	if err != nil {
		panic(err)
	}

	go func() {
		lnetSoGin.Start()
		wg.Done()
	}()

	wg.Wait()
}
