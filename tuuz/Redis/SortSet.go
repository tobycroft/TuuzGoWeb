package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"main.go/config/app_conf"
)

func SortSet_Add(key string, score float64, value interface{}) error {
	var z redis.Z
	z.Score = score
	z.Member = value
	return goredis.ZAdd(context.Background(), app_conf.Project+":"+key, z).Err()
}

func SortSet_Count(key string, min, max interface{}) int64 {
	if min == nil || max == nil {
		count, err := goredis.ZCard(context.Background(), app_conf.Project+":"+key).Result()
		if err != nil {
			return 0
		}
		return count
	} else {
		minuim, ok := min.(string)
		if !ok {
			return 0
		}
		maxium, ok := max.(string)
		if !ok {
			return 0
		}
		count, err := goredis.ZCount(context.Background(), app_conf.Project+":"+key, minuim, maxium).Result()
		if err != nil {
			return 0
		}
		return count
	}
}

func SortSet_Increase(key string, incr float64, value string) (float64, error) {
	return goredis.ZIncrBy(context.Background(), app_conf.Project+":"+key, incr, value).Result()
}

func SortSet_rank(key, rank_by string) (int64, error) {
	return goredis.ZRank(context.Background(), app_conf.Project+":"+key, rank_by).Result()
}

func SortSet_list_asc(key string) ([]string, error) {
	return SortSet_list_asc_min_max(app_conf.Project+":"+key, 0, -1)
}

func SortSet_list_desc(key string) ([]string, error) {
	return SortSet_list_desc_min_max(app_conf.Project+":"+key, 0, -1)
}

func SortSet_list_asc_min_max(key string, start, end int64) ([]string, error) {
	return goredis.ZRange(context.Background(), app_conf.Project+":"+key, start, end).Result()
}

func SortSet_list_desc_min_max(key string, start, end int64) ([]string, error) {
	return goredis.ZRevRange(context.Background(), app_conf.Project+":"+key, start, end).Result()
}

func SortSet_search(key, search string, limit int) (ret []string, err error) {
	ret, _, err = goredis.ZScan(context.Background(), app_conf.Project+":"+key, 0, search, int64(limit)).Result()
	return
}
