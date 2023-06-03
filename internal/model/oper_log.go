package model

import "github.com/mytrix-technology/admintrix/internal/model/internal"

// OperLog is the golang structure for table sys_oper_log.
type OperLog internal.OperLog

type OperLogPageReq struct {
	Username string `p:"username"` // 操作账号
	Model    string `p:"model"`    // 操作模块
	Page     int    `p:"page"`     // 页码
	Limit    int    `p:"limit"`    // 每页数
}
