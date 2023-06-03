package dao

import "github.com/mytrix-technology/admintrix/internal/dao/internal"

type LoginLogDao struct {
	internal.LoginLogDao
}

var (
	// LoginLog is globally public accessible object for table sys_login_log operations.
	LoginLog = LoginLogDao{
		internal.LoginLog,
	}
)
