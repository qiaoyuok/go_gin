package utils

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_gin/config"
	"time"
)

var ZapSugarLogger *zap.SugaredLogger
var ZapLogger *zap.Logger

func RegisterLogger() {
	path := config.C.Log.Path
	appLog, _ := rotatelogs.New(
		path+config.C.Log.AppLogFileName+".%Y%m%d",
		rotatelogs.WithMaxAge(config.C.Log.MaxAge*24*time.Hour), // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24),               // 24小时切割一次
	)
	errLog, _ := rotatelogs.New(
		path+config.C.Log.ErrLogFileName+".%Y%m%d",
		rotatelogs.WithMaxAge(config.C.Log.MaxAge*24*time.Hour), // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24),               // 24小时切割一次
	)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	appL := zapcore.NewCore(encoder, zapcore.AddSync(appLog), zapcore.InfoLevel)
	errL := zapcore.NewCore(encoder, zapcore.AddSync(errLog), zap.ErrorLevel)

	core := zapcore.NewTee(appL, errL)
	ZapLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	ZapSugarLogger = ZapLogger.Sugar()
}
