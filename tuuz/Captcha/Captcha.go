package Captcha

import (
	"github.com/afocus/captcha"
	"github.com/tobycroft/Calc"
	"image"
	"image/color"
	"log"
	"main.go/tuuz/Redis"
	"time"
)

var cap *captcha.Captcha

func Create(num int) (image.Image, string) {
	cap := captcha.New()
	cap.SetFont("comic.ttf")
	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	return cap.Create(num, captcha.NUM)
}

/*
Need Redis Support
*/
func AutoCreate() (image.Image, string) {
	num := Calc.Rand(10000000, 99999999)
	img, err := ManualCreate(4, Calc.Int2String(num))
	if err != nil {
		log.Print(err)
	}
	return img, Calc.Int2String(num)
}

func ManualCreate(lon int, ident string) (image.Image, error) {
	img, str := Create(lon)
	err := Redis.String_set("__captcha__"+Calc.Md5(ident), str, 600*time.Second)
	if err != nil {
		log.Print(err)
	}
	return img, err
}

func AutoVerify(ident string, cap_string string) bool {
	ret, err := Redis.String_get("__captcha__" + Calc.Md5(ident))
	if err != nil {
		return false
	} else {
		str := ret
		//fmt.Println(str)
		if str == cap_string {
			return true
		} else {
			return false
		}
	}
}
