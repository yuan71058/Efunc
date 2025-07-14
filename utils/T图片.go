package utils

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

func T图片_生成二维码base64(内容 string) string {
	局_二维码base64 := ""
	png, err := qrcode.Encode(内容, qrcode.Medium, 256)
	if err == nil {
		局_二维码base64 = base64.StdEncoding.EncodeToString(png)
	}
	return 局_二维码base64
}
