package logger

import (
	f "github.com/madeiramadeirabr/lib-go-logger/pkg/formatter"
	h "github.com/madeiramadeirabr/lib-go-logger/pkg/handler"
	l "github.com/madeiramadeirabr/lib-go-logger/pkg/log_level"
)

type Interface interface {
	Emergency(message string, logMessageOptions f.LogMessageOptions)
	Error(message string, logMessageOptions f.LogMessageOptions)
	Warning(message string, logMessageOptions f.LogMessageOptions)
	Info(message string, logMessageOptions f.LogMessageOptions)
	Debug(message string, logMessageOptions f.LogMessageOptions)
	Trace(message string, logMessageOptions f.LogMessageOptions)
}

type Logger struct {
	formatter f.Interface
	handler   h.Interface
}

type Config struct {
	ServiceName string
	Level       l.LogLevelEnum
}

func New(
	handler h.Interface,
	formatter f.Interface,
) *Logger {

	return &Logger{
		formatter: formatter,
		handler:   handler,
	}
}

func (logger Logger) Emergency(message string, logMessageOptions f.LogMessageOptions) {
	logger.Log(l.LogLevelEmergency, message, logMessageOptions)
}
func (logger Logger) Error(message string, logMessageOptions f.LogMessageOptions) {
	logger.Log(l.LogLevelError, message, logMessageOptions)
}
func (logger Logger) Warning(message string, logMessageOptions f.LogMessageOptions) {
	logger.Log(l.LogLevelWarning, message, logMessageOptions)
}
func (logger Logger) Info(message string, logMessageOptions f.LogMessageOptions) {
	logger.Log(l.LogLevelInfo, message, logMessageOptions)
}
func (logger Logger) Debug(message string, logMessageOptions f.LogMessageOptions) {
	logger.Log(l.LogLevelDebug, message, logMessageOptions)
}
func (logger Logger) Trace(message string, logMessageOptions f.LogMessageOptions) {
	logger.Log(l.LogLevelTrace, message, logMessageOptions)
}

func (logger Logger) Log(level l.LogLevelEnum, message string, logMessageOptions f.LogMessageOptions) bool {
	if !logger.handler.IsHandling(level) {
		return false
	}

	formattedMessage, err := logger.formatter.Format(level, message, logMessageOptions)

	if err != nil {
		return false
	}

	return logger.handler.Write(formattedMessage)
}
