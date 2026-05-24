package utils

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

// B表达式_计算 计算一个数学或逻辑表达式字符串并返回结果。
// 支持基本四则运算、比较运算、逻辑运算和内置函数。
// 表达式示例: "1 + 2 * 3"、"10 > 5 && 3 < 7"、"pow(2, 10)"。
//
// 参数:
//   - 表达式: 表达式字符串
//
// 返回:
//   - interface{}: 计算结果，类型取决于表达式（int64/float64/bool 等）
//   - error: 表达式语法错误或计算错误时返回
func B表达式_计算(表达式 string) (interface{}, error) {
	expr, err := govaluate.NewEvaluableExpression(表达式)
	if err != nil {
		return nil, err
	}
	return expr.Evaluate(nil)
}

// B表达式_计算带参数 计算带参数的表达式。
// 参数以 map 形式传入，表达式中用变量名引用。
// 示例: 表达式 "a + b * c"，参数 {"a": 1, "b": 2, "c": 3}，结果为 7。
//
// 参数:
//   - 表达式: 包含变量的表达式字符串
//   - 参数: 变量名到值的映射，key 为变量名（string），value 为变量值
//
// 返回:
//   - interface{}: 计算结果
//   - error: 表达式语法错误、变量未定义或计算错误时返回
func B表达式_计算带参数(表达式 string, 参数 map[string]interface{}) (interface{}, error) {
	expr, err := govaluate.NewEvaluableExpression(表达式)
	if err != nil {
		return nil, err
	}
	return expr.Evaluate(参数)
}

// B表达式_编译 预编译表达式，返回可重复使用的表达式对象。
// 适用于需要多次计算同一表达式的场景，避免重复解析。
//
// 参数:
//   - 表达式: 表达式字符串
//
// 返回:
//   - *govaluate.EvaluableExpression: 编译后的表达式对象
//   - error: 表达式语法错误时返回
func B表达式_编译(表达式 string) (*govaluate.EvaluableExpression, error) {
	return govaluate.NewEvaluableExpression(表达式)
}

// B表达式_编译并计算 编译表达式后立即计算（无参数）。
// 如果需要多次计算同一表达式，建议使用 B表达式_编译 获取对象后重复使用。
//
// 参数:
//   - 表达式: 表达式字符串
//
// 返回:
//   - interface{}: 计算结果
//   - error: 编译或计算错误时返回
func B表达式_编译并计算(表达式 string) (interface{}, error) {
	expr, err := B表达式_编译(表达式)
	if err != nil {
		return nil, fmt.Errorf("表达式编译失败: %w", err)
	}
	return expr.Evaluate(nil)
}

// B表达式_求值 对已编译的表达式对象进行求值（无参数）。
//
// 参数:
//   - 表达式对象: 通过 B表达式_编译 获得的表达式对象
//
// 返回:
//   - interface{}: 计算结果
//   - error: 计算错误时返回
func B表达式_求值(表达式对象 *govaluate.EvaluableExpression) (interface{}, error) {
	return 表达式对象.Evaluate(nil)
}

// B表达式_求值带参数 对已编译的表达式对象进行求值（带参数）。
//
// 参数:
//   - 表达式对象: 通过 B表达式_编译 获得的表达式对象
//   - 参数: 变量名到值的映射
//
// 返回:
//   - interface{}: 计算结果
//   - error: 计算错误时返回
func B表达式_求值带参数(表达式对象 *govaluate.EvaluableExpression, 参数 map[string]interface{}) (interface{}, error) {
	return 表达式对象.Evaluate(参数)
}
