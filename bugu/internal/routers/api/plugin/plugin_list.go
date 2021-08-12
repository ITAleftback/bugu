/**
 * @Author: Anpw
 * @Description:
 * @File:  plugin_about_list
 * @Version: 1.0.0
 * @Date: 2021/6/3 20:02
 */

package plugin

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/convert"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (p Plugin) List(ctx *gin.Context) {
	param := service.PluginListRequest{ChipID:convert.StrTo(ctx.Param("chip_id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(ctx.Request.Context())
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	totalRows, err := svc.CountPlugin(&service.CountPluginRequest{
		ChipID: param.ChipID,
	})
	plugins, err := svc.GetPluginList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetPluginList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetPluginsFail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(plugins, totalRows)
	return
}