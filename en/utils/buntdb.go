// BuntDB 键值库封装
// 基于 tidwall/buntdb 库，提供基于内存的键值数据库操作。
// 支持 JSON 索引、事务操作、过期时间设置和范围遍历等功能。
// 数据会持久化到文件，文件不存在时自动创建。
package utils

import (
	"time"

	"github.com/tidwall/buntdb"
)

// BuntDB_Open 打开或创建一个 BuntDB 数据库文件。
//
// 参数:
//   - filePath: 数据库文件路径，如 "data.db"；空字符串则使用纯内存模式
//
// 返回:
//   - *buntdb.DB: 数据库实例
//   - error: 打开失败时返回错误
func BuntDB_Open(filePath string) (*buntdb.DB, error) {
	return buntdb.Open(filePath)
}

// BuntDB_Get 从数据库中获取指定键的值。
//
// 参数:
//   - db: 数据库实例
//   - key: 要获取的键名
//
// 返回:
//   - string: 键对应的值；键不存在时返回空串
//   - error: 键不存在或读取失败时返回错误
func BuntDB_Get(db *buntdb.DB, key string) (string, error) {
	var val string
	err := db.View(func(tx *buntdb.Tx) error {
		var err error
		val, err = tx.Get(key)
		return err
	})
	return val, err
}

// BuntDB_Set 设置指定键的值。键不存在则创建，已存在则覆盖。
func BuntDB_Set(db *buntdb.DB, key string, value string) error {
	return db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

// BuntDB_SetWithTTL 设置指定键的值，并设置过期时间。过期后键值自动删除。
func BuntDB_SetWithTTL(db *buntdb.DB, key string, value string, expireSeconds float64) error {
	return db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, &buntdb.SetOptions{
			Expires: true,
			TTL:     time.Duration(expireSeconds * float64(time.Second)),
		})
		return err
	})
}

// BuntDB_Delete 删除指定键。
func BuntDB_Delete(db *buntdb.DB, key string) error {
	return db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(key)
		return err
	})
}

// BuntDB_Scan 遍历指定范围内的所有键值对。
func BuntDB_Scan(db *buntdb.DB, startKey string, endKey string, fn func(key, value string) bool) error {
	return db.View(func(tx *buntdb.Tx) error {
		return tx.Ascend("", func(key, value string) bool {
			return fn(key, value)
		})
	})
}

// BuntDB_CreateIndex 为数据库创建 JSON 字段索引。
func BuntDB_CreateIndex(db *buntdb.DB, indexName string, pattern string) error {
	return db.CreateIndex(indexName, "*", buntdb.IndexJSON(pattern))
}

// BuntDB_Close 关闭数据库，释放资源。关闭前会自动将内存数据持久化到文件。
func BuntDB_Close(db *buntdb.DB) error {
	return db.Close()
}