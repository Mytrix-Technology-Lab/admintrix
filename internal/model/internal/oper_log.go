package internal

import "github.com/gogf/gf/os/gtime"

type OperLog struct {
	Id           int64       `orm:"id,primary"    json:"id"`           // 主键ID
	Model        string      `orm:"model"         json:"model"`        // 操作模块
	OperType     int         `orm:"oper_type"     json:"operType"`     // 操作类型：0其它 1新增 2修改 3删除 4查询 5设置状态 6导入 7导出 8设置权限 9设置密码
	OperMethod   string      `orm:"oper_method"   json:"operMethod"`   // 操作方法
	Username     string      `orm:"username"      json:"username"`     // 操作账号
	OperName     string      `orm:"oper_name"     json:"operName"`     // 操作用户
	OperUrl      string      `orm:"oper_url"      json:"operUrl"`      // 请求URL
	OperIp       string      `orm:"oper_ip"       json:"operIp"`       // 主机地址
	OperLocation string      `orm:"oper_location" json:"operLocation"` // 操作地点
	RequestParam string      `orm:"request_param" json:"requestParam"` // 请求参数
	Result       string      `orm:"result"        json:"result"`       // 返回参数
	Status       int         `orm:"status"        json:"status"`       // 日志状态：0正常日志 1错误日志
	UserAgent    string      `orm:"user_agent"    json:"userAgent"`    // 代理信息
	Note         string      `orm:"note"          json:"note"`         // 备注
	CreateUser   int         `orm:"create_user"   json:"createUser"`   // 添加人
	CreateTime   *gtime.Time `orm:"create_time"   json:"createTime"`   // 操作时间
	UpdateUser   int         `orm:"update_user"   json:"updateUser"`   // 更新人
	UpdateTime   *gtime.Time `orm:"update_time"   json:"updateTime"`   // 更新时间
	Mark         int         `orm:"mark"          json:"mark"`         // 有效标识
}
