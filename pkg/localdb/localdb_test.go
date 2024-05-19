package localdb

import (
	"testing"

	"github.com/kardianos/service"
)

func TestDb(t *testing.T) {
	logger := service.ConsoleLogger
	InitDb(logger)
	defer CloseDb()
	PingNetwork()
	GetMonitorData()
}
