package Redis

import (
	"time"
)

func Set(key string, value interface{}, exp time.Duration) error {
	return goredis.Set(goredis_ctx, key, value, exp).Err()
}

func Get(key string) (string, error) {
	return goredis.Get(goredis_ctx, key).Result()
}

func GetInt(key string) (int, error) {
	return goredis.Get(goredis_ctx, key).Int()
}

func GetInt64(key string) (int64, error) {
	return goredis.Get(goredis_ctx, key).Int64()
}

func GetFloat64(key string) (float64, error) {
	return goredis.Get(goredis_ctx, key).Float64()
}

func GetBytes(key string) ([]byte, error) {
	return goredis.Get(goredis_ctx, key).Bytes()
}

func GetTime(key string) (time.Time, error) {
	return goredis.Get(goredis_ctx, key).Time()
}
