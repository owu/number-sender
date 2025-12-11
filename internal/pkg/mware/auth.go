package mware

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/owu/number-sender/internal/pkg/config"
	"github.com/owu/number-sender/internal/pkg/utils"
)

func AuthMw(configs *config.LoadConfigs) gin.HandlerFunc {
	return func(c *gin.Context) {

		if configs.Env() == "test" {
			c.Next()
			return
		}

		milli := c.GetHeader("Milli")
		token := c.GetHeader("Token")

		if milli == "" || token == "" {
			utils.Fail(c, utils.MissingParams)
			return
		}

		ms, err := strconv.ParseInt(milli, 10, 64)
		if err != nil {
			utils.Fail(c, utils.TimestampError)
			return
		}

		if ms < time.Now().UnixMilli()-60000 || ms > time.Now().UnixMilli()+60000 {
			utils.Fail(c, utils.TimestampExpired)
			return
		}

		if token != fmt.Sprintf("%x", md5.Sum([]byte(milli+","+configs.Encrypt()))) {
			utils.Fail(c, utils.AuthFailed)
			return
		}
		c.Next()
	}
}
