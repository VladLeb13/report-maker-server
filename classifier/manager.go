package classifier

import (
	fault_tolerant "report-maker-server/classifier/fault-tolerant"
	"report-maker-server/classifier/perfomance"
	"report-maker-server/tools"
)

func Manager(ctx *tools.AppContex) {
	go startPerformanceAnalyzer(ctx)
	go startFaultTolerantAnalyzer(ctx)

}

func startPerformanceAnalyzer(ctx *tools.AppContex) {
	perfomance.Worker(ctx)
}

func startFaultTolerantAnalyzer(ctx *tools.AppContex) {
	fault_tolerant.Worker(ctx)
}
