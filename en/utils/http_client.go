// HTTP 客户端模块
// 基于 resty 库实现，提供链式调用风格的 HTTP 请求 API。
// 支持 JSON/XML 自动序列化、请求重试、代理等高级功能。
package utils

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// HTTPClient_New 创建一个新的 HTTP 客户端实例。
// 基于 resty 库实现，提供链式调用风格的 HTTP 请求 API。
// 支持 JSON/XML 自动序列化、请求重试、代理等高级功能。
//
// 返回:
//   - *resty.Request: 可链式调用的请求对象
func HTTPClient_New() *resty.Request {
	return resty.New().R()
}

// HTTPClient_Get 发送 HTTP GET 请求。
//
// 参数:
//   - url: 请求的 URL 地址
//
// 返回:
//   - *resty.Response: 响应对象，包含状态码、响应体等信息
//   - error: 请求失败时返回错误
func HTTPClient_Get(url string) (*resty.Response, error) {
	return resty.New().R().Get(url)
}

// HTTPClient_Post 发送 HTTP POST 请求（JSON 格式）。
// 自动将数据序列化为 JSON 并设置 Content-Type 头。
//
// 参数:
//   - url: 请求的 URL 地址
//   - data: 要发送的数据，会自动序列化为 JSON
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func HTTPClient_Post(url string, data interface{}) (*resty.Response, error) {
	return resty.New().R().SetBody(data).Post(url)
}

// HTTPClient_Put 发送 HTTP PUT 请求（JSON 格式）。
//
// 参数:
//   - url: 请求的 URL 地址
//   - data: 要发送的数据
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func HTTPClient_Put(url string, data interface{}) (*resty.Response, error) {
	return resty.New().R().SetBody(data).Put(url)
}

// HTTPClient_Delete 发送 HTTP DELETE 请求。
//
// 参数:
//   - url: 请求的 URL 地址
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func HTTPClient_Delete(url string) (*resty.Response, error) {
	return resty.New().R().Delete(url)
}

// HTTPClient_GetText 发送 GET 请求并返回响应体文本。
//
// 参数:
//   - url: 请求的 URL 地址
//
// 返回:
//   - string: 响应体文本
//   - error: 请求失败时返回错误
func HTTPClient_GetText(url string) (string, error) {
	resp, err := resty.New().R().Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// HTTPClient_GetWithHeaders 带自定义请求头发送 GET 请求。
//
// 参数:
//   - url: 请求的 URL 地址
//   - headers: HTTP 请求头键值对
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func HTTPClient_GetWithHeaders(url string, headers map[string]string) (*resty.Response, error) {
	return resty.New().R().SetHeaders(headers).Get(url)
}

// HTTPClient_GetWithParams 带查询参数发送 GET 请求。
//
// 参数:
//   - url: 请求的 URL 地址
//   - params: URL 查询参数键值对
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func HTTPClient_GetWithParams(url string, params map[string]string) (*resty.Response, error) {
	return resty.New().R().SetQueryParams(params).Get(url)
}

// HTTPClient_SetTimeout 创建带超时设置的 HTTP 客户端。
// 超时后请求会自动取消并返回错误。
//
// 参数:
//   - timeoutSec: 请求超时时间（秒）
//
// 返回:
//   - *resty.Client: 配置好超时的客户端实例
func HTTPClient_SetTimeout(timeoutSec int) *resty.Client {
	return resty.New().SetTimeout(time.Duration(timeoutSec) * time.Second)
}

// HTTPClient_SetProxy 创建带代理设置的 HTTP 客户端。
//
// 参数:
//   - proxyAddr: 代理服务器地址，如 "http://127.0.0.1:7890"
//
// 返回:
//   - *resty.Client: 配置好代理的客户端实例
func HTTPClient_SetProxy(proxyAddr string) *resty.Client {
	return resty.New().SetProxy(proxyAddr)
}

// HTTPClient_PostForm 以表单形式提交 POST 请求（application/x-www-form-urlencoded）。
//
// 参数:
//   - url: 请求的 URL 地址
//   - formData: 表单字段键值对
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func HTTPClient_PostForm(url string, formData map[string]string) (*resty.Response, error) {
	return resty.New().R().SetFormData(formData).Post(url)
}

// HTTPClient_DownloadFile 下载文件到指定路径。
//
// 参数:
//   - url: 文件下载 URL
//   - savePath: 本地保存路径
//
// 返回:
//   - error: 下载失败时返回错误
func HTTPClient_DownloadFile(url string, savePath string) error {
	_, err := resty.New().R().SetOutput(savePath).Get(url)
	return err
}