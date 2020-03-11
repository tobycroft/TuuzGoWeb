package Cache

import (
	"time"
)

var CacheMap map[string]interface{}
var CacheList map[float64][]string

func Cache(key string, value interface{}, duration float64) {
	switch config() {
	case 1:
		memory(key, value, duration)
		break

	case 2:
		red(key, value, duration)
		break

	default:
		break
	}

}

func memory(key string, value interface{}, duration float64) bool {
	time := float64(time.Now().Unix())
	CacheList[time] = append(CacheList[time+duration], key)
	CacheMap[key] = value
	return true
}

func red(key interface{}, value interface{}, duration float64) {

}

func Timer() {
	for {
		for key, value := range CacheList {
			if key < float64(time.Now().Unix()) {
				if len(value) > 0 {
					for _, v := range value {
						if CacheMap[v] != nil {
							delete(CacheMap, v)
						}
					}
				}
				delete(CacheList, key)
			}
		}
		time.Sleep(time.Second)
	}
}
