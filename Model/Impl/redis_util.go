package Impl

import (
	"log"
	"strconv"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func getBalanceRedis(walletId int) (float64, bool) {
	var balance float64
	err = balanceCache.Get(ctx, strconv.Itoa(walletId), &balance)
	if err != nil {
		if err == redis.Nil {
			log.Println("Redis: Cache miss")
		} else {
			log.Println("Redis: Cannot fetch from cache:", err)
		}

		return -1, false
	}
	return balance, true
}

func setBalanceRedis(walletId int, balance float64, expiry time.Duration) {
	err := balanceCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   strconv.Itoa(walletId),
		Value: balance,
	})
	if err != nil {
		log.Println("Redis: Cannot set on cache:", err)
	}
}

func deleteBalanceRedis(walletId int) {
	_, err = rdb.Del(ctx, strconv.Itoa(walletId)).Result()
	if err != nil {
		log.Println("Redis: Cannot delete Redis key:", walletId, err)
	}
}
