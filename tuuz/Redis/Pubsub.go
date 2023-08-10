package Redis

import (
	"context"
	"fmt"
	"main.go/config/app_conf"
)

func Suscribe(channel string) {

	pubsub := goredis.Subscribe(context.Background(), app_conf.Project+":"+channel)

	_, err := pubsub.Receive(context.Background())
	if err != nil {
		return
	}

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}
