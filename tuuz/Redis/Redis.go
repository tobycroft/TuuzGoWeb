package Redis

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"main.go/config/app_conf"
	"main.go/tuuz/Jsong"
)

func Set(key string, value interface{}, duration int) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()
	str, err := Jsong.Encode(value)
	if err != nil {
		fmt.Println("redis set failed1json:", err)
		return str, err

	}
	status, errs := RRedis.Do("SET", app_conf.Project+":"+key, str, "EX", duration)
	if errs != nil {
		fmt.Println("redis set failed2:", errs)
		return status, errs
	}
	return status, err
}

func Set_permenent(key string, value interface{}) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()
	str, err := Jsong.Encode(value)
	if err != nil {
		fmt.Println("redis set failed1json:", err)
	}
	status, err := RRedis.Do("SET", app_conf.Project+":"+key, str)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	return status, err
}

func Get(key string) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()

	status, err := RRedis.Do("GET", app_conf.Project+":"+key)
	if err != nil {
		//fmt.Println("redis get failed1:", err)
		return nil, err
	}
	status2, err := redigo.String(status, err)
	if err != nil {
		//fmt.Println("redis get failed2:", err)
		return nil, err
	}
	ret, err := Jsong.JToken(status2)
	if err != nil {
		fmt.Println("redis get failed3:", err, status2)
		return nil, err
	}
	return ret, err
}

func Del(key string) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()
	status, err := RRedis.Do("DEL", app_conf.Project+":"+key)
	if err != nil {
		fmt.Println("redis delete fail", err)
	}
	return status, err
}

func Expire(key string, duration float64) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()
	status, err := RRedis.Do("EXPIRE", app_conf.Project+":"+key, duration)
	if err != nil {
		fmt.Println("err while change duration:", err)
	}
	return status, err
}
