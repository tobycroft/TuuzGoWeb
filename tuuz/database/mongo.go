package database

import (
	"fmt"
	"log"
	"time"

	"github.com/Unknwon/goconfig"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"main.go/tuuz/Log"
)

var MongoDB *mongo.Client

func init() {
	_ready_mongo()
	_conn_mongo()
}

func NewOrm() *mongo.Database {
	return MongoDB.Database(mongo_dbname)
}

func _ready_mongo() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
		_ready_mongo()
	} else {
		value, err := cfg.GetSection("mongo")
		if err != nil {
			cfg.SetValue("mongo", "need", "false")
			cfg.SetValue("mongo", "retry", "false")
			cfg.SetValue("mongo", "dbauth", "")
			cfg.SetValue("mongo", "dbuser", "")
			cfg.SetValue("mongo", "dbpass", "")
			cfg.SetValue("mongo", "dbhost", "")
			cfg.SetValue("mongo", "dbport", "")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("mongo_ready")
			_ready_mongo()
		}
		mongo_need = value["need"]
		mongo_retry = value["retry"]

		mongo_auth = value["dbauth"]
		mongo_dbname = value["dbname"]
		mongo_dbuser = value["dbuser"]
		mongo_dbpass = value["dbpass"]
		mongo_dbhost = value["dbhost"]
		mongo_dbport = value["dbport"]
	}
}

func _conn_mongo() {
	if mongo_need == "true" {
		var err error
		MongoDB, err = mongo.Connect(Mongoconfig())
		if err != nil {
			if mongo_retry == "true" {
				Log.Dbrr(err, "mongo not connect")
				time.Sleep(1)
				_conn()
			} else {
				log.Panic(err)
			}
		}
	}

}

func Mongoconfig() *options.ClientOptions {
	var conf options.ClientOptions
	conf.ApplyURI("mongodb://" +
		mongo_dbuser + ":" +
		mongo_dbpass +
		"@" + mongo_dbhost + ":" + mongo_dbport + "/?authSource=" + mongo_auth)
	return &conf
}

var mongo_need string
var mongo_retry string
var mongo_auth string
var mongo_dbname string
var mongo_dbuser string
var mongo_dbpass string
var mongo_dbhost string
var mongo_dbport string
