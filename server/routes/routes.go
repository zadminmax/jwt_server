package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/handler"
	"server/utils"
)

// SetupRoutes 注册路由
func SetupRoutes(app *gin.Engine) {
	app.Use(Cors())
	v1 := app.Group("/api")
	{
		v1.POST("/login", handler.LoginHandler)
		v1.POST("/register", handler.RegisterHandler)
	}
	v2 := app.Group("/data")
	v2.Use(utils.VerifyJWT())
	{
		v2.GET("/user", handler.UserHandler)
		v2.GET("/logout", handler.LogoutHandler)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header["Origin"]
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Origin", origin[0])
		c.Header("Access-Control-Allow-Headers", "Content-Type,token,*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT,PATCH,HEAD,TRACE,*")
		c.Header("Access-Control-Expose-Headers", "Content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "3600")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
