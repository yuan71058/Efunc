// 数据验证工具
// 基于 go-playground/validator/v10，支持结构体字段验证、单变量验证、自定义规则注册。
// 常用标签: required(必填), email(邮箱), min/max(最小/最大值), len(长度), oneof(枚举值)。
package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var globalValidator = validator.New()

// Validation_Struct 验证结构体字段的标签约束。
// 结构体字段需使用 validate 标签定义校验规则，如 `validate:"required,email"`。
//
// 参数:
//   - s: 待验证的结构体实例（非指针）
//
// 返回:
//   - error: 验证失败时返回详细的字段错误信息；通过时返回 nil
func Validation_Struct(s interface{}) error {
	return globalValidator.Struct(s)
}

// Validation_Var 验证单个变量的值是否符合标签约束。
// 适用于不需要定义结构体的简单校验场景。
//
// 参数:
//   - field: 待验证的变量值
//   - tag: 校验规则标签，如 "required,email"、"min=1,max=100"
//
// 返回:
//   - error: 验证失败时返回错误信息；通过时返回 nil
func Validation_Var(field interface{}, tag string) error {
	return globalValidator.Var(field, tag)
}

// Validation_VarWithValue 验证结构体中指定字段的值。
// 仅验证单个字段，不验证整个结构体。
//
// 参数:
//   - s: 包含待验证字段的结构体
//   - fieldName: 要验证的字段名（结构体中的字段名，非 JSON 名）
//   - tag: 校验规则标签
//
// 返回:
//   - error: 验证失败时返回错误信息
func Validation_VarWithValue(s interface{}, fieldName string, tag string) error {
	return globalValidator.VarWithValue(s, fieldName, tag)
}

// Validation_GetError 从验证错误中提取可读的错误信息。
// 将 validator.ValidationErrors 转换为中文友好的错误描述。
//
// 参数:
//   - err: Validation_Struct 返回的错误
//
// 返回:
//   - string: 格式化的错误信息，如 "字段 Email 验证 required 失败"
func Validation_GetError(err error) string {
	if err == nil {
		return ""
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	result := ""
	for _, e := range errs {
		result += fmt.Sprintf("字段 %s 验证 %s 失败, 当前值: %v; ", e.Field(), e.Tag(), e.Value())
	}
	return result
}

// Validation_Register 注册自定义的验证函数。
// 注册后可在标签中使用自定义的验证规则名。
//
// 参数:
//   - ruleName: 自定义验证规则的名称，如 "phone"
//   - fn: 验证函数，接收字段值返回是否通过
//
// 返回:
//   - error: 注册失败时返回错误
func Validation_Register(ruleName string, fn func(fl validator.FieldLevel) bool) error {
	return globalValidator.RegisterValidation(ruleName, fn)
}

// Validation_IsEmail 验证字符串是否为有效的邮箱地址格式。
//
// 参数:
//   - email: 待验证的邮箱字符串
//
// 返回:
//   - bool: true 表示是有效的邮箱格式
func Validation_IsEmail(email string) bool {
	err := globalValidator.Var(email, "required,email")
	return err == nil
}

// Validation_IsURL 验证字符串是否为有效的 URL 格式。
//
// 参数:
//   - url: 待验证的 URL 字符串
//
// 返回:
//   - bool: true 表示是有效的 URL 格式
func Validation_IsURL(url string) bool {
	err := globalValidator.Var(url, "required,url")
	return err == nil
}

// Validation_IsIP 验证字符串是否为有效的 IP 地址（支持 IPv4 和 IPv6）。
//
// 参数:
//   - ip: 待验证的 IP 地址字符串
//
// 返回:
//   - bool: true 表示是有效的 IP 地址
func Validation_IsIP(ip string) bool {
	err := globalValidator.Var(ip, "required,ip")
	return err == nil
}

// Validation_IsJSON 验证字符串是否为有效的 JSON 格式。
//
// 参数:
//   - jsonText: 待验证的 JSON 字符串
//
// 返回:
//   - bool: true 表示是有效的 JSON 格式
func Validation_IsJSON(jsonText string) bool {
	err := globalValidator.Var(jsonText, "required,json")
	return err == nil
}

// Validation_GetStructTag 获取结构体字段的 validate 标签值。
// 用于调试和日志，查看字段绑定的校验规则。
//
// 参数:
//   - s: 结构体实例
//   - fieldName: 字段名称
//
// 返回:
//   - string: validate 标签的值；字段不存在或无标签时返回空串
func Validation_GetStructTag(s interface{}, fieldName string) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	field, ok := t.FieldByName(fieldName)
	if !ok {
		return ""
	}
	return field.Tag.Get("validate")
}