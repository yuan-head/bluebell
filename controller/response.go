package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 10001, // 程序中的错误码
	"msg": xx,     // 提示信息
    "data": {},    // 响应数据
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"` // 错误码
	Msg  interface{} `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 响应数据
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
