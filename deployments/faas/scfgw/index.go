package main

import (
	"context"
	"fmt"

	_ "github.com/cheetah-fun-gs/lego/cmd/scfgw"
	"github.com/cheetah-fun-gs/lego/internal/common"
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

	// 根据 action 加载资源
	logs := common.ParseLogger()
	for _, logger := range logs {
		logger.Path = ""                // 强行关闭 scf 不需要文件日志
		logger.IsDisableConsole = false // 强行打开
	}

	redigos := common.ParseRedigo()
	if len(redigos) > 0 {
		common.InitRedigo(redigos)
	}

	sqls := common.ParseSQLDB()
	if len(sqls) > 0 {
		common.InitSQLDB(sqls)
	}

	mgos := common.ParseMgo()
	if len(mgos) > 0 {
		common.InitMgo(mgos)
	}

	return legoscfgw.Handle(ctx, event,
		svcscfgw.BeforeHandleFunc, svcscfgw.BehindHandleFunc, handler)
}

func main() {
	cloudfunction.Start(hello)
}
