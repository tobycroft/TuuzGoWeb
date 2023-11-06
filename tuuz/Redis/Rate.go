package Redis

import (
	"context"
	"github.com/go-redis/redis_rate/v10"
	"main.go/config/app_conf"
)

var RedisRateLimiter *redis_rate.Limiter

func newLimitter() {
	RedisRateLimiter = redis_rate.NewLimiter(goredis)
}

func RedisRateLimit_second(channel string, ratio int) (*redis_rate.Result, error) {
	return RedisRateLimiter.Allow(context.Background(), app_conf.Project+":"+channel, redis_rate.PerSecond(ratio))
}

func RedisRateLimit_minute(channel string, ratio int) (*redis_rate.Result, error) {
	return RedisRateLimiter.Allow(context.Background(), app_conf.Project+":"+channel, redis_rate.PerMinute(ratio))
}

func RedisRateLimit_hour(channel string, ratio int) (*redis_rate.Result, error) {
	return RedisRateLimiter.Allow(context.Background(), app_conf.Project+":"+channel, redis_rate.PerHour(ratio))
}
