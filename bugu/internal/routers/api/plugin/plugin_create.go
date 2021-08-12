/**
 * @Author: Anpw
 * @Description:
 * @File:  plugin_about_create
 * @Version: 1.0.0
 * @Date: 2021/6/3 20:01
 */

package plugin

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (p Plugin) Create(ctx *gin.Context) {
	param := service.CreatePluginRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.CreatePlugin(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreatePlugin err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreatePluginFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建插件成功"})
	return
}
