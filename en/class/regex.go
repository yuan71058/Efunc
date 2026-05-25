package class

import "regexp"

// Regex 正则表达式封装类，提供正则匹配和子表达式提取功能。
// 创建时编译正则表达式并执行匹配，后续可通过方法获取匹配结果。
//
// 使用示例:
//
//	re, ok := NewRegex(`(\d+)-(\d+)`, "abc 123-456 def")
//	if ok {
//	    count := re.MatchCount()
//	    text := re.MatchText(0)
//	    sub := re.SubmatchText(0, 1)
//	}
type Regex struct {
	r              *regexp.Regexp
	res            [][]string
	Count          int
	SubmatchCount2 int
}

// NewRegex 创建并初始化正则表达式对象。
// 编译正则表达式文本，并在被搜索文本中执行匹配。
//
// 参数:
//   - pattern: 正则表达式模式字符串
//   - text: 要执行匹配的目标文本
//
// 返回:
//   - *Regex: 正则表达式对象
//   - bool: true 表示至少有一个匹配，false 表示无匹配
func NewRegex(pattern string, text string) (*Regex, bool) {
	t := new(Regex)
	b := t.Create(pattern, text)
	return t, b
}

// Create 编译正则表达式并执行匹配。
// 如果匹配成功，Count 和 SubmatchCount2 会被赋值。
//
// 参数:
//   - pattern: 正则表达式模式字符串
//   - text: 要执行匹配的目标文本
//
// 返回:
//   - bool: true 表示至少有一个匹配，false 表示无匹配
func (this *Regex) Create(pattern string, text string) bool {
	this.r = regexp.MustCompile(pattern)
	this.res = this.r.FindAllStringSubmatch(text, -1)
	this.Count = len(this.res)
	if this.Count == 0 {
		return false
	}
	this.SubmatchCount2 = len(this.res[0])
	return true
}

// MatchCount 获取正则匹配到的结果数量。
//
// 返回:
//   - int: 匹配到的结果数量
func (this *Regex) MatchCount() int {
	return len(this.res)
}

// MatchText 获取指定索引的完整匹配文本。
// 索引 0 表示第一个匹配结果的完整文本。
//
// 参数:
//   - matchIndex: 匹配结果的索引（从 0 开始）
//
// 返回:
//   - string: 完整匹配文本；索引越界时可能 panic
func (this *Regex) MatchText(matchIndex int) string {
	return this.res[matchIndex][0]
}

// SubmatchText 获取指定匹配结果中某个子表达式的匹配文本。
// 子表达式索引 0 为完整匹配，1 为第一个捕获组，以此类推。
//
// 参数:
//   - matchIndex: 匹配结果的索引（从 0 开始）
//   - submatchIndex: 子表达式的索引（从 0 开始，0 为完整匹配）
//
// 返回:
//   - string: 子匹配文本；索引越界时返回空串
func (this *Regex) SubmatchText(matchIndex int, submatchIndex int) string {
	if -1 >= matchIndex || -1 >= submatchIndex {
		return ""
	}
	if this.Count <= matchIndex || this.SubmatchCount2 <= submatchIndex {
		return ""
	}

	return this.res[matchIndex][submatchIndex]
}

// SubmatchCount 获取每个匹配结果中子表达式的数量（含完整匹配）。
//
// 返回:
//   - int: 子表达式数量
func (this *Regex) SubmatchCount() int {
	return this.SubmatchCount2
}

// GetResult 获取所有匹配结果的原始二维字符串数组。
// 外层为匹配索引，内层为子表达式索引。
//
// 返回:
//   - [][]string: 所有匹配结果
func (this *Regex) GetResult() [][]string {
	return this.res
}