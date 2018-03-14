package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-slim-router/examples/controllers"
	"github.com/hhxsv5/gin-slim-router"
)

func main() {
	engine := gin.New()
	router := gsr.NewSlimRouter(engine)
	//router.Use(middlewares.Global()) // Global middleware

	// Register route: POST http://127.0.0.1:5200/xxx
	// Register route: GET http://127.0.0.1:5200/yyy
	router.RegisterController(controllers.Test{}.NewController())

	// Register route: POST http://127.0.0.1:5200/user/create
	// Register route: GET http://127.0.0.1:5200/user/list
	router.RegisterGroup("user" /*, middlewares.Auth()*/).RegisterController(controllers.User{}.NewController())
	engine.Run(":5200")
}
