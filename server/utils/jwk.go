package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"server/model"
	"server/public"
	"time"
)

// VerifyJWT 验证JWT
func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		response := model.NewResponse(c)
		if token == "" {
			response.SendError(public.ERR_TOKEN_EMPTY)
			c.Abort()
			return
		}
		j := NewJwt()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.SendError(public.ERR_TOKEN_EXPIRED)
				c.Abort()
				return
			}
			log.Println(err.Error())
			response.SendError(public.ERR_TOKEN_NOT_HANDLE)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token")
)

type UserClaims struct {
	UserId       uint `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func NewJwt() *JWT {
	return &JWT{
		[]byte(public.TOKEN_SIGIN_KEY),
	}
}

func (j *JWT) CreateToken(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(tokenStr string,hours int) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(hours) * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
