/**
 * @Author: Anpw
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/5/28 19:30
 */

package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*gorm.Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "bugu_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
