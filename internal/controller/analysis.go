package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

var Analysis = new(analysisCtl)

type analysisCtl struct{}

func (c *analysisCtl) Index(r *ghttp.Request) {
	// 渲染模板
	//response.BuildTpl(r, "analysis/index.html").WriteTpl()
}
