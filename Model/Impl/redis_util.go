package Impl

import (
	"log"
	"strconv"
	"time"
)

func getBalanceRedis(walletId int) (float64, bool) {
	vals, err := rdb.Get(ctx, strconv.Itoa(walletId)).Result()
	if err != nil {
		log.Println("Cannot fetch from catche:", err)
		return -1, false
	}
	if len(vals) > 0 {
		val, err := strconv.ParseFloat(vals, 64)
		if err != nil {
			log.Println("Redis: Cannot convert to float", err)
			return -1, false
		}
		return val, true
	}
	return -1, false
}

func setBalanceRedis(walletId int, balance float64, expiry time.Duration) {
	_, err = rdb.Set(ctx, strconv.Itoa(walletId),
		strconv.FormatFloat(balance, 'f', 6, 64), expiry).Result()

	if err != nil {
		log.Println("Cannot set on cache:", err)
	}
}

func deleteBalanceRedis(walletId int) {
	_, err = rdb.Del(ctx, strconv.Itoa(walletId)).Result()
	if err != nil {
		log.Println("Cannot delete Redis key:", walletId, err)
	}
}
