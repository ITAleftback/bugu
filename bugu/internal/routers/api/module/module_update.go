/**
 * @Author: Anpw
 * @Description:
 * @File:  module_update
 * @Version: 1.0.0
 * @Date: 2021/6/3 22:26
 */

package module

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/convert"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (m Module) Update(ctx *gin.Context) {
	param := service.UpdateModuleRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.UpdateModule(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateModule err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateModuleFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新模块成功"})
	return
}