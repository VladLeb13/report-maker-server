package tests

import (
	"report-maker-server/server"
	"testing"
)

func TestApp(t *testing.T) {
	server.Serve()
}
