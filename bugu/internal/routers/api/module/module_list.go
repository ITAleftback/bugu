/**
 * @Author: Anpw
 * @Description:
 * @File:  module_list
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

func (m Module) List(ctx *gin.Context) {
	param := service.ModuleListRequest{PluginID:convert.StrTo(ctx.Param("plugin_id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	totalRows, err := svc.CountModule(&service.CountModuleRequest{
		PluginID: param.PluginID,
	})
	modules, err := svc.GetModuleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetModule err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetModulesFail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(modules, totalRows)
	return
}
