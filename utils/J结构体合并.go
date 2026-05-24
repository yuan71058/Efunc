package utils

import (
	"dario.cat/mergo"
	"github.com/jinzhu/copier"
)

// J结构体_合并 将源结构体的非零值字段合并到目标结构体。
// 仅覆盖目标结构体中的零值字段，非零值保持不变。
// 基于 mergo 库实现，适用于配置合并、默认值填充等场景。
//
// 参数:
//   - 目标: 目标结构体指针，合并结果写入此结构体
//   - 源: 源结构体，其非零值将覆盖目标中的零值
//
// 返回:
//   - error: 合并失败时返回错误（如类型不匹配）
func J结构体_合并(目标 interface{}, 源 interface{}) error {
	return mergo.Merge(目标, 源)
}

// J结构体_合并覆盖 将源结构体的所有字段合并到目标结构体。
// 与 J结构体_合并 不同，此函数会覆盖目标中的所有字段（包括非零值）。
//
// 参数:
//   - 目标: 目标结构体指针
//   - 源: 源结构体
//
// 返回:
//   - error: 合并失败时返回错误
func J结构体_合并覆盖(目标 interface{}, 源 interface{}) error {
	return mergo.Merge(目标, 源, mergo.WithOverride)
}

// J结构体_合并不带空值 将源结构体的非零值字段合并到目标，跳过空值。
// 源结构体中为零值的字段不会覆盖目标中的对应字段。
//
// 参数:
//   - 目标: 目标结构体指针
//   - 源: 源结构体
//
// 返回:
//   - error: 合并失败时返回错误
func J结构体_合并不带空值(目标 interface{}, 源 interface{}) error {
	return mergo.Merge(目标, 源, mergo.WithoutDereference)
}

// J结构体_合并Map 将源 map 的非零值合并到目标 map。
//
// 参数:
//   - 目标: 目标 map 指针
//   - 源: 源 map
//
// 返回:
//   - error: 合并失败时返回错误
func J结构体_合并Map(目标 interface{}, 源 interface{}) error {
	return mergo.Map(目标, 源)
}

// J结构体_合并Map覆盖 将源 map 的所有值合并到目标 map，覆盖已有键。
//
// 参数:
//   - 目标: 目标 map 指针
//   - 源: 源 map
//
// 返回:
//   - error: 合并失败时返回错误
func J结构体_合并Map覆盖(目标 interface{}, 源 interface{}) error {
	return mergo.Map(目标, 源, mergo.WithOverride)
}

// J结构体_复制 将源结构体的字段值复制到目标结构体。
// 基于 copier 库实现，支持不同结构体之间的字段复制（按字段名匹配）。
// 支持切片、嵌套结构体等复杂类型的复制。
//
// 参数:
//   - 目标: 目标结构体指针
//   - 源: 源结构体实例
//
// 返回:
//   - error: 复制失败时返回错误
func J结构体_复制(目标 interface{}, 源 interface{}) error {
	return copier.Copy(目标, 源)
}

// J结构体_复制带选项 带选项复制结构体字段。
// 可控制是否忽略空值、是否深拷贝等行为。
//
// 参数:
//   - 目标: 目标结构体指针
//   - 源: 源结构体实例
//   - 忽略空值: true 时源中的零值字段不会覆盖目标
//
// 返回:
//   - error: 复制失败时返回错误
func J结构体_复制带选项(目标 interface{}, 源 interface{}, 忽略空值 bool) error {
	if 忽略空值 {
		return copier.CopyWithOption(目标, 源, copier.Option{IgnoreEmpty: true})
	}
	return copier.Copy(目标, 源)
}
