/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_create
 * @Version: 1.0.0
 * @Date: 2021/6/4 15:15
 */

package overload

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (o Overload) Create(ctx *gin.Context) {
	param := service.CreateOverloadRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.CreateOverload(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateOverload err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateOverloadFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建重载成功"})
	return
}
