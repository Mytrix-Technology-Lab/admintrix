// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/mytrix-technology/admintrix/internal/model"
	"time"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// MenuDao is the data access object for table sys_menu.
type MenuDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	Columns MenuColumns // columns contains all the column names of Table for convenient usage.
}

// MenuColumns defines and stores column names for table sys_menu.
type MenuColumns struct {
	Id          string // 主键ID
	ParentId    string // 父级ID
	Title       string // 菜单标题
	Icon        string // 图标
	Path        string // 菜单路径
	Component   string // 菜单组件
	Target      string // 打开方式：0组件 1内链 2外链
	Permission  string // 权限标识
	Type        string // 类型：0菜单 1节点
	Method      string // 请求方式
	Status      string // 状态：1正常 2禁用
	Hide        string // 是否可见：1是 2否
	Note        string // 备注
	Sort        string // 显示顺序
	CreateUser  string // 添加人
	CreateTime  string // 创建时间
	UpdateUser  string // 更新人
	UpdateTime  string // 更新时间
	Mark        string // 有效标识
}

// MenuColumns holds the columns for table sys_menu.
var (
	MenuColumnsX = MenuColumns{}
	Menu = MenuDao{
		M:     g.DB("default").Model("sys_menu").Safe(),
		DB:    g.DB("default"),
		Table: "sys_menu",
		Columns: MenuColumns{
			Id:         "id",
			ParentId:   "parent_id",
			Title:      "title",
			Icon:       "icon",
			Path:       "path",
			Component:  "component",
			Target:     "target",
			Permission: "permission",
			Type:       "type",
			Method:     "method",
			Status:     "status",
			Hide:       "hide",
			Note:       "note",
			Sort:       "sort",
			CreateUser: "create_user",
			CreateTime: "create_time",
			UpdateUser: "update_user",
			UpdateTime: "update_time",
			Mark:       "mark",
		},
	}
)

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao() *MenuDao {
	return &MenuDao{
		group:   "default",
		table:   "sys_menu",
		Columns: MenuColumnsX,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MenuDao) DBO() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
//func (dao *MenuDao) Table() string {
//	return dao.table
//}

// Columns returns all column names of current dao.
func (dao *MenuDao) Columnsx() MenuColumns {
	return dao.Columns
}

// Group returns the configuration group name of database of current dao.
//func (dao *MenuDao) Group() string {
//	return dao.group
//}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenuDao) CtxO(ctx context.Context) *gdb.Model {
	return dao.DBO().Model(dao.table).Safe().Ctx(ctx)
}
func (d *MenuDao) Ctx(ctx context.Context) *MenuDao {
	return &MenuDao{M: d.M.Ctx(ctx)}
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
//func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
//	return dao.Ctx(ctx).Transaction(ctx, f)
//}

// As sets an alias name for current table.
func (d *MenuDao) As(as string) *MenuDao {
	return &MenuDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *MenuDao) TX(tx *gdb.TX) *MenuDao {
	return &MenuDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *MenuDao) Master() *MenuDao {
	return &MenuDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *MenuDao) Slave() *MenuDao {
	return &MenuDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *MenuDao) Args(args ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MenuDao) LeftJoin(table ...string) *MenuDao {
	return &MenuDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MenuDao) RightJoin(table ...string) *MenuDao {
	return &MenuDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MenuDao) InnerJoin(table ...string) *MenuDao {
	return &MenuDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *MenuDao) Fields(fieldNamesOrMapStruct ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *MenuDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *MenuDao) Option(option int) *MenuDao {
	return &MenuDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *MenuDao) OmitEmpty() *MenuDao {
	return &MenuDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *MenuDao) Filter() *MenuDao {
	return &MenuDao{M: d.M.Filter()}
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
func (d *MenuDao) Where(where interface{}, args ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *MenuDao) WherePri(where interface{}, args ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *MenuDao) And(where interface{}, args ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *MenuDao) Or(where interface{}, args ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *MenuDao) Group(groupBy string) *MenuDao {
	return &MenuDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *MenuDao) Order(orderBy ...string) *MenuDao {
	return &MenuDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *MenuDao) Limit(limit ...int) *MenuDao {
	return &MenuDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *MenuDao) Offset(offset int) *MenuDao {
	return &MenuDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *MenuDao) Page(page, limit int) *MenuDao {
	return &MenuDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *MenuDao) Batch(batch int) *MenuDao {
	return &MenuDao{M: d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *MenuDao) Cache(duration time.Duration, name ...string) *MenuDao {
	return &MenuDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *MenuDao) Data(data ...interface{}) *MenuDao {
	return &MenuDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.Menu.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *MenuDao) All(where ...interface{}) ([]*model.Menu, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Menu
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.Menu.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *MenuDao) One(where ...interface{}) (*model.Menu, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Menu
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *MenuDao) FindOne(where ...interface{}) (*model.Menu, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Menu
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *MenuDao) FindAll(where ...interface{}) ([]*model.Menu, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Menu
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *MenuDao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *MenuDao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *MenuDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *MenuDao) Chunk(limit int, callback func(entities []*model.Menu, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.Menu
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *MenuDao) LockUpdate() *MenuDao {
	return &MenuDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *MenuDao) LockShared() *MenuDao {
	return &MenuDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *MenuDao) Unscoped() *MenuDao {
	return &MenuDao{M: d.M.Unscoped()}
}