package lib

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"golang-CICD/config"
)

var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Conf.Redis.Host, config.Conf.Redis.Port),
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to ping redis, err:" + err.Error())
	}
}

func RedisClient() *redis.Client {
	return client
}
