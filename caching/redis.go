package caching

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService() *RedisService {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Redis:", pong)

	return &RedisService{client: client}
}

func (rs *RedisService) Set(key string, value string, expiration time.Duration) error {
	err := rs.client.Set(context.Background(), key, value, expiration).Err()
	return err
}

func (rs *RedisService) Get(key string) (string, error) {
	value, err := rs.client.Get(context.Background(), key).Result()
	return value, err
}

func (rs *RedisService) Exists(key string) (bool, error) {
	exists, err := rs.client.Exists(context.Background(), key).Result()
	return exists == 1, err
}

func (rs *RedisService) Close() error {
	err := rs.client.Close()
	return err
}
