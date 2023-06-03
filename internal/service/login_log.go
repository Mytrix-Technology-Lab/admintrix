package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/mytrix-technology/admintrix/internal/dao"
	"github.com/mytrix-technology/admintrix/internal/model"
	"github.com/mytrix-technology/admintrix/pkg/utils"
	"github.com/mytrix-technology/admintrix/pkg/utils/convert"
)

var LoginLog = new(loginLogService)

type loginLogService struct{}

func (s *loginLogService) GetList(req *model.LoginLogPageReq) ([]model.LoginLog, int, error) {
	query := dao.LoginLog.Clone()
	query = query.Where("mark=1")

	if req != nil {
		if req.Username != "" {
			query = query.Where("username=?", req.Username)
		}
	}

	count, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	query = query.Order("id desc")

	query = query.Page(req.Page, req.Limit)

	var list []model.LoginLog
	query.Structs(&list)
	return list, count, nil
}

func (s *loginLogService) Delete(Ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, gerror.New("演示环境，暂无权限操作")
	}
	idsArr := convert.ToInt64Array(Ids, ",") //nolint:typecheck
	result, err := dao.LoginLog.Delete("id in (?)", idsArr)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	return rows, nil
}
