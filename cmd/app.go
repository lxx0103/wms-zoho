package cmd

import (
	"wms.com/api/v1/auth"
	"wms.com/api/v1/setting"
	"wms.com/api/v1/user"
	"wms.com/core/cache"
	"wms.com/core/config"
	"wms.com/core/database"
	"wms.com/core/event"
	"wms.com/core/log"
	"wms.com/core/router"
)

func Run() {
	config.LoadConfig("config.toml")
	log.ConfigLogger()
	cache.ConfigCache()
	database.ConfigMysql()
	event.Subscribe(user.Subscribe, auth.Subscribe)
	r := router.InitRouter()
	router.InitPublicRouter(r, auth.Routers)
	router.InitAuthRouter(r, user.Routers, setting.Routers)
	router.RunServer(r)
}
