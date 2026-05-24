package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var 全局校验器 = validator.New()

// V校验_验证结构体 验证结构体字段的标签约束。
// 结构体字段需使用 validate 标签定义校验规则，如 `validate:"required,email"`。
// 常用标签: required(必填), email(邮箱), min/max(最小/最大值), len(长度), oneof(枚举值)。
//
// 参数:
//   - 结构体: 待验证的结构体实例（非指针）
//
// 返回:
//   - error: 验证失败时返回详细的字段错误信息；通过时返回 nil
func V校验_验证结构体(结构体 interface{}) error {
	return 全局校验器.Struct(结构体)
}

// V校验_验证变量 验证单个变量的值是否符合标签约束。
// 适用于不需要定义结构体的简单校验场景。
//
// 参数:
//   - 变量: 待验证的变量值
//   - 标签: 校验规则标签，如 "required,email"、"min=1,max=100"
//
// 返回:
//   - error: 验证失败时返回错误信息；通过时返回 nil
func V校验_验证变量(变量 interface{}, 标签 string) error {
	return 全局校验器.Var(变量, 标签)
}

// V校验_验证字段 验证结构体中指定字段的值。
// 仅验证单个字段，不验证整个结构体。
//
// 参数:
//   - 结构体: 包含待验证字段的结构体
//   - 字段名: 要验证的字段名（结构体中的字段名，非 JSON 名）
//   - 标签: 校验规则标签
//
// 返回:
//   - error: 验证失败时返回错误信息
func V校验_验证字段(结构体 interface{}, 字段名 string, 标签 string) error {
	return 全局校验器.VarWithValue(结构体, 字段名, 标签)
}

// V校验_取错误信息 从验证错误中提取可读的错误信息。
// 将 validator.ValidationErrors 转换为中文友好的错误描述。
//
// 参数:
//   - 错误: V校验_验证结构体 返回的错误
//
// 返回:
//   - string: 格式化的错误信息，如 "字段 Email 验证 required 失败"
func V校验_取错误信息(错误 error) string {
	if 错误 == nil {
		return ""
	}
	errs, ok := 错误.(validator.ValidationErrors)
	if !ok {
		return 错误.Error()
	}
	result := ""
	for _, e := range errs {
		result += fmt.Sprintf("字段 %s 验证 %s 失败, 当前值: %v; ", e.Field(), e.Tag(), e.Value())
	}
	return result
}

// V校验_注册自定义验证 注册自定义的验证函数。
// 注册后可在标签中使用自定义的验证规则名。
//
// 参数:
//   - 规则名: 自定义验证规则的名称，如 "phone"
//   - 验证函数: 验证函数，接收字段值返回是否通过
//
// 返回:
//   - error: 注册失败时返回错误
func V校验_注册自定义验证(规则名 string, 验证函数 func(fl validator.FieldLevel) bool) error {
	return 全局校验器.RegisterValidation(规则名, 验证函数)
}

// V校验_是邮箱 验证字符串是否为有效的邮箱地址格式。
//
// 参数:
//   - 邮箱: 待验证的邮箱字符串
//
// 返回:
//   - bool: true 表示是有效的邮箱格式
func V校验_是邮箱(邮箱 string) bool {
	err := 全局校验器.Var(邮箱, "required,email")
	return err == nil
}

// V校验_是URL 验证字符串是否为有效的 URL 格式。
//
// 参数:
//   - 网址: 待验证的 URL 字符串
//
// 返回:
//   - bool: true 表示是有效的 URL 格式
func V校验_是URL(网址 string) bool {
	err := 全局校验器.Var(网址, "required,url")
	return err == nil
}

// V校验_是IP 验证字符串是否为有效的 IP 地址（支持 IPv4 和 IPv6）。
//
// 参数:
//   - ip: 待验证的 IP 地址字符串
//
// 返回:
//   - bool: true 表示是有效的 IP 地址
func V校验_是IP(ip string) bool {
	err := 全局校验器.Var(ip, "required,ip")
	return err == nil
}

// V校验_是JSON 验证字符串是否为有效的 JSON 格式。
//
// 参数:
//   - json文本: 待验证的 JSON 字符串
//
// 返回:
//   - bool: true 表示是有效的 JSON 格式
func V校验_是JSON(json文本 string) bool {
	err := 全局校验器.Var(json文本, "required,json")
	return err == nil
}

// V校验_取结构体标签 获取结构体字段的 validate 标签值。
// 用于调试和日志，查看字段绑定的校验规则。
//
// 参数:
//   - 结构体: 结构体实例
//   - 字段名: 字段名称
//
// 返回:
//   - string: validate 标签的值；字段不存在或无标签时返回空串
func V校验_取结构体标签(结构体 interface{}, 字段名 string) string {
	t := reflect.TypeOf(结构体)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	field, ok := t.FieldByName(字段名)
	if !ok {
		return ""
	}
	return field.Tag.Get("validate")
}
