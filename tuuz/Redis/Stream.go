package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"main.go/config/app_conf"
)

func Stream_publish(stream_channel string, value any) error {
	var xa redis.XAddArgs
	xa.Stream = app_conf.Project + ":" + stream_channel
	xa.Values = value
	return goredis.XAdd(context.Background(), &xa).Err()
}

func Stream_xLength(stream_channel string) (int64, error) {
	return goredis.XLen(context.Background(), app_conf.Project+":"+stream_channel).Result()
}

func Stream_xRange(stream_key string) ([]redis.XMessage, error) {
	return goredis.XRange(context.Background(), app_conf.Project+":"+stream_key, "-", "+").Result()
}
