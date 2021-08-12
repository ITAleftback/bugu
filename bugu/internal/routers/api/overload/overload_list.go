/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_list
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

func (o Overload) List(ctx *gin.Context) {
	param := service.OverloadListRequest{ModuleID:convert.StrTo(ctx.Param("module_id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	totalRows, err := svc.CountOverload(&service.CountOverloadRequest{
		ModuleID: param.ModuleID,
	})
	overloads, err := svc.GetOverloadList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetOverload err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOverloadsFail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(overloads, totalRows)
	return
}


