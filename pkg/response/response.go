package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    MyCode      `json:"code"` //把错误码都拉过来
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code MyCode) { //错误的，传入状态码，直接用写好的错误信息
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseErrorWithMsg(c *gin.Context, code MyCode, msg interface{}) { //错误的，传入状态码，且需要自定义错误信息的情况
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseSuccess(ctx *gin.Context, data interface{}) { //成功的，只有success一个状态码，要传入data：因为成功了，前端需要查看数据
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
