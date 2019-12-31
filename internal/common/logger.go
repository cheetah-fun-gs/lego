package common

import (
	log4gopulus "github.com/cheetah-fun-gs/goplus/logger/log4go"
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	mlogger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
	"github.com/cheetah-fun-gs/goplus/structure"
)

// initLogger 初始化日志器
func initLogger() {
	// 初始化默认日志
	defaultLoggerConfig := &log4gopulus.Config{}
	ok, err := mconfiger.GetStructN("logger", "logs.default", defaultLoggerConfig)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("logger logs.default not configuration")
	}

	defaultLoggerConfig.IsDebugMode = GlobalIsDebugMode // 以全局变量为准
	mlogger.Init(log4gopulus.New("default", defaultLoggerConfig))

	// 初始化其他日志
	_, logs, _ := mconfiger.GetMapN("logger", "logs")
	for name, data := range logs {
		if name != "default" {
			logConfig := &log4gopulus.Config{}
			if err := structure.MapToStruct(data.(map[string]interface{}), logConfig); err != nil {
				panic(err)
			}

			logConfig.IsDebugMode = GlobalIsDebugMode // 以全局变量为准
			if err := mlogger.Register(name, log4gopulus.New(name, logConfig)); err != nil {
				panic(err)
			}
		}
	}
}
