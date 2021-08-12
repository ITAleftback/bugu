/**
 * @Author: Anpw
 * @Description:
 * @File:  chip_delete
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

func (c Chip) Delete(ctx *gin.Context) {
	param := service.DeleteChipRequest{ID:convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteChip(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteChip err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteChipFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "删除芯片成功"})
	return
}
