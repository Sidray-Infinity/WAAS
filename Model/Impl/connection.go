package Impl

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var address string = "root:root@tcp(127.0.0.1:3306)/waas?charset=utf8&parseTime=True&loc=Local"
var err error
var db *gorm.DB

var rdb *redis.Client
var ctx = context.Background()

var cronMutex *redsync.Mutex
var balanceCache *cache.Cache

func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	balanceCache = cache.New(&cache.Options{
		Redis:      rdb,
		LocalCache: cache.NewTinyLFU(10000, time.Hour),
	})

	// Initialize the mutex
	pool := goredis.NewPool(rdb)
	rs := redsync.New(pool)
	mutexname := "cron-mutex" // Should be common in all instances
	option := redsync.WithExpiry(15 * time.Second)
	cronMutex = rs.NewMutex(mutexname, option)
}

func CloseRedis() {
	err = rdb.Close()
	if err != nil {
		log.Println("Cannot close redis client:", err)
	}
}

func ConnnectToDB() {
	db, err = gorm.Open(mysql.Open(address), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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
