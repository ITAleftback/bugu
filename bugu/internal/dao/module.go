/**
 * @Author: Anpw
 * @Description:
 * @File:  module
 * @Version: 1.0.0
 * @Date: 2021/6/3 22:04
 */

package dao

import (
	"bugu/internal/model"
	"bugu/pkg/app"
	"errors"
	"github.com/jinzhu/gorm"
)

func (d *Dao) CreateModule(moduleName, moduleIntroduction, pluginName string) error {
	pa := model.Plugin{
		PluginName: pluginName,
	}
	plugin, err := pa.GetByPluginName(d.engine)
	if err != nil {
		return err
	}
	if plugin.ID == 0 {
		return errors.New("不存在该插件")
	}
	m := model.Module{
		ModuleName:         moduleName,
		ModuleIntroduction: moduleIntroduction,
		PluginID:           plugin.Model.ID,
	}
	module, err := m.GetByModuleName(d.engine)
	if err != nil {
		return err
	}
	if module.ID != 0 {
		return errors.New("模块已存在")
	}

	return m.Create(d.engine)
}

func (d *Dao) UpdateModule(id uint, moduleName, moduleIntroduction, pluginName string) error {
	m := model.Module{
		ModuleName: moduleName,
	}
	module, err := m.GetByModuleName(d.engine)
	if err != nil {
		return err
	}
	if module.ID != 0 {
		return errors.New("模块已存在")
	}
	p := model.Plugin{
		PluginName: pluginName,
	}
	plugin, err := p.GetByPluginName(d.engine)
	if err != nil {
		return err
	}
	if plugin.ID == 0 {
		return errors.New("不存在该插件")
	}

	module = model.Module{
		Model: &gorm.Model{ID: id},
	}

	values := map[string]interface{}{
		"module_name":         moduleName,
		"module_introduction": moduleIntroduction,
		"plugin_id":     plugin.ID,
	}
	return module.Update(d.engine, values)
}

func (d *Dao) DeleteModule(id uint) error {
	module := model.Module{Model: &gorm.Model{ID: id}}
	return module.Delete(d.engine)
}

func (d *Dao) GetModule(id uint) (model.Module, error) {
	module := model.Module{Model: &gorm.Model{ID: id}}
	return module.Get(d.engine)
}

func (d *Dao) GetModuleList(pluginID uint, page, pageSize int) ([]*model.Module, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	module := model.Module{PluginID: pluginID}
	return module.GetModuleList(d.engine, pageOffset, pageSize)
}

func (d *Dao) CountModule(pluginID uint) (int, error) {
	module := model.Module{PluginID: pluginID}
	return module.Count(d.engine)
}