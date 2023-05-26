package log_level

type LogLevelEnum string

const (
	LogLevelEmergency LogLevelEnum = "EMERGENCY"
	LogLevelError     LogLevelEnum = "ERROR"
	LogLevelWarning   LogLevelEnum = "WARNING"
	LogLevelInfo      LogLevelEnum = "INFO"
	LogLevelDebug     LogLevelEnum = "DEBUG"
	LogLevelTrace     LogLevelEnum = "TRACE"
)

func LevelStringToInt(level LogLevelEnum) int {
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
