package utils

import (
	"fmt"
	"sort"
	"strings"
)

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

func S数组_合并文本(文本数组 []string, 连接字符 string) string {
	return strings.Join(文本数组, 连接字符)
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
