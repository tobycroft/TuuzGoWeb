package SystemParamModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "system_param"

func Api_find_val(key interface{}) interface{} {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"key": key,
	}
	db.Where(where)
	ret, err := db.Value("val")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_KV() map[string]string {
	db := tuuz.Db().Table(table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		arr := map[string]string{}
		for _, data := range ret {
			arr[data["key"].(string)] = data["val"].(string)
		}
		return arr
	}
}
