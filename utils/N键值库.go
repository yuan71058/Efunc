package utils

import (
	"time"

	"github.com/tidwall/buntdb"
)

// N键值_打开 打开或创建一个 BuntDB 数据库文件。
// BuntDB 是基于内存的键值数据库，支持 JSON 操作、事务和索引。
// 数据会持久化到指定文件，文件不存在时自动创建。
//
// 参数:
//   - 文件路径: 数据库文件路径，如 "data.db"；空字符串则使用纯内存模式
//
// 返回:
//   - *buntdb.DB: 数据库实例
//   - error: 打开失败时返回错误
func N键值_打开(文件路径 string) (*buntdb.DB, error) {
	return buntdb.Open(文件路径)
}

// N键值_取值 从数据库中获取指定键的值。
//
// 参数:
//   - 数据库: 数据库实例
//   - 键: 要获取的键名
//
// 返回:
//   - string: 键对应的值；键不存在时返回空串
//   - error: 键不存在或读取失败时返回错误
func N键值_取值(数据库 *buntdb.DB, 键 string) (string, error) {
	var 值 string
	err := 数据库.View(func(tx *buntdb.Tx) error {
		var err error
		值, err = tx.Get(键)
		return err
	})
	return 值, err
}

// N键值_置值 设置指定键的值。键不存在则创建，已存在则覆盖。
//
// 参数:
//   - 数据库: 数据库实例
//   - 键: 键名
//   - 值: 键值
//
// 返回:
//   - error: 设置失败时返回错误
func N键值_置值(数据库 *buntdb.DB, 键 string, 值 string) error {
	return 数据库.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(键, 值, nil)
		return err
	})
}

// N键值_置值带过期 设置指定键的值，并设置过期时间。
// 过期后键值自动删除。
//
// 参数:
//   - 数据库: 数据库实例
//   - 键: 键名
//   - 值: 键值
//   - 过期秒数: 过期时间（秒），0 表示永不过期
//
// 返回:
//   - error: 设置失败时返回错误
func N键值_置值带过期(数据库 *buntdb.DB, 键 string, 值 string, 过期秒数 float64) error {
	return 数据库.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(键, 值, &buntdb.SetOptions{Expires: true, TTL: time.Duration(过期秒数 * float64(time.Second))})
		return err
	})
}

// N键值_删除 删除指定键。
//
// 参数:
//   - 数据库: 数据库实例
//   - 键: 要删除的键名
//
// 返回:
//   - error: 删除失败时返回错误
func N键值_删除(数据库 *buntdb.DB, 键 string) error {
	return 数据库.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(键)
		return err
	})
}

// N键值_遍历 遍历指定范围内的所有键值对。
// 起始和结束为空字符串时遍历所有键。
//
// 参数:
//   - 数据库: 数据库实例
//   - 起始键: 遍历起始键（包含），空字符串从头开始
//   - 结束键: 遍历结束键（不包含），空字符串到末尾
//   - 回调函数: 每个键值对的回调，返回 false 停止遍历
//
// 返回:
//   - error: 遍历失败时返回错误
func N键值_遍历(数据库 *buntdb.DB, 起始键 string, 结束键 string, 回调函数 func(键, 值 string) bool) error {
	return 数据库.View(func(tx *buntdb.Tx) error {
		return tx.Ascend("", func(键, 值 string) bool {
			return 回调函数(键, 值)
		})
	})
}

// N键值_创建索引 为数据库创建索引，支持按 JSON 字段索引。
// 创建索引后可使用索引名进行高效查询。
//
// 参数:
//   - 数据库: 数据库实例
//   - 索引名: 索引名称
//   - 模式: 索引模式，如 "name" 或 "user.name"（JSON 路径）
//
// 返回:
//   - error: 创建失败时返回错误
func N键值_创建索引(数据库 *buntdb.DB, 索引名 string, 模式 string) error {
	return 数据库.CreateIndex(索引名, "*", buntdb.IndexJSON(模式))
}

// N键值_关闭 关闭数据库，释放资源。
// 关闭前会自动将内存数据持久化到文件。
//
// 参数:
//   - 数据库: 数据库实例
//
// 返回:
//   - error: 关闭失败时返回错误
func N键值_关闭(数据库 *buntdb.DB) error {
	return 数据库.Close()
}
