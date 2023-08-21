package factory

import (
	"github.com/madeiramadeirabr/lib-go-logger/clock"
	f "github.com/madeiramadeirabr/lib-go-logger/formatter"
	h "github.com/madeiramadeirabr/lib-go-logger/handler"
	l "github.com/madeiramadeirabr/lib-go-logger/logger"
)

func MakeLogger(config l.Config) *l.Logger {

	handler := h.New(config.Level)
	formatter := f.New(config.ServiceName, clock.Clock{})

	return l.New(handler, formatter)
}
