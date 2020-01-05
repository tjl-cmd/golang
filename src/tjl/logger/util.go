package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogData struct {
	Message      string
	TimeStr      string
	levelStr     string
	FileName     string
	FuncName     string
	LineNo       int
	WarnAndFatal bool
}

/*
当业务调用打日志的方法是，我们把日志相关的数据写入到chan(队列)
有个后台的线程不断的从chan里面获取这些日志，最终写入到文件中
*/
func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}
func writeLog(level int, format string, args ...interface{}) *LogData {
	now := time.Now()
	nowStr := now.Format("2006-01-01 15:04:05.999")
	LevelStr := getLevelText(level)
	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)
	logData := &LogData{
		Message:      msg,
		TimeStr:      nowStr,
		levelStr:     LevelStr,
		FileName:     fileName,
		FuncName:     funcName,
		LineNo:       lineNo,
		WarnAndFatal: false,
	}
	if level == LogLeveError || level == LogLeveWarn || level == LogLeveFatal {
		logData.WarnAndFatal = true
	}
	return logData
	//fmt.Fprintf(file, "%s %s (%s %s:%d)  %s\n", nowStr, LevelStr, fileName, funcName, lineNo, msg)
}
