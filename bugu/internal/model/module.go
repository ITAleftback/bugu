/**
 * @Author: Anpw
 * @Description:
 * @File:  module
 * @Version: 1.0.0
 * @Date: 2021/5/26 16:43
 */

package model

import (
	"github.com/jinzhu/gorm"
)

type Module struct {
	*gorm.Model
	ModuleName         string `json:"module_name"`
	ModuleIntroduction string `json:"module_introduction"`
	PluginID           uint   `json:"plugin_id"`
}

func (m Module) TableName() string {
	return "bugu_module"
}

func (m Module) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

func (m Module) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&m).Updates(values).Where("id = ?", m.ID).Error
}

func (m Module) Get(db *gorm.DB) (Module, error) {
	var module Module
	db = db.Where("id = ?", m.ID)
	err := db.First(&module).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return module, err
	}
	return module, nil
}

func (m Module) GetModuleList(db *gorm.DB, pageOffset, pageSize int) ([]*Module, error) {
	var modules []*Module
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if m.PluginID != 0 {
		db = db.Where("plugin_id = ?", m.PluginID)
	}
	if err = db.Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}

func (m Module) GetByModuleName(db *gorm.DB) (Module, error) {
	var module Module
	db = db.Where("module_name = ?", m.ModuleName)
	err := db.First(&module).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return module, err
	}
	return module, nil
}

func (m Module) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", m.Model.ID).Delete(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m Module) Count(db *gorm.DB) (int, error) {
	var count  int
	if m.PluginID != 0 {
		db = db.Where("plugin_id = ?", m.PluginID)
	}
	if err := db.Model(&m).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}