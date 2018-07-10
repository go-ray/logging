package logging

import (
	"fmt"
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger(logpath, name, level, format string) {
	baseLogPath := path.Join(logpath, name)
	baseLogPath, _ = filepath.Abs(baseLogPath)
	fmt.Println("base log path:", baseLogPath)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}
	var formatter logrus.Formatter
	switch format {
	case "json":
		formatter = &logrus.JSONFormatter{}
	default:
		formatter = &logrus.TextFormatter{}
	}
	logger = logrus.New()
	logger.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		},
		formatter,
	))

	switch level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	default:
		logger.SetLevel(logrus.DebugLevel)
	}

}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Print(args ...interface{}) {
	logger.Print(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Println(args ...interface{}) {
	logger.Println(args...)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

func Warningln(args ...interface{}) {
	logger.Warningln(args...)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}
