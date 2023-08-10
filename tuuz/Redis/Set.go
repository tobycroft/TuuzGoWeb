package Redis

import (
	"context"
	"github.com/shopspring/decimal"
	"main.go/config/app_conf"
	"time"
)

func Set_add(key string, value any) error {
	return goredis.SAdd(context.Background(), app_conf.Project+":"+key, value).Err()
}

func Set_count(key string) int64 {
	count, err := goredis.SCard(context.Background(), app_conf.Project+":"+key).Result()
	if err != nil {
		return 0
	}
	return count
}

func Set_add_more[T []string | []any | []int64 | []int | []decimal.Decimal | []float64](key string, value T, duration time.Duration) error {
	return goredis.SAdd(context.Background(), app_conf.Project+":"+key, value, time.Second*duration).Err()
}

func Set_list(key string) ([]string, error) {
	return goredis.SMembers(context.Background(), app_conf.Project+":"+key).Result()
}

func Set_list_map(key string) (map[string]struct{}, error) {
	return goredis.SMembersMap(context.Background(), app_conf.Project+":"+key).Result()
}

func Set_isMember(key string, value any) bool {
	has, err := goredis.SIsMember(context.Background(), app_conf.Project+":"+key, value).Result()
	if err != nil {
		return false
	}
	return has
}

func Set_isMembers(key string, value ...any) []bool {
	has, err := goredis.SMIsMember(context.Background(), app_conf.Project+":"+key, value).Result()
	if err != nil {
		return []bool{}
	}
	return has
}

func Set_search(key string, search string, limit int) ([]string, error) {
	strs, _, err := goredis.SScan(context.Background(), app_conf.Project+":"+key, 0, search, int64(limit)).Result()
	if err != nil {
		return strs, err
	}
	return strs, err
}

func Set_del_get(key string, value interface{}) (string, error) {
	return goredis.SPop(context.Background(), app_conf.Project+":"+key).Result()
}

func Set_same_with_others(set1_key string, set2_key string) ([]string, error) {
	return goredis.SInter(context.Background(), app_conf.Project+":"+set1_key, app_conf.Project+":"+set2_key).Result()
}

func Set_diff_with_others(set1_key, set2_key string) ([]string, error) {
	return goredis.SDiff(context.Background(), app_conf.Project+":"+set1_key, app_conf.Project+":"+set2_key).Result()
}

func Set_store_same_with_others(dest_key, set1_key, set2_key string) error {
	return goredis.SInterStore(context.Background(), app_conf.Project+":"+set1_key, app_conf.Project+":"+set2_key).Err()
}

func Set_store_diff_with_others(dest_key, set1_key, set2_key string) error {
	return goredis.SDiffStore(context.Background(), app_conf.Project+":"+dest_key, app_conf.Project+":"+set1_key, app_conf.Project+":"+set2_key).Err()
}

func Set_del_get_random(key string) (string, error) {
	return goredis.SPop(context.Background(), app_conf.Project+":"+key).Result()
}

func Set_del(key string, value ...interface{}) error {
	return goredis.SRem(context.Background(), app_conf.Project+":"+key, value).Err()
}

func Set_combine_with_others(set1_key, set2_key string) ([]string, error) {
	return goredis.SUnion(context.Background(), app_conf.Project+":"+set1_key, app_conf.Project+":"+set2_key).Result()
}

func Set_store_combine_with_others(dest_key, set1_key, set2_key string) error {
	return goredis.SUnionStore(context.Background(), app_conf.Project+":"+dest_key, app_conf.Project+":"+set1_key, app_conf.Project+":"+set2_key).Err()
}
