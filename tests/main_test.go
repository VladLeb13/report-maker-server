package tests

import (
	"testing"

	"report-maker-server/server"
)

func TestApp(t *testing.T) {
	server.Serve()
}
