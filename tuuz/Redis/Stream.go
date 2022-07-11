package Redis

import "main.go/config/app_conf"

func Stream(stream_key string) {
	goredis.XInfoStream(goredis_ctx, app_conf.Project+":"+stream_key).Result()
}
