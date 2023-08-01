package model

import "github.com/mytrix-technology/admintrix/internal/model/internal"

// RoleMenu is the golang structure for table sys_role_menu.
type RoleMenu internal.RoleMenu

// Fill with you ideas below.

// 角色权限菜单列表
type RoleMenuInfo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ParentId int    `json:"parentId"`
	Checked  bool   `json:"checked"`
	Open     bool   `json:"open"`
}

type RoleMenuSaveReq struct {
	RoleId  int   `p:"roleId" v:"required#角色ID不能为空"`
	MenuIds []int `p:"menuIds" v:"required#菜单权限节点不能为空"`
}
