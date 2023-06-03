package dao

import "github.com/mytrix-technology/admintrix/internal/dao/internal"

type OperLogDao struct {
	internal.OperLogDao
}

var (
	// OperLog is globally public accessible object for table sys_oper_log operations.
	OperLog = OperLogDao{
		internal.OperLog,
	}
)
