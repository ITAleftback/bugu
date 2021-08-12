/**
 * @Author: Anpw
 * @Description:
 * @File:  chip
 * @Version: 1.0.0
 * @Date: 2021/6/2 5:28
 */

package dao

import (
	"bugu/internal/model"
	"bugu/pkg/app"
	"errors"
	"github.com/jinzhu/gorm"
)

func (d *Dao) CreateChip(chipName, manufacturerName, types string) error {
	c := model.Chip{
		ChipName:         chipName,
		ManufacturerName: manufacturerName,
		Type:             types,
	}
	chip, err := c.GetByChipName(d.engine)
	if err != nil {
		return err
	}
	if chip.ID != 0 {
		return errors.New("芯片已存在")
	}
	return c.Create(d.engine)
}

func (d *Dao) UpdateChip(id uint, chipName, manufacturerName, types string) error {
	c := model.Chip{
		ChipName: chipName,
	}
	chip, err := c.GetByChipName(d.engine)
	if err != nil {
		return err
	}
	if chip.ID != 0 {
		return errors.New("芯片已存在")
	}
	chip = model.Chip{
		Model: &gorm.Model{ID: id},
	}
	values := map[string]interface{}{
		"chip_name":         chipName,
		"manufacturer_name": manufacturerName,
		"type":              types,
	}
	return chip.Update(d.engine, values)
}

func (d *Dao) DeleteChip(id uint) error {
	chip := model.Chip{Model: &gorm.Model{ID: id}}
	return chip.Delete(d.engine)
}

func (d *Dao) GetChip(id uint) (model.Chip, error) {
	chip := model.Chip{Model: &gorm.Model{ID: id}}
	return chip.Get(d.engine)
}

func (d *Dao) GetChipList(chipName string, page, pageSize int) ([]*model.Chip, error) {
	chip := model.Chip{ChipName: chipName}
	pageOffset := app.GetPageOffset(page, pageSize)
	return chip.GetChipList(d.engine, pageOffset, pageSize)
}

func (d *Dao) CountChip(chipName string) (int, error) {
	chip := model.Chip{ChipName: chipName}
	return chip.Count(d.engine)
}