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
	Consumer string
}

type stream struct {
	*Stream
}
type group struct {
	*stream
}

func (self *Stream) New(stream_name string) *stream {
	self.StreamChannelName = app_conf.Project + ":" + stream_name
	//self.Producer = make(chan any)
	//self.Consumer = make(chan any)
	return &stream{self}
}

func (self *stream) Publish(value map[string]any) (string, error) {
	return goredis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: self.StreamChannelName,
		Values: value,
	}).Result()
}

func (self *stream) XLength() (int64, error) {
	return goredis.XLen(context.Background(), self.StreamChannelName).Result()
}

func (self *stream) XRange() ([]redis.XMessage, error) {
	return goredis.XRange(context.Background(), self.StreamChannelName, "-", "+").Result()
}

func (self *stream) XRevRange() ([]redis.XMessage, error) {
	return goredis.XRevRange(context.Background(), self.StreamChannelName, "-", "+").Result()
}

func (self *stream) XRead() ([]redis.XStream, error) {
	var xr redis.XReadArgs
	return goredis.XRead(context.Background(), &xr).Result()
}

func (self *stream) XGroupCreate(group string, start string) error {
	self.Group = group
	gps, err := self.XInfoGroups()
	if err != nil {
		return err
	}
	for _, gp := range gps {
		if gp.Name == group {
			return nil
		}
	}
	return goredis.XGroupCreateMkStream(context.Background(), self.StreamChannelName, self.Group, start).Err()
}

func (self *stream) XGroupCreateConsumer(Consumer string) error {
	self.Consumer = Consumer
	coms, err := self.XInfoConsumers(self.Group)
	if err != nil {
		return err
	}
	for _, com := range coms {
		if com.Name == Consumer {
			return nil
		}
	}
	return goredis.XGroupCreateConsumer(context.Background(), self.StreamChannelName, self.Group, Consumer).Err()
}

func (self *stream) XDelete() error {
	return goredis.XGroupDestroy(context.Background(), self.StreamChannelName, self.Group).Err()
}
func (self *stream) XInfoGroups() ([]redis.XInfoGroup, error) {
	return goredis.XInfoGroups(context.Background(), self.StreamChannelName).Result()
}

func (self *stream) XInfoConsumers(group string) ([]redis.XInfoConsumer, error) {
	self.Group = group
	return goredis.XInfoConsumers(context.Background(), self.StreamChannelName, self.Group).Result()
}

func (self *stream) XReadGroup() ([]redis.XStream, error) {
	//self.Group = Group
	//self.Consumer = Consumer
	return goredis.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Group:    self.Group,
		Consumer: self.Consumer,
		Streams:  []string{self.StreamChannelName},
		Count:    1,
		Block:    0,
		NoAck:    false,
	}).Result()
}
