package BalanceRecordModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

const Table = "ps_balance_record"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, student_id, coin_id, Type, order_id, before, amount, after, extra, remark1, remark2 interface{}) bool {
	db := self.Db.Table(Table)
	data := map[string]interface{}{
		"uid":        uid,
		"student_id": student_id,
		"coin_id":    coin_id,
		"type":       Type,
		"order_id":   order_id,
		"before":     before,
		"amount":     amount,
		"after":      after,
		"extra":      extra,
		"remark1":    remark1,
		"remark2":    remark2,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_insert_special_date(uid, student_id, coin_id, Type, order_id, before, amount, after, extra, remark1, remark2, date interface{}) bool {
	db := self.Db.Table(Table)
	data := map[string]interface{}{
		"uid":        uid,
		"student_id": student_id,
		"coin_id":    coin_id,
		"type":       Type,
		"order_id":   order_id,
		"before":     before,
		"amount":     amount,
		"after":      after,
		"extra":      extra,
		"remark1":    remark1,
		"remark2":    remark2,
		"date":       date,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_find(uid, student_id, id interface{}) gorose.Data {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.Where("id", id)
	db.OrderBy("id desc")
	db.LockForUpdate()
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_find_last(uid, student_id, coin_id interface{}) gorose.Data {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.Where("coin_id", coin_id)
	db.OrderBy("id desc")
	db.LockForUpdate()
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_like(uid, student_id any, extra string) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.Where("extra", "like", extra)
	db.OrderBy("id desc")
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select(uid interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(Table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_paginator(uid, student_id any, limit, page int) gorose.Paginate {
	db := tuuz.Db().Table(Table)
	if uid != nil {
		db.Where("uid", uid)
	}
	if student_id != nil {
		db.Where("student_id", student_id)
	}
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Paginator()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return gorose.Paginate{}
	} else {
		return ret
	}
}

func (self *Interface) Api_select_byCoinId(uid, coin_id interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(Table)
	where := map[string]interface{}{
		"uid":     uid,
		"coin_id": coin_id,
	}
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select_byType(uid, Type interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(Table)
	where := map[string]interface{}{
		"uid":  uid,
		"type": Type,
	}
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select_inType(uid interface{}, Type []interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(Table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	db.WhereIn("type", Type)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select_inTypeAndCoinId(uid interface{}, coin_id interface{}, Type []interface{}, limit, page int) []gorose.Data {
	db := self.Db.Table(Table)
	where := map[string]interface{}{
		"uid":     uid,
		"coin_id": coin_id,
	}
	db.Where(where)
	db.WhereIn("type", Type)
	db.Limit(limit)
	db.Page(page)
	db.OrderBy("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_sum_balance_likeExtra(student_ids []interface{}, coin_id any, extra string, start_date, end_date time.Time) interface{} {
	db := tuuz.Db().Table(Table)
	db.WhereIn("student_id", student_ids)
	db.Where("date", ">=", start_date)
	if coin_id != nil {
		db.Where("coin_id", coin_id)
	}
	db.Where("date", "<", end_date)
	if extra != "" {
		db.Where("extra", "like", extra+"%")
	}
	ret, err := db.Sum("amount")
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_sum_balance_likeExtra_byUid(uid []any, student_ids []interface{}, coin_id any, extra string, start_date, end_date time.Time) interface{} {
	db := tuuz.Db().Table(Table)
	if uid != nil {
		db.WhereIn("uid", uid)
	}
	db.WhereIn("student_id", student_ids)
	db.Where("date", ">=", start_date)
	if coin_id != nil {
		db.Where("coin_id", coin_id)
	}
	db.Where("date", "<", end_date)
	if extra != "" {
		db.Where("extra", "like", extra+"%")
	}
	ret, err := db.Sum("amount")
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_sum_balance_likeExtra(student_ids []interface{}, coin_id any, extra string, start_date, end_date time.Time) []gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Fields("*,date(date) as date_day")
	db.WhereIn("student_id", student_ids)
	db.Where("date", ">=", start_date)
	if coin_id != nil {
		db.Where("coin_id", coin_id)
	}
	db.Where("date", "<", end_date)
	if extra != "" {
		db.Where("extra", "like", extra+"%")
	}
	db.Group("date(date)")
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_update_date_byId(id, date any) bool {
	db := tuuz.Db().Table(Table)
	db.Where("id", id)
	db.Data(map[string]any{
		"date": date,
	})
	_, err := db.Update()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
