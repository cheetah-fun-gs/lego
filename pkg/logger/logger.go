package logger

import (
	"context"
	"fmt"
	"log"
)

// Logger 默认日志器
type Logger struct {
	DebugMode bool // 是否 debug 模式
}

// Debug 级别日志
func (logger *Logger) Debug(ctx context.Context, format string, v ...interface{}) {
	if logger.DebugMode {
		log.SetPrefix(fmt.Sprintf("[Debug] %v ", ctx))
		log.Printf(format, v...)
	}
	return
}

// Info 级别日志
func (logger *Logger) Info(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix(fmt.Sprintf("[Info] %v ", ctx))
	log.Printf(format, v...)
	return
}

// Warn 级别日志
func (logger *Logger) Warn(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix(fmt.Sprintf("[Warn] %v ", ctx))
	log.Printf(format, v...)
	return
}

// Error 级别日志
func (logger *Logger) Error(ctx context.Context, format string, v ...interface{}) {
	log.SetPrefix(fmt.Sprintf("[Error] %v ", ctx))
	log.Printf(format, v...)
	return
}

// New 一个新的日志器
func New() *Logger {
	return &Logger{}
}
