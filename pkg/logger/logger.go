package logger

import (
	"context"
	"log"
)

func logPrintf(ctx context.Context, format string, v ...interface{}) {
	log.Printf("%v "+format, append([]interface{}{ctx}, v...)...)
}

// Logger 默认日志器
type Logger struct {
	IsDebugMode bool // 是否 debug 模式
}

// Debug 级别日志
func (logger *Logger) Debug(ctx context.Context, format string, v ...interface{}) {
	if logger.IsDebugMode {
		log.SetPrefix("[Debug] ")
		logPrintf(ctx, format, v...)
	}
	return
}

// Info 级别日志
func (logger *Logger) Info(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix("[Info] ")
	logPrintf(ctx, format, v...)
	return
}

// Warn 级别日志
func (logger *Logger) Warn(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix("[Warn] ")
	logPrintf(ctx, format, v...)
	return
}

// Error 级别日志
func (logger *Logger) Error(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix("[Error] ")
	logPrintf(ctx, format, v...)
	return
}

// New 一个新的日志器
func New() *Logger {
	return &Logger{}
}
