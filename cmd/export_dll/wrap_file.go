package main

import (
	"encoding/json"

	"github.com/yuan71058/Efunc/utils"
)

func registerFileFunctions() {
	r := globalRegistry

	r.Register("W文件_是否存在", []string{"路径"}, "判断文件或目录是否存在",
		func(p json.RawMessage) *CallResult {
			var v struct {
				路径 string `json:"路径"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文件_是否存在(v.路径))
		})

	r.Register("W文件_写到文件", []string{"文件名", "欲写入文件的数据"}, "写入字节数据到文件(Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string `json:"文件名"`
				数据   string `json:"欲写入文件的数据"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data, err := decodeBase64(v.数据)
			if err != nil {
				return errResultf("base64解码失败: %v", err)
			}
			if err := utils.W文件_写到文件(v.文件名, data); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})

	r.Register("W文件_读入文本", []string{"文件名"}, "读取文件全部内容为文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string `json:"文件名"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文件_读入文本(v.文件名))
		})

	r.Register("W文件_读入文件", []string{"文件名"}, "读取文件全部内容为字节集(Base64)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string `json:"文件名"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			data := utils.W文件_读入文件(v.文件名)
			return okResult(encodeBase64(data))
		})

	r.Register("W文件_删除", []string{"欲删除的文件名"}, "删除文件",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string `json:"欲删除的文件名"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			if err := utils.W文件_删除(v.文件名); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})

	r.Register("W文件_更名", []string{"欲更名的原文件或目录名", "欲更改为的现文件或目录名"}, "重命名文件或目录",
		func(p json.RawMessage) *CallResult {
			var v struct {
				原名 string `json:"欲更名的原文件或目录名"`
				新名 string `json:"欲更改为的现文件或目录名"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			if err := utils.W文件_更名(v.原名, v.新名); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})

	r.Register("W文件_取文件名", []string{"路径"}, "从路径提取文件名",
		func(p json.RawMessage) *CallResult {
			var v struct {
				路径 string `json:"路径"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文件_取文件名(v.路径))
		})

	r.Register("W文件_取父目录", []string{"dirpath"}, "获取父目录",
		func(p json.RawMessage) *CallResult {
			var v struct {
				路径 string `json:"dirpath"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文件_取父目录(v.路径))
		})

	r.Register("W文件_取大小", []string{"文件名"}, "获取文件大小",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string `json:"文件名"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文件_取大小(v.文件名))
		})

	r.Register("W文件_追加文本", []string{"文件名", "欲追加的文本"}, "追加文本到文件末尾",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string `json:"文件名"`
				文本   string `json:"欲追加的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			if err := utils.W文件_追加文本(v.文件名, v.文本); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})

	r.Register("W文件_写出", []string{"文件名", "欲写入文件的数据"}, "写出数据到文件(支持任意类型)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string          `json:"文件名"`
				数据   json.RawMessage `json:"欲写入文件的数据"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var data interface{}
			if err := json.Unmarshal(v.数据, &data); err != nil {
				return errResult(err.Error())
			}
			if err := utils.W文件_写出(v.文件名, data); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})

	r.Register("W文件_保存", []string{"文件名", "欲写入文件的数据"}, "智能保存(内容不同才写入)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文件名 string          `json:"文件名"`
				数据   json.RawMessage `json:"欲写入文件的数据"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			var data interface{}
			if err := json.Unmarshal(v.数据, &data); err != nil {
				return errResult(err.Error())
			}
			if err := utils.W文件_保存(v.文件名, data); err != nil {
				return errResult(err.Error())
			}
			return okResult(true)
		})
}
