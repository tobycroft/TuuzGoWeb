package Redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"main.go/config/app_conf"
)

type Pubsub[T string | int | int32 | int64 | float32 | float64 | byte] struct {
}

func (self *Pubsub[T]) Subscribe(channel string) <-chan *redis.Message {
	return goredis.Subscribe(context.Background(), app_conf.Project+":"+channel).Channel()
}

func (self *Pubsub[T]) Publish(channel string, message any) error {
	return goredis.Publish(context.Background(), app_conf.Project+":"+channel, message).Err()
}

func (self *Pubsub[T]) Publish_struct(channel string, message_struct any) error {
	message, err := sonic.MarshalString(message_struct)
	if err != nil {
		return err
	}
	return goredis.Publish(context.Background(), app_conf.Project+":"+channel, message).Err()
}
