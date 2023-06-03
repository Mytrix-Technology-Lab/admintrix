package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gview"
	"github.com/gogf/swagger"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})

	gview.Instance().BindFuncMap(gview.FuncMap{})
}
