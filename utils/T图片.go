package utils

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

// T图片_生成二维码base64 生成指定内容的二维码图片，并返回 Base64 编码字符串。
// 二维码大小为 256x256 像素，容错等级为 Medium。
// 返回的 Base64 字符串可直接用于 HTML img 标签的 src 属性。
//
// 参数:
//   - 内容: 二维码中编码的内容文本
//
// 返回:
//   - string: Base64 编码的 PNG 图片字符串；生成失败返回空串
//
// 示例:
//
//	base64 := T图片_生成二维码base64("https://example.com")
//	// 可用于: <img src="data:image/png;base64,xxx">
func T图片_生成二维码base64(内容 string) string {
	局_二维码base64 := ""
	png, err := qrcode.Encode(内容, qrcode.Medium, 256)
	if err == nil {
		局_二维码base64 = base64.StdEncoding.EncodeToString(png)
	}
	return 局_二维码base64
}
