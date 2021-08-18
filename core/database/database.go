package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"vandacare.com/core/config"
)

func InitMySQL() *sqlx.DB {
	host := config.ReadConfig("database.host")
	user := config.ReadConfig("database.user")
	password := config.ReadConfig("database.password")
	port := config.ReadConfig("database.port")
	dbname := config.ReadConfig("database.dbname")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return db
}
