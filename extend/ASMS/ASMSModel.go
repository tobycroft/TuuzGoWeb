package ASMS

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "tuuz_asms"

func Api_insert(phone, code interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"phone": phone,
		"code":  code,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(phone, code interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"phone": phone,
		"code":  code,
	}
	db.Where(where)
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_in10(phone, code interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"phone": phone,
		"code":  code,
	}
	db.Where(where)
	db.Where("date>Date_SUB(NOW(),INTERVAL 10 MINUTE)")
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_in5(phone, code interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"phone": phone,
		"code":  code,
	}
	db.Where(where)
	db.Where("date>Date_SUB(NOW(),INTERVAL 5 MINUTE)")
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_in1(phone interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"phone": phone,
	}
	db.Where(where)
	db.Where("date>Date_SUB(NOW(),INTERVAL 1 MINUTE)")
	ret, err := db.Find()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_delete(phone interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"phone": phone,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
