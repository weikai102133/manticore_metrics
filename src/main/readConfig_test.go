package main


import (
	"strings"
	"testing"
)

func Test_ReadConfig(t *testing.T) {
	var filepath  = "D:/manticore.metrics/src/main/manticore.conf"
	config, e := ReadConfig(filepath)
	path := config.Username +":"+config.Password +"@("+config.Ip +":"+config.Port +")/"+config.DatabaseName
	if e!= nil || strings.Compare(path,"root:root@(127.0.0.1:9306)/Manticore")!=0{
		t.Fatalf("readConfig失败")
	}
	t.Logf("test readConfig成功")
}