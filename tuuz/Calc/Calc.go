package Calc

import (
	"crypto/rand"
	"math"
	"math/big"
	rand2 "math/rand"
	"strconv"
	"strings"
	"time"
)

func Mt_rand(min, max int64) int64 {
	rand2.Seed(Seed())
	if min == max {
		return min
	} else {
		r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
		return r.Int63n(max-min+1) + min
	}
}

func Seed() int64 {
	num, _ := rand.Int(rand.Reader, big.NewInt(999999999))
	return num.Int64() + time.Now().UnixNano()
}

func Rand(min, max int) int {
	rand2.Seed(Seed())
	if min == max {
		return min
	} else {
		var randNum int
		if max-min < 0 {
			randNum = rand2.Intn(min-max) + min
		} else {
			randNum = rand2.Intn(max-min) + min
		}
		return randNum
	}
}

func Any2Int64(any interface{}) int64 {
	ret, err := Any2Int64_2(any)
	if err != nil {
		return -99999998
	}
	return ret
}

func Any2Int64_2(any interface{}) (int64, error) {
	return String2Int64(Any2String(any))
}

func Any2Float64(any interface{}) float64 {
	ret, err := String2Float64(Any2String(any))
	if err != nil {
		return 0
	}
	return ret
}
func Any2Float64_2(any interface{}) (float64, error) {
	return String2Float64(Any2String(any))
}

func Any2Int(any interface{}) int {
	ret, err := Any2Int_2(any)
	if err != nil {
		return -99999998
	}
	return ret
}

func Any2Int_2(any interface{}) (int, error) {
	return String2Int(Any2String(any))
}

func Hex2Dec(val string) int64 {
	val = strings.TrimLeft(val, "0x")
	if val == "" {
		return 0
	}
	n := new(big.Int)
	n, _ = n.SetString(val, 16)

	return n.Int64()
}

func Dec2Hex(val int64) string {
	n := strconv.FormatInt(val, 16)
	return n
}

func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

func Transfer2Eth(value float64, decimal int) float64 {
	return value / math.Pow10(Any2Int(decimal))
}

func Round(x float64, decimal int) float64 {
	n := math.Pow10(decimal)
	return math.Trunc((x+0.5/n)*n) / n
}

func Decimal(x float64, decimal int) string {
	value := strconv.FormatFloat(x, 'f', decimal, 64)
	return value
}
