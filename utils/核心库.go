package utils

import "github.com/gogf/gf/v2/util/gconv"
import "fmt"

// 字节数组
func D到字节集(value interface{}) []byte {
	return gconv.Bytes(value)
}
func D到字节(value interface{}) byte {
	return gconv.Byte(value)
}
func D到整数(value interface{}) int {
	return gconv.Int(value)
}

func D到数值(value interface{}) float64 {
	return gconv.Float64(value)
}
func D到文本(value interface{}) string {
	return gconv.String(value)
}
func D到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error {
	return gconv.Struct(待转换的参数, 结构体指针)
}

func S三元[T any](value bool, string1, string2 T) T {
	if value {
		return string1
	}
	return string2
}

// index 从0开始
func D多项选择[T any](index int, arr []T, 默认值 T) T {
	if len(arr) < index+1 {
		return 默认值
	}
	return arr[index]
}

func G格式化文本(str string, 参数 ...interface{}) string {
	return fmt.Sprintf(str, 参数...)
}
