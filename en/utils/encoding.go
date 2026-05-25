package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"html"
	"io"
	"mime"
	"mime/quotedprintable"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/idna"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ============================================================
// MD5 编码
// ============================================================

// Encoding_MD5 计算字节集的 MD5 哈希值并返回十六进制字符串（小写）。
//
// 参数:
//   - data: 待计算哈希的字节集
//
// 返回:
//   - string: 小写 MD5 十六进制字符串
func Encoding_MD5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// Encoding_MD5FromText 计算文本的 MD5 哈希值并返回十六进制字符串（小写）。
//
// 参数:
//   - text: 待计算哈希的文本
//
// 返回:
//   - string: 小写 MD5 十六进制字符串
func Encoding_MD5FromText(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// ============================================================
// SHA 编码
// ============================================================

// Encoding_SHA1 计算字节集的 SHA1 哈希值并返回十六进制字符串（小写）。
//
// 参数:
//   - data: 待计算哈希的字节集
//
// 返回:
//   - string: 小写 SHA1 十六进制字符串
func Encoding_SHA1(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

// Encoding_SHA256 计算字节集的 SHA256 哈希值并返回十六进制字符串（小写）。
//
// 参数:
//   - data: 待计算哈希的字节集
//
// 返回:
//   - string: 小写 SHA256 十六进制字符串
func Encoding_SHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// Encoding_SHA512 计算字节集的 SHA512 哈希值并返回十六进制字符串（小写）。
//
// 参数:
//   - data: 待计算哈希的字节集
//
// 返回:
//   - string: 小写 SHA512 十六进制字符串
func Encoding_SHA512(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

// ============================================================
// CRC 编码
// ============================================================

// Encoding_CRC32 计算字节集的 CRC32 校验值（IEEE 多项式）。
//
// 参数:
//   - data: 待计算校验的字节集
//
// 返回:
//   - uint32: CRC32 校验值
func Encoding_CRC32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

// Encoding_CRC64 计算字节集的 CRC64 校验值（ISO 多项式）。
//
// 参数:
//   - data: 待计算校验的字节集
//
// 返回:
//   - uint64: CRC64 校验值
func Encoding_CRC64(data []byte) uint64 {
	table := crc64.MakeTable(crc64.ISO)
	return crc64.Checksum(data, table)
}

// ============================================================
// Base64 编码
// ============================================================

// Encoding_Base64Encode 将字节集编码为 Base64 字符串（标准编码）。
//
// 参数:
//   - data: 待编码的字节集
//
// 返回:
//   - string: Base64 编码字符串
func Encoding_Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Encoding_Base64Decode 将 Base64 字符串解码为字节集。
//
// 参数:
//   - text: Base64 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func Encoding_Base64Decode(text string) []byte {
	result, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return []byte{}
	}
	return result
}

// Encoding_Base64URLEncode 将字节集编码为 URL 安全的 Base64 字符串。
// 使用 - 和 _ 替代 + 和 /，去掉尾部 = 填充符。
//
// 参数:
//   - data: 待编码的字节集
//
// 返回:
//   - string: URL 安全的 Base64 字符串
func Encoding_Base64URLEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

// Encoding_Base64URLDecode 将 URL 安全的 Base64 字符串解码为字节集。
//
// 参数:
//   - text: URL 安全的 Base64 文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func Encoding_Base64URLDecode(text string) []byte {
	result, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(text)
	if err != nil {
		return []byte{}
	}
	return result
}

// ============================================================
// Base32 编码
// ============================================================

// Encoding_Base32Encode 将字节集编码为 Base32 字符串（标准编码）。
//
// 参数:
//   - data: 待编码的字节集
//
// 返回:
//   - string: Base32 编码字符串
func Encoding_Base32Encode(data []byte) string {
	return base32.StdEncoding.EncodeToString(data)
}

// Encoding_Base32Decode 将 Base32 字符串解码为字节集。
//
// 参数:
//   - text: Base32 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func Encoding_Base32Decode(text string) []byte {
	result, err := base32.StdEncoding.DecodeString(text)
	if err != nil {
		return []byte{}
	}
	return result
}

// ============================================================
// URL 编码
// ============================================================

// Encoding_URLEncode 对文本进行 URL 编码（百分号编码）。
// 将特殊字符（中文、空格等）转换为 %XX 格式。
//
// 参数:
//   - text: 待编码的文本
//
// 返回:
//   - string: URL 编码字符串
func Encoding_URLEncode(text string) string {
	return url.QueryEscape(text)
}

// Encoding_URLDecode 将 URL 编码的文本解码为原始文本。
//
// 参数:
//   - text: URL 编码的文本
//
// 返回:
//   - string: 解码后的文本；解码失败返回空串
func Encoding_URLDecode(text string) string {
	result, err := url.QueryUnescape(text)
	if err != nil {
		return ""
	}
	return result
}

// Encoding_URLFullEncode 对文本进行完整的 URL 编码（编码所有特殊字符）。
//
// 参数:
//   - text: 待编码的文本
//
// 返回:
//   - string: 完整 URL 编码字符串
func Encoding_URLFullEncode(text string) string {
	var result strings.Builder
	for _, b := range []byte(text) {
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') || b == '-' || b == '_' || b == '.' || b == '~' {
			result.WriteByte(b)
		} else {
			result.WriteString(fmt.Sprintf("%%%02X", b))
		}
	}
	return result.String()
}

// ============================================================
// Hex 编码
// ============================================================

// Encoding_HexEncode 将字节集编码为十六进制字符串（小写）。
//
// 参数:
//   - data: 待编码的字节集
//
// 返回:
//   - string: 十六进制编码字符串
func Encoding_HexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

// Encoding_HexDecode 将十六进制字符串解码为字节集。
//
// 参数:
//   - text: 十六进制编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func Encoding_HexDecode(text string) []byte {
	result, err := hex.DecodeString(text)
	if err != nil {
		return []byte{}
	}
	return result
}

// Encoding_HexEncodeUpper 将字节集编码为大写十六进制字符串。
//
// 参数:
//   - data: 待编码的字节集
//
// 返回:
//   - string: 大写十六进制编码字符串
func Encoding_HexEncodeUpper(data []byte) string {
	return strings.ToUpper(hex.EncodeToString(data))
}

// ============================================================
// Unicode / USC2 编码
// ============================================================

// Encoding_USC2ToText 将 USC2/Unicode 转义序列转换为中文文本。
// 例如将 \u4e2d\u6587 转换为 "中文"。
//
// 参数:
//   - text: 包含 USC2 转义序列的字符串（如 "\\u4e2d\\u6587"）
//
// 返回:
//   - string: 转换后的中文文本；失败返回空串
func Encoding_USC2ToText(text string) string {
	result, err := strconv.Unquote(`"` + text + `"`)
	if err != nil {
		return ""
	}
	return result
}

// Encoding_TextToUSC2 将中文文本转换为 USC2/Unicode 转义序列。
// 例如将 "中文" 转换为 \u4e2d\u6587。
//
// 参数:
//   - text: 待转换的文本
//
// 返回:
//   - string: USC2 转义序列字符串
func Encoding_TextToUSC2(text string) string {
	var result strings.Builder
	for _, r := range text {
		if r > 127 {
			result.WriteString(fmt.Sprintf("\\u%04x", r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// Encoding_TextToUnicode 将文本转换为 \\uXXXX 格式的 Unicode 转义序列（所有字符）。
//
// 参数:
//   - text: 待转换的文本
//
// 返回:
//   - string: 所有字符的 Unicode 转义序列
func Encoding_TextToUnicode(text string) string {
	var result strings.Builder
	for _, r := range text {
		result.WriteString(fmt.Sprintf("\\u%04x", r))
	}
	return result.String()
}

// ============================================================
// HTML 编码
// ============================================================

// Encoding_HTMLEncode 将文本中的 HTML 特殊字符进行转义编码。
// 将 & < > " ' 转换为 HTML 实体。
//
// 参数:
//   - text: 待编码的文本
//
// 返回:
//   - string: HTML 编码后的文本
func Encoding_HTMLEncode(text string) string {
	return html.EscapeString(text)
}

// Encoding_HTMLDecode 将 HTML 实体解码为原始字符。
// 将 &amp; &lt; &gt; &quot; &#39; 等实体还原。
//
// 参数:
//   - text: 包含 HTML 实体的文本
//
// 返回:
//   - string: 解码后的文本
func Encoding_HTMLDecode(text string) string {
	return html.UnescapeString(text)
}

// ============================================================
// Quoted-Printable 编码
// ============================================================

// Encoding_QPEncode 将字节集进行 Quoted-Printable 编码。
// 常用于邮件传输中的非 ASCII 字符编码。
//
// 参数:
//   - data: 待编码的字节集
//
// 返回:
//   - string: QP 编码后的字符串
func Encoding_QPEncode(data []byte) string {
	var buf bytes.Buffer
	writer := quotedprintable.NewWriter(&buf)
	writer.Write(data)
	writer.Close()
	return buf.String()
}

// Encoding_QPDecode 将 Quoted-Printable 编码的文本解码为字节集。
//
// 参数:
//   - text: QP 编码的文本
//
// 返回:
//   - []byte: 解码后的字节集；解码失败返回空字节集
func Encoding_QPDecode(text string) []byte {
	reader := quotedprintable.NewReader(strings.NewReader(text))
	data, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return data
}

// ============================================================
// JSON 编码
// ============================================================

// Encoding_JSONEncode 将 Go 值编码为 JSON 字符串。
//
// 参数:
//   - value: 待编码的 Go 值（结构体、Map 等）
//
// 返回:
//   - string: JSON 编码字符串；编码失败返回空串
func Encoding_JSONEncode(value interface{}) string {
	data, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(data)
}

// Encoding_JSONEncodeIndent 将 Go 值编码为带缩进的 JSON 字符串。
//
// 参数:
//   - value: 待编码的 Go 值
//   - prefix: 每行前缀（可为空）
//   - indent: 缩进字符串（如 "  " 或 "\t"）
//
// 返回:
//   - string: 带缩进的 JSON 字符串；编码失败返回空串
func Encoding_JSONEncodeIndent(value interface{}, prefix string, indent string) string {
	data, err := json.MarshalIndent(value, prefix, indent)
	if err != nil {
		return ""
	}
	return string(data)
}

// Encoding_JSONDecode 将 JSON 字符串解码到目标变量。
//
// 参数:
//   - text: JSON 编码的文本
//   - target: 解码目标的指针
//
// 返回:
//   - error: 解码失败时返回错误
func Encoding_JSONDecode(text string, target interface{}) error {
	return json.Unmarshal([]byte(text), target)
}

// ============================================================
// MIME 编码
// ============================================================

// Encoding_MIMEEncode 对文本进行 MIME 编码（用于邮件头部）。
// 自动选择 Quoted-Printable 或 Base64 编码方式。
//
// 参数:
//   - text: 待编码的文本
//
// 返回:
//   - string: MIME 编码字符串
func Encoding_MIMEEncode(text string) string {
	return mime.QEncoding.Encode("utf-8", text)
}

// Encoding_MIMEB64Encode 对文本进行 MIME Base64 编码（用于邮件头部）。
//
// 参数:
//   - text: 待编码的文本
//
// 返回:
//   - string: MIME Base64 编码字符串
func Encoding_MIMEB64Encode(text string) string {
	return mime.BEncoding.Encode("utf-8", text)
}

// ============================================================
// 字节序编码
// ============================================================

// Encoding_IntToBigEndian 将 uint16/uint32/uint64 整数编码为大端字节序。
//
// 参数:
//   - value: 待编码的整数值
//
// 返回:
//   - []byte: 大端字节序的字节集
func Encoding_IntToBigEndian(value interface{}) []byte {
	switch v := value.(type) {
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

// Encoding_IntToLittleEndian 将 uint16/uint32/uint64 整数编码为小端字节序。
//
// 参数:
//   - value: 待编码的整数值
//
// 返回:
//   - []byte: 小端字节序的字节集
func Encoding_IntToLittleEndian(value interface{}) []byte {
	switch v := value.(type) {
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

// Encoding_BigEndianToInt 将大端字节序解码为 uint64 整数。
//
// 参数:
//   - data: 大端字节序的字节集
//
// 返回:
//   - uint64: 解码后的整数值
func Encoding_BigEndianToInt(data []byte) uint64 {
	switch len(data) {
	case 2:
		return uint64(binary.BigEndian.Uint16(data))
	case 4:
		return uint64(binary.BigEndian.Uint32(data))
	case 8:
		return binary.BigEndian.Uint64(data)
	default:
		return 0
	}
}

// Encoding_LittleEndianToInt 将小端字节序解码为 uint64 整数。
//
// 参数:
//   - data: 小端字节序的字节集
//
// 返回:
//   - uint64: 解码后的整数值
func Encoding_LittleEndianToInt(data []byte) uint64 {
	switch len(data) {
	case 2:
		return uint64(binary.LittleEndian.Uint16(data))
	case 4:
		return uint64(binary.LittleEndian.Uint32(data))
	case 8:
		return binary.LittleEndian.Uint64(data)
	default:
		return 0
	}
}

// ============================================================
// Punycode 编码
// ============================================================

// Encoding_PunycodeEncode 将国际化域名进行 Punycode 编码。
// 例如 "中文.com" 编码为 "xn--fiq228c.com"。
//
// 参数:
//   - domain: 国际化域名
//
// 返回:
//   - string: Punycode 编码后的域名；编码失败返回原串
func Encoding_PunycodeEncode(domain string) string {
	result, err := idna.ToASCII(domain)
	if err != nil {
		return domain
	}
	return result
}

// Encoding_PunycodeDecode 将 Punycode 编码的域名解码为原始域名。
//
// 参数:
//   - domain: Punycode 编码的域名
//
// 返回:
//   - string: 解码后的域名；解码失败返回原串
func Encoding_PunycodeDecode(domain string) string {
	result, err := idna.ToUnicode(domain)
	if err != nil {
		return domain
	}
	return result
}

// ============================================================
// ANSI/GBK 编码
// ============================================================

// Encoding_UTF8ToGBK 将 UTF-8 编码的文本转换为 GBK（ANSI）编码。
// 中文 Windows 系统常用 GBK 编码。
//
// 参数:
//   - text: UTF-8 编码的文本
//
// 返回:
//   - []byte: GBK 编码的字节集；转换失败返回空字节集
func Encoding_UTF8ToGBK(text string) []byte {
	reader := transform.NewReader(strings.NewReader(text), simplifiedchinese.GBK.NewEncoder())
	result, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return result
}

// Encoding_GBKToUTF8 将 GBK（ANSI）编码的字节集转换为 UTF-8 文本。
//
// 参数:
//   - data: GBK 编码的字节集
//
// 返回:
//   - string: UTF-8 编码的文本；转换失败返回空串
func Encoding_GBKToUTF8(data []byte) string {
	reader := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewDecoder())
	result, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}
	return string(result)
}

// Encoding_UTF8ToGB18030 将 UTF-8 编码的文本转换为 GB18030 编码。
// GB18030 是 GBK 的超集，支持更多中文字符。
//
// 参数:
//   - text: UTF-8 编码的文本
//
// 返回:
//   - []byte: GB18030 编码的字节集；转换失败返回空字节集
func Encoding_UTF8ToGB18030(text string) []byte {
	reader := transform.NewReader(strings.NewReader(text), simplifiedchinese.GB18030.NewEncoder())
	result, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return result
}

// Encoding_GB18030ToUTF8 将 GB18030 编码的字节集转换为 UTF-8 文本。
//
// 参数:
//   - data: GB18030 编码的字节集
//
// 返回:
//   - string: UTF-8 编码的文本；转换失败返回空串
func Encoding_GB18030ToUTF8(data []byte) string {
	reader := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewDecoder())
	result, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}
	return string(result)
}

// ============================================================
// Gzip 压缩
// ============================================================

// Encoding_GzipCompress 对字节集进行 Gzip 压缩。
//
// 参数:
//   - data: 待压缩的字节集
//
// 返回:
//   - []byte: Gzip 压缩后的字节集；压缩失败返回空字节集
func Encoding_GzipCompress(data []byte) []byte {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write(data)
	if err != nil {
		return []byte{}
	}
	writer.Close()
	return buf.Bytes()
}

// Encoding_GzipDecompress 对 Gzip 压缩的字节集进行解压。
//
// 参数:
//   - data: Gzip 压缩的字节集
//
// 返回:
//   - []byte: 解压后的字节集；解压失败返回空字节集
func Encoding_GzipDecompress(data []byte) []byte {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return []byte{}
	}
	defer reader.Close()
	result, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return result
}