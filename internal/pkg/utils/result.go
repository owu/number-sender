package utils

import (
	"github.com/gin-gonic/gin"
)

type Result struct {
	Error int32       `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func Fail(c *gin.Context, result Result) {
	c.JSON(401, result)
	c.Abort()
}

func Success(c *gin.Context, data interface{}) {
	result := Succeed
	result.Data = data
	c.JSON(200, result)
}
