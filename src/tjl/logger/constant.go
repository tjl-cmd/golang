package logger

const (
	LogLevelDebug = iota
	LogLeveTrace
	LogLevelInfo
	LogLeveWarn
	LogLeveError
	LogLeveFatal
)
const (
	LogSplitTypeHour = iota
	LogSplitTypeSize
)

func getLevelText(level int) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLeveTrace:
		return "TRACE"
	case LogLevelInfo:
		return "INFO"
	case LogLeveWarn:
		return "WARN"
	case LogLeveError:
		return "ERROR"
	case LogLeveFatal:
		return "FATAL"
	}
	return "UNKNOWN"
}

func getLevel(level string) int {
	switch level {
	case "debug":
		return LogLevelDebug
	case "trace":
		return LogLeveTrace
	case "info":
		return LogLevelInfo
	case "warn":
		return LogLeveWarn
	case "error":
		return LogLeveError
	case "fatal":
		return LogLeveFatal
	}
	return LogLevelDebug
}
