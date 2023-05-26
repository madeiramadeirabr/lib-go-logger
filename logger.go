package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

type LoggerInterface interface {
	Emergency(message string, logMessageOptions LogMessageOptions)
	Error(message string, logMessageOptions LogMessageOptions)
	Warning(message string, logMessageOptions LogMessageOptions)
	Info(message string, logMessageOptions LogMessageOptions)
	Debug(message string, logMessageOptions LogMessageOptions)
	Trace(message string, logMessageOptions LogMessageOptions)
}

type LogLevelEnum string

const (
	LogLevelEmergency LogLevelEnum = "EMERGENCY"
	LogLevelError     LogLevelEnum = "ERROR"
	LogLevelWarning   LogLevelEnum = "WARNING"
	LogLevelInfo      LogLevelEnum = "INFO"
	LogLevelDebug     LogLevelEnum = "DEBUG"
	LogLevelTrace     LogLevelEnum = "TRACE"
)

type Logger struct {
	serviceName string
	level       LogLevelEnum
}

type LogMessage struct {
	GlobalEventTimestamp string       `json:"global_event_timestamp"`
	GlobalEventName      string       `json:"global_event_name,omitempty"`
	Level                LogLevelEnum `json:"level"`
	Context              string       `json:"context,omitempty"`
	Message              string       `json:"message"`
	ServiceName          string       `json:"service_name"`
	TraceId              string       `json:"trace_id,omitempty"`
	SessionId            string       `json:"session_id,omitempty"`
}

type LogMessageOptions struct {
	GlobalEventName string      `json:"global_event_name,omitempty"`
	Context         interface{} `json:"context,omitempty"`
	TraceId         string      `json:"trace_id,omitempty"`
	SessionId       string      `json:"session_id,omitempty"`
}

func New(
	serviceName string,
	level LogLevelEnum,
) *Logger {

	return &Logger{
		serviceName: serviceName,
		level:       level,
	}
}

func (logger Logger) levelStringToInt(level LogLevelEnum) int {
	switch level {
	case LogLevelTrace:
		return 0
	case LogLevelDebug:
		return 1
	case LogLevelInfo:
		return 2
	case LogLevelWarning:
		return 3
	case LogLevelError:
		return 4
	case LogLevelEmergency:
		return 5
	default:
		return 0
	}
}

func (logger Logger) isHandling(levelMessage LogLevelEnum) bool {
	return logger.levelStringToInt(levelMessage) >= logger.levelStringToInt(logger.level)
}

func (logger Logger) Log(level LogLevelEnum, message string, logMessageOptions LogMessageOptions) bool {
	if !logger.isHandling(level) {
		return false
	}

	formattedMessage, err := logger.formatMessage(level, message, logMessageOptions)

	if err != nil {
		return false
	}

	fmt.Println(formattedMessage)

	return true
}

func (logger Logger) formatMessage(level LogLevelEnum, message string, logMessageOptions LogMessageOptions) (string, error) {
	log := LogMessage{
		GlobalEventName:      logMessageOptions.GlobalEventName,
		TraceId:              logMessageOptions.TraceId,
		SessionId:            logMessageOptions.SessionId,
		ServiceName:          logger.serviceName,
		Message:              message,
		Level:                level,
		GlobalEventTimestamp: time.Now().Format(time.RFC3339),
	}

	if logMessageOptions.Context != nil {
		if context := fmt.Sprintf("%s", logMessageOptions.Context); context != "" {
			log.Context = context
		}
	}

	logMarshal, err := json.Marshal(log)

	if err != nil {
		return "", err
	}

	return string(logMarshal), nil
}

func (logger Logger) Emergency(message string, logMessageOptions LogMessageOptions) {
	logger.Log(LogLevelEmergency, message, logMessageOptions)
}
func (logger Logger) Error(message string, logMessageOptions LogMessageOptions) {
	logger.Log(LogLevelError, message, logMessageOptions)
}
func (logger Logger) Warning(message string, logMessageOptions LogMessageOptions) {
	logger.Log(LogLevelWarning, message, logMessageOptions)
}
func (logger Logger) Info(message string, logMessageOptions LogMessageOptions) {
	logger.Log(LogLevelInfo, message, logMessageOptions)
}
func (logger Logger) Debug(message string, logMessageOptions LogMessageOptions) {
	logger.Log(LogLevelDebug, message, logMessageOptions)
}
func (logger Logger) Trace(message string, logMessageOptions LogMessageOptions) {
	logger.Log(LogLevelTrace, message, logMessageOptions)
}
