// 配置管理工具
// 基于 spf13/viper 库，支持 JSON/TOML/YAML/INI/properties 等格式的配置文件读写。
// 支持配置热更新、环境变量绑定、配置写入等功能。
package utils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config_LoadFile 从指定配置文件中读取配置项。
// 支持 JSON/TOML/YAML/INI/properties 等格式，自动根据文件扩展名识别。
//
// 参数:
//   - filePath: 配置文件的完整路径
//
// 返回:
//   - *viper.Viper: viper 实例，可用于后续取值操作
//   - error: 读取失败时返回错误
func Config_LoadFile(filePath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

// Config_LoadKey 从配置文件中读取指定键的值，返回字符串。
//
// 参数:
//   - filePath: 配置文件的完整路径
//   - key: 配置项的键名，支持点号分隔的层级路径，如 "database.host"
//
// 返回:
//   - string: 配置项的值
//   - error: 读取失败时返回错误
func Config_LoadKey(filePath string, key string) (string, error) {
	v, err := Config_LoadFile(filePath)
	if err != nil {
		return "", err
	}
	return v.GetString(key), nil
}

// Config_GetString 从 viper 实例中获取字符串值。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//
// 返回:
//   - string: 配置项的值
func Config_GetString(v *viper.Viper, key string) string {
	return v.GetString(key)
}

// Config_GetInt 从 viper 实例中获取整数值。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//
// 返回:
//   - int: 配置项的整数值
func Config_GetInt(v *viper.Viper, key string) int {
	return v.GetInt(key)
}

// Config_GetFloat64 从 viper 实例中获取浮点数值。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//
// 返回:
//   - float64: 配置项的浮点数值
func Config_GetFloat64(v *viper.Viper, key string) float64 {
	return v.GetFloat64(key)
}

// Config_GetBool 从 viper 实例中获取布尔值。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//
// 返回:
//   - bool: 配置项的布尔值
func Config_GetBool(v *viper.Viper, key string) bool {
	return v.GetBool(key)
}

// Config_GetStringSlice 从 viper 实例中获取字符串切片。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//
// 返回:
//   - []string: 字符串切片
func Config_GetStringSlice(v *viper.Viper, key string) []string {
	return v.GetStringSlice(key)
}

// Config_GetIntSlice 从 viper 实例中获取整数切片。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//
// 返回:
//   - []int: 整数切片
func Config_GetIntSlice(v *viper.Viper, key string) []int {
	return v.GetIntSlice(key)
}

// Config_Set 在 viper 实例中设置指定键的值。
// 设置的值不会自动写入文件，需要调用 Config_Save 持久化。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//   - value: 要设置的值
func Config_Set(v *viper.Viper, key string, value interface{}) {
	v.Set(key, value)
}

// Config_Save 将 viper 实例中的配置写回文件。
//
// 参数:
//   - v: viper 实例
//
// 返回:
//   - error: 写入失败时返回错误
func Config_Save(v *viper.Viper) error {
	return v.WriteConfig()
}

// Config_Watch 监听配置文件的变更，当文件被修改时自动重新加载。
// 需配合 Config_OnChange 使用处理变更事件。
//
// 参数:
//   - v: viper 实例
func Config_Watch(v *viper.Viper) {
	v.WatchConfig()
}

// Config_OnChange 注册配置文件变更时的回调函数。
// 当监听到配置文件被修改时，自动调用回调函数。
//
// 参数:
//   - v: viper 实例
//   - fn: 变更时执行的回调函数
func Config_OnChange(v *viper.Viper, fn func()) {
	v.OnConfigChange(func(e fsnotify.Event) {
		fn()
	})
}

// Config_BindEnv 将指定键绑定到环境变量。
// 读取配置时，若文件中无该键，则从环境变量中获取。
//
// 参数:
//   - v: viper 实例
//   - key: 配置项的键名
//   - envName: 对应的环境变量名
func Config_BindEnv(v *viper.Viper, key string, envName string) {
	v.BindEnv(key, envName)
}

// Config_AutoEnv 开启自动环境变量绑定。
// 设置前缀后，所有键名会自动映射为 "前缀_键名" 格式的环境变量。
//
// 参数:
//   - v: viper 实例
//   - prefix: 环境变量前缀，如 "MYAPP"
func Config_AutoEnv(v *viper.Viper, prefix string) {
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
}