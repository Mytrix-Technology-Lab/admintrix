package service

import (
	"errors"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/mytrix-technology/admintrix/internal/dao"
	"github.com/mytrix-technology/admintrix/pkg/utils" //nolint:typecheck
	jwt "github.com/mytrix-technology/admintrix/pkg/utils"
)

var Login = new(loginService)

type loginService struct{}

var SessionList = gmap.New(true)

func (s *loginService) UserLogin(username, password string, r *ghttp.Request) (string, error) {
	user, err := dao.User.FindOne("username=? and mark=1", username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("用户名或者密码不正确")
	}

	pwd, _ := utils.Md5(password + user.Username)
	if user.Password != pwd {
		return "", errors.New("密码不正确")
	}

	if user.Status != 1 {
		return "", errors.New("您的账号已被禁用,请联系管理员")
	}

	dao.User.Data(g.Map{
		"login_time":  gtime.Now(),
		"login_ip":    utils.GetClientIp(r),
		"update_time": gtime.Now(),
	})

	token, _ := jwt.GenerateToken(user.Id, user.Username, user.Password)

	r.Session.Set("userId", user.Id)
	r.Session.Set("userInfo", user)
	sessionId := r.Session.Id()
	SessionList.Set(sessionId, r.Session)
	return token, nil
}
