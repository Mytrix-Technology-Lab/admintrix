package middleware

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

func CORS(r *ghttp.Request) {
	fmt.Println("跨域处理中间件")
	r.Response.CORSDefault()

	r.Middleware.Next()
}
