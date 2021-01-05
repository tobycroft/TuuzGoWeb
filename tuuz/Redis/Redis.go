package Redis

import (
	"errors"
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
	"main.go/config/app_conf"
	"main.go/tuuz/Jsong"
)

func Add(key string, value string, duration int) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()
	str, err := Jsong.Encode(value)
	if err != nil {
		fmt.Println("redis set failed1json:", err)
		return str, err

	}
	status, errs := RRedis.Do("SADD", app_conf.Project+":"+key, str, "EX", duration)
	if errs != nil {
		fmt.Println("redis set failed2:", errs)
		return status, errs
	}
	return status, err
}

func IsMember(key, value string) bool {
	RRedis := Conn()
	defer RRedis.Close()
	ismember, err := redigo.Bool(RRedis.Do("sIsMember", key, value))
	if err != nil {
		return false
	} else {
		return ismember
	}
}

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

func SetRaw(key string, value interface{}, duration int) error {
	redis := Conn()
	defer redis.Close()
	_, err := redis.Do("SET", app_conf.Project+":"+key, value, "EX", duration)
	if err != nil {
		return err
	} else {
		return nil
	}
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

func CheckExists(key string) bool {
	redis := Conn()
	defer redis.Close()
	is_key_exit, err := redigo.Bool(redis.Do("EXISTS", app_conf.Project+":"+key))
	if err != nil {
		return false
	}
	if is_key_exit {
		return true
	} else {
		return false
	}
}

func GetString(key string) (string, error) {
	redis := Conn()
	defer redis.Close()
	is_key_exit, err := redigo.Bool(redis.Do("EXISTS", app_conf.Project+":"+key))
	if !is_key_exit {
		return "", errors.New("not exists")
	}
	if err != nil {
		return "", err
	}
	n, err := redis.Do("GET", app_conf.Project+":"+key)
	return redigo.String(n, err)
}

func GetBool(key string) (bool, error) {
	redis := Conn()
	defer redis.Close()
	is_key_exit, err := redigo.Bool(redis.Do("EXISTS", app_conf.Project+":"+key))
	if !is_key_exit {
		return false, errors.New("not exists")
	}
	if err != nil {
		return false, err
	}
	n, err := redis.Do("GET", app_conf.Project+":"+key)
	return redigo.Bool(n, err)
}

func GetInt(key string) (int, error) {
	redis := Conn()
	defer redis.Close()
	is_key_exit, err := redigo.Bool(redis.Do("EXISTS", app_conf.Project+":"+key))
	if !is_key_exit {
		return 0, errors.New("not exists")
	}
	if err != nil {
		return 0, err
	}
	n, err := redis.Do("GET", app_conf.Project+":"+key)
	return redigo.Int(n, err)
}

func GetFloat64(key string) (float64, error) {
	redis := Conn()
	defer redis.Close()
	is_key_exit, err := redigo.Bool(redis.Do("EXISTS", app_conf.Project+":"+key))
	if !is_key_exit {
		return 0, errors.New("not exists")
	}
	if err != nil {
		return 0, err
	}
	n, err := redis.Do("GET", app_conf.Project+":"+key)
	return redigo.Float64(n, err)
}

func GetAny(key string, value *interface{}) error {
	redis := Conn()
	defer redis.Close()
	is_key_exit, err := redigo.Bool(redis.Do("EXISTS", app_conf.Project+":"+key))
	if !is_key_exit {
		return errors.New("not exists")
	}
	if err != nil {
		return err
	}
	*value, err = redis.Do("GET", app_conf.Project+":"+key)
	if err != nil {
		return err
	}
	return nil
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

func Rpush(key string, value interface{}, duration interface{}) error {
	RRedis := Conn()
	defer RRedis.Close()
	_, err := RRedis.Do("rPush", app_conf.Project+":"+key, value)
	if err != nil {
		fmt.Println("redis set fail:", err)
	}
	return err
}

func Lrange(key string) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()
	data, err := RRedis.Do("LRANGE", app_conf.Project+":"+key, "0", "-1")
	if err != nil {
		fmt.Println("redis set fail:", err)
	}
	return data, err
}

func Lpop(key string) (interface{}, error) {
	RRedis := Conn()
	defer RRedis.Close()
	data, err := RRedis.Do("LPOP", app_conf.Project+":"+key)
	if err != nil {
		fmt.Println("redis set fail:", err)
	}
	return data, err
}

func Lpush(key string, value interface{}) error {
	RRedis := Conn()
	defer RRedis.Close()
	_, err := RRedis.Do("lPush", app_conf.Project+":"+key, value)
	if err != nil {
		fmt.Println("redis set fail:", err)
	}
	return err
}
