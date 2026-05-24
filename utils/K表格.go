package utils

import (
	"bytes"
	"strings"

	"github.com/scylladb/termtables"
)

// K表格_创建 创建一个新的控制台表格。
// 创建后可通过 K表格_添加表头 和 K表格_添加行 填充数据。
//
// 返回:
//   - *termtables.Table: 表格对象指针
func K表格_创建() *termtables.Table {
	return termtables.CreateTable()
}

// K表格_添加表头 为表格添加表头行。
// 表头会以粗体或分隔线形式突出显示。
//
// 参数:
//   - 表格: 表格对象指针
//   - 表头: 表头列名列表
func K表格_添加表头(表格 *termtables.Table, 表头 ...string) {
	表格.AddHeaders(表头)
}

// K表格_添加行 为表格添加一行数据。
// 列数应与表头列数一致。
//
// 参数:
//   - 表格: 表格对象指针
//   - 行数据: 该行各列的值
func K表格_添加行(表格 *termtables.Table, 行数据 ...interface{}) {
	表格.AddRow(行数据)
}

// K表格_添加分隔线 在表格中添加一条水平分隔线。
//
// 参数:
//   - 表格: 表格对象指针
func K表格_添加分隔线(表格 *termtables.Table) {
	表格.AddSeparator()
}

// K表格_输出 将表格渲染为字符串输出。
// 自动对齐列宽，适合在控制台显示。
//
// 参数:
//   - 表格: 表格对象指针
//
// 返回:
//   - string: 渲染后的表格字符串
func K表格_输出(表格 *termtables.Table) string {
	return 表格.Render()
}

// K表格_快速创建 快速创建并输出一个表格。
// 传入表头和行数据，直接返回渲染后的表格字符串。
//
// 参数:
//   - 表头: 表头列名列表
//   - 行数据: 二维字符串数组，每行一个切片
//
// 返回:
//   - string: 渲染后的表格字符串
func K表格_快速创建(表头 []string, 行数据 [][]string) string {
	表格 := termtables.CreateTable()
	表格.AddHeaders(表头)
	for _, 行 := range 行数据 {
		行interface := make([]interface{}, len(行))
		for i, v := range 行 {
			行interface[i] = v
		}
		表格.AddRow(行interface...)
	}
	return 表格.Render()
}

// K表格_输出Markdown 将表格渲染为 Markdown 格式。
// 使用 | 分隔列，适合写入 Markdown 文档。
//
// 参数:
//   - 表头: 表头列名列表
//   - 行数据: 二维字符串数组
//
// 返回:
//   - string: Markdown 格式的表格字符串
func K表格_输出Markdown(表头 []string, 行数据 [][]string) string {
	var buf bytes.Buffer
	分隔线 := strings.Repeat("---|", len(表头))

	buf.WriteString("| ")
	buf.WriteString(strings.Join(表头, " | "))
	buf.WriteString(" |\n")
	buf.WriteString("| ")
	buf.WriteString(分隔线)
	buf.WriteString("\n")

	for _, 行 := range 行数据 {
		buf.WriteString("| ")
		buf.WriteString(strings.Join(行, " | "))
		buf.WriteString(" |\n")
	}
	return buf.String()
}
