package utils

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// D数据库_连接MySQL 连接 MySQL 数据库并返回引擎实例。
// 使用 xorm 作为 ORM 引擎，支持结构体映射和链式查询。
// 驱动自动注册 github.com/go-sql-driver/mysql。
//
// 参数:
//   - 用户名: 数据库用户名
//   - 密码: 数据库密码
//   - 主机地址: 数据库地址，如 "127.0.0.1:3306"
//   - 数据库名: 要连接的数据库名称
//
// 返回:
//   - *xorm.Engine: xorm 引擎实例
//   - error: 连接失败时返回错误
func D数据库_连接MySQL(用户名 string, 密码 string, 主机地址 string, 数据库名 string) (*xorm.Engine, error) {
	连接字符串 := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		用户名, 密码, 主机地址, 数据库名)
	return xorm.NewEngine("mysql", 连接字符串)
}

// D数据库_连接SQLite 连接 SQLite 数据库并返回引擎实例。
// 文件不存在时自动创建。
//
// 参数:
//   - 文件路径: SQLite 数据库文件路径
//
// 返回:
//   - *xorm.Engine: xorm 引擎实例
//   - error: 连接失败时返回错误
func D数据库_连接SQLite(文件路径 string) (*xorm.Engine, error) {
	return xorm.NewEngine("sqlite3", 文件路径)
}

// D数据库_测试连接 测试数据库连接是否正常。
// 执行一次 Ping 操作验证连接可用性。
//
// 参数:
//   - 引擎: xorm 引擎实例
//
// 返回:
//   - error: 连接失败时返回错误
func D数据库_测试连接(引擎 *xorm.Engine) error {
	return 引擎.Ping()
}

// D数据库_同步表 自动同步结构体到数据库表。
// 根据结构体定义创建或更新数据库表结构。
// 会自动创建不存在的表，添加不存在的列，但不会删除已有的列。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 结构体: 要同步的结构体实例（可传入多个）
//
// 返回:
//   - error: 同步失败时返回错误
func D数据库_同步表(引擎 *xorm.Engine, 结构体 ...interface{}) error {
	return 引擎.Sync2(结构体...)
}

// D数据库_插入 插入一条或多条记录到数据库。
// 传入结构体指针或结构体指针切片。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 记录: 要插入的数据，结构体指针或切片
//
// 返回:
//   - int64: 影响的行数
//   - error: 插入失败时返回错误
func D数据库_插入(引擎 *xorm.Engine, 记录 interface{}) (int64, error) {
	return 引擎.Insert(记录)
}

// D数据库_查询 查询记录并存入结构体切片。
// 传入结构体切片指针，查询结果会填充到切片中。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 结果切片: 存放查询结果的切片指针，如 &[]User{}
//
// 返回:
//   - error: 查询失败时返回错误
func D数据库_查询(引擎 *xorm.Engine, 结果切片 interface{}) error {
	return 引擎.Find(结果切片)
}

// D数据库_查询单条 查询单条记录。
// 传入结构体指针，查询结果会填充到结构体中。
// 如果查询到多条记录，只返回第一条。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 记录: 存放查询结果的结构体指针
//
// 返回:
//   - bool: true 表示查到记录，false 表示未查到
//   - error: 查询失败时返回错误
func D数据库_查询单条(引擎 *xorm.Engine, 记录 interface{}) (bool, error) {
	return 引擎.Get(记录)
}

// D数据库_更新 更新记录。
// 传入要更新的结构体指针，非零值字段会被更新。
// 建议配合 Where 条件使用，避免全表更新。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 记录: 包含更新值的结构体指针
//
// 返回:
//   - int64: 影响的行数
//   - error: 更新失败时返回错误
func D数据库_更新(引擎 *xorm.Engine, 记录 interface{}) (int64, error) {
	return 引擎.Update(记录)
}

// D数据库_条件更新 带条件的更新记录。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 记录: 包含更新值的结构体指针
//   - 条件: WHERE 条件字符串
//   - 参数: 条件参数值
//
// 返回:
//   - int64: 影响的行数
//   - error: 更新失败时返回错误
func D数据库_条件更新(引擎 *xorm.Engine, 记录 interface{}, 条件 string, 参数 ...interface{}) (int64, error) {
	return 引擎.Where(条件, 参数...).Update(记录)
}

// D数据库_删除 删除记录。
// 传入结构体作为删除条件，非零值字段作为 WHERE 条件。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 条件结构体: 作为删除条件的结构体
//
// 返回:
//   - int64: 影响的行数
//   - error: 删除失败时返回错误
func D数据库_删除(引擎 *xorm.Engine, 条件结构体 interface{}) (int64, error) {
	return 引擎.Delete(条件结构体)
}

// D数据库_条件查询 带条件查询记录。
// 使用 SQL 条件字符串和参数进行查询。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 结果切片: 存放查询结果的切片指针
//   - 条件: SQL WHERE 条件，如 "age > ? AND name LIKE ?"
//   - 参数: 条件中的参数值
//
// 返回:
//   - error: 查询失败时返回错误
func D数据库_条件查询(引擎 *xorm.Engine, 结果切片 interface{}, 条件 string, 参数 ...interface{}) error {
	return 引擎.Where(条件, 参数...).Find(结果切片)
}

// D数据库_统计 获取满足条件的记录数。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 条件结构体: 作为统计条件的结构体
//
// 返回:
//   - int64: 记录数
//   - error: 查询失败时返回错误
func D数据库_统计(引擎 *xorm.Engine, 条件结构体 interface{}) (int64, error) {
	return 引擎.Count(条件结构体)
}

// D数据库_执行SQL 执行原生 SQL 语句。
// 适用于复杂的 SQL 操作，如批量更新、多表关联等。
// 参数以 sqlOrArgs 形式传入，第一个参数为 SQL 语句，后续为参数值。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - sql语句: SQL 语句
//   - 参数: SQL 中的参数值
//
// 返回:
//   - sql.Result: SQL 执行结果
//   - error: 执行失败时返回错误
func D数据库_执行SQL(引擎 *xorm.Engine, sql语句 string, 参数 ...interface{}) (sql.Result, error) {
	args := make([]interface{}, 0, len(参数)+1)
	args = append(args, sql语句)
	args = append(args, 参数...)
	return 引擎.Exec(args...)
}

// D数据库_事务 执行数据库事务。
// 事务中的所有操作要么全部成功，要么全部回滚。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 事务函数: 事务中执行的操作，接收 Session 对象
//
// 返回:
//   - interface{}: 事务返回值
//   - error: 事务失败时返回错误（自动回滚）
func D数据库_事务(引擎 *xorm.Engine, 事务函数 func(session *xorm.Session) (interface{}, error)) (interface{}, error) {
	return 引擎.Transaction(事务函数)
}

// D数据库_设置连接池 设置数据库连接池参数。
// 合理配置连接池可提高并发性能。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 最大空闲数: 连接池最大空闲连接数
//   - 最大连接数: 连接池最大连接数，0 表示不限制
//   - 最大生存时间: 连接最大生存时间（秒），0 表示不限制
func D数据库_设置连接池(引擎 *xorm.Engine, 最大空闲数 int, 最大连接数 int, 最大生存时间 int) {
	引擎.SetMaxIdleConns(最大空闲数)
	引擎.SetMaxOpenConns(最大连接数)
	if 最大生存时间 > 0 {
		引擎.SetConnMaxLifetime(time.Duration(最大生存时间) * time.Second)
	}
}

// D数据库_关闭 关闭数据库连接，释放资源。
//
// 参数:
//   - 引擎: xorm 引擎实例
//
// 返回:
//   - error: 关闭失败时返回错误
func D数据库_关闭(引擎 *xorm.Engine) error {
	return 引擎.Close()
}

// D数据库_取表名 获取结构体对应的数据库表名。
// 需要结构体定义了 `xorm:"'table_name'"` 标签或遵循 xorm 命名规则。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 结构体: 结构体实例
//
// 返回:
//   - string: 表名
func D数据库_取表名(引擎 *xorm.Engine, 结构体 interface{}) string {
	return 引擎.TableName(结构体, true)
}

// D数据库_表是否存在 检查数据库中是否存在指定表。
//
// 参数:
//   - 引擎: xorm 引擎实例
//   - 表名: 表名
//
// 返回:
//   - bool: true 表示表存在
//   - error: 查询失败时返回错误
func D数据库_表是否存在(引擎 *xorm.Engine, 表名 string) (bool, error) {
	return 引擎.IsTableExist(表名)
}

// D数据库_取反射类型 辅助函数：获取 interface{} 的 reflect.Type。
// 用于 xorm 的泛型查询接口。
func D数据库_取反射类型(值 interface{}) reflect.Type {
	return reflect.TypeOf(值)
}
