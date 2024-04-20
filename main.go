package main

import (
	"github.com/antony0016/sw-api/api/controller"
	"github.com/antony0016/sw-api/config"
	"github.com/antony0016/sw-api/inits"
	"github.com/antony0016/sw-api/internal/repository"
	"github.com/antony0016/sw-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.Env{}
	env.EnvLoad()

	db, err := inits.DBInit(env)

	if err != nil {
		logger.Fatal("Can't init db.", err)
	} else {
		logger.Println("Connect to db success.")
	}

	testRepository := repository.NewTestRepository(db)
	testController := controller.NewTestController(testRepository)

	r := gin.Default()
	inits.WebInit(r)
	inits.MiddlewareInit(r)
	inits.RouterInit(r)
	testController.Register(r)
	err = r.Run("0.0.0.0:8080")
	if err != nil {
		logger.Fatal("Can't run server.", err)
	} else {
		logger.Println("Server is running on :8080")
	}
}
