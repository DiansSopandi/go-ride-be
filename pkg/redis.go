package pkg

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

var (
	once         sync.Once
	redisClient  *redis.Client
	redisLimiter *redis_rate.Limiter
	rdb          *redis.Client
)

// InitRedis untuk inisialisasi Redis client & limiter sekali saja
func InitRedis() {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: Cfg.Redis.Host + ":" + fmt.Sprintf("%d", Cfg.Redis.Port), // "localhost:6379",
		})

		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("❌ Redis connection failed: %v", err)
		}
		log.Println("✅ Redis connected")

		redisLimiter = redis_rate.NewLimiter(redisClient)
	})
}

// GetRedisClient return redis client
func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitRedis()
	}
	return redisClient
}

// GetLimiter return rate limiter
func GetLimiter() *redis_rate.Limiter {
	if redisLimiter == nil {
		InitRedis()
	}
	return redisLimiter
}

func CloseRedis() error {
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}
