package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"main.go/config/app_conf"
)

type Stream struct {
	stream_channel string
	Producer       chan any
	Consumer       chan any
	group          string
	consumer       string
}

func StreamNew(stream_name string) Stream {
	return Stream{}.New(stream_name)
}

func (self Stream) New(stream_name string) Stream {
	self.stream_channel = app_conf.Project + ":" + stream_name
	//self.Producer = make(chan any)
	//self.Consumer = make(chan any)
	return self
}

func (self Stream) Publish(value map[string]any) (string, error) {
	return goredis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: self.stream_channel,
		Values: value,
	}).Result()
}

func (self Stream) XLength() (int64, error) {
	return goredis.XLen(context.Background(), self.stream_channel).Result()
}

func (self Stream) XRange() ([]redis.XMessage, error) {
	return goredis.XRange(context.Background(), self.stream_channel, "-", "+").Result()
}

func (self Stream) XRevRange() ([]redis.XMessage, error) {
	return goredis.XRevRange(context.Background(), self.stream_channel, "-", "+").Result()
}

func (self Stream) XRead() ([]redis.XStream, error) {
	var xr redis.XReadArgs
	return goredis.XRead(context.Background(), &xr).Result()
}

func (self Stream) XGroupCreateConsumer(group, consumer string) error {
	return goredis.XGroupCreateConsumer(context.Background(), self.stream_channel, group, consumer).Err()
}

func (self Stream) XReadGroup() ([]redis.XStream, error) {
	return goredis.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Group:    self.group,
		Consumer: self.consumer,
		Streams:  []string{self.stream_channel},
		Count:    1,
		Block:    0,
		NoAck:    false,
	}).Result()
}
