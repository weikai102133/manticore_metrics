package main

import (
	"gopkg.in/ini.v1"
	"log"
)

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
