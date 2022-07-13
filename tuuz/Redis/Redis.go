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

func ExpireTime(key string) (time.Duration, error) {
	return goredis.TTL(goredis_ctx, app_conf.Project+":"+key).Result()
}

func ExpireAt(key string, expire_at time.Time) error {
	return goredis.ExpireAt(goredis_ctx, app_conf.Project+":"+key, expire_at).Err()
}

func CheckExists(key string) bool {
	row, err := goredis.Exists(goredis_ctx, app_conf.Project+":"+key).Result()
	if err != nil {
		return false
	}
	return row > 0
}
