package mware

import (
	"github.com/gin-gonic/gin"
	"github.com/owu/number-sender/internal/pkg/limit"
	"github.com/owu/number-sender/internal/pkg/utils"
)

func LimitMw(limit *limit.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limit.TakeAvailable(c.FullPath()) {
			utils.Fail(c, utils.LimitError)
			return
		}
		c.Next()
	}
}
