package logger

import (
	"encoding/json"
	"fmt"
)

type LoggerInterface interface {
	Emergency(message string, args interface{}, globalEventName string)
	Error(message string, args interface{}, globalEventName string)
	Warning(message string, args interface{}, globalEventName string)
	Info(message string, args interface{}, globalEventName string)
	Debug(message string, args interface{}, globalEventName string)
	Trace(message string, args interface{}, globalEventName string)
}

type LogLevelEnum int

const (
	LogLevelEmergency LogLevelEnum = 5
	LogLevelError     LogLevelEnum = 4
	LogLevelWarning   LogLevelEnum = 3
	LogLevelInfo      LogLevelEnum = 2
	LogLevelDebug     LogLevelEnum = 1
	LogLevelTrace     LogLevelEnum = 0
)

type Logger struct {
	serviceName string
	logLevel    LogLevelEnum
}

type LogMessage struct {
	GlobalEventTimestamp string `json:"global_event_timestamp"`
	GlobalEventName      string `json:"global_event_name,omitempty"`
	Level                string `json:"level"`
	Context              string `json:"context,omitempty"`
	Message              string `json:"message"`
	ServiceName          string `json:"service_name"`
	TraceId              string `json:"trace_id"`
	SessionId            string `json:"session_id"`
}

func New(
	serviceName string,
	logLevel LogLevelEnum,
) *Logger {
	return &Logger{
		serviceName: serviceName,
		logLevel:    logLevel,
	}
}

func (logger Logger) Log(level LogLevelEnum, message string, args interface{}, globalEventName string) bool {
	if level > logger.logLevel {
		return false
	}

	formattedMessage, err := logger.formatMessage(message, globalEventName)

	if err == nil {
		return false
	}

	fmt.Println(formattedMessage)

	return true
}

func (logger Logger) formatMessage(message string, globalEventName string) (string, error) {

	log := LogMessage{
		GlobalEventName:      globalEventName,
		GlobalEventTimestamp: "",
		Message:              message,
		TraceId:              "",
		SessionId:            "",
		Context:              "",
		Level:                "",
	}

	logMarshal, err := json.Marshal(log)

	if err != nil {
		return "", err
	}

	return string(logMarshal), nil
}

func (logger Logger) Emergency(message string, args interface{}, globalEventName string) {
	logger.Log(LogLevelEmergency, message, args, globalEventName)
}
func (logger Logger) Error(message string, args interface{}, globalEventName string) {
	logger.Log(LogLevelError, message, args, globalEventName)
}
func (logger Logger) Warning(message string, args interface{}, globalEventName string) {
	logger.Log(LogLevelWarning, message, args, globalEventName)
}
func (logger Logger) Info(message string, args interface{}, globalEventName string) {
	logger.Log(LogLevelInfo, message, args, globalEventName)
}
func (logger Logger) Debug(message string, args interface{}, globalEventName string) {
	logger.Log(LogLevelDebug, message, args, globalEventName)
}
func (logger Logger) Trace(message string, args interface{}, globalEventName string) {
	logger.Log(LogLevelTrace, message, args, globalEventName)
}
