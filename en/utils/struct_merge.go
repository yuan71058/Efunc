// 结构体合并/复制工具
// 基于 dario.cat/mergo 和 jinzhu/copier 库，提供结构体合并与复制功能。
// 适用于配置合并、默认值填充、结构体克隆等场景。
package utils

import (
	"dario.cat/mergo"
	"github.com/jinzhu/copier"
)

// Struct_Merge 将源结构体的非零值字段合并到目标结构体。
// 仅覆盖目标结构体中的零值字段，非零值保持不变。
// 基于 mergo 库实现，适用于配置合并、默认值填充等场景。
//
// 参数:
//   - dst: 目标结构体指针，合并结果写入此结构体
//   - src: 源结构体，其非零值将覆盖目标中的零值
//
// 返回:
//   - error: 合并失败时返回错误（如类型不匹配）
func Struct_Merge(dst interface{}, src interface{}) error {
	return mergo.Merge(dst, src)
}

// Struct_MergeOverride 将源结构体的所有字段合并到目标结构体。
// 与 Struct_Merge 不同，此函数会覆盖目标中的所有字段（包括非零值）。
//
// 参数:
//   - dst: 目标结构体指针
//   - src: 源结构体
//
// 返回:
//   - error: 合并失败时返回错误
func Struct_MergeOverride(dst interface{}, src interface{}) error {
	return mergo.Merge(dst, src, mergo.WithOverride)
}

// Struct_MergeWithoutDereference 将源结构体的非零值字段合并到目标，跳过空值。
// 源结构体中为零值的字段不会覆盖目标中的对应字段。
//
// 参数:
//   - dst: 目标结构体指针
//   - src: 源结构体
//
// 返回:
//   - error: 合并失败时返回错误
func Struct_MergeWithoutDereference(dst interface{}, src interface{}) error {
	return mergo.Merge(dst, src, mergo.WithoutDereference)
}

// Struct_MergeMap 将源 map 的非零值合并到目标 map。
//
// 参数:
//   - dst: 目标 map 指针
//   - src: 源 map
//
// 返回:
//   - error: 合并失败时返回错误
func Struct_MergeMap(dst interface{}, src interface{}) error {
	return mergo.Map(dst, src)
}

// Struct_MergeMapOverride 将源 map 的所有值合并到目标 map，覆盖已有键。
//
// 参数:
//   - dst: 目标 map 指针
//   - src: 源 map
//
// 返回:
//   - error: 合并失败时返回错误
func Struct_MergeMapOverride(dst interface{}, src interface{}) error {
	return mergo.Map(dst, src, mergo.WithOverride)
}

// Struct_Copy 将源结构体的字段值复制到目标结构体。
// 基于 copier 库实现，支持不同结构体之间的字段复制（按字段名匹配）。
// 支持切片、嵌套结构体等复杂类型的复制。
//
// 参数:
//   - dst: 目标结构体指针
//   - src: 源结构体实例
//
// 返回:
//   - error: 复制失败时返回错误
func Struct_Copy(dst interface{}, src interface{}) error {
	return copier.Copy(dst, src)
}

// Struct_CopyWithOption 带选项复制结构体字段。
// 可控制是否忽略空值、是否深拷贝等行为。
//
// 参数:
//   - dst: 目标结构体指针
//   - src: 源结构体实例
//   - ignoreEmpty: true 时源中的零值字段不会覆盖目标
//
// 返回:
//   - error: 复制失败时返回错误
func Struct_CopyWithOption(dst interface{}, src interface{}, ignoreEmpty bool) error {
	if ignoreEmpty {
		return copier.CopyWithOption(dst, src, copier.Option{IgnoreEmpty: true})
	}
	return copier.Copy(dst, src)
}