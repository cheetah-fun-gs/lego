package cmd

import (
	"sync"

	"github.com/cheetah-fun-gs/lego/internal/net/gnet"
	"github.com/cheetah-fun-gs/lego/internal/net/lnet"
	"github.com/cheetah-fun-gs/lego/internal/net/nnet"
)

// Execute 主函数
func Execute() {
	wg := sync.WaitGroup{}
	wg.Add(3)

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

	nnetSoHTTP, err := nnet.SoHTTP()
	if err != nil {
		panic(err)
	}

	go func() {
		nnetSoHTTP.Start()
		wg.Done()
	}()

	wg.Wait()
}
