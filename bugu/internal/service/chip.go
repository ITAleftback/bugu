/**
 * @Author: Anpw
 * @Description:
 * @File:  chip
 * @Version: 1.0.0
 * @Date: 2021/6/2 5:45
 */

package service

import (
	"bugu/internal/model"
	"bugu/pkg/app"
	"github.com/jinzhu/gorm"
)

type CreateChipRequest struct {
	ChipName         string `form:"chip_name" binding:"required"`
	ManufacturerName string `form:"manufacturer_name" binding:"required"`
	Type             string `form:"type" binding:"required"`
}

type UpdateChipRequest struct {
	ID               uint   `form:"id" binding:"required,gte=1"`
	ChipName         string `form:"chip_name" binding:"required"`
	ManufacturerName string `form:"manufacturer_name" binding:"required"`
	Type             string `form:"type" binding:"required"`
}

type DeleteChipRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type ChipRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type ChipListRequest struct {
	ChipName string `form:"chip_name" binding:"max=100"`
}

type CountChipRequest struct {
	ChipName string `form:"chip_name" binding:"max=100"`
}

type Chip struct {
	*gorm.Model
	ChipName         string `json:"chip_name"`
	ManufacturerName string `json:"manufacturer_name"`
	Type             string `json:"type"`
}

func (svc *Service) CreateChip(param *CreateChipRequest) error {
	return svc.dao.CreateChip(param.ChipName, param.ManufacturerName, param.Type)
}

func (svc *Service) UpdateChip(param *UpdateChipRequest) error {
	return svc.dao.UpdateChip(param.ID, param.ChipName, param.ManufacturerName, param.Type)
}

func (svc *Service) DeleteChip(param *DeleteChipRequest) error {
	return svc.dao.DeleteChip(param.ID)
}

func (svc *Service) GetChip(param *ChipRequest) (*Chip, error) {
	chip, err := svc.dao.GetChip(param.ID)
	if err != nil {
		return nil, err
	}
	return &Chip{
		Model:            &gorm.Model{ID: chip.ID, CreatedAt: chip.CreatedAt},
		ChipName:         chip.ChipName,
		ManufacturerName: chip.ManufacturerName,
		Type:             chip.Type,
	}, nil
}

func (svc *Service) GetChipList(param *ChipListRequest, pager *app.Pager) ([]*model.Chip, error) {
	return svc.dao.GetChipList(param.ChipName, pager.Page, pager.PageSize)
}

func (svc *Service) CountChip(param *CountChipRequest) (int, error) {
	return svc.dao.CountChip(param.ChipName)
}
