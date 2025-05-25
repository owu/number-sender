package server

import (
	"github.com/gin-gonic/gin"
	"github.com/owu/number-sender/internal/pkg/consts"
	"github.com/owu/number-sender/internal/pkg/model"
	"github.com/owu/number-sender/internal/pkg/redis"
	"github.com/owu/number-sender/internal/pkg/utils"
	"go.uber.org/zap"
	"strings"
	"time"
)

func Ping(c *gin.Context) {
	utils.Success(c, map[string]int64{
		"time": time.Now().UnixMilli(),
	})
}

func Len(c *gin.Context, defaultRedis *redis.DefaultRedis) {
	data := model.Fetch{
		Starter:  defaultRedis.Len(consts.Starter),
		Standard: defaultRedis.Len(consts.Standard),
		Premium:  defaultRedis.Len(consts.Premium),
		Ultimate: defaultRedis.Len(consts.Ultimate),
	}

	zap.L().Info("Len.data",
		zap.Any("data", data),
	)

	utils.Success(c, data)
}

func Pop(c *gin.Context, defaultRedis *redis.DefaultRedis) {

	plan := consts.Plans(strings.ToLower(c.Param("plan")))

	data := model.Fetch{}

	switch plan {
	case consts.Starter:
		data.Starter = defaultRedis.Pop(plan)
	case consts.Standard:
		data.Standard = defaultRedis.Pop(plan)
	case consts.Premium:
		data.Premium = defaultRedis.Pop(plan)
	case consts.Ultimate:
		data.Ultimate = defaultRedis.Pop(plan)
	default:
		c.JSON(200, utils.MissingParams)
		return
	}

	utils.Success(c, data)
}
