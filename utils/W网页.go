package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// W网页_取域名 从 URL 中提取域名部分。
// 支持 http:// 和 https:// 两种协议前缀的 URL。
//
// 参数:
//   - Url: 完整的 URL 地址，如 "https://www.example.com/path"
//
// 返回:
//   - string: 域名部分，如 "www.example.com"；无法识别时返回空串
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

// W网页_访问_对象 发送 HTTP 请求并返回响应数据。
// 这是网页访问的核心函数，支持多种 HTTP 方法、代理、Cookie 管理等高级功能。
//
// 参数:
//   - 网址: 请求的完整 URL
//   - 访问方式: HTTP 方法，0=GET, 1=POST, 2=HEAD, 3=PUT, 4=OPTIONS, 5=DELETE, 6=TRACE, 7=CONNECT
//   - 提交信息: POST 表单数据（当前未使用，请使用 字节集提交 参数）
//   - 提交Cookies: 请求时携带的 Cookie 字符串
//   - 返回Cookies: 用于接收响应 Set-Cookie 头的指针
//   - 附加协议头: 自定义请求头，每行一个，格式为 "键名:值值"
//   - 返回协议头: 用于接收响应 Content-Type 头的指针
//   - 返回状态代码: 用于接收 HTTP 状态码的指针
//   - 禁止重定向: true 时不自动跟随重定向
//   - 字节集提交: POST/PUT 请求的原始字节流数据
//   - 代理地址: 代理服务器地址，如 "http://127.0.0.1:8080"
//   - 超时: 超时时间（秒），-1 表示不限，0 或小于 1 自动设为 15 秒
//   - 代理用户名: 代理认证用户名（当前未使用）
//   - 代理密码: 代理认证密码（当前未使用）
//   - 代理标识: 代理类型标识（当前未使用）
//   - 对象继承: 继承对象（当前未使用）
//   - 是否自动合并更新Cookie: 是否合并更新 Cookie（当前未使用）
//   - 是否补全必要协议头: 是否自动补全 User-Agent 等必要头（当前未使用）
//   - 是否处理协议头大小写: 是否规范化协议头大小写（当前未使用）
//
// 返回:
//   - []byte: 响应体字节集；请求失败时返回空字节集
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

// 内部_协议头取名值 从协议头字符串中分离键名和值。
// 按第一个冒号分割，如 "Content-Type:application/json" → ("Content-Type", "application/json")。
//
// 参数:
//   - 协议头: 协议头字符串，格式为 "键名:值值"
//
// 返回:
//   - string: 键名
//   - string: 键值；格式不正确时返回空串
func 内部_协议头取名值(协议头 string) (string, string) {
	头名值 := strings.SplitN(协议头, ":", 2)
	if len(头名值) == 2 {
		return 头名值[0], 头名值[1]
	}
	return "", ""
}

// Q取单条Cookie 从 Cookie 字符串中提取指定名称的单条 Cookie 值。
// Cookie 字符串格式为 "name1=value1; name2=value2"。
//
// 参数:
//   - 原Cookies: 完整的 Cookie 字符串
//   - 单条Cookie名称: 要提取的 Cookie 名称（不区分大小写）
//
// 返回:
//   - string: Cookie 的值；未找到时返回空串
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

// W网页_Cookie合并更新 将新 Cookie 合并到旧 Cookie 中。
// 合并规则：旧 Cookie 中与新 Cookie 同名的项会被新值覆盖，
// 旧 Cookie 中独有的项保留，值为 "=deleted" 的项会被移除，
// 最后去除重复的分号。
//
// 参数:
//   - 旧Cookie: 原有的 Cookie 字符串
//   - 新Cookie: 新获取的 Cookie 字符串
//
// 返回:
//   - string: 合并后的 Cookie 字符串
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
		if F_取文本右边(cookie, 8) != "=deleted" {
			旧Cookie += cookie + "; "
		}
	}
	旧Cookie = F_取文本左边(旧Cookie, len(旧Cookie)-2)
	旧Cookie = W文本_去重复文本(旧Cookie, "; ")

	return 旧Cookie
}

// 内部_数组成员是否存在1 检查 Cookie 数组中是否存在指定名称的 Cookie。
// 通过 内部_Cookie取名 提取每个 Cookie 的名称进行比较。
//
// 参数:
//   - 数组: Cookie 字符串数组
//   - 要判断值: 要查找的 Cookie 名称
//
// 返回:
//   - bool: 存在返回 true，否则返回 false
func 内部_数组成员是否存在1(数组 []string, 要判断值 string) bool {
	for _, 成员 := range 数组 {
		if 内部_Cookie取名(成员) == 要判断值 {
			return true
		}
	}
	return false
}

// 内部_Cookie取值 从单条 Cookie 字符串中提取值部分。
// 如 "name=value" → "value"。
//
// 参数:
//   - Cookie: 单条 Cookie 字符串
//
// 返回:
//   - string: Cookie 的值；无等号时返回空串
func 内部_Cookie取值(Cookie string) string {
	位置 := strings.Index(Cookie, "=")
	if 位置 != -1 {
		结果 := F_取文本右边(Cookie, len(Cookie)-位置)
		return 结果
	}
	return ""
}

// 内部_Cookie取名 从单条 Cookie 字符串中提取名称部分。
// 如 "name=value" → "name"，会去除首尾空格。
//
// 参数:
//   - Cookie: 单条 Cookie 字符串
//
// 返回:
//   - string: Cookie 的名称；无等号时返回去除空格后的完整字符串
func 内部_Cookie取名(Cookie string) string {
	位置 := strings.Index(Cookie, "=")
	if 位置 != -1 {
		结果 := F_取文本左边(Cookie, 位置-1)
		return F_删首尾空(结果)
	}
	return F_删首尾空(Cookie)
}

// 内部_数组成员是否存在_文本 在字符串数组中查找指定值，返回其索引位置。
//
// 参数:
//   - 数组: 源字符串数组
//   - 要判断值: 要查找的值
//
// 返回:
//   - int: 找到的索引（从 0 开始）；未找到返回 -1
func 内部_数组成员是否存在_文本(数组 []string, 要判断值 string) int {
	for i, 成员 := range 数组 {
		if 成员 == 要判断值 {
			return i
		}
	}
	return -1
}

// W网页_处理协议头 规范化 HTTP 协议头字符串的格式。
// 将协议头键名中每个以连字符分隔的单词首字母大写（Title Case），
// 如 "content-type" → "Content-Type"，"accept-encoding" → "Accept-Encoding"。
//
// 参数:
//   - 原始协议头: 原始协议头字符串，每行一个，格式为 "键名:值值"
//
// 返回:
//   - string: 格式化后的协议头字符串
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

// 网页_访问_对象 发送 HTTP 请求并返回响应数据（无 W 前缀版本）。
// 功能与 W网页_访问_对象 完全相同，为使用习惯提供别名。
//
// 参数:
//   - 网址: 请求的完整 URL
//   - 访问方式: HTTP 方法，0=GET, 1=POST, 2=HEAD, 3=PUT, 4=OPTIONS, 5=DELETE, 6=TRACE, 7=CONNECT
//   - 提交信息: POST 表单数据（当前未使用）
//   - 提交Cookies: 请求时携带的 Cookie 字符串
//   - 返回Cookies: 用于接收响应 Set-Cookie 头的指针
//   - 附加协议头: 自定义请求头，每行一个
//   - 返回协议头: 用于接收响应 Content-Type 头的指针
//   - 返回状态代码: 用于接收 HTTP 状态码的指针
//   - 禁止重定向: true 时不自动跟随重定向
//   - 字节集提交: POST/PUT 请求的原始字节流数据
//   - 代理地址: 代理服务器地址
//   - 超时: 超时时间（秒），-1 表示不限
//   - 代理用户名: 代理认证用户名（当前未使用）
//   - 代理密码: 代理认证密码（当前未使用）
//   - 代理标识: 代理类型标识（当前未使用）
//   - 对象继承: 继承对象（当前未使用）
//   - 是否自动合并更新Cookie: 是否合并更新 Cookie（当前未使用）
//   - 是否补全必要协议头: 是否自动补全必要头（当前未使用）
//   - 是否处理协议头大小写: 是否规范化协议头大小写（当前未使用）
//
// 返回:
//   - []byte: 响应体字节集；请求失败时返回空字节集
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
