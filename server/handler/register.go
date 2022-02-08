package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/inits"
	"server/model"
	"server/public"
	"server/utils"
)

// RegisterHandler 注册请求处理
func RegisterHandler(c *gin.Context)  {
	user:=new(model.User)
	response:=model.NewResponse(c)
	if err:=c.ShouldBindJSON(&user);err!=nil{
		log.Println(err.Error())
		response.SendError(public.ERR_BAD_DATA)
		return
	}
	user.Password=utils.HashAndSalt([]byte(user.Password))  //对密码进行哈希处理
	err := inits.UserData.AddUser(user)
	if err!=nil{
		log.Println(err.Error())
		response.SendError(public.ERR_USERNAME_REPEAT)
		return
	}
	resData:=make(map[string]interface{})
	resData["user_name"]=user.UserName
	resData["name"]=user.Name
	resData["email"]=user.Email
	response.SendData(resData)
}
