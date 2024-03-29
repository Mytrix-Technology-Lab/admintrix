package internal

import "github.com/gogf/gf/os/gtime"

type Menu struct {
	Id         int         `orm:"id,primary"  json:"id"`         // 主键ID
	ParentId   int         `orm:"parent_id"   json:"parentId"`   // 父级ID
	Title      string      `orm:"title"       json:"title"`      // 菜单标题
	Icon       string      `orm:"icon"        json:"icon"`       // 图标
	Path       string      `orm:"path"        json:"path"`       // 菜单路径
	Component  string      `orm:"component"   json:"component"`  // 菜单组件
	Target     string      `orm:"target"      json:"target"`     // 打开方式：0组件 1内链 2外链
	Permission string      `orm:"permission"  json:"permission"` // 权限标识
	Type       int         `orm:"type"        json:"type"`       // 类型：0菜单 1节点
	Method     string      `orm:"method"      json:"method"`     // 请求方式
	Status     int         `orm:"status"      json:"status"`     // 状态：1正常 2禁用
	Hide       int         `orm:"hide"        json:"hide"`       // 是否可见：1是 2否
	Note       string      `orm:"note"        json:"note"`       // 备注
	Sort       int         `orm:"sort"        json:"sort"`       // 显示顺序
	CreateUser int         `orm:"create_user" json:"createUser"` // 添加人
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 创建时间
	UpdateUser int         `orm:"update_user" json:"updateUser"` // 更新人
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
	Mark       int         `orm:"mark"        json:"mark"`       // 有效标识
}
