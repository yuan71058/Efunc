package utils

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
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

	buf.WriteString("| ")
	buf.WriteString(strings.Join(表头, " | "))
	buf.WriteString(" |\n")

	buf.WriteString("| ")
	for i := range 表头 {
		if i > 0 {
			buf.WriteString(" | ")
		}
		buf.WriteString("---")
	}
	buf.WriteString(" |\n")

	for _, 行 := range 行数据 {
		buf.WriteString("| ")
		buf.WriteString(strings.Join(行, " | "))
		buf.WriteString(" |\n")
	}
	return buf.String()
}

// K表格_输出CSV 将表格输出为 CSV 格式。
// 每行用逗号分隔，字段含逗号时自动加引号。
//
// 参数:
//   - 表头: 表头列名列表
//   - 行数据: 二维字符串数组
//
// 返回:
//   - string: CSV 格式字符串
func K表格_输出CSV(表头 []string, 行数据 [][]string) string {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	writer.Write(表头)
	writer.WriteAll(行数据)
	writer.Flush()

	return buf.String()
}

// K表格_输出TSV 将表格输出为 TSV 格式（制表符分隔）。
//
// 参数:
//   - 表头: 表头列名列表
//   - 行数据: 二维字符串数组
//
// 返回:
//   - string: TSV 格式字符串
func K表格_输出TSV(表头 []string, 行数据 [][]string) string {
	var buf bytes.Buffer

	buf.WriteString(strings.Join(表头, "\t"))
	buf.WriteString("\n")

	for _, 行 := range 行数据 {
		buf.WriteString(strings.Join(行, "\t"))
		buf.WriteString("\n")
	}
	return buf.String()
}

// K表格_输出JSON 将表格输出为 JSON 数组格式。
// 每行数据与表头组合成对象，所有行组成数组。
//
// 参数:
//   - 表头: 表头列名列表
//   - 行数据: 二维字符串数组
//
// 返回:
//   - string: JSON 格式字符串
func K表格_输出JSON(表头 []string, 行数据 [][]string) string {
	结果 := make([]map[string]string, 0, len(行数据))
	for _, 行 := range 行数据 {
		行map := make(map[string]string)
		for i, 值 := range 行 {
			if i < len(表头) {
				行map[表头[i]] = 值
			}
		}
		结果 = append(结果, 行map)
	}
	数据, _ := json.MarshalIndent(结果, "", "  ")
	return string(数据)
}

// K表格_输出HTML 将表格输出为 HTML 表格格式。
//
// 参数:
//   - 表头: 表头列名列表
//   - 行数据: 二维字符串数组
//
// 返回:
//   - string: HTML 表格字符串
func K表格_输出HTML(表头 []string, 行数据 [][]string) string {
	var buf bytes.Buffer

	buf.WriteString("<table>\n  <thead>\n    <tr>\n")
	for _, 列 := range 表头 {
		buf.WriteString(fmt.Sprintf("      <th>%s</th>\n", 列))
	}
	buf.WriteString("    </tr>\n  </thead>\n  <tbody>\n")

	for _, 行 := range 行数据 {
		buf.WriteString("    <tr>\n")
		for _, 值 := range 行 {
			buf.WriteString(fmt.Sprintf("      <td>%s</td>\n", 值))
		}
		buf.WriteString("    </tr>\n")
	}

	buf.WriteString("  </tbody>\n</table>")
	return buf.String()
}

// K表格_从CSV读取 从 CSV 字符串解析表格数据。
//
// 参数:
//   - csv文本: CSV 格式的文本
//
// 返回:
//   - []string: 表头
//   - [][]string: 行数据
//   - error: 解析失败时返回错误
func K表格_从CSV读取(csv文本 string) ([]string, [][]string, error) {
	reader := csv.NewReader(strings.NewReader(csv文本))
	记录, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	if len(记录) == 0 {
		return []string{}, [][]string{}, nil
	}
	return 记录[0], 记录[1:], nil
}

// K表格_从TSV读取 从 TSV 字符串解析表格数据。
//
// 参数:
//   - tsv文本: TSV 格式的文本
//
// 返回:
//   - []string: 表头
//   - [][]string: 行数据
func K表格_从TSV读取(tsv文本 string) ([]string, [][]string) {
	行列表 := strings.Split(strings.TrimSpace(tsv文本), "\n")
	if len(行列表) == 0 {
		return []string{}, [][]string{}
	}
	表头 := strings.Split(行列表[0], "\t")
	行数据 := make([][]string, 0, len(行列表)-1)
	for _, 行 := range 行列表[1:] {
		行数据 = append(行数据, strings.Split(行, "\t"))
	}
	return 表头, 行数据
}

// K表格_转置 将表格数据进行转置（行列互换）。
//
// 参数:
//   - 表头: 表头列名列表
//   - 行数据: 二维字符串数组
//
// 返回:
//   - []string: 新表头
//   - [][]string: 新行数据
func K表格_转置(表头 []string, 行数据 [][]string) ([]string, [][]string) {
	最大列数 := len(表头)
	for _, 行 := range 行数据 {
		if len(行) > 最大列数 {
			最大列数 = len(行)
		}
	}

	新表头 := make([]string, len(行数据)+1)
	新表头[0] = ""
	for i := range 行数据 {
		if i < len(表头) {
			新表头[i+1] = fmt.Sprintf("行%d", i+1)
		}
	}

	新行数据 := make([][]string, 最大列数)
	for i := 0; i < 最大列数; i++ {
		新行数据[i] = make([]string, len(行数据)+1)
		if i < len(表头) {
			新行数据[i][0] = 表头[i]
		}
		for j, 行 := range 行数据 {
			if i < len(行) {
				新行数据[i][j+1] = 行[i]
			}
		}
	}
	return 新表头, 新行数据
}

// K表格_过滤行 按条件过滤表格行数据。
// 条件函数返回 true 时保留该行。
//
// 参数:
//   - 行数据: 二维字符串数组
//   - 条件: 过滤条件函数，参数为行索引和行数据
//
// 返回:
//   - [][]string: 过滤后的行数据
func K表格_过滤行(行数据 [][]string, 条件 func(int, []string) bool) [][]string {
	结果 := make([][]string, 0)
	for i, 行 := range 行数据 {
		if 条件(i, 行) {
			结果 = append(结果, 行)
		}
	}
	return 结果
}

// K表格_排序列 按指定列排序表格行数据。
//
// 参数:
//   - 行数据: 二维字符串数组
//   - 列索引: 排序依据的列索引
//   - 升序: true 升序，false 降序
//
// 返回:
//   - [][]string: 排序后的行数据
func K表格_排序列(行数据 [][]string, 列索引 int, 升序 bool) [][]string {
	结果 := make([][]string, len(行数据))
	copy(结果, 行数据)

	for i := 0; i < len(结果)-1; i++ {
		for j := i + 1; j < len(结果); j++ {
			需交换 := false
			if 列索引 < len(结果[i]) && 列索引 < len(结果[j]) {
				if 升序 {
					需交换 = 结果[i][列索引] > 结果[j][列索引]
				} else {
					需交换 = 结果[i][列索引] < 结果[j][列索引]
				}
			}
			if 需交换 {
				结果[i], 结果[j] = 结果[j], 结果[i]
			}
		}
	}
	return 结果
}

// K表格_取列 提取表格指定列的所有值。
//
// 参数:
//   - 行数据: 二维字符串数组
//   - 列索引: 要提取的列索引
//
// 返回:
//   - []string: 该列所有值
func K表格_取列(行数据 [][]string, 列索引 int) []string {
	结果 := make([]string, 0, len(行数据))
	for _, 行 := range 行数据 {
		if 列索引 < len(行) {
			结果 = append(结果, 行[列索引])
		}
	}
	return 结果
}

// K表格_取行 提取表格指定行的数据。
//
// 参数:
//   - 行数据: 二维字符串数组
//   - 行索引: 要提取的行索引
//
// 返回:
//   - []string: 该行数据；索引越界返回空切片
func K表格_取行(行数据 [][]string, 行索引 int) []string {
	if 行索引 < 0 || 行索引 >= len(行数据) {
		return []string{}
	}
	return 行数据[行索引]
}

// K表格_行数 获取表格行数。
//
// 参数:
//   - 行数据: 二维字符串数组
//
// 返回:
//   - int: 行数
func K表格_行数(行数据 [][]string) int {
	return len(行数据)
}

// K表格_列数 获取表格列数（基于表头）。
//
// 参数:
//   - 表头: 表头列名列表
//
// 返回:
//   - int: 列数
func K表格_列数(表头 []string) int {
	return len(表头)
}

// K表格_合并 合并两个表格（行追加）。
// 两个表格的列数应一致。
//
// 参数:
//   - 行数据1: 第一个表格的行数据
//   - 行数据2: 第二个表格的行数据
//
// 返回:
//   - [][]string: 合并后的行数据
func K表格_合并(行数据1 [][]string, 行数据2 [][]string) [][]string {
	结果 := make([][]string, 0, len(行数据1)+len(行数据2))
	结果 = append(结果, 行数据1...)
	结果 = append(结果, 行数据2...)
	return 结果
}

// K表格_去重 对表格行数据进行去重。
//
// 参数:
//   - 行数据: 二维字符串数组
//
// 返回:
//   - [][]string: 去重后的行数据
func K表格_去重(行数据 [][]string) [][]string {
	seen := make(map[string]bool)
	结果 := make([][]string, 0)
	for _, 行 := range 行数据 {
		key := strings.Join(行, "\x00")
		if !seen[key] {
			seen[key] = true
			结果 = append(结果, 行)
		}
	}
	return 结果
}
