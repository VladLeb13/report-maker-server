package tests

import (
	"context"
	"log"
	"testing"

	"report-maker-server/config"
	"report-maker-server/database"
	"report-maker-server/server"
	"report-maker-server/tools"
)

func TestApp(t *testing.T) {
	ctx := tools.AppContex{
		Context: context.Background(),
	}

	cnf, err := config.Parse()
	if err != nil {
		log.Fatal("Error parsing config file")
	}
	ctx.Context = context.WithValue(ctx.Context, "config", cnf)

	db := database.Get(&ctx)
	ctx.Context = context.WithValue(ctx.Context, "database", db)

	ctx.Context = context.WithValue(ctx.Context, "test-test", "Это тестовое сообщение переданное из контекста приложения")

	err = server.Serve(&ctx)
	if err != nil {
		log.Println(err)
	}
}
