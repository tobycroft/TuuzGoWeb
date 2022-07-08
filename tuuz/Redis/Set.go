package Redis

import (
	"main.go/config/app_conf"
	"time"
)

func SAdd(key string, value string, duration time.Duration) error {
	return goredis.SAdd(goredis_ctx, app_conf.Project+":"+key, value, time.Second*duration).Err()
}

func SMember(key string) ([]string, error) {
	return goredis.SMembers(goredis_ctx, app_conf.Project+":"+key).Result()
}

func SMemberMap(key string) (map[string]struct{}, error) {
	return goredis.SMembersMap(goredis_ctx, app_conf.Project+":"+key).Result()
}

func IsMember(key, value string) (bool, error) {
	return goredis.SIsMember(goredis_ctx, app_conf.Project+":"+key, value).Result()

}
