package global_return

import (
	"github.com/gin-gonic/gin"
	"github.com/stark-sim/serializer/code_msgs"
	"net/http"
)

var EmptyData struct{}

type ResponseData struct {
	*code_msgs.CodeMsg
	Data interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, c *code_msgs.CodeMsg) {
	rd := &ResponseData{
		CodeMsg: c,
		Data:    EmptyData,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithCustomMsg(ctx *gin.Context, c *code_msgs.CodeMsg, errMsg string) {
	if errMsg != "" {
		c.Msg = errMsg
	}
	rd := &ResponseData{
		CodeMsg: c,
		Data:    EmptyData,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, codeMsg *code_msgs.CodeMsg, data interface{}) {
	rd := &ResponseData{
		CodeMsg: codeMsg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
