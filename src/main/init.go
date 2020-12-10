package main

import (
	"database/sql"
	"log"
	"os"
)

//初始化数据库
func InitSql(){
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
	//配置日志设置
	file,er:=os.Open("D:/manticore.metrics/src/log.txt")
	defer func(){ _ = file.Close() }()
	if er!=nil && !os.IsExist(er){
		file, _ = os.Create("D:/manticore.metrics/src/log.txt")
	}
	filepath := "D:/manticore.metrics/src/log.txt"
	logFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		loger.Println(err)
	}
	loger = log.New(logFile, "[orcale_query]", log.LstdFlags|log.Llongfile|log.LUTC) // 将文件设置为loger作为输出
}

