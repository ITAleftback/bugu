/**
 * @Author: Anpw
 * @Description:
 * @File:  overload_output
 * @Version: 1.0.0
 * @Date: 2021/6/13 2:25
 */

package dao

import (
	"bugu/internal/model"
	"github.com/jinzhu/gorm"
)

func (d *Dao) DeleteOverloadOutput(id uint) error {
	oo := model.OverloadOutput{Model: &gorm.Model{ID: id}}
	return oo.Delete(d.engine)
}

func (d *Dao) GetOverloadOutputList(overloadID uint) ([]*model.OverloadOutput, error) {
	oo := model.OverloadOutput{OverloadID: overloadID}
	return oo.GetOverloadOutputList(d.engine)
}
