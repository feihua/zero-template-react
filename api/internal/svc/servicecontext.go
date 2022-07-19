package svc

import (
	"zero-template-react/api/internal/config"
	"zero-template-react/api/internal/middleware"
	"zero-template-react/api/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.SysUserModel
	RoleModel     model.SysRoleModel
	RoleUserModel model.SysRoleUserModel
	MenuModel     model.SysMenuModel
	MenuRoleModel model.SysMenuRoleModel
	Redis         *redis.Redis
	CheckUrl      rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	return &ServiceContext{
		Config: c,

		UserModel:     model.NewSysUserModel(sqlConn, c.CacheRedis),
		RoleModel:     model.NewSysRoleModel(sqlConn, c.CacheRedis),
		RoleUserModel: model.NewSysRoleUserModel(sqlConn, c.CacheRedis),
		MenuModel:     model.NewSysMenuModel(sqlConn, c.CacheRedis),
		MenuRoleModel: model.NewSysMenuRoleModel(sqlConn, c.CacheRedis),
		Redis:         newRedis,
		CheckUrl:      middleware.NewCheckUrlMiddleware(newRedis).Handle,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
