package app_conf

import (
	"github.com/Unknwon/goconfig"
	"time"
)

var Redicon_address = "10.0.0.171"
var Redicon_port = "6379"

const Redicon_proto = "tcp"
const Redicon_username = ""
const Redicon_password = ""
const Redicon_poolsize = 20
const Redicon_MinIdleConn = 10

var Recicon_on = false

const Recicon_panic_on_link_error = true
const Recion_timeout_dial = 5 * time.Second
const Recion_timeout_read = 5 * time.Second
const Recion_timeout_write = 5 * time.Second
const Recion_timeout_pool = 5 * time.Second
const Recion_db = 0

func init() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
	} else {
		value, err := cfg.GetSection("redis")
		if err != nil {
		} else {
			if value["address"] != "" && value["address"] != "" && value["database"] != "" {
				Redicon_address = value["address"]
				Redicon_port = value["port"]
				Recicon_on = true
			}
		}
	}
}
