package utils

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func S数组_取随机成员(源数组 []string, 数量 int) []string {
	if 数量 > len(源数组) {
		数量 = len(源数组)
	}

	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 创建一个切片来存储结果
	结果 := make([]string, 数量)

	// 复制源数组，避免修改原始数组
	复制数组 := make([]string, len(源数组))
	copy(复制数组, 源数组)

	// 随机选取成员
	for i := 0; i < 数量; i++ {
		// 生成一个随机索引
		索引 := rand.Intn(len(复制数组))
		// 将选中的成员添加到结果中
		结果[i] = 复制数组[索引]
		// 从复制数组中移除已选中的成员，避免重复选取
		复制数组 = append(复制数组[:索引], 复制数组[索引+1:]...)
	}

	return 结果
}
func S数组_到文本(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

// 反转的数组切片[:]  换成切片即可 传递指针,
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

func S数组_合并文本[T comparable](数组 []T, 连接字符 string) string {
	result := ""
	for _, item := range 数组 {
		result += D到文本(item) + 连接字符
	}
	result = strings.TrimSuffix(result, 连接字符)
	return result

}

func S数组_取文本出现次数(参数_数组 []string, 参数_成员 string) int {
	n := 0
	for i := 0; i < len(参数_数组); i++ {
		if 参数_数组[i] == 参数_成员 {
			n++
		}
	}
	return n
}

// 寻找 文本在数组中的索引,失败返回-1
func S数组_取文本索引(文本数组 []string, 文本 string) int {
	for i := 0; i < len(文本数组); i++ {
		if 文本数组[i] == 文本 {
			return i
		}
	}
	return -1
}

func S数组_整数是否存在(数组 []int, 整数 int) bool {
	for _, num := range 数组 {
		if num == 整数 {
			return true
		}
	}
	return false
}
func S数组_是否存在[T comparable](数组 []T, 元素 T) bool {
	for _, item := range 数组 {
		if item == 元素 {
			return true
		}
	}
	return false
}
func S数组_求平均值(参数 []int) int {
	var 总和 int
	for _, v := range 参数 {
		总和 += v
	}
	return 总和 / len(参数)
}

// 判断数组各元素是否是空字符串或空格
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

func S数组_排序整数(arr []int) []int {
	局_arr := arr
	sort.Ints(局_arr)
	return 局_arr
}

func S数组_排序文本(arr []string) []string {
	局_arr := arr
	sort.Strings(局_arr)
	return 局_arr
}
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

func S数组_乱序[T comparable](数组 []T) []T {
	// 创建一个新的切片用于存储乱序后的结果
	result := make([]T, len(数组))
	copy(result, 数组) // 复制原数组，避免修改原始数据

	// 使用 Fisher-Yates 算法进行乱序
	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                       // 生成 [0, i] 范围内的随机索引
		result[i], result[j] = result[j], result[i] // 交换元素
	}

	return result
}

// 获取数组2有但是数组1没有的成员数组
func S数组_整数取差集(int1 []int, int2 []int) []int {
	existingMembers := make(map[int]bool)
	nonExistingMembers := []int{}

	// 将 int1 中的成员添加到 existingMembers 中
	for _, member := range int1 {
		existingMembers[member] = true
	}

	// 检查 int2 中的成员是否存在于 existingMembers 中
	for _, member := range int2 {
		if _, exists := existingMembers[member]; !exists {
			nonExistingMembers = append(nonExistingMembers, member)
		}
	}

	return nonExistingMembers
}

// 差集函数，返回切片a中有但切片b中没有的元素
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
