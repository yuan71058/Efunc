package utils

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// S数组_取随机成员 从源数组中随机选取指定数量的成员，不重复选取。
// 如果请求数量超过数组长度，自动调整为数组长度。
//
// 参数:
//   - 源数组: 候选字符串数组
//   - 数量: 要选取的成员数量
//
// 返回:
//   - []string: 随机选取的成员数组
func S数组_取随机成员(源数组 []string, 数量 int) []string {
	if 数量 > len(源数组) {
		数量 = len(源数组)
	}

	rand.Seed(time.Now().UnixNano())

	结果 := make([]string, 数量)

	复制数组 := make([]string, len(源数组))
	copy(复制数组, 源数组)

	for i := 0; i < 数量; i++ {
		索引 := rand.Intn(len(复制数组))
		结果[i] = 复制数组[索引]
		复制数组 = append(复制数组[:索引], 复制数组[索引+1:]...)
	}

	return 结果
}

// S数组_到文本 将 interface{} 数组转换为逗号分隔的文本。
// 去除方括号，用逗号替换空格分隔符。
//
// 参数:
//   - array: 待转换的 interface{} 数组
//
// 返回:
//   - string: 逗号分隔的文本
func S数组_到文本(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// S数组_反转 原地反转 interface{} 数组的元素顺序。
// 直接修改传入的切片，无需返回值。
//
// 参数:
//   - 反转的数组切片: 待反转的数组切片指针
func S数组_反转(反转的数组切片 []interface{}) {
	成员数量 := len(反转的数组切片)
	折中数量 := 成员数量 / 2

	for N := 0; N < 折中数量; N++ {
		临时成员 := 反转的数组切片[N]
		反转的数组切片[N] = 反转的数组切片[成员数量-1]
		反转的数组切片[成员数量-1] = 临时成员
		成员数量--
	}
}

// S数组_合并文本 泛型数组合并为文本，用指定连接字符连接。
// 每个元素通过 D到文本 转换为字符串后拼接。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - 数组: 待合并的泛型数组
//   - 连接字符: 元素之间的连接符
//
// 返回:
//   - string: 合并后的文本
func S数组_合并文本[T comparable](数组 []T, 连接字符 string) string {
	result := ""
	for _, item := range 数组 {
		result += D到文本(item) + 连接字符
	}
	result = strings.TrimSuffix(result, 连接字符)
	return result

}

// S数组_取文本出现次数 统计指定成员在字符串数组中出现的次数。
//
// 参数:
//   - 参数_数组: 源字符串数组
//   - 参数_成员: 要统计的成员
//
// 返回:
//   - int: 出现次数
func S数组_取文本出现次数(参数_数组 []string, 参数_成员 string) int {
	n := 0
	for i := 0; i < len(参数_数组); i++ {
		if 参数_数组[i] == 参数_成员 {
			n++
		}
	}
	return n
}

// S数组_取文本索引 查找文本在字符串数组中的索引位置。
//
// 参数:
//   - 文本数组: 源字符串数组
//   - 文本: 要查找的文本
//
// 返回:
//   - int: 找到的索引（从 0 开始）；未找到返回 -1
func S数组_取文本索引(文本数组 []string, 文本 string) int {
	for i := 0; i < len(文本数组); i++ {
		if 文本数组[i] == 文本 {
			return i
		}
	}
	return -1
}

// S数组_整数是否存在 检查整数数组中是否包含指定整数。
//
// 参数:
//   - 数组: 源整数数组
//   - 整数: 要查找的整数
//
// 返回:
//   - bool: true 表示存在
func S数组_整数是否存在(数组 []int, 整数 int) bool {
	for _, num := range 数组 {
		if num == 整数 {
			return true
		}
	}
	return false
}

// S数组_是否存在 泛型检查数组中是否包含指定元素。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - 数组: 源数组
//   - 元素: 要查找的元素
//
// 返回:
//   - bool: true 表示存在
func S数组_是否存在[T comparable](数组 []T, 元素 T) bool {
	for _, item := range 数组 {
		if item == 元素 {
			return true
		}
	}
	return false
}

// S数组_求平均值 计算整数数组的平均值（整数除法，向下取整）。
//
// 参数:
//   - 参数: 整数数组
//
// 返回:
//   - int: 平均值
func S数组_求平均值(参数 []int) int {
	var 总和 int
	for _, v := range 参数 {
		总和 += v
	}
	return 总和 / len(参数)
}

// S数组_是否为空 判断字符串数组是否为空或所有元素都是空白字符串。
// 空数组返回 true，所有元素均为空格/空串也返回 true。
//
// 参数:
//   - list: 待判断的字符串数组
//
// 返回:
//   - bool: true 表示数组为空或全为空白
func S数组_是否为空(list []string) (isEmpty bool) {

	if len(list) == 0 {
		return true
	}

	isEmpty = true
	for _, f := range list {

		if strings.TrimSpace(f) != "" {
			isEmpty = false
			break
		}
	}

	return isEmpty
}

// S数组_排序整数 对整数数组进行升序排序，返回排序后的新数组。
//
// 参数:
//   - arr: 待排序的整数数组
//
// 返回:
//   - []int: 排序后的新数组
func S数组_排序整数(arr []int) []int {
	局_arr := arr
	sort.Ints(局_arr)
	return 局_arr
}

// S数组_排序文本 对字符串数组进行字典序升序排序，返回排序后的新数组。
//
// 参数:
//   - arr: 待排序的字符串数组
//
// 返回:
//   - []string: 排序后的新数组
func S数组_排序文本(arr []string) []string {
	局_arr := arr
	sort.Strings(局_arr)
	return 局_arr
}

// S数组_去重复 泛型数组去重，保持元素首次出现的顺序。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - 数组: 待去重的数组
//
// 返回:
//   - []T: 去重后的新数组
func S数组_去重复[T comparable](数组 []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, v := range 数组 {
		if _, exists := seen[v]; !exists {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// S数组_乱序 泛型数组随机打乱顺序（Fisher-Yates 洗牌算法）。
// 返回乱序后的新数组，不修改原数组。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - 数组: 待打乱的数组
//
// 返回:
//   - []T: 乱序后的新数组
func S数组_乱序[T comparable](数组 []T) []T {
	result := make([]T, len(数组))
	copy(result, 数组)

	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// S数组_整数取差集 获取数组2中有但数组1中没有的成员。
//
// 参数:
//   - int1: 基准整数数组
//   - int2: 待比较的整数数组
//
// 返回:
//   - []int: 差集结果数组
func S数组_整数取差集(int1 []int, int2 []int) []int {
	existingMembers := make(map[int]bool)
	nonExistingMembers := []int{}

	for _, member := range int1 {
		existingMembers[member] = true
	}

	for _, member := range int2 {
		if _, exists := existingMembers[member]; !exists {
			nonExistingMembers = append(nonExistingMembers, member)
		}
	}

	return nonExistingMembers
}

// S数组_取差集 获取数组 a 中有但数组 b 中没有的元素。
//
// 参数:
//   - a: 源整数数组
//   - b: 待排除的整数数组
//
// 返回:
//   - []int: 差集结果数组
func S数组_取差集(a, b []int) []int {
	m := make(map[int]bool)
	for _, v := range b {
		m[v] = true
	}

	var 结果 []int
	for _, v := range a {
		if !m[v] {
			结果 = append(结果, v)
		}
	}

	return 结果
}
