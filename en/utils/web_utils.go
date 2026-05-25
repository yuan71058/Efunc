// Web 网页工具模块
// 提供 HTTP 请求发送、URL 解析、Cookie 管理、协议头处理等功能。
// 基于 Go 标准库 net/http 实现，支持代理、超时、重定向控制等。
package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// WebUtils_GetDomain 从 URL 中提取域名部分。
// 支持 http:// 和 https:// 两种协议前缀的 URL。
//
// 参数:
//   - rawUrl: 完整的 URL 地址，如 "https://www.example.com/path"
//
// 返回:
//   - string: 域名部分，如 "www.example.com"；无法识别时返回空串
func WebUtils_GetDomain(rawUrl string) string {
	var domain string

	if rawUrl[:8] == "https://" {
		domain = Text_ExtractBetween(rawUrl, "https://", "/")
	}
	if rawUrl[:7] == "http://" {
		domain = Text_ExtractBetween(rawUrl, "http://", "/")
	}
	return domain
}

// WebUtils_Request 发送 HTTP 请求并返回响应数据。
// 这是网页访问的核心函数，支持多种 HTTP 方法、代理、Cookie 管理等高级功能。
//
// 参数:
//   - url: 请求的完整 URL
//   - method: HTTP 方法，0=GET, 1=POST, 2=HEAD, 3=PUT, 4=OPTIONS, 5=DELETE, 6=TRACE, 7=CONNECT
//   - postData: POST 表单数据（当前未使用，请使用 postBytes 参数）
//   - cookies: 请求时携带的 Cookie 字符串
//   - respCookies: 用于接收响应 Set-Cookie 头的指针
//   - headers: 自定义请求头，每行一个，格式为 "键名:值值"
//   - respContentType: 用于接收响应 Content-Type 头的指针
//   - respStatusCode: 用于接收 HTTP 状态码的指针
//   - noRedirect: true 时不自动跟随重定向
//   - postBytes: POST/PUT 请求的原始字节流数据
//   - proxy: 代理服务器地址，如 "http://127.0.0.1:8080"
//   - timeout: 超时时间（秒），-1 表示不限，0 或小于 1 自动设为 15 秒
//   - proxyUser: 代理认证用户名（当前未使用）
//   - proxyPass: 代理认证密码（当前未使用）
//   - proxyType: 代理类型标识（当前未使用）
//   - inherit: 继承对象（当前未使用）
//   - autoMergeCookie: 是否合并更新 Cookie（当前未使用）
//   - autoCompleteHeaders: 是否自动补全 User-Agent 等必要头（当前未使用）
//   - normalizeHeaders: 是否规范化协议头大小写（当前未使用）
//
// 返回:
//   - []byte: 响应体字节集；请求失败时返回空字节集
func WebUtils_Request(url string, method int, postData string, cookies string, respCookies *string, headers string, respContentType *string, respStatusCode *int, noRedirect bool, postBytes []byte, proxy string, timeout int, proxyUser string, proxyPass string, proxyType int, inherit interface{}, autoMergeCookie bool, autoCompleteHeaders bool, normalizeHeaders bool) []byte {
	client := &http.Client{}
	if timeout != -1 {
		if timeout < 1 {
			timeout = 15
		}
		client.Timeout = time.Duration(timeout) * time.Second
	}

	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	var httpMethod string
	switch method {
	case 0:
		httpMethod = "GET"
	case 1:
		httpMethod = "POST"
	case 2:
		httpMethod = "HEAD"
	case 3:
		httpMethod = "PUT"
	case 4:
		httpMethod = "OPTIONS"
	case 5:
		httpMethod = "DELETE"
	case 6:
		httpMethod = "TRACE"
	case 7:
		httpMethod = "CONNECT"
	default:
		httpMethod = "GET"
	}

	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(postBytes))
	if err != nil {
		return []byte{}
	}

	if noRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	if headers != "" {
		headerList := strings.Split(headers, "\n")
		for _, header := range headerList {
			if strings.TrimSpace(header) != "" {
				name, value := webUtilsParseHeader(header)
				req.Header.Set(name, value)
			}
		}
	}

	if cookies != "" {
		req.Header.Set("Cookie", cookies)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	*respStatusCode = resp.StatusCode
	*respContentType = resp.Header.Get("Content-Type")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}

	if respCookies != nil {
		*respCookies = resp.Header.Get("Set-Cookie")
	}

	return body
}

// webUtilsParseHeader 从协议头字符串中分离键名和值。
// 按第一个冒号分割，如 "Content-Type:application/json" → ("Content-Type", "application/json")。
func webUtilsParseHeader(header string) (string, string) {
	parts := strings.SplitN(header, ":", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return "", ""
}

// WebUtils_GetCookie 从 Cookie 字符串中提取指定名称的单条 Cookie 值。
// Cookie 字符串格式为 "name1=value1; name2=value2"。
//
// 参数:
//   - rawCookies: 完整的 Cookie 字符串
//   - cookieName: 要提取的 Cookie 名称（不区分大小写）
//
// 返回:
//   - string: Cookie 的值；未找到时返回空串
func WebUtils_GetCookie(rawCookies, cookieName string) string {
	cookies := strings.Split(rawCookies, ";")
	for i := 0; i < len(cookies); i++ {
		name := strings.Trim(strings.Split(cookies[i], "=")[0], " ")
		if strings.ToLower(name) == strings.ToLower(cookieName) {
			if strings.Index(strings.Trim(cookies[i], " "), ":") != -1 {
				return strings.TrimSpace(strings.Split(strings.Trim(cookies[i], " "), "=")[1])
			}
			return strings.TrimSpace(strings.Split(strings.Trim(cookies[i], " "), "=")[1])
		}
	}
	return ""
}

// WebUtils_MergeCookies 将新 Cookie 合并到旧 Cookie 中。
// 合并规则：旧 Cookie 中与新 Cookie 同名的项会被新值覆盖，
// 旧 Cookie 中独有的项保留，值为 "=deleted" 的项会被移除，
// 最后去除重复的分号。
//
// 参数:
//   - oldCookie: 原有的 Cookie 字符串
//   - newCookie: 新获取的 Cookie 字符串
//
// 返回:
//   - string: 合并后的 Cookie 字符串
func WebUtils_MergeCookies(oldCookie, newCookie string) string {
	oldCookie = strings.TrimSpace(oldCookie)
	newCookie = strings.TrimSpace(newCookie)

	oldCookieList := strings.Split(oldCookie, ";")
	newCookieList := strings.Split(newCookie, ";")

	for _, cookie := range oldCookieList {
		if !webUtilsArrayContains(newCookieList, webUtilsGetCookieName(cookie)) {
			newCookieList = append(newCookieList, cookie)
		}
	}

	oldCookie = ""
	for _, cookie := range newCookieList {
		if Core_GetTextRight(cookie, 8) != "=deleted" {
			oldCookie += cookie + "; "
		}
	}
	oldCookie = Core_GetTextLeft(oldCookie, len(oldCookie)-2)
	oldCookie = Text_RemoveDuplicateText(oldCookie, "; ")

	return oldCookie
}

// webUtilsArrayContains 检查 Cookie 数组中是否存在指定名称的 Cookie。
func webUtilsArrayContains(arr []string, target string) bool {
	for _, item := range arr {
		if webUtilsGetCookieName(item) == target {
			return true
		}
	}
	return false
}

// webUtilsGetCookieValue 从单条 Cookie 字符串中提取值部分。
func webUtilsGetCookieValue(cookie string) string {
	pos := strings.Index(cookie, "=")
	if pos != -1 {
		result := Core_GetTextRight(cookie, len(cookie)-pos)
		return result
	}
	return ""
}

// webUtilsGetCookieName 从单条 Cookie 字符串中提取名称部分。
func webUtilsGetCookieName(cookie string) string {
	pos := strings.Index(cookie, "=")
	if pos != -1 {
		result := Core_GetTextLeft(cookie, pos-1)
		return Core_TrimSpaces(result)
	}
	return Core_TrimSpaces(cookie)
}

// webUtilsArrayIndexOf 在字符串数组中查找指定值，返回其索引位置。
func webUtilsArrayIndexOf(arr []string, target string) int {
	for i, item := range arr {
		if item == target {
			return i
		}
	}
	return -1
}

// WebUtils_NormalizeHeaders 规范化 HTTP 协议头字符串的格式。
// 将协议头键名中每个以连字符分隔的单词首字母大写（Title Case），
// 如 "content-type" → "Content-Type"，"accept-encoding" → "Accept-Encoding"。
//
// 参数:
//   - rawHeaders: 原始协议头字符串，每行一个，格式为 "键名:值值"
//
// 返回:
//   - string: 格式化后的协议头字符串
func WebUtils_NormalizeHeaders(rawHeaders string) string {
	lines := strings.Split(rawHeaders, "\n")
	result := ""
	for i := 0; i < len(lines); i++ {
		colonPos := strings.Index(lines[i], ":")
		if colonPos == -1 {
			break
		}

		key := strings.TrimSpace(lines[i][:colonPos])
		if strings.Contains(key, "-") {
			keyParts := strings.Split(key, "-")
			total := len(keyParts)
			key = ""
			for x := 0; x < total; x++ {
				if x == total-1 {
					key += strings.Title(keyParts[x])
				} else {
					key += strings.Title(keyParts[x]) + "-"
				}
			}
		} else {
			key = strings.Title(key)
		}

		value := strings.TrimSpace(lines[i][colonPos+1:])
		if len(value) > 0 && value[0] == ' ' {
			value = value[1:]
		}

		result += key + value + "\n"
	}

	return strings.TrimRight(result, "\n")
}

// Web访问_对象 发送 HTTP 请求并返回响应数据（无前缀版本）。
// 功能与 WebUtils_Request 完全相同，为使用习惯提供别名。
func Web访问_对象(url string, method int, postData string, cookies string, respCookies *string, headers string, respContentType *string, respStatusCode *int, noRedirect bool, postBytes []byte, proxy string, timeout int, proxyUser string, proxyPass string, proxyType int, inherit interface{}, autoMergeCookie bool, autoCompleteHeaders bool, normalizeHeaders bool) []byte {
	client := &http.Client{}
	if timeout != -1 {
		if timeout < 1 {
			timeout = 15
		}
		client.Timeout = time.Duration(timeout) * time.Second
	}

	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	var httpMethod string
	switch method {
	case 0:
		httpMethod = "GET"
	case 1:
		httpMethod = "POST"
	case 2:
		httpMethod = "HEAD"
	case 3:
		httpMethod = "PUT"
	case 4:
		httpMethod = "OPTIONS"
	case 5:
		httpMethod = "DELETE"
	case 6:
		httpMethod = "TRACE"
	case 7:
		httpMethod = "CONNECT"
	default:
		httpMethod = "GET"
	}

	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(postBytes))
	if err != nil {
		return []byte{}
	}

	if noRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	if headers != "" {
		headerList := strings.Split(headers, "\n")
		for _, header := range headerList {
			if strings.TrimSpace(header) != "" {
				name, value := webUtilsParseHeader(header)
				req.Header.Set(name, value)
			}
		}
	}

	if cookies != "" {
		req.Header.Set("Cookie", cookies)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}
	}
	defer resp.Body.Close()

	*respStatusCode = resp.StatusCode
	*respContentType = resp.Header.Get("Content-Type")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}

	if respCookies != nil {
		*respCookies = resp.Header.Get("Set-Cookie")
	}

	return body
}