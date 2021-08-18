package cmd

import (
	"vandacare.com/api/v1/auth"
	"vandacare.com/api/v1/user"
	"vandacare.com/core/cache"
	"vandacare.com/core/config"
	"vandacare.com/core/database"
	"vandacare.com/core/event"
	"vandacare.com/core/log"
	"vandacare.com/core/router"
)

func Run() {
	config.LoadConfig("config.toml")
	log.ConfigLogger()
	cache.ConfigCache()
	database.InitMySQL()
	event.Subscribe(user.Subscribe, auth.Subscribe)
	r := router.InitRouter()
	router.InitPublicRouter(r, auth.Routers)
	router.InitAuthRouter(r, user.Routers)
	router.RunServer(r)
}
