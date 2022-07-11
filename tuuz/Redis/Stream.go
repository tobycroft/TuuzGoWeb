package Redis

func Stream(stream_key string) {
	goredis.XInfoStream(goredis_ctx, stream_key).Result()
}
