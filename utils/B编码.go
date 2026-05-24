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

	"golang.org/x/net/idna"
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
