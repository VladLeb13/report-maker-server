package tests

import (
	"context"
	"log"
	"testing"

	"report-maker-server/classifier"
	"report-maker-server/config"
	"report-maker-server/database"
	"report-maker-server/server"
	"report-maker-server/tools"

	class "github.com/VladLeb13/gophernet/run"
)

func TestServe(t *testing.T) {
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

func TestStartWorker(t *testing.T) {

	ctx := tools.AppContex{
		Context: context.Background(),
	}
	perfStruct := tools.PerformanceAnalyzerStruct{
		Data:   make(chan class.Data),
		Result: make(chan []string),
	}
	faultTlStruct := tools.FaultTolerantAnalyzerStruct{
		Data:   make(chan tools.FaultTolerantParameters),
		Result: make(chan float64),
	}
	//Используются внутри воркеров
	ctx.Context = context.WithValue(ctx.Context, "PerformanceAnalyzerStruct", perfStruct)
	ctx.Context = context.WithValue(ctx.Context, "FaultTolerantAnalyzerStruct", faultTlStruct)

	perfChan := tools.PerfomanceChan{
		DataForPerformanceAnalyze: make(chan tools.DataForPerformanceAnalyze),
		AnalysisResult:            make(chan tools.PerformanceAnalysisResult),
	}
	ctx.Context = context.WithValue(ctx.Context, "PerfomanceChan", perfChan)

	faultTlChan := tools.FaultTolerantChan{
		DataFaultTolerantAnalyze: make(chan tools.FaultTolerantParameters),
		AnalysisResult:           make(chan float64),
	}
	ctx.Context = context.WithValue(ctx.Context, "FaultTolerantChan", faultTlChan)

	classifier.Manager(&ctx)

}
