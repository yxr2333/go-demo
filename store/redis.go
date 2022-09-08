package store

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func InitRedisClient() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       6,
		PoolSize: 20,
	})

	_, err = RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
