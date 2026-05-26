package main

import (
	"encoding/json"

	"github.com/yuan71058/Efunc/utils"
)

func registerCoreFunctions() {
	r := globalRegistry

	r.Register("D到文本", []string{"value"}, "转换为文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"value"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.D到文本(val))
		})

	r.Register("D到整数", []string{"value"}, "转换为整数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"value"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.D到整数(val))
		})

	r.Register("D到数值", []string{"value"}, "转换为浮点数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"value"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.D到数值(val))
		})

	r.Register("D到字节集", []string{"value"}, "转换为字节集(返回Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"value"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			data := utils.D到字节集(val)
			return okResult(encodeBase64(data))
		})

	r.Register("Q取随机数", []string{"min", "max"}, "获取随机整数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Min int `json:"min"`
				Max int `json:"max"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.Q取随机数(v.Min, v.Max))
		})

	r.Register("C类型_到文本", []string{"值"}, "类型转换到文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.C类型_到文本(val))
		})

	r.Register("C类型_到整数", []string{"值"}, "类型转换到整数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.C类型_到整数(val))
		})

	r.Register("C类型_到浮点数", []string{"值"}, "类型转换到浮点数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.C类型_到浮点数(val))
		})

	r.Register("C类型_到逻辑型", []string{"值"}, "类型转换到布尔",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.C类型_到逻辑型(val))
		})

	r.Register("C类型_安全到文本", []string{"值", "默认值"}, "安全转换到文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值   json.RawMessage `json:"值"`
				默认值 string          `json:"默认值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.C类型_安全到文本(val, v.默认值))
		})

	r.Register("C类型_进制转换", []string{"文本", "进制"}, "进制转换",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"文本"`
				进制 int    `json:"进制"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			result, err := utils.C类型_进制转换(v.文本, v.进制)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(result)
		})
}

func registerArrayFunctions() {
	r := globalRegistry

	r.Register("S数组_是否为空", []string{"list"}, "判断字符串数组是否为空",
		func(p json.RawMessage) *CallResult {
			var v struct {
				List []string `json:"list"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_是否为空(v.List))
		})

	r.Register("S数组_取文本索引", []string{"文本数组", "文本"}, "查找文本索引",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数组 []string `json:"文本数组"`
				文本 string   `json:"文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_取文本索引(v.数组, v.文本))
		})

	r.Register("S数组_取文本出现次数", []string{"参数_数组", "参数_成员"}, "统计成员出现次数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数组 []string `json:"参数_数组"`
				成员 string   `json:"参数_成员"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_取文本出现次数(v.数组, v.成员))
		})

	r.Register("S数组_整数是否存在", []string{"数组", "整数"}, "检查整数是否存在",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数组 []int `json:"数组"`
				整数 int   `json:"整数"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_整数是否存在(v.数组, v.整数))
		})

	r.Register("S数组_排序整数", []string{"arr"}, "整数数组升序排序",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Arr []int `json:"arr"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_排序整数(v.Arr))
		})

	r.Register("S数组_排序文本", []string{"arr"}, "字符串数组排序",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Arr []string `json:"arr"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_排序文本(v.Arr))
		})

	r.Register("S数组_求平均值", []string{"参数"}, "整数数组求平均值",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Arr []int `json:"参数"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_求平均值(v.Arr))
		})

	r.Register("S数组_取随机成员", []string{"源数组", "数量"}, "随机选取成员",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数组 []string `json:"源数组"`
				数量 int      `json:"数量"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_取随机成员(v.数组, v.数量))
		})

	r.Register("S数组_整数取差集", []string{"int1", "int2"}, "整数数组差集",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Arr1 []int `json:"int1"`
				Arr2 []int `json:"int2"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_整数取差集(v.Arr1, v.Arr2))
		})

	r.Register("S数组_取差集", []string{"a", "b"}, "整数数组差集",
		func(p json.RawMessage) *CallResult {
			var v struct {
				A []int `json:"a"`
				B []int `json:"b"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S数组_取差集(v.A, v.B))
		})
}
