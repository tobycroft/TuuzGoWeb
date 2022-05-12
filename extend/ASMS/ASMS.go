package ASMS

import (
	"errors"
	"github.com/Unknwon/goconfig"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
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

func Sms_single(phone interface{}, quhao, code interface{}) error {
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
	param := map[string]interface{}{
		"phone": phone,
		"quhao": quhao,
		"text":  code,
		"ts":    ts,
		"name":  name,
		"sign":  Calc.Md5(token + Calc.Any2String(ts)),
	}
	ret, err := Net.Post(url+"/asms/send", nil, param, nil, nil)
	//fmt.Println(ret, err)
	if err != nil {
		return err
	} else {
		rtt, errs := Jsong.JObject(ret)
		if errs != nil {
			return errs
		} else {
			if rtt["code"].(float64) == 0 {
				return nil
			} else {
				return errors.New(Calc.Any2String(rtt["echo"]))
			}
		}
	}
}
