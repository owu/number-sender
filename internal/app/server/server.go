package server

import (
	"github.com/gin-gonic/gin"
	"github.com/owu/number-sender/internal/app/workers"
	"github.com/owu/number-sender/internal/pkg/calculate"
	"github.com/owu/number-sender/internal/pkg/config"
	"github.com/owu/number-sender/internal/pkg/limit"
	"github.com/owu/number-sender/internal/pkg/mware"
	"github.com/owu/number-sender/internal/pkg/redis"

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
	//gin.SetMode(gin.ReleaseMode)

	if _, err := options.workersInstance.Cron.AddFunc("@every 5s", func() {
		Task(options.defaultRedis, options.chains)
	}); err != nil {
		panic("server Cron.AddFunc failed," + err.Error())
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
