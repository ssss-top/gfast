package router

import (
	"gfast/app/quantify/api"
	"gfast/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//加载路由
func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Group("/quantify", func(group *ghttp.RouterGroup) {
			//context拦截器
			group.Middleware(middleware.Ctx, middleware.Auth)
			group.Group("/define", func(group *ghttp.RouterGroup) {
				group.GET("/list", api.Trade.QuantifyTaskDefineList)
				group.GET("/details", api.Trade.QuantifyTaskDefineDetails)
			})
		})
	})
}
