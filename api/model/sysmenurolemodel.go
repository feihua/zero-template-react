package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysMenuRoleModel = (*customSysMenuRoleModel)(nil)

type (
	// SysMenuRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuRoleModel.
	SysMenuRoleModel interface {
		sysMenuRoleModel
	}

	customSysMenuRoleModel struct {
		*defaultSysMenuRoleModel
	}
)

// NewSysMenuRoleModel returns a model for the database table.
func NewSysMenuRoleModel(conn sqlx.SqlConn, c cache.CacheConf) SysMenuRoleModel {
	return &customSysMenuRoleModel{
		defaultSysMenuRoleModel: newSysMenuRoleModel(conn, c),
	}
}
