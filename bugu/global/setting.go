/**
 * @Author: Anpw
 * @Description:
 * @File:  setting
 * @Version: 1.0.0
 * @Date: 2021/5/13 21:27
 */

package global

import (
	"bugu/pkg/logger"
	"bugu/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
