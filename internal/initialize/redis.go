package initialize

import (
	"context"
	"fmt"
	"simplebank/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()
func InitRedis() {

	// Load Redis configuration from yaml file
	r := global.Config.Redis

	rdb  := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis: %v", zap.Error(err))
	}

	fmt.Println("Connected to Redis successfully")
	global.Rdb = rdb

	setRedisExample()
			
}

func setRedisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		global.Logger.Error("Failed to set key in Redis: %v", zap.Error(err))
	}

	val, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		global.Logger.Error("Failed to get key from Redis: %v", zap.Error(err))
	}
	global.Logger.Info("Value of 'score' in Redis: %s", zap.String("score", val))
}

