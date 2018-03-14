package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-slim-router"
)

type Test struct {
	// HttpMethod: Any/GET/POST/DELETE/PATCH/PUT/OPTIONS/HEAD
	// Format: {HttpMethod}Action func(*gin.Context) `path:"/xxxpath"`
	PostXxx func(*gin.Context) `path:"/xxx"`
	GetYyy  func(*gin.Context) `path:"/yyy"`
}

func (t Test) NewController() gsr.Controller {
	t.PostXxx = xxx
	t.GetYyy = yyy
	return t
}
func xxx(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "change xxx")
}
func yyy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get yyy")
}
