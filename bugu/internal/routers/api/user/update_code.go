/**
 * @Author: Anpw
 * @Description:
 * @File:  update_code
 * @Version: 1.0.0
 * @Date: 2021/6/2 5:14
 */

package user

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (u User) UpdateCode(c *gin.Context) {
	param := service.UpdatePwdRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateCode(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateCode err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateCodeFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{"msg": "修改密码成功"})
	return
}
