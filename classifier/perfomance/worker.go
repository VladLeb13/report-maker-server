package perfomance

import (
	"log"
	"time"

	"report-maker-server/config"
	"report-maker-server/tools"

	"github.com/VladLeb13/gophernet/run"
)

func Worker(ctx *tools.AppContex) {
	perfStruct := ctx.Context.Value("PerformanceAnalyzerStruct").(tools.PerformanceAnalyzerStruct)
	perfChan := ctx.Context.Value("PerfomanceChan").(tools.PerfomanceChan)
	cnf := ctx.Context.Value("config").(config.Config)

	in := perfStruct.Data
	result := perfStruct.Result
	status := make(chan bool)

	go run.Classifier(in, result, status)

	time.Sleep(time.Duration(cnf.Start_timeout) * time.Second)

	var init_manager int
	for init_manager != 1 {
		select {
		case ok := <-status:
			if ok {
				log.Println("Init success")
				init_manager = 1
				break
			}
		default:
			status <- false
			init_manager = 1
			return
		}
	}

	for {
		select {
		case d := <-perfChan.DataForPerformanceAnalyze:
			in <- d.Data
		case raw_data := <-result:
			res := tools.PerformanceAnalysisResult{}
			res.Set(raw_data)
			perfChan.AnalysisResult <- res
		default:
			time.Sleep(1 * time.Second)
		}
	}

}
