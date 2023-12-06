package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func W网页_取域名(Url string) string {
	var 域名 string

	if Url[:8] == "https://" {
		域名 = W文本_取出中间文本(Url, "https://", "/")
	}
	if Url[:7] == "http://" {
		域名 = W文本_取出中间文本(Url, "http://", "/")
	}
	return 域名
}

func W网页_访问_对象(网址 string, 访问方式 int, 提交信息 string, 提交Cookies string, 返回Cookies *string, 附加协议头 string, 返回协议头 *string, 返回状态代码 *int, 禁止重定向 bool, 字节集提交 []byte, 代理地址 string, 超时 int, 代理用户名 string, 代理密码 string, 代理标识 int, 对象继承 interface{}, 是否自动合并更新Cookie bool, 是否补全必要协议头 bool, 是否处理协议头大小写 bool) []byte {
	client := &http.Client{}
	if 超时 != -1 {
		if 超时 < 1 {
			超时 = 15
		}
		client.Timeout = time.Duration(超时) * time.Second
	}

	if 代理地址 != "" {
		proxyURL, _ := url.Parse(代理地址)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	var method string
	switch 访问方式 {
	case 0:
		method = "GET"
	case 1:
		method = "POST"
	case 2:
		method = "HEAD"
	case 3:
		method = "PUT"
	case 4:
		method = "OPTIONS"
	case 5:
		method = "DELETE"
	case 6:
		method = "TRACE"
	case 7:
		method = "CONNECT"
	default:
		method = "GET"
	}

	req, err := http.NewRequest(method, 网址, bytes.NewBuffer(字节集提交))
	if err != nil {
		return []byte{}
	}

	if 禁止重定向 {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	if 附加协议头 != "" {
		协议头列表 := strings.Split(附加协议头, "\n")
		for _, 协议头 := range 协议头列表 {
			if strings.TrimSpace(协议头) != "" {
				头名, 头值 := 内部_协议头取名值(协议头)
				req.Header.Set(头名, 头值)
			}
		}
	}

	if 提交Cookies != "" {
		req.Header.Set("Cookie", 提交Cookies)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	返回状态代码 = &resp.StatusCode
	*返回协议头 = resp.Header.Get("Content-Type")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}

	if 返回Cookies != nil {
		*返回Cookies = resp.Header.Get("Set-Cookie")
	}

	return body
}

func 内部_协议头取名值(协议头 string) (string, string) {
	头名值 := strings.SplitN(协议头, ":", 2)
	if len(头名值) == 2 {
		return 头名值[0], 头名值[1]
	}
	return "", ""
}
func Q取单条Cookie(原Cookies, 单条Cookie名称 string) string {
	Cookies := strings.Split(原Cookies, ";")
	for i := 0; i < len(Cookies); i++ {
		Name := strings.Trim(strings.Split(Cookies[i], "=")[0], " ")
		if strings.ToLower(Name) == strings.ToLower(单条Cookie名称) {
			if strings.Index(strings.Trim(Cookies[i], " "), ":") != -1 {
				return strings.TrimSpace(strings.Split(strings.Trim(Cookies[i], " "), "=")[1])
			}
			return strings.TrimSpace(strings.Split(strings.Trim(Cookies[i], " "), "=")[1])
		}
	}
	return ""
}
func W网页_Cookie合并更新(旧Cookie, 新Cookie string) string {
	旧Cookie = strings.TrimSpace(旧Cookie)
	新Cookie = strings.TrimSpace(新Cookie)

	局_旧Cookie组 := strings.Split(旧Cookie, ";")
	局_新Cookie组 := strings.Split(新Cookie, ";")

	for _, cookie := range 局_旧Cookie组 {
		if !内部_数组成员是否存在1(局_新Cookie组, 内部_Cookie取名(cookie)) {
			局_新Cookie组 = append(局_新Cookie组, cookie)
		}
	}

	旧Cookie = ""
	for _, cookie := range 局_新Cookie组 {
		if 取文本右边(cookie, 8) != "=deleted" {
			旧Cookie += cookie + "; "
		}
	}
	旧Cookie = 取文本左边(旧Cookie, len(旧Cookie)-2)
	旧Cookie = W文本_去重复文本(旧Cookie, "; ")

	return 旧Cookie
}
func 内部_数组成员是否存在1(数组 []string, 要判断值 string) bool {
	for _, 成员 := range 数组 {
		if 内部_Cookie取名(成员) == 要判断值 {
			return true
		}
	}
	return false
}

func 内部_Cookie取值(Cookie string) string {
	位置 := strings.Index(Cookie, "=")
	if 位置 != -1 {
		结果 := 取文本右边(Cookie, len(Cookie)-位置)
		return 结果
	}
	return ""
}

func 内部_Cookie取名(Cookie string) string {
	位置 := strings.Index(Cookie, "=")
	if 位置 != -1 {
		结果 := 取文本左边(Cookie, 位置-1)
		return 删首尾空(结果)
	}
	return 删首尾空(Cookie)
}
func 内部_数组成员是否存在_文本(数组 []string, 要判断值 string) int {
	for i, 成员 := range 数组 {
		if 成员 == 要判断值 {
			return i
		}
	}
	return -1
}

func W网页_处理协议头(原始协议头 string) string {
	数组 := strings.Split(原始协议头, "\n")
	协议头 := ""
	for i := 0; i < len(数组); i++ {
		冒号位置 := strings.Index(数组[i], ":")
		if 冒号位置 == -1 {
			break
		}

		键名 := strings.TrimSpace(数组[i][:冒号位置])
		if strings.Contains(键名, "-") {
			键名拼接 := ""
			键名数组 := strings.Split(键名, "-")
			总数 := len(键名数组)
			for x := 0; x < 总数; x++ {
				if x == 总数-1 {
					键名拼接 += strings.Title(键名数组[x])
				} else {
					键名拼接 += strings.Title(键名数组[x]) + "-"
				}
			}
			键名 = 键名拼接
		} else {
			键名 = strings.Title(键名)
		}

		键值 := strings.TrimSpace(数组[i][冒号位置+1:])
		if len(键值) > 0 && 键值[0] == ' ' {
			键值 = 键值[1:]
		}

		协议头 += 键名 + 键值 + "\n"
	}

	return strings.TrimRight(协议头, "\n")
}

func 网页_访问_对象(网址 string, 访问方式 int, 提交信息 string, 提交Cookies string, 返回Cookies *string, 附加协议头 string, 返回协议头 *string, 返回状态代码 *int, 禁止重定向 bool, 字节集提交 []byte, 代理地址 string, 超时 int, 代理用户名 string, 代理密码 string, 代理标识 int, 对象继承 interface{}, 是否自动合并更新Cookie bool, 是否补全必要协议头 bool, 是否处理协议头大小写 bool) []byte {
	client := &http.Client{}
	if 超时 != -1 {
		if 超时 < 1 {
			超时 = 15
		}
		client.Timeout = time.Duration(超时) * time.Second
	}

	if 代理地址 != "" {
		proxyURL, _ := url.Parse(代理地址)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	var method string
	switch 访问方式 {
	case 0:
		method = "GET"
	case 1:
		method = "POST"
	case 2:
		method = "HEAD"
	case 3:
		method = "PUT"
	case 4:
		method = "OPTIONS"
	case 5:
		method = "DELETE"
	case 6:
		method = "TRACE"
	case 7:
		method = "CONNECT"
	default:
		method = "GET"
	}

	req, err := http.NewRequest(method, 网址, bytes.NewBuffer(字节集提交))
	if err != nil {
		return []byte{}
	}

	if 禁止重定向 {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	if 附加协议头 != "" {
		协议头列表 := strings.Split(附加协议头, "\n")
		for _, 协议头 := range 协议头列表 {
			if strings.TrimSpace(协议头) != "" {
				头名, 头值 := 内部_协议头取名值(协议头)
				req.Header.Set(头名, 头值)
			}
		}
	}

	if 提交Cookies != "" {
		req.Header.Set("Cookie", 提交Cookies)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}
	}
	defer resp.Body.Close()

	返回状态代码 = &resp.StatusCode
	*返回协议头 = resp.Header.Get("Content-Type")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}

	if 返回Cookies != nil {
		*返回Cookies = resp.Header.Get("Set-Cookie")
	}

	return body
}
