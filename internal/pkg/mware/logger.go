package mware

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"number-sender/internal/pkg/logger"
	"time"
)

func LoggerMw(r *gin.Engine) {
	// Add a ginzap ware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	// r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.GinzapWithConfig(logger.Log, &ginzap.Config{TimeFormat: time.RFC3339, UTC: true, DefaultLevel: zapcore.InfoLevel}))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger.Log, true))
}
