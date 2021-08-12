/**
 * @Author: Anpw
 * @Description:
 * @File:  chip_create
 * @Version: 1.0.0
 * @Date: 2021/6/2 21:06
 */

package chip

import (
	"bugu/global"
	"bugu/internal/service"
	"bugu/pkg/app"
	"bugu/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func (c Chip) Create(ctx *gin.Context) {
	param := service.CreateChipRequest{}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.CreateChip(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateChip err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateChipFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "创建芯片成功"})
	return
}
