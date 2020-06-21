package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)


var (
	Log *logrus.Logger
)

func SettingInit(logFile string){
	//logfile, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil{
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level:level,
		Out:os.Stdout,
		Formatter: &logrus.JSONFormatter{},
	}
}

func Info(msg string, tags ...string){
	if Log.Level < logrus.InfoLevel{
		return
	}
	Log.WithFields(parseFields(tags...)).Info(msg)
}

func Debug(msg string, tags ...string){
	if Log.Level < logrus.DebugLevel{
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(msg)
}
func Error(msg string, err error, tags ...string){
	if Log.Level < logrus.ErrorLevel{
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.WithFields(parseFields(tags...)).Error(msg)
}

func parseFields(tags ...string)logrus.Fields{
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags{
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}
