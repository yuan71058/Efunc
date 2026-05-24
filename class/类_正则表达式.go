package class

import "regexp"

// L_正则表达式 正则表达式封装类，提供正则匹配和子表达式提取功能。
// 创建时编译正则表达式并执行匹配，后续可通过方法获取匹配结果。
//
// 使用示例:
//
//	re, ok := New正则表达式类(`(\d+)-(\d+)`, "abc 123-456 def")
//	if ok {
//	    count := re.Q取匹配数量()
//	    text := re.Q取匹配文本(0)
//	    sub := re.Q取子匹配文本(0, 1)
//	}
type L_正则表达式 struct {
	r              *regexp.Regexp
	res            [][]string
	Count          int
	SubmatchCount2 int
}

// New正则表达式类 创建并初始化正则表达式对象。
// 编译正则表达式文本，并在被搜索文本中执行匹配。
//
// 参数:
//   - 正则表达式文本: 正则表达式模式字符串
//   - 被搜索的文本: 要执行匹配的目标文本
//
// 返回:
//   - *L_正则表达式: 正则表达式对象
//   - bool: true 表示至少有一个匹配，false 表示无匹配
func New正则表达式类(正则表达式文本 string, 被搜索的文本 string) (*L_正则表达式, bool) {
	t := new(L_正则表达式)
	b := t.E创建(正则表达式文本, 被搜索的文本)
	return t, b
}

// E创建 编译正则表达式并执行匹配。
// 如果匹配成功，Count 和 SubmatchCount2 会被赋值。
//
// 参数:
//   - 正则表达式文本: 正则表达式模式字符串
//   - 被搜索的文本: 要执行匹配的目标文本
//
// 返回:
//   - bool: true 表示至少有一个匹配，false 表示无匹配
func (this *L_正则表达式) E创建(正则表达式文本 string, 被搜索的文本 string) bool {
	this.r = regexp.MustCompile(正则表达式文本)
	this.res = this.r.FindAllStringSubmatch(被搜索的文本, -1)
	this.Count = len(this.res)
	if this.Count == 0 {
		return false
	}
	this.SubmatchCount2 = len(this.res[0])
	return true
}

// Q取匹配数量 获取正则匹配到的结果数量。
//
// 返回:
//   - int: 匹配到的结果数量
func (this *L_正则表达式) Q取匹配数量() int {
	return len(this.res)
}

// Q取匹配文本 获取指定索引的完整匹配文本。
// 索引 0 表示第一个匹配结果的完整文本。
//
// 参数:
//   - 匹配索引: 匹配结果的索引（从 0 开始）
//
// 返回:
//   - string: 完整匹配文本；索引越界时可能 panic
func (this *L_正则表达式) Q取匹配文本(匹配索引 int) string {
	return this.res[匹配索引][0]
}

// Q取子匹配文本 获取指定匹配结果中某个子表达式的匹配文本。
// 子表达式索引 0 为完整匹配，1 为第一个捕获组，以此类推。
//
// 参数:
//   - 匹配索引: 匹配结果的索引（从 0 开始）
//   - 子表达式索引: 子表达式的索引（从 0 开始，0 为完整匹配）
//
// 返回:
//   - string: 子匹配文本；索引越界时返回空串
func (this *L_正则表达式) Q取子匹配文本(匹配索引 int, 子表达式索引 int) string {
	if -1 >= 匹配索引 || -1 >= 子表达式索引 {
		return ""
	}
	if this.Count <= 匹配索引 || this.SubmatchCount2 <= 子表达式索引 {
		return ""
	}

	return this.res[匹配索引][子表达式索引]
}

// Q取子匹配数量 获取每个匹配结果中子表达式的数量（含完整匹配）。
//
// 返回:
//   - int: 子表达式数量
func (this *L_正则表达式) Q取子匹配数量() int {
	return this.SubmatchCount2
}

// GetResult 获取所有匹配结果的原始二维字符串数组。
// 外层为匹配索引，内层为子表达式索引。
//
// 返回:
//   - [][]string: 所有匹配结果
func (this *L_正则表达式) GetResult() [][]string {
	return this.res
}
