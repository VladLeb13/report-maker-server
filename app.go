package main

import (
	"context"
	"log"

	"report-maker-server/classifier"
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

	perfStruct := tools.PerformanceAnalyzerStruct{
		Data:   make(chan class.Data),
		Result: make(chan []string),
	}
	ctx.Context = context.WithValue(ctx.Context, "PerformanceAnalyzerStruct", perfStruct)

	perfChan := tools.PerfomanceChan{
		DataForPerformanceAnalyze: make(chan tools.DataForPerformanceAnalyze),
		AnalysisResult:            make(chan tools.PerformanceAnalysisResult),
	}
	ctx.Context = context.WithValue(ctx.Context, "PerfomanceChan", perfChan)

	go classifier.Manager(&ctx)

	err := server.Serve()
	if err != nil {
		log.Println(err)
	}

}

//answer := strings.Join(res, " ")
//for i, v := range d.RawData {
//elem := strings.Join(v, " ")
//if answer == elem {
//log.Println("bingo!!! утверждение номер: " + strconv.Itoa(i) + " верно")
//log.Println(v)
//}
//}
