package app_conf

import "github.com/Unknwon/goconfig"

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
			cfg.SetValue("app", "project", Project)
			cfg.SetValue("app", "debug", Debug)
			cfg.SetValue("app", "testmode", "false")
			cfg.SetValue("app", "appmode", AppMode)
			cfg.SetValue("app", "websocketkey", WebsocketKey)
			goconfig.SaveConfigFile(cfg, "conf.ini")
			_ready()
		}
		Project = value["project"]
		Debug = value["debug"]

		TestMode = value["testmode"] == "true"
		AppMode = value["appmode"]
		WebsocketKey = value["websocketKey"]
	}
}
