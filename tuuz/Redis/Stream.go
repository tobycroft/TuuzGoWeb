package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"main.go/config/app_conf"
)

type Stream struct {
	StreamChannelName string
	//Producer          chan any
	//Consumer          chan any
	Group    string
	consumer string
}

type stream struct {
	Stream
}

func StreamNew(stream_name string) stream {
	return stream{}.New(stream_name)
}

func (self Stream) New(stream_name string) stream {
	self.StreamChannelName = app_conf.Project + ":" + stream_name
	//self.Producer = make(chan any)
	//self.Consumer = make(chan any)
	return stream{self}
}

func (self stream) Publish(value map[string]any) (string, error) {
	return goredis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: self.StreamChannelName,
		Values: value,
	}).Result()
}

func (self stream) XLength() (int64, error) {
	return goredis.XLen(context.Background(), self.StreamChannelName).Result()
}

func (self stream) XRange() ([]redis.XMessage, error) {
	return goredis.XRange(context.Background(), self.StreamChannelName, "-", "+").Result()
}

func (self stream) XRevRange() ([]redis.XMessage, error) {
	return goredis.XRevRange(context.Background(), self.StreamChannelName, "-", "+").Result()
}

func (self stream) XRead() ([]redis.XStream, error) {
	var xr redis.XReadArgs
	return goredis.XRead(context.Background(), &xr).Result()
}

func (self stream) XGroupCreateConsumer(group, consumer string) error {
	return goredis.XGroupCreateConsumer(context.Background(), self.StreamChannelName, group, consumer).Err()
}

func (self stream) XReadGroup() ([]redis.XStream, error) {
	return goredis.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Group:    self.Group,
		Consumer: self.consumer,
		Streams:  []string{self.StreamChannelName},
		Count:    1,
		Block:    0,
		NoAck:    false,
	}).Result()
}
