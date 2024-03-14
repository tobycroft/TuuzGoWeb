package BalanceAction

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/tobycroft/gorose-pro"
	"main.go/extend/Balance/BalanceModel"
	"main.go/extend/Balance/BalanceRecordModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
)

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) App_check_balance(uid, student_id, coin_id interface{}) decimal.Decimal {
	if self.Db == nil {
		self.Db = tuuz.Db()
	}
	var balmodel BalanceModel.Interface
	balmodel.Db = self.Db
	userbalance := balmodel.Api_find(uid, student_id, coin_id)
	if len(userbalance) > 0 {
		return Calc.ToDecimal(userbalance["balance"])
	} else {
		balmodel.Api_insert(uid, student_id, coin_id, 0)
		return decimal.Zero
	}
}

func (self *Interface) App_single_balance(uid, student_id, coin_id, Type, order_id interface{}, amount decimal.Decimal, extra, remark1, remark2 interface{}) error {
	if self.Db == nil {
		self.Db = tuuz.Db()
	}
	db := self.Db
	if order_id == nil {
		order_id = Calc.GenerateOrderId()
	}
	db.Begin()
	balance_left := self.App_check_balance(uid, student_id, coin_id)
	if balance_left.Add(amount).LessThan(decimal.Zero) {
		db.Rollback()
		return errors.New("剩余余额不足，不够扣减")
	}
	var balance BalanceModel.Interface
	balance.Db = db
	if !balance.Api_incr(uid, student_id, coin_id, amount) {
		db.Rollback()
		return errors.New("数据库显示余额扣减失败：余额不足")
	}

	//插入变动数据
	var balancerecord BalanceRecordModel.Interface
	balancerecord.Db = db
	last_record := balancerecord.Api_find_last(uid, student_id, coin_id)
	after := decimal.Zero
	before := "0"
	if len(last_record) > 0 {
		after = Calc.Bc_add(last_record["after"], amount)
		before = last_record["after"].(string)
	} else {
		after = amount
	}
	if after.LessThan(decimal.Zero) {
		db.Rollback()
		return errors.New("余额记录不足")
	}
	if !balancerecord.Api_insert(uid, student_id, coin_id, Type, order_id, before, amount, after, extra, remark1, remark2) {
		db.Rollback()
		return errors.New("balance_record添加失败")
	}
	db.Commit()
	return nil
}
