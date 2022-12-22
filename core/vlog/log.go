package vlog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = 0
	LogLevelInfo  LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 3
	LogLevelPanic LogLevel = 4
)

var (
	LogLevelNameMap = map[LogLevel]string{
		LogLevelDebug: "debug",
		LogLevelInfo:  "info",
		LogLevelWarn:  "warn",
		LogLevelError: "error",
		LogLevelPanic: "panic",
	}
)

type ILog interface {
	LogDebug(format string, v ...interface{})
	LogInfo(format string, v ...interface{})
	LogWarn(format string, v ...interface{})
	LogErr(format string, v ...interface{})
	LogPanic(format string, v ...interface{})
}

type Log struct {
	FilePrefix string // 日志名称前缀
	FileSize   int    // 日志文件最大占用空间
	FileCount  int    // 单日最多日志文件数
	LogMask    int    // 日志级别
	sync.Mutex
}

func NewLogEngine(filePrefix string, fileSize, fileCount int, logMask int) ILog {
	return &Log{
		FilePrefix: filePrefix,
		FileSize:   fileSize,
		FileCount:  fileCount,
		LogMask:    logMask,
	}
}

func (l *Log) LogDebug(format string, v ...interface{}) {
	l.log(LogLevelDebug, format, v)
}

func (l *Log) LogInfo(format string, v ...interface{}) {
	l.log(LogLevelInfo, format, v)
}

func (l *Log) LogWarn(format string, v ...interface{}) {
	l.log(LogLevelWarn, format, v)
}

func (l *Log) LogErr(format string, v ...interface{}) {
	l.log(LogLevelError, format, v)
}

func (l *Log) LogPanic(format string, v ...interface{}) {
	l.log(LogLevelPanic, format, v)
}

func (l *Log) ShouldDoLog(logLevel LogLevel) bool {
	return 0 != (l.LogMask & (1 << logLevel))
}

func (l *Log) log(logLevel LogLevel, format string, v ...interface{}) {
	if !l.ShouldDoLog(logLevel) {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	l.Lock()
	defer l.Unlock()

	curDate := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s_%s.log", l.FilePrefix, curDate)
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}

	logLevelName := LogLevelNameMap[logLevel]
	curTime := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fd.WriteString(fmt.Sprintf("[%s][%s][%s]", curTime, logLevelName, FileLine()))
	_, _ = fd.WriteString(fmt.Sprintf(format+"\n", v...))
	_ = fd.Close()

	if runtime.GOOS == "windows" {
		fmt.Printf("[[%s]][%s]", curTime, logLevelName)
		fmt.Printf(format+"\n", v...)
	}

	l.shiftFile(fileName)
}

func (l *Log) shiftFile(filename string) {
	info, err := os.Stat(filename)
	if err != nil {
		return
	}
	if info.Size() < int64(l.FileSize) {
		return
	}

	extFileName := path.Ext(filename)
	baseFileName := strings.TrimSuffix(filename, extFileName)
	for i := l.FileCount - 1; i >= 0; i-- {
		tmpFileName := fmt.Sprintf("%s_%d%s", baseFileName, i, extFileName)
		newFileName := fmt.Sprintf("%s_%d%s", baseFileName, i+1, extFileName)
		if i == 0 {
			_ = os.Rename(filename, newFileName)
		}
		if _, err := os.Stat(tmpFileName); err != nil {
			continue
		}
		if i+1 == l.FileCount {
			_ = os.Remove(tmpFileName)
		} else {
			_ = os.Rename(tmpFileName, newFileName)
		}
	}
}

// 获取到代码所在的文件名、行数
func FileLine() string {
	if _, fileName, fileLine, ok := runtime.Caller(3); ok {
		return fmt.Sprintf("%s:%d", path.Base(fileName), fileLine)
	}
	return ""
}
