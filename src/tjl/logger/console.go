package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger(config map[string]string) (log LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not foun log_level")
		return
	}
	level := getLevel(logLevel)
	log = &ConsoleLogger{
		level: level,
	}

	//logger.init()
	return
}
func (c *ConsoleLogger) Init() {

}
func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLeveFatal {
		level = LogLevelDebug
	}
	c.level = level
}
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	logData := writeLog(LogLevelDebug, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s %s:%d)  %s\n", logData.TimeStr, logData.levelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)

}
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLeveTrace {
		return
	}
	//writeLog(os.Stdout, LogLeveTrace, format, args...)
	logData := writeLog(LogLeveTrace, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s %s:%d)  %s\n", logData.TimeStr, logData.levelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
	//判断队列是否满
	//select {
	//case c.LogDataChan <- logData:
	//default:
	//}
}
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s %s:%d)  %s\n", logData.TimeStr, logData.levelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)

}
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLeveWarn {
		return
	}
	logData := writeLog(LogLeveWarn, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s %s:%d)  %s\n", logData.TimeStr, logData.levelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)

}
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLeveError {
		return
	}
	logData := writeLog(LogLeveError, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s %s:%d)  %s\n", logData.TimeStr, logData.levelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)

}
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLeveFatal {
		return
	}
	logData := writeLog(LogLeveFatal, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s %s:%d)  %s\n", logData.TimeStr, logData.levelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)

}
func (c *ConsoleLogger) Close() {

}
