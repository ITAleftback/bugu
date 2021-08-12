/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_get
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

func (o Overload) Get(ctx *gin.Context) {
	param := service.OverloadRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	overload, err := svc.GetOverload(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetOverload err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOverloadFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(overload)
	return
}