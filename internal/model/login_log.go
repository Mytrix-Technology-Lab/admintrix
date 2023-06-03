package model

import "github.com/mytrix-technology/admintrix/internal/model/internal"

type LoginLog internal.LoginLog

type LoginLogPageReq struct {
	Username string `p:"username"` // 操作账号
	Page     int    `p:"page"`     // 页码
	Limit    int    `p:"limit"`    // 每页数
}

type LoginLogDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}
