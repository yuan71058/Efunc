package utils

import (
	"io"

	"github.com/valyala/fasttemplate"
)

// T模板_执行 使用 fasttemplate 执行模板替换。
// 模板中使用 {占位符} 格式的占位符，通过标签映射替换为实际值。
// 相比 strings.Replace 和 fmt.Sprintf，fasttemplate 在高频调用时性能更优。
//
// 参数:
//   - 模板文本: 包含占位符的模板字符串，占位符格式为 {key}
//   - 标签: 占位符名到替换值的映射，如 {"name": "张三", "age": "25"}
//
// 返回:
//   - string: 替换后的结果字符串
//   - error: 模板语法错误时返回
func T模板_执行(模板文本 string, 标签 map[string]interface{}) (string, error) {
	t, err := fasttemplate.NewTemplate(模板文本, "{", "}")
	if err != nil {
		return "", err
	}
	return t.ExecuteString(标签), nil
}

// T模板_执行自定义分隔符 使用自定义分隔符执行模板替换。
// 可自定义占位符的起始和结束标记，如 {{ 和 }} 或 <% 和 %>。
//
// 参数:
//   - 模板文本: 包含占位符的模板字符串
//   - 起始标记: 占位符的起始标记，如 "{{"
//   - 结束标记: 占位符的结束标记，如 "}}"
//   - 标签: 占位符名到替换值的映射
//
// 返回:
//   - string: 替换后的结果字符串
//   - error: 模板语法错误时返回
func T模板_执行自定义分隔符(模板文本 string, 起始标记 string, 结束标记 string, 标签 map[string]interface{}) (string, error) {
	t, err := fasttemplate.NewTemplate(模板文本, 起始标记, 结束标记)
	if err != nil {
		return "", err
	}
	return t.ExecuteString(标签), nil
}

// T模板_写入 将模板替换结果写入指定的 io.Writer。
// 适用于输出到文件、网络连接等场景，避免生成中间字符串。
//
// 参数:
//   - 模板文本: 包含占位符的模板字符串
//   - 写入器: 实现了 io.Writer 接口的对象
//   - 标签: 占位符名到替换值的映射
//
// 返回:
//   - int64: 写入的字节数
//   - error: 模板语法错误或写入失败时返回
func T模板_写入(模板文本 string, 写入器 io.Writer, 标签 map[string]interface{}) (int64, error) {
	t, err := fasttemplate.NewTemplate(模板文本, "{", "}")
	if err != nil {
		return 0, err
	}
	n, err := t.Execute(写入器, 标签)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// T模板_创建 创建可重复使用的模板对象。
// 适用于需要多次执行同一模板的场景，避免重复解析模板。
//
// 参数:
//   - 模板文本: 包含占位符的模板字符串
//   - 起始标记: 占位符的起始标记
//   - 结束标记: 占位符的结束标记
//
// 返回:
//   - *fasttemplate.Template: 模板对象，可重复调用 ExecuteString
//   - error: 模板语法错误时返回
func T模板_创建(模板文本 string, 起始标记 string, 结束标记 string) (*fasttemplate.Template, error) {
	return fasttemplate.NewTemplate(模板文本, 起始标记, 结束标记)
}

// T模板_替换 简单的模板替换，使用默认 {key} 格式。
// 是 T模板_执行 的简化版本，忽略错误直接返回结果。
//
// 参数:
//   - 模板文本: 包含占位符的模板字符串
//   - 标签: 占位符名到替换值的映射
//
// 返回:
//   - string: 替换后的结果字符串；出错时返回原模板文本
func T模板_替换(模板文本 string, 标签 map[string]interface{}) string {
	result, err := T模板_执行(模板文本, 标签)
	if err != nil {
		return 模板文本
	}
	return result
}
