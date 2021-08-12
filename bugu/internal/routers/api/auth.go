/**
 * @Author: Anpw
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/5/29 16:59
 */

package api

//
//func GetAuth(c *gin.Context) {
//	param := service.AuthRequest{}
//	response := app.NewResponse(c)
//	valid, errs := app.BindAndValid(c, &param)
//	if valid == true {
//		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
//		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
//		return
//	}
//
//	svc := service.New(c.Request.Context())
//	err := svc.CheckAuth(&param)
//	if err != nil {
//		global.Logger.Errorf("svc.CheckAuth err: %v", err)
//		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
//		return
//	}
//
//	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
//	if err != nil {
//		global.Logger.Errorf("svc.GenerateToken err: %v", err)
//		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
//		return
//	}
//
//	response.ToResponse(gin.H{"token": token})
//
//}
