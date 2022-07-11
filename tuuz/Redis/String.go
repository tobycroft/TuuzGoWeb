package Redis

import (
	"time"
)

func String_set(key string, value interface{}, exp time.Duration) error {
	return goredis.Set(goredis_ctx, key, value, exp).Err()
}

func String_get(key string) (string, error) {
	return goredis.Get(goredis_ctx, key).Result()
}

func String_getset(key string, value interface{}) (string, error) {
	return goredis.GetSet(goredis_ctx, key, value).Result()
}

func String_getInt(key string) (int, error) {
	return goredis.Get(goredis_ctx, key).Int()
}

func String_getInt64(key string) (int64, error) {
	return goredis.Get(goredis_ctx, key).Int64()
}

func String_getFloat64(key string) (float64, error) {
	return goredis.Get(goredis_ctx, key).Float64()
}

func String_getBytes(key string) ([]byte, error) {
	return goredis.Get(goredis_ctx, key).Bytes()
}

func String_getTime(key string) (time.Time, error) {
	return goredis.Get(goredis_ctx, key).Time()
}

func String_length(key string) (int64, error) {
	return goredis.StrLen(goredis_ctx, key).Result()
}

func String_Set_byMap(KV map[string]interface{}) error {
	return goredis.MSet(goredis_ctx, KV).Err()
}

func String_float64_incr(key string, incr float64) (float64, error) {
	return goredis.IncrByFloat(goredis_ctx, key, incr).Result()
}

func String_int64_incr(key string, incr int64) (int64, error) {
	return goredis.IncrBy(goredis_ctx, key, incr).Result()
}

func String_int64_decr(key string, decr int64) (int64, error) {
	return goredis.DecrBy(goredis_ctx, key, decr).Result()
}
