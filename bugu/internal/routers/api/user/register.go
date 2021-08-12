/**
 * @Author: Anpw
 * @Description:
 * @File:  register
 * @Version: 1.0.0
 * @Date: 2021/6/2 5:13
 */

package user

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (u User) Register(c *gin.Context) {
	param := service.CreateUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "注册成功"})
	return
}
