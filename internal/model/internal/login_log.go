package internal

import (
	"github.com/gogf/gf/os/gtime"
)

type LoginLog struct {
	Id           int64       `orm:"id,primary"    json:"id"`           // 主键ID
	Username     string      `orm:"username"      json:"username"`     // 操作账号
	Method       string      `orm:"method"        json:"method"`       // 操作方法
	OperUrl      string      `orm:"oper_url"      json:"operUrl"`      // 请求URL
	OperIp       string      `orm:"oper_ip"       json:"operIp"`       // 主机地址
	OperLocation string      `orm:"oper_location" json:"operLocation"` // 操作地点
	Os           string      `orm:"os"            json:"os"`           // 操作系统
	RequestParam string      `orm:"request_param" json:"requestParam"` // 请求参数
	Browser      string      `orm:"browser"       json:"browser"`      // 浏览器
	Result       string      `orm:"result"        json:"result"`       // 返回参数
	Status       uint        `orm:"status"        json:"status"`       // 操作状态：0操作成功 1操作失败
	Type         uint        `orm:"type"          json:"type"`         // 操作类型：1登录成功 2登录失败 3注销成功 2注销失败
	UserAgent    string      `orm:"user_agent"    json:"userAgent"`    // 代理信息
	Note         string      `orm:"note"          json:"note"`         // 备注
	CreateUser   int         `orm:"create_user"   json:"createUser"`   // 添加人
	CreateTime   *gtime.Time `orm:"create_time"   json:"createTime"`   // 操作时间
	UpdateUser   int         `orm:"update_user"   json:"updateUser"`   // 更新人
	UpdateTime   *gtime.Time `orm:"update_time"   json:"updateTime"`   // 更新时间
	Mark         int         `orm:"mark"          json:"mark"`         // 有效标识
}
