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
	MaxLen   int64
}

type stream struct {
	*Stream
}
type group struct {
	*stream
}

func (self *Stream) New(stream_name string) *stream {
	self.StreamChannelName = app_conf.Project + ":" + stream_name
	self.MaxLen = 10000
	//self.Producer = make(chan any)
	//self.Consumer = make(chan any)
	return &stream{self}
}

func (self *stream) SetMaxLen(MaxLen int64) *stream {
	self.MaxLen = MaxLen
	return self
}

func (self *stream) Publish(value map[string]any) (string, error) {
	return goredis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: self.StreamChannelName,
		Values: value,
		MaxLen: self.MaxLen,
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
	if err == nil {
		for _, gp := range gps {
			if gp.Name == group {
				return nil
			}
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

func (self *stream) XGroupDestroy() error {
	return goredis.XGroupDestroy(context.Background(), self.StreamChannelName, self.Group).Err()
}

func (self *stream) XGroupDelConsumer() error {
	return goredis.XGroupDelConsumer(context.Background(), self.StreamChannelName, self.Group, self.Consumer).Err()
}

func (self *stream) XInfoGroups() ([]redis.XInfoGroup, error) {
	return goredis.XInfoGroups(context.Background(), self.StreamChannelName).Result()
}

func (self *stream) XInfoConsumers(group string) ([]redis.XInfoConsumer, error) {
	self.Group = group
	return goredis.XInfoConsumers(context.Background(), self.StreamChannelName, self.Group).Result()
}

func (self *stream) XReadGroup() (redis.XStream, error) {
	data, err := goredis.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Group:    self.Group,
		Consumer: self.Consumer,
		Streams:  []string{self.StreamChannelName, ">"},
		Count:    1,
		Block:    0,
		NoAck:    false,
	}).Result()
	if err != nil {
		return redis.XStream{}, err
	}
	return data[0], err
}

func (self *stream) XReadGroupMore(count int64) ([]redis.XStream, error) {
	return goredis.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Group:    self.Group,
		Consumer: self.Consumer,
		Streams:  []string{self.StreamChannelName, ">"},
		Count:    count,
		Block:    0,
		NoAck:    false,
	}).Result()
}

func (self *stream) XPending() (*redis.XPending, error) {
	return goredis.XPending(context.Background(), self.StreamChannelName, self.Group).Result()
}

func (self *stream) XTrim(MaxLen int64) error {
	self.MaxLen = MaxLen
	return goredis.XTrimMaxLen(context.Background(), self.StreamChannelName, self.MaxLen).Err()
}

func (self *stream) XAck(ids ...string) error {
	return goredis.XAck(context.Background(), self.StreamChannelName, self.Group, ids...).Err()
}

func (self *stream) XDel(ids ...string) error {
	return goredis.XDel(context.Background(), self.StreamChannelName, ids...).Err()
}
