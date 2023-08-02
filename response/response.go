package response

import (
	"github.com/gin-gonic/gin"
	"github.com/stark-sim/serializer/code"
	"google.golang.org/grpc/status"
	"net/http"
)

var EmptyData struct{}

type RespData struct {
	Code code.MyCode `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PaginateResp struct {
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
	Total     int64       `json:"total"`
	List      interface{} `json:"list"`
}

func RespSuccess(ctx *gin.Context, data interface{}) {
	rd := &RespData{
		Code: code.Success,
		Msg:  code.Success.Msg(),
		Data: data,
	}
	ctx.JSON(http.StatusOK, rd)
}

// RespSuccessPagination 分页数据返回
func RespSuccessPagination(ctx *gin.Context, pageIndex, pageSize int, total int64, data interface{}) {
	respData := PaginateResp{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Total:     total,
		List:      data,
	}
	RespSuccess(ctx, respData)
}

func RespSuccessWithMsg(ctx *gin.Context, data interface{}, msg string) {
	rd := &RespData{
		Code: code.Success,
		Msg:  msg,
		Data: data,
	}
	ctx.JSON(http.StatusOK, rd)
}

func RespError(ctx *gin.Context, c code.MyCode) {
	rd := &RespData{
		Code: c,
		Msg:  c.Msg(),
		Data: EmptyData,
	}
	ctx.JSON(http.StatusOK, rd)
}

// RespErrorInvalidParams 参数校验不通过(gin should bind)
func RespErrorInvalidParams(ctx *gin.Context, err error) {
	msg := validError(err)
	rd := &RespData{
		Code: code.InvalidParams,
		Msg:  msg,
		Data: EmptyData,
	}
	ctx.JSON(http.StatusOK, rd)
}

func RespErrorWithMsg(ctx *gin.Context, code code.MyCode, errMsg string) {
	rd := &RespData{
		Code: code,
		Msg:  errMsg,
		Data: EmptyData,
	}
	ctx.JSON(http.StatusOK, rd)
}

func RespGrpcErrorWithMsg(ctx *gin.Context, code code.MyCode, err error) {
	ret := status.Convert(err)
	rd := &RespData{
		Code: code,
		Msg:  ret.Message(),
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}
