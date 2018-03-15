package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/router"
)

type Test struct {
	// HttpMethod: Any/GET/POST/DELETE/PATCH/PUT/OPTIONS/HEAD
	// Format: {HttpMethod}Action func(*gin.Context) `path:"/xxxpath"`
	PostXxx func(*gin.Context) `path:"/xxx"`
	GetYyy  func(*gin.Context) `path:"/yyy"`
	Any     func(*gin.Context) `path:"/any"`
}

func (t Test) NewController() router.Controller {
	t.PostXxx = xxx
	t.GetYyy = yyy
	t.Any = any
	return t
}

func xxx(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "change xxx")
}

func yyy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get yyy")
}

func any(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "any method")
}
