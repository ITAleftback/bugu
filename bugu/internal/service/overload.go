/**
 * @Author: Anpw
 * @Description:
 * @File:  overload
 * @Version: 1.0.0
 * @Date: 2021/6/4 14:55
 */

package service

import (
	"bugu/internal/dao"
	"bugu/internal/model"
	"bugu/pkg/app"
	"github.com/jinzhu/gorm"
)

type CreateOverloadRequest struct {
	OverloadIntroduction string `form:"overload_introduction" binding:"required"`
	OverloadMaskCode     string `form:"overload_mask_code" binding:"required"`
	JsonOverloadInput    string `form:"json_overload_input" binding:"required"`
	JsonOverloadOutput   string `form:"json_overload_output" binding:"required"`
	ModuleName           string `form:"module_name" binding:"required"`
}
type UpdateOverloadRequest struct {
	ID                   uint   `form:"id" binding:"required,gte=1"`
	OverloadIntroduction string `form:"overload_introduction" binding:"required"`
	OverloadMaskCode     string `form:"overload_mask_code" binding:"required"`
	ModuleName           string `form:"module_name" binding:"required"`
}

type DeleteOverloadRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type OverloadRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type OverloadListRequest struct {
	ModuleID uint `form:"module_id" binding:"max=100"`
}

type Overload struct {
	*gorm.Model
	OverloadIntroduction string `json:"overload_introduction"`
	OverloadMaskCode     string `json:"overload_mask_code"`
	OverloadInput        []*model.OverloadInput
	OverloadOutput       []*model.OverloadOutput
	ModuleID             uint `json:"module_id"`
}

type OverloadInput struct {
	OIIntroduction string `json:"oi_introduction"`
	OIType         string `json:"oi_type"`
}

type OverloadOutput struct {
	OOIntroduction string `json:"oo_introduction"`
	OOType         string `json:"oo_type"`
}

type CountOverloadRequest struct {
	ModuleID uint `form:"module_id" binding:"max=100"`
}

func (svc *Service) CreateOverload(param *CreateOverloadRequest) error {
	return svc.dao.CreateOverload(&dao.Overload{
		OverloadIntroduction: param.OverloadIntroduction,
		OverloadMaskCode:     param.OverloadMaskCode,
	}, param.ModuleName, param.JsonOverloadInput, param.JsonOverloadOutput)
}

func (svc *Service) UpdateOverload(param *UpdateOverloadRequest) error {
	return svc.dao.UpdateOverload(&dao.Overload{
		Model:                &gorm.Model{ID: param.ID},
		OverloadIntroduction: param.OverloadIntroduction,
		OverloadMaskCode:     param.OverloadMaskCode,
	}, param.ModuleName)
}

func (svc *Service) DeleteOverload(param *DeleteOverloadRequest) error {
	return svc.dao.DeleteOverload(param.ID)
}

func (svc *Service) GetOverload(param *OverloadRequest) (*Overload, error) {
	overload, err := svc.dao.GetOverload(param.ID)
	if err != nil {
		return nil, err
	}
	overloadInputs, err := svc.dao.GetOverloadInputList(param.ID)
	if err != nil {
		return nil, err
	}
	overloadOutputs, err := svc.dao.GetOverloadOutputList(param.ID)
	if err != nil {
		return nil, err
	}
	return &Overload{
		Model:            &gorm.Model{ID: overload.ID, CreatedAt: overload.CreatedAt},
		OverloadIntroduction: overload.OverloadIntroduction,
		OverloadMaskCode:     overload.OverloadMaskCode,
		OverloadInput:        overloadInputs,
		OverloadOutput:       overloadOutputs,
		ModuleID:             overload.ModuleID,
	}, nil
}

func (svc *Service) GetOverloadList(param *OverloadListRequest, pager *app.Pager) ([]*model.Overload, error) {
	return svc.dao.GetOverloadList(param.ModuleID, pager.Page, pager.PageSize)
}

func (svc *Service) CountOverload(param *CountOverloadRequest) (int, error) {
	return svc.dao.CountOverload(param.ModuleID)
}