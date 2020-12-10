package main


import (
	"database/sql"
	"log"
	"testing"
)

func Test_init(t *testing.T){
	var filepath  = "D:/manticore.metrics/src/main/manticore.conf"
	config,err := ReadConfig(filepath)   //也可以通过os.arg或flag从命令行指定配置文件路径

	if err != nil {
		t.Fatal(err)
	}

	log.Println(config.ManticoreDriverName)
	log.Println(config.Username +":"+config.Password +"@("+config.Ip +":"+config.Port +")/"+config.DatabaseName)
	// 设置连接数据库的参数
	db, err =sql.Open(config.ManticoreDriverName,config.Username+":"+config.Password+"@("+config.Ip+":"+config.Port+")/"+config.DatabaseName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("init成功")
}