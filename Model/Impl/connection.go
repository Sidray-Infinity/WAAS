package Impl

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var address string = "root:root@tcp(127.0.0.1:3306)/waas?charset=utf8&parseTime=True&loc=Local"
var err error
var db *gorm.DB

func ConnnectToDB() {
	db, err = gorm.Open(mysql.Open(address), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot Connect to DB", err)
		os.Exit(0)
	}
}

func CloseDB() {
	sqlDb, err := db.DB()
	if err != nil {
		log.Println("Cannot extract SQL db")
		return
	}
	if err = sqlDb.Ping(); err != nil {
		sqlDb.Close()
	} else {
		log.Println("Cannot close DB Connection:", err)
	}
}
