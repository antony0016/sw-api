package router

import (
	"github.com/antony0016/sw-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	testRouterGroup := r.Group("/ping")
	InitTestRouter(testRouterGroup)
}

func InitTestRouter(r *gin.RouterGroup) {
	r.GET("/", controller.Ping)
}
