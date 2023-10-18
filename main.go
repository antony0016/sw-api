package main

import (
	"fmt"
	"github.com/antony0016/sw-api/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println("Run server failed")
	}
}
