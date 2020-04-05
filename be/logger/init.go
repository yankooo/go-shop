package logger

import (
	"github.com/yankooo/school-eco/be/config"
	"os"
	"time"
	"github.com/yankooo/school-eco/be/model"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"

	"github.com/sirupsen/logrus"
)

const _TIME_LAYOUT = "2006-01-02 15:04:05.000000000"

var _G_log *logrus.Logger

func Debugf(format string, args ...interface{}) {
	_G_log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	_G_log.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	_G_log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	_G_log.Fatalf(format, args...)
}

func Error(args ...interface{}) {
	_G_log.Error(args...)
}

func Debugln(args ...interface{}) {
	_G_log.Debugln(args...)
}

/*根据配置文件初始化日志*/
func InitLogger() error {
	return newLogger(config.GlobalConf().LoggerConfig)
}

func newLogger(loggerConfig *model.LoggerConfig) (err error) {
	_G_log = logrus.New()
	if loggerConfig.Mode == "dev" {
		_G_log.Out = os.Stdout
	} else if loggerConfig.Mode == "release" {
		_G_log.Out, err = os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}

		var apiLogPath = loggerConfig.Path
		logWriter, _ := rotatelogs.New(
			apiLogPath+".%Y-%m-%d-%H-%M._G_log",
			rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
			rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.FatalLevel: logWriter,
		}
		lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: _TIME_LAYOUT,
		})
		_G_log.AddHook(lfHook)
	}

	// 设置日志级别
	SetLoggerLevel(loggerConfig)

	_G_log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	return nil
}

func SetLoggerLevel(config *model.LoggerConfig) {
	switch config.Level {
	case "info":
		_G_log.SetLevel(logrus.InfoLevel)
	case "debug":
		_G_log.SetLevel(logrus.DebugLevel)
	default:
		_G_log.SetLevel(logrus.DebugLevel)
	}
}
