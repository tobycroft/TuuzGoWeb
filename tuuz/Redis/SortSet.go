package Redis

import "github.com/go-redis/redis/v8"

func ZADD(key string, score float64, value interface{}) error {
	var z redis.Z
	z.Score = score
	z.Member = value
	return goredis.ZAdd(goredis_ctx, key, &z).Err()
}

func ZCount(key string, min, max interface{}) int64 {
	if min == nil || max == nil {
		count, err := goredis.ZCard(goredis_ctx, key).Result()
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
		count, err := goredis.ZCount(goredis_ctx, key, minuim, maxium).Result()
		if err != nil {
			return 0
		}
		return count
	}

}
