package Redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"main.go/config/app_conf"
	"time"
)

var goredis_ctx = context.Background()

var goredis *redis.Client

func init() {
	if !app_conf.Recicon_on {
		return
	}
	options := redis.Options{
		Addr:         app_conf.Redicon_address + ":" + app_conf.Redicon_port,
		Username:     app_conf.Redicon_username,
		Password:     app_conf.Redicon_password, // no password set
		Network:      app_conf.Redicon_proto,
		DB:           0, // use default DB
		MinIdleConns: app_conf.Redicon_MinIdleConn,
	}
	if app_conf.Redicon_poolsize > 0 {
		options.PoolSize = app_conf.Redicon_poolsize
	}
	goredis = redis.NewClient(&options)
	go func() {
		for {
			time.Sleep(3 * time.Second)
			ret, err := goredis.Ping(goredis_ctx).Result()
			if err != nil {
				log.Println("redis_out", ret, err)
				if app_conf.Recicon_panic_on_link_error {
					panic("redis_out")
				}
			}
		}
	}()
}
