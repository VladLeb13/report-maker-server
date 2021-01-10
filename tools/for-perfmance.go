package tools

import (
	class "github.com/VladLeb13/classifier/run"
)

type (
	//Используется внутри воркеров
	PerformanceAnalyzerStruct struct {
		Data   chan class.Data
		Result chan []string
	}

	PerformanceParameters struct {
		RAMVolume float64
		RAMSpeed  float64
		CPUFreq   float64
		CPUCore   float64
		CPUThread float64
		DiskType  float64
		Cluster   PerformanceParametersCluster
	}

	PerformanceParametersCluster struct {
		High   float64
		Normal float64
		Low    float64
	}
)

//Используется внутри приложения
type PerfomanceChan struct {
	DataForPerformanceAnalyze chan DataForPerformanceAnalyze
	AnalysisResult            chan PerformanceAnalysisResult
}

type PerformanceAnalysisResult struct {
	Data PerformanceParameters
}

func (res *PerformanceAnalysisResult) Set(raw_data []string) {
	if len(raw_data) == 9 {
		res.Data.RAMVolume = stringToFloat(raw_data[0])
		res.Data.RAMSpeed = stringToFloat(raw_data[1])
		res.Data.CPUFreq = stringToFloat(raw_data[2])
		res.Data.CPUCore = stringToFloat(raw_data[3])
		res.Data.CPUThread = stringToFloat(raw_data[4])
		res.Data.DiskType = stringToFloat(raw_data[5])
		res.Data.Cluster.High = stringToFloat(raw_data[6])
		res.Data.Cluster.Normal = stringToFloat(raw_data[7])
		res.Data.Cluster.Low = stringToFloat(raw_data[8])
	}
}

type DataForPerformanceAnalyze struct {
	Data class.Data
}

func (d *DataForPerformanceAnalyze) Set(params []PerformanceParameters) {
	for _, v := range params {
		d.Data.RawData = append(d.Data.RawData, []string{
			floatToString(v.RAMVolume),
			floatToString(v.RAMSpeed),
			floatToString(v.CPUFreq),
			floatToString(v.CPUCore),
			floatToString(v.CPUThread),
			floatToString(v.DiskType),
			floatToString(v.Cluster.High),
			floatToString(v.Cluster.Normal),
			floatToString(v.Cluster.Low),
		})
	}
}
