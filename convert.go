package log

import "os"

const (
	ErrorLevelString = "error"
	WarnLevelString  = "warn"
	InfoLevelString  = "info"
	DebugLevelString = "debug"
)

// LevelEnvironmentVariable get the level of logging
var LevelEnvironmentVariable = "LOG_LEVEL"

func LevelToUint() uint8 {
	switch os.Getenv(LevelEnvironmentVariable) {
	case ErrorLevelString:
		return ErrorLevel
	case WarnLevelString:
		return WarnLevel
	case InfoLevelString:
		return InfoLevel
	case DebugLevelString:
		return DebugLevel
	default:
		return DebugLevel
	}
}

func SetLevelEnvironmentVariable(l string) {
	LevelEnvironmentVariable = l
}
