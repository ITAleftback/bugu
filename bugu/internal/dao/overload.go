/**
 * @Author: Anpw
 * @Description:
 * @File:  overload
 * @Version: 1.0.0
 * @Date: 2021/6/4 14:30
 */

package dao

import (
	"bugu/internal/model"
	"bugu/pkg/app"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
)

type Overload struct {
	*gorm.Model
	OverloadIntroduction string `json:"overload_introduction"`
	OverloadMaskCode     string `json:"overload_mask_code"`
	ModuleID             uint   `json:"module_id"`
}

func (d *Dao) CreateOverload(param *Overload, moduleName, jsonOverloadInput, jsonOverloadOutput string) error {
	m := model.Module{
		ModuleName: moduleName,
	}
	module, err := m.GetByModuleName(d.engine)
	if err != nil {
		return err
	}
	if module.ID == 0 {
		return errors.New("不存在该模块")
	}
	o := model.Overload{
		OverloadIntroduction: param.OverloadIntroduction,
		OverloadMaskCode:     param.OverloadMaskCode,
		ModuleID:             module.ID,
	}
	if err = o.Create(d.engine); err != nil {
		return err
	}
	//	jsonStr := `[
	//{"oi_introduction": "test", "oi_type": "test"},
	//{"oi_introduction": "test2", "oi_type": "test2"}
	//]`
	var overloadInput []model.OverloadInput
	err = json.Unmarshal([]byte(jsonOverloadInput), &overloadInput)
	if err != nil {
		return err
	}
	var overloadInputs model.OverloadInput
	for i := 0; i < len(overloadInput); i++ {
		overloadInputs = model.OverloadInput{
			OIIntroduction: overloadInput[i].OIIntroduction,
			OIType:         overloadInput[i].OIType,
			OverloadID:     o.ID,
		}

		err = overloadInputs.Create(d.engine)
		if err != nil {
			return err
		}
	}
	//overloadOutput
	var overloadOutput []model.OverloadOutput
	err = json.Unmarshal([]byte(jsonOverloadOutput), &overloadOutput)
	if err != nil {
		return err
	}
	var overloadOutputs model.OverloadOutput
	for i := 0; i < len(overloadOutput); i++ {
		overloadOutputs = model.OverloadOutput{
			OOIntroduction: overloadOutput[i].OOIntroduction,
			OOType:         overloadOutput[i].OOType,
			OverloadID:     o.ID,
		}
		err = overloadOutputs.Create(d.engine)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dao) UpdateOverload(param *Overload, moduleName string) error {
	m := model.Module{
		ModuleName: moduleName,
	}
	module, err := m.GetByModuleName(d.engine)
	if err != nil {
		return err
	}
	if module.ID == 0 {
		return errors.New("不存在该模块")
	}
	overload := model.Overload{
		Model: &gorm.Model{ID: param.ID},
	}

	values := map[string]interface{}{
		"overload_introduction": param.OverloadIntroduction,
		"overload_mask_code":    param.OverloadMaskCode,
		"module_id":             module.ID,
	}
	return overload.Update(d.engine, values)
}

func (d *Dao) DeleteOverload(id uint) error {
	overload := model.Overload{Model: &gorm.Model{ID: id}}
	return overload.Delete(d.engine)
}

func (d *Dao) GetOverload(id uint) (model.Overload, error) {
	overload := model.Overload{Model: &gorm.Model{ID: id}}
	return overload.Get(d.engine)
}

func (d *Dao) GetOverloadList(moduleID uint, page, pageSize int) ([]*model.Overload, error) {
	overload := model.Overload{ModuleID: moduleID}
	pageOffset := app.GetPageOffset(page, pageSize)
	return overload.GetOverloadList(d.engine, pageOffset, pageSize)
}

func (d *Dao) CountOverload(moduleID uint) (int, error) {
	overload := model.Overload{ModuleID: moduleID}
	return overload.Count(d.engine)
}
