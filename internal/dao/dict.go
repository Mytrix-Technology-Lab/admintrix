// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mytrix-technology/admintrix/internal/dao/internal"
)

// internalDictDao is internal type for wrapping internal DAO implements.
type internalDictDao = *internal.DictDao

// DictDao is the data access object for table sys_dict.
// You can define custom methods on it to extend its functionality as you wish.
type DictDao struct {
	internalDictDao
}

var (
	// Dict is globally public accessible object for table sys_dict operations.
	Dict = DictDao{
		internal.NewDictDao(),
	}
)

// Fill with you ideas below.