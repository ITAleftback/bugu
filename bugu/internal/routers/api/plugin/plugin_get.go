/**
 * @Author: Anpw
 * @Description:
 * @File:  plugin_about_get
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

func (p Plugin) Get(ctx *gin.Context) {
	param := service.PluginRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	plugin, err := svc.GetPlugin(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetPlugin err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetPluginFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(plugin)
	return
}
