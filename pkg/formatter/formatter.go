package formatter

import (
	"encoding/json"
	"fmt"
	"lib-go-logger/pkg/clock"
	l "lib-go-logger/pkg/log_level"
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
	clock       clock.Interface
}

func New(serviceName string, clock clock.Interface) *Formatter {

	return &Formatter{
		serviceName: serviceName,
		clock:       clock,
	}
}

func (f Formatter) Format(level l.LogLevelEnum, message string, logMessageOptions LogMessageOptions) (string, error) {
	log := LogMessage{
		GlobalEventTimestamp: f.clock.GetCurrentTimestamp().String(),
		Level:                level,
		Message:              message,
		ServiceName:          f.serviceName,
		SessionId:            logMessageOptions.SessionId,
		TraceId:              logMessageOptions.TraceId,
		GlobalEventName:      logMessageOptions.GlobalEventName,
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