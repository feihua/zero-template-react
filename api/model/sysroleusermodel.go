package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleUserModel = (*customSysRoleUserModel)(nil)

type (
	// SysRoleUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleUserModel.
	SysRoleUserModel interface {
		sysRoleUserModel
	}

	customSysRoleUserModel struct {
		*defaultSysRoleUserModel
	}
)

// NewSysRoleUserModel returns a model for the database table.
func NewSysRoleUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleUserModel {
	return &customSysRoleUserModel{
		defaultSysRoleUserModel: newSysRoleUserModel(conn, c),
	}
}
