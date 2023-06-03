package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/mssola/useragent"
	"github.com/mytrix-technology/admintrix/internal/dao"
	"github.com/mytrix-technology/admintrix/internal/model"
	"github.com/mytrix-technology/admintrix/pkg/utils"
)

func LoginLog(r *ghttp.Request) {
	r.Middleware.Next()
	fmt.Println("登录日志中间件")

	urlItem := []string{"/login", "/logout"}
	if utils.InStringArray(r.RequestURI, urlItem) {

		// 获取浏览器信息
		userAgent := r.Header.Get("User-Agent")
		ua := useragent.New(userAgent)

		var entity model.LoginLog
		entity.Method = r.Method
		entity.OperUrl = r.URL.String()
		entity.OperIp = r.GetClientIp()
		entity.OperLocation = utils.GetIpCity(entity.OperIp)
		entity.RequestParam = string(r.GetBody())
		entity.Status = 0
		entity.Os = ua.OS()
		entity.Browser, _ = ua.Browser()
		entity.UserAgent = r.UserAgent()
		entity.CreateTime = gtime.Now()
		entity.Mark = 1

		if r.RequestURI == "/login" {
			// 登录成功
			var jsonObj map[string]interface{}
			json.Unmarshal(r.GetBody(), &jsonObj)

			user, err := dao.User.Where("username=?", jsonObj["username"]).FindOne()
			if err != nil {
				return
			}
			entity.Type = 1
			entity.Username = user.Username
			entity.CreateUser = user.Id
		} else if r.RequestURI == "/logout" {
			entity.Type = 3
			entity.Username = utils.UInfo(r).Username
			entity.CreateUser = utils.Uid(r)
		}

		dao.LoginLog.Insert(entity)
	}
}
