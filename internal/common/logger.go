package common

import (
	log4gopulus "github.com/cheetah-fun-gs/goplus/logger/log4go"
	mconfiger "github.com/cheetah-fun-gs/goplus/multier/multiconfiger"
	mlogger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
	"github.com/cheetah-fun-gs/goplus/structure"
)

// ParseLogger 解析所有日志器配置
func ParseLogger() map[string]*log4gopulus.Config {
	_, logs, err := mconfiger.GetMapN("logger", "logs")
	if err != nil {
		panic(err)
	}

	loggers := map[string]*log4gopulus.Config{}
	for name, data := range logs {
		logConfig := &log4gopulus.Config{}
		if err := structure.MapToStruct(data.(map[string]interface{}), logConfig); err != nil {
			panic(err)
		}
		loggers[name] = logConfig
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
		mlogger.Init(log4gopulus.New("default", logConfig))
	}

	// 初始化其他日志
	for name, logConfig := range logs {
		if name != "default" {
			logConfig.IsDebugMode = GlobalIsDebugMode // 以全局变量为准
			if err := mlogger.Register(name, log4gopulus.New(name, logConfig)); err != nil {
				panic(err)
			}
		}
	}
}
