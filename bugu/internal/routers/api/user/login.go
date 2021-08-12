/**
 * @Author: Anpw
 * @Description:
 * @File:  login
 * @Version: 1.0.0
 * @Date: 2021/6/2 5:14
 */

package user

import (
	"bugu/global"
	"bugu/internal/model"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (u User) Login(c *gin.Context) {
	param := service.LoginUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.LoginUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.LoginUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorLoginUserFail.WithDetails(err.Error()))
		return
	}
	/**
	 * TODO-Anpw: 2021/8/7 2:33 为了业务简单，我直接把user用来发放token，破坏了些架构
	 * Description:
	 */
	var user model.User
	global.DBEngine.Where("username=?", param.Username).First(&user)
	token, err := app.GenerateToken(user)
	if err != nil {
		global.Logger.Errorf("svc.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"msg": "登录成功", "token": token})
	return
}
