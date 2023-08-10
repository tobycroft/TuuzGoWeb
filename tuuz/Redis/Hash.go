package Redis

import (
	"context"
	"errors"
	"github.com/tobycroft/gorose-pro"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

func Hash_add(key string, field, value any) error {
	return goredis.HSet(context.Background(), app_conf.Project+":"+key, field, value).Err()
}

func Hash_add_more[T map[string]string | map[string]any | gorose.Data](key string, maps T) error {
	return goredis.HSet(context.Background(), app_conf.Project+":"+key, maps).Err()
}

func Hash_field_exist(key string, field string) bool {
	ok, err := goredis.HExists(context.Background(), app_conf.Project+":"+key, field).Result()
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return false
	}
	return ok
}

func Hash_field_incr[T int64 | float64](key string, field string, incr_num T) (T, error) {
	switch any(incr_num).(type) {
	case int64:
		res, err := goredis.HIncrBy(context.Background(), app_conf.Project+":"+key, field, int64(incr_num)).Result()
		return T(res), err

	case float64:
		res, err := goredis.HIncrByFloat(context.Background(), app_conf.Project+":"+key, field, float64(incr_num)).Result()
		return T(res), err
	}
	return 0, errors.New("类型不匹配")
}

func Hash_list_keys(key string) ([]string, error) {
	return goredis.HKeys(context.Background(), app_conf.Project+":"+key).Result()
}

func Hash_list_values(key string) ([]string, error) {
	return goredis.HVals(context.Background(), app_conf.Project+":"+key).Result()
}

func Hash_get(key string, field string) (string, error) {
	return goredis.HGet(context.Background(), app_conf.Project+":"+key, field).Result()
}

func Hash_count(key string) int64 {
	count, err := goredis.HLen(context.Background(), app_conf.Project+":"+key).Result()
	if err != nil {
		return 0
	}
	return count
}

func Hash_all(key string) (map[string]string, error) {
	data, err := goredis.HGetAll(context.Background(), app_conf.Project+":"+key).Result()
	if err != nil {
		return nil, err
	}
	if len(data) < 1 {
		return nil, errors.New("empty")
	}
	return data, err
}

func Hash_delete(key string, field string) error {
	return goredis.HDel(context.Background(), app_conf.Project+":"+key, field).Err()
}

func Hash_search(key string, cursor uint64, search_pattern string, count int64) ([]string, uint64, error) {
	return goredis.HScan(context.Background(), app_conf.Project+":"+key, cursor, search_pattern, count).Result()
}
