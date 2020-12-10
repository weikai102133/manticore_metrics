package main

import (
	"database/sql"
	"testing"
)



func Test_splitValue(t *testing.T)  {
	value := SplitValue("this is splitValuetest:1200,this is splitValuetest:200,this is splitValuetest:100")
	if value != "1200" {
		t.Fatal("TestsplitValue失败")
	}
	t.Logf("TestsplitValue成功")
}

func Test_splitRowValue(t *testing.T) {
	value := SplitRowValue("this is splitValuetest:1200,this is splitValuetest:200,this is splitValuetest:100")
	if value != "200" {
		t.Fatal("TestsplitValue失败")
	}
	t.Logf("TestsplitValue成功")
}

func Test_matchAndSetProto(t *testing.T) {
	name := []string{"testrt","products"}
	proto := []string{"sphinx","mysql","http","ssl","compressed","replication","combination","other"}
	for _,s := range name{
		for _,pro := range proto{
			MatchAndSetProto(pro,s)
		}
	}
}

func Test_matchAndSetThreadStatus(t *testing.T) {
	name := []string{"testrt","products"}
	status := []string{"handshake","net_read","net_write","query","net_idle","other"}
	for _,s := range name{
		for _,state := range status{
			MatchAndSetThreadStatus(state,s)
		}
	}
}

func Test_matchAndSetIndexType(t *testing.T) {
	name := []string{"testrt","products"}
	types := []string{"disk","rt","percolate","template","distributed","other"}
	for _,s := range name{
		for _,ty := range types{
			MatchAndSetIndexType(ty,s)
		}
	}
}

func Test_matchAndSetIndexMetric(t *testing.T) {
	name := []string{"testrt","products"}
	variableName := []string{"index_type","indexed_documents","indexed_bytes","ram_bytes","disk_bytes","other"}
	value := "1240"
	for _,s := range name{
		for _,variable := range variableName{
			MatchAndSetIndexMetric(s,variable,value)
		}
	}
}

func Test_showIndexStatus(t *testing.T) {
	var filepath  = "D:/manticore.metrics/src/main/manticore.conf"
	config,err := ReadConfig(filepath)   //也可以通过os.arg或flag从命令行指定配置文件路径

	if err != nil {
		t.Fatal(err)
	}

	t.Logf(config.ManticoreDriverName)
	t.Logf(config.Username +":"+config.Password +"@("+config.Ip +":"+config.Port +")/"+config.DatabaseName)
	// 设置连接数据库的参数
	db, err =sql.Open(config.ManticoreDriverName,config.Username+":"+config.Password+"@("+config.Ip+":"+config.Port+")/"+config.DatabaseName)
	if err != nil {
		t.Fatal(err)
	}
	name := []string{"testrt","products"}

	for _,s := range name{
		ShowIndexStatus(s)
	}
}

