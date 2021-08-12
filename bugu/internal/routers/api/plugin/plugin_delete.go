/**
 * @Author: Anpw
 * @Description:
 * @File:  plugin_about_delete
 * @Version: 1.0.0
 * @Date: 2021/6/3 20:01
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

func (p Plugin) Delete(ctx *gin.Context) {
	param := service.DeletePluginRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeletePlugin(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeletePlugin err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeletePluginFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "删除插件成功"})
	return
}
