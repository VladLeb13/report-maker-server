package tools

import "time"

type (
	//Используется внутри воркеров
	FaultTolerantAnalyzerStruct struct {
		Data   chan FaultTolerantParameters
		Result chan float64
	}
	FaultTolerantParameters struct {
		CommissioningDate time.Time
		Backup            bool
		ErrorCount        int
		PerfomanceCluster float64
	}

	//Используется внутри приложения
	FaultTolerantChan struct {
		DataFaultTolerantAnalyze chan FaultTolerantParameters
		AnalysisResult           chan float64
	}
)
