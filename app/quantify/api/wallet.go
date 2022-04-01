package api

import (
	"gfast/app/quantify/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"ss-quant-trade-common/model/rest_model"
)

type wallet struct {
	CommonBase
}

var Wallet = new(wallet)

// QuantifyTaskDefineList 查询支持的机器人策略
func (c *wallet) QuantifyTaskDefineList(r *ghttp.Request) {
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
