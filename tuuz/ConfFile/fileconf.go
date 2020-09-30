package ConfFile

import "github.com/Unknwon/goconfig"

func Load(section, key string) string {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		SaveConf(section, key, "")
		return ""
	}
	value, err := cfg.GetValue(section, key)
	if err != nil {
		SaveConf(section, key, "")
		return ""
	} else {
		return value
	}
}

func LoadSec(section string) map[string]string {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	value, err := cfg.GetSection(section)
	if err != nil {
		return nil
	} else {
		return value
	}
}

func SaveConf(section string, key string, value string) bool {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		return false
	}
	cfg.SetValue(section, key, value)
	err = goconfig.SaveConfigFile(cfg, "conf.ini")
	if err != nil {
		return false
	} else {
		return true
	}
}
