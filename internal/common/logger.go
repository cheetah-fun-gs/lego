package common

import (
	log4gopulus "github.com/cheetah-fun-gs/goplus/logger/log4go"
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	mlogger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
)

// ParseLogger 解析所有日志器配置
func ParseLogger() map[string]*log4gopulus.Config {
	loggers := map[string]*log4gopulus.Config{}

	ok, err := mconfiger.GetAnyN("logger", "logs", &loggers)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("logger logs not configure")
	}

	return loggers
}

// InitLogger 初始化日志器
func InitLogger(logs map[string]*log4gopulus.Config) {
	// 初始化默认日志
	if logConfig, ok := logs["default"]; !ok {
		panic("logs.default not configure")
	} else {
		logConfig.IsDebugMode = GlobalIsDebugMode // 以全局变量为准
		if logConfig.CallerDepth >= 0 {
			logConfig.CallerDepth++ // 使用 mlogger calldepth需要+1
		}
		mlogger.Init(log4gopulus.New("default", logConfig))
	}

	// 初始化其他日志
	for name, logConfig := range logs {
		if name != "default" {
			logConfig.IsDebugMode = GlobalIsDebugMode // 以全局变量为准
			if logConfig.CallerDepth >= 0 {
				logConfig.CallerDepth++ // 使用 mlogger calldepth需要+1
			}
			if err := mlogger.Register(name, log4gopulus.New(name, logConfig)); err != nil {
				panic(err)
			}
		}
	}
}
