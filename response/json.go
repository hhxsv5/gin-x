package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hhxsv5/gin-x/enums/codes"
)

type Response struct {
	Success   bool        `json:"success"`
	Code      string      `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func (r Response) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}

func SuccessJSON(c *gin.Context, data interface{}) {
	rsp := &Response{
		Success:   true,
		Code:      codes.Success,
		Msg:       codes.ErrorMap[codes.Success],
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      data,
	}
	c.JSON(http.StatusOK, rsp)
}

func FailJSON(c *gin.Context, code string, msg string) {
	if len(msg) == 0 {
		if t, ok := codes.ErrorMap[code]; ok {
			msg = t
		}
	}
	rsp := &Response{
		Success:   false,
		Code:      code,
		Msg:       msg,
		Timestamp: time.Now().UnixNano() / 1000000,
		Data:      nil,
	}
	c.JSON(http.StatusOK, rsp)
}
