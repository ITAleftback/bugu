/**
 * @Author: Anpw
 * @Description:
 * @File:  module
 * @Version: 1.0.0
 * @Date: 2021/6/3 22:13
 */

package service

import (
	"bugu/internal/model"
	"bugu/pkg/app"
	"github.com/jinzhu/gorm"
)

type CreateModuleRequest struct {
	ModuleName         string `form:"module_name" binding:"required"`
	ModuleIntroduction string `form:"module_introduction" binding:"required"`
	PluginName         string `form:"plugin_name" binding:"required"`
}

type UpdateModuleRequest struct {
	ID                 uint   `form:"id" binding:"required,gte=1"`
	ModuleName         string `form:"module_name" binding:"required"`
	ModuleIntroduction string `form:"module_introduction" binding:"required"`
	PluginName         string `form:"plugin_name" binding:"required"`
}

type DeleteModuleRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type ModuleRequest struct {
	ID uint `form:"id" binding:"required,gte=1"`
}

type ModuleListRequest struct {
	PluginID uint `form:"plugin_id" binding:"max=100"`
}
type CountModuleRequest struct {
	PluginID uint `form:"plugin_id" binding:"max=100"`
}


type Module struct {
	*gorm.Model
	ModuleName         string `json:"module_name"`
	ModuleIntroduction string `json:"module_introduction"`
	PluginID           uint   `json:"plugin_id"`
}

func (svc *Service) CreateModule(param *CreateModuleRequest) error {
	return svc.dao.CreateModule(param.ModuleName, param.ModuleIntroduction, param.PluginName)
}

func (svc *Service) UpdateModule(param *UpdateModuleRequest) error {
	return svc.dao.UpdateModule(param.ID, param.ModuleName, param.ModuleIntroduction, param.PluginName)
}

func (svc *Service) DeleteModule(param *DeleteModuleRequest) error {
	return svc.dao.DeleteModule(param.ID)
}

func (svc *Service) GetModule(param *ModuleRequest) (*Module, error) {
	module, err := svc.dao.GetModule(param.ID)
	if err != nil {
		return nil, err
	}
	return &Module{
		Model:            &gorm.Model{ID: module.ID, CreatedAt: module.CreatedAt},
		ModuleName:         module.ModuleName,
		ModuleIntroduction: module.ModuleIntroduction,
		PluginID:           module.PluginID,
	}, nil
}

func (svc *Service) GetModuleList(param *ModuleListRequest, pager *app.Pager) ([]*model.Module, error) {
	return svc.dao.GetModuleList(param.PluginID, pager.Page, pager.PageSize)
}

func (svc *Service) CountModule(param *CountModuleRequest) (int, error) {
	return svc.dao.CountModule(param.PluginID)
}