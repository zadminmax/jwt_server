package handler

import (
	"github.com/gin-gonic/gin"
	"server/inits"
	"server/model"
	"server/public"
	"server/utils"
)

func UserHandler(c *gin.Context) {
	claims := c.Value("claims").(*utils.UserClaims)
	user := inits.UserData.GetUserById(claims.UserId)
	response := model.NewResponse(c)
	if user == nil {
		response.SendError(public.ERR_USER_NOT_EXIST)
		return
	}
	resData := make(map[string]interface{})
	resData["user_name"] = user.UserName
	resData["name"] = user.Name
	resData["email"] = user.Email
	response.SendData(resData)
}
