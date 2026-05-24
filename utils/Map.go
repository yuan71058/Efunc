package utils

import (
	"net/url"
	"reflect"
)

// Map_取key整数数组 从 map[int]string 中提取所有键名，返回整数数组。
// 预分配了 map 长度的数组容量，避免 append 时的内存重分配。
//
// 参数:
//   - m: 整数键的 map
//
// 返回:
//   - []int: 所有键名组成的数组
func Map_取key整数数组(m map[int]string) []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

// Map_Struct转Map 利用反射将结构体转换为 map[string]interface{}。
// 结构体字段如果带有 "mapstructure" tag，则使用 tag 值作为 map 的键名；
// 否则使用结构体字段的原始名称。
//
// 参数:
//   - obj: 结构体实例（非指针）
//
// 返回:
//   - map[string]interface{}: 转换后的 map
//
// 示例:
//
//	type User struct {
//	    Name string `mapstructure:"user_name"`
//	    Age  int
//	}
//	Map_Struct转Map(User{Name: "张三", Age: 25})
//	// map[string]interface{}{"user_name": "张三", "Age": 25}
func Map_Struct转Map(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

// Map_转post数据 将 map[string]string 转换为 POST 表单格式的字符串。
// 格式：key1=value1&key2=value2
//
// 参数:
//   - URL参数: 键值对 map
//   - 是否url编码: true 时使用 url.Values.Encode 进行标准 URL 编码
//
// 返回:
//   - string: 拼接后的 POST 数据字符串
func Map_转post数据(URL参数 map[string]string, 是否url编码 bool) string {
	局_返回 := ""
	if 是否url编码 {
		queryString := url.Values{}
		for key, value := range URL参数 {
			queryString.Set(key, value)
		}
		局_返回 = queryString.Encode()
	} else {
		for key, value := range URL参数 {
			if 局_返回 == "" {
				局_返回 += key + "=" + value
				continue
			}
			局_返回 += "&" + key + "=" + value
		}
	}

	return 局_返回
}

// Map_键名是否存在 检查 map[int]string 中是否存在指定的键名。
//
// 参数:
//   - m: 待检查的 map
//   - key: 要查找的键名
//
// 返回:
//   - bool: true 表示键名存在
func Map_键名是否存在(m map[int]string, key int) bool {
	_, ok := m[key]
	return ok
}
