/**
 * @Author: Anpw
 * @Description:
 * @File:  plugin
 * @Version: 1.0.0
 * @Date: 2021/6/8 23:50
 */

package dao

import (
	"bugu/internal/model"
	"bugu/pkg/app"
	"errors"
	"github.com/jinzhu/gorm"
)

func (d *Dao) CreatePlugin(pluginName, pluginIntroduction, fileMaskCode, chipName string) error {
	c := model.Chip{
		ChipName: chipName,
	}
	chip, err := c.GetByChipName(d.engine)
	if err != nil {
		return err
	}
	if chip.ID == 0 {
		return errors.New("不存在该芯片")
	}
	p := model.Plugin{
		PluginName:         pluginName,
		PluginIntroduction: pluginIntroduction,
		FileMaskCode:       fileMaskCode,
		ChipID:             chip.ID,
	}
	plugin, err := p.GetByPluginName(d.engine)
	if err != nil {
		return err
	}
	if plugin.ID != 0 {
		return errors.New("插件已存在")
	}

	return p.Create(d.engine)
}

func (d *Dao) UpdatePlugin(id uint, pluginName, pluginIntroduction, fileMaskCode, chipName string) error {
	p := model.Plugin{
		PluginName: pluginName,
	}
	plugin, err := p.GetByPluginName(d.engine)
	if err != nil {
		return err
	}
	if plugin.ID != 0 {
		return errors.New("插件已存在")
	}

	c := model.Chip{
		ChipName: chipName,
	}
	chip, err := c.GetByChipName(d.engine)
	if err != nil {
		return err
	}
	if chip.ID == 0 {
		return errors.New("不存在该芯片")
	}
	plugin = model.Plugin{
		Model: &gorm.Model{ID: id},
	}

	values := map[string]interface{}{
		"plugin_name":         pluginName,
		"plugin_introduction": pluginIntroduction,
		"file_mask_code":      fileMaskCode,
		"chip_id":             chip.ID,
	}
	return plugin.Update(d.engine, values)
}

func (d *Dao) DeletePlugin(id uint) error {
	plugin := model.Plugin{Model: &gorm.Model{ID: id}}
	return plugin.Delete(d.engine)
}

func (d *Dao) GetPlugin(id uint) (model.Plugin, error) {
	plugin := model.Plugin{Model: &gorm.Model{ID: id}}
	return plugin.Get(d.engine)
}

func (d *Dao) GetPluginList(chipID uint, page, pageSize int) ([]*model.Plugin, error) {
	plugin := model.Plugin{ChipID: chipID}
	pageOffset := app.GetPageOffset(page, pageSize)
	return plugin.GetPluginList(d.engine, pageOffset, pageSize)
}

func (d *Dao) CountPlugin(ChipID uint) (int, error) {
	plugin := model.Plugin{ChipID: ChipID}
	return plugin.Count(d.engine)
}