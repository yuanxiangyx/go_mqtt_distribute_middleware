package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack2 "gopkg.in/natefinch/lumberjack.v2"
	"mqtt_pro/config"
)

func InitLogger(logName string, cfg *config.LogOption) (log *zap.Logger) {
	writeSyncer := getLogWriter(logName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	return zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack2.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

type LogGroupS struct {
	ErrLog  *zap.Logger
	InfoLog *zap.Logger
}

var LogGroup *LogGroupS

func InitGroupLog(option *config.LogOption) {
	LogGroup = &LogGroupS{
		InfoLog: InitLogger("logs/info.log", option),
		ErrLog:  InitLogger("logs/err.log", option),
	}
}

func WInfo(msg string) {
	LogGroup.InfoLog.Info(msg)
}

func WError(msg string) {
	LogGroup.ErrLog.Error(msg)
}
