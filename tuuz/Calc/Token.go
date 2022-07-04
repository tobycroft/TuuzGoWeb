package Calc

import (
	"strconv"
	"sync/atomic"
	"time"
)

func GenerateToken() string {
	unix := time.Now().UnixNano()
	rand := Rand(0, 99999999)
	str := strconv.FormatInt(unix, 10) + strconv.FormatInt(int64(rand), 10)
	return Md5(str)
}

var BaseNum = int64(0)

func GenerateOrderId() string {
	new_num := atomic.AddInt64(&BaseNum, 1)
	str := Int642String(time.Now().Unix())
	return time.Now().Format("D20060102T150405U") + str + "R" + Int642String(new_num)
}

func RefreshBaseNum() {
	ticker := time.NewTicker(1 * time.Second)
	//done := make(chan bool)
	go func() {
		for {
			select {
			//case <-done:
			//	return
			case <-ticker.C:
				atomic.StoreInt64(&BaseNum, 0)
			}
		}
	}()
}
