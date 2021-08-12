/**
 * @Author: Anpw
 * @Description:
 * @File:  chip_get
 * @Version: 1.0.0
 * @Date: 2021/6/3 14:11
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

func (c Chip) Get(ctx *gin.Context) {
	param := service.ChipRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	chip, err := svc.GetChip(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetChip err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetChipFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(chip)
	return
}
