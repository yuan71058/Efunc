package utils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// P配置_从文件读取 从指定配置文件中读取配置项。
// 支持 JSON/TOML/YAML/INI/properties 等格式，自动根据文件扩展名识别。
//
// 参数:
//   - 文件路径: 配置文件的完整路径
//
// 返回:
//   - *viper.Viper: viper 实例，可用于后续取值操作
//   - error: 读取失败时返回错误
func P配置_从文件读取(文件路径 string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(文件路径)
	if 错误 := v.ReadInConfig(); 错误 != nil {
		return nil, 错误
	}
	return v, nil
}

// P配置_从文件读取指定项 从配置文件中读取指定键的值，返回字符串。
//
// 参数:
//   - 文件路径: 配置文件的完整路径
//   - 键名: 配置项的键名，支持点号分隔的层级路径，如 "database.host"
//
// 返回:
//   - string: 配置项的值
//   - error: 读取失败时返回错误
func P配置_从文件读取指定项(文件路径 string, 键名 string) (string, error) {
	v, 错误 := P配置_从文件读取(文件路径)
	if 错误 != nil {
		return "", 错误
	}
	return v.GetString(键名), nil
}

// P配置_取值 从 viper 实例中获取字符串值。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//
// 返回:
//   - string: 配置项的值
func P配置_取值(v *viper.Viper, 键名 string) string {
	return v.GetString(键名)
}

// P配置_取整数值 从 viper 实例中获取整数值。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//
// 返回:
//   - int: 配置项的整数值
func P配置_取整数值(v *viper.Viper, 键名 string) int {
	return v.GetInt(键名)
}

// P配置_取浮点数值 从 viper 实例中获取浮点数值。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//
// 返回:
//   - float64: 配置项的浮点数值
func P配置_取浮点数值(v *viper.Viper, 键名 string) float64 {
	return v.GetFloat64(键名)
}

// P配置_取逻辑值 从 viper 实例中获取布尔值。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//
// 返回:
//   - bool: 配置项的布尔值
func P配置_取逻辑值(v *viper.Viper, 键名 string) bool {
	return v.GetBool(键名)
}

// P配置_取字符串切片 从 viper 实例中获取字符串切片。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//
// 返回:
//   - []string: 字符串切片
func P配置_取字符串切片(v *viper.Viper, 键名 string) []string {
	return v.GetStringSlice(键名)
}

// P配置_取整数切片 从 viper 实例中获取整数切片。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//
// 返回:
//   - []int: 整数切片
func P配置_取整数切片(v *viper.Viper, 键名 string) []int {
	return v.GetIntSlice(键名)
}

// P配置_设置值 在 viper 实例中设置指定键的值。
// 设置的值不会自动写入文件，需要调用 P配置_写回文件 持久化。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//   - 值: 要设置的值
func P配置_设置值(v *viper.Viper, 键名 string, 值 interface{}) {
	v.Set(键名, 值)
}

// P配置_写回文件 将 viper 实例中的配置写回文件。
//
// 参数:
//   - v: viper 实例
//
// 返回:
//   - error: 写入失败时返回错误
func P配置_写回文件(v *viper.Viper) error {
	return v.WriteConfig()
}

// P配置_监听变更 监听配置文件的变更，当文件被修改时自动重新加载。
// 需配合 P配置_变更回调 使用处理变更事件。
//
// 参数:
//   - v: viper 实例
func P配置_监听变更(v *viper.Viper) {
	v.WatchConfig()
}

// P配置_变更回调 注册配置文件变更时的回调函数。
// 当监听到配置文件被修改时，自动调用回调函数。
//
// 参数:
//   - v: viper 实例
//   - 回调: 变更时执行的回调函数
func P配置_变更回调(v *viper.Viper, 回调 func()) {
	v.OnConfigChange(func(e fsnotify.Event) {
		回调()
	})
}

// P配置_绑定环境变量 将指定键绑定到环境变量。
// 读取配置时，若文件中无该键，则从环境变量中获取。
//
// 参数:
//   - v: viper 实例
//   - 键名: 配置项的键名
//   - 环境变量名: 对应的环境变量名
func P配置_绑定环境变量(v *viper.Viper, 键名 string, 环境变量名 string) {
	v.BindEnv(键名, 环境变量名)
}

// P配置_自动环境变量 开启自动环境变量绑定。
// 设置前缀后，所有键名会自动映射为 "前缀_键名" 格式的环境变量。
//
// 参数:
//   - v: viper 实例
//   - 前缀: 环境变量前缀，如 "MYAPP"
func P配置_自动环境变量(v *viper.Viper, 前缀 string) {
	v.SetEnvPrefix(前缀)
	v.AutomaticEnv()
}
