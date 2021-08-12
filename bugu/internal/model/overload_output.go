/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_output
 * @Version: 1.0.0
 * @Date: 2021/6/9 21:34
 */

package model

import "github.com/jinzhu/gorm"

type OverloadOutput struct {
	*gorm.Model
	OOIntroduction string `json:"oo_introduction"`
	OOType         string `json:"oo_type"`
	OverloadID     uint   `json:"overload_id"`
}

func (oo OverloadOutput) TableName() string {
	return "bugu_overload_output"
}

func (oo OverloadOutput) Create(db *gorm.DB) error {
	return db.Create(&oo).Error
}

func (oo OverloadOutput) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&oo).Updates(values).Where("id = ?", oo.ID).Error
}

func (oo OverloadOutput) Get(db *gorm.DB) (OverloadOutput, error) {
	var overloadOutput OverloadOutput
	db = db.Where("id = ?", oo.ID)
	err := db.First(&overloadOutput).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return overloadOutput, err
	}
	return overloadOutput, nil
}

func (oo OverloadOutput) GetOverloadOutputList(db *gorm.DB) ([]*OverloadOutput, error) {
	var overloadOutputs []*OverloadOutput
	var err error
	if oo.OverloadID != 0 {
		db = db.Where("overload_id = ?", oo.OverloadID)
	}
	if err = db.Find(&overloadOutputs).Error; err != nil {
		return nil, err
	}
	return overloadOutputs, nil
}

func (oo OverloadOutput) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", oo.Model.ID).Delete(&oo).Error; err != nil {
		return err
	}
	return nil
}
