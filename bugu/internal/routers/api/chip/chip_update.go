/**
 * @Author: Anpw
 * @Description:
 * @File:  chip_update
 * @Version: 1.0.0
 * @Date: 2021/6/2 21:07
 */

package chip

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/convert"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (c Chip) Update(ctx *gin.Context) {
	param := service.UpdateChipRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.UpdateChip(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateChip err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateChipFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "更新芯片成功"})
	return
}
