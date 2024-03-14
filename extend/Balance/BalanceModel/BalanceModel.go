package BalanceModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "ps_balance"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_find(uid, student_id, coin_id interface{}) gorose.Data {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.Where("coin_id", coin_id)
	db.LockForUpdate()
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_sum(student_id, coin_id interface{}) interface{} {
	db := tuuz.Db().Table(Table)
	db.Where("student_id", student_id)
	db.Where("coin_id", coin_id)
	ret, err := db.Sum("balance")
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_select(uid, student_id interface{}) []gorose.Data {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.LockForUpdate()
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_count_user_gt(student_ids []interface{}, balance interface{}) int64 {
	db := tuuz.Db().Table(Table)
	db.WhereIn("student_id", student_ids)
	db.Where("balance", "<=", balance)
	db.GroupBy("student_id")
	ret, err := db.Counts()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func (self *Interface) Api_value(uid, student_id, coin_id interface{}) interface{} {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.Where("coin_id", coin_id)
	ret, err := db.Value("balance")
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_insert(uid, student_id, coin_id, balance interface{}) bool {
	db := self.Db.Table(Table)
	data := map[string]interface{}{
		"uid":        uid,
		"student_id": student_id,
		"coin_id":    coin_id,
		"balance":    balance,
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

func (self *Interface) Api_update(uid, student_id, coin_id, balance interface{}) bool {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.Where("coin_id", coin_id)
	db.Where("balance", ">", 0)
	data := map[string]interface{}{
		"balance": balance,
	}
	db.Data(data)
	num, err := db.Update()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return num > 0
	}
}

func (self *Interface) Api_incr(uid, student_id, coin_id, incr_balance interface{}) bool {
	db := self.Db.Table(Table)
	db.Where("uid", uid)
	db.Where("student_id", student_id)
	db.Where("coin_id", coin_id)
	db.Where("balance", ">", 0)
	num, err := db.Increment("balance", incr_balance)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return num > 0
	}
}

func Api_sum_balance_byStudentIds(student_ids []interface{}) interface{} {
	db := tuuz.Db().Table(Table)
	db.WhereIn("student_id", student_ids)
	ret, err := db.Sum("balance")
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_select_byCoinId(coin_id, balance_gt any) []gorose.Data {
	db := tuuz.Db().Table(Table)
	if coin_id != nil {
		db.Where("coin_id", coin_id)
	}
	if balance_gt != nil {
		db.Where("balance", ">=", balance_gt)
	}
	ret, err := db.Get()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
