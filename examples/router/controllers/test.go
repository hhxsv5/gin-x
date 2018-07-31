package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/router"
)

type Test struct {
	// HttpMethod: Any/GET/POST/DELETE/PATCH/PUT/OPTIONS/HEAD
	// Format: Action func(*gin.Context) `request:"{HttpMethod} /xxxpath"`
	Xxx func(*gin.Context) `request:"POST /xxx"`
	Yyy func(*gin.Context) `request:"GET /yyy"`
	Any func(*gin.Context) `request:"ANY /any"`
}

func (t Test) NewController() router.Controller {
	t.Xxx = xxx
	t.Yyy = yyy
	t.Any = any
	return t
}

func xxx(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "post /xxx")
}

func yyy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get /yyy")
}

func any(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "any /any")
}
