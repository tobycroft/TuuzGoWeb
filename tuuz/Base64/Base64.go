package Base64

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Decode(data string) ([]byte, error) {
	ret, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return ret, err
}

func EncodePng(imgStream image.Image) string {
	bt := new(bytes.Buffer)
	png.Encode(bt, imgStream)
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(bt.Bytes())
}
