package utils

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Array_RandomMembers 从源数组中随机选取指定数量的成员，不重复选取。
// 如果请求数量超过数组长度，自动调整为数组长度。
//
// 参数:
//   - source: 候选字符串数组
//   - count: 要选取的成员数量
//
// 返回:
//   - []string: 随机选取的成员数组
func Array_RandomMembers(source []string, count int) []string {
	if count > len(source) {
		count = len(source)
	}

	rand.Seed(time.Now().UnixNano())

	result := make([]string, count)

	copied := make([]string, len(source))
	copy(copied, source)

	for i := 0; i < count; i++ {
		index := rand.Intn(len(copied))
		result[i] = copied[index]
		copied = append(copied[:index], copied[index+1:]...)
	}

	return result
}

// Array_ToString 将 interface{} 数组转换为逗号分隔的文本。
// 去除方括号，用逗号替换空格分隔符。
//
// 参数:
//   - array: 待转换的 interface{} 数组
//
// 返回:
//   - string: 逗号分隔的文本
func Array_ToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// Array_Reverse 原地反转 interface{} 数组的元素顺序。
// 直接修改传入的切片，无需返回值。
//
// 参数:
//   - arr: 待反转的数组切片指针
func Array_Reverse(arr []interface{}) {
	count := len(arr)
	mid := count / 2

	for n := 0; n < mid; n++ {
		tmp := arr[n]
		arr[n] = arr[count-1]
		arr[count-1] = tmp
		count--
	}
}

// Array_Join 泛型数组合并为文本，用指定连接字符连接。
// 每个元素通过 ToString 转换为字符串后拼接。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - arr: 待合并的泛型数组
//   - separator: 元素之间的连接符
//
// 返回:
//   - string: 合并后的文本
func Array_Join[T comparable](arr []T, separator string) string {
	result := ""
	for _, item := range arr {
		result += ToString(item) + separator
	}
	result = strings.TrimSuffix(result, separator)
	return result
}

// Array_CountText 统计指定成员在字符串数组中出现的次数。
//
// 参数:
//   - arr: 源字符串数组
//   - member: 要统计的成员
//
// 返回:
//   - int: 出现次数
func Array_CountText(arr []string, member string) int {
	n := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == member {
			n++
		}
	}
	return n
}

// Array_FindText 查找文本在字符串数组中的索引位置。
//
// 参数:
//   - arr: 源字符串数组
//   - text: 要查找的文本
//
// 返回:
//   - int: 找到的索引（从 0 开始）；未找到返回 -1
func Array_FindText(arr []string, text string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == text {
			return i
		}
	}
	return -1
}

// Array_ContainsInt 检查整数数组中是否包含指定整数。
//
// 参数:
//   - arr: 源整数数组
//   - value: 要查找的整数
//
// 返回:
//   - bool: true 表示存在
func Array_ContainsInt(arr []int, value int) bool {
	for _, num := range arr {
		if num == value {
			return true
		}
	}
	return false
}

// Array_Contains 泛型检查数组中是否包含指定元素。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - arr: 源数组
//   - elem: 要查找的元素
//
// 返回:
//   - bool: true 表示存在
func Array_Contains[T comparable](arr []T, elem T) bool {
	for _, item := range arr {
		if item == elem {
			return true
		}
	}
	return false
}

// Array_Average 计算整数数组的平均值（整数除法，向下取整）。
//
// 参数:
//   - arr: 整数数组
//
// 返回:
//   - int: 平均值
func Array_Average(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum / len(arr)
}

// Array_IsEmpty 判断字符串数组是否为空或所有元素都是空白字符串。
// 空数组返回 true，所有元素均为空格/空串也返回 true。
//
// 参数:
//   - list: 待判断的字符串数组
//
// 返回:
//   - bool: true 表示数组为空或全为空白
func Array_IsEmpty(list []string) (isEmpty bool) {
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

// Array_SortInt 对整数数组进行升序排序，返回排序后的新数组。
//
// 参数:
//   - arr: 待排序的整数数组
//
// 返回:
//   - []int: 排序后的新数组
func Array_SortInt(arr []int) []int {
	sorted := arr
	sort.Ints(sorted)
	return sorted
}

// Array_SortText 对字符串数组进行字典序升序排序，返回排序后的新数组。
//
// 参数:
//   - arr: 待排序的字符串数组
//
// 返回:
//   - []string: 排序后的新数组
func Array_SortText(arr []string) []string {
	sorted := arr
	sort.Strings(sorted)
	return sorted
}

// Array_Unique 泛型数组去重，保持元素首次出现的顺序。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - arr: 待去重的数组
//
// 返回:
//   - []T: 去重后的新数组
func Array_Unique[T comparable](arr []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, v := range arr {
		if _, exists := seen[v]; !exists {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Array_Shuffle 泛型数组随机打乱顺序（Fisher-Yates 洗牌算法）。
// 返回乱序后的新数组，不修改原数组。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - arr: 待打乱的数组
//
// 返回:
//   - []T: 乱序后的新数组
func Array_Shuffle[T comparable](arr []T) []T {
	result := make([]T, len(arr))
	copy(result, arr)

	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Array_DiffInt 获取数组2中有但数组1中没有的成员。
//
// 参数:
//   - arr1: 基准整数数组
//   - arr2: 待比较的整数数组
//
// 返回:
//   - []int: 差集结果数组
func Array_DiffInt(arr1 []int, arr2 []int) []int {
	existingMembers := make(map[int]bool)
	nonExistingMembers := []int{}

	for _, member := range arr1 {
		existingMembers[member] = true
	}

	for _, member := range arr2 {
		if _, exists := existingMembers[member]; !exists {
			nonExistingMembers = append(nonExistingMembers, member)
		}
	}

	return nonExistingMembers
}

// Array_Diff 获取数组 a 中有但数组 b 中没有的元素。
//
// 参数:
//   - a: 源整数数组
//   - b: 待排除的整数数组
//
// 返回:
//   - []int: 差集结果数组
func Array_Diff(a, b []int) []int {
	m := make(map[int]bool)
	for _, v := range b {
		m[v] = true
	}

	var result []int
	for _, v := range a {
		if !m[v] {
			result = append(result, v)
		}
	}

	return result
}