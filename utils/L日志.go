package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// L日志_创建开发日志 创建适用于开发环境的日志实例。
// 输出到控制台，使用人类可读的格式，日志级别为 Debug。
//
// 返回:
//   - *zap.Logger: zap 日志实例
//   - error: 创建失败时返回错误
func L日志_创建开发日志() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

// L日志_创建生产日志 创建适用于生产环境的日志实例。
// 输出 JSON 格式日志，日志级别为 Info，自带调用者信息和堆栈跟踪。
//
// 返回:
//   - *zap.Logger: zap 日志实例
//   - error: 创建失败时返回错误
func L日志_创建生产日志() (*zap.Logger, error) {
	return zap.NewProduction()
}

// L日志_创建自定义日志 创建自定义配置的日志实例。
// 可指定日志级别、输出路径、编码格式等。
//
// 参数:
//   - 日志级别: 0=Debug, 1=Info, 2=Warn, 3=Error
//   - 输出路径: 日志文件路径列表，如 []string{"stdout", "/var/log/app.log"}
//   - JSON格式: true 使用 JSON 编码，false 使用 Console 编码
//
// 返回:
//   - *zap.Logger: zap 日志实例
//   - error: 创建失败时返回错误
func L日志_创建自定义日志(日志级别 int, 输出路径 []string, JSON格式 bool) (*zap.Logger, error) {
	var 级别 zapcore.Level
	switch 日志级别 {
	case 0:
		级别 = zapcore.DebugLevel
	case 1:
		级别 = zapcore.InfoLevel
	case 2:
		级别 = zapcore.WarnLevel
	case 3:
		级别 = zapcore.ErrorLevel
	default:
		级别 = zapcore.InfoLevel
	}

	编码 := "console"
	if JSON格式 {
		编码 = "json"
	}

	配置 := zap.Config{
		Level:            zap.NewAtomicLevelAt(级别),
		Development:      false,
		Encoding:         编码,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      输出路径,
		ErrorOutputPaths: 输出路径,
	}

	return 配置.Build()
}

// L日志_调试 输出 Debug 级别的日志。
//
// 参数:
//   - 日志: zap 日志实例
//   - 消息: 日志消息
//   - 字段: 可选的结构化字段，如 zap.String("key", "value")
func L日志_调试(日志 *zap.Logger, 消息 string, 字段 ...zap.Field) {
	日志.Debug(消息, 字段...)
}

// L日志_信息 输出 Info 级别的日志。
//
// 参数:
//   - 日志: zap 日志实例
//   - 消息: 日志消息
//   - 字段: 可选的结构化字段
func L日志_信息(日志 *zap.Logger, 消息 string, 字段 ...zap.Field) {
	日志.Info(消息, 字段...)
}

// L日志_警告 输出 Warn 级别的日志。
//
// 参数:
//   - 日志: zap 日志实例
//   - 消息: 日志消息
//   - 字段: 可选的结构化字段
func L日志_警告(日志 *zap.Logger, 消息 string, 字段 ...zap.Field) {
	日志.Warn(消息, 字段...)
}

// L日志_错误 输出 Error 级别的日志。
//
// 参数:
//   - 日志: zap 日志实例
//   - 消息: 日志消息
//   - 字段: 可选的结构化字段
func L日志_错误(日志 *zap.Logger, 消息 string, 字段 ...zap.Field) {
	日志.Error(消息, 字段...)
}

// L日志_字符串 创建字符串类型的结构化日志字段。
//
// 参数:
//   - 键: 字段键名
//   - 值: 字段值
//
// 返回:
//   - zap.Field: zap 字段
func L日志_字符串(键 string, 值 string) zap.Field {
	return zap.String(键, 值)
}

// L日志_整数 创建整数类型的结构化日志字段。
//
// 参数:
//   - 键: 字段键名
//   - 值: 字段值
//
// 返回:
//   - zap.Field: zap 字段
func L日志_整数(键 string, 值 int) zap.Field {
	return zap.Int(键, 值)
}

// L日志_错误类型 创建 error 类型的结构化日志字段。
//
// 参数:
//   - 键: 字段键名
//   - 值: 错误值
//
// 返回:
//   - zap.Field: zap 字段
func L日志_错误类型(键 string, 值 error) zap.Field {
	return zap.NamedError(键, 值)
}

// L日志_同步 刷新日志缓冲区，确保所有日志都已写入。
// 程序退出前应调用此函数。
//
// 参数:
//   - 日志: zap 日志实例
//
// 返回:
//   - error: 同步失败时返回错误
func L日志_同步(日志 *zap.Logger) error {
	return 日志.Sync()
}
