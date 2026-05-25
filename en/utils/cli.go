// 命令行参数工具
// 基于 Go 标准库 flag 包，提供命令行参数解析功能。
// 支持字符串、整数、布尔、浮点数等类型的参数定义与解析。
package utils

import (
	"flag"
	"os"
)

// CLI_Parse 解析命令行参数。
// 调用 flag.Parse() 解析 os.Args[1:] 中的命令行参数。
// 应在定义所有命令行参数后调用此函数。
func CLI_Parse() {
	flag.Parse()
}

// CLI_GetString 定义并获取一个字符串类型的命令行参数。
// 支持短参数名和长参数名两种形式，如 -n 和 -name。
//
// 参数:
//   - shortName: 短参数名，如 "n"
//   - longName: 长参数名，如 "name"（保留参数位，供文档使用）
//   - defaultValue: 参数未提供时的默认值
//   - usage: 参数的用法说明
//
// 返回:
//   - *string: 指向参数值的指针，解析后通过 *指针 获取实际值
func CLI_GetString(shortName string, longName string, defaultValue string, usage string) *string {
	return flag.String(shortName, defaultValue, usage)
}

// CLI_GetInt 定义并获取一个整数类型的命令行参数。
//
// 参数:
//   - shortName: 短参数名，如 "p"
//   - longName: 长参数名，如 "port"（保留参数位，供文档使用）
//   - defaultValue: 参数未提供时的默认值
//   - usage: 参数的用法说明
//
// 返回:
//   - *int: 指向参数值的指针
func CLI_GetInt(shortName string, longName string, defaultValue int, usage string) *int {
	return flag.Int(shortName, defaultValue, usage)
}

// CLI_GetBool 定义并获取一个布尔类型的命令行参数。
//
// 参数:
//   - shortName: 短参数名，如 "v"
//   - longName: 长参数名，如 "verbose"（保留参数位，供文档使用）
//   - defaultValue: 参数未提供时的默认值
//   - usage: 参数的用法说明
//
// 返回:
//   - *bool: 指向参数值的指针
func CLI_GetBool(shortName string, longName string, defaultValue bool, usage string) *bool {
	return flag.Bool(shortName, defaultValue, usage)
}

// CLI_GetFloat64 定义并获取一个浮点数类型的命令行参数。
//
// 参数:
//   - shortName: 短参数名，如 "r"
//   - longName: 长参数名，如 "rate"（保留参数位，供文档使用）
//   - defaultValue: 参数未提供时的默认值
//   - usage: 参数的用法说明
//
// 返回:
//   - *float64: 指向参数值的指针
func CLI_GetFloat64(shortName string, longName string, defaultValue float64, usage string) *float64 {
	return flag.Float64(shortName, defaultValue, usage)
}

// CLI_GetArgs 获取命令行中未被 flag 解析的位置参数。
// 必须在 CLI_Parse 之后调用。
//
// 返回:
//   - []string: 未被解析的位置参数列表
func CLI_GetArgs() []string {
	return flag.Args()
}

// CLI_GetProgramName 获取当前运行的程序名称（os.Args[0]）。
//
// 返回:
//   - string: 程序名称或路径
func CLI_GetProgramName() string {
	return os.Args[0]
}

// CLI_SetUsage 设置命令行参数的用法说明。
// 当用户输入 -h 或 --help 时显示此说明。
//
// 参数:
//   - fn: 自定义的用法说明函数
func CLI_SetUsage(fn func()) {
	flag.Usage = fn
}

// CLI_ShowUsage 显示默认的命令行参数用法说明。
func CLI_ShowUsage() {
	flag.Usage()
}