/**
 * @Author: Anpw
 * @Description:
 * @File:  module_delete
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

func (m Module) Delete(ctx *gin.Context) {
	param := service.DeleteModuleRequest{ID: convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteModule(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteModule err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteModuleFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "删除模块成功"})
	return
}
