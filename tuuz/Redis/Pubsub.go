package Redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"main.go/config/app_conf"
)

type PubSub struct {
}

func (self PubSub) Subscribe(channel string) <-chan *redis.Message {
	return goredis.Subscribe(context.Background(), app_conf.Project+":"+channel).Channel()
}

func (self PubSub) Publish(channel string, message any) error {
	return goredis.Publish(context.Background(), app_conf.Project+":"+channel, message).Err()
}

func (self PubSub) Publish_struct(channel string, message_struct any) error {
	message, err := sonic.MarshalString(message_struct)
	if err != nil {
		return err
	}
	return goredis.Publish(context.Background(), app_conf.Project+":"+channel, message).Err()
}
