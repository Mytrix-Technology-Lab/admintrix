package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	/* 案例演示 */
	s.Group("example", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controllers.Example.List)
		//group.POST("/add", controllers.Example.Add)
		//group.PUT("/update", controllers.Example.Update)
		//group.DELETE("/delete/:ids", controllers.Example.Delete)
		//group.PUT("/status", controllers.Example.Status)
		//group.PUT("/isVip", controllers.Example.IsVip)
	})

	s.Group("example2", func(group *ghttp.RouterGroup) {
		//group.GET("/list", controllers.Example2.List)
		//group.POST("/add", controllers.Example2.Add)
		//group.PUT("/update", controllers.Example2.Update)
		//group.DELETE("/delete/:ids", controllers.Example2.Delete)
		//group.PUT("/status", controllers.Example2.Status)
	})
}
