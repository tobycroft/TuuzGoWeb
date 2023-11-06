package Redis

import (
	"context"
	"main.go/config/app_conf"
)

func List_append(key string, value interface{}) error {
	return goredis.LPush(context.Background(), app_conf.Project+":"+key, value).Err()
}

func List_insert(key string, value interface{}) error {
	return goredis.RPush(context.Background(), app_conf.Project+":"+key, value).Err()
}

func List_count(key string) int64 {
	count, err := goredis.LLen(context.Background(), app_conf.Project+":"+key).Result()
	if err != nil {
		return 0
	}
	return count
}

func List_del_get_fist(key string) (string, error) {
	return goredis.LPop(context.Background(), app_conf.Project+":"+key).Result()
}

func List_del_get_last(key string) (string, error) {
	return goredis.RPop(context.Background(), app_conf.Project+":"+key).Result()
}

func List_del(key string, value interface{}) error {
	return goredis.LRem(context.Background(), app_conf.Project+":"+key, 0, value).Err()
}
