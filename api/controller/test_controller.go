package controller

import (
	"github.com/antony0016/sw-api/api/http"
	"github.com/antony0016/sw-api/internal/repository"
	"github.com/antony0016/sw-api/internal/usercase"
	"github.com/antony0016/sw-api/utils/logger"
	"github.com/gin-gonic/gin"
)

type TestController struct {
	testRepository *repository.TestRepository
}

func (tr *TestController) Register(r *gin.Engine) {
	testRouterGroup := r.Group("/test")
	testRouterGroup.GET("/ping", tr.Ping)
	testRouterGroup.POST("/insert", tr.TestInsert)
}

func NewTestController(testRepository *repository.TestRepository) *TestController {
	return &TestController{testRepository: testRepository}
}

func (tr *TestController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (tr *TestController) TestInsert(c *gin.Context) {
	var testReq http.TestRequest
	err := c.Bind(&http.TestRequest{})
	if err != nil {
		c.JSON(400, gin.H{
			"message": "request error",
		})
		logger.Printf("Bind error: %v", err)
		return
	}
	if usercase.SaveTestItem(tr.testRepository, testReq.Id, testReq.Test) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(500, gin.H{
			"message": "fail",
		})
	}
}
