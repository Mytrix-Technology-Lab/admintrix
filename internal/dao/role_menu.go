// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mytrix-technology/admintrix/internal/dao/internal"
)

// internalRoleMenuDao is internal type for wrapping internal DAO implements.
type internalRoleMenuDao = *internal.RoleMenuDao

// RoleMenuDao is the data access object for table sys_role_menu.
// You can define custom methods on it to extend its functionality as you wish.
type RoleMenuDao struct {
	internalRoleMenuDao
}

var (
	// RoleMenu is globally public accessible object for table sys_role_menu operations.
	RoleMenu = RoleMenuDao{
		internal.NewRoleMenuDao(),
	}
)

// Fill with you ideas below.
