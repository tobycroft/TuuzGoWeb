package database

import (
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tobycroft/gorose-pro"
	"log"
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
		return dsn()
	}
	value, err := cfg.GetSection("database")
	if err != nil {
		return dsn()
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

func dsn() string {
	dbname := "1"
	dbuser := "1"
	dbpass := "1"
	dbhost := "10.0.0.170"
	conntype := "tcp"
	dbport := "3306"
	charset := "utf8mb4"
	return dbuser + ":" + dbpass + "@" + conntype + "(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + charset + "&parseTime=true"
}
