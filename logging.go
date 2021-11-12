package go_logging

import (
	"os"
	"strings"

	logrus "github.com/sirupsen/logrus"
)

var (
	//Log is
	Log = logrus.New()
)

func init() {
	InitLog("")
}

//=============log=========================
func convertStringToLogLevel(logLevel string) logrus.Level {
	switch strings.ToLower(logLevel) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}

//ChangeLogLevel is
func ChangeLogLevel(logLevel string) {
	Log.SetLevel(convertStringToLogLevel(logLevel))
}

//InitLog is
func InitLog(logLevel string) {
	Log.SetFormatter(&logrus.JSONFormatter{FieldMap: logrus.FieldMap{
		logrus.FieldKeyTime:  "timestamp",
		logrus.FieldKeyLevel: "level",
		logrus.FieldKeyMsg:   "message",
		logrus.FieldKeyFunc:  "caller"},
	})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(convertStringToLogLevel(logLevel))
	Log.ReportCaller = true
}
