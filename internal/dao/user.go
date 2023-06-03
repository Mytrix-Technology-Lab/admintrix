// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mytrix-technology/admintrix/internal/dao/internal"
)

// internalUserDao is internal type for wrapping internal DAO implements.
type internalUserDao = *internal.UserDao

// UserDao is the data access object for table sys_user.
// You can define custom methods on it to extend its functionality as you wish.
type UserDao struct {
	internalUserDao
}

var (
	// User is globally public accessible object for table sys_user operations.
	User = UserDao{
		internal.NewUserDao(),
	}
)

// Fill with you ideas below.