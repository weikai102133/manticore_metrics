package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	//"github.com/go-delve/delve/cmd/dlv"
)

var (

	loger *log.Logger

	db = &sql.DB{}

	/*uptime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_uptime",
		Help: "manticore运行时间",
	})

	threadsNum = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_current_threads_num",
		Help: "manticore当前线程数量",
	})

	queueTaskNum = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_jobs_in_queue_num",
		Help: "manticore当前队列中任务数",
	})

	connectionNum = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_connections_num",
		Help: "manticore当前客户端连接数",
	})

	taskNum = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_current_tasks_num",
		Help: "manticore当前处理的任务数",
	})

	queryNum = promauto.NewCounter(prometheus.CounterOpts{
		Name: "manticore_queries_num",
		Help: "manticore启动后请求总数",
	})

	SSL = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_ssl_available",
		Help: "SSL证书是否启用 0为false，1为true",
	})

	socketConnectType = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_ssl_socket",
		Help: "host TCP socket连接方式 0为port，1为unix",
	})

	jobsPerQueue = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_jobs_num_per_queue",
		Help: "队列中任务数和线程数的比值 ",
	})


	taskPerThread = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_tasks_num_per_thread",
		Help: "处理的任务数和线程数的比值 ",
	})*/

	taskCompletedNum = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_metric_thread_num",
		Help: "显示这个线程已经完成的任务总数",
	},
		[]string{"threadName"})

	threadIsIdle = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_thread_in_idle",
		Help: "显示当前线程是否空闲 0为空闲 1为working",
	},
		[]string{"threadName"})

	idleThreadNum = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_thread_in_idle_num",
		Help: "显示空闲线程总数 ",
	})

	workingThreadNum = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_thread_in_working_num",
		Help: "显示工作线程总数 ",
	})

	connectProto = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_thread_proto",
		Help: "显示连接协议 0（sphinx）,1（mysql）,2（http）,3（ssl）,4（compressed ）,5（replication ）,6（combination ）,7线程空闲或其它协议",
	},
		[]string{"threadName"})

	threadStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_thread_state",
		Help: "显示线程状态 0（handshake），1（net_read），2（net_write），3（query），4（net_idle）,5(线程空闲)",
	},
		[]string{"threadName"})

	threadDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_time_seconds",
		Help: "显示此线程当前任务执行时间 单位s",

	},
		[]string{"threadName"})

	/*threadUptime = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_work_time_seconds",
		Help: "显示此线程开始时间",
	},
		[]string{"threadName"})

	effectiveCpuTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "manticore_work_time_CPU_seconds",
		Help: "显示有效cpu时间",
	})*/

	indexType = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_index_type",
		Help: "显示此索引类型 0（disk）, 1（rt）, 2（percolate）,3（template） ,  4（distributed），5（其它）",
	},
		[]string{"indexName"})

	docNum = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_documents_num",
		Help: "显示此索引文档的数量",
	},
		[]string{"indexName"})

	docBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_documents_bytes",
		Help: "显示此索引文档的大小",
	},
		[]string{"indexName"})

	ramBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_ram_bytes",
		Help: "显示常驻ram索引部分的总大小",
	},
		[]string{"indexName"})

	diskBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_disk_bytes",
		Help: "该索引的文件总大小",
	},
		[]string{"indexName"})

	mappedBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_disk_mapped_bytes",
		Help: "该索引文件映射的总大小",
	},
		[]string{"indexName"})

	mappedCacheBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_disk_mapped_cached_bytes",
		Help: "该索引实际缓存在RAM中的文件映射的总大小",
	},
		[]string{"indexName"})

	docListsBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_disk_mapped_doclists_bytes",
		Help: "显示文档列表映射部分大小",
	},
		[]string{"indexName"})

	docListsCacheBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_disk_mapped_cached_doclists_bytes",
		Help: "显示文档列表的缓存映射部分大小",
	},
		[]string{"indexName"})

	hitListsBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_disk_mapped_hitlists_bytes",
		Help: "显示命中列表的映射部分大小",
	},
		[]string{"indexName"})

	hitListsCacheBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_disk_mapped_cached_hitlists_bytes",
		Help: "显示命中列表的缓存映射部分大小",
	},
		[]string{"indexName"})

	killedDocNum = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_killed_documents_num",
		Help: "显示该索引删除文档总数",
	},
		[]string{"indexName"})

	killedDocRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_killed_rate",
		Help: "显示该索引删除文档所占比例   %",
	},
		[]string{"indexName"})

	ramChunkBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_ram_chunk_bytes",
		Help: "real-time or percolate index的RAM chunk大小",
	},
		[]string{"indexName"})

	ramChunkSegment = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_ram_chunk_segments_count",
		Help: "显示实时索引的磁盘块数",
	},
		[]string{"indexName"})

	diskChunkNum = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_disk_chunks_num",
		Help: "显示实时索引的磁盘块数",
	},
		[]string{"indexName"})

	menLimitBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_mem_limit_bytes",
		Help: "显示索引的rt_mem_limit的实际值",
	},
		[]string{"indexName"})

	garbageBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_ram_bytes_retired_bytes",
		Help: "显示该索引在RAM chunkst中垃圾的大小",
	},
		[]string{"indexName"})

	queryNumPerMin = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_query_time_1min_num",
		Help: "显示该索引最近一分钟内的查询次数统计",
	},
		[]string{"indexName"})

	queryAvgNumPerSecond = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_query_time_1min_avg",
		Help: "显示该索引最近一分钟内的查询执行次数均值 单位 : 次数/s",
	},
		[]string{"indexName"})

	/*queryMaxNumPerMin = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_query_time_1min_max",
		Help: "显示该索引最近一分钟内的查询执行次数统计最大值",
	},
		[]string{"indexName"})

	queryAvgNumPerMin = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_query_time_1min_avg",
		Help: "显示该索引最近一分钟内的查询执行时间统计平均值",
	},
		[]string{"indexName"})*/

	queryRowsNumPer = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_index_found_rows_1min_avg",
		Help: "显示该索引最近一分钟每次查询行数平均值",
	},
		[]string{"indexName"})

	/*queryRowsMinNumPerMin = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_index_found_rows_1min_min",
		Help: "显示该索引最近一分钟内的查询行数最小的那次查询行数值",
	},
		[]string{"indexName"})

	queryRowsMaxNumPerMin = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_index_found_rows_1min_nax",
		Help: "显示该索引最近一分钟内的查询行数最大的那次查询行数值",
	},
		[]string{"indexName"})

	queryRowsAvgNumPerMin = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "manticore_indexed_index_found_rows_1min_avg",
		Help: "显示该索引最近一分钟内的每次查询行数的平均值",
	},
		[]string{"indexName"})*/

)

type Config struct {
	ManticoreDriverName string `ini:"manticoreDriverName"`
	Username            string `ini:"username"`
	Password            string `ini:"password"`
	Ip                  string `ini:"ip"`
	Port                string `ini:"port"`
	DatabaseName        string `ini:"databaseName"`
}

func recordMetrics() {
	//查询并收集manticore当前线程状态
	go func() {
		for {
			time.Sleep(2 * time.Second)
			rows,err:=db.Query("show threads option format = all")       //获取所有数据
			if  err != nil{
				fmt.Println(err)
				return
			}

			threadIsIdle.Reset()
			taskCompletedNum.Reset()
			threadStatus.Reset()
			connectProto.Reset()
			threadDuration.Reset()

			var tid,name,proto,state,host,connid,time,worktime,worktimecpu,thd,jobs,lastjob,idle,info string

			idleCount := 0
			workingCount := 0
			for rows.Next(){        //循环显示所有的数据
				rows.Scan(&tid,&name,&proto,&state,&host,&connid,&time,&worktime,&worktimecpu,&thd,&jobs,&lastjob,&idle,&info)

				if idle == "No (working)"  {
					workingCount++
					threadIsIdle.With(prometheus.Labels{"threadName":name}).Set(1)  //线程工作中
					float,_ := strconv.ParseFloat(time,64)
					threadDuration.With(prometheus.Labels{"threadName":name}).Set(float)  //线程当前任务执行时间
				} else {
					idleCount++
					threadIsIdle.With(prometheus.Labels{"threadName":name}).Set(0)   //线程空闲中
				}

				float,_ := strconv.ParseFloat(jobs,64)
				taskCompletedNum.With(prometheus.Labels{"threadName":name}).Set(float)  //线程完成任务总数

				matchAndSetProto(proto,name)						//线程连接协议

				matchAndSetThreadStatus(state,name)                 //线程状态

			}

			idleThreadNum.Set(float64(idleCount))
			workingThreadNum.Set(float64(workingCount))
		}
	}()

	//查询并收集manticore当前索引状态
	go func() {
		for {
			time.Sleep(2 * time.Second)
			rows,err:=db.Query("show tables")       //获取所有索引
			//fmt.Println(time.Now())
			if  err != nil{
				fmt.Println(err)
				return
			}

			var index,value string

			indexType.Reset()
			docNum.Reset()
			docBytes.Reset()
			ramBytes.Reset()
			diskBytes.Reset()
			mappedBytes.Reset()
			mappedCacheBytes.Reset()
			queryNumPerMin.Reset()
			docListsBytes.Reset()
			docListsCacheBytes.Reset()
			hitListsBytes.Reset()
			hitListsCacheBytes.Reset()
			killedDocNum.Reset()
			killedDocRate.Reset()
			ramChunkBytes.Reset()
			ramChunkSegment.Reset()
			diskChunkNum.Reset()
			menLimitBytes.Reset()
			garbageBytes.Reset()
			queryNumPerMin.Reset()
			queryAvgNumPerSecond.Reset()
			queryRowsNumPer.Reset()


			for rows.Next(){        //循环显示所有的数据
				rows.Scan(&index,&value)
				showIndexStatus(index)
				//fmt.Println(index,"--",value)
			}

		}
	}()
}

//索引收集入口函数
func showIndexStatus(indexName string) {
	rows,err:=db.Query("show index "+indexName+" status")       //获取所有数据

	if  err != nil{
		loger.Println(err)
		return
	}

	var variableName ,value string

	for rows.Next(){        //循环显示所有的数据
		rows.Scan(&variableName,&value)
		matchAndSetIndexMetric(indexName,variableName,value)
	}

}

//收集多种索引指标
func matchAndSetIndexMetric(indexName string,variableName string,value string) {
	switch variableName {
	case "index_type":
		matchAndSetIndexType(value,indexName)
	case "indexed_documents":
		float,_ := strconv.ParseFloat(value,64)
		docNum.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "indexed_bytes":
		float,_ := strconv.ParseFloat(value,64)
		docBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "ram_bytes":
		float,_ := strconv.ParseFloat(value,64)
		ramBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_bytes":
		float,_ := strconv.ParseFloat(value,64)
		diskBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_mapped":
		float,_ := strconv.ParseFloat(value,64)
		mappedBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_mapped_cached":
		float,_ := strconv.ParseFloat(value,64)
		mappedCacheBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_mapped_doclists":
		float,_ := strconv.ParseFloat(value,64)
		docListsBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_mapped_cached_doclists":
		float,_ := strconv.ParseFloat(value,64)
		docListsCacheBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_mapped_hitlists":
		float,_ := strconv.ParseFloat(value,64)
		hitListsBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_mapped_cached_hitlists":
		float,_ := strconv.ParseFloat(value,64)
		hitListsCacheBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "killed_documents":
		float,_ := strconv.ParseFloat(value,64)
		killedDocNum.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "killed_rate":
		float,_ := strconv.ParseFloat(value[0:len(value)-1],64)
		fmt.Println(float)
		killedDocRate.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "ram_chunk":
		float,_ := strconv.ParseFloat(value,64)
		ramChunkBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "ram_chunk_segments_count":
		float,_ := strconv.ParseFloat(value,64)
		ramChunkSegment.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "disk_chunks":
		float,_ := strconv.ParseFloat(value,64)
		diskChunkNum.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "mem_limit":
		float,_ := strconv.ParseFloat(value,64)
		menLimitBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "ram_bytes_retired":
		float,_ := strconv.ParseFloat(value,64)
		garbageBytes.With(prometheus.Labels{"indexName":indexName}).Set(float)
	case "query_time_1min":
		val := splitValue(value)
		if  val!= ""{
			float,_ := strconv.ParseFloat(val,64)
			//fmt.Println(float/60)
			queryNumPerMin.With(prometheus.Labels{"indexName":indexName}).Set(float)
			queryAvgNumPerSecond.With(prometheus.Labels{"indexName":indexName}).Set(float/60)
		}
	case "found_rows_1min":
		val := splitRowValue(value)
		if  val!= ""{
			float,_ := strconv.ParseFloat(val,64)
			//fmt.Println(float)
			queryRowsNumPer.With(prometheus.Labels{"indexName":indexName}).Set(float)
		}
	default:
		//log.Println("matchAndSetIndexMetric没有"+variableName+"类型")
	}

}

//收集索引类别
func matchAndSetIndexType(types string,indexName string)  {
	switch types {
	case "disk":
		indexType.With(prometheus.Labels{"indexName":indexName}).Set(0)
	case "rt":
		indexType.With(prometheus.Labels{"indexName":indexName}).Set(1)
	case "percolate":
		indexType.With(prometheus.Labels{"indexName":indexName}).Set(2)
	case "template":
		indexType.With(prometheus.Labels{"indexName":indexName}).Set(3)
	case "distributed":
		indexType.With(prometheus.Labels{"indexName":indexName}).Set(4)
	default:
		indexType.With(prometheus.Labels{"indexName":indexName}).Set(5)
		//log.Println("matchAndSetIndexType没有"+types+"类型索引")
	}
}

/*func matchAndSetMetric(metricName string,value string)  {
	switch metricName {
	case "uptime":
		float,_ := strconv.ParseFloat(value,64)
		uptime.Set(float)
	}
}*/

//收集线程状态
func matchAndSetThreadStatus(state string ,name string)  {
	switch state {
	case "handshake":
		threadStatus.With(prometheus.Labels{"threadName": name}).Set(0)
	case "net_read":
		threadStatus.With(prometheus.Labels{"threadName": name}).Set(1)
	case "net_write":
		threadStatus.With(prometheus.Labels{"threadName": name}).Set(2)
	case "query":
		threadStatus.With(prometheus.Labels{"threadName": name}).Set(3)
	case "net_idle":
		threadStatus.With(prometheus.Labels{"threadName": name}).Set(4)
	default:
		threadStatus.With(prometheus.Labels{"threadName": name}).Set(5)
		//log.Println("matchAndSetThreadStatus没有"+state+"类型状态")
	}
}

//收集连接协议
func matchAndSetProto(proto string,name string) {
	switch proto {
	case "sphinx":
		connectProto.With(prometheus.Labels{"threadName": name}).Set(0)
	case "mysql":
		connectProto.With(prometheus.Labels{"threadName": name}).Set(1)
	case "http":
		connectProto.With(prometheus.Labels{"threadName": name}).Set(2)
	case "ssl":
		connectProto.With(prometheus.Labels{"threadName": name}).Set(3)
	case "compressed":
		connectProto.With(prometheus.Labels{"threadName": name}).Set(4)
	case "replication":
		connectProto.With(prometheus.Labels{"threadName": name}).Set(5)
	case "combination":
		connectProto.With(prometheus.Labels{"threadName": name}).Set(6)
	default:
		connectProto.With(prometheus.Labels{"threadName": name}).Set(7)
	}
}

//分割字符串查询最近一分钟查询总次数
func splitValue(value string) string{
	if value == ""{
		return ""
	}
	split := strings.Split(value, ",")
	secondSplit := strings.Split(split[0], ":")
	if len(secondSplit)<=1 {
		return ""
	}
	return  secondSplit[1]
}

//分割字符串获取最近一分钟每次查询行数平均值
func splitRowValue(value string) string{
	if value == ""{
		return ""
	}
	split := strings.Split(value, ",")
	if len(split) <=1 {
		return ""
	}
	secondSplit := strings.Split(split[1], ":")
	if len(secondSplit)<=1 {
		return ""
	}
	return  secondSplit[1]
}

//读取配置文件并转成结构体
func ReadConfig(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path)   //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config)   //解析成结构体
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}

//初始化数据库
func initSql(){
	var filepath  = "src/main/manticore.conf"
	config,err := ReadConfig(filepath)   //也可以通过os.arg或flag从命令行指定配置文件路径

	if err != nil {
		loger.Fatal(err)
	}

	loger.Println(config.ManticoreDriverName)
	loger.Println(config.Username +":"+config.Password +"@("+config.Ip +":"+config.Port +")/"+config.DatabaseName)
	// 设置连接数据库的参数
	db, err =sql.Open(config.ManticoreDriverName,config.Username+":"+config.Password+"@("+config.Ip+":"+config.Port+")/"+config.DatabaseName)
	if err != nil {
		loger.Fatal(err)
	}
}

func init() {
	file,er:=os.Open("D:/manticore.metrics/src/log.txt")
	defer func(){ _ = file.Close() }()
	if er!=nil && !os.IsExist(er){
		file, _ = os.Create("D:/manticore.metrics/src/log.txt")

	}
	filepath := "D:/manticore.metrics/src/log.txt"
	logFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	loger = log.New(logFile, "[orcale_query]", log.LstdFlags|log.Llongfile|log.LUTC) // 将文件设置为loger作为输出
}

func main() {
	initSql()
	defer db.Close()    //关闭数据库
	err:=db.Ping()      //连接数据库

	if err != nil{
		defer db.Close()    //关闭数据库
		loger.Fatal("数据库连接失败")
		return
	}

	recordMetrics()  //开始记录manticore指标

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
