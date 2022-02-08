package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
)

func LogoutHandler(c *gin.Context) {
	ClearCookie(c)
	response:=model.NewResponse(c)
	response.SendData(nil)
}

func ClearCookie(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
}
