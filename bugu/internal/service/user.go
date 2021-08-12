/**
 * @Author: Anpw
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2021/5/26 21:07
 */

package service

type CreateUserRequest struct {
	Username     string `form:"username" binding:"required,min=3,max=10"`
	Password     string `form:"password" binding:"required,min=6,max=10"`
	SecurityCode string `form:"securityCode" binding:"required,min=3,max=6"`
	Admin        uint   `form:"admin,default=0" binding:"oneof=0 1"`
}

type LoginUserRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UpdatePwdRequest struct {
	Username     string `form:"username" binding:"required"`
	NewPwd       string `form:"newPwd" binding:"required,min=6,max=10"`
	SecurityCode string `form:"securityCode" binding:"required"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Username, param.Password, param.SecurityCode, param.Admin)
}

func (svc *Service) LoginUser(param *LoginUserRequest) error {
	return svc.dao.LoginUser(param.Username, param.Password)
}

func (svc *Service) UpdateCode(param *UpdatePwdRequest) error {
	return svc.dao.UpdateCode(param.Username, param.NewPwd, param.SecurityCode)
}
