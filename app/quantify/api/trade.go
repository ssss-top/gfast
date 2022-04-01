package api

import (
	"gfast/app/quantify/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"ss-quant-trade-common/model/rest_model"
)

type trade struct {
	CommonBase
}

var Trade = new(trade)

// QuantifyTaskDefineList 查询支持的机器人策略
func (c *trade) QuantifyTaskDefineList(r *ghttp.Request) {
	user := c.GetCurrentUser(r.GetCtx())
	req := rest_model.UserIDRequest{
		UserID: int64(user.GetUserId()),
	}
	list, err := service.Trade.ListQTaskDefine(&req)
	if err != nil {
		g.Log().Error(err)
		c.FailJsonExit(r, "获取数据失败")
	}
	c.SusJsonExit(r, g.Map{
		"list": list,
	})
}

// QuantifyTaskDefineDetails 查询策略详情
func (c *trade) QuantifyTaskDefineDetails(r *ghttp.Request) {
	user := c.GetCurrentUser(r.GetCtx())
	req := rest_model.FindTaskDefineDetailRequest{
		UserIDRequest: rest_model.UserIDRequest{
			UserID: int64(user.GetUserId()),
		},
		TaskDefineID: r.GetQueryString("id"),
	}
	details, err := service.Trade.QTaskDefineDetails(&req)
	if err != nil {
		g.Log().Error(err)
		c.FailJsonExit(r, "获取数据失败")
	}
	c.SusJsonExit(r, g.Map{
		"details": details,
	})
}
