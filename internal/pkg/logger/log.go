package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
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
	defer func() {
		if err := v2.Close(); err != nil {
			log.Printf("failed to close log file: %v", err)
		}
	}()

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(v2),
		zap.InfoLevel,
	)

	Log = zap.New(core)
	zap.ReplaceGlobals(Log)

	defer func() {
		if err := Log.Sync(); err != nil {
			log.Printf("failed to close log file: %v", err)
		}
	}()
}
