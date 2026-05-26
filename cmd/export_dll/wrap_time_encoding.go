package main

import (
	"encoding/base64"
	"encoding/json"

	"github.com/yuan71058/Efunc/utils"
)

func decodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func encodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func registerTimeFunctions() {
	r := globalRegistry

	r.Register("S时间_取现行时间戳", nil, "获取当前Unix时间戳(秒)",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.S时间_取现行时间戳())
		})

	r.Register("S时间_取现行时间", nil, "获取当前时间字符串",
		func(p json.RawMessage) *CallResult {
			return okResult(utils.S时间_取现行时间())
		})

	r.Register("S时间_时间戳格式化", []string{"format", "时间戳"}, "时间戳格式化",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Format string `json:"format"`
				时间戳  int64  `json:"时间戳"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S时间_时间戳格式化(v.Format, v.时间戳))
		})

	r.Register("S时间_是否闰年", []string{"年"}, "判断是否闰年",
		func(p json.RawMessage) *CallResult {
			var v struct {
				年 int `json:"年"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S时间_是否闰年(v.年))
		})

	r.Register("S时间_取月份天数", []string{"年", "月"}, "获取月份天数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				年 int `json:"年"`
				月 int `json:"月"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.S时间_取月份天数(v.年, v.月))
		})
}

func registerEncodingFunctions() {
	r := globalRegistry

	r.Register("B编码_URL编码", []string{"欲编码的文本"}, "URL编码",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"欲编码的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.B编码_URL编码(v.文本))
		})

	r.Register("B编码_URL解码", []string{"欲解码的文本"}, "URL解码",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"欲解码的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.B编码_URL解码(v.文本))
		})

	r.Register("B编码_BASE64编码", []string{"字节集"}, "Base64编码(输入Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"字节集"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			return okResult(utils.B编码_BASE64编码(data))
		})

	r.Register("B编码_BASE64解码", []string{"文本"}, "Base64解码(返回Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data := utils.B编码_BASE64解码(v.文本)
			return okResult(encodeBase64(data))
		})

	r.Register("B编码_十六进制编码", []string{"字节集"}, "十六进制编码(输入Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"字节集"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			return okResult(utils.B编码_十六进制编码(data))
		})

	r.Register("B编码_十六进制解码", []string{"文本"}, "十六进制解码(返回Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data := utils.B编码_十六进制解码(v.文本)
			return okResult(encodeBase64(data))
		})

	r.Register("B编码_UTF8到GBK", []string{"文本"}, "UTF8转GBK(返回Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data := utils.B编码_UTF8到GBK(v.文本)
			return okResult(encodeBase64(data))
		})

	r.Register("B编码_GBK到UTF8", []string{"数据"}, "GBK转UTF8(输入Base64字节集)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"数据"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			return okResult(utils.B编码_GBK到UTF8(data))
		})

	r.Register("B编码_JSON编码", []string{"值"}, "JSON编码",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 json.RawMessage `json:"值"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var val interface{}
			json.Unmarshal(v.值, &val)
			return okResult(utils.B编码_JSON编码(val))
		})

	r.Register("Z字节集_十六进制到字节集", []string{"原始16进制文本"}, "十六进制文本转字节集(返回Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"原始16进制文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data := utils.Z字节集_十六进制到字节集(v.文本)
			return okResult(encodeBase64(data))
		})

	r.Register("Z字节集_字节集到十六进制", []string{"字节集"}, "字节集转十六进制(输入Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"字节集"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			return okResult(utils.Z字节集_字节集到十六进制(data))
		})

	r.Register("Z字节集_Gzip解压", []string{"字节集"}, "Gzip解压(输入Base64,返回Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				数据 string `json:"字节集"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			result, err := utils.Z字节集_Gzip解压(data)
			if err != nil {
				return errResult(err.Error())
			}
			return okResult(encodeBase64(result))
		})
}
