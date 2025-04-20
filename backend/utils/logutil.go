package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// LogLevel 定义日志级别
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

// Log 日志类结构体
type Log struct {
	level  LogLevel
	file   *os.File
	logger *log.Logger
	debug  bool
}

// NewLog 创建日志实例
func NewLog(levelStr string, outputPath string) (*Log, error) {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		return nil, err
	}

	var out io.Writer = os.Stdout
	if outputPath != "" {
		file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		out = io.MultiWriter(os.Stdout, file)
	}

	l := &Log{
		level:  level,
		logger: log.New(out, "", log.LstdFlags),
		debug:  level <= LevelDebug,
	}

	return l, nil
}

func parseLogLevel(level string) (LogLevel, error) {
	lvl := strings.ToLower(level)
	switch lvl {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	default:
		return LevelInfo, nil // 默认使用Info级别
	}
}

func (l *Log) getCallerInfo() (string, int) {
	_, file, line, ok := runtime.Caller(3) // 跳过3层调用栈
	if !ok {
		return "unknown", 0
	}
	return path.Base(file), line
}

func (l *Log) log(level LogLevel, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	file, line := l.getCallerInfo()
	msg := fmt.Sprintf(format, args...)
	levelStr := "INFO"
	switch level {
	case LevelDebug:
		levelStr = "DEBUG"
	case LevelWarn:
		levelStr = "WARN "
	case LevelError:
		levelStr = "ERROR"
	}

	output := fmt.Sprintf("[%s] [%s] %s:%d - %s\n", now, levelStr, file, line, msg)
	l.logger.Output(2, output) // 输出到日志
}

// Debugf 调试日志
func (l *Log) Debugf(format string, args ...interface{}) {
	l.log(LevelDebug, format, args...)
}

// Infof 信息日志
func (l *Log) Infof(format string, args ...interface{}) {
	l.log(LevelInfo, format, args...)
}

// Warnf 警告日志
func (l *Log) Warnf(format string, args ...interface{}) {
	l.log(LevelWarn, format, args...)
}

// Errorf 错误日志
func (l *Log) Errorf(format string, args ...interface{}) {
	l.log(LevelError, format, args...)
}

// Close 关闭日志文件
func (l *Log) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
