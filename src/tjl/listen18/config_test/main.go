package main

import (
	"fmt"
	"tjl/iniconfig"
)

type Config struct {
	ServerConf SeverConfig `ini:"server"`
	MysqlConf  MysqlConfig `ini:"mysql"`
}
type SeverConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}
type MysqlConfig struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

func main() {
	filename := "C:/logs/test.conf"
	var conf2 Config
	conf2.MysqlConf.Username = "root"
	err := iniconfig.MarshalFile(filename, conf2)
	if err != nil {
		fmt.Println("marshal filed err", err)
		return
	}
	var conf Config
	err = iniconfig.UnMarshalFile(filename, &conf)
	if err != nil {
		fmt.Println("unmarshal filed err :", err)
		return
	}

	fmt.Printf("conf:%#v\n", conf)
}
