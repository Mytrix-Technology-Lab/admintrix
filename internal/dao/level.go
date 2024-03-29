// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/mytrix-technology/admintrix/internal/dao/internal"
)

// internalLevelDao is internal type for wrapping internal DAO implements.
type internalLevelDao = *internal.LevelDao

// LevelDao is the data access object for table sys_level.
// You can define custom methods on it to extend its functionality as you wish.
type LevelDao struct {
	internalLevelDao
}

var (
	// Level is globally public accessible object for table sys_level operations.
	Level = LevelDao{
		internal.NewLevelDao(),
	}
)

// Fill with you ideas below.
