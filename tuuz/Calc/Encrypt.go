package Calc

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Md5Up(str string) string {
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(str))))
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	//	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}
