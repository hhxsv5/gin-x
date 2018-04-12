package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("See helper tests")

	qs := "{\\x22data\\x22:{\\x22id\\x22:16,\\x22title\\x22:\\x22\\x5Cu6dfb\\x5Cu52a0\\x5Cu4e00\\x5Cu4e2a\\x22,\\x22online_time\\x22:1523448000,\\x22pv\\x22:0,\\x22desc\\x22:\\x22\\x5Cu6dfb\\x5Cu52a0\\x5Cu4e00\\x5Cu4e2a\\x22,\\x22cover\\x22:\\x22http:\\x5C/\\x5C/7xpdel.com1.z0.glb.clouddn.com\\x5C/Fl1s7j3Md71MpIL-Gib10UDAElPq\\x22,\\x22type\\x22:3,\\x22isFeather\\x22:1,\\x22tsCount\\x22:2},\\x22errcode\\x22:0,\\x22errmsg\\x22:\\x22success\\x22}"
	s, err := strconv.Unquote(fmt.Sprintf(`"%s"`, qs))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
