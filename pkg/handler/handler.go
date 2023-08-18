package handler

import (
	"fmt"
	l "github.com/madeiramadeirabr/lib-go-logger/pkg/log_level"
)

type Interface interface {
	IsHandling(levelMessage l.LogLevelEnum) bool
	Write(message string) bool
}

type Stdout struct {
	level l.LogLevelEnum
}

func New(level l.LogLevelEnum) *Stdout {

	return &Stdout{
		level: level,
	}
}

func (h Stdout) IsHandling(levelMessage l.LogLevelEnum) bool {
	return l.LevelStringToInt(levelMessage) >= l.LevelStringToInt(h.level)
}

func (h Stdout) Write(message string) bool {

	fmt.Println(message)

	return true
}
