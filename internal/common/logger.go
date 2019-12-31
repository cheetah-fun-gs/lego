package common

import (
	log4gopulus "github.com/cheetah-fun-gs/goplus/logger/log4go"
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	mlogger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
)

// InitLogger 初始化日志器
func InitLogger() {
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

	// 初始化 access 日志
	accessLoggerConfig := &log4gopulus.Config{}
	ok, err = mconfiger.GetStructN("logger", "logs.access", accessLoggerConfig)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("logger logs.access not configuration")
	}

	accessLoggerConfig.CallerDepth = -1                // 强行禁用 因为完全一致
	accessLoggerConfig.IsDebugMode = GlobalIsDebugMode // 以全局变量为准
	if err := mlogger.Register("access", log4gopulus.New("access", accessLoggerConfig)); err != nil {
		panic(err)
	}
}
