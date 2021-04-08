package meilogrus

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/andytyc/goutils/meifile"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// MeiFileLog MeiFileLog
type MeiFileLog struct {
	// 日志名称 | 默认{all.log}
	LogFileName string

	// 日志级别 | 默认{Info}
	LogLevel logrus.Level

	// 日志路径 | 默认{pwdDir + "/logs/ + LogFileName}
	LogFilePath        string
	FuncGetLogFilePath func(logFileName string) (logFilePath string, err error)

	// 日志格式 | 默认{时间戳，级别，文件:行数，消息}
	LogFormatter logrus.Formatter

	// 日志写入器 | 默认{控制台，文件输出：日志每隔 8 小时轮转一个新文件，保留最近 7 天的日志文件}
	LogFileWriter        io.Writer
	FuncGetLogFileWriter func(logFilePath string) (writer io.Writer, err error)

	// 日志退出回调函数 | 默认{采用logrus默认回调：os.Exit}
	FuncLogExit func(code int)
}

// GetLogger 创建并返回日志实例
func (m *MeiFileLog) GetLogger() (logger *logrus.Logger, err error) {
	err = m.checkFileds()
	if err != nil {
		return
	}

	logger = logrus.New()
	logger.SetOutput(m.LogFileWriter)
	logger.SetLevel(m.LogLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(m.LogFormatter)
	if m.FuncLogExit != nil {
		logger.ExitFunc = m.FuncLogExit
	}

	return
}

func (m *MeiFileLog) checkFileds() (err error) {
	if m.LogFileName == "" {
		m.LogFileName = DefaultLogFileName
	}

	if m.LogFilePath == "" {
		if m.FuncGetLogFilePath == nil {
			m.FuncGetLogFilePath = m.DefaultFuncGetLogFilePath
		}
		m.LogFilePath, err = m.FuncGetLogFilePath(m.LogFileName)
		if err != nil {
			return
		}
	}

	if m.LogLevel == 0 {
		m.LogLevel = logrus.InfoLevel
	}

	if m.LogFormatter == nil {
		m.LogFormatter = &DefaultLogFormatter{}
	}

	if m.LogFileWriter == nil {
		if m.FuncGetLogFileWriter == nil {
			m.FuncGetLogFileWriter = m.DefaultFuncGetLogFileWriter
		}
		m.LogFileWriter, err = m.FuncGetLogFileWriter(m.LogFilePath)
		if err != nil {
			return
		}
	}

	return
}

// *****************************************************************************
// Default
// *****************************************************************************

var (
	DefaultLogFileName = "all.log"
)

// 默认日志路径
func (m *MeiFileLog) DefaultFuncGetLogFilePath(logFileName string) (logFilePath string, err error) {
	var pwdDir string
	pwdDir, err = os.Getwd()
	if err != nil {
		return
	}
	logFilePath = pwdDir + "/logs/" + logFileName
	err = meifile.CreateFile(logFilePath)
	if err != nil {
		return
	}
	return
}

// DefaultLogFormatter 默认日志格式
type DefaultLogFormatter struct{}

// Format 时间戳，级别，文件:行数，消息
func (s *DefaultLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	loglevel := strings.ToUpper(entry.Level.String())
	msg := fmt.Sprintf("%s [%s] [%s:%d] %s\n", timestamp, loglevel, entry.Caller.Function, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}

// DefaultFuncGetLogFileWriter 默认输出{文件、控制台}
// 文件：每隔 24 小时轮转一个新文件，保留最近 7 天的日志文件，多余的自动清理掉。
// 年月日时分秒：logFilePath+".%Y%m%d%H%M%S"
func (m *MeiFileLog) DefaultFuncGetLogFileWriter(logFilePath string) (writer io.Writer, err error) {
	fileWriter, err := rotatelogs.New(
		logFilePath+".%Y%m%d",
		rotatelogs.WithLinkName(logFilePath),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		rotatelogs.WithMaxAge(time.Duration(24*7)*time.Hour),
	)
	if err != nil {
		return
	}

	writers := []io.Writer{fileWriter, os.Stdout}
	writer = io.MultiWriter(writers...)

	return
}
