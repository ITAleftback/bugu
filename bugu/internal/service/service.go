/**
 * @Author: Anpw
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/5/26 23:16
 */

package service

import (
	"bugu/global"
	"bugu/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
