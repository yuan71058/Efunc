package utils

import (
	"encoding/base64"
	"net/url"
	"strconv"
)

func B编码_URL编码(欲编码的文本 string) string {
	return url.QueryEscape(欲编码的文本)
}

func B编码_URL解码(URL string) string {

	decodedURL, err := url.QueryUnescape(URL)
	if err != nil {
		return ""
	}
	return decodedURL
}
func B编码_usc2到文本(字符串 string) string {

	解码文本, err := strconv.Unquote(`"` + 字符串 + `"`)
	if err != nil {
		return ""
	}
	return 解码文本
}
func B编码_BASE64编码(字节集 []byte) string {
	return base64.StdEncoding.EncodeToString(字节集)
}

func B编码_BASE64解码(文本 string) []byte {
	解码字节集, err := base64.StdEncoding.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}
