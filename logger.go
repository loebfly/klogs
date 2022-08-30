package klogs

import (
	"fmt"
	"runtime"
	"time"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
)

// Logger 日志接口
type Logger struct {
	category string
	level    string
}

func (logger Logger) Debug(format string, args ...interface{}) {
	logger.level = LevelDebug
	logger.outPut(format, args...)
}

func (logger Logger) Info(format string, args ...interface{}) {
	logger.level = LevelInfo
	logger.outPut(format, args...)
}

func (logger Logger) Warn(format string, args ...interface{}) {
	logger.level = LevelWarn
	logger.outPut(format, args...)
}

func (logger Logger) Error(format string, args ...interface{}) {
	logger.level = LevelError
	logger.outPut(format, args...)
}

func (logger Logger) outPut(format string, args ...interface{}) {
	// Determine caller func
	createdAt := time.Now()
	pc, _, line, ok := runtime.Caller(2)
	src := "(UNKNOWN)"
	if ok {
		src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line)
	}

	msg := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] [%s] (%s) [%s] %s\n",
		createdAt.Format("2006-01-02 15:04:05"),
		logger.category,
		logger.level,
		src, msg)
}
