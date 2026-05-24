package utils

import (
	"github.com/joho/godotenv"
	"os"
)

// K环境_加载 从指定文件加载环境变量到当前进程。
// 不会覆盖已存在的环境变量。
//
// 参数:
//   - 文件路径: .env 文件路径，如 ".env" 或 "config/.env"
//
// 返回:
//   - error: 加载失败时返回错误
func K环境_加载(文件路径 string) error {
	return godotenv.Load(文件路径)
}

// K环境_加载并覆盖 从指定文件加载环境变量到当前进程。
// 会覆盖已存在的同名环境变量。
//
// 参数:
//   - 文件路径: .env 文件路径
//
// 返回:
//   - error: 加载失败时返回错误
func K环境_加载并覆盖(文件路径 string) error {
	return godotenv.Overload(文件路径)
}

// K环境_取值 获取指定环境变量的值。
//
// 参数:
//   - 名称: 环境变量名
//
// 返回:
//   - string: 环境变量的值；变量不存在时返回空串
func K环境_取值(名称 string) string {
	return os.Getenv(名称)
}

// K环境_取值带默认值 获取指定环境变量的值，不存在时返回默认值。
//
// 参数:
//   - 名称: 环境变量名
//   - 默认值: 变量不存在时返回的默认值
//
// 返回:
//   - string: 环境变量的值或默认值
func K环境_取值带默认值(名称 string, 默认值 string) string {
	值 := os.Getenv(名称)
	if 值 == "" {
		return 默认值
	}
	return 值
}

// K环境_设置值 设置环境变量的值。
// 设置的环境变量仅在当前进程及其子进程中有效。
//
// 参数:
//   - 名称: 环境变量名
//   - 值: 环境变量的值
//
// 返回:
//   - error: 设置失败时返回错误
func K环境_设置值(名称 string, 值 string) error {
	return os.Setenv(名称, 值)
}

// K环境_删除值 删除指定的环境变量。
//
// 参数:
//   - 名称: 环境变量名
//
// 返回:
//   - error: 删除失败时返回错误
func K环境_删除值(名称 string) error {
	return os.Unsetenv(名称)
}

// K环境_是否存在 判断指定环境变量是否存在。
//
// 参数:
//   - 名称: 环境变量名
//
// 返回:
//   - bool: 存在返回 true，否则返回 false
func K环境_是否存在(名称 string) bool {
	_, 存在 := os.LookupEnv(名称)
	return 存在
}

// K环境_取所有 获取所有环境变量。
//
// 返回:
//   - []string: 所有环境变量，格式为 "KEY=VALUE"
func K环境_取所有() []string {
	return os.Environ()
}

// K环境_从文本读取 从文本内容解析环境变量（无需文件）。
// 文本格式与 .env 文件相同，每行一个 "KEY=VALUE"。
//
// 参数:
//   - 文本: 环境变量文本内容
//
// 返回:
//   - map[string]string: 解析后的键值对
//   - error: 解析失败时返回错误
func K环境_从文本读取(文本 string) (map[string]string, error) {
	return godotenv.Unmarshal(文本)
}
