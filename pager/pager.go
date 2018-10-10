package pager

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	MaxPageSize     = 100
	DefaultPageSize = 20
)

func ParsePager(ctx *gin.Context) (start, limit uint64) {
	page, err := strconv.ParseUint(ctx.DefaultQuery("page", ctx.DefaultPostForm("page", "1")), 10, 64)
	if err != nil {
		page = 1
	}
	if page < 1 {
		page = 1
	}

	limit, err = strconv.ParseUint(ctx.DefaultQuery("page_size", ctx.DefaultPostForm("page_size", strconv.FormatUint(DefaultPageSize, 10))), 10, 64)
	if err != nil {
		limit = DefaultPageSize
	}
	if limit <= 0 {
		limit = DefaultPageSize
	} else if limit > MaxPageSize {
		limit = MaxPageSize
	}

	start = (page - 1) * limit
	return
}

func BuildPager(list interface{}, total interface{}) (pager gin.H) {
	return gin.H{"list": list, "total": total}
}
