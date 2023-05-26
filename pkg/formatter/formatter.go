package formatter

import (
	"encoding/json"
	"fmt"
	l "lib-go-logger/pkg/log_level"
	"time"
)

type LogMessage struct {
	GlobalEventTimestamp string         `json:"global_event_timestamp"`
	GlobalEventName      string         `json:"global_event_name,omitempty"`
	Level                l.LogLevelEnum `json:"level"`
	Context              string         `json:"context,omitempty"`
	Message              string         `json:"message"`
	ServiceName          string         `json:"service_name"`
	TraceId              string         `json:"trace_id,omitempty"`
	SessionId            string         `json:"session_id,omitempty"`
}

type LogMessageOptions struct {
	GlobalEventName string      `json:"global_event_name,omitempty"`
	Context         interface{} `json:"context,omitempty"`
	TraceId         string      `json:"trace_id,omitempty"`
	SessionId       string      `json:"session_id,omitempty"`
}

type Interface interface {
	Format(level l.LogLevelEnum, message string, logMessageOptions LogMessageOptions) (string, error)
}

type Formatter struct {
	serviceName string
}

func New(serviceName string) *Formatter {

	return &Formatter{
		serviceName: serviceName,
	}
}

func (f Formatter) Format(level l.LogLevelEnum, message string, logMessageOptions LogMessageOptions) (string, error) {
	log := LogMessage{
		GlobalEventName:      logMessageOptions.GlobalEventName,
		TraceId:              logMessageOptions.TraceId,
		SessionId:            logMessageOptions.SessionId,
		ServiceName:          f.serviceName,
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
