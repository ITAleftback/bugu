/**
 * @Author: Anpw
 * @Description:
 * @File:  module_get
 * @Version: 1.0.0
 * @Date: 2021/6/3 22:27
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

func (m Module) Get(ctx *gin.Context) {
	param := service.ModuleRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	module, err := svc.GetModule(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetModule err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetModuleFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(module)
	return
}
