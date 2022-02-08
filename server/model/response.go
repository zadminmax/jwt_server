package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/public"
)

// ResponseEntity 定义向客户端反馈数据统一格式
type ResponseEntity struct {
	ErrCode uint          `json:"err_code"`
	Msg     string       `json:"msg"`
	Data    interface{}  `json:"data"`
}

type Response struct {
	Context *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{Context: c}
}

// SendData 返回成功数据
func (m *Response)SendData(data interface{}) {
	entity:=new(ResponseEntity)
	entity.ErrCode=0
	entity.Msg="ok"
	entity.Data=data
	m.Context.JSON(http.StatusOK,entity)
}

// SendError 返回错误信息
func (m *Response)SendError(errCode uint) {
	entity:=new(ResponseEntity)
	entity.ErrCode=errCode
	if msg,ok:=public.ErrMsg[errCode];ok{
		entity.Msg=msg
	}else {
		entity.Msg="未知错误"
	}
	m.Context.JSON(http.StatusOK,entity)
}