package utils

import (
	"net/url"
	"reflect"
)

func Map_取key整数数组(m map[int]string) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

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

func Map_转post数据(URL参数 map[string]string, 是否url编码 bool) string {
	局_返回 := ""
	if 是否url编码 {
		queryString := url.Values{}
		for key, value := range URL参数 {
			queryString.Set(key, value)
		}
		局_返回 = queryString.Encode()
	} else {
		//?origin=https:%252F%252Fstore.steampowered.com&input_protobuf_encoded=CgoxMTExMTExMTEx
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
