package utils

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// H客户端_创建 创建一个新的 HTTP 客户端实例。
// 基于 resty 库实现，提供链式调用风格的 HTTP 请求 API。
// 支持 JSON/XML 自动序列化、请求重试、代理等高级功能。
//
// 返回:
//   - *resty.Request: 可链式调用的请求对象
func H客户端_创建() *resty.Request {
	return resty.New().R()
}

// H客户端_Get 发送 HTTP GET 请求。
//
// 参数:
//   - 网址: 请求的 URL 地址
//
// 返回:
//   - *resty.Response: 响应对象，包含状态码、响应体等信息
//   - error: 请求失败时返回错误
func H客户端_Get(网址 string) (*resty.Response, error) {
	return resty.New().R().Get(网址)
}

// H客户端_Post 发送 HTTP POST 请求（JSON 格式）。
// 自动将数据序列化为 JSON 并设置 Content-Type 头。
//
// 参数:
//   - 网址: 请求的 URL 地址
//   - 数据: 要发送的数据，会自动序列化为 JSON
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func H客户端_Post(网址 string, 数据 interface{}) (*resty.Response, error) {
	return resty.New().R().SetBody(数据).Post(网址)
}

// H客户端_Put 发送 HTTP PUT 请求（JSON 格式）。
//
// 参数:
//   - 网址: 请求的 URL 地址
//   - 数据: 要发送的数据
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func H客户端_Put(网址 string, 数据 interface{}) (*resty.Response, error) {
	return resty.New().R().SetBody(数据).Put(网址)
}

// H客户端_Delete 发送 HTTP DELETE 请求。
//
// 参数:
//   - 网址: 请求的 URL 地址
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func H客户端_Delete(网址 string) (*resty.Response, error) {
	return resty.New().R().Delete(网址)
}

// H客户端_取文本 发送 GET 请求并返回响应体文本。
//
// 参数:
//   - 网址: 请求的 URL 地址
//
// 返回:
//   - string: 响应体文本
//   - error: 请求失败时返回错误
func H客户端_取文本(网址 string) (string, error) {
	resp, err := resty.New().R().Get(网址)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// H客户端_带请求头发送 带自定义请求头发送 GET 请求。
//
// 参数:
//   - 网址: 请求的 URL 地址
//   - 请求头: HTTP 请求头键值对
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func H客户端_带请求头发送(网址 string, 请求头 map[string]string) (*resty.Response, error) {
	return resty.New().R().SetHeaders(请求头).Get(网址)
}

// H客户端_带参数发送 带查询参数发送 GET 请求。
//
// 参数:
//   - 网址: 请求的 URL 地址
//   - 参数: URL 查询参数键值对
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func H客户端_带参数发送(网址 string, 参数 map[string]string) (*resty.Response, error) {
	return resty.New().R().SetQueryParams(参数).Get(网址)
}

// H客户端_设置超时 创建带超时设置的 HTTP 客户端。
// 超时后请求会自动取消并返回错误。
//
// 参数:
//   - 超时秒数: 请求超时时间（秒）
//
// 返回:
//   - *resty.Client: 配置好超时的客户端实例
func H客户端_设置超时(超时秒数 int) *resty.Client {
	return resty.New().SetTimeout(time.Duration(超时秒数) * time.Second)
}

// H客户端_设置代理 创建带代理设置的 HTTP 客户端。
//
// 参数:
//   - 代理地址: 代理服务器地址，如 "http://127.0.0.1:7890"
//
// 返回:
//   - *resty.Client: 配置好代理的客户端实例
func H客户端_设置代理(代理地址 string) *resty.Client {
	return resty.New().SetProxy(代理地址)
}

// H客户端_表单提交 以表单形式提交 POST 请求（application/x-www-form-urlencoded）。
//
// 参数:
//   - 网址: 请求的 URL 地址
//   - 表单数据: 表单字段键值对
//
// 返回:
//   - *resty.Response: 响应对象
//   - error: 请求失败时返回错误
func H客户端_表单提交(网址 string, 表单数据 map[string]string) (*resty.Response, error) {
	return resty.New().R().SetFormData(表单数据).Post(网址)
}

// H客户端_下载文件 下载文件到指定路径。
//
// 参数:
//   - 网址: 文件下载 URL
//   - 保存路径: 本地保存路径
//
// 返回:
//   - error: 下载失败时返回错误
func H客户端_下载文件(网址 string, 保存路径 string) error {
	_, err := resty.New().R().SetOutput(保存路径).Get(网址)
	return err
}
