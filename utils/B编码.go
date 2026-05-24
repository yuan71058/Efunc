package utils

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"mime"
	"mime/quotedprintable"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"golang.org/x/net/idna"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ============================================================
// URL 编码
// ============================================================

// B编码_URL编码 对文本进行 URL 编码（百分号编码）。
// 将特殊字符转换为 %XX 格式，中文等非 ASCII 字符也会被编码。
//
// 参数:
//   - 欲编码的文本: 待编码的文本
//
// 返回:
//   - string: URL 编码后的文本
func B编码_URL编码(欲编码的文本 string) string {
	return url.QueryEscape(欲编码的文本)
}

// B编码_URL解码 对 URL 编码的文本进行解码。
// 将 %XX 格式的编码还原为原始字符。
//
// 参数:
//   - URL: URL 编码的文本
//
// 返回:
//   - string: 解码后的文本；解码失败返回空串
func B编码_URL解码(URL string) string {
	decodedURL, err := url.QueryUnescape(URL)
	if err != nil {
		return ""
	}
	return decodedURL
}

// B编码_URL路径编码 对 URL 路径进行编码，保留 / 等路径分隔符。
// 与 B编码_URL编码 不同，此函数不会编码 / 、& 、= 等路径合法字符。
//
// 参数:
//   - 路径: 待编码的 URL 路径
//
// 返回:
//   - string: 编码后的路径
func B编码_URL路径编码(路径 string) string {
	return url.PathEscape(路径)
}

// B编码_URL路径解码 对 URL 路径编码进行解码。
//
// 参数:
//   - 路径: URL 编码的路径
//
// 返回:
//   - string: 解码后的路径；失败返回空串
func B编码_URL路径解码(路径 string) string {
	结果, err := url.PathUnescape(路径)
	if err != nil {
		return ""
	}
	return 结果
}

// B编码_URL组件编码 编码 URL 组件（对路径和查询参数分别编码）。
// 返回编码后的完整 URL 字符串。
//
// 参数:
//   - 网址: 原始 URL 字符串
//
// 返回:
//   - string: 编码后的 URL；解析失败返回原串
func B编码_URL组件编码(网址 string) string {
	u, err := url.Parse(网址)
	if err != nil {
		return 网址
	}
	return u.String()
}

// ============================================================
// Base64 编码
// ============================================================

// B编码_BASE64编码 将字节集进行 Base64 编码。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: Base64 编码后的字符串
func B编码_BASE64编码(字节集 []byte) string {
	return base64.StdEncoding.EncodeToString(字节集)
}

// B编码_BASE64解码 将 Base64 编码的文本解码为字节集。
//
// 参数:
//   - 文本: Base64 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_BASE64解码(文本 string) []byte {
	解码字节集, err := base64.StdEncoding.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}

// B编码_BASE64URL编码 将字节集进行 URL 安全的 Base64 编码。
// 使用 - 替换 +，_ 替换 /，不含 = 填充，适合用于 URL 参数。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: URL 安全的 Base64 编码字符串
func B编码_BASE64URL编码(字节集 []byte) string {
	return base64.RawURLEncoding.EncodeToString(字节集)
}

// B编码_BASE64URL解码 将 URL 安全的 Base64 编码文本解码为字节集。
//
// 参数:
//   - 文本: URL 安全的 Base64 编码文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_BASE64URL解码(文本 string) []byte {
	解码字节集, err := base64.RawURLEncoding.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}

// B编码_BASE64无填充编码 将字节集进行标准 Base64 编码（无 = 填充）。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: 无填充的 Base64 编码字符串
func B编码_BASE64无填充编码(字节集 []byte) string {
	return base64.RawStdEncoding.EncodeToString(字节集)
}

// B编码_BASE64无填充解码 将无填充的 Base64 编码文本解码为字节集。
//
// 参数:
//   - 文本: 无填充的 Base64 编码文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_BASE64无填充解码(文本 string) []byte {
	解码字节集, err := base64.RawStdEncoding.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}

// ============================================================
// Base32 编码
// ============================================================

// B编码_BASE32编码 将字节集进行 Base32 编码。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: Base32 编码后的字符串
func B编码_BASE32编码(字节集 []byte) string {
	return base32.StdEncoding.EncodeToString(字节集)
}

// B编码_BASE32解码 将 Base32 编码的文本解码为字节集。
//
// 参数:
//   - 文本: Base32 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_BASE32解码(文本 string) []byte {
	解码字节集, err := base32.StdEncoding.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}

// B编码_BASE32HEX编码 将字节集进行 Base32 Hex 编码（使用 0-9 A-V 字符集）。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: Base32 Hex 编码后的字符串
func B编码_BASE32HEX编码(字节集 []byte) string {
	return base32.HexEncoding.EncodeToString(字节集)
}

// B编码_BASE32HEX解码 将 Base32 Hex 编码的文本解码为字节集。
//
// 参数:
//   - 文本: Base32 Hex 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_BASE32HEX解码(文本 string) []byte {
	解码字节集, err := base32.HexEncoding.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}

// ============================================================
// Hex 编码
// ============================================================

// B编码_十六进制编码 将字节集编码为十六进制字符串（小写）。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: 十六进制编码字符串
func B编码_十六进制编码(字节集 []byte) string {
	return hex.EncodeToString(字节集)
}

// B编码_十六进制解码 将十六进制字符串解码为字节集。
//
// 参数:
//   - 文本: 十六进制编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_十六进制解码(文本 string) []byte {
	解码字节集, err := hex.DecodeString(文本)
	if err != nil {
		return []byte{}
	}
	return 解码字节集
}

// B编码_十六进制大写 将字节集编码为大写十六进制字符串。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: 大写十六进制编码字符串
func B编码_十六进制大写(字节集 []byte) string {
	return strings.ToUpper(hex.EncodeToString(字节集))
}

// ============================================================
// Unicode / USC2 编码
// ============================================================

// B编码_usc2到文本 将 USC2/Unicode 转义序列转换为中文文本。
// 例如将 \u4e2d\u6587 转换为 "中文"。
//
// 参数:
//   - 字符串: 包含 USC2 转义序列的字符串（如 "\\u4e2d\\u6587"）
//
// 返回:
//   - string: 转换后的中文文本；失败返回空串
func B编码_usc2到文本(字符串 string) string {
	解码文本, err := strconv.Unquote(`"` + 字符串 + `"`)
	if err != nil {
		return ""
	}
	return 解码文本
}

// B编码_文本到USC2 将中文文本转换为 USC2/Unicode 转义序列。
// 例如将 "中文" 转换为 \u4e2d\u6587。
//
// 参数:
//   - 文本: 待转换的文本
//
// 返回:
//   - string: USC2 转义序列字符串
func B编码_文本到USC2(文本 string) string {
	var 结果 strings.Builder
	for _, r := range 文本 {
		if r > 127 {
			结果.WriteString(fmt.Sprintf("\\u%04x", r))
		} else {
			结果.WriteRune(r)
		}
	}
	return 结果.String()
}

// B编码_文本到Unicode 将文本转换为 \\uXXXX 格式的 Unicode 转义序列（所有字符）。
//
// 参数:
//   - 文本: 待转换的文本
//
// 返回:
//   - string: 所有字符的 Unicode 转义序列
func B编码_文本到Unicode(文本 string) string {
	var 结果 strings.Builder
	for _, r := range 文本 {
		结果.WriteString(fmt.Sprintf("\\u%04x", r))
	}
	return 结果.String()
}

// ============================================================
// HTML 编码
// ============================================================

// B编码_HTML编码 将文本中的 HTML 特殊字符进行转义编码。
// 将 & < > " ' 转换为 HTML 实体。
//
// 参数:
//   - 文本: 待编码的文本
//
// 返回:
//   - string: HTML 编码后的文本
func B编码_HTML编码(文本 string) string {
	return html.EscapeString(文本)
}

// B编码_HTML解码 将 HTML 实体解码为原始字符。
// 将 &amp; &lt; &gt; &quot; &#39; 等实体还原。
//
// 参数:
//   - 文本: 包含 HTML 实体的文本
//
// 返回:
//   - string: 解码后的文本
func B编码_HTML解码(文本 string) string {
	return html.UnescapeString(文本)
}

// ============================================================
// Quoted-Printable 编码
// ============================================================

// B编码_QP编码 将字节集进行 Quoted-Printable 编码。
// 常用于邮件传输中的非 ASCII 字符编码。
//
// 参数:
//   - 字节集: 待编码的字节集
//
// 返回:
//   - string: QP 编码后的字符串
func B编码_QP编码(字节集 []byte) string {
	var buf bytes.Buffer
	writer := quotedprintable.NewWriter(&buf)
	writer.Write(字节集)
	writer.Close()
	return buf.String()
}

// B编码_QP解码 将 Quoted-Printable 编码的文本解码为字节集。
//
// 参数:
//   - 文本: QP 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func B编码_QP解码(文本 string) []byte {
	reader := quotedprintable.NewReader(strings.NewReader(文本))
	数据, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return 数据
}

// ============================================================
// JSON 编码
// ============================================================

// B编码_JSON编码 将 Go 值编码为 JSON 字符串。
//
// 参数:
//   - 值: 待编码的 Go 值（结构体、Map 等）
//
// 返回:
//   - string: JSON 编码字符串；编码失败返回空串
func B编码_JSON编码(值 interface{}) string {
	数据, err := json.Marshal(值)
	if err != nil {
		return ""
	}
	return string(数据)
}

// B编码_JSON编码缩进 将 Go 值编码为带缩进的 JSON 字符串。
//
// 参数:
//   - 值: 待编码的 Go 值
//   - 前缀: 每行前缀（可为空）
//   - 缩进: 缩进字符串（如 "  " 或 "\t"）
//
// 返回:
//   - string: 带缩进的 JSON 字符串；编码失败返回空串
func B编码_JSON编码缩进(值 interface{}, 前缀 string, 缩进 string) string {
	数据, err := json.MarshalIndent(值, 前缀, 缩进)
	if err != nil {
		return ""
	}
	return string(数据)
}

// B编码_JSON解码 将 JSON 字符串解码到目标变量。
//
// 参数:
//   - 文本: JSON 编码的文本
//   - 目标: 解码目标的指针
//
// 返回:
//   - error: 解码失败时返回错误
func B编码_JSON解码(文本 string, 目标 interface{}) error {
	return json.Unmarshal([]byte(文本), 目标)
}

// ============================================================
// MIME 编码
// ============================================================

// B编码_MIME编码 对文本进行 MIME 编码（用于邮件头部）。
// 自动选择 Quoted-Printable 或 Base64 编码方式。
//
// 参数:
//   - 文本: 待编码的文本
//
// 返回:
//   - string: MIME 编码字符串
func B编码_MIME编码(文本 string) string {
	return mime.QEncoding.Encode("utf-8", 文本)
}

// B编码_MIMEB64编码 对文本进行 MIME Base64 编码（用于邮件头部）。
//
// 参数:
//   - 文本: 待编码的文本
//
// 返回:
//   - string: MIME Base64 编码字符串
func B编码_MIMEB64编码(文本 string) string {
	return mime.BEncoding.Encode("utf-8", 文本)
}

// ============================================================
// 字节序编码
// ============================================================

// B编码_整数到大端 将 uint16/uint32/uint64 整数编码为大端字节序。
//
// 参数:
//   - 值: 待编码的整数值
//
// 返回:
//   - []byte: 大端字节序的字节集
func B编码_整数到大端(值 interface{}) []byte {
	switch v := 值.(type) {
	case uint16:
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, v)
		return buf
	case uint32:
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, v)
		return buf
	case uint64:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, v)
		return buf
	default:
		return []byte{}
	}
}

// B编码_整数到小端 将 uint16/uint32/uint64 整数编码为小端字节序。
//
// 参数:
//   - 值: 待编码的整数值
//
// 返回:
//   - []byte: 小端字节序的字节集
func B编码_整数到小端(值 interface{}) []byte {
	switch v := 值.(type) {
	case uint16:
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, v)
		return buf
	case uint32:
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, v)
		return buf
	case uint64:
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, v)
		return buf
	default:
		return []byte{}
	}
}

// B编码_大端到整数 将大端字节序解码为 uint64 整数。
//
// 参数:
//   - 字节集: 大端字节序的字节集
//
// 返回:
//   - uint64: 解码后的整数值
func B编码_大端到整数(字节集 []byte) uint64 {
	switch len(字节集) {
	case 2:
		return uint64(binary.BigEndian.Uint16(字节集))
	case 4:
		return uint64(binary.BigEndian.Uint32(字节集))
	case 8:
		return binary.BigEndian.Uint64(字节集)
	default:
		return 0
	}
}

// B编码_小端到整数 将小端字节序解码为 uint64 整数。
//
// 参数:
//   - 字节集: 小端字节序的字节集
//
// 返回:
//   - uint64: 解码后的整数值
func B编码_小端到整数(字节集 []byte) uint64 {
	switch len(字节集) {
	case 2:
		return uint64(binary.LittleEndian.Uint16(字节集))
	case 4:
		return uint64(binary.LittleEndian.Uint32(字节集))
	case 8:
		return binary.LittleEndian.Uint64(字节集)
	default:
		return 0
	}
}

// ============================================================
// Punycode 编码
// ============================================================

// B编码_Punycode编码 将国际化域名进行 Punycode 编码。
// 例如 "中文.com" 编码为 "xn--fiq228c.com"。
//
// 参数:
//   - 域名: 国际化域名
//
// 返回:
//   - string: Punycode 编码后的域名；编码失败返回原串
func B编码_Punycode编码(域名 string) string {
	结果, err := idnaToASCII(域名)
	if err != nil {
		return 域名
	}
	return 结果
}

// B编码_Punycode解码 将 Punycode 编码的域名解码为原始域名。
//
// 参数:
//   - 域名: Punycode 编码的域名
//
// 返回:
//   - string: 解码后的域名；解码失败返回原串
func B编码_Punycode解码(域名 string) string {
	结果, err := idnaToUnicode(域名)
	if err != nil {
		return 域名
	}
	return 结果
}

// idnaToASCII 将域名转换为 ASCII 形式（Punycode）。
func idnaToASCII(域名 string) (string, error) {
	return idna.ToASCII(域名)
}

// idnaToUnicode 将 Punycode 域名转换为 Unicode 形式。
func idnaToUnicode(域名 string) (string, error) {
	return idna.ToUnicode(域名)
}

// ============================================================
// ANSI/GBK 编码
// ============================================================

// B编码_UTF8到GBK 将 UTF-8 编码的文本转换为 GBK（ANSI）编码。
// 中文 Windows 系统常用 GBK 编码。
//
// 参数:
//   - 文本: UTF-8 编码的文本
//
// 返回:
//   - []byte: GBK 编码的字节集；转换失败返回空字节集
func B编码_UTF8到GBK(文本 string) []byte {
	reader := transform.NewReader(strings.NewReader(文本), simplifiedchinese.GBK.NewEncoder())
	结果, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return 结果
}

// B编码_GBK到UTF8 将 GBK（ANSI）编码的字节集转换为 UTF-8 文本。
//
// 参数:
//   - 数据: GBK 编码的字节集
//
// 返回:
//   - string: UTF-8 编码的文本；转换失败返回空串
func B编码_GBK到UTF8(数据 []byte) string {
	reader := transform.NewReader(bytes.NewReader(数据), simplifiedchinese.GBK.NewDecoder())
	结果, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}
	return string(结果)
}

// B编码_UTF8到GB18030 将 UTF-8 编码的文本转换为 GB18030 编码。
// GB18030 是 GBK 的超集，支持更多中文字符。
//
// 参数:
//   - 文本: UTF-8 编码的文本
//
// 返回:
//   - []byte: GB18030 编码的字节集；转换失败返回空字节集
func B编码_UTF8到GB18030(文本 string) []byte {
	reader := transform.NewReader(strings.NewReader(文本), simplifiedchinese.GB18030.NewEncoder())
	结果, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return 结果
}

// B编码_GB18030到UTF8 将 GB18030 编码的字节集转换为 UTF-8 文本。
//
// 参数:
//   - 数据: GB18030 编码的字节集
//
// 返回:
//   - string: UTF-8 编码的文本；转换失败返回空串
func B编码_GB18030到UTF8(数据 []byte) string {
	reader := transform.NewReader(bytes.NewReader(数据), simplifiedchinese.GB18030.NewDecoder())
	结果, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}
	return string(结果)
}

// ============================================================
// UTF-16 编码
// ============================================================

// B编码_UTF8到UTF16 将 UTF-8 文本转换为 UTF-16 小端序字节集。
// Windows API 常用 UTF-16LE 编码。
//
// 参数:
//   - 文本: UTF-8 编码的文本
//
// 返回:
//   - []byte: UTF-16LE 编码的字节集（含 BOM 头 FF FE）
func B编码_UTF8到UTF16(文本 string) []byte {
	runes := []rune(文本)
	u16 := utf16.Encode(runes)
	结果 := make([]byte, 0, len(u16)*2+2)
	结果 = append(结果, 0xFF, 0xFE)
	for _, v := range u16 {
		结果 = append(结果, byte(v), byte(v>>8))
	}
	return 结果
}

// B编码_UTF16到UTF8 将 UTF-16 字节集转换为 UTF-8 文本。
// 自动识别 BOM 头判断字节序，无 BOM 默认小端。
//
// 参数:
//   - 数据: UTF-16 编码的字节集
//
// 返回:
//   - string: UTF-8 编码的文本
func B编码_UTF16到UTF8(数据 []byte) string {
	if len(数据) < 2 {
		return ""
	}
	大端 := false
	偏移 := 0
	if 数据[0] == 0xFE && 数据[1] == 0xFF {
		大端 = true
		偏移 = 2
	} else if 数据[0] == 0xFF && 数据[1] == 0xFE {
		偏移 = 2
	}

	u16 := make([]uint16, 0, (len(数据)-偏移)/2)
	for i := 偏移; i+1 < len(数据); i += 2 {
		if 大端 {
			u16 = append(u16, uint16(数据[i])<<8|uint16(数据[i+1]))
		} else {
			u16 = append(u16, uint16(数据[i+1])<<8|uint16(数据[i]))
		}
	}
	runes := utf16.Decode(u16)
	return string(runes)
}

// B编码_UTF8到UTF16大端 将 UTF-8 文本转换为 UTF-16 大端序字节集。
//
// 参数:
//   - 文本: UTF-8 编码的文本
//
// 返回:
//   - []byte: UTF-16BE 编码的字节集（含 BOM 头 FE FF）
func B编码_UTF8到UTF16大端(文本 string) []byte {
	runes := []rune(文本)
	u16 := utf16.Encode(runes)
	结果 := make([]byte, 0, len(u16)*2+2)
	结果 = append(结果, 0xFE, 0xFF)
	for _, v := range u16 {
		结果 = append(结果, byte(v>>8), byte(v))
	}
	return 结果
}

// ============================================================
// UTF-8 BOM 处理
// ============================================================

// B编码_添加UTF8BOM 为 UTF-8 字节集添加 BOM 头（EF BB BF）。
// 某些 Windows 程序需要 BOM 头来识别 UTF-8 编码。
//
// 参数:
//   - 数据: UTF-8 编码的字节集
//
// 返回:
//   - []byte: 带 BOM 头的 UTF-8 字节集
func B编码_添加UTF8BOM(数据 []byte) []byte {
	if len(数据) >= 3 && 数据[0] == 0xEF && 数据[1] == 0xBB && 数据[2] == 0xBF {
		return 数据
	}
	结果 := make([]byte, 0, len(数据)+3)
	结果 = append(结果, 0xEF, 0xBB, 0xBF)
	结果 = append(结果, 数据...)
	return 结果
}

// B编码_移除UTF8BOM 移除 UTF-8 字节集的 BOM 头。
//
// 参数:
//   - 数据: 可能包含 BOM 头的 UTF-8 字节集
//
// 返回:
//   - []byte: 移除 BOM 头后的字节集
func B编码_移除UTF8BOM(数据 []byte) []byte {
	if len(数据) >= 3 && 数据[0] == 0xEF && 数据[1] == 0xBB && 数据[2] == 0xBF {
		return 数据[3:]
	}
	return 数据
}

// B编码_是否有UTF8BOM 检查字节集是否包含 UTF-8 BOM 头。
//
// 参数:
//   - 数据: 待检查的字节集
//
// 返回:
//   - bool: 包含 BOM 返回 true
func B编码_是否有UTF8BOM(数据 []byte) bool {
	return len(数据) >= 3 && 数据[0] == 0xEF && 数据[1] == 0xBB && 数据[2] == 0xBF
}

// ============================================================
// Unicode 码点操作
// ============================================================

// B编码_Unicode解码 将 \uXXXX 格式的 Unicode 转义序列解码为文本。
// 支持 \uXXXX 和 \UXXXXXXXX 两种格式。
//
// 参数:
//   - 文本: 包含 Unicode 转义序列的文本
//
// 返回:
//   - string: 解码后的文本
func B编码_Unicode解码(文本 string) string {
	var 结果 strings.Builder
	i := 0
	for i < len(文本) {
		if i+5 < len(文本) && 文本[i] == '\\' && 文本[i+1] == 'u' {
			hex := 文本[i+2 : i+6]
			val, err := strconv.ParseUint(hex, 16, 32)
			if err == nil {
				结果.WriteRune(rune(val))
				i += 6
				continue
			}
		}
		if i+9 < len(文本) && 文本[i] == '\\' && 文本[i+1] == 'U' {
			hex := 文本[i+2 : i+10]
			val, err := strconv.ParseUint(hex, 16, 32)
			if err == nil {
				结果.WriteRune(rune(val))
				i += 10
				continue
			}
		}
		结果.WriteByte(文本[i])
		i++
	}
	return 结果.String()
}

// B编码_取Unicode码点 获取文本中指定位置字符的 Unicode 码点值。
//
// 参数:
//   - 文本: 输入文本
//   - 位置: 字符位置（从 0 开始，按 rune 计数）
//
// 返回:
//   - int: Unicode 码点值；位置越界返回 -1
func B编码_取Unicode码点(文本 string, 位置 int) int {
	runes := []rune(文本)
	if 位置 < 0 || 位置 >= len(runes) {
		return -1
	}
	return int(runes[位置])
}

// B编码_码点到文本 将 Unicode 码点值转换为对应的字符。
//
// 参数:
//   - 码点: Unicode 码点值
//
// 返回:
//   - string: 对应的字符；无效码点返回空串
func B编码_码点到文本(码点 int) string {
	if 码点 < 0 || 码点 > 0x10FFFF {
		return ""
	}
	return string(rune(码点))
}

// B编码_取UTF8字节数 获取文本的 UTF-8 编码字节数。
//
// 参数:
//   - 文本: 输入文本
//
// 返回:
//   - int: UTF-8 编码后的字节长度
func B编码_取UTF8字节数(文本 string) int {
	return len([]byte(文本))
}

// B编码_取字符数 获取文本的 Unicode 字符数（按 rune 计数）。
// 与 len() 不同，此函数正确计算多字节字符。
//
// 参数:
//   - 文本: 输入文本
//
// 返回:
//   - int: 字符数量
func B编码_取字符数(文本 string) int {
	return utf8.RuneCountInString(文本)
}

// B编码_是否有效UTF8 检查字节集是否为有效的 UTF-8 编码。
//
// 参数:
//   - 数据: 待检查的字节集
//
// 返回:
//   - bool: 有效返回 true
func B编码_是否有效UTF8(数据 []byte) bool {
	return utf8.Valid(数据)
}
