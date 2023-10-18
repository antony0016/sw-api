package controller

import (
	"github.com/antony0016/sw-api/internal/service"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	testService := service.NewTestService("test")
	c.JSON(200, gin.H{
		"message": testService.Ping() + " " + testService.GetName(),
	})
}
