package database

import (
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tobycroft/gorose-pro"
	"log"
	"main.go/config/db_conf"
)

var Database *gorose.Engin

func init() {
	var err error
	Database, err = gorose.Open(DbConfig())
	if err != nil {
		log.Panic(err)
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
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		return db_conf.Dsn()
	}
	value, err := cfg.GetSection("database")
	if err != nil {
		return db_conf.Dsn()
	} else {
		dbname := value["dbname"]
		dbuser := value["dbuser"]
		dbpass := value["dbpass"]
		dbhost := value["dbhost"]
		dbport := value["dbport"]
		conntype := "tcp"
		charset := "utf8mb4"
		return dbuser + ":" + dbpass + "@" + conntype + "(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
	}
}
