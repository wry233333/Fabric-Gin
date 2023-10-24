package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConn struct {
	ctx context.Context
	rdb *redis.Client
}

func InitRedis() *RedisConn {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123654", // no password set
		DB:       0,        // use default DB
	})

	redisConn := &RedisConn{
		ctx: ctx,
		rdb: rdb,
	}

	return redisConn
}

func (r *RedisConn) Set(key any, value any) bool {
	err := r.rdb.Set(r.ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	return true
}
