/**
 * @Author: Anpw
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2021/5/26 23:11
 */

package dao

import (
	"bugu/internal/model"
	"bugu/pkg/util"
	"errors"
)

func (d *Dao) CreateUser(username, password, securityCode string, admin uint) error {
	hashPass, err := util.GeneratePassHash(password)
	if err != nil {
		return err
	}
	u := model.User{
		Username:     username,
		Password:     hashPass,
		SecurityCode: securityCode,
		Admin:        admin,
	}

	user, err := u.GetByUsername(d.engine)
	if err != nil {
		return err
	}
	if user.ID != 0 {
		return errors.New("用户已存在")
	}
	return u.Create(d.engine)
}

func (d *Dao) LoginUser(username string, password string) error {
	u := model.User{
		Username: username,
	}
	user, err := u.GetByUsername(d.engine)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	flag := util.CompareHash(user.Password, password)
	if !flag {
		return errors.New("账号或密码错误")
	}
	return nil
}

func (d *Dao) UpdateCode(username, newPwd, securityCode string) error {
	u := model.User{
		Username: username,
	}
	user, err := u.GetByUsername(d.engine)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	if user.SecurityCode != securityCode {
		return errors.New("安全码错误")
	}
	hashPwd, err := util.GeneratePassHash(newPwd)
	if err != nil {
		return err
	}
	values := map[string]interface{}{
		"password": hashPwd,
	}
	return user.Update(d.engine, values)
}
