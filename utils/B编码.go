package utils

import (
	"encoding/base64"
	"net/url"
	"strconv"
)

// B编码_URL编码 对文本进行 URL 编码（百分号编码）。
// 将特殊字符转换为 %XX 格式，中文等非 ASCII 字符也会被编码。
//
// 参数:
//   - 欲编码的文本: 待编码的文本
//
// 返回:
//   - string: URL 编码后的文本
//
// 示例:
//
//	B编码_URL编码("go语言")  // "go%E8%AF%AD%E8%A8%80"
func B编码_URL编码(欲编码的文本 string) string {
	return url.QueryEscape(欲编码的文本)
}

// B编码_URL解码 对 URL 编码的文本进行解码。
// 将 %XX 格式的编码还原为原始字符。
//
// 参数:
//   - URL: URL 编码的文本
//
// 返回:
//   - string: 解码后的文本；解码失败返回空串
func B编码_URL解码(URL string) string {
	decodedURL, err := url.QueryUnescape(URL)
	if err != nil {
		return ""
	}
	return decodedURL
}

// B编码_usc2到文本 将 USC2/Unicode 转义序列转换为中文文本。
// 例如将 \u4e2d\u6587 转换为 "中文"。
//
// 参数:
//   - 字符串: 包含 USC2 转义序列的字符串（如 "\\u4e2d\\u6587"）
//
// 返回:
//   - string: 转换后的中文文本；失败返回空串
func B编码_usc2到文本(字符串 string) string {
	解码文本, err := strconv.Unquote(`"` + 字符串 + `"`)
	if err != nil {
		return ""
	}
	return 解码文本
}

// B编码_BASE64编码 将字节集进行 Base64 编码。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: Base64 编码后的字符串
//
// 示例:
//
//	B编码_BASE64编码([]byte("hello"))  // "aGVsbG8="
func B编码_BASE64编码(字节集 []byte) string {
	return base64.StdEncoding.EncodeToString(字节集)
}

// B编码_BASE64解码 将 Base64 编码的文本解码为字节集。
//
// 参数:
//   - 文本: Base64 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_BASE64解码(文本 string) []byte {
	解码字节集, err := base64.StdEncoding.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}
