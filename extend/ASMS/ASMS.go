package ASMS

import (
	"errors"
	"github.com/Unknwon/goconfig"
	"github.com/tobycroft/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
	"time"
)

/*
 *
 * 请将本文件与文件夹放入TuuzGo的extend文件夹下
 *
 * 上面的token为API系统的标识符，用于区分项目
 *

 *
 */

const url = "http://asms.tuuz.cc:10081"

func Sms_send(phone any, quhao, text any) error {
	conf, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		return err
	}
	value, err := conf.GetSection("asms")
	if err != nil {
		return err
	}
	name := value["name"]
	token := value["token"]
	ts := time.Now().Unix()
	param := map[string]any{
		"phone": phone,
		"quhao": quhao,
		"text":  text,
		"ts":    ts,
		"name":  name,
		"sign":  Calc.Md5(token + Calc.Any2String(ts)),
	}
	ret, err := Net.Post(url+"/asms/send", nil, param, nil, nil)
	//fmt.Println(ret, err)
	if err != nil {
		Log.Crrs(err, ret)
		return err
	} else {
		rtt, errs := Jsong.JObject[string, any](ret)
		if errs != nil {
			return errors.New(ret)
		} else {
			if Calc.Any2String(rtt["code"]) == "0" {
				return nil
			} else {
				Log.Crrs(err, ret)
				return errors.New(Calc.Any2String(rtt["echo"]))
			}
		}
	}
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
