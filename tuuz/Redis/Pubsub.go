package Redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"main.go/config/app_conf"
)

type Pubsub struct {
}

func (self *Pubsub) Subscribe(channel string) <-chan *redis.Message {
	return goredis.Subscribe(context.Background(), app_conf.Project+":"+channel).Channel()
}
func (self *Pubsub) Publish_string(channel string, message string) error {
	return goredis.Publish(context.Background(), app_conf.Project+":"+channel, message).Err()
}

func (self *Pubsub) Publish_struct(channel string, message_struct any) error {
	message, err := sonic.MarshalString(message_struct)
	if err != nil {
		return err
	}
	return goredis.Publish(context.Background(), app_conf.Project+":"+channel, message).Err()
}
