/**
 * @Author: Anpw
 * @Description:
 * @File:  overload
 * @Version: 1.0.0
 * @Date: 2021/5/26 16:44
 */

package model

import (
	"github.com/jinzhu/gorm"
)

type Overload struct {
	*gorm.Model
	OverloadIntroduction string `json:"overload_introduction"`
	OverloadMaskCode     string `json:"overload_mask_code"`
	ModuleID             uint `json:"module_id"`
}

func (o Overload) TableName() string {
	return "bugu_overload"
}

func (o *Overload) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

func (o Overload) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&o).Updates(values).Where("id = ?", o.ID).Error
}

func (o Overload) Get(db *gorm.DB) (Overload, error) {
	var overload Overload
	db = db.Where("id = ?", o.ID)
	err := db.First(&overload).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return overload, err
	}
	return overload, nil
}

func (o Overload) GetByModuleIntroduction(db *gorm.DB) (Overload, error) {
	var overload Overload
	db = db.Where("overload_introduction = ?", o.OverloadIntroduction)
	err := db.First(&overload).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return overload, err
	}
	return overload, nil
}

func (o Overload) GetOverloadList(db *gorm.DB, pageOffset, pageSize int) ([]*Overload, error) {
	var overloads []*Overload
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if o.ModuleID != 0 {
		db = db.Where("module_id = ?", o.ModuleID)
	}
	if err = db.Find(&overloads).Error; err != nil {
		return nil, err
	}
	return overloads, nil
}

func (o Overload) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", o.Model.ID).Delete(&o).Error; err != nil {
		return err
	}
	return nil
}

func (o Overload) Count(db *gorm.DB) (int, error) {
	var count  int
	if o.ModuleID != 0 {
		db = db.Where("module_id = ?", o.ModuleID)
	}
	if err := db.Model(&o).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}