package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// W文本_是否存在关键字  关键字为空 直接返回 真
func W文本_是否包含关键字(内容, 关键字 string) bool {
	return strings.Contains(内容, 关键字)
}

func W文本_是否为英数字母(s string) bool {
	pattern := "^[A-Za-z0-9]+$"
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func W文本_是否为字母(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func W文本_是否为数字(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

/*
.版本 2

.子程序 W文本_倒取出中间文本, 文本型, 公开, 比如：欲取全文本为“0012345”,现在要取出“123”，<123>的右边为“4”，<123>的左边为“0”，注意这里是倒取
.参数 欲取全文本, 文本型, , 比如：欲取全文本为“0012345”
.参数 右边文本, 文本型, , 123的右边为“4”，引号直接用 #引号，如："<font color=#引号red#引号>" 注意左右
.参数 左边文本, 文本型, , 123的左边为“0”，引号直接用 #引号，如："<font color=#引号red#引号>" 注意左右
.参数 倒数搜寻位置, 整数型, 可空, 可空,这里是指搜寻 参数二 右边文本的开始位置
.参数 是否不区分大小写, 逻辑型, 可空, 默认为假：区分大小写 真：不区分大小写
*/
func W文本_倒取出中间文本(欲取全文本 string, 右边文本 string, 左边文本 string, 倒数搜寻位置 int, 是否不区分大小写 bool) string {
	倒数搜寻位置 = 选择(倒数搜寻位置 <= 0, -1, 倒数搜寻位置).(int)
	倒数搜寻位置 = len(欲取全文本) - 倒数搜寻位置
	rPos := strings.LastIndex(欲取全文本, 右边文本)
	if rPos != -1 {
		lPos := strings.LastIndex(欲取全文本, 左边文本)
		if lPos != -1 {
			lPos += len(左边文本)
		}
		return 欲取全文本[lPos:rPos]
	}
	return ""
}

/*
.子程序 W文本_取文本所在行, 整数型, 公开, 查找某段字或关键中在文本中的哪一行出现，成功返回行数，失败返回0
.参数 源文本, 文本型
.参数 欲查找的文本, 文本型
.参数 是否区分大小写, 逻辑型, 可空
*/
func W文本_取文本所在行(源文本 string, 欲查找的文本 string, 是否区分大小写 bool) int {
	局文本 := strings.Split(源文本, "\n")
	for 局计次 := 0; 局计次 < len(局文本); 局计次++ {
		if strings.Index(局文本[局计次], 欲查找的文本) != -1 {
			return 局计次
		}
	}
	return 0
}

/*
.版本 2

.子程序 W文本_删除指定文本行, 文本型, 公开, 删除指定文本的一行文本，返回删行后的文本
.参数 源文本, 文本型
.参数 行数, 整数型, , 输入你想删除的行数，如：想删除第3行的整行文本就直接输3
*/
func W文本_删除指定文本行(源文本 string, 行数 int) string {
	临时文本 := strings.Split(源文本, "\n")
	输出文本 := ""
	if strings.Index(输出文本, "\n") == -1 {
		输出文本 += "\n"
	}
	临时文本 = append(临时文本[:行数], 临时文本[行数+1:]...)

	for 计次 := 0; 计次 < len(临时文本); 计次++ {
		if len(临时文本) != 计次 {
			输出文本 += 临时文本[计次] + "\n"
		} else {
			输出文本 += 临时文本[计次]
		}
	}
	return 输出文本
}

func W文本_取随机范围数字(起始数, 结束数, 单双选择 int) string {
	临时整数 := H汇编_取随机数(起始数, 结束数)
	if 单双选择 == 1 {
		if 临时整数%2 == 0 {
			if 临时整数 == 结束数 {
				临时整数 = 临时整数 - 1
			} else {
				临时整数 = 临时整数 + 1
			}
		}
	} else if 单双选择 == 2 {
		if 临时整数%2 == 1 {
			if 临时整数 == 结束数 {
				临时整数 = 临时整数 - 1
			} else {
				临时整数 = 临时整数 + 1
			}
		}
	}
	return fmt.Sprintf("%d", 临时整数)
}
func W文本_取指定变量文本行(文本 string, 行号 int) string {
	文本数组 := strings.Split(文本, "\n")
	if 行号 <= 0 {
		return ""
	}
	if 行号 > len(文本数组) {
		return ""
	}
	return 文本数组[行号-1]
}
func W文本_颠倒(欲转换文本 string, 带有中文 bool) string {
	if 带有中文 {
		字节集 := []byte(欲转换文本)
		局_结果 := make([]byte, 0, len(字节集))
		for i := len(字节集) - 1; i >= 0; i -= 2 {
			局_结果 = append(局_结果, 字节集[i-1], 字节集[i])
		}
		return string(局_结果)
	}

	//字符数 := utf8.RuneCountInString(欲转换文本)  可能后续要改
	倒序内容 := ""
	for _, r := range 欲转换文本 {
		倒序内容 = string(r) + 倒序内容
	}
	return 倒序内容
}

func W文本_取出现次数(被搜索文本 string, 欲搜索文本 string) int {
	i := 0
	位置_ := strings.IndexFunc(被搜索文本, func(r rune) bool {
		return strings.HasPrefix(被搜索文本, 欲搜索文本)
	})
	for 位置_ != -1 {
		i++
		位置_ = strings.IndexFunc(被搜索文本[位置_+len(欲搜索文本):], func(r rune) bool {
			return strings.HasPrefix(被搜索文本[位置_+len(欲搜索文本):], 欲搜索文本)
		})
	}
	return i
}

func W文本_首字母改大写(英文文本 string) string {
	return strings.ToUpper(string(英文文本[0])) + 英文文本[1:]
}
func W文本_替换(源文本, 旧文本, 新文本 string) string {
	return strings.Replace(源文本, 旧文本, 新文本, -1)
}

func W文本_替换2(源文本 string, 替换内容 map[string]string) string {
	局_临时 := 源文本
	for k, v := range 替换内容 {
		局_临时 = strings.Replace(源文本, k, v, -1)
	}
	return 局_临时
}

// 成功返回 位置,失败返回-1
func W文本_寻找(源文本, 要寻找的文本 string) int {
	return strings.Index(源文本, 要寻找的文本)
}
func W文本_取随机IP() string {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10) + 1

	switch x {
	case 1:
		return IP_10进制转IP(取随机数(607649792, 608174079))
	case 2:
		return IP_10进制转IP(取随机数(1038614528, 1039007743))
	case 3:
		return IP_10进制转IP(取随机数(1783627776, 1784676351))
	case 4:
		return IP_10进制转IP(取随机数(2035023872, 1039007743))
	case 5:
		return IP_10进制转IP(取随机数(2078801920, 2079064063))
	case 6:
		return IP_10进制转IP(-1 * 取随机数(1948778497, 1950089216))
	case 7:
		return IP_10进制转IP(-1 * 取随机数(1425014785, 1425539072))
	case 8:
		return IP_10进制转IP(-1 * 取随机数(1235419137, 1236271104))
	case 9:
		return IP_10进制转IP(-1 * 取随机数(768606209, 770113536))
	case 10:
		return IP_10进制转IP(-1 * 取随机数(564133889, 569376768))
	}

	return "0.0.0.0"
}

func W文本_取行数(文本 string) int {
	scanner := bufio.NewScanner(strings.NewReader(文本))
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func W文本_取文本右边2(被查找的文本 string, 欲寻找的文本 string, 起始寻找位置 int, 是否不区分大小写 bool) string {
	起始寻找位置 = func() int {
		if 起始寻找位置 == 0 {
			return len(被查找的文本) + 1
		}
		return 起始寻找位置
	}()

	找到的位置 := strings.LastIndex(被查找的文本, 欲寻找的文本)
	if 是否不区分大小写 {
		找到的位置 = strings.LastIndex(strings.ToLower(被查找的文本), strings.ToLower(欲寻找的文本))
	}

	if 找到的位置 == -1 {
		return ""
	}

	结果文本 := 被查找的文本[len(被查找的文本)-len(欲寻找的文本)-找到的位置+1:]
	return 结果文本
}

func W文本_删除空行(要操作的文本 string) string {
	if strings.HasPrefix(要操作的文本, "\n") {
		要操作的文本 = 要操作的文本[1:]
	}

	if strings.HasSuffix(要操作的文本, "\n") {
		要操作的文本 = 要操作的文本[:len(要操作的文本)-1]
	}

	正则 := regexp.MustCompile(`(?:^\r|\n$)`)
	结果 := 正则.ReplaceAllString(要操作的文本, "")

	return 结果
}
func W文本_逐字分割(原文本 string) []string {
	return strings.Split(原文本, "")
}
func W文本_去重复文本(原文本 string, 分割符 string) string {
	var 局_数组 []string
	var 局_数组1 []string
	var 局_文本 string

	if 分割符 == "" {
		局_数组 = W文本_逐字分割(原文本)
	} else {
		局_数组 = 分割文本(原文本, 分割符)
	}

	for _, 成员 := range 局_数组 {
		if 内部_数组成员是否存在_文本(局_数组1, 成员) == -1 {
			局_数组1 = 加入成员(局_数组1, 成员)
			局_文本 += 成员 + 分割符
		}
	}

	局_文本 = 取文本左边(局_文本, 取文本长度(局_文本)-取文本长度(分割符))
	return 局_文本
}

// 文本取出中间文本_正则方式
func W文本_取出中间文本_批量正则(内容 string, 左边文本 string, 右边文本 string) []string {

	re := regexp.MustCompile(regexp.QuoteMeta(左边文本) + `(.*?)` + regexp.QuoteMeta(右边文本))
	result := re.FindAllStringSubmatch(内容, -1)
	var 局_临时 = make([]string, 0, len(result))
	for i, _ := range result {
		局_临时 = append(局_临时, result[i][1])
	}
	return 局_临时
}

// 文本取出中间文本
func W文本_取出中间文本(内容 string, 左边文本 string, 右边文本 string) string {
	左边位置 := strings.Index(内容, 左边文本)
	if 左边位置 == -1 {
		return ""
	}
	左边位置 = 左边位置 + len(左边文本)
	内容 = string([]byte(内容)[左边位置:])

	var 右边位置 int
	if 右边文本 == "" {
		右边位置 = len(内容)
	} else {
		右边位置 = strings.Index(内容, 右边文本)
		if 右边位置 == -1 {
			return ""
		}
	}
	内容 = string([]byte(内容)[:右边位置])
	return 内容
}

// 获取关键字左边文本 含关键字
func W文本_取文本左边2(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}

	位置 = 位置 + len(关键字) - 1
	内容 = string([]byte(内容)[:位置])
	return 内容
}

// 获取关键字左边文本
func W文本_取文本左边(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}

	位置 = 位置
	内容 = string([]byte(内容)[:位置])
	return 内容
}

// 获取关键字右边文本
func W文本_取文本右边(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}
	内容 = string([]byte(内容)[位置+len(关键字):])
	return 内容
}

// 获取关键字右边文本
func W文本_取文本右边_带关键字(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}
	内容 = string([]byte(内容)[位置+len(关键字):])
	return 内容
}
func W文本_取随机字符串(字符串长度 int) string {
	var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, 字符串长度)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 字符串长度; i++ {
		bytes[i] = strByte[r.Intn(strByteLen-1)]
	}
	if bytes[0] == strByte[strByteLen-1] { //第一位不能是0 防止意外
		bytes[0] = strByte[strByteLen-2]
	}

	return string(bytes)
}

func W文本_取随机字符串_数字(字符串长度 int) string {
	var strByte = []byte("1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, 字符串长度)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 字符串长度; i++ {
		bytes[i] = strByte[r.Intn(strByteLen-1)]
	}
	if bytes[0] == strByte[strByteLen-1] { //第一位不能是0 防止意外
		bytes[0] = strByte[strByteLen-2]
	}

	return string(bytes)
}

// 调用格式： 〈文本型数组〉 分割文本 （文本型 待分割文本，［文本型 用作分割的文本］，［整数型 要返回的子文本数目］） - 系统核心支持库->文本操作
// 英文名称：split
// 将指定文本进行分割，返回分割后的一维文本数组。本命令为初级命令。
// 参数<1>的名称为“待分割文本”，类型为“文本型（text）”。如果参数值是一个长度为零的文本，则返回一个空数组，即没有任何成员的数组。
// 参数<2>的名称为“用作分割的文本”，类型为“文本型（text）”，可以被省略。参数值用于标识子文本边界。如果被省略，则默认使用半角逗号字符作为分隔符。如果是一个长度为零的文本，则返回的数组仅包含一个成员，即完整的“待分割文本”。
// 参数<3>的名称为“要返回的子文本数目”，类型为“整数型（int）”，可以被省略。如果被省略，则默认返回所有的子文本。
//
// 操作系统需求： Windows、Linux
func W文本_分割文本(待分割文本 string, 用作分割的文本 string) []string {
	return strings.Split(待分割文本, 用作分割的文本)
}
func W文本_gbk到utf8(src string) string {
	srcDecoder := mahonia.NewDecoder("gbk")
	desDecoder := mahonia.NewDecoder("utf-8")
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

func W文本_utf8到gbk(src string) string {
	srcDecoder := mahonia.NewDecoder("utf-8")
	desDecoder := mahonia.NewDecoder("gbk")
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

// 调用格式： 〈文本型〉 取文本左边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
// 英文名称：left
// 返回一个文本，其中包含指定文本中从左边算起指定数量的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func W文本_取左边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	if len([]rune(欲取其部分的文本)) < 欲取出字符的数目 {
		欲取出字符的数目 = len(欲取其部分的文本)
	}
	return string([]rune(欲取其部分的文本)[:欲取出字符的数目])
}

//调用格式： 〈文本型〉 取文本右边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
//英文名称：right
//返回一个文本，其中包含指定文本中从右边算起指定数量的字符。本命令为初级命令。
//参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
//参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
//操作系统需求： Windows、Linux

func W文本_取右边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	l := len([]rune(欲取其部分的文本))
	lpos := l - 欲取出字符的数目
	if lpos < 0 {
		lpos = 0
	}
	return string([]rune(欲取其部分的文本)[lpos:l])
}

func W文本_删首尾空(内容 string) string {
	return strings.TrimSpace(内容)
}
func W文本_是否JSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
func W文本_删首空(欲删除空格的文本 string) string {
	return strings.TrimLeft(欲删除空格的文本, " ")
}

//
//调用格式： 〈文本型〉 删尾空 （文本型 欲删除空格的文本） - 系统核心支持库->文本操作
//英文名称：RTrim
//返回一个文本，其中包含被删除了尾部全角或半角空格的指定文本。本命令为初级命令。
//参数<1>的名称为“欲删除空格的文本”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func W文本_删尾空(欲删除空格的文本 string) string {
	return strings.TrimRight(欲删除空格的文本, " ")
}

func W文本_子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string {

	return strings.Replace(欲被替换的文本, 欲被替换的子文本, 用作替换的子文本, -1)
}

func W文本_取随机ip() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}
func W文本_到大写(value string) string {
	return strings.ToUpper(value)
}

func W文本_到小写(value string) string {
	return strings.ToLower(value)
}

// 中文占多个字节但是这里算一个长度
func W文本_取长度(value string) int {
	return utf8.RuneCountInString(value)
}

// 调用格式： 〈文本型〉 字符 （字节型 欲取其字符的字符代码） - 系统核心支持库->文本操作
// 英文名称：chr
// 返回一个文本，其中包含有与指定字符代码相关的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其字符的字符代码”，类型为“字节型（byte）”。
//
// 操作系统需求： Windows、Linux
func W文本_字符(字节型 int8) string {
	return string(byte(字节型))
}

// 查找关键字位置,失败返回-1
func W文本_寻找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.Index(被搜寻的文本, 欲寻找的文本)
}

// 从后往前查找关键字位置,失败返回-1
func W文本_倒找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.LastIndex(被搜寻的文本, 欲寻找的文本)
}

// 调用格式： 〈文本型〉 取空白文本 （整数型 重复次数） - 系统核心支持库->文本操作
// 英文名称：space
// 返回具有指定数目半角空格的文本。本命令为初级命令。
// 参数<1>的名称为“重复次数”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func W文本_取空白(重复次数 int) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + " "
	}
	return str
}

//调用格式： 〈文本型〉 取重复文本 （整数型 重复次数，文本型 待重复文本） - 系统核心支持库->文本操作
//英文名称：string
//返回一个文本，其中包含指定次数的文本重复结果。本命令为初级命令。
//参数<1>的名称为“重复次数”，类型为“整数型（int）”。
//参数<2>的名称为“待重复文本”，类型为“文本型（text）”。该文本将用于建立返回的文本。如果为空，将返回一个空文本。
//
//操作系统需求： Windows、Linux

func W文本_取重复(重复次数 int, 待重复文本 string) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + 待重复文本
	}
	return str
}
