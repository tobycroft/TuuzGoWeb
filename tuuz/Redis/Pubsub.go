package Redis

import (
	"fmt"
	"main.go/config/app_conf"
)

func Suscribe(channel string) {

	pubsub := goredis.Subscribe(goredis_ctx, app_conf.Project+":"+channel)

	_, err := pubsub.Receive(goredis_ctx)
	if err != nil {
		return
	}

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}
