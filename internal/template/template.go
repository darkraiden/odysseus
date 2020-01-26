package template

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Template struct {
	Name   string
	logger *logrus.Logger
}

// New is a factory function and accepts a logging function and some data
func New(logger *logrus.Logger, data string) *Template {
	return &Template{data, logger}
}

// Log should do something and log using the given logger
func (t *Template) Log(m string) {
	t.logger.Info(fmt.Sprintf("%s", m))
}
