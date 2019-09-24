package logger

import (
	"context"
	"github.com/cheetah-fun-gs/goso/pkg/so"
	"log"
)

// Logger 默认日志器
type Logger struct {
	DebugMode bool // 是否 debug 模式
}

// Debug 级别日志
func (logger *Logger) Debug(ctx context.Context, format string, v ...interface{}) {
	if logger.DebugMode {
		log.SetPrefix("[Debug] ")
		log.Printf(format, v...)
	}
	return
}

// Info 级别日志
func (logger *Logger) Info(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix("[Info] ")
	log.Printf(format, v...)
	return
}

// Warn 级别日志
func (logger *Logger) Warn(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix("[Warn] ")
	log.Printf(format, v...)
	return
}

// Error 级别日志
func (logger *Logger) Error(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix("[Error] ")
	log.Printf(format, v...)
	return
}

// New 一个新的日志器
func New() so.Logger {
	return &Logger{}
}
