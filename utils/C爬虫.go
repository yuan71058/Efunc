package utils

import (
	"github.com/gocolly/colly/v2"
)

// C爬虫_创建 创建一个新的网页采集器。
// 基于 colly 库实现，支持回调式的网页爬取。
// 通过注册回调函数处理不同类型的 HTML 元素和事件。
//
// 返回:
//   - *colly.Collector: 采集器实例
func C爬虫_创建() *colly.Collector {
	return colly.NewCollector()
}

// C爬虫_访问 访问指定的 URL。
// 采集器会发送 HTTP 请求并触发已注册的回调函数。
//
// 参数:
//   - 采集器: 采集器实例
//   - 网址: 要访问的 URL
//
// 返回:
//   - error: 访问失败时返回错误
func C爬虫_访问(采集器 *colly.Collector, 网址 string) error {
	return 采集器.Visit(网址)
}

// C爬虫_注册HTML回调 注册 HTML 元素选择器回调。
// 当采集器发现匹配 CSS 选择器的元素时，调用回调函数。
// 示例: C爬虫_注册HTML回调(c, "div.title", func(e *colly.HTMLElement) {...})
//
// 参数:
//   - 采集器: 采集器实例
//   - 选择器: CSS 选择器，如 "a[href]"、"div.content"
//   - 回调函数: 匹配元素时的回调，参数为 HTMLElement
func C爬虫_注册HTML回调(采集器 *colly.Collector, 选择器 string, 回调函数 func(e *colly.HTMLElement)) {
	采集器.OnHTML(选择器, 回调函数)
}

// C爬虫_注册请求回调 注册 HTTP 请求回调。
// 每次发送请求前调用，可用于修改请求头、添加认证等。
//
// 参数:
//   - 采集器: 采集器实例
//   - 回调函数: 请求回调，参数为 *colly.Request
func C爬虫_注册请求回调(采集器 *colly.Collector, 回调函数 func(r *colly.Request)) {
	采集器.OnRequest(回调函数)
}

// C爬虫_注册响应回调 注册 HTTP 响应回调。
// 每次收到响应时调用，可用于检查状态码、记录日志等。
//
// 参数:
//   - 采集器: 采集器实例
//   - 回调函数: 响应回调，参数为 *colly.Response
func C爬虫_注册响应回调(采集器 *colly.Collector, 回调函数 func(r *colly.Response)) {
	采集器.OnResponse(回调函数)
}

// C爬虫_注册错误回调 注册错误回调。
// 请求失败时调用，可用于错误处理和重试逻辑。
//
// 参数:
//   - 采集器: 采集器实例
//   - 回调函数: 错误回调，参数为 *colly.Response 和 error
func C爬虫_注册错误回调(采集器 *colly.Collector, 回调函数 func(r *colly.Response, err error)) {
	采集器.OnError(回调函数)
}

// C爬虫_限制并发 设置采集器的并发限制和请求间隔。
// 防止请求过于频繁被目标网站封禁。
//
// 参数:
//   - 采集器: 采集器实例
//   - 域名: 限制规则适用的域名
//   - 并发数: 最大并发请求数
//   - 间隔毫秒: 同一域名两次请求之间的间隔（毫秒）
func C爬虫_限制并发(采集器 *colly.Collector, 域名 string, 并发数 int, 间隔毫秒 int) {
	采集器.Limit(&colly.LimitRule{
		DomainGlob:  域名,
		Parallelism: 并发数,
		Delay:       0,
	})
}

// C爬虫_设置请求头 设置采集器的默认请求头。
//
// 参数:
//   - 采集器: 采集器实例
//   - 请求头: HTTP 请求头键值对
func C爬虫_设置请求头(采集器 *colly.Collector, 请求头 map[string]string) {
	采集器.OnRequest(func(r *colly.Request) {
		for k, v := range 请求头 {
			r.Headers.Set(k, v)
		}
	})
}

// C爬虫_取元素文本 从 HTMLElement 中提取文本内容。
// 是 colly.HTMLElement.Text 的便捷封装。
//
// 参数:
//   - 元素: HTML 元素对象
//
// 返回:
//   - string: 元素的文本内容
func C爬虫_取元素文本(元素 *colly.HTMLElement) string {
	return 元素.Text
}

// C爬虫_取元素属性 从 HTMLElement 中提取指定属性的值。
//
// 参数:
//   - 元素: HTML 元素对象
//   - 属性名: HTML 属性名，如 "href"、"src"、"class"
//
// 返回:
//   - string: 属性值；属性不存在时返回空串
func C爬虫_取元素属性(元素 *colly.HTMLElement, 属性名 string) string {
	return 元素.Attr(属性名)
}
