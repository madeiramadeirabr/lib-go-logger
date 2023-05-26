package container

import (
	f "lib-go-logger/pkg/formatter"
	h "lib-go-logger/pkg/handler"
	l "lib-go-logger/pkg/logger"
)

func makeLogger(config l.Config) *l.Logger {

	handler := h.New(config.Level)
	formatter := f.New(config.ServiceName)

	return l.New(handler, formatter)
}
