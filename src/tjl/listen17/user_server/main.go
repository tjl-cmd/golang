package main

import (
	"time"
	"tjl/logger"
)

func initLogger(name, logPath, logName string, level string) (err error) {
	m := make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = logName
	m["log_level"] = level
	m["log_split_type"] = "size"
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	//log = logger.NewConsoleLog(level)
	//log.Debug("init logger success")
	return
}
func Run() {
	for {
		logger.Debug("user server is running C:\\Users\\27409\\Desktop\\golang\\src\\tjl\\logger>")
		time.Sleep(time.Second)
	}
}
func main() {
	_ = initLogger("file", "c:/logs/", "user_server", "debug")
	//initLogger("console", "c:/logs/", "user_server", "debug")
	Run()
	return
	/*
		file := log.NewFileLog("c:/a.log")
		file.LogDebug("this is a debug log")
		file.LogWarn("this is a warn log ")
	*/
	/*
		console := log.NewConsoleLog("xxxx")
		console.LogConsoleDebug("this is a console log")
		console.LogConsoleWarn("this is a warn log")
	*/
	//log1 := log.NewFileLog("c:/a.log")
	//log1.LogWarn("this is a console log")
	//log1.LogDebug("this is a warn log")
}
