package services

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

const (
	ErrorLog = "error.log"
	DebugLog = "debug.log"
)

var Logger *logrus.Logger
var errorFile *os.File
var debugFile *os.File

func CreateLogger() (err error) {
	errorfile, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	debugfile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return

	}

	Logger = logrus.New()
	Logger.SetOutput(ioutil.Discard)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.AddHook(&writer.Hook{
		Writer: errorfile,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})

	Logger.AddHook(&writer.Hook{
		Writer: debugfile,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})

	return
}
func LogToFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		Logger.Infof(
			"| %3d| %5v | %10s | %s | %s |",
			ctx.Writer.Status(),
			latencyTime,
			ctx.ClientIP(),
			ctx.Request.Method,
			ctx.Request.RequestURI,
		)
	}
}
func CloseLogger() {
	if errorFile != nil {
		errorFile.Close()
	}
	if debugFile != nil {
		debugFile.Close()
	}
}
