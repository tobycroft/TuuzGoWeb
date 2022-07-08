package Redis

import (
	"github.com/shopspring/decimal"
	"main.go/config/app_conf"
	"time"
)

func SAdd(key string, value any, duration time.Duration) error {
	return goredis.SAdd(goredis_ctx, app_conf.Project+":"+key, value, time.Second*duration).Err()
}

func SAdds[T []string | []any | []int64 | []int | []decimal.Decimal | []float64](key string, value T, duration time.Duration) error {
	return goredis.SAdd(goredis_ctx, app_conf.Project+":"+key, value, time.Second*duration).Err()
}

func SMember(key string) ([]string, error) {
	return goredis.SMembers(goredis_ctx, app_conf.Project+":"+key).Result()
}

func SMemberMap(key string) (map[string]struct{}, error) {
	return goredis.SMembersMap(goredis_ctx, app_conf.Project+":"+key).Result()
}

func IsMember(key string, value any) bool {
	has, err := goredis.SIsMember(goredis_ctx, app_conf.Project+":"+key, value).Result()
	if err != nil {
		return false
	}
	return has
}

func IsMembers(key string, value ...any) []bool {
	has, err := goredis.SMIsMember(goredis_ctx, app_conf.Project+":"+key, value).Result()
	if err != nil {
		return []bool{}
	}
	return has
}
