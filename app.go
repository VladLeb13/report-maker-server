package main

import (
	"context"
	"log"

	"report-maker-server/classifier"
	"report-maker-server/config"
	"report-maker-server/database"
	"report-maker-server/scheduler"
	"report-maker-server/server"
	"report-maker-server/tools"

	class "github.com/VladLeb13/classifier/run"
)

func main() {
	//TODO: make serve
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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

	go classifier.Manager(&ctx)

	go scheduler.Start(&ctx)

	err = server.Serve(&ctx)
	if err != nil {
		log.Println(err)
	}

}
