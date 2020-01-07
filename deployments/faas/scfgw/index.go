package main

import (
	"context"
	"fmt"

	allhandler "github.com/cheetah-fun-gs/lego/internal/generated/handler"
	svcscfgw "github.com/cheetah-fun-gs/lego/internal/svc/scfgw"
	legoscfgw "github.com/cheetah-fun-gs/lego/pkg/svc/scfgw"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

func hello(ctx context.Context, event legoscfgw.Event) (interface{}, error) {
	action, err := legoscfgw.ParseAction(ctx, event)
	if err != nil {
		return nil, err
	}

	handler, ok := allhandler.Handlers[action]
	if !ok {
		return nil, fmt.Errorf("action %v is not found", action)
	}

	return legoscfgw.Handle(ctx, event,
		svcscfgw.BeforeHandleFunc, svcscfgw.BehindHandleFunc, handler)
}

func main() {
	cloudfunction.Start(hello)
}
