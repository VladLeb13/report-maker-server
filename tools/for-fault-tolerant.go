package tools

import "time"

type (
	//Используется внутри воркеров
	FaultTolerantAnalyzerStruct struct {
		Data   chan FaultTolerantParameters
		Result chan float64
	}

	//Используется внутри приложения
	FaultTolerantChan struct {
		DataFaultTolerantAnalyze chan FaultTolerantParameters
		AnalysisResult           chan float64
	}
)

type FaultTolerantParameters struct {
	CommissioningDate time.Time
	Backup            bool
	ErrorCount        int
	PerfomanceCluster float64
}

func (ftp *FaultTolerantParameters) Set(commission_date string, backup int, error_count int, perf_cluster int) {
	t, err := time.Parse(time.RFC3339, commission_date)
	if err != nil {
		ftp.CommissioningDate = *new(time.Time)
	} else {
		ftp.CommissioningDate = t
	}

	if backup == 1 {
		ftp.Backup = true
	} else {
		ftp.Backup = false
	}

	ftp.ErrorCount = error_count

	ftp.PerfomanceCluster = float64(perf_cluster)
}
