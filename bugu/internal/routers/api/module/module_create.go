/**
 * @Author: Anpw
 * @Description:
 * @File:  module_create
 * @Version: 1.0.0
 * @Date: 2021/6/3 22:26
 */

package module

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (m Module) Create(ctx *gin.Context) {
	param := service.CreateModuleRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.CreateModule(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateModule err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateModuleFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建模块成功"})
	return
}
