package common

import (
	"github.com/cheetah-fun-gs/goso/pkg/logger"
)

// GlobalLogger 全局日志器
var GlobalLogger = logger.New()

// 端口定义
var (
	PortsHTTPGNet = []int{8080}
	PortsHTTPLNet = []int{8000}
	PortsHTTPNNet = []int{8086}
)
