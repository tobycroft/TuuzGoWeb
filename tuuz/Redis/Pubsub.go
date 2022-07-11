package Redis

import "fmt"

func Suscribe(channel string) {

	pubsub := goredis.Subscribe(goredis_ctx, channel)

	_, err := pubsub.Receive(goredis_ctx)
	if err != nil {
		return
	}

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}
