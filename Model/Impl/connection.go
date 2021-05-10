package Impl

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"sync"

	"github.com/jinzhu/gorm"
)

var address string = "root:root@tcp(127.0.0.1:3306)/waas?charset=utf8&parseTime=True&loc=Local"
var err error
var walletMutx = make(map[int]*sync.Mutex)
var db *gorm.DB

func ConnnectToDB() {
	db, err = gorm.Open("mysql", address)
	if err != nil {
		fmt.Println("Cannot Connect to DB", err)
		os.Exit(0)
	}
}
