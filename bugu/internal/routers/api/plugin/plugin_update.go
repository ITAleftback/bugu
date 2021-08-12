/**
 * @Author: Anpw
 * @Description:
 * @File:  plugin_about_update
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

func (p Plugin) Update(ctx *gin.Context) {
	param := service.UpdatePluginRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.UpdatePlugin(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdatePlugin err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdatePluginFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新插件成功"})
	return
}