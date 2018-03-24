package helper

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

var (
	MaxLimit     = uint64(100)
	DefaultLimit = uint64(10)
)

func ParsePager(ctx *gin.Context) (start, limit uint64) {
	is, err := strconv.ParseUint(ctx.DefaultQuery("start", ctx.DefaultPostForm("start", "0")), 10, 32)
	if err == nil {
		start = uint64(is)
	} else {
		start = 0
	}
	il, err := strconv.ParseUint(ctx.DefaultQuery("limit", ctx.DefaultPostForm("limit", strconv.FormatUint(DefaultLimit, 10))), 10, 32)
	if err == nil {
		limit = uint64(il)
	} else {
		limit = DefaultLimit
	}

	if start < 0 {
		start = 0
	}
	if limit <= 0 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	return start, limit
}

func BuildPager(list interface{}, total interface{}) map[string]interface{} {
	pager := map[string]interface{}{}
	pager["list"] = list
	pager["total"] = total
	return pager
}
