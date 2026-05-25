// Package utils 提供中文命名的通用工具函数库，涵盖类型转换、编码处理、文本操作、文件管理、
// 网络请求、加密校验、并发安全等常用功能，是 Go 语言版的精易模块。
package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	mrand "math/rand"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// D到字节集 将任意类型的值转换为字节切片（[]byte）。
// 内部使用 gconv.Bytes 进行转换，支持字符串、整数、浮点数等基本类型。
//
// 参数:
//   - value: 待转换的值，支持 string/int/float/bool 等基本类型
//
// 返回:
//   - []byte: 转换后的字节切片
func D到字节集(value interface{}) []byte {
	return gconv.Bytes(value)
}

// D到字节 将任意类型的值转换为单个字节（byte）。
// 如果值超出 byte 范围（0-255），会进行截断处理。
//
// 参数:
//   - value: 待转换的值
//
// 返回:
//   - byte: 转换后的字节值
func D到字节(value interface{}) byte {
	return gconv.Byte(value)
}

// D到整数 将任意类型的值转换为 int 类型。
// 字符串 "123" → 123，浮点数 3.14 → 3，布尔值 true → 1。
//
// 参数:
//   - value: 待转换的值
//
// 返回:
//   - int: 转换后的整数值
func D到整数(value interface{}) int {
	return gconv.Int(value)
}

// D到整数64 将任意类型的值转换为 int64 类型。
// 适用于需要大整数范围的场景。
//
// 参数:
//   - value: 待转换的值
//
// 返回:
//   - int64: 转换后的 64 位整数值
func D到整数64(value interface{}) int64 {
	return gconv.Int64(value)
}

// D到数值 将任意类型的值转换为 float64 类型。
// 字符串 "3.14" → 3.14，整数 42 → 42.0。
//
// 参数:
//   - value: 待转换的值
//
// 返回:
//   - float64: 转换后的浮点数值
func D到数值(value interface{}) float64 {
	return gconv.Float64(value)
}

// D到文本 将任意类型的值转换为字符串。
// 整数 123 → "123"，浮点数 3.14 → "3.14"，布尔值 true → "true"。
//
// 参数:
//   - value: 待转换的值
//
// 返回:
//   - string: 转换后的字符串
func D到文本(value interface{}) string {
	return gconv.String(value)
}

// D到结构体 将通用类型（如 Map、JSON 字符串）转换为指定的结构体。
// 内部使用 gconv.Struct，支持 map、JSON 字符串等作为输入源。
// 结构体字段可通过 tag（如 json/c/gconv）指定映射名称。
//
// 参数:
//   - 待转换的参数: 源数据，支持 map、JSON 字符串等
//   - 结构体指针: 目标结构体的指针，转换结果将写入该结构体
//
// 返回:
//   - error: 转换失败时返回错误信息
func D到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error {
	return gconv.Struct(待转换的参数, 结构体指针)
}

// S三元 泛型三元运算符，根据条件返回两个值中的一个。
// 类似其他语言的 condition ? a : b 语法。
//
// 类型参数:
//   - T: 任意类型
//
// 参数:
//   - value: 条件表达式，true 时返回 string1，false 时返回 string2
//   - string1: 条件为真时的返回值
//   - string2: 条件为假时的返回值
//
// 返回:
//   - T: 根据条件选中的值
//
// 示例:
//
//	S三元(true, "是", "否")  // "是"
//	S三元(false, 1, 2)       // 2
func S三元[T any](value bool, string1, string2 T) T {
	if value {
		return string1
	}
	return string2
}

// D多项选择 泛型多项选择器，根据索引从数组中选取对应元素。
// 索引从 0 开始，越界时返回默认值。
//
// 类型参数:
//   - T: 任意类型
//
// 参数:
//   - index: 选择的索引位置（从 0 开始）
//   - arr: 候选值数组
//   - 默认值: 索引越界时返回的默认值
//
// 返回:
//   - T: 选中的值或默认值
//
// 示例:
//
//	D多项选择(1, []string{"a","b","c"}, "x")  // "b"
//	D多项选择(5, []int{1,2,3}, 0)              // 0（越界返回默认值）
func D多项选择[T any](index int, arr []T, 默认值 T) T {
	if len(arr) < index+1 {
		return 默认值
	}
	return arr[index]
}

// G格式化文本 使用格式化字符串生成文本，等同于 fmt.Sprintf。
// 支持所有 fmt 包的格式化动词（%d、%s、%v 等）。
//
// 参数:
//   - str: 格式化字符串模板
//   - 参数: 格式化参数，可变数量
//
// 返回:
//   - string: 格式化后的文本
//
// 示例:
//
//	G格式化文本("姓名:%s,年龄:%d", "张三", 25)  // "姓名:张三,年龄:25"
func G格式化文本(str string, 参数 ...interface{}) string {
	return fmt.Sprintf(str, 参数...)
}

// G格式化_JSON 将 JSON 字符串进行缩进格式化，便于阅读。
// 使用 4 空格缩进，如果格式化失败则返回原始字符串。
//
// 参数:
//   - data: 待格式化的 JSON 字符串
//
// 返回:
//   - string: 格式化后的 JSON 文本；失败时返回原字符串
func G格式化_JSON(data string) string {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return string(data)
	}
	return string(b)
}

// D到文本数组 将通用型变量（interface{}）转换为字符串数组。
// 支持数组、切片类型的元素逐一转换；非数组类型作为单元素数组返回。
// 使用反射（reflect）实现类型判断和元素提取。
//
// 参数:
//   - 通用型变量: 待转换的值，支持 []string、[]int 等数组/切片类型
//
// 返回:
//   - []string: 转换后的字符串数组
func D到文本数组(通用型变量 interface{}) []string {
	var 局_文本数组 []string
	aa := reflect.TypeOf(通用型变量).Kind()
	if aa == reflect.Array || aa == reflect.Slice {
		for v := range reflect.ValueOf(通用型变量).Len() {
			index := reflect.ValueOf(通用型变量).Index(v)
			nameValue, ok := index.Interface().(string)
			if ok {
				局_文本数组 = append(局_文本数组, nameValue)
			}
		}
	} else {
		局_文本数组 = append(局_文本数组, 通用型变量.(string))
	}
	return 局_文本数组
}

// S是否为数组 判断通用型变量是否为数组或切片类型。
// 使用反射检查值的 Kind 是否为 Array 或 Slice。
//
// 参数:
//   - 通用型变量: 待判断的值
//
// 返回:
//   - bool: true 表示是数组/切片类型
func S是否为数组(通用型变量 interface{}) bool {
	aa := reflect.TypeOf(通用型变量).Kind()
	return (aa == reflect.Array || aa == reflect.Slice)
}

// W文本到utf8 将 GBK 编码的字符串转换为 UTF-8 编码。
// 使用 golang.org/x/text 库进行编码转换，转换失败时返回原字符串。
//
// 参数:
//   - src: GBK 编码的源字符串
//
// 返回:
//   - string: UTF-8 编码的字符串；失败时返回原字符串
func W文本到utf8(src string) string {
	reader := strings.NewReader(src)
	transformer := simplifiedchinese.GBK.NewDecoder()
	result, err := ioutil.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		return src
	}
	return string(result)
}

// Utf8到文本 将 UTF-8 编码的字符串转换为 GBK 编码。
// 使用 golang.org/x/text 库进行编码转换，转换失败时返回原字符串。
//
// 参数:
//   - src: UTF-8 编码的源字符串
//
// 返回:
//   - string: GBK 编码的字符串；失败时返回原字符串
func Utf8到文本(src string) string {
	reader := strings.NewReader(src)
	transformer := simplifiedchinese.GBK.NewEncoder()
	result, err := ioutil.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		return src
	}
	return string(result)
}

// Q取随机数 生成指定范围内的随机整数，包含边界值 [min, max]。
// 注意：Go 1.20+ 之前版本未自动初始化随机种子，高频调用可能产生相同结果。
//
// 参数:
//   - min: 最小值（包含）
//   - max: 最大值（包含）
//
// 返回:
//   - int: [min, max] 范围内的随机整数
func Q取随机数(min, max int) int {
	return mrand.Intn(max-min+1) + min
}

var h汇编随机源 = mrand.New(mrand.NewSource(time.Now().UnixNano()))
var h汇编随机锁 sync.Mutex

func H汇编_取随机数(起始数, 结束数 int) int {
	h汇编随机锁.Lock()
	defer h汇编随机锁.Unlock()
	return h汇编随机源.Intn(结束数-起始数+1) + 起始数
}

func H汇编_取随机字节(长度 int) ([]byte, error) {
	b := make([]byte, 长度)
	_, err := rand.Read(b)
	return b, err
}
