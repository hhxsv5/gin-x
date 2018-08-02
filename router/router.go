package router

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	RequestTag = "request"

	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodHead    = "HEAD"
	MethodPatch   = "PATCH"
	MethodOptions = "OPTIONS"
	MethodAny     = "ANY"
)

var (
	actionType     = reflect.TypeOf(func(*gin.Context) {})
	supportMethods = map[string]func(gin.IRouter, string, func(*gin.Context)){
		MethodGet:     func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.GET(path, handler) },
		MethodPost:    func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.POST(path, handler) },
		MethodPut:     func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.PUT(path, handler) },
		MethodDelete:  func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.DELETE(path, handler) },
		MethodHead:    func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.HEAD(path, handler) },
		MethodPatch:   func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.PATCH(path, handler) },
		MethodOptions: func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.OPTIONS(path, handler) },
		MethodAny:     func(rg gin.IRouter, path string, handler func(*gin.Context)) { rg.Any(path, handler) },
	}
	parseAction = func(rg gin.IRouter, field reflect.StructField, value reflect.Value) bool {
		if field.Type != actionType {
			return false
		}
		request := strings.Trim(field.Tag.Get(RequestTag), " ")
		if request == "-" {
			return false
		}
		s := strings.SplitN(request, " ", 2)
		method := strings.Trim(s[0], " ")
		path := ""
		if len(s) >= 2 {
			path = strings.Trim(s[1], " ")
			if path == "/" {
				path = ""
			}
		}
		return guessMethod(rg, method, path, value.Interface().(func(*gin.Context)))
	}
	guessMethod = func(rg gin.IRouter, method string, path string, handler func(*gin.Context)) bool {
		method = strings.ToUpper(method)
		if call, ok := supportMethods[method]; ok {
			call(rg, path, handler)
			return true
		}
		return false
	}
)

type Controller interface {
	NewController() Controller
}

type SlimRouter struct {
	engine *gin.Engine
}

func NewSlimRouter(engine *gin.Engine) *SlimRouter {
	return &SlimRouter{engine}
}

func (sr *SlimRouter) Use(middleware ...gin.HandlerFunc) {
	sr.engine.Use(middleware...)
}

func (sr *SlimRouter) RegisterGroup(group string, middleware ...gin.HandlerFunc) *routerGroup {
	rg := sr.engine.Group(group)
	rg.Use(middleware...)
	return &routerGroup{rg: rg}
}

func (sr *SlimRouter) RegisterController(ctrl ...Controller) {
	for _, tc := range ctrl {
		t := reflect.TypeOf(tc)
		v := reflect.ValueOf(tc)
		fc := t.NumField()
		for i := 0; i < fc; i++ {
			parseAction(sr.engine, t.Field(i), v.Field(i))
		}
	}
}

type routerGroup struct {
	rg *gin.RouterGroup
}

func (g *routerGroup) RegisterController(ctrl ...Controller) {
	for _, tc := range ctrl {
		t := reflect.TypeOf(tc)
		v := reflect.ValueOf(tc)
		fc := t.NumField()
		for i := 0; i < fc; i++ {
			parseAction(g.rg, t.Field(i), v.Field(i))
		}
	}
}

func (g *routerGroup) RegisterGroup(group string, middleware ...gin.HandlerFunc) *routerGroup {
	rg := g.rg.Group(group)
	rg.Use(middleware...)
	return &routerGroup{rg}
}
