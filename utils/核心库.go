package utils

import "github.com/gogf/gf/v2/util/gconv"
import "fmt"
import "reflect"
import "encoding/json"
import "golang.org/x/text/encoding/simplifiedchinese"
import "golang.org/x/text/transform"
import (
	"io/ioutil"
	"math/rand"
	"strings"
)

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
func D到整数64(value interface{}) int64 {
	return gconv.Int64(value)
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
func G格式化_JSON(data string) string {
	b, err := json.MarshalIndent(data, "", "    ") // 4空格缩进
	if err != nil {
		return string(data)
	}
	return string(b)
}

// 可以将通用型的文本数组转换成文本数组  还挺麻烦查了很多资料
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

// 可以将通用型的文本数组转换成文本数组  还挺麻烦查了很多资料
func S是否为数组(通用型变量 interface{}) bool {
	aa := reflect.TypeOf(通用型变量).Kind()
	return (aa == reflect.Array || aa == reflect.Slice)

}

func W文本到utf8(src string) string {
	reader := strings.NewReader(src)
	transformer := simplifiedchinese.GBK.NewDecoder()
	result, err := ioutil.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		return src // 转换失败返回原字符串
	}
	return string(result)
}

func Utf8到文本(src string) string {
	reader := strings.NewReader(src)
	transformer := simplifiedchinese.GBK.NewEncoder()
	result, err := ioutil.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		return src // 转换失败返回原字符串
	}
	return string(result)
}
func Q取随机数(min, max int) int {
	return rand.Intn(max-min+1) + min
}
