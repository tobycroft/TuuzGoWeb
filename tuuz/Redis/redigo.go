package Redis

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"log"
	"main.go/config/app_conf"
)

var pool *redigo.Pool

func init() {
	pool_size := app_conf.Redicon_poolsize
	pool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%s", app_conf.Redicon_address, app_conf.Redicon_port),
			redigo.DialUsername(app_conf.Redicon_username),
			redigo.DialPassword(app_conf.Redicon_password),
			redigo.DialConnectTimeout(app_conf.Redicon_timeout),
			redigo.DialClientName(app_conf.Project),
		)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		return c, nil
	}, pool_size)
}

func Conn() redigo.Conn {
	return pool.Get()
}
