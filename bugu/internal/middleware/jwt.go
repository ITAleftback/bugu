/**
 * @Author: Anpw
 * @Description:
 * @File:  jwt
 * @Version: 1.0.0
 * @Date: 2021/5/29 17:17
 */

package middleware

import (
	"bugu/global"
	"bugu/internal/model"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")
		}
		if token == "" || !strings.HasPrefix(token,"Bearer ") {
			ecode = errcode.InvalidParams
		} else {
			token = token[7:]
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}else {
				var user model.User
				global.DBEngine.Where("id=?", claims.UserID).First(&user)
				if user.ID == 0 {
					ecode = errcode.UnauthorizedAuthNotExist
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
