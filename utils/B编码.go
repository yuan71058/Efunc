package utils

import (
	"net/url"
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
