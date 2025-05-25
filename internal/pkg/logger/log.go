package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func init() {

	v2 := &lumberjack.Logger{
		Filename:   "./logs/number.sender.log",
		MaxSize:    10, // megabytes
		MaxBackups: 15,
		MaxAge:     15,   //days
		Compress:   true, // disabled by default
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(v2),
		zap.InfoLevel,
	)

	Log = zap.New(core)
	zap.ReplaceGlobals(Log)
}
