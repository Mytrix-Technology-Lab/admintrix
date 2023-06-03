package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/mojocn/base64Captcha"
	"github.com/mytrix-technology/admintrix/internal/service"
	"github.com/mytrix-technology/admintrix/pkg/utils/common"
)

var Login = new(loginCtl)

type loginCtl struct{}

type LoginReq struct {
	UserName string `p:"username" v:"required|length:5,30#请输入登录账号|账号长度为：min-max位"`
	Password string `p:"password" v:"required|length:5,12#请输入密码|密码长度为：min-max位"`
	Captcha  string `p:"captcha" v:"required|length:4,6#请输入验证码|验证码长度不够"`
	IdKey    string `p:"idKey" v:"required#验证码KEY不能为空"`
}

func (c *loginCtl) Login(r *ghttp.Request) {
	if r.Method == "POST" {
		var req *LoginReq

		// 获取参数并验证
		if err := r.Parse(&req); err != nil {
			// 返回错误信息
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 校验验证码
		verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		if !verifyRes {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  "验证码不正确",
			})
		}

		// 系统登录
		if token, err := service.Login.UserLogin(req.UserName, req.Password, r); err != nil {
			// 登录错误
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		} else {
			// 登录成功
			r.Response.WriteJsonExit(common.JsonResult{
				Code: 0,
				Msg:  "登录成功",
				Data: g.Map{
					"access_token": token,
				},
			})
		}
	}
}
