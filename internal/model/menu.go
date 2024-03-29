package model

import "github.com/mytrix-technology/admintrix/internal/model/internal"

// Menu is the golang structure for table sys_menu.
type Menu internal.Menu

// Fill with you ideas below.

// 菜单Vo
type TreeNode struct {
	Menu
	Children []*TreeNode `json:"children"` // 子菜单
}

// 列表查询条件
type MenuQueryReq struct {
	Title string `p:"name"` // 菜单标题
}

// 添加菜单
type MenuAddReq struct {
	ParentId    int    `p:"parent_id"`                         // 上级ID
	Title       string `p:"title"       v:"required#菜单名称不能为空"` // 菜单标题
	Icon        string `p:"icon"        v:"required#菜单图标不能为空"` // 图标
	Path        string `p:"path"        v:"required#菜单路径不能为空"` // URL地址
	Component   string `p:"component"`                         // 菜单组件
	Target      string `p:"target"`                            // 打开方式：0组件 1内链 2外链
	Permission  string `p:"permission"`                        // 权限标识
	Type        int    `p:"type"        v:"required#请选择菜单类型"`  // 类型：1模块 2导航 3菜单 4节点
	Status      int    `p:"status"      v:"required#请选择菜单状态"`  // 状态：1正常 2禁用
	Hide        int    `p:"hide"`                              // 是否可见：1是 2否
	Note        string `p:"note"`                              // 菜单备注
	Sort        int    `p:"sort"        v:"required#请输入菜单排序号"` // 显示顺序
	CheckedList []int  `p:"checkedList"`                       // 权限节点
}

// 更新菜单
type MenuUpdateReq struct {
	Id          int    `p:"id" 		   v:"required#主键ID不能为空"`
	ParentId    int    `p:"parent_id"`                         // 上级ID
	Title       string `p:"title"       v:"required#菜单名称不能为空"` // 菜单标题
	Icon        string `p:"icon"        v:"required#菜单图标不能为空"` // 图标
	Path        string `p:"path"        v:"required#菜单路径不能为空"` // URL地址
	Component   string `p:"component"`                         // 菜单组件
	Target      string `p:"target"`                            // 打开方式：0组件 1内链 2外链
	Permission  string `p:"permission"`                        // 权限标识
	Type        int    `p:"type"        v:"required#请选择菜单类型"`  // 类型：1模块 2导航 3菜单 4节点
	Status      int    `p:"status"      v:"required#请选择菜单状态"`  // 是否显示：1显示 2不显示
	Hide        int    `p:"hide"`                              // 是否可见：1是 2否
	Note        string `p:"note"`                              // 菜单备注
	Sort        int    `p:"sort"        v:"required#请输入菜单排序号"` // 显示顺序
	CheckedList []int  `p:"checkedList"`                       // 权限节点
}

// 菜单删除
type MenuDeleteReq struct {
	Ids string `p:"ids" v:"required#请选择需要删除的数据记录"`
}

// 菜单信息
type MenuInfoVo struct {
	Menu
	CheckedList []interface{} `json:"checkedList"` // 权限节点列表
}

// 菜单树结构
type MenuTreeNode struct {
	Menu
	Children []*MenuTreeNode `json:"children"` // 子菜单
}
