package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/router"
	"github.com/hhxsv5/gin-x/examples/router/controllers"
)

func main() {
	ng := gin.New()
	r := router.NewSlimRouter(ng)
	//r.Use(middlewares.Global()) // Global middleware

	// Register route: POST http://127.0.0.1:5200/xxx
	// Register route: GET http://127.0.0.1:5200/yyy
	r.RegisterController(controllers.Test{}.NewController())

	// Register route: POST http://127.0.0.1:5200/user/create
	// Register route: GET http://127.0.0.1:5200/user/list
	r.RegisterGroup("user" /*, middlewares.Auth()*/).RegisterController(controllers.User{}.NewController())
	ng.Run(":5200")
}
