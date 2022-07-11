package Redis

import "main.go/config/app_conf"

func Hash_add(key string, value interface{}) error {
	return goredis.HSet(goredis_ctx, app_conf.Project+":"+key, value).Err()
}

func Hash_slices(key string, value []interface{}) error {
	return goredis.HSet(goredis_ctx, app_conf.Project+":"+key, value).Err()
}

func Hash_list_keys(key string) ([]string, error) {
	return goredis.HKeys(goredis_ctx, app_conf.Project+":"+key).Result()
}

func Hash_get(key string, field string) (string, error) {
	return goredis.HGet(goredis_ctx, app_conf.Project+":"+key, field).Result()
}

func Hash_get_list(key string, fields ...string) ([]interface{}, error) {
	return goredis.HMGet(goredis_ctx, app_conf.Project+":"+key, fields...).Result()
}

func Hash_count(key string) int64 {
	count, err := goredis.HLen(goredis_ctx, app_conf.Project+":"+key).Result()
	if err != nil {
		return 0
	}
	return count
}

func Hash_map_set(key string, value map[string]interface{}) error {
	return Hash_add(app_conf.Project+":"+key, value)
}

func Hash_map_get(key string) map[string]string {
	ret, err := goredis.HGetAll(goredis_ctx, app_conf.Project+":"+key).Result()
	if err != nil {
		return nil
	}
	return ret
}
