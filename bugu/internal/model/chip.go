/**
 * @Author: Anpw
 * @Description:
 * @File:  chip
 * @Version: 1.0.0
 * @Date: 2021/5/26 16:21
 */

package model

import (
	"github.com/jinzhu/gorm"
)

type Chip struct {
	*gorm.Model
	ChipName         string `json:"chip_name"`
	ManufacturerName string `json:"manufacturer_name"`
	Type             string `json:"type"`
}

func (c Chip) TableName() string {
	return "bugu_chip"
}

func (c Chip) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c Chip) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&c).Updates(values).Where("id = ?", c.ID).Error
}

func (c Chip) Get(db *gorm.DB) (Chip, error) {
	var chip Chip
	db = db.Where("id = ?", c.ID)
	err := db.First(&chip).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return chip, err
	}
	return chip, nil
}

func (c Chip) GetChipList(db *gorm.DB, pageOffset, pageSize int) ([]*Chip, error) {
	var chips []*Chip
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if c.ChipName != "" {
		db = db.Where("chip_name = ?", c.ChipName)
	}
	if err = db.Find(&chips).Error; err != nil {
		return nil, err
	}
	return chips, nil
}

func (c Chip) GetByChipName(db *gorm.DB) (Chip, error) {
	var chip Chip
	db = db.Where("chip_name = ?", c.ChipName)
	err := db.First(&chip).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return chip, err
	}
	return chip, nil
}

func (c Chip) Delete(db *gorm.DB) error {
	if err := db.Where("id=?", c.ID).Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c Chip) Count(db *gorm.DB) (int, error) {
	var count  int
	if c.ChipName != "" {
		db = db.Where("chip_name = ?", c.ChipName)
	}
	if err := db.Model(&c).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}