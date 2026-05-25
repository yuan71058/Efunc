package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTP_GetDomain 从 URL 中提取域名部分。
// 支持 http:// 和 https:// 两种协议前缀的 URL。
//
// 参数:
//   - rawURL: 完整的 URL 地址，如 "https://www.example.com/path"
//
// 返回:
//   - string: 域名部分，如 "www.example.com"；无法识别时返回空串
func HTTP_GetDomain(rawURL string) string {
	var domain string

	if strings.HasPrefix(rawURL, "https://") {
		domain = Text_Mid(rawURL, "https://", "/")
	}
	if strings.HasPrefix(rawURL, "http://") {
		domain = Text_Mid(rawURL, "http://", "/")
	}
	return domain
}

// HTTP_Request 发送 HTTP 请求并返回响应数据。
// 这是网页访问的核心函数，支持多种 HTTP 方法、代理、Cookie 管理等高级功能。
//
// 参数:
//   - urlStr: 请求的完整 URL
//   - method: HTTP 方法，0=GET, 1=POST, 2=HEAD, 3=PUT, 4=OPTIONS, 5=DELETE, 6=TRACE, 7=CONNECT
//   - formData: POST 表单数据（当前未使用，请使用 postBody 参数）
//   - cookies: 请求时携带的 Cookie 字符串
//   - respCookies: 用于接收响应 Set-Cookie 头的指针
//   - headers: 自定义请求头，每行一个，格式为 "键名:值值"
//   - respContentType: 用于接收响应 Content-Type 头的指针
//   - statusCode: 用于接收 HTTP 状态码的指针
//   - noRedirect: true 时不自动跟随重定向
//   - postBody: POST/PUT 请求的原始字节流数据
//   - proxy: 代理服务器地址，如 "http://127.0.0.1:8080"
//   - timeout: 超时时间（秒），-1 表示不限，0 或小于 1 自动设为 15 秒
//   - proxyUser: 代理认证用户名（当前未使用）
//   - proxyPass: 代理认证密码（当前未使用）
//   - proxyType: 代理类型标识（当前未使用）
//   - inheritedObject: 继承对象（当前未使用）
//   - autoMergeCookies: 是否合并更新 Cookie（当前未使用）
//   - autoFillHeaders: 是否自动补全 User-Agent 等必要头（当前未使用）
//   - normalizeHeaders: 是否规范化协议头大小写（当前未使用）
//
// 返回:
//   - []byte: 响应体字节集；请求失败时返回空字节集
func HTTP_Request(urlStr string, method int, formData string, cookies string, respCookies *string, headers string, respContentType *string, statusCode *int, noRedirect bool, postBody []byte, proxy string, timeout int, proxyUser string, proxyPass string, proxyType int, inheritedObject interface{}, autoMergeCookies bool, autoFillHeaders bool, normalizeHeaders bool) []byte {
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

	var methodStr string
	switch method {
	case 0:
		methodStr = "GET"
	case 1:
		methodStr = "POST"
	case 2:
		methodStr = "HEAD"
	case 3:
		methodStr = "PUT"
	case 4:
		methodStr = "OPTIONS"
	case 5:
		methodStr = "DELETE"
	case 6:
		methodStr = "TRACE"
	case 7:
		methodStr = "CONNECT"
	default:
		methodStr = "GET"
	}

	req, err := http.NewRequest(methodStr, urlStr, bytes.NewBuffer(postBody))
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
				key, value := parseHeaderKeyValue(header)
				req.Header.Set(key, value)
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

	if statusCode != nil {
		*statusCode = resp.StatusCode
	}
	if respContentType != nil {
		*respContentType = resp.Header.Get("Content-Type")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}

	if respCookies != nil {
		*respCookies = resp.Header.Get("Set-Cookie")
	}

	return body
}

// parseHeaderKeyValue 从协议头字符串中分离键名和值。
// 按第一个冒号分割，如 "Content-Type:application/json" → ("Content-Type", "application/json")。
//
// 参数:
//   - header: 协议头字符串，格式为 "键名:值值"
//
// 返回:
//   - string: 键名
//   - string: 键值；格式不正确时返回空串
func parseHeaderKeyValue(header string) (string, string) {
	parts := strings.SplitN(header, ":", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return "", ""
}

// HTTP_GetSingleCookie 从 Cookie 字符串中提取指定名称的单条 Cookie 值。
// Cookie 字符串格式为 "name1=value1; name2=value2"。
//
// 参数:
//   - cookieStr: 完整的 Cookie 字符串
//   - cookieName: 要提取的 Cookie 名称（不区分大小写）
//
// 返回:
//   - string: Cookie 的值；未找到时返回空串
func HTTP_GetSingleCookie(cookieStr, cookieName string) string {
	cookies := strings.Split(cookieStr, ";")
	for i := 0; i < len(cookies); i++ {
		name := strings.Trim(strings.Split(cookies[i], "=")[0], " ")
		if strings.EqualFold(name, cookieName) {
			parts := strings.SplitN(strings.TrimSpace(cookies[i]), "=", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

// HTTP_MergeCookies 将新 Cookie 合并到旧 Cookie 中。
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
func HTTP_MergeCookies(oldCookie, newCookie string) string {
	oldCookie = strings.TrimSpace(oldCookie)
	newCookie = strings.TrimSpace(newCookie)

	oldCookies := strings.Split(oldCookie, ";")
	newCookies := strings.Split(newCookie, ";")

	for _, cookie := range oldCookies {
		if !cookieNameExistsInSlice(newCookies, parseCookieName(cookie)) {
			newCookies = append(newCookies, cookie)
		}
	}

	oldCookie = ""
	for _, cookie := range newCookies {
		if !strings.HasSuffix(cookie, "=deleted") {
			oldCookie += cookie + "; "
		}
	}
	oldCookie = strings.TrimSuffix(oldCookie, "; ")
	oldCookie = Text_Deduplicate(oldCookie, "; ")

	return oldCookie
}

// cookieNameExistsInSlice 检查 Cookie 数组中是否存在指定名称的 Cookie。
func cookieNameExistsInSlice(cookies []string, name string) bool {
	for _, cookie := range cookies {
		if parseCookieName(cookie) == name {
			return true
		}
	}
	return false
}

// parseCookieValue 从单条 Cookie 字符串中提取值部分。
// 如 "name=value" → "value"。
func parseCookieValue(cookie string) string {
	pos := strings.Index(cookie, "=")
	if pos != -1 {
		return cookie[pos+1:]
	}
	return ""
}

// parseCookieName 从单条 Cookie 字符串中提取名称部分。
// 如 "name=value" → "name"，会去除首尾空格。
func parseCookieName(cookie string) string {
	pos := strings.Index(cookie, "=")
	if pos != -1 {
		return strings.TrimSpace(cookie[:pos])
	}
	return strings.TrimSpace(cookie)
}

// HTTP_NormalizeHeaders 规范化 HTTP 协议头字符串的格式。
// 将协议头键名中每个以连字符分隔的单词首字母大写（Title Case），
// 如 "content-type" → "Content-Type"，"accept-encoding" → "Accept-Encoding"。
//
// 参数:
//   - rawHeaders: 原始协议头字符串，每行一个，格式为 "键名:值值"
//
// 返回:
//   - string: 格式化后的协议头字符串
func HTTP_NormalizeHeaders(rawHeaders string) string {
	lines := strings.Split(rawHeaders, "\n")
	headers := ""
	for i := 0; i < len(lines); i++ {
		colonPos := strings.Index(lines[i], ":")
		if colonPos == -1 {
			break
		}

		key := strings.TrimSpace(lines[i][:colonPos])
		if strings.Contains(key, "-") {
			keyParts := strings.Split(key, "-")
			keyResult := ""
			total := len(keyParts)
			for x := 0; x < total; x++ {
				if x == total-1 {
					keyResult += strings.Title(keyParts[x])
				} else {
					keyResult += strings.Title(keyParts[x]) + "-"
				}
			}
			key = keyResult
		} else {
			key = strings.Title(key)
		}

		value := strings.TrimSpace(lines[i][colonPos+1:])
		if len(value) > 0 && value[0] == ' ' {
			value = value[1:]
		}

		headers += key + ":" + value + "\n"
	}

	return strings.TrimRight(headers, "\n")
}