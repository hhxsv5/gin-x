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
