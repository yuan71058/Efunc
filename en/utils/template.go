// 快速模板替换工具
// 基于 valyala/fasttemplate 库，提供高性能模板字符串替换。
// 模板中使用 {占位符} 格式的占位符，通过标签映射替换为实际值。
// 相比 strings.Replace 和 fmt.Sprintf，fasttemplate 在高频调用时性能更优。
package utils

import (
	"io"

	"github.com/valyala/fasttemplate"
)

// Template_Execute 使用 fasttemplate 执行模板替换。
// 模板中使用 {占位符} 格式的占位符，通过标签映射替换为实际值。
//
// 参数:
//   - templateText: 包含占位符的模板字符串，占位符格式为 {key}
//   - tags: 占位符名到替换值的映射，如 {"name": "张三", "age": "25"}
//
// 返回:
//   - string: 替换后的结果字符串
//   - error: 模板语法错误时返回
func Template_Execute(templateText string, tags map[string]interface{}) (string, error) {
	t, err := fasttemplate.NewTemplate(templateText, "{", "}")
	if err != nil {
		return "", err
	}
	return t.ExecuteString(tags), nil
}

// Template_ExecuteDelim 使用自定义分隔符执行模板替换。
// 可自定义占位符的起始和结束标记，如 {{ 和 }} 或 <% 和 %>。
//
// 参数:
//   - templateText: 包含占位符的模板字符串
//   - startTag: 占位符的起始标记，如 "{{"
//   - endTag: 占位符的结束标记，如 "}}"
//   - tags: 占位符名到替换值的映射
//
// 返回:
//   - string: 替换后的结果字符串
//   - error: 模板语法错误时返回
func Template_ExecuteDelim(templateText string, startTag string, endTag string, tags map[string]interface{}) (string, error) {
	t, err := fasttemplate.NewTemplate(templateText, startTag, endTag)
	if err != nil {
		return "", err
	}
	return t.ExecuteString(tags), nil
}

// Template_Write 将模板替换结果写入指定的 io.Writer。
// 适用于输出到文件、网络连接等场景，避免生成中间字符串。
//
// 参数:
//   - templateText: 包含占位符的模板字符串
//   - w: 实现了 io.Writer 接口的对象
//   - tags: 占位符名到替换值的映射
//
// 返回:
//   - int64: 写入的字节数
//   - error: 模板语法错误或写入失败时返回
func Template_Write(templateText string, w io.Writer, tags map[string]interface{}) (int64, error) {
	t, err := fasttemplate.NewTemplate(templateText, "{", "}")
	if err != nil {
		return 0, err
	}
	n, err := t.Execute(w, tags)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Template_New 创建可重复使用的模板对象。
// 适用于需要多次执行同一模板的场景，避免重复解析模板。
//
// 参数:
//   - templateText: 包含占位符的模板字符串
//   - startTag: 占位符的起始标记
//   - endTag: 占位符的结束标记
//
// 返回:
//   - *fasttemplate.Template: 模板对象，可重复调用 ExecuteString
//   - error: 模板语法错误时返回
func Template_New(templateText string, startTag string, endTag string) (*fasttemplate.Template, error) {
	return fasttemplate.NewTemplate(templateText, startTag, endTag)
}

// Template_Replace 简单的模板替换，使用默认 {key} 格式。
// 是 Template_Execute 的简化版本，忽略错误直接返回结果。
//
// 参数:
//   - templateText: 包含占位符的模板字符串
//   - tags: 占位符名到替换值的映射
//
// 返回:
//   - string: 替换后的结果字符串；出错时返回原模板文本
func Template_Replace(templateText string, tags map[string]interface{}) string {
	result, err := Template_Execute(templateText, tags)
	if err != nil {
		return templateText
	}
	return result
}