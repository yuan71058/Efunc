package main

import (
	"encoding/json"

	"github.com/yuan71058/Efunc/utils"
)

func registerTextFunctions() {
	r := globalRegistry

	r.Register("W文本_取长度", []string{"value"}, "获取文本字符数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Value string `json:"value"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取长度(v.Value))
		})

	r.Register("W文本_是否包含关键字", []string{"内容", "关键字"}, "检查是否包含关键字",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容  string `json:"内容"`
				关键字 string `json:"关键字"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否包含关键字(v.内容, v.关键字))
		})

	r.Register("W文本_是否存在", []string{"内容", "关键字"}, "检查是否包含关键字(别名)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容  string `json:"内容"`
				关键字 string `json:"关键字"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否存在(v.内容, v.关键字))
		})

	r.Register("W文本_是否存在_任意", []string{"内容", "关键字"}, "检查是否包含任意关键字",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容  string   `json:"内容"`
				关键字 []string `json:"关键字"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否存在_任意(v.内容, v.关键字))
		})

	r.Register("W文本_是否存在_同时", []string{"内容", "关键字"}, "检查是否同时包含所有关键字",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容  string   `json:"内容"`
				关键字 []string `json:"关键字"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否存在_同时(v.内容, v.关键字))
		})

	r.Register("W文本_取出中间文本", []string{"内容", "左边文本", "右边文本"}, "提取左右标记之间的文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容   string `json:"内容"`
				左边文本 string `json:"左边文本"`
				右边文本 string `json:"右边文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取出中间文本(v.内容, v.左边文本, v.右边文本))
		})

	r.Register("W文本_取出中间文本_批量正则", []string{"内容", "左边文本", "右边文本"}, "正则批量提取中间文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容   string `json:"内容"`
				左边文本 string `json:"左边文本"`
				右边文本 string `json:"右边文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取出中间文本_批量正则(v.内容, v.左边文本, v.右边文本))
		})

	r.Register("W文本_取左边", []string{"欲取其部分的文本", "欲取出字符的数目"}, "从左侧截取指定字符数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本     string `json:"欲取其部分的文本"`
				字符数目 int    `json:"欲取出字符的数目"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取左边(v.文本, v.字符数目))
		})

	r.Register("W文本_取右边", []string{"欲取其部分的文本", "欲取出字符的数目"}, "从右侧截取指定字符数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本     string `json:"欲取其部分的文本"`
				字符数目 int    `json:"欲取出字符的数目"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取右边(v.文本, v.字符数目))
		})

	r.Register("W文本_取文本左边", []string{"内容", "关键字"}, "获取关键字左侧文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容  string `json:"内容"`
				关键字 string `json:"关键字"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取文本左边(v.内容, v.关键字))
		})

	r.Register("W文本_取文本右边", []string{"内容", "关键字"}, "获取关键字右侧文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容  string `json:"内容"`
				关键字 string `json:"关键字"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取文本右边(v.内容, v.关键字))
		})

	r.Register("W文本_替换", []string{"源文本", "旧文本", "新文本"}, "替换文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				源文本 string `json:"源文本"`
				旧文本 string `json:"旧文本"`
				新文本 string `json:"新文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_替换(v.源文本, v.旧文本, v.新文本))
		})

	r.Register("W文本_子文本替换", []string{"欲被替换的文本", "欲被替换的子文本", "用作替换的子文本"}, "替换子文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				源文本 string `json:"欲被替换的文本"`
				旧文本 string `json:"欲被替换的子文本"`
				新文本 string `json:"用作替换的子文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_子文本替换(v.源文本, v.旧文本, v.新文本))
		})

	r.Register("W文本_删首尾空", []string{"内容"}, "去除首尾空白",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容 string `json:"内容"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_删首尾空(v.内容))
		})

	r.Register("W文本_删首空", []string{"欲删除空格的文本"}, "去除左侧空格",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"欲删除空格的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_删首空(v.文本))
		})

	r.Register("W文本_删尾空", []string{"欲删除空格的文本"}, "去除右侧空格",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"欲删除空格的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_删尾空(v.文本))
		})

	r.Register("W文本_到大写", []string{"value"}, "转换为大写",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Value string `json:"value"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_到大写(v.Value))
		})

	r.Register("W文本_到小写", []string{"value"}, "转换为小写",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Value string `json:"value"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_到小写(v.Value))
		})

	r.Register("W文本_首字母改大写", []string{"英文文本"}, "首字母大写",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"英文文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_首字母改大写(v.文本))
		})

	r.Register("W文本_分割文本", []string{"待分割文本", "用作分割的文本"}, "按分隔符分割",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本   string `json:"待分割文本"`
				分隔符 string `json:"用作分割的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_分割文本(v.文本, v.分隔符))
		})

	r.Register("W文本_逐字分割", []string{"原文本"}, "逐字符拆分",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"原文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_逐字分割(v.文本))
		})

	r.Register("W文本_寻找", []string{"源文本", "要寻找的文本"}, "查找文本位置",
		func(p json.RawMessage) *CallResult {
			var v struct {
				源文本 string `json:"源文本"`
				目标   string `json:"要寻找的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_寻找(v.源文本, v.目标))
		})

	r.Register("W文本_寻找文本", []string{"被搜寻的文本", "欲寻找的文本"}, "查找文本位置(别名)",
		func(p json.RawMessage) *CallResult {
			var v struct {
				源文本 string `json:"被搜寻的文本"`
				目标   string `json:"欲寻找的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_寻找文本(v.源文本, v.目标))
		})

	r.Register("W文本_倒找文本", []string{"被搜寻的文本", "欲寻找的文本"}, "从后查找文本位置",
		func(p json.RawMessage) *CallResult {
			var v struct {
				源文本 string `json:"被搜寻的文本"`
				目标   string `json:"欲寻找的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_倒找文本(v.源文本, v.目标))
		})

	r.Register("W文本_取出现次数", []string{"被搜索文本", "欲搜索文本"}, "统计出现次数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				源文本 string `json:"被搜索文本"`
				目标   string `json:"欲搜索文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取出现次数(v.源文本, v.目标))
		})

	r.Register("W文本_取重复", []string{"重复次数", "待重复文本"}, "重复文本",
		func(p json.RawMessage) *CallResult {
			var v struct {
				次数 int    `json:"重复次数"`
				文本 string `json:"待重复文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取重复(v.次数, v.文本))
		})

	r.Register("W文本_取空白", []string{"重复次数"}, "生成空格字符串",
		func(p json.RawMessage) *CallResult {
			var v struct {
				次数 int `json:"重复次数"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取空白(v.次数))
		})

	r.Register("W文本_取行数", []string{"文本"}, "统计行数",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取行数(v.文本))
		})

	r.Register("W文本_删除空行", []string{"要操作的文本"}, "删除空行",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"要操作的文本"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_删除空行(v.文本))
		})

	r.Register("W文本_去重复文本", []string{"原文本", "分割符"}, "去重",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本   string `json:"原文本"`
				分隔符 string `json:"分割符"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_去重复文本(v.文本, v.分隔符))
		})

	r.Register("W文本_是否JSON", []string{"s"}, "检查是否为JSON",
		func(p json.RawMessage) *CallResult {
			var v struct {
				S string `json:"s"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否JSON(v.S))
		})

	r.Register("W文本_是否为英数字母", []string{"s"}, "检查是否仅含英数字母",
		func(p json.RawMessage) *CallResult {
			var v struct {
				S string `json:"s"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否为英数字母(v.S))
		})

	r.Register("W文本_是否为数字", []string{"s"}, "检查是否仅含数字",
		func(p json.RawMessage) *CallResult {
			var v struct {
				S string `json:"s"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否为数字(v.S))
		})

	r.Register("W文本_是否为字母", []string{"s"}, "检查是否仅含字母",
		func(p json.RawMessage) *CallResult {
			var v struct {
				S string `json:"s"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_是否为字母(v.S))
		})

	r.Register("W文本_取随机字符串", []string{"字符串长度"}, "生成随机字符串",
		func(p json.RawMessage) *CallResult {
			var v struct {
				长度 int `json:"字符串长度"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取随机字符串(v.长度))
		})

	r.Register("W文本_取随机字符串_数字", []string{"字符串长度"}, "生成随机数字字符串",
		func(p json.RawMessage) *CallResult {
			var v struct {
				长度 int `json:"字符串长度"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取随机字符串_数字(v.长度))
		})

	r.Register("W文本_去除敏感信息", []string{"内容"}, "文本脱敏",
		func(p json.RawMessage) *CallResult {
			var v struct {
				内容 string `json:"内容"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_去除敏感信息(v.内容))
		})

	r.Register("W文本_颠倒", []string{"欲转换文本", "带有中文"}, "文本反转",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本   string `json:"欲转换文本"`
				中文模式 bool   `json:"带有中文"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_颠倒(v.文本, v.中文模式))
		})

	r.Register("W文本_取指定变量文本行", []string{"文本", "行号"}, "获取指定行内容",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"文本"`
				行号 int    `json:"行号"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取指定变量文本行(v.文本, v.行号))
		})

	r.Register("W文本_删除指定文本行", []string{"源文本", "行数"}, "删除指定行",
		func(p json.RawMessage) *CallResult {
			var v struct {
				文本 string `json:"源文本"`
				行号 int    `json:"行数"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_删除指定文本行(v.文本, v.行号))
		})

	r.Register("W文本_取文本所在行", []string{"源文本", "欲查找的文本", "是否区分大小写"}, "查找文本所在行号",
		func(p json.RawMessage) *CallResult {
			var v struct {
				源文本 string `json:"源文本"`
				目标   string `json:"欲查找的文本"`
				区分   bool   `json:"是否区分大小写"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_取文本所在行(v.源文本, v.目标, v.区分))
		})

	r.Register("W文本_字符", []string{"字节型"}, "字节值转字符",
		func(p json.RawMessage) *CallResult {
			var v struct {
				值 int8 `json:"字节型"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_字符(v.值))
		})

	r.Register("W文本_gbk到utf8", []string{"src"}, "GBK转UTF8",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Src string `json:"src"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_gbk到utf8(v.Src))
		})

	r.Register("W文本_utf8到gbk", []string{"src"}, "UTF8转GBK",
		func(p json.RawMessage) *CallResult {
			var v struct {
				Src string `json:"src"`
			}
			if err := json.Unmarshal(p, &v); err != nil {
				return errResult(err.Error())
			}
			return okResult(utils.W文本_utf8到gbk(v.Src))
		})
}
