package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	/* 案例演示 */
	s.Group("example", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Example.List)
		//group.POST("/add", controller.Example.Add)
		//group.PUT("/update", controller.Example.Update)
		//group.DELETE("/delete/:ids", controller.Example.Delete)
		//group.PUT("/status", controller.Example.Status)
		//group.PUT("/isVip", controller.Example.IsVip)
	})

	s.Group("example2", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controller.Example2.List)
		//group.POST("/add", controller.Example2.Add)
		//group.PUT("/update", controller.Example2.Update)
		//group.DELETE("/delete/:ids", controller.Example2.Delete)
		//group.PUT("/status", controller.Example2.Status)
	})
}
