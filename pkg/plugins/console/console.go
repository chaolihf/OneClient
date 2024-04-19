package console

import (
	"go.uber.org/zap"
)

type Console struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Console {
	return &Console{logger: logger}
}

func (console *Console) Log(msg ...interface{}) {
	console.logger.Info(msg[0].(string))
}
