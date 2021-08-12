/**
 * @Author: Anpw
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/5/29 17:00
 */

package dao

import "bugu/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	return auth.Get(d.engine)
}
