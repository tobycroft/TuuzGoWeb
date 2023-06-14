package RSA

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"strings"
)

const (
	PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	PEM_END   = "\n-----END RSA PRIVATE KEY-----"
)

// RsaSign sign = RsaSign(signContent, privateKey, crypto.SHA256)
func RsaSign(signContent string, key string, hash crypto.Hash) string {
	// Convert private key from PEM format to RSA private key
	privateKeyBytes := []byte(FormatPrivateKey(key)) // Replace with your actual private key
	block, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return ""
	}

	hashes := sha256.Sum256([]byte(signContent))
	//有些公司这里使用二重签，这里记得先hash
	hashHex := hex.EncodeToString(hashes[:])
	hashes = sha256.Sum256([]byte(hashHex))

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashes[:])
	if err != nil {
		panic(err)
	}

	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	return signatureBase64
}

func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateKey = FormatPrivateKey(privateKey)
	// 2、解码私钥字节，生成加密对象
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}

func FormatPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
		privateKey = PEM_BEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEM_END) {
		privateKey = privateKey + PEM_END
	}
	return privateKey
}
