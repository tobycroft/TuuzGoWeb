package app_conf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var Project = "tuuzgoweb"
var Debug = "tgw"
var TestMode = true
var AppMode = "debug"
var WebsocketKey = ""

func init() {
	_ready()
}

func _ready() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
		_ready()
	} else {
		value, err := cfg.GetSection("app")
		if err != nil {
			cfg.SetValue("app", "Project", Project)
			cfg.SetValue("app", "Debug", Debug)
			cfg.SetValue("app", "TestMode", "false")
			cfg.SetValue("app", "AppMode", AppMode)
			cfg.SetValue("app", "WebsocketKey", WebsocketKey)
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("app_ready")
			_ready()
		}
		Project = value["Project"]
		Debug = value["Debug"]

		TestMode = value["TestMode"] == "true"
		AppMode = value["AppMode"]
		WebsocketKey = value["WebsocketKey"]
	}
}
