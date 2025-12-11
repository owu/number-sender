package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/owu/number-sender/internal/app/workers"
	"github.com/owu/number-sender/internal/pkg/calculate"
	"github.com/owu/number-sender/internal/pkg/config"
	"github.com/owu/number-sender/internal/pkg/limit"
	"github.com/owu/number-sender/internal/pkg/logger"
	"github.com/owu/number-sender/internal/pkg/mware"
	"github.com/owu/number-sender/internal/pkg/redis"
	"go.uber.org/zap"

	"log"
)

type Options struct {
	configs         *config.LoadConfigs
	defaultRedis    *redis.DefaultRedis
	workersInstance *workers.Workers
	chains          *calculate.Chains
	limiter         *limit.Limiter
}

func initApp(options Options) error {
	// 根据环境配置设置Gin模式
	if options.configs.Env() == "production" {
		gin.SetMode(gin.ReleaseMode)
		logger.Log.Info("Gin mode set to ReleaseMode")
	} else {
		gin.SetMode(gin.DebugMode)
		logger.Log.Info("Gin mode set to DebugMode")
	}

	if _, err := options.workersInstance.Cron.AddFunc("@every 5s", func() {
		Task(options.defaultRedis, options.chains)
	}); err != nil {
		logger.Log.Error("server Cron.AddFunc failed", zap.Error(err))
		return fmt.Errorf("server Cron.AddFunc failed: %w", err)
	}

	//gin init
	r := gin.New()

	//init LoggerMw
	mware.LoggerMw(r)

	//api router
	RegMux(r, options.configs, options.defaultRedis, options.limiter)

	//run
	log.Fatal(r.Run(options.configs.HttpPort()))

	return nil
}
