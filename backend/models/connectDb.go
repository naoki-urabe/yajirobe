package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
	"yajirobe/config"
)

var Db *sqlx.DB
var err error

func ConnectDb() {
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)
	for i := 0; i < 10; i++ {
		time.Sleep(30 * time.Second)
		Db, err = sqlx.Connect("mysql", dbConnectInfo)
		if err != nil {
			log.Fatalln(err)
		} else {
			break
		}
	}
	fmt.Println("Successfully connect database...")
}
