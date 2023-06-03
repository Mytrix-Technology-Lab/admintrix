package middleware

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/mytrix-technology/admintrix/pkg/utils"
	"github.com/mytrix-technology/admintrix/pkg/utils/common"
	"time"
)

func CheckLogin(r *ghttp.Request) {
	fmt.Println("登录验证中间件")

	urlItem := []string{"/captcha", "/login"}
	if !utils.InStringArray(r.RequestURI, urlItem) {
		token := r.GetHeader("Authorization")
		token = gstr.Replace(token, "Bearer ", "")

		claim, err := utils.ParseToken(token)
		if err != nil {
			fmt.Println("解析token出现错误：", err)
			r.Response.WriteJsonExit(common.JsonResult{
				Code: 401,
				Msg:  "Token已过期",
			})
		} else if time.Now().Unix() > claim.ExpiresAt {
			fmt.Println("时间超时")
			r.Response.WriteJsonExit(common.JsonResult{
				Code: 401,
				Msg:  "时间超时",
			})
		}
	}

	r.Middleware.Next()
}
