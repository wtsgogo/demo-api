package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx context.Context
var rdb *redis.Client

func init() {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func SetValue(key, value string, second int) error {
	err := rdb.Set(ctx, key, value, time.Duration(second)*time.Second).Err()
	if err != nil {
		log.Println("无法设置缓存信息:", err.Error())
		return err
	}
	return nil
}

func GetValue(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("缓存信息不存在或已过期:", key)
		return "", nil
	} else if err != nil {
		log.Println("无法读取缓存信息:", err.Error())
		return "", err
	}
	return val, nil
}

func DelValue(key string) {
	rdb.Del(ctx, key)
}
