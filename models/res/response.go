package res

import (
	CODE "GalleryMing/models/res/code"
	"GalleryMing/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response json格式消息封装
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	SUCCESS = 0
	ERR     = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func ResultOkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

// ResultFail 失败响应
func ResultFail(data any, msg string, c *gin.Context) {
	Result(ERR, data, msg, c)
}

func ResultFailWithMsg(msg string, c *gin.Context) {
	Result(ERR, map[string]any{}, msg, c)
}

func ResultFailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	ResultFailWithMsg(msg, c)
}

func ResultFailWithCode(code CODE.ErrorCode, c *gin.Context) {
	if msg, ok := CODE.ErrorCodeMap[code]; ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERR, map[string]any{}, "未知错误", c)
}
