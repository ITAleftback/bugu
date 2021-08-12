/**
 * @Author: Anpw
 * @Description:
 * @File:  plugin
 * @Version: 1.0.0
 * @Date: 2021/6/3 17:32
 */

package service

import (
	"bugu/internal/model"
	"bugu/pkg/app"
	"github.com/jinzhu/gorm"
)

type CreatePluginRequest struct {
	PluginName         string `form:"plugin_name" binding:"required"`
	PluginIntroduction string `form:"plugin_introduction" binding:"required"`
	FileMaskCode       string `form:"file_mask_code" binding:"required"`
	ChipName           string `form:"chip_name" binding:"required"`
}

type UpdatePluginRequest struct {
	ID                 uint   `form:"id" binding:"required,gte=1"`
	PluginName         string `form:"plugin_name" binding:"required"`
	PluginIntroduction string `form:"plugin_introduction" binding:"required"`
	FileMaskCode       string `form:"file_mask_code" binding:"required"`
	ChipName           string `form:"chip_name" binding:"required"`
}

type DeletePluginRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type PluginRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type CountPluginRequest struct {
	ChipID uint `form:"chip_id" binding:"max=100"`
}

type PluginListRequest struct {
	ChipID uint `form:"chip_id" binding:"max=100"`
}

type Plugin struct {
	*gorm.Model
	PluginName         string `json:"plugin_name"`
	PluginIntroduction string `json:"plugin_introduction"`
	FileMaskCode       string `json:"file_mask_code"`
	ChipID             uint   `json:"chip_id"`
}

func (svc *Service) CreatePlugin(param *CreatePluginRequest) error {
	return svc.dao.CreatePlugin(param.PluginName, param.PluginIntroduction, param.FileMaskCode, param.ChipName)
}

func (svc *Service) UpdatePlugin(param *UpdatePluginRequest) error {
	return svc.dao.UpdatePlugin(param.ID, param.PluginName, param.PluginIntroduction, param.FileMaskCode, param.ChipName)
}

func (svc *Service) DeletePlugin(param *DeletePluginRequest) error {
	return svc.dao.DeletePlugin(param.ID)
}

func (svc *Service) GetPlugin(param *PluginRequest) (*Plugin, error) {
	plugin, err := svc.dao.GetPlugin(param.ID)
	if err != nil {
		return nil, err
	}
	return &Plugin{
		Model:            &gorm.Model{ID: plugin.ID, CreatedAt: plugin.CreatedAt},
		PluginName:         plugin.PluginName,
		PluginIntroduction: plugin.PluginIntroduction,
		FileMaskCode:       plugin.FileMaskCode,
		ChipID:             plugin.ChipID,
	}, nil
}

func (svc *Service) GetPluginList(param *PluginListRequest, pager *app.Pager) ([]*model.Plugin, error) {
	return svc.dao.GetPluginList(param.ChipID, pager.Page, pager.PageSize)
}

func (svc *Service) CountPlugin(param *CountPluginRequest) (int, error) {
	return svc.dao.CountPlugin(param.ChipID)
}