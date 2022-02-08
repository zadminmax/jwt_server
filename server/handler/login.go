package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/inits"
	"server/model"
	"server/public"
	"server/utils"
	"time"
)

type Login struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResult struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

// LoginHandler 登录请求处理
func LoginHandler(c *gin.Context) {
	login := new(Login)
	response := model.NewResponse(c)
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Println(err.Error())
		response.SendError(public.ERR_BAD_DATA)
		return
	}
	user := inits.UserData.GetUserByUserName(login.UserName)
	if user == nil {
		response.SendError(public.ERR_USERNAME_NOT_EXIST)
		return
	}
	if !utils.ValidatePasswords(user.Password, []byte(login.Password)) {
		response.SendError(public.ERR_PASSWORD_NOT_MATCH)
		return
	}
	GenerateToken(c, *user)
}

func GenerateToken(c *gin.Context, user model.User) {
	j := utils.NewJwt()
	claims := utils.UserClaims{
		UserId:   user.Id,
		UserName: user.UserName,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(public.TOKEN_EXPIRED_HOURS * time.Hour).Unix(),
			Issuer:    public.TOKEN_SIGIN_KEY,
			NotBefore: time.Now().Unix() - 1000,
		},
	}
	response := model.NewResponse(c)
	token, err := j.CreateToken(claims)
	if err != nil {
		response.SendError(public.ERR_SERVER_PROCESS_MISS)
		return
	}
	res := &LoginResult{
		UserName: user.UserName,
		Token:    token,
	}
	//SetCookie(c,token)
	response.SendData(res)

}

func SetCookie(c *gin.Context, token string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().AddDate(0, 0, 3),
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
}
