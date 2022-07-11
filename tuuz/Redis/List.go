package Redis

import "main.go/config/app_conf"

func List_append(key string, value interface{}) error {
	return goredis.LPush(goredis_ctx, app_conf.Project+":"+key, value).Err()
}

func List_insert(key string, value interface{}) error {
	return goredis.RPush(goredis_ctx, app_conf.Project+":"+key, value).Err()
}

func List_count(key string) int64 {
	count, err := goredis.LLen(goredis_ctx, app_conf.Project+":"+key).Result()
	if err != nil {
		return 0
	}
	return count
}

func List_del_get_fist(key string) (string, error) {
	return goredis.LPop(goredis_ctx, app_conf.Project+":"+key).Result()
}

func List_del_get_last(key string) (string, error) {
	return goredis.RPop(goredis_ctx, app_conf.Project+":"+key).Result()
}

func List_del(key string, value interface{}) error {
	return goredis.LRem(goredis_ctx, app_conf.Project+":"+key, 0, value).Err()
}
