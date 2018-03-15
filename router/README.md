gin-slim-router
======
A slim router for gin framework

## Usage

1. Create controller
```Go
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/router"
)

type User struct {
	// HttpMethod: Any/GET/POST/DELETE/PATCH/PUT/OPTIONS/HEAD
	// Format: {HttpMethod}Action func(*gin.Context) `path:"/xxxpath"`
	PostCreate func(*gin.Context) `path:"/create"`
	GetList    func(*gin.Context) `path:"/list"`
}

func (u User) NewController() router.Controller {
	u.PostCreate = create
	u.GetList = getList
	return u
}

func create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "create user")
}

func getList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get user list")
}
```

2. Register route
```Go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/router"
	"github.com/hhxsv5/gin-x/examples/controllers"
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
```

## License

[MIT](https://github.com/hhxsv5/gin-slim-router/blob/master/LICENSE)
