package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type IAppLogger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type AppLogger struct {
	logger *log.Logger
}

func (l *AppLogger) Info(args ...interface{}) {
	l.logger.Infoln(args)
}

func (l *AppLogger) Error(args ...interface{}) {
	l.logger.Errorln(args)
}

func (l *AppLogger) Warning(args ...interface{}) {
	l.logger.Warnln(args)
}

func (l *AppLogger) Fatal(args ...interface{}) {
	l.logger.Fatalln(args)
}

func (l *AppLogger) Debug(args ...interface{}) {
	l.logger.Debugln(args)
}

func (l *AppLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args)
}

func (l *AppLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorln(format, args)
}

func (l *AppLogger) Warningf(format string, args ...interface{}) {
	l.logger.Warnln(format, args)
}

func (l *AppLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalln(format, args)
}

func (l *AppLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugln(format, args)
}

func New() IAppLogger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			nextPtr, nextFile, _, ok := runtime.Caller(9)
			nextFunc := runtime.FuncForPC(nextPtr)
			if ok {
				wd, _ := os.Getwd()
				nextFilePath, _ := filepath.Rel(wd, nextFile)
				_, nextFileLine := nextFunc.FileLine(nextPtr)
				funcName := strings.Split(nextFunc.Name(), ".")
				return fmt.Sprintf(" %s:%s:%d", nextFilePath, funcName[len(funcName)-1], nextFileLine), ""
			}
			return "", ""
		},
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.InfoLevel)
	logger.SetReportCaller(true)
	defLogger := logger
	return &AppLogger{
		logger: defLogger,
	}
}
