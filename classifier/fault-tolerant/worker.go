package fault_tolerant

import (
	"time"

	"report-maker-server/tools"
)

func Worker(ctx *tools.AppContex) {
	faultTl := ctx.Context.Value("FaultTolerantAnalyzerStruct").(tools.FaultTolerantAnalyzerStruct)
	faultTlChan := ctx.Context.Value("FaultTolerantChan").(tools.FaultTolerantChan)

	in := faultTl.Data
	result := faultTl.Result

	go analyzer(in, result)

	for {
		select {
		case d := <-faultTlChan.DataFaultTolerantAnalyze:
			in <- d
		case res := <-result:
			faultTlChan.AnalysisResult <- res
		default:
			time.Sleep(1 * time.Second)
		}
	}

}

func analyzer(in chan tools.FaultTolerantParameters, result chan float64) {

	for {
		select {
		case d := <-in:
			result <- analysis(d)
		default:
			time.Sleep(1 * time.Second)
		}

	}

}

const (
	backupIdx            = 3
	errorIdx             = 1
	commissioningDateIdx = 1
)

func analysis(data tools.FaultTolerantParameters) (result float64) {
	startIdx := data.PerfomanceCluster

	if !data.Backup {
		startIdx = -backupIdx
	}

	if data.ErrorCount >= 15 {
		startIdx = -errorIdx
	}

	t := time.Now()
	t = t.AddDate(-5, 0, 0)
	if data.CommissioningDate.Before(t) {
		startIdx = -commissioningDateIdx
	}

	if startIdx <= 0 {
		startIdx = 1
	}

	result = startIdx

	return
}
