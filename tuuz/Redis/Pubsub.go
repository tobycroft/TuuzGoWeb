package Redis

import "github.com/go-redis/redis/v8"

func Suscribe(channel string) {

	goredis.Subscribe(goredis_ctx, channel).Channel(
		redis.WithChannelSize(100),
	)

}
