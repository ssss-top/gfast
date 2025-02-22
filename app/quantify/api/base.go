package api

import (
	"context"
	"gfast/app/common/api"
	"gfast/app/system/dao"
	"gfast/app/system/service"
)

type CommonBase struct {
	api.CommonBase
}

// GetCurrentUser 获取当前登陆用户信息
func (c *CommonBase) GetCurrentUser(ctx context.Context) *dao.CtxUser {
	context := service.Context.Get(ctx)
	if context == nil {
		return nil
	}
	return context.User
}
