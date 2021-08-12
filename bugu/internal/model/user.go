/**
 * @Author: Anpw
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2021/5/26 16:10
 */

package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	*gorm.Model
	Username     string `json:"username"`
	Password     string `json:"password"`
	SecurityCode string `json:"security_code"`
	Admin        uint `json:"admin"`
}

func (u User) TableName() string {
	return "bugu_user"
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&u).Updates(values).Where("id = ?", u.ID).Error
}

func (u User) Get(db *gorm.DB) (User, error) {
	var user User
	db = db.Where("id = ?", u.ID)
	err := db.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}

func (u User) GetByUsername(db *gorm.DB) (User, error) {
	var user User
	db = db.Where("username = ?", u.Username)
	err := db.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}
	return user, nil
}
