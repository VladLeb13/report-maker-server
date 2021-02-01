package scheduler

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"report-maker-server/config"
	"report-maker-server/tools"
)

const (
	get_Data_For_Analysis = ` 	SELECT    HardwareID as hard_id
 										, PerfomanceID as perf_id
 										, Fault_toleranceID as flt_id 				
								FROM Workstation
								WHERE Allow_analysis = 1`

	get_Data_For_Perfomance = ` 	SELECT  RAM.Size AS ram_size
								, RAM.Frequency AS ram_freq
								, CPU.Frequency AS cpu_freq
								, CPU.Number_cores AS cpu_cores
								, CPU.Number_threads AS cpu_thread
								, HDD.Type AS disk_type
								FROM Hardware
									INNER JOIN Hardware ON CPU_list.ID = Hardware.CPU_listID
									INNER JOIN CPU ON CPU.ID = CPU_list.CPUID
									INNER JOIN Hardware ON Hardware.ID = Hardware.RAM_listID
									INNER JOIN CPU ON RAM.ID = RAM_list.RAMID
									INNER JOIN Hardware ON Hardware.ID = Hardware.HDD_listID
									INNER JOIN CPU ON HDD.ID = HDD_list.HDDID
								WHERE Hardware.ID = $1`

	get_Data_For_FLT = ` 	SELECT Fault_tolerance.Commissioning_date AS commission_date 
 								, Fault_tolerance.Backup AS backup
								, Fault_tolerance.Number_of_error AS error_count 
								, Perfomance.Cluster AS perf_cluster
						FROM Fault_tolerance
							INNER JOIN Perfomance ON Perfomance.ID = $1
						WHERE Fault_tolerance.ID = $2`
)

type (
	data_For_Analysis struct {
		hard_id string
		perf_id string
		flt_id  string
	}

	data_For_Perfomance_Analysis struct {
		ram_sie    int
		ram_freq   int
		cpu_freq   int
		cpu_cores  int
		cpu_thread int
		disk_type  int
		id         string
	}

	data_For_FLT_Analysis struct {
		commission_date string
		backup          int
		error_count     int
		perf_cluster    int
		id              string
	}
)

func Start(ctx *tools.AppContex) {
	cnf := ctx.Context.Value("config").(config.Config)
	db := ctx.Context.Value("database").(*sql.DB)

	timeout := cnf.Scheduler_cycle
	for { //общий цикл для воркера
		time.Sleep(time.Duration(timeout) * time.Minute)

		rows, err := db.Query(get_Data_For_Analysis)
		if err != nil {
			log.Println("Error in query get_Data_For_Analysis ", err)
		}

		var resp []data_For_Analysis
		for rows.Next() {
			var v data_For_Analysis
			rows.Scan(&v)

			resp = append(resp, v)
		}
		rows.Close()

		for _, v := range resp {

			var perfomanceParam data_For_Perfomance_Analysis
			err = db.QueryRow(get_Data_For_Perfomance, v.hard_id).Scan(&perfomanceParam)
			if err != nil {
				log.Println("Error in scan values get_DataForPerfomance ", err)
			}
			perfomanceParam.id = v.perf_id
			if err := PerfomanceAnalysis(ctx, perfomanceParam); err != nil {
				log.Println(err)
			}

			var fltParam data_For_FLT_Analysis
			err = db.QueryRow(get_Data_For_FLT, v.perf_id, v.flt_id).Scan(&fltParam)
			if err != nil {
				log.Println("Error in scan values get_DataForFLT ", err)
			}
			fltParam.id = v.flt_id
			if err := FaultTolerantAnalysis(ctx, fltParam); err != nil {
				log.Println(err)
			}

		}

	}

}

const set_Perfomance_Analisis_Result = `UPDATE Perfomance
									    SET Cluster = $1
									    WHERE ID = $2`

func PerfomanceAnalysis(ctx *tools.AppContex, d data_For_Perfomance_Analysis) (err error) {
	db := ctx.Context.Value("database").(*sql.DB)
	perfChan := ctx.Context.Value("PerfomanceChan").(tools.PerfomanceChan)

	perf_params := tools.PerformanceParameters{}
	params := perf_params.Set(d.ram_sie, d.ram_freq, d.cpu_freq, d.cpu_cores, d.cpu_thread, d.disk_type)

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

	_, err = db.Exec(set_Perfomance_Analisis_Result, cluster, d.id)
	if err != nil {
		log.Println("Error exec query \"set_Fault_Tolerant_Analisis_Result\"", err)
	}

	return
}

const set_Fault_Tolerant_Analisis_Result = ` UPDATE Fault_tolerance
								             SET Cluster = $1
								             WHERE ID = $2`

func FaultTolerantAnalysis(ctx *tools.AppContex, d data_For_FLT_Analysis) (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	db := ctx.Context.Value("database").(*sql.DB)
	faultTl := ctx.Context.Value("FaultTolerantAnalyzerStruct").(tools.FaultTolerantAnalyzerStruct)

	data := tools.FaultTolerantParameters{
		CommissioningDate: time.Time{},
		Backup:            false,
		ErrorCount:        0,
		PerfomanceCluster: 0,
	}
	data.Set(d.commission_date, d.backup, d.error_count, d.perf_cluster)

	faultTl.Data <- data
	result, ok := <-faultTl.Result
	if !ok {
		err = errors.New("Error change chanel ")
		return
	}

	_, err = db.Exec(set_Fault_Tolerant_Analisis_Result, int(result), d.id)
	if err != nil {
		log.Println("Error exec query \"set_Fault_Tolerant_Analisis_Result\"")
		return
	}
	return
}
