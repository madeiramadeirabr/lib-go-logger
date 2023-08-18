package container

import (
	"github.com/madeiramadeirabr/lib-go-logger/pkg/clock"
	f "github.com/madeiramadeirabr/lib-go-logger/pkg/formatter"
	h "github.com/madeiramadeirabr/lib-go-logger/pkg/handler"
	l "github.com/madeiramadeirabr/lib-go-logger/pkg/logger"
)

func MakeLogger(config l.Config) *l.Logger {

	handler := h.New(config.Level)
	formatter := f.New(config.ServiceName, clock.Clock{})

	return l.New(handler, formatter)
}
