// 控制台表格输出工具
// 基于 scylladb/termtables 库，支持表格创建、渲染，以及导出为 Markdown/CSV/TSV/JSON/HTML 格式。
package utils

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/scylladb/termtables"
)

func Table_New() *termtables.Table { return termtables.CreateTable() }
func Table_AddHeaders(t *termtables.Table, headers ...string) { t.AddHeaders(headers) }
func Table_AddRow(t *termtables.Table, row ...interface{}) { t.AddRow(row) }
func Table_AddSeparator(t *termtables.Table) { t.AddSeparator() }
func Table_Render(t *termtables.Table) string { return t.Render() }

// Table_Quick 快速创建并输出一个表格。
func Table_Quick(headers []string, rows [][]string) string {
	t := termtables.CreateTable()
	t.AddHeaders(headers)
	for _, row := range rows {
		rowInterface := make([]interface{}, len(row))
		for i, v := range row {
			rowInterface[i] = v
		}
		t.AddRow(rowInterface...)
	}
	return t.Render()
}

// Table_ToMarkdown 将表格渲染为 Markdown 格式。
func Table_ToMarkdown(headers []string, rows [][]string) string {
	var buf bytes.Buffer
	buf.WriteString("| ")
	buf.WriteString(strings.Join(headers, " | "))
	buf.WriteString(" |\n")
	buf.WriteString("| ")
	for i := range headers {
		if i > 0 {
			buf.WriteString(" | ")
		}
		buf.WriteString("---")
	}
	buf.WriteString(" |\n")
	for _, row := range rows {
		buf.WriteString("| ")
		buf.WriteString(strings.Join(row, " | "))
		buf.WriteString(" |\n")
	}
	return buf.String()
}

// Table_ToCSV 将表格输出为 CSV 格式。
func Table_ToCSV(headers []string, rows [][]string) string {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	writer.Write(headers)
	writer.WriteAll(rows)
	writer.Flush()
	return buf.String()
}

// Table_ToTSV 将表格输出为 TSV 格式（制表符分隔）。
func Table_ToTSV(headers []string, rows [][]string) string {
	var buf bytes.Buffer
	buf.WriteString(strings.Join(headers, "\t"))
	buf.WriteString("\n")
	for _, row := range rows {
		buf.WriteString(strings.Join(row, "\t"))
		buf.WriteString("\n")
	}
	return buf.String()
}

// Table_ToJSON 将表格输出为 JSON 数组格式。
func Table_ToJSON(headers []string, rows [][]string) string {
	result := make([]map[string]string, 0, len(rows))
	for _, row := range rows {
		rowMap := make(map[string]string)
		for i, val := range row {
			if i < len(headers) {
				rowMap[headers[i]] = val
			}
		}
		result = append(result, rowMap)
	}
	data, _ := json.MarshalIndent(result, "", "  ")
	return string(data)
}

// Table_ToHTML 将表格输出为 HTML 表格格式。
func Table_ToHTML(headers []string, rows [][]string) string {
	var buf bytes.Buffer
	buf.WriteString("<table>\n  <thead>\n    <tr>\n")
	for _, col := range headers {
		buf.WriteString(fmt.Sprintf("      <th>%s</th>\n", col))
	}
	buf.WriteString("    </tr>\n  </thead>\n  <tbody>\n")
	for _, row := range rows {
		buf.WriteString("    <tr>\n")
		for _, col := range row {
			buf.WriteString(fmt.Sprintf("      <td>%s</td>\n", col))
		}
		buf.WriteString("    </tr>\n")
	}
	buf.WriteString("  </tbody>\n</table>\n")
	return buf.String()
}