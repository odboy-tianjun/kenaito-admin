package main

import (
	"github.com/odboy-tianjun/kenaito-admin/server/config"
	"github.com/odboy-tianjun/kenaito-admin/server/global"
)

func main() {
	// 加载配置
	global.CONFIG = config.LoadConfig(".")
	// TODO 连接redis
	// TODO 连接mysql
	// TODO 注册gin路由
	// TODO 启动gin服务
}
