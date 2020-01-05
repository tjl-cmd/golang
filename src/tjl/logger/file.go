package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileLogger struct {
	level         int
	logPath       string
	logName       string
	file          *os.File
	warnFile      *os.File
	LogDataChan   chan *LogData
	LogSplitType  int
	LogSplitSize  int64
	lastSplitHour int
}

func NewFileLogger(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not foun log_path")
		return
	}
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not foun log_name")
		return
	}
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not foun log_level")
		return
	}
	logChanSize, ok := config["log_level"]
	if !ok {
		logChanSize = "50000"
	}

	if !ok {
		logChanSize = "50000"
	}
	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}
	var logSplitType int = LogSplitTypeHour
	var logSplitSize int64
	logSplitStr, ok := config["log_split_type"]
	if !ok {
		logSplitStr = "hour"
	} else {
		if logSplitStr == "size" {
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok {
				logSplitSizeStr = "104857600"
			}
			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 104857600
			}
			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}
	level := getLevel(logLevel)
	log = &FileLogger{
		level:         level,
		logPath:       logPath,
		logName:       logName,
		LogDataChan:   make(chan *LogData, chanSize),
		LogSplitSize:  logSplitSize,
		LogSplitType:  logSplitType,
		lastSplitHour: time.Now().Hour(),
	}
	fmt.Println(logSplitSize, logSplitType, logSplitStr)
	log.Init()
	return
}
func (f *FileLogger) Init() {
	filename := fmt.Sprintf("%s%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 07755)
	if err != nil {
		panic(fmt.Sprintf("open faile  %s failed ,err:%v", filename, err))
	}
	f.file = file
	//写错误日志和fatal日志文件
	filename = fmt.Sprintf("%s%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 07755)
	if err != nil {
		panic(fmt.Sprintf("open faile  %s failed ,err:%v", filename, err))
	}
	f.warnFile = file
	go f.writeLogBackground()
}

//按小时进行切分
func (f *FileLogger) splitFileHour(warnFile bool) {
	now := time.Now()
	hour := now.Hour()
	if hour == f.lastSplitHour {
		return
	}
	f.lastSplitHour = hour
	//备份日志文件
	var backupFilename string
	var filename string
	if warnFile {
		backupFilename = fmt.Sprintf("%s%s.log.wf_%04d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s%s.log.wf", f.logPath, f.logName)
	} else {
		backupFilename = fmt.Sprintf("%s%s.log_%04d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s%s.log", f.logPath, f.logName)
	}
	file := f.file
	if warnFile {
		file = f.warnFile
	}
	file.Close()
	os.Rename(filename, backupFilename)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 07755)
	if err != nil {
		return
	}
	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}

//按大小进行切分
func (f *FileLogger) splitFileSize(warnFile bool) {
	file := f.file
	if warnFile {
		file = f.warnFile
	}
	statInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := statInfo.Size()
	if fileSize <= f.LogSplitSize {
		return
	}
	//备份日志文件
	var backupFilename string
	var filename string
	now := time.Now()
	if warnFile {
		backupFilename = fmt.Sprintf("%s%s.log.wf_%04d%02d%02d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s%s.log.wf", f.logPath, f.logName)
	} else {
		backupFilename = fmt.Sprintf("%s%s.log_%04d%02d%02d%02d%02d%02d", f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s%s.log", f.logPath, f.logName)
	}
	file.Close()
	os.Rename(filename, backupFilename)

	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 07755)
	if err != nil {
		return
	}
	if warnFile {
		f.warnFile = file
	} else {
		f.file = file
	}
}
func (f *FileLogger) checkSplitFile(warnFile bool) {

	if f.LogSplitType == LogSplitTypeHour {
		f.splitFileHour(warnFile)
		return
	}
	f.splitFileSize(warnFile)
}
func (f *FileLogger) writeLogBackground() {
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnFile
		}
		f.checkSplitFile(logData.WarnAndFatal)
		fmt.Fprintf(file, "%s %s (%s %s:%d)  %s\n", logData.TimeStr, logData.levelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)

	}
}
func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}

	logData := writeLog(LogLevelDebug, format, args...)
	//判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLeveTrace {
		return
	}
	logData := writeLog(LogLeveTrace, format, args...)
	//判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	//判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLeveWarn {
		return
	}
	logData := writeLog(LogLeveWarn, format, args...)
	//判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLeveError {
		return
	}
	logData := writeLog(LogLeveError, format, args...)
	//判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLeveFatal {
		return
	}
	logData := writeLog(LogLeveFatal, format, args...)
	//判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLeveFatal {
		level = LogLevelDebug
	}
	f.level = level

}
func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
