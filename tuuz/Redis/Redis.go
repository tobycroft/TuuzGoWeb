package Redis

import (
	"main.go/config/app_conf"
	"time"
)

func Del(key string) error {
	return goredis.Del(goredis_ctx, app_conf.Project+":"+key).Err()
}

func Expire(key string, duration time.Duration) error {
	return goredis.Expire(goredis_ctx, app_conf.Project+":"+key, duration).Err()
}

func WhenWillExpire(key string) (time.Duration, error) {
	return goredis.TTL(goredis_ctx, key).Result()
}
