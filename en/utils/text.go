package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/axgle/mahonia"
)

// Text_Contains 判断被搜索文本中是否包含待搜索文本（区分大小写）。
//
// 参数:
//   - searchIn: 被搜索文本
//   - searchFor: 待搜索文本
//
// 返回:
//   - bool: true 表示存在
func Text_Contains(searchIn string, searchFor string) bool {
	return strings.Contains(searchIn, searchFor)
}

// Text_HasPrefix 判断文本是否以指定前缀开头。
//
// 参数:
//   - s: 原始文本
//   - prefix: 前缀
//
// 返回:
//   - bool: true 表示以该前缀开头
func Text_HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// Text_HasSuffix 判断文本是否以指定后缀结尾。
//
// 参数:
//   - s: 原始文本
//   - suffix: 后缀
//
// 返回:
//   - bool: true 表示以该后缀结尾
func Text_HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Text_Index 在源文本中查找指定文本的首次出现位置。
//
// 参数:
//   - source: 被搜索的源文本
//   - target: 待查找的文本
//
// 返回:
//   - int: 首次出现的位置（字节偏移，从 0 开始）；未找到返回 -1
func Text_Index(source string, target string) int {
	return strings.Index(source, target)
}

// Text_LastIndex 在源文本中查找指定文本的最后一次出现位置。
//
// 参数:
//   - source: 被搜索的源文本
//   - target: 待查找的文本
//
// 返回:
//   - int: 最后一次出现的位置（字节偏移，从 0 开始）；未找到返回 -1
func Text_LastIndex(source string, target string) int {
	return strings.LastIndex(source, target)
}

// Text_Count 统计指定子串在源文本中的出现次数。
//
// 参数:
//   - source: 源文本
//   - substr: 要统计的子串
//
// 返回:
//   - int: 出现次数
func Text_Count(source string, substr string) int {
	return strings.Count(source, substr)
}

// Text_Repeat 将指定文本重复连接指定次数。
//
// 参数:
//   - text: 要重复的文本
//   - count: 重复次数
//
// 返回:
//   - string: 重复连接后的文本
func Text_Repeat(text string, count int) string {
	return strings.Repeat(text, count)
}

// Text_Replace 将源文本中所有匹配的旧文本替换为新文本（全部替换）。
//
// 参数:
//   - source: 原始文本
//   - oldStr: 要被替换的子串
//   - newStr: 替换后的子串
//
// 返回:
//   - string: 替换后的文本
func Text_Replace(source, oldStr, newStr string) string {
	return strings.Replace(source, oldStr, newStr, -1)
}

// Text_ReplaceN 将源文本中前 N 个匹配的旧文本替换为新文本。
//
// 参数:
//   - source: 原始文本
//   - oldStr: 要被替换的子串
//   - newStr: 替换后的子串
//   - n: 替换的次数上限
//
// 返回:
//   - string: 替换后的文本
func Text_ReplaceN(source, oldStr, newStr string, n int) string {
	return strings.Replace(source, oldStr, newStr, n)
}

// Text_ReplaceMulti 使用 map 批量替换文本中的多个子串。
// map 的键为旧文本，值为新文本。
// 注意：当前实现每次替换都从原始源文本开始，可能导致后续替换覆盖前面的结果。
//
// 参数:
//   - source: 原始文本
//   - replacements: 替换映射，key=旧文本，value=新文本
//
// 返回:
//   - string: 替换后的文本
func Text_ReplaceMulti(source string, replacements map[string]string) string {
	result := source
	for k, v := range replacements {
		result = strings.Replace(source, k, v, -1)
	}
	return result
}

// Text_ToLower 将文本转换为小写。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 小写文本
func Text_ToLower(text string) string {
	return strings.ToLower(text)
}

// Text_ToUpper 将文本转换为大写。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 大写文本
func Text_ToUpper(text string) string {
	return strings.ToUpper(text)
}

// Text_ToTitle 将文本中每个单词的首字母转为大写。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 首字母大写的文本
func Text_ToTitle(text string) string {
	return strings.ToTitle(text)
}

// Text_Capitalize 将文本的首字母转为大写。
//
// 参数:
//   - text: 英文原始文本
//
// 返回:
//   - string: 首字母大写后的文本
func Text_Capitalize(text string) string {
	return strings.ToUpper(string(text[0])) + text[1:]
}

// Text_Trim 去除文本首尾的空白字符（空格、制表符、换行等）。
// 等同于 strings.TrimSpace。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 去除首尾空白后的文本
func Text_Trim(text string) string {
	return strings.TrimSpace(text)
}

// Text_TrimLeft 去除文本左侧的空白字符。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 去除左侧空白后的文本
func Text_TrimLeft(text string) string {
	return strings.TrimLeft(text, " \t\r\n")
}

// Text_TrimRight 去除文本右侧的空白字符。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 去除右侧空白后的文本
func Text_TrimRight(text string) string {
	return strings.TrimRight(text, " \t\r\n")
}

// Text_TrimSpaceLeft 去除文本左侧的空格字符。
// 仅去除半角空格，不去除制表符和换行符。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 去除左侧空格后的文本
func Text_TrimSpaceLeft(text string) string {
	return strings.TrimLeft(text, " ")
}

// Text_TrimSpaceRight 去除文本右侧的空格字符。
// 仅去除半角空格，不去除制表符和换行符。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 去除右侧空格后的文本
func Text_TrimSpaceRight(text string) string {
	return strings.TrimRight(text, " ")
}

// Text_Split 按指定分隔符将文本分割为字符串数组。
// 等同于 strings.Split。
//
// 参数:
//   - text: 待分割的文本
//   - separator: 用作分割的分隔符
//
// 返回:
//   - []string: 分割后的字符串数组
func Text_Split(text string, separator string) []string {
	return strings.Split(text, separator)
}

// Text_Join 泛型数组合并为文本，用指定连接字符连接。
// 每个元素通过 ToString 转换为字符串后拼接。
//
// 类型参数:
//   - T: 可比较的类型
//
// 参数:
//   - arr: 待合并的泛型数组
//   - separator: 元素之间的连接符
//
// 返回:
//   - string: 合并后的文本
func Text_Join[T comparable](arr []T, separator string) string {
	result := ""
	for _, item := range arr {
		result += ToString(item) + separator
	}
	result = strings.TrimSuffix(result, separator)
	return result
}

// Text_Reverse 将文本内容反转。
// 带有中文时按双字节反转（适用于 GBK 编码），否则按 Unicode 字符反转。
//
// 参数:
//   - text: 待反转的文本
//   - hasChineseChars: true 时按双字节反转，false 时按 Unicode 字符反转
//
// 返回:
//   - string: 反转后的文本
func Text_Reverse(text string, hasChineseChars bool) string {
	if hasChineseChars {
		bytes := []byte(text)
		result := make([]byte, 0, len(bytes))
		for i := len(bytes) - 1; i >= 0; i -= 2 {
			result = append(result, bytes[i-1], bytes[i])
		}
		return string(result)
	}

	reversed := ""
	for _, r := range text {
		reversed = string(r) + reversed
	}
	return reversed
}

// Text_CharCount 获取文本的字符数（非字节数）。
// 中文占多个字节但算一个字符长度，使用 utf8.RuneCountInString 实现。
//
// 参数:
//   - text: 源文本
//
// 返回:
//   - int: 字符数
func Text_CharCount(text string) int {
	return utf8.RuneCountInString(text)
}

// Text_Left 从文本左侧截取指定字符数的子串。
// 使用 rune 截取，中文安全，每个中文算一个字符。
//
// 参数:
//   - text: 源文本
//   - count: 截取的字符数
//
// 返回:
//   - string: 左侧指定字符数的文本；请求数超过长度时返回原文本
func Text_Left(text string, count int) string {
	runes := []rune(text)
	if len(runes) < count {
		count = len(runes)
	}
	return string(runes[:count])
}

// Text_Right 从文本右侧截取指定字符数的子串。
// 使用 rune 截取，中文安全，每个中文算一个字符。
//
// 参数:
//   - text: 源文本
//   - count: 截取的字符数
//
// 返回:
//   - string: 右侧指定字符数的文本；请求数超过长度时返回原文本
func Text_Right(text string, count int) string {
	runes := []rune(text)
	lpos := len(runes) - count
	if lpos < 0 {
		lpos = 0
	}
	return string(runes[lpos:])
}

// Text_Substring 获取文本中指定范围的子文本。
// 使用 rune 截取，中文安全。
//
// 参数:
//   - text: 源文本
//   - start: 起始位置（从 0 开始）
//   - length: 取出长度；-1 表示取到末尾
//
// 返回:
//   - string: 指定范围的子文本
func Text_Substring(text string, start int, length int) string {
	runes := []rune(text)
	strLen := len(runes)
	if start >= strLen {
		return ""
	}
	if length == -1 || start+length > strLen {
		return string(runes[start:])
	}
	return string(runes[start : start+length])
}

// Text_Mid 按左右标识取出中间文本（首次匹配）。
// 从左向右搜索，返回第一个匹配结果。
//
// 参数:
//   - content: 源文本
//   - leftBound: 中间文本左侧的标记
//   - rightBound: 中间文本右侧的标记，为空时取到文本末尾
//
// 返回:
//   - string: 提取的中间文本；未找到标记时返回空串
func Text_Mid(content string, leftBound string, rightBound string) string {
	leftPos := strings.Index(content, leftBound)
	if leftPos == -1 {
		return ""
	}
	leftPos = leftPos + len(leftBound)
	content = string([]byte(content)[leftPos:])

	var rightPos int
	if rightBound == "" {
		rightPos = len(content)
	} else {
		rightPos = strings.Index(content, rightBound)
		if rightPos == -1 {
			return ""
		}
	}
	content = string([]byte(content)[:rightPos])
	return content
}

// Text_MidAll 使用正则表达式批量提取左右标记之间的所有中间文本。
// 左右标记会自动进行正则转义，中间部分使用非贪婪匹配。
//
// 参数:
//   - content: 源文本
//   - leftBound: 中间文本左侧的标记
//   - rightBound: 中间文本右侧的标记
//
// 返回:
//   - []string: 所有匹配的中间文本数组
func Text_MidAll(content string, leftBound string, rightBound string) []string {
	re := regexp.MustCompile(regexp.QuoteMeta(leftBound) + `(.*?)` + regexp.QuoteMeta(rightBound))
	match := re.FindAllStringSubmatch(content, -1)
	result := make([]string, 0, len(match))
	for i := range match {
		result = append(result, match[i][1])
	}
	return result
}

// Text_LeftBefore 获取关键字左侧的文本（不含关键字）。
// 按字节截取，注意中文占多个字节。
//
// 参数:
//   - content: 源文本
//   - keyword: 分隔关键字
//
// 返回:
//   - string: 关键字左侧的文本；未找到返回空串
func Text_LeftBefore(content string, keyword string) string {
	pos := strings.Index(content, keyword)
	if pos == -1 {
		return ""
	}
	content = string([]byte(content)[:pos])
	return content
}

// Text_RightAfter 获取关键字右侧的文本（不含关键字）。
// 按字节截取，注意中文占多个字节。
//
// 参数:
//   - content: 源文本
//   - keyword: 分隔关键字
//
// 返回:
//   - string: 关键字右侧的文本；未找到返回空串
func Text_RightAfter(content string, keyword string) string {
	pos := strings.Index(content, keyword)
	if pos == -1 {
		return ""
	}
	content = string([]byte(content)[pos+len(keyword):])
	return content
}

// Text_LeftToKeyword 获取关键字左侧的文本（含关键字本身）。
//
// 参数:
//   - content: 源文本
//   - keyword: 分隔关键字
//
// 返回:
//   - string: 关键字及其左侧的文本；未找到返回空串
func Text_LeftToKeyword(content string, keyword string) string {
	pos := strings.Index(content, keyword)
	if pos == -1 {
		return ""
	}
	pos = pos + len(keyword)
	content = string([]byte(content)[:pos])
	return content
}

// Text_LineCount 统计文本的行数。
// 按换行符统计，空文本返回 0。
//
// 参数:
//   - text: 待统计的文本
//
// 返回:
//   - int: 行数
func Text_LineCount(text string) int {
	scanner := bufio.NewScanner(strings.NewReader(text))
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

// Text_LineAt 获取文本中指定行号的内容（行号从 1 开始）。
//
// 参数:
//   - text: 原始多行文本
//   - lineNum: 行号（从 1 开始）
//
// 返回:
//   - string: 指定行的文本；行号无效时返回空串
func Text_LineAt(text string, lineNum int) string {
	lines := strings.Split(text, "\n")
	if lineNum <= 0 {
		return ""
	}
	if lineNum > len(lines) {
		return ""
	}
	return lines[lineNum-1]
}

// Text_RemoveBlankLines 删除文本中的空行。
// 移除文本首尾的换行符，以及每行开头的 \r 和末尾的 \n。
//
// 参数:
//   - text: 原始文本
//
// 返回:
//   - string: 删除空行后的文本
func Text_RemoveBlankLines(text string) string {
	if strings.HasPrefix(text, "\n") {
		text = text[1:]
	}
	if strings.HasSuffix(text, "\n") {
		text = text[:len(text)-1]
	}
	reg := regexp.MustCompile(`(?:^\r|\n$)`)
	result := reg.ReplaceAllString(text, "")
	return result
}

// Text_ToRuneSlice 将文本逐字符拆分为字符串数组。
// 每个字符（含中文）作为数组的一个元素。
//
// 参数:
//   - text: 待拆分的文本
//
// 返回:
//   - []string: 字符数组
func Text_ToRuneSlice(text string) []string {
	return strings.Split(text, "")
}

// Text_Deduplicate 去除文本中重复的元素。
// 按指定分隔符分割后去重，保持元素首次出现的顺序。
//
// 参数:
//   - text: 源文本
//   - separator: 元素之间的分隔符，为空时按字符去重
//
// 返回:
//   - string: 去重后用分隔符连接的文本
func Text_Deduplicate(text string, separator string) string {
	var arr []string
	var arrResult []string
	var resultText string

	if separator == "" {
		arr = Text_ToRuneSlice(text)
	} else {
		arr = Text_Split(text, separator)
	}

	for _, member := range arr {
		if Array_FindText(arrResult, member) == -1 {
			arrResult = AppendMember(arrResult, member)
			resultText += member + separator
		}
	}

	resultText = Text_LeftTextBytes(resultText, TextLen(resultText)-TextLen(separator))
	return resultText
}

// Text_LeftTextBytes 取文本左侧指定字节数的子串。
// 注意：按字节截取，中文可能被截断。中文安全截取请用 Text_Left。
//
// 参数:
//   - text: 源文本
//   - n: 截取的字节数
//
// 返回:
//   - string: 左侧 n 个字节的文本；n 超过长度时返回原文本
func Text_LeftTextBytes(text string, n int) string {
	if n >= len(text) {
		return text
	}
	return text[:n]
}

// Text_RandomString 生成指定长度的随机字母数字字符串。
// 字符集包含大小写字母和数字，首位不会为 '0'。
//
// 参数:
//   - length: 生成的字符串长度
//
// 返回:
//   - string: 随机字符串
func Text_RandomString(length int) string {
	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var charsetLen = len(charset)
	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = charset[r.Intn(charsetLen-1)]
	}
	if bytes[0] == charset[charsetLen-1] {
		bytes[0] = charset[charsetLen-2]
	}
	return string(bytes)
}

// Text_RandomDigits 生成指定长度的随机数字字符串。
// 字符集仅包含数字 1-9 和 0，首位不会为 '0'。
//
// 参数:
//   - length: 生成的字符串长度
//
// 返回:
//   - string: 随机数字字符串
func Text_RandomDigits(length int) string {
	var charset = []byte("1234567890")
	var charsetLen = len(charset)
	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		bytes[i] = charset[r.Intn(charsetLen-1)]
	}
	if bytes[0] == charset[charsetLen-1] {
		bytes[0] = charset[charsetLen-2]
	}
	return string(bytes)
}

// Text_RandomIP 生成一个随机的 IP 地址。
// 每段为 0-254 之间的随机数，可能生成保留地址。
//
// 返回:
//   - string: 随机 IP 地址，如 "192.168.1.1"
func Text_RandomIP() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

// Text_RandomPublicIP 生成一个随机的公网 IP 地址。
// 从预定义的 IP 段中随机选取，确保生成的 IP 属于公网地址范围。
//
// 返回:
//   - string: 随机 IP 地址，如 "203.0.113.42"；生成失败返回 "0.0.0.0"
func Text_RandomPublicIP() string {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10) + 1

	switch x {
	case 1:
		return IP_DecToIP(RandomNum(607649792, 608174079))
	case 2:
		return IP_DecToIP(RandomNum(1038614528, 1039007743))
	case 3:
		return IP_DecToIP(RandomNum(1783627776, 1784676351))
	case 4:
		return IP_DecToIP(RandomNum(2035023872, 1039007743))
	case 5:
		return IP_DecToIP(RandomNum(2078801920, 2079064063))
	case 6:
		return IP_DecToIP(-1 * RandomNum(1948778497, 1950089216))
	case 7:
		return IP_DecToIP(-1 * RandomNum(1425014785, 1425539072))
	case 8:
		return IP_DecToIP(-1 * RandomNum(1235419137, 1236271104))
	case 9:
		return IP_DecToIP(-1 * RandomNum(768606209, 770113536))
	case 10:
		return IP_DecToIP(-1 * RandomNum(564133889, 569376768))
	}

	return "0.0.0.0"
}

// Text_IsJSON 检查字符串是否为有效的 JSON 对象格式。
// 仅检查是否为 {"key": "value"} 格式的 JSON 对象，不检查 JSON 数组。
//
// 参数:
//   - s: 待检查的字符串
//
// 返回:
//   - bool: 是有效的 JSON 对象返回 true，否则返回 false
func Text_IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

// Text_GBKToUTF8 将 GBK 编码的字符串转换为 UTF-8 编码。
// 使用 mahonia 库进行编码转换。
//
// 参数:
//   - src: GBK 编码的源字符串
//
// 返回:
//   - string: UTF-8 编码的字符串
func Text_GBKToUTF8(src string) string {
	srcDecoder := mahonia.NewDecoder("gbk")
	desDecoder := mahonia.NewDecoder("utf-8")
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

// Text_UTF8ToGBK 将 UTF-8 编码的字符串转换为 GBK 编码。
// 使用 mahonia 库进行编码转换。
//
// 参数:
//   - src: UTF-8 编码的源字符串
//
// 返回:
//   - string: GBK 编码的字符串
func Text_UTF8ToGBK(src string) string {
	srcDecoder := mahonia.NewDecoder("utf-8")
	desDecoder := mahonia.NewDecoder("gbk")
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

// Text_CountOccurrences 统计指定文本在源文本中出现的次数。
//
// 参数:
//   - searchIn: 源文本
//   - searchFor: 要统计的文本
//
// 返回:
//   - int: 出现次数
func Text_CountOccurrences(searchIn string, searchFor string) int {
	return strings.Count(searchIn, searchFor)
}

// Text_ToByte 将字节值转换为对应的单字符字符串。
// 等同于 string(byte(code))。
//
// 参数:
//   - code: ASCII 字符代码（-128 到 127）
//
// 返回:
//   - string: 单字符字符串
func Text_ToByte(code int8) string {
	return string(byte(code))
}