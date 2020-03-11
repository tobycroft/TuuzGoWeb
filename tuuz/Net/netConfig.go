package Net

import "strings"

const cookie_tag = "sid,JSESSIONID,DedeUserID,DedeUserID__ckMd5,SESSDATA,bili_jct,_dfcaptcha,rpdid,INTVER,_uuid,buvid3,buvid2,LIVE_BUVID,rpdid"

func CookieTagChecker(cookie_key string) bool {
	if cookie_tag == "" {
		return true
	} else {
		arr := strings.Split(cookie_tag, ",")
		for _, v := range arr {
			if v == cookie_key {
				return true
			}
		}
		return false
	}
}
