// 数据库操作模块
// 提供 MySQL/SQLite 数据库连接、表结构同步、增删改查、事务处理、连接池配置等功能。
// 基于 xorm 作为 ORM 引擎，支持结构体映射和链式查询。
package utils

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// Database_ConnectMySQL 连接 MySQL 数据库并返回引擎实例。
// 使用 xorm 作为 ORM 引擎，支持结构体映射和链式查询。
// 驱动自动注册 github.com/go-sql-driver/mysql。
//
// 参数:
//   - user: 数据库用户名
//   - password: 数据库密码
//   - host: 数据库地址，如 "127.0.0.1:3306"
//   - dbName: 要连接的数据库名称
//
// 返回:
//   - *xorm.Engine: xorm 引擎实例
//   - error: 连接失败时返回错误
func Database_ConnectMySQL(user string, password string, host string, dbName string) (*xorm.Engine, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user, password, host, dbName)
	return xorm.NewEngine("mysql", connStr)
}

// Database_ConnectSQLite 连接 SQLite 数据库并返回引擎实例。
// 文件不存在时自动创建。
//
// 参数:
//   - filePath: SQLite 数据库文件路径
//
// 返回:
//   - *xorm.Engine: xorm 引擎实例
//   - error: 连接失败时返回错误
func Database_ConnectSQLite(filePath string) (*xorm.Engine, error) {
	return xorm.NewEngine("sqlite3", filePath)
}

// Database_Ping 测试数据库连接是否正常。
// 执行一次 Ping 操作验证连接可用性。
//
// 参数:
//   - engine: xorm 引擎实例
//
// 返回:
//   - error: 连接失败时返回错误
func Database_Ping(engine *xorm.Engine) error {
	return engine.Ping()
}

// Database_SyncTables 自动同步结构体到数据库表。
// 根据结构体定义创建或更新数据库表结构。
// 会自动创建不存在的表，添加不存在的列，但不会删除已有的列。
//
// 参数:
//   - engine: xorm 引擎实例
//   - beans: 要同步的结构体实例（可传入多个）
//
// 返回:
//   - error: 同步失败时返回错误
func Database_SyncTables(engine *xorm.Engine, beans ...interface{}) error {
	return engine.Sync2(beans...)
}

// Database_Insert 插入一条或多条记录到数据库。
// 传入结构体指针或结构体指针切片。
//
// 参数:
//   - engine: xorm 引擎实例
//   - record: 要插入的数据，结构体指针或切片
//
// 返回:
//   - int64: 影响的行数
//   - error: 插入失败时返回错误
func Database_Insert(engine *xorm.Engine, record interface{}) (int64, error) {
	return engine.Insert(record)
}

// Database_Find 查询记录并存入结构体切片。
// 传入结构体切片指针，查询结果会填充到切片中。
//
// 参数:
//   - engine: xorm 引擎实例
//   - resultSlice: 存放查询结果的切片指针，如 &[]User{}
//
// 返回:
//   - error: 查询失败时返回错误
func Database_Find(engine *xorm.Engine, resultSlice interface{}) error {
	return engine.Find(resultSlice)
}

// Database_Get 查询单条记录。
// 传入结构体指针，查询结果会填充到结构体中。
// 如果查询到多条记录，只返回第一条。
//
// 参数:
//   - engine: xorm 引擎实例
//   - record: 存放查询结果的结构体指针
//
// 返回:
//   - bool: true 表示查到记录，false 表示未查到
//   - error: 查询失败时返回错误
func Database_Get(engine *xorm.Engine, record interface{}) (bool, error) {
	return engine.Get(record)
}

// Database_Update 更新记录。
// 传入要更新的结构体指针，非零值字段会被更新。
// 建议配合 Where 条件使用，避免全表更新。
//
// 参数:
//   - engine: xorm 引擎实例
//   - record: 包含更新值的结构体指针
//
// 返回:
//   - int64: 影响的行数
//   - error: 更新失败时返回错误
func Database_Update(engine *xorm.Engine, record interface{}) (int64, error) {
	return engine.Update(record)
}

// Database_UpdateWhere 带条件的更新记录。
//
// 参数:
//   - engine: xorm 引擎实例
//   - record: 包含更新值的结构体指针
//   - cond: WHERE 条件字符串
//   - args: 条件参数值
//
// 返回:
//   - int64: 影响的行数
//   - error: 更新失败时返回错误
func Database_UpdateWhere(engine *xorm.Engine, record interface{}, cond string, args ...interface{}) (int64, error) {
	return engine.Where(cond, args...).Update(record)
}

// Database_Delete 删除记录。
// 传入结构体作为删除条件，非零值字段作为 WHERE 条件。
//
// 参数:
//   - engine: xorm 引擎实例
//   - condBean: 作为删除条件的结构体
//
// 返回:
//   - int64: 影响的行数
//   - error: 删除失败时返回错误
func Database_Delete(engine *xorm.Engine, condBean interface{}) (int64, error) {
	return engine.Delete(condBean)
}

// Database_FindWhere 带条件查询记录。
// 使用 SQL 条件字符串和参数进行查询。
//
// 参数:
//   - engine: xorm 引擎实例
//   - resultSlice: 存放查询结果的切片指针
//   - cond: SQL WHERE 条件，如 "age > ? AND name LIKE ?"
//   - args: 条件中的参数值
//
// 返回:
//   - error: 查询失败时返回错误
func Database_FindWhere(engine *xorm.Engine, resultSlice interface{}, cond string, args ...interface{}) error {
	return engine.Where(cond, args...).Find(resultSlice)
}

// Database_Count 获取满足条件的记录数。
//
// 参数:
//   - engine: xorm 引擎实例
//   - condBean: 作为统计条件的结构体
//
// 返回:
//   - int64: 记录数
//   - error: 查询失败时返回错误
func Database_Count(engine *xorm.Engine, condBean interface{}) (int64, error) {
	return engine.Count(condBean)
}

// Database_ExecSQL 执行原生 SQL 语句。
// 适用于复杂的 SQL 操作，如批量更新、多表关联等。
// 参数以 sqlOrArgs 形式传入，第一个参数为 SQL 语句，后续为参数值。
//
// 参数:
//   - engine: xorm 引擎实例
//   - sql: SQL 语句
//   - args: SQL 中的参数值
//
// 返回:
//   - sql.Result: SQL 执行结果
//   - error: 执行失败时返回错误
func Database_ExecSQL(engine *xorm.Engine, sql string, args ...interface{}) (sql.Result, error) {
	allArgs := make([]interface{}, 0, len(args)+1)
	allArgs = append(allArgs, sql)
	allArgs = append(allArgs, args...)
	return engine.Exec(allArgs...)
}

// Database_Transaction 执行数据库事务。
// 事务中的所有操作要么全部成功，要么全部回滚。
//
// 参数:
//   - engine: xorm 引擎实例
//   - fn: 事务中执行的操作，接收 Session 对象
//
// 返回:
//   - interface{}: 事务返回值
//   - error: 事务失败时返回错误（自动回滚）
func Database_Transaction(engine *xorm.Engine, fn func(session *xorm.Session) (interface{}, error)) (interface{}, error) {
	return engine.Transaction(fn)
}

// Database_SetPool 设置数据库连接池参数。
// 合理配置连接池可提高并发性能。
//
// 参数:
//   - engine: xorm 引擎实例
//   - maxIdle: 连接池最大空闲连接数
//   - maxOpen: 连接池最大连接数，0 表示不限制
//   - maxLifetime: 连接最大生存时间（秒），0 表示不限制
func Database_SetPool(engine *xorm.Engine, maxIdle int, maxOpen int, maxLifetime int) {
	engine.SetMaxIdleConns(maxIdle)
	engine.SetMaxOpenConns(maxOpen)
	if maxLifetime > 0 {
		engine.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	}
}

// Database_Close 关闭数据库连接，释放资源。
//
// 参数:
//   - engine: xorm 引擎实例
//
// 返回:
//   - error: 关闭失败时返回错误
func Database_Close(engine *xorm.Engine) error {
	return engine.Close()
}

// Database_GetTableName 获取结构体对应的数据库表名。
// 需要结构体定义了 `xorm:"'table_name'"` 标签或遵循 xorm 命名规则。
//
// 参数:
//   - engine: xorm 引擎实例
//   - bean: 结构体实例
//
// 返回:
//   - string: 表名
func Database_GetTableName(engine *xorm.Engine, bean interface{}) string {
	return engine.TableName(bean, true)
}

// Database_IsTableExist 检查数据库中是否存在指定表。
//
// 参数:
//   - engine: xorm 引擎实例
//   - tableName: 表名
//
// 返回:
//   - bool: true 表示表存在
//   - error: 查询失败时返回错误
func Database_IsTableExist(engine *xorm.Engine, tableName string) (bool, error) {
	return engine.IsTableExist(tableName)
}

// Database_GetReflectType 辅助函数：获取 interface{} 的 reflect.Type。
// 用于 xorm 的泛型查询接口。
func Database_GetReflectType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}