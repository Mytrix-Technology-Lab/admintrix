package model

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/mytrix-technology/admintrix/internal/model/internal"
)

type User internal.User

type UserPageReq struct {
	Realname string `p:"realname"` // 用户姓名
	Gender   int    `p:"gender"`   // 性别:1男 2女 3保密
	Page     int    `p:"page"`     // 页码
	Limit    int    `p:"limit"`    // 每页数
}

type UserAddReq struct {
	Realname   string      `p:"realname"      v:"required#真实姓名不能为空"`  // 真实姓名
	Nickname   string      `p:"nickname"      v:"required#昵称不能为空"`    // 昵称
	Gender     int         `p:"gender"        v:"required#性别不能为空"`    // 性别:1男 2女 3保密
	Avatar     string      `p:"avatar"        v:"required#请上传头像"`     // 头像
	Mobile     string      `p:"mobile"        v:"required#手机号不能为空"`   // 手机号码
	Email      string      `p:"email"         v:"required#电子邮件不能为空"`  // 邮箱地址
	Birthday   *gtime.Time `p:"birthday"      v:"required#出生日期不能为空"`  // 出生日期
	DeptId     int         `p:"dept_id"       v:"required#请选择部门"`     // 部门ID
	LevelId    int         `p:"level_id"      v:"required#请选择职级"`     // 职级ID
	PositionId int         `p:"position_id"   v:"required#请选择岗位"`     // 岗位ID
	City       []string    `p:"city"		  v:"required#请选择省市区"`          // 省市区
	Address    string      `p:"address"`                              // 详细地址
	Username   string      `p:"username"      v:"required#登录用户名不能为空"` // 登录用户名
	Password   string      `p:"password"`                             // 登录密码
	Intro      string      `p:"intro"`                                // 个人简介
	Status     int         `p:"status"        v:"required#请选择状态"`     // 状态：1正常 2禁用
	Note       string      `p:"note"`                                 // 备注
	Sort       int         `p:"sort"          v:"required#排序号不能为空"`   // 排序号
	RoleIds    []int       `p:"roleIds"		  v:"required#请选择用户角色"`      // 用户角色
}

type UserUpdateReq struct {
	Id         int         `p:"id" v:"required#主键ID不能为空"`
	Realname   string      `p:"realname"      v:"required#真实姓名不能为空"`  // 真实姓名
	Nickname   string      `p:"nickname"      v:"required#昵称不能为空"`    // 昵称
	Gender     int         `p:"gender"        v:"required#性别不能为空"`    // 性别:1男 2女 3保密
	Avatar     string      `p:"avatar"        v:"required#请上传头像"`     // 头像
	Mobile     string      `p:"mobile"        v:"required#手机号不能为空"`   // 手机号码
	Email      string      `p:"email"         v:"required#电子邮件不能为空"`  // 邮箱地址
	Birthday   *gtime.Time `p:"birthday"      v:"required#出生日期不能为空"`  // 出生日期
	DeptId     int         `p:"dept_id"       v:"required#请选择部门"`     // 部门ID
	LevelId    int         `p:"level_id"      v:"required#请选择职级"`     // 职级ID
	PositionId int         `p:"position_id"   v:"required#请选择岗位"`     // 岗位ID
	City       []string    `p:"city"		  v:"required#请选择省市区"`          // 省市区
	Address    string      `p:"address"`                              // 详细地址
	Username   string      `p:"username"      v:"required#登录用户名不能为空"` // 登录用户名
	Password   string      `p:"password"`                             // 登录密码
	Intro      string      `p:"intro"`                                // 个人简介
	Status     int         `p:"status"        v:"required#请选择状态"`     // 状态：1正常 2禁用
	Note       string      `p:"note"`                                 // 备注
	Sort       int         `p:"sort"          v:"required#排序号不能为空"`   // 排序号
	RoleIds    []int       `p:"roleIds"		  v:"required#请选择用户角色"`      // 用户角色
}

type UserDeleteReq struct {
	Ids string `p:"ids"  v:"required#请选择要删除的数据记录"`
}

type UserStatusReq struct {
	Id     int `p:"id" v:"required#主键ID不能为空"`
	Status int `p:"status"    v:"required#状态不能为空"`
}

type UserResetPwdReq struct {
	Id int `p:"id" v:"required#主键ID不能为空"`
}

type UserInfoReq struct {
	Avatar   string `p:"avatar"`                              // 头像
	Realname string `p:"realname"      v:"required#真实姓名不能为空"` // 真实姓名
	Nickname string `p:"nickname"      v:"required#昵称不能为空"`   // 昵称
	Gender   int    `p:"gender"        v:"required#性别不能为空"`   // 性别:1男 2女 3保密
	Mobile   string `p:"mobile"        v:"required#手机号不能为空"`  // 手机号码
	Email    string `p:"email"         v:"required#电子邮件不能为空"` // 邮箱地址
	Address  string `p:"address"`                             // 详细地址
	Intro    string `p:"intro"`                               // 个人简介
}

type UserInfoVo struct {
	User
	GenderName   string      `json:"genderName"`   // 性别
	LevelName    string      `json:"levelName"`    // 职级
	PositionName string      `json:"positionName"` // 岗位
	DeptName     string      `json:"deptName"`     // 部门
	RoleIds      interface{} `json:"RoleIds"`      // 角色ID
	RoleList     interface{} `json:"roleList"`     // 角色列表
	City         interface{} `json:"city"`         // 省市区
}

type UpdatePwd struct {
	OldPassword string `p:"oldPassword"      v:"required#旧密码不能为空"` // 旧密码
	NewPassword string `p:"newPassword"      v:"required#新密码不能为空"` // 新密码
	RePassword  string `p:"rePassword"      v:"required#确认密码不能为空"` // 确认密码
}

type CheckUserReq struct {
	Username string `p:"username"      v:"required#用户名不能为空"` // 用户名
}

type ProfileInfoVo struct {
	Realname       string        `json:"realname"`       // 真实姓名
	Nickname       string        `json:"nickname"`       // 昵称
	Gender         int           `json:"gender"`         // 性别:1男 2女 3保密
	Avatar         string        `json:"avatar"`         // 头像
	Mobile         string        `json:"mobile"`         // 手机号码
	Email          string        `json:"email"`          // 邮箱地址
	City           []string      `json:"city"`           // 省市区
	Address        string        `json:"address"`        // 详细地址
	Intro          string        `json:"intro"`          // 个人简介
	Roles          []interface{} `json:"roles"`          // 用户角色
	Authorities    []interface{} `json:"authorities"`    // 用户权限
	PermissionList []string      `json:"permissionList"` // 权限列表

}
