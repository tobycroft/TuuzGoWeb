package Redis

import (
	"context"
	"main.go/config/app_conf"
	"time"
)

func String_set(key string, value interface{}, exp time.Duration) error {
	return goredis.Set(context.Background(), app_conf.Project+":"+key, value, exp).Err()
}

func String_get(key string) (string, error) {
	return goredis.Get(context.Background(), app_conf.Project+":"+key).Result()
}

func String_getset(key string, value interface{}) (string, error) {
	return goredis.GetSet(context.Background(), app_conf.Project+":"+key, value).Result()
}

func String_getInt(key string) (int, error) {
	return goredis.Get(context.Background(), app_conf.Project+":"+key).Int()
}

func String_getInt64(key string) (int64, error) {
	return goredis.Get(context.Background(), app_conf.Project+":"+key).Int64()
}

func String_getFloat64(key string) (float64, error) {
	return goredis.Get(context.Background(), app_conf.Project+":"+key).Float64()
}

func String_getBytes(key string) ([]byte, error) {
	return goredis.Get(context.Background(), app_conf.Project+":"+key).Bytes()
}

func String_getBool(key string) (bool, error) {
	return goredis.Get(context.Background(), app_conf.Project+":"+key).Bool()
}

func String_getTime(key string) (time.Time, error) {
	return goredis.Get(context.Background(), app_conf.Project+":"+key).Time()
}

func String_length(key string) (int64, error) {
	return goredis.StrLen(context.Background(), app_conf.Project+":"+key).Result()
}

func String_float64_incr(key string, incr float64) (float64, error) {
	return goredis.IncrByFloat(context.Background(), app_conf.Project+":"+key, incr).Result()
}

func String_int64_incr(key string, incr int64) (int64, error) {
	return goredis.IncrBy(context.Background(), app_conf.Project+":"+key, incr).Result()
}

func String_int64_decr(key string, decr int64) (int64, error) {
	return goredis.DecrBy(context.Background(), app_conf.Project+":"+key, decr).Result()
}
