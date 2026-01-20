package mware

import (
	"github.com/gin-gonic/gin"
	"number-sender/internal/pkg/limit"
	"number-sender/internal/pkg/utils"
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
