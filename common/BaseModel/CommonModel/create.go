package CommonModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

type Interface struct {
	Table string
	Db    gorose.IOrm
}

func (self *Interface) Create(data map[string]interface{}) int64 {
	db := self.Db.Table(self.Table)
	db.Data(data)
	ret, err := db.InsertGetId()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Create(table string, data map[string]interface{}) int64 {
	db := tuuz.Db().Table(table)
	db.Data(data)
	ret, err := db.InsertGetId()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}
