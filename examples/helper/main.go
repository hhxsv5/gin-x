package main

import (
	"log"
	"github.com/hhxsv5/gin-x/helper"
)

func main() {

	// Md5
	m := helper.Md5("123")
	log.Println(m)

	// Random string
	r := helper.RandStr(12)
	log.Println(r)

	type T struct {
		A string
		B int
	}

	// Json2String
	s := T{"a", 123}
	ss, err := helper.Json2String(s)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(ss)
	}

	// String2Json
	sss := T{}
	err = helper.String2Json(ss, &sss)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(sss)
	}

	// Check file exist
	e := helper.FileExists("./main.go")
	log.Println(e)

	// Check is dir
	id := helper.FileIsDir("./main.go")
	log.Println(id)

	// Pager usage
	//ctx: *gin.Context
	//start, limit := helper.ParsePager(ctx)
	//list := make([]string, 10)
	//total := 10
	//pager := helper.BuildPager(list, total)
	//ctx.JSON(http.StatusOK, pager)

	// Time
	t1 := helper.TodayStart()
	log.Println(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond())
	t2 := helper.TodayEnd()
	log.Println(t2.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute(), t2.Second(), t2.Nanosecond())

	// Shell
	str, err := helper.ExecShell("php -r 'echo time();'")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("now time", str)
	}
}
