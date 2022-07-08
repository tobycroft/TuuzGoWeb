package Redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"main.go/config/app_conf"
)

var goredis = context.Background()

func init() {
	options := redis.Options{
		Addr:         app_conf.Redicon_address + ":" + app_conf.Redicon_port,
		Username:     app_conf.Redicon_username,
		Password:     app_conf.Redicon_password, // no password set
		Network:      app_conf.Redicon_proto,
		DialTimeout:  app_conf.Redicon_timeout,
		PoolTimeout:  30,
		ReadTimeout:  app_conf.Redicon_readtimeout,
		WriteTimeout: app_conf.Redicon_writetimeout,
		DB:           0, // use default DB
		MinIdleConns: app_conf.Redicon_MinIdleConn,
	}
	if app_conf.Redicon_poolsize > 0 {
		options.PoolSize = app_conf.Redicon_poolsize
	}
	//rdb := redis.NewClient(&options)
}

func Conn() {

}
