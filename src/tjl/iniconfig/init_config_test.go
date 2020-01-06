package iniconfig

import (
	"fmt"
	"io/ioutil"
	"testing"
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

func (c *Config) ConfigSet(ip string, port int, username string, passwd string, database string, host string, port1 int, timeout float32) {
	c.ServerConf.IP = ip
	c.ServerConf.Port = port
	c.MysqlConf.Username = username
	c.MysqlConf.Passwd = passwd
	c.MysqlConf.Database = database
	c.MysqlConf.Host = host
	c.MysqlConf.Port = port1
	c.MysqlConf.Timeout = timeout
}
func TestIniConfig(t *testing.T) {
	fmt.Println("hello world")
	//t.Error("http")
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("read file failed")
	}
	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Error("Unmarshal failed err:", err)
	}

	confData, err := Marshal(conf)
	t.Logf("marshal success,conf:%s", string(confData))
	t.Logf("unmarshal success conf:%#v", conf)
	MarshalFile("C:/logs/test.conf", conf)
}
func TestIniConfigFile(t *testing.T) {
	var conf Config
	filename := "C:/logs/test.conf"
	err := MarshalFile(filename, conf)
	if err != nil {
		t.Errorf("unmarshal failed ,err :%v", err)
		return
	}
	var conf2 Config
	err = UnMarshalFile(filename, &conf2)
	if err != nil {
		t.Errorf("unmarshal failed err :%v", err)
	}
	t.Logf("unmarshal success conf:%#v", conf2)
}
