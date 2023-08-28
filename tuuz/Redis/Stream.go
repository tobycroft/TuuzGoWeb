package Redis

import (
	"context"
	"main.go/config/app_conf"
)

func Stream_publish(stream_key string, value interface{}) error {
	var xa redis.XAddArgs
	xa.Stream = app_conf.Project + ":" + stream_key
	xa.Values = value
	return goredis.XAdd(context.Background(), &xa).Err()
}
