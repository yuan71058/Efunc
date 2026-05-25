// Spider/网页采集模块
// 基于 colly 库实现，提供回调式的网页爬取功能。
// 通过注册回调函数处理不同类型的 HTML 元素和事件。
package utils

import (
	"github.com/gocolly/colly/v2"
)

// Spider_New 创建一个新的网页采集器。
// 基于 colly 库实现，支持回调式的网页爬取。
// 通过注册回调函数处理不同类型的 HTML 元素和事件。
//
// 返回:
//   - *colly.Collector: 采集器实例
func Spider_New() *colly.Collector {
	return colly.NewCollector()
}

// Spider_Visit 访问指定的 URL。
// 采集器会发送 HTTP 请求并触发已注册的回调函数。
//
// 参数:
//   - collector: 采集器实例
//   - url: 要访问的 URL
//
// 返回:
//   - error: 访问失败时返回错误
func Spider_Visit(collector *colly.Collector, url string) error {
	return collector.Visit(url)
}

// Spider_OnHTML 注册 HTML 元素选择器回调。
// 当采集器发现匹配 CSS 选择器的元素时，调用回调函数。
// 示例: Spider_OnHTML(c, "div.title", func(e *colly.HTMLElement) {...})
//
// 参数:
//   - collector: 采集器实例
//   - selector: CSS 选择器，如 "a[href]"、"div.content"
//   - callback: 匹配元素时的回调，参数为 HTMLElement
func Spider_OnHTML(collector *colly.Collector, selector string, callback func(e *colly.HTMLElement)) {
	collector.OnHTML(selector, callback)
}

// Spider_OnRequest 注册 HTTP 请求回调。
// 每次发送请求前调用，可用于修改请求头、添加认证等。
//
// 参数:
//   - collector: 采集器实例
//   - callback: 请求回调，参数为 *colly.Request
func Spider_OnRequest(collector *colly.Collector, callback func(r *colly.Request)) {
	collector.OnRequest(callback)
}

// Spider_OnResponse 注册 HTTP 响应回调。
// 每次收到响应时调用，可用于检查状态码、记录日志等。
//
// 参数:
//   - collector: 采集器实例
//   - callback: 响应回调，参数为 *colly.Response
func Spider_OnResponse(collector *colly.Collector, callback func(r *colly.Response)) {
	collector.OnResponse(callback)
}

// Spider_OnError 注册错误回调。
// 请求失败时调用，可用于错误处理和重试逻辑。
//
// 参数:
//   - collector: 采集器实例
//   - callback: 错误回调，参数为 *colly.Response 和 error
func Spider_OnError(collector *colly.Collector, callback func(r *colly.Response, err error)) {
	collector.OnError(callback)
}

// Spider_LimitConcurrency 设置采集器的并发限制和请求间隔。
// 防止请求过于频繁被目标网站封禁。
//
// 参数:
//   - collector: 采集器实例
//   - domain: 限制规则适用的域名
//   - parallelism: 最大并发请求数
//   - delayMs: 同一域名两次请求之间的间隔（毫秒）
func Spider_LimitConcurrency(collector *colly.Collector, domain string, parallelism int, delayMs int) {
	collector.Limit(&colly.LimitRule{
		DomainGlob:  domain,
		Parallelism: parallelism,
		Delay:       0,
	})
}

// Spider_SetHeaders 设置采集器的默认请求头。
//
// 参数:
//   - collector: 采集器实例
//   - headers: HTTP 请求头键值对
func Spider_SetHeaders(collector *colly.Collector, headers map[string]string) {
	collector.OnRequest(func(r *colly.Request) {
		for k, v := range headers {
			r.Headers.Set(k, v)
		}
	})
}

// Spider_GetElementText 从 HTMLElement 中提取文本内容。
// 是 colly.HTMLElement.Text 的便捷封装。
//
// 参数:
//   - elem: HTML 元素对象
//
// 返回:
//   - string: 元素的文本内容
func Spider_GetElementText(elem *colly.HTMLElement) string {
	return elem.Text
}

// Spider_GetElementAttr 从 HTMLElement 中提取指定属性的值。
//
// 参数:
//   - elem: HTML 元素对象
//   - attrName: HTML 属性名，如 "href"、"src"、"class"
//
// 返回:
//   - string: 属性值；属性不存在时返回空串
func Spider_GetElementAttr(elem *colly.HTMLElement, attrName string) string {
	return elem.Attr(attrName)
}