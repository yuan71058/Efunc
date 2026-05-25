// 环境变量工具
// 基于 godotenv 库，支持 .env 文件加载、环境变量设置/读取/删除等操作。
// 环境变量在当前进程及其子进程中有效。
package utils

import (
	"github.com/joho/godotenv"
	"os"
)

// Env_Load 从指定文件加载环境变量到当前进程。
// 不会覆盖已存在的环境变量。
//
// 参数:
//   - filePath: .env 文件路径，如 ".env" 或 "config/.env"
//
// 返回:
//   - error: 加载失败时返回错误
func Env_Load(filePath string) error {
	return godotenv.Load(filePath)
}

// Env_Overload 从指定文件加载环境变量到当前进程。
// 会覆盖已存在的同名环境变量。
//
// 参数:
//   - filePath: .env 文件路径
//
// 返回:
//   - error: 加载失败时返回错误
func Env_Overload(filePath string) error {
	return godotenv.Overload(filePath)
}

// Env_Get 获取指定环境变量的值。
//
// 参数:
//   - name: 环境变量名
//
// 返回:
//   - string: 环境变量的值；变量不存在时返回空串
func Env_Get(name string) string {
	return os.Getenv(name)
}

// Env_GetWithDefault 获取指定环境变量的值，不存在时返回默认值。
//
// 参数:
//   - name: 环境变量名
//   - defaultValue: 变量不存在时返回的默认值
//
// 返回:
//   - string: 环境变量的值或默认值
func Env_GetWithDefault(name string, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

// Env_Set 设置环境变量的值。
// 设置的环境变量仅在当前进程及其子进程中有效。
//
// 参数:
//   - name: 环境变量名
//   - value: 环境变量的值
//
// 返回:
//   - error: 设置失败时返回错误
func Env_Set(name string, value string) error {
	return os.Setenv(name, value)
}

// Env_Unset 删除指定的环境变量。
//
// 参数:
//   - name: 环境变量名
//
// 返回:
//   - error: 删除失败时返回错误
func Env_Unset(name string) error {
	return os.Unsetenv(name)
}

// Env_Exists 判断指定环境变量是否存在。
//
// 参数:
//   - name: 环境变量名
//
// 返回:
//   - bool: 存在返回 true，否则返回 false
func Env_Exists(name string) bool {
	_, exists := os.LookupEnv(name)
	return exists
}

// Env_GetAll 获取所有环境变量。
//
// 返回:
//   - []string: 所有环境变量，格式为 "KEY=VALUE"
func Env_GetAll() []string {
	return os.Environ()
}

// Env_Unmarshal 从文本内容解析环境变量（无需文件）。
// 文本格式与 .env 文件相同，每行一个 "KEY=VALUE"。
//
// 参数:
//   - text: 环境变量文本内容
//
// 返回:
//   - map[string]string: 解析后的键值对
//   - error: 解析失败时返回错误
func Env_Unmarshal(text string) (map[string]string, error) {
	return godotenv.Unmarshal(text)
}