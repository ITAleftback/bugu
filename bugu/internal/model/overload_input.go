/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_input
 * @Version: 1.0.0
 * @Date: 2021/6/9 21:34
 */

package model

import "github.com/jinzhu/gorm"

type OverloadInput struct {
	*gorm.Model
	OIIntroduction string `json:"oi_introduction"`
	OIType         string `json:"oi_type"`
	OverloadID     uint   `json:"overload_id"`
}

func (oi OverloadInput) TableName() string {
	return "bugu_overload_input"
}

func (oi OverloadInput) Create(db *gorm.DB) error {
	return db.Create(&oi).Error
}

func (oi OverloadInput) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&oi).Updates(values).Where("id = ?", oi.ID).Error
}

func (oi OverloadInput) Get(db *gorm.DB) (OverloadInput, error) {
	var overloadInput OverloadInput
	db = db.Where("id = ?", oi.ID)
	err := db.First(&overloadInput).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return overloadInput, err
	}
	return overloadInput, nil
}

func (oi OverloadInput) GetOverloadInputList(db *gorm.DB) ([]*OverloadInput, error) {
	var overloadInputs []*OverloadInput
	var err error
	if oi.OverloadID != 0 {
		db = db.Where("overload_id = ?", oi.OverloadID)
	}
	if err = db.Find(&overloadInputs).Error; err != nil {
		return nil, err
	}
	return overloadInputs, nil
}

func (oi OverloadInput) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", oi.Model.ID).Delete(&oi).Error; err != nil {
		return err
	}
	return nil
}
