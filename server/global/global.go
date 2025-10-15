package global

import (
	"github.com/odboy-tianjun/kenaito-admin/server/config"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

var (
	// CONFIG 全局配置
	CONFIG *config.Config

	// DB 全局数据库连接
	DB *xorm.Engine

	// REDIS 全局Redis客户端
	REDIS *redis.Client
)
