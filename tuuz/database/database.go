package database

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tobycroft/gorose-pro"
	"log"
	"main.go/config/db_conf"
	"main.go/tuuz/Log"
	"time"
)

var Database *gorose.Engin

func init() {
	_ready()
	_conn()
}

func _ready() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
		_ready()
	} else {
		value, err := cfg.GetSection("database")
		if err != nil {
			cfg.SetValue("database", "need", "true")
			cfg.SetValue("database", "retry", "false")
			cfg.SetValue("database", "dbname", "")
			cfg.SetValue("database", "dbuser", "")
			cfg.SetValue("database", "dbpass", "")
			cfg.SetValue("database", "dbhost", "")
			cfg.SetValue("database", "dbport", "")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("database_ready")
			_ready()
		}
		need = value["need"]
		retry = value["retry"]

		dbname = value["dbname"]
		dbuser = value["dbuser"]
		dbpass = value["dbpass"]
		dbhost = value["dbhost"]
		dbport = value["dbport"]
	}
}

func _conn() {
	if need == "true" {
		var err error
		Database, err = gorose.Open(DbConfig())
		if err != nil {
			if retry == "true" {
				Log.Dbrr(err, "database not connect")
				time.Sleep(1)
				_conn()
			} else {
				log.Panic(err)
			}
		}
	}

}

func DbConfig() *gorose.Config {
	var conf gorose.Config
	conf.Driver = "mysql"
	conf.SetMaxIdleConns = 20
	conf.SetMaxOpenConns = 300
	conf.Prefix = ""
	conf.Dsn = dsn_local()
	return &conf
}

func dsn_local() string {
	if need == "true" {
		if dbhost == "" || dbport == "" {
			return db_conf.Dsn()
		}
	}
	conntype := "tcp"
	charset := "utf8mb4"
	return dbuser + ":" + dbpass + "@" + conntype + "(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
}

var need string
var retry string
var dbname string
var dbuser string
var dbpass string
var dbhost string
var dbport string
