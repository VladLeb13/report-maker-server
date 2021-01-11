package scheduler

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"report-maker-server/config"
	"report-maker-server/tools"
)

const get_DataForAnalysis = `
`

func Start(ctx *tools.AppContex) {
	cnf := ctx.Context.Value("config").(config.Config)
	db := ctx.Context.Value("database").(*sql.DB)

	//timeout:=cnf.Scheduler_cycle
	//for { //общий цикл для воркера
	//time.Sleep(time.Duration(timeout) * time.Minute)
	//TODO: Запрос на выбор всех uuid из таблицы Workstation
	//for rows in uuid из Workstation {
	//time.Sleep(time.Duration(timeout) * time.Minute)
	var (
		ram_size   int
		ram_freq   int
		cpu_freq   int
		cpu_cores  int
		cpu_thread int
		disk_type  int
	)

	//TODO: Запрос на выбор параметров по uuid для  PerfomanceAnalysis
	//select
	if err := PerfomanceAnalysis(ctx, ram_size, ram_freq, cpu_freq, cpu_cores, cpu_thread, disk_type); err != nil {
		log.Println(err)
	}

	var (
		commission_date string
		backup          int
		error_count     int
		perf_cluster    int
	)

	//TODO: Запрос на выбор параметров по uuid для  FaultTolerantAnalysis
	//select

	if err := FaultTolerantAnalysis(ctx, commission_date, backup, error_count, perf_cluster); err != nil {
		log.Println(err)
	}

	//}
	//}

}

const set_Perfomance_Analisis_Result = `
`

func PerfomanceAnalysis(ctx *tools.AppContex, ram_sie int, ram_freq int, cpu_freq int, cpu_cores int, cpu_thread int, disk_type int) (err error) {
	db := ctx.Context.Value("database").(*sql.DB)
	perfChan := ctx.Context.Value("PerfomanceChan").(tools.PerfomanceChan)

	params := tools.PerformanceParameters{}.Set(ram_sie, ram_freq, cpu_freq, cpu_cores, cpu_thread, disk_type)

	data := tools.DataForPerformanceAnalyze{}
	data.Set(params)

	perfChan.DataForPerformanceAnalyze <- data

	result := <-perfChan.AnalysisResult

	var cluster int
	for _, v := range params {
		if v.Cluster.Low == result.Data.Cluster.Low && v.Cluster.Normal == result.Data.Cluster.Normal && v.Cluster.High == result.Data.Cluster.High {
			if result.Data.Cluster.Low == 1.0 {
				cluster = 3
			}
			if result.Data.Cluster.Normal == 1.0 {
				cluster = 4
			}
			if result.Data.Cluster.High == 1.0 {
				cluster = 5
			}
		}
	}

	//TODO: Запрос на запись значения анализа производительности
	_, err = db.Exec(set_Perfomance_Analisis_Result, cluster)
	if err != nil {
		log.Println("Error exec query \"set_Fault_Tolerant_Analisis_Result\"", err)
	}

	return
}

const set_Fault_Tolerant_Analisis_Result = `
`

func FaultTolerantAnalysis(ctx *tools.AppContex, commission_date string, backup int, error_count int, perf_cluster int) (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	db := ctx.Context.Value("database").(*sql.DB)
	faultTl := ctx.Context.Value("FaultTolerantAnalyzerStruct").(tools.FaultTolerantAnalyzerStruct)

	data := tools.FaultTolerantParameters{
		CommissioningDate: time.Time{},
		Backup:            false,
		ErrorCount:        0,
		PerfomanceCluster: 0,
	}
	data.Set(commission_date, backup, error_count, perf_cluster)

	faultTl.Data <- data
	result, ok := <-faultTl.Result
	if !ok {
		err = errors.New("Error change chanel ")
		return
	}

	//TODO: Запрос на запись значения анализа отказоустойчивости
	_, err = db.Exec(set_Fault_Tolerant_Analisis_Result, int(result))
	if err != nil {
		log.Println("Error exec query \"set_Fault_Tolerant_Analisis_Result\"")
		return
	}
	return
}
