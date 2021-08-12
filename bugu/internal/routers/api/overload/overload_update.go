/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_update
 * @Version: 1.0.0
 * @Date: 2021/6/4 15:15
 */

package overload

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/convert"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (o Overload) Update(ctx *gin.Context) {
	param := service.UpdateOverloadRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.UpdateOverload(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateOverload err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateOverloadFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新模块成功"})
	return
}