package Net

import (
	"net/http"
)

//func CookieUpdater(new_cookie map[string]interface{}, ident string) {
//	user_cookie_map, err := CookieSelector(ident)
//	if err != nil {
//		fmt.Println(err)
//		Log.Err(err)
//		user_cookie_map = new_cookie
//	} else {
//		user_cookie_map = Array.Merge(user_cookie_map, new_cookie)
//	}
//	err = Redis.Hash_add("__cookie__"+ident, user_cookie_map)
//	if err != nil {
//		fmt.Println(err)
//		Log.Err(err)
//		return
//	}
//}

//func CookieSelector(ident string) (map[string]interface{}, error) {
//	user_cookie_map, err := Redis.Hash_map_get("__cookie__" + ident)
//	if err != nil {
//		return make(map[string]interface{}), err
//	}
//	arr := make(map[string]interface{})
//	//fmt.Println(user_cookie_map)
//	for s, s2 := range user_cookie_map {
//		arr[s] = s2
//	}
//	return arr, err
//}

func (self Curl) cookieHandler(resp_headers []*http.Cookie) map[string]interface{} {
	cookie_arr := make(map[string]interface{})
	for _, resp_header := range resp_headers {
		cookie_arr[resp_header.Name] = resp_header.Value
	}
	return cookie_arr
}
