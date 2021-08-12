/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_input
 * @Version: 1.0.0
 * @Date: 2021/6/13 1:55
 */

package dao

import (
	"bugu/internal/model"
	"github.com/jinzhu/gorm"
)

func (d *Dao) DeleteOverloadInput(id uint) error {
	oi := model.OverloadInput{Model: &gorm.Model{ID: id}}
	return oi.Delete(d.engine)
}

func (d *Dao) GetOverloadInputList(overloadID uint) ([]*model.OverloadInput, error) {
	oi := model.OverloadInput{OverloadID: overloadID}
	return oi.GetOverloadInputList(d.engine)
}