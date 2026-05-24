package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// W文本_是否包含关键字 检查内容中是否包含指定关键字。
// 关键字为空时直接返回 true。
//
// 参数:
//   - 内容: 被搜索的文本
//   - 关键字: 要查找的关键字
//
// 返回:
//   - bool: 包含返回 true，否则返回 false
func W文本_是否包含关键字(内容, 关键字 string) bool {
	return strings.Contains(内容, 关键字)
}

// W文本_是否存在 检查内容中是否包含指定关键字。
// 功能与 W文本_是否包含关键字 相同，为使用习惯提供别名。
//
// 参数:
//   - 内容: 被搜索的文本
//   - 关键字: 要查找的关键字
//
// 返回:
//   - bool: 包含返回 true，否则返回 false
func W文本_是否存在(内容, 关键字 string) bool {
	return strings.Contains(内容, 关键字)
}

// W文本_是否存在_任意 检查内容中是否包含关键字数组中的任意一个。
// 只要有一个关键字匹配即返回 true。
//
// 参数:
//   - 内容: 被搜索的文本
//   - 关键字: 要查找的关键字数组
//
// 返回:
//   - bool: 任意一个关键字存在返回 true，全部不存在返回 false
func W文本_是否存在_任意(内容 string, 关键字 []string) bool {
	for _, v := range 关键字 {
		if strings.Contains(内容, v) {
			return true
		}
	}
	return false
}

// W文本_是否存在_同时 检查内容中是否同时包含关键字数组中的所有关键字。
// 所有关键字都必须存在才返回 true。
//
// 参数:
//   - 内容: 被搜索的文本
//   - 关键字: 要查找的关键字数组
//
// 返回:
//   - bool: 所有关键字都存在返回 true，任一不存在返回 false
func W文本_是否存在_同时(内容 string, 关键字 []string) bool {
	for _, v := range 关键字 {
		if !strings.Contains(内容, v) {
			return false
		}
	}
	return true
}

// W文本_是否为英数字母 检查字符串是否仅包含英文字母和数字。
//
// 参数:
//   - s: 待检查的字符串
//
// 返回:
//   - bool: 仅包含英文字母和数字返回 true，否则返回 false
func W文本_是否为英数字母(s string) bool {
	pattern := "^[A-Za-z0-9]+$"
	match, _ := regexp.MatchString(pattern, s)
	return match
}

// W文本_是否为字母 检查字符串是否仅包含字母字符（含中文等 Unicode 字母）。
//
// 参数:
//   - s: 待检查的字符串
//
// 返回:
//   - bool: 仅包含字母返回 true，否则返回 false
func W文本_是否为字母(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// W文本_是否为数字 检查字符串是否仅包含数字字符（含 Unicode 数字）。
//
// 参数:
//   - s: 待检查的字符串
//
// 返回:
//   - bool: 仅包含数字返回 true，否则返回 false
func W文本_是否为数字(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// W文本_倒取出中间文本 从文本末尾开始搜索，提取左右标记之间的中间文本。
// 与 W文本_取出中间文本 不同，此函数从右向左搜索标记位置。
// 例如："0012345" 中，右边文本为 "4"，左边文本为 "0"，则返回 "123"。
//
// 参数:
//   - 欲取全文本: 完整的源文本
//   - 右边文本: 中间文本右侧的标记文本
//   - 左边文本: 中间文本左侧的标记文本
//   - 倒数搜寻位置: 搜索右边文本的起始位置（从右往左），0 或负数表示从末尾开始
//   - 是否不区分大小写: true 时不区分大小写搜索
//
// 返回:
//   - string: 提取的中间文本；未找到标记时返回空串
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

// W文本_取文本所在行 查找指定文本在源文本中的行号。
// 按换行符 \n 分割源文本，返回第一次匹配的行号。
//
// 参数:
//   - 源文本: 被搜索的多行文本
//   - 欲查找的文本: 要查找的文本
//   - 是否区分大小写: 当前未使用，始终区分大小写
//
// 返回:
//   - int: 行号（从 0 开始）；未找到返回 0
func W文本_取文本所在行(源文本 string, 欲查找的文本 string, 是否区分大小写 bool) int {
	局文本 := strings.Split(源文本, "\n")
	for 局计次 := 0; 局计次 < len(局文本); 局计次++ {
		if strings.Index(局文本[局计次], 欲查找的文本) != -1 {
			return 局计次
		}
	}
	return 0
}

// W文本_删除指定文本行 删除源文本中指定行号的整行文本。
// 行号从 0 开始计数。
//
// 参数:
//   - 源文本: 多行文本
//   - 行数: 要删除的行号（从 0 开始）
//
// 返回:
//   - string: 删除指定行后的文本
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

// W文本_取随机范围数字 生成指定范围内的随机数字字符串。
// 可选择只生成奇数或偶数。
//
// 参数:
//   - 起始数: 最小值（包含）
//   - 结束数: 最大值（包含）
//   - 单双选择: 0=不限, 1=只取奇数, 2=只取偶数
//
// 返回:
//   - string: 随机数字字符串
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

// W文本_取指定变量文本行 获取多行文本中指定行号的内容。
// 行号从 1 开始计数。
//
// 参数:
//   - 文本: 多行文本
//   - 行号: 要获取的行号（从 1 开始）
//
// 返回:
//   - string: 指定行的文本；行号越界时返回空串
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

// W文本_颠倒 将文本内容反转。
// 带有中文时按双字节反转（适用于 GBK 编码），否则按 Unicode 字符反转。
//
// 参数:
//   - 欲转换文本: 待反转的文本
//   - 带有中文: true 时按双字节反转，false 时按 Unicode 字符反转
//
// 返回:
//   - string: 反转后的文本
func W文本_颠倒(欲转换文本 string, 带有中文 bool) string {
	if 带有中文 {
		字节集 := []byte(欲转换文本)
		局_结果 := make([]byte, 0, len(字节集))
		for i := len(字节集) - 1; i >= 0; i -= 2 {
			局_结果 = append(局_结果, 字节集[i-1], 字节集[i])
		}
		return string(局_结果)
	}

	倒序内容 := ""
	for _, r := range 欲转换文本 {
		倒序内容 = string(r) + 倒序内容
	}
	return 倒序内容
}

// W文本_取出现次数 统计指定文本在源文本中出现的次数。
//
// 参数:
//   - 被搜索文本: 源文本
//   - 欲搜索文本: 要统计的文本
//
// 返回:
//   - int: 出现次数
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

// W文本_首字母改大写 将文本的首字母转换为大写。
//
// 参数:
//   - 英文文本: 待转换的文本
//
// 返回:
//   - string: 首字母大写后的文本
func W文本_首字母改大写(英文文本 string) string {
	return strings.ToUpper(string(英文文本[0])) + 英文文本[1:]
}

// W文本_替换 将源文本中所有匹配的旧文本替换为新文本。
// 全部替换，等同于 strings.Replace 的 -1 模式。
//
// 参数:
//   - 源文本: 原始文本
//   - 旧文本: 要被替换的子串
//   - 新文本: 替换后的子串
//
// 返回:
//   - string: 替换后的文本
func W文本_替换(源文本, 旧文本, 新文本 string) string {
	return strings.Replace(源文本, 旧文本, 新文本, -1)
}

// W文本_替换2 使用 map 批量替换文本中的多个子串。
// map 的键为旧文本，值为新文本。
// 注意：当前实现每次替换都从原始源文本开始，可能导致后续替换覆盖前面的结果。
//
// 参数:
//   - 源文本: 原始文本
//   - 替换内容: 替换映射，key=旧文本，value=新文本
//
// 返回:
//   - string: 替换后的文本
func W文本_替换2(源文本 string, 替换内容 map[string]string) string {
	局_临时 := 源文本
	for k, v := range 替换内容 {
		局_临时 = strings.Replace(源文本, k, v, -1)
	}
	return 局_临时
}

// W文本_寻找 在源文本中查找指定文本的位置。
//
// 参数:
//   - 源文本: 被搜索的文本
//   - 要寻找的文本: 要查找的文本
//
// 返回:
//   - int: 找到的位置（字节偏移，从 0 开始）；未找到返回 -1
func W文本_寻找(源文本, 要寻找的文本 string) int {
	return strings.Index(源文本, 要寻找的文本)
}

// W文本_取随机IP 生成一个随机的公网 IP 地址。
// 从预定义的 IP 段中随机选取，确保生成的 IP 属于公网地址范围。
//
// 返回:
//   - string: 随机 IP 地址，如 "203.0.113.42"；生成失败返回 "0.0.0.0"
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

// W文本_取行数 统计文本的行数。
// 按换行符统计，空文本返回 0。
//
// 参数:
//   - 文本: 待统计的文本
//
// 返回:
//   - int: 行数
func W文本_取行数(文本 string) int {
	scanner := bufio.NewScanner(strings.NewReader(文本))
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

// W文本_取文本右边2 从文本右侧截取指定关键字之后的内容。
// 支持从指定位置开始搜索，可选择是否区分大小写。
//
// 参数:
//   - 被查找的文本: 源文本
//   - 欲寻找的文本: 分隔关键字
//   - 起始寻找位置: 搜索起始位置，0 表示从末尾开始
//   - 是否不区分大小写: true 时不区分大小写
//
// 返回:
//   - string: 关键字右侧的文本；未找到返回空串
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

// W文本_删除空行 删除文本中的空行。
// 移除文本首尾的换行符，以及每行开头的 \r 和末尾的 \n。
//
// 参数:
//   - 要操作的文本: 源文本
//
// 返回:
//   - string: 删除空行后的文本
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

// W文本_逐字分割 将文本逐字符拆分为字符串数组。
// 每个字符（含中文）作为数组的一个元素。
//
// 参数:
//   - 原文本: 待拆分的文本
//
// 返回:
//   - []string: 字符数组
func W文本_逐字分割(原文本 string) []string {
	return strings.Split(原文本, "")
}

// W文本_去重复文本 去除文本中重复的元素。
// 按指定分隔符分割后去重，保持元素首次出现的顺序。
//
// 参数:
//   - 原文本: 源文本
//   - 分割符: 元素之间的分隔符，为空时按字符去重
//
// 返回:
//   - string: 去重后用分隔符连接的文本
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

// W文本_取出中间文本_批量正则 使用正则表达式批量提取左右标记之间的所有中间文本。
// 左右标记会自动进行正则转义，中间部分使用非贪婪匹配。
//
// 参数:
//   - 内容: 源文本
//   - 左边文本: 中间文本左侧的标记
//   - 右边文本: 中间文本右侧的标记
//
// 返回:
//   - []string: 所有匹配的中间文本数组
func W文本_取出中间文本_批量正则(内容 string, 左边文本 string, 右边文本 string) []string {

	re := regexp.MustCompile(regexp.QuoteMeta(左边文本) + `(.*?)` + regexp.QuoteMeta(右边文本))
	result := re.FindAllStringSubmatch(内容, -1)
	var 局_临时 = make([]string, 0, len(result))
	for i, _ := range result {
		局_临时 = append(局_临时, result[i][1])
	}
	return 局_临时
}

// W文本_取出中间文本 从文本中提取左右标记之间的中间文本（首次匹配）。
// 从左向右搜索，返回第一个匹配结果。
//
// 参数:
//   - 内容: 源文本
//   - 左边文本: 中间文本左侧的标记
//   - 右边文本: 中间文本右侧的标记，为空时取到文本末尾
//
// 返回:
//   - string: 提取的中间文本；未找到标记时返回空串
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

// W文本_取文本左边2 获取关键字左侧的文本（含关键字本身）。
//
// 参数:
//   - 内容: 源文本
//   - 关键字: 分隔关键字
//
// 返回:
//   - string: 关键字及其左侧的文本；未找到返回空串
func W文本_取文本左边2(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}

	位置 = 位置 + len(关键字) - 1
	内容 = string([]byte(内容)[:位置])
	return 内容
}

// W文本_取文本左边 获取关键字左侧的文本（不含关键字）。
// 按字节截取，注意中文占多个字节。
//
// 参数:
//   - 内容: 源文本
//   - 关键字: 分隔关键字
//
// 返回:
//   - string: 关键字左侧的文本；未找到返回空串
func W文本_取文本左边(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}

	内容 = string([]byte(内容)[:位置])
	return 内容
}

// W文本_取文本右边 获取关键字右侧的文本（不含关键字）。
// 按字节截取，注意中文占多个字节。
//
// 参数:
//   - 内容: 源文本
//   - 关键字: 分隔关键字
//
// 返回:
//   - string: 关键字右侧的文本；未找到返回空串
func W文本_取文本右边(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}
	内容 = string([]byte(内容)[位置+len(关键字):])
	return 内容
}

// W文本_取文本右边_带关键字 获取关键字右侧的文本（不含关键字）。
// 功能与 W文本_取文本右边 相同，为使用习惯提供别名。
//
// 参数:
//   - 内容: 源文本
//   - 关键字: 分隔关键字
//
// 返回:
//   - string: 关键字右侧的文本；未找到返回空串
func W文本_取文本右边_带关键字(内容 string, 关键字 string) string {
	位置 := strings.Index(内容, 关键字)
	if 位置 == -1 {
		return ""
	}
	内容 = string([]byte(内容)[位置+len(关键字):])
	return 内容
}

// W文本_取随机字符串 生成指定长度的随机字符串。
// 字符集包含大小写字母和数字，首位不会为 '0'。
//
// 参数:
//   - 字符串长度: 生成的字符串长度
//
// 返回:
//   - string: 随机字符串
func W文本_取随机字符串(字符串长度 int) string {
	var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, 字符串长度)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 字符串长度; i++ {
		bytes[i] = strByte[r.Intn(strByteLen-1)]
	}
	if bytes[0] == strByte[strByteLen-1] {
		bytes[0] = strByte[strByteLen-2]
	}

	return string(bytes)
}

// W文本_取随机字符串_数字 生成指定长度的随机数字字符串。
// 字符集仅包含数字 1-9 和 0，首位不会为 '0'。
//
// 参数:
//   - 字符串长度: 生成的字符串长度
//
// 返回:
//   - string: 随机数字字符串
func W文本_取随机字符串_数字(字符串长度 int) string {
	var strByte = []byte("1234567890")
	var strByteLen = len(strByte)
	bytes := make([]byte, 字符串长度)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 字符串长度; i++ {
		bytes[i] = strByte[r.Intn(strByteLen-1)]
	}
	if bytes[0] == strByte[strByteLen-1] {
		bytes[0] = strByte[strByteLen-2]
	}

	return string(bytes)
}

// W文本_分割文本 按指定分隔符将文本分割为字符串数组。
// 等同于 strings.Split。
//
// 参数:
//   - 待分割文本: 待分割的文本
//   - 用作分割的文本: 分隔符
//
// 返回:
//   - []string: 分割后的字符串数组
func W文本_分割文本(待分割文本 string, 用作分割的文本 string) []string {
	return strings.Split(待分割文本, 用作分割的文本)
}

// W文本_gbk到utf8 将 GBK 编码的字符串转换为 UTF-8 编码。
// 使用 mahonia 库进行编码转换。
//
// 参数:
//   - src: GBK 编码的源字符串
//
// 返回:
//   - string: UTF-8 编码的字符串
func W文本_gbk到utf8(src string) string {
	srcDecoder := mahonia.NewDecoder("gbk")
	desDecoder := mahonia.NewDecoder("utf-8")
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

// W文本_utf8到gbk 将 UTF-8 编码的字符串转换为 GBK 编码。
// 使用 mahonia 库进行编码转换。
//
// 参数:
//   - src: UTF-8 编码的源字符串
//
// 返回:
//   - string: GBK 编码的字符串
func W文本_utf8到gbk(src string) string {
	srcDecoder := mahonia.NewDecoder("utf-8")
	desDecoder := mahonia.NewDecoder("gbk")
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

// W文本_取左边 从文本左侧截取指定字符数的子串。
// 使用 rune 截取，中文安全，每个中文算一个字符。
//
// 参数:
//   - 欲取其部分的文本: 源文本
//   - 欲取出字符的数目: 截取的字符数
//
// 返回:
//   - string: 左侧指定字符数的文本；请求数超过长度时返回原文本
func W文本_取左边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	if len([]rune(欲取其部分的文本)) < 欲取出字符的数目 {
		欲取出字符的数目 = len([]rune(欲取其部分的文本))
	}
	return string([]rune(欲取其部分的文本)[:欲取出字符的数目])
}

// W文本_取右边 从文本右侧截取指定字符数的子串。
// 使用 rune 截取，中文安全，每个中文算一个字符。
//
// 参数:
//   - 欲取其部分的文本: 源文本
//   - 欲取出字符的数目: 截取的字符数
//
// 返回:
//   - string: 右侧指定字符数的文本；请求数超过长度时返回原文本
func W文本_取右边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	l := len([]rune(欲取其部分的文本))
	lpos := l - 欲取出字符的数目
	if lpos < 0 {
		lpos = 0
	}
	return string([]rune(欲取其部分的文本)[lpos:l])
}

// W文本_删首尾空 去除文本首尾的空白字符（空格、制表符、换行等）。
// 等同于 strings.TrimSpace。
//
// 参数:
//   - 内容: 源文本
//
// 返回:
//   - string: 去除首尾空白后的文本
func W文本_删首尾空(内容 string) string {
	return strings.TrimSpace(内容)
}

// W文本_是否JSON 检查字符串是否为有效的 JSON 对象格式。
// 仅检查是否为 {"key": "value"} 格式的 JSON 对象，不检查 JSON 数组。
//
// 参数:
//   - s: 待检查的字符串
//
// 返回:
//   - bool: 是有效的 JSON 对象返回 true，否则返回 false
func W文本_是否JSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

// W文本_删首空 去除文本左侧的空格字符。
// 仅去除半角空格，不去除制表符和换行符。
//
// 参数:
//   - 欲删除空格的文本: 源文本
//
// 返回:
//   - string: 去除左侧空格后的文本
func W文本_删首空(欲删除空格的文本 string) string {
	return strings.TrimLeft(欲删除空格的文本, " ")
}

// W文本_删尾空 去除文本右侧的空格字符。
// 仅去除半角空格，不去除制表符和换行符。
//
// 参数:
//   - 欲删除空格的文本: 源文本
//
// 返回:
//   - string: 去除右侧空格后的文本
func W文本_删尾空(欲删除空格的文本 string) string {
	return strings.TrimRight(欲删除空格的文本, " ")
}

// W文本_子文本替换 将源文本中所有匹配的子文本替换为新子文本。
// 全部替换，等同于 strings.Replace 的 -1 模式。
//
// 参数:
//   - 欲被替换的文本: 原始文本
//   - 欲被替换的子文本: 要被替换的子串
//   - 用作替换的子文本: 替换后的子串
//
// 返回:
//   - string: 替换后的文本
func W文本_子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string {

	return strings.Replace(欲被替换的文本, 欲被替换的子文本, 用作替换的子文本, -1)
}

// W文本_取随机ip 生成一个随机的 IP 地址。
// 每段为 0-254 之间的随机数，可能生成保留地址。
//
// 返回:
//   - string: 随机 IP 地址，如 "192.168.1.1"
func W文本_取随机ip() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}

// W文本_到大写 将文本转换为大写。
//
// 参数:
//   - value: 源文本
//
// 返回:
//   - string: 大写文本
func W文本_到大写(value string) string {
	return strings.ToUpper(value)
}

// W文本_到小写 将文本转换为小写。
//
// 参数:
//   - value: 源文本
//
// 返回:
//   - string: 小写文本
func W文本_到小写(value string) string {
	return strings.ToLower(value)
}

// W文本_取长度 获取文本的字符数（非字节数）。
// 中文占多个字节但算一个字符长度，使用 utf8.RuneCountInString 实现。
//
// 参数:
//   - value: 源文本
//
// 返回:
//   - int: 字符数
func W文本_取长度(value string) int {
	return utf8.RuneCountInString(value)
}

// W文本_字符 将字节值转换为对应的单字符字符串。
// 等同于 string(byte(字节型))。
//
// 参数:
//   - 字节型: ASCII 字符代码（-128 到 127）
//
// 返回:
//   - string: 对应的单字符字符串
func W文本_字符(字节型 int8) string {
	return string(byte(字节型))
}

// W文本_寻找文本 在源文本中查找指定文本的位置。
// 从左向右搜索，返回第一次出现的位置。
//
// 参数:
//   - 被搜寻的文本: 被搜索的文本
//   - 欲寻找的文本: 要查找的文本
//
// 返回:
//   - int: 找到的位置（字节偏移，从 0 开始）；未找到返回 -1
func W文本_寻找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.Index(被搜寻的文本, 欲寻找的文本)
}

// W文本_倒找文本 从后往前查找指定文本的位置。
// 返回最后一次出现的位置。
//
// 参数:
//   - 被搜寻的文本: 被搜索的文本
//   - 欲寻找的文本: 要查找的文本
//
// 返回:
//   - int: 找到的位置（字节偏移，从 0 开始）；未找到返回 -1
func W文本_倒找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.LastIndex(被搜寻的文本, 欲寻找的文本)
}

// W文本_取空白 生成指定数量的半角空格字符串。
//
// 参数:
//   - 重复次数: 空格的数量
//
// 返回:
//   - string: 由指定数量空格组成的字符串
func W文本_取空白(重复次数 int) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + " "
	}
	return str
}

// W文本_取重复 生成指定次数重复的文本。
//
// 参数:
//   - 重复次数: 重复的次数
//   - 待重复文本: 要重复的文本
//
// 返回:
//   - string: 重复后的文本；待重复文本为空时返回空串
func W文本_取重复(重复次数 int, 待重复文本 string) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + 待重复文本
	}
	return str
}

// W文本_取随机数字数组 生成指定数量的不重复随机数字字符串数组。
// 数字范围在 [最小值, 最大值) 之间。
//
// 参数:
//   - 最小值: 随机数的最小值（包含）
//   - 最大值: 随机数的上限（不包含）
//   - 数量: 要生成的数字数量
//
// 返回:
//   - []string: 不重复的随机数字字符串数组
func W文本_取随机数字数组(最小值, 最大值 int, 数量 int) []string {
	局_数组 := make([]string, 0, 数量)
	局_已存在 := make(map[string]bool, 数量)

	for len(局_数组) < 数量 {
		局_随机 := rand.Intn(最大值)
		if 局_随机 < 最小值 {
			continue
		}
		局_随机2 := strconv.Itoa(局_随机)
		if _, ok := 局_已存在[局_随机2]; !ok {
			局_已存在[局_随机2] = true
			局_数组 = append(局_数组, 局_随机2)
		}
	}
	return 局_数组
}

// W文本_去除敏感信息 对文本进行脱敏处理，将中间一半的字符替换为 *。
// 长度小于等于 2 时，仅保留第一个字符；长度大于 2 时，中间一半替换为 *。
// 使用 rune 处理，中文安全。
//
// 参数:
//   - 内容: 待脱敏的文本
//
// 返回:
//   - string: 脱敏后的文本，如 "张**"、"138****1234"
func W文本_去除敏感信息(内容 string) string {
	if len(内容) == 0 {
		return 内容
	}

	runes := []rune(内容)
	length := len(runes)

	if length <= 2 {
		if length == 1 {
			return "*"
		}
		return string(runes[0]) + "*"
	}

	replaceCount := length / 2
	startIndex := (length - replaceCount) / 2

	for i := 0; i < replaceCount; i++ {
		runes[startIndex+i] = '*'
	}

	return string(runes)
}
