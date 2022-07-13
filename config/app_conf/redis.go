package app_conf

import "time"

const Redicon_address = "10.0.0.171"
const Redicon_port = "6379"
const Redicon_proto = "tcp"
const Redicon_username = ""
const Redicon_password = ""
const Redicon_poolsize = 20
const Redicon_MinIdleConn = 10
const Recicon_on = false
const Recicon_panic_on_link_error = true
const Recion_timeout_dial = 5 * time.Second
const Recion_timeout_read = 5 * time.Second
const Recion_timeout_write = 5 * time.Second
const Recion_timeout_pool = 5 * time.Second

const Recion_db = 0
const Recion_idle_check_feq = 60 * time.Second
const Recion_idle_check_feq = 60 * time.Second
