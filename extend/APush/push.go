package APush

import (
	"errors"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Net"
)

/*
 *
 * 请将本文件与文件夹放入TuuzGo的extend文件夹下
 *
 * 上面的token为API系统的标识符，用于区分项目
 *
 *
 * user_sync是用户同步模块，应该在用户每次打开APP的时候调用本功能接口uid_or_username的意义为如果你的系统使用的用户唯一的标识符为user表自增id，那么uid_or_username就填写你系统中该用户的user_id即可，rid为设备识别码id由前端提供，详情可以参考极光registration_id说明，必须要先使用本接口同步用户信息才可推送，你方系统中可以不需要对rid进行保留处理，下方接口会说明
 *
 *
 *
 * 可是使用single方法做单人推送，使用push_more方法做多人推送，使用push_all方法做全员推送
 *
 * push_single/push_more方法中，uid_or_username为你在sync接口中传入的用户标识数据（例如uid或者username），content为内容主体，苹果是没有title的，所以在苹果上content就是title，安卓上有title（部分系统没有），所以content意义大于title，请不要设定过长content避免推送后因为用户手机顶栏长度有限而无法完全显示你希望推送消息的全文
 * 在more方法中，为uids_or_users，这里因该传入array信息例如[20,21,56,43]这样的uid数据即可同时推送给这部分人
 * 如果需要特殊推送，可以加入extra消息，注意extra一定是一个object类型的数据，不能是array否则推送将会直x f接报错无法执行，另外extra是额外消息，本消息只能由apicloud通过eventlisterner
接收，随推顶栏推送接收，一般用于跳转落地页，有需要请自行设定
 *
 * push_all方法为全员推送，无定向性，所有人都会收到，应该避免线上测试误操作调用本接口
 *
 * 如果有错将会返回err内容，如果没错err返回就是nil直接判断err如果有就显示（or not）
 *
*/

const token = "gochat"

const url = "http://push.tuuz.cc:10080"

func Push_single(uid interface{}, content, title, extra interface{}) error {
	users, err := Jsong.Encode([]interface{}{uid})
	if err != nil {
		return err
	}
	param := map[string]interface{}{
		"users":   users,
		"content": content,
		"title":   title,
		"extra":   extra,
		"token":   token,
	}
	ret, err := Net.Post(url+"/push", nil, param, nil, nil)
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
				return errors.New(Calc.Any2String(rtt["data"]))
			}
		}
	}
}

func Push_more(uid []interface{}, content, title, extra interface{}) error {
	users, err := Jsong.Encode(uid)
	if err != nil {
		Log.Errs(err, tuuz.FUNCTION_ALL())
		return err
	}
	param := map[string]interface{}{
		"users":   users,
		"content": content,
		"title":   title,
		"extra":   extra,
		"token":   token,
	}
	ret, err := Net.Post(url+"/push", nil, param, nil, nil)
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
				return errors.New(Calc.Any2String(rtt["data"]))
			}
		}
	}
}

func Push_message(uid interface{}, content, extra interface{}) error {
	users, err := Jsong.Encode([]interface{}{uid})
	if err != nil {
		return err
	}
	param := map[string]interface{}{
		"users":   users,
		"content": content,
		"extra":   extra,
		"token":   token,
	}
	ret, err := Net.Post(url+"/message", nil, param, nil, nil)
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
				return errors.New(Calc.Any2String(rtt["data"]))
			}
		}
	}
}

func Push_message_more(uids []interface{}, content, extra interface{}) error {
	users, err := Jsong.Encode(uids)
	if err != nil {
		return err
	}
	param := map[string]interface{}{
		"users":   users,
		"content": content,
		"extra":   extra,
		"token":   token,
	}
	ret, err := Net.Post(url+"/message", nil, param, nil, nil)
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
				return errors.New(Calc.Any2String(rtt["data"]))
			}
		}
	}
}

func Push_all(content, title, extra interface{}) error {
	param := map[string]interface{}{
		"content": content,
		"title":   title,
		"extra":   extra,
		"token":   token,
	}
	ret, err := Net.Post(url+"/push_all", nil, param, nil, nil)
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
				return errors.New(Calc.Any2String(rtt["data"]))
			}
		}
	}
}

func User_sync(user, rid interface{}) error {
	param := map[string]interface{}{
		"user":  user,
		"rid":   rid,
		"token": token,
	}
	ret, err := Net.Post(url+"/sync", nil, param, nil, nil)
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
				return errors.New(Calc.Any2String(rtt["data"]))
			}
		}
	}
}
