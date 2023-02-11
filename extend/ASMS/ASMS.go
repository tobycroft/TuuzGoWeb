package ASMS

import (
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/tobycroft/AossGoSdk"
)

/*
 *
 * 请将本文件与文件夹放入TuuzGo的extend文件夹下
 *
 * 上面的token为API系统的标识符，用于区分项目
 *

 *
 */

func init() {
	_ready()
}

var name string
var token string

func _ready() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		goconfig.SaveConfigFile(&goconfig.ConfigFile{}, "conf.ini")
		_ready()
	} else {
		value, err := cfg.GetSection("asms")
		if err != nil {
			cfg.SetValue("asms", "name", "")
			cfg.SetValue("asms", "token", "")
			goconfig.SaveConfigFile(cfg, "conf.ini")
			fmt.Println("asms_ready")
			_ready()
		} else {
			if value["name"] != "" && value["token"] != "" {
				name = value["name"]
				token = value["token"]
			}
		}
	}
}

func Sms_send(phone, quhao, text any) error {
	var sms AossGoSdk.ASMS
	sms.Name = name
	sms.Token = token
	return sms.Sms_send(phone, quhao, text)
}

func Sms_single(phone any, quhao, text any, code int64) error {
	if len(Api_find_in1(phone)) > 0 {
		return errors.New("你已经发送验证码，请稍后再次发送")
	}
	err := Sms_send(phone, quhao, text)
	if err != nil {
		return err
	} else {
		if !Api_insert(phone, code) {
			return errors.New("验证码数据库插入失败")
		}
		return nil
	}
}

func Sms_verify_in10(phone, code any) error {
	data := Api_find_in10(phone, code)
	if len(data) > 0 {
		return nil
	} else {
		return errors.New("验证码出错")
	}
}

func Sms_verify_in5(phone, code any) error {
	data := Api_find_in10(phone, code)
	if len(data) > 0 {
		return nil
	} else {
		return errors.New("验证码出错")
	}
}
