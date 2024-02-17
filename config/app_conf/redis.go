package app_conf

import (
	"time"
)

var Redicon_address = "10.0.0.171"
var Redicon_port = "6379"

const Redicon_proto = "tcp"

var Redicon_username = ""
var Redicon_password = ""

const Redicon_poolsize = 20
const Redicon_MinIdleConn = 10

var Redicon_on = false

const Recicon_panic_on_link_error = true
const Recion_timeout_dial = 5 * time.Second
const Recion_timeout_read = 5 * time.Second
const Recion_timeout_write = 5 * time.Second
const Recion_timeout_pool = 5 * time.Second

var Recion_db = 0
