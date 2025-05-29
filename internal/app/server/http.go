package server

import (
	"github.com/gin-gonic/gin"
	"github.com/owu/number-sender/internal/pkg/config"
	"github.com/owu/number-sender/internal/pkg/limit"
	"github.com/owu/number-sender/internal/pkg/mware"
	"github.com/owu/number-sender/internal/pkg/redis"
)

func RegMux(r *gin.Engine, configs *config.LoadConfigs, defaultRedis *redis.DefaultRedis, limit *limit.Limiter) {

	//r.SetTrustedProxies([]string{"127.0.0.1"})

	//health
	r.GET("/ping", Ping)

	api := r.Group("/api", mware.AuthMw(configs))
	{
		//len
		api.GET("/len", func(c *gin.Context) {
			Len(c, defaultRedis)
		})

		//pop, {:plan} ->  starter,standard,premium,ultimate
		api.GET("/pop/:plan", mware.LimitMw(limit), func(c *gin.Context) {
			Pop(c, defaultRedis)
		})

	}
}
