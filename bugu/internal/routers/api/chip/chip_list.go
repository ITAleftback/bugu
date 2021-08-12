/**
 * @Author: Anpw
 * @Description:
 * @File:  chip_list
 * @Version: 1.0.0
 * @Date: 2021/6/3 14:11
 */

package chip

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (c Chip) List(ctx *gin.Context) {
	param := service.ChipListRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	totalRows, err := svc.CountChip(&service.CountChipRequest{
		ChipName:  param.ChipName,
	})
	chips, err := svc.GetChipList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetChip err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetChipsFail.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(chips, totalRows)
	return
}
