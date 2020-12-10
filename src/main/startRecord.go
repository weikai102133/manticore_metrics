package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	InitSql()
	defer db.Close()    //关闭数据库
	err:=db.Ping()      //连接数据库


	if err != nil{
		defer db.Close()    //关闭数据库
		loger.Fatal("数据库连接失败")
		return
	}

	RecordMetrics() //开始记录manticore指标

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
