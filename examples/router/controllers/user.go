package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/router"
)

type User struct {
	// HttpMethod: Any/GET/POST/DELETE/PATCH/PUT/OPTIONS/HEAD
	// Format: Action func(*gin.Context) `request:"{HttpMethod} /xxxpath"`
	Create func(*gin.Context) `request:"post /create"`
	List   func(*gin.Context) `request:"get /list"`
}

func (u User) NewController() router.Controller {
	u.Create = create
	u.List = getList
	return u
}

func create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "create user")
}

func getList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get user list")
}
