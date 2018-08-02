package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/examples/router/controllers"
	"github.com/hhxsv5/gin-x/router"
)

func main() {
	ng := gin.New()
	r := router.NewSlimRouter(ng)
	//r.Use(middlewares.Global()) // Global middleware

	// Register routes: POST http://127.0.0.1:5200/xxx
	// Register routes: GET http://127.0.0.1:5200/yyy
	// Register routes: ANY http://127.0.0.1:5200/any
	r.RegisterController(controllers.Test{}.NewController())

	// Register routes: POST http://127.0.0.1:5200/user/create
	// Register routes: GET http://127.0.0.1:5200/user/list
	r.RegisterGroup("api").RegisterGroup("user" /*, middlewares.Auth()*/).RegisterController(controllers.User{}.NewController())
	ng.Run(":5200")
}
