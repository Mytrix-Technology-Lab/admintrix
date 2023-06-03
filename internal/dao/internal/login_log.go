package internal

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/mytrix-technology/admintrix/internal/model"
	"time"
)

type LoginLogDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns LoginLogColumns
}

type LoginLogColumns struct {
	Id           string // 主键ID
	Username     string // 操作账号
	Method       string // 操作方法
	OperUrl      string // 请求URL
	OperIp       string // 主机地址
	OperLocation string // 操作地点
	Os           string // 操作系统
	RequestParam string // 请求参数
	Browser      string // 浏览器
	Result       string // 返回参数
	Status       string // 操作状态：0操作成功 1操作失败
	Type         string // 操作类型：1登录成功 2登录失败 3注销成功 2注销失败
	UserAgent    string // 代理信息
	Note         string // 备注
	CreateUser   string // 添加人
	CreateTime   string // 操作时间
	UpdateUser   string // 更新人
	UpdateTime   string // 更新时间
	Mark         string // 有效标识
}

var (
	// LoginLog is globally public accessible object for table sys_login_log operations.
	LoginLog = LoginLogDao{
		M:     g.DB("default").Model("sys_login_log").Safe(),
		DB:    g.DB("default"),
		Table: "sys_login_log",
		Columns: LoginLogColumns{
			Id:           "id",
			Username:     "username",
			Method:       "method",
			OperUrl:      "oper_url",
			OperIp:       "oper_ip",
			OperLocation: "oper_location",
			Os:           "os",
			RequestParam: "request_param",
			Browser:      "browser",
			Result:       "result",
			Status:       "status",
			Type:         "type",
			UserAgent:    "user_agent",
			Note:         "note",
			CreateUser:   "create_user",
			CreateTime:   "create_time",
			UpdateUser:   "update_user",
			UpdateTime:   "update_time",
			Mark:         "mark",
		},
	}
)

// a global or package variable for long using.
func (d *LoginLogDao) Ctx(ctx context.Context) *LoginLogDao {
	return &LoginLogDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *LoginLogDao) As(as string) *LoginLogDao {
	return &LoginLogDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *LoginLogDao) TX(tx *gdb.TX) *LoginLogDao {
	return &LoginLogDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *LoginLogDao) Master() *LoginLogDao {
	return &LoginLogDao{M: d.M.Master()}
}

func (d *LoginLogDao) Slave() *LoginLogDao {
	return &LoginLogDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *LoginLogDao) Args(args ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.Args(args...)}
}

func (d *LoginLogDao) LeftJoin(table ...string) *LoginLogDao {
	return &LoginLogDao{M: d.M.LeftJoin(table...)}
}

func (d *LoginLogDao) RightJoin(table ...string) *LoginLogDao {
	return &LoginLogDao{M: d.M.RightJoin(table...)}
}

func (d *LoginLogDao) InnerJoin(table ...string) *LoginLogDao {
	return &LoginLogDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *LoginLogDao) Fields(fieldNamesOrMapStruct ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *LoginLogDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *LoginLogDao) Option(option int) *LoginLogDao {
	return &LoginLogDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *LoginLogDao) OmitEmpty() *LoginLogDao {
	return &LoginLogDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *LoginLogDao) Filter() *LoginLogDao {
	return &LoginLogDao{M: d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *LoginLogDao) Where(where interface{}, args ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *LoginLogDao) WherePri(where interface{}, args ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *LoginLogDao) And(where interface{}, args ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *LoginLogDao) Or(where interface{}, args ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *LoginLogDao) Group(groupBy string) *LoginLogDao {
	return &LoginLogDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *LoginLogDao) Order(orderBy ...string) *LoginLogDao {
	return &LoginLogDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *LoginLogDao) Limit(limit ...int) *LoginLogDao {
	return &LoginLogDao{M: d.M.Limit(limit...)}
}

func (d *LoginLogDao) Offset(offset int) *LoginLogDao {
	return &LoginLogDao{M: d.M.Offset(offset)}
}

func (d *LoginLogDao) Page(page, limit int) *LoginLogDao {
	return &LoginLogDao{M: d.M.Page(page, limit)}
}

func (d *LoginLogDao) Batch(batch int) *LoginLogDao {
	return &LoginLogDao{M: d.M.Batch(batch)}
}

func (d *LoginLogDao) Cache(duration time.Duration, name ...string) *LoginLogDao {
	return &LoginLogDao{M: d.M.Cache(duration, name...)}
}

func (d *LoginLogDao) Data(data ...interface{}) *LoginLogDao {
	return &LoginLogDao{M: d.M.Data(data...)}
}

func (d *LoginLogDao) All(where ...interface{}) ([]*model.LoginLog, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.LoginLog
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

func (d *LoginLogDao) One(where ...interface{}) (*model.LoginLog, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.LoginLog
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

func (d *LoginLogDao) FindOne(where ...interface{}) (*model.LoginLog, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.LoginLog
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *LoginLogDao) FindAll(where ...interface{}) ([]*model.LoginLog, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.LoginLog
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

func (d *LoginLogDao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

func (d *LoginLogDao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

func (d *LoginLogDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *LoginLogDao) Chunk(limit int, callback func(entities []*model.LoginLog, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.LoginLog
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

func (d *LoginLogDao) LockUpdate() *LoginLogDao {
	return &LoginLogDao{M: d.M.LockUpdate()}
}

func (d *LoginLogDao) LockShared() *LoginLogDao {
	return &LoginLogDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *LoginLogDao) Unscoped() *LoginLogDao {
	return &LoginLogDao{M: d.M.Unscoped()}
}
