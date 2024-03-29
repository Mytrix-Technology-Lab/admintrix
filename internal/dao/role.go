// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mytrix-technology/admintrix/internal/dao/internal"
)

// internalRoleDao is internal type for wrapping internal DAO implements.
type internalRoleDao = *internal.RoleDao

// RoleDao is the data access object for table sys_role.
// You can define custom methods on it to extend its functionality as you wish.
type RoleDao struct {
	internalRoleDao
}

var (
	// Role is globally public accessible object for table sys_role operations.
	Role = RoleDao{
		internal.NewRoleDao(),
	}
)

// Fill with you ideas below.
