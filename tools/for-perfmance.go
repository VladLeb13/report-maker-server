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

type PerformanceParameters struct {
	RAMVolume float64
	RAMSpeed  float64
	CPUFreq   float64
	CPUCore   float64
	CPUThread float64
	DiskType  float64
	Cluster   PerformanceParametersCluster
}

func (pp *PerformanceParameters) Set(ram_sie int, ram_freq int,
	cpu_freq int, cpu_cores int, cpu_thread int, disk_type int) (params []PerformanceParameters) {

	pp.RAMVolume = float64(ram_sie)
	pp.RAMSpeed = float64(ram_freq)
	pp.CPUFreq = float64(cpu_freq)
	pp.CPUCore = float64(cpu_cores)
	pp.CPUThread = float64(cpu_thread)
	pp.DiskType = float64(disk_type)

	pp.normalize()

	param1 := PerformanceParameters{
		RAMVolume: pp.RAMVolume,
		RAMSpeed:  pp.RAMSpeed,
		CPUFreq:   pp.CPUFreq,
		CPUCore:   pp.CPUCore,
		CPUThread: pp.CPUThread,
		DiskType:  pp.DiskType,
		Cluster: PerformanceParametersCluster{
			High:   1.0,
			Normal: 0.0,
			Low:    0.0,
		},
	}

	params = append(params, param1)

	param2 := PerformanceParameters{
		RAMVolume: pp.RAMVolume,
		RAMSpeed:  pp.RAMSpeed,
		CPUFreq:   pp.CPUFreq,
		CPUCore:   pp.CPUCore,
		CPUThread: pp.CPUThread,
		DiskType:  pp.DiskType,
		Cluster: PerformanceParametersCluster{
			High:   0.0,
			Normal: 1.0,
			Low:    0.0,
		},
	}

	params = append(params, param2)

	param3 := PerformanceParameters{
		RAMVolume: pp.RAMVolume,
		RAMSpeed:  pp.RAMSpeed,
		CPUFreq:   pp.CPUFreq,
		CPUCore:   pp.CPUCore,
		CPUThread: pp.CPUThread,
		DiskType:  pp.DiskType,
		Cluster: PerformanceParametersCluster{
			High:   0.0,
			Normal: 0.0,
			Low:    1.0,
		},
	}

	params = append(params, param3)
	return
}
func (pp *PerformanceParameters) normalize() {
	ram_volume := pp.RAMVolume * 1024
	ram_volume = ram_volume / 100000
	pp.RAMVolume = ram_volume

	ram_speed := pp.RAMSpeed / 10000
	pp.RAMSpeed = ram_speed

	cpu_freq := pp.CPUFreq / 10000
	pp.CPUFreq = cpu_freq

	cpu_core := pp.CPUCore / 10
	pp.CPUCore = cpu_core

	cpu_thread := pp.CPUThread / 10
	pp.CPUThread = cpu_thread

	disk_type := pp.DiskType / 10
	pp.DiskType = disk_type
}
