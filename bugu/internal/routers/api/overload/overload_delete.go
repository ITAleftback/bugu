/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_delete
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

func (o Overload) Delete(ctx *gin.Context) {
	param := service.DeleteOverloadRequest{ID: convert.StrTo(ctx.Param("id")).MustUint()}
	response := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(ctx.Request.Context())
	err := svc.DeleteOverload(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteOverload err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteOverloadFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"msg": "删除重载成功"})
	return
}