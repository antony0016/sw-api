package inits

import "github.com/gin-gonic/gin"

func WebInit(r *gin.Engine) {
	// log.Println("Web init")
}

func MiddlewareInit(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}

func RouterInit(r *gin.Engine) {
	// log.Println("Router init")
}
