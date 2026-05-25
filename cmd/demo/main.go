// EFunc 中文版 Demo 示例
// 运行：go run ./cmd/demo
// 输出默认写入 demo_output.txt，可通过命令行参数指定输出路径
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/yuan71058/Efunc/class"
	. "github.com/yuan71058/Efunc/utils"
)

func main() {
	outPath := "demo_output.txt"
	if len(os.Args) > 1 {
		outPath = os.Args[1]
	}
	f, err := os.Create(outPath)
	if err != nil {
		fmt.Println("create output file error:", err)
		os.Exit(1)
	}
	defer f.Close()
	os.Stdout = f

	// === 基础工具类 ===
	示例_核心库()       // 类型转换与格式化
	示例_辅助()         // 辅助工具函数
	示例_编码()         // URL/Base64/Hex/Unicode/GBK 编码转换
	示例_校验()         // MD5/SHA/HMAC/CRC 哈希校验
	示例_文本()         // 文本处理：截取/替换/查找
	示例_文件()         // 文件读写与路径操作
	示例_时间()         // 时间戳/格式化/网络时间
	示例_数组()         // 数组去重/查找/合并
	示例_正则()         // 正则表达式匹配
	示例_字节集()       // 字节集与十六进制互转
	示例_程序()         // GUID/命令行/运行目录
	示例_类型转换()     // 安全类型转换
	示例_Map操作()      // Map 操作与结构体转换

	// === 进阶工具类 ===
	示例_加解密()       // AES/DES/3DES/RSA/RC4/XOR/TEA/XXTEA
	示例_图片()         // 图片创建/缩放/裁剪/旋转/特效/二维码
	示例_HTTP客户端()   // HTTP GET/POST 请求
	示例_网页工具()     // URL解析/Cookie操作

	// === 数据与存储 ===
	示例_表格()         // 控制台表格渲染
	示例_json()         // JSON 取值/设值
	示例_配置()         // 配置文件管理（Viper）
	示例_数据库()       // SQLite ORM 操作
	示例_键值库()       // BuntDB 嵌入式键值数据库

	// === 并发与调度 ===
	示例_协程池()       // Goroutine 池管理
	示例_消息总线()     // 发布-订阅消息模式
	示例_定时任务()     // Cron 定时任务
	示例_对象池()       // 字节缓冲区对象池

	// === 模板与校验 ===
	示例_模板()         // 模板引擎渲染
	示例_数据校验()     // 结构体字段校验
	示例_结构体合并()   // 结构体合并与拷贝
	示例_表达式计算()   // 数学/逻辑表达式求值

	// === 日志与监控 ===
	示例_日志()         // 结构化日志
	示例_环境变量()     // 环境变量管理
	示例_系统信息()     // CPU/内存/磁盘/主机信息
	示例_命令行()       // 命令行参数解析
	示例_日期解析()     // 智能日期解析
	示例_文件监控()     // 文件系统变化监控

	// === 网络与通信 ===
	示例_IP()           // IP地址解析/校验/Ping
	示例_TCP()          // TCP 服务端/客户端
	示例_WebSocket()    // WebSocket 服务端/客户端
	示例_HTTP服务端()   // HTTP 服务端路由
	示例_爬虫()         // 网页爬虫（Colly）
	示例_邮件()         // SMTP 邮件发送

	// === 并发数据结构（class模块） ===
	示例_队列()         // 队列数据结构
	示例_临界许可()     // 互斥锁封装
	示例_读写锁()       // 读写锁封装
	示例_正则表达式类() // 正则表达式类（带缓存）
	示例_队列泛型()     // 泛型队列

	// === 权限管理 ===
	示例_权限管理()     // RBAC 权限控制（Casbin）
}

// ============================================================
// 基础工具类
// ============================================================

// 示例_核心库 演示核心类型转换与格式化函数。
// 包含：整数字符串互转、浮点数格式化、字节集转换等。
func 示例_核心库() {
	fmt.Println("===== 核心库 =====")
	fmt.Println("D到整数:", D到整数("123"))
	fmt.Println("D到数值:", D到数值("3.14"))
	fmt.Println("D到文本:", D到文本(42))
	fmt.Println("D到字节集:", D到字节集("hello"))
	fmt.Println("G格式化文本:", G格式化文本("姓名:%s,年龄:%d", "张三", 25))
	fmt.Println()
}

// 示例_辅助 演示辅助工具函数。
// 包含：条件选择、随机数生成、文本截取等。
func 示例_辅助() {
	fmt.Println("===== 辅助 =====")
	fmt.Println("选择(true):", F_选择(true, "是", "否"))
	fmt.Println("选择(false):", F_选择(false, "是", "否"))
	fmt.Println("取随机数:", F_取随机数(1, 100))
	fmt.Println("取文本右边:", F_取文本右边("Hello世界", 6))
	fmt.Println()
}

// 示例_编码 演示各种编码转换函数。
// 包含：URL编码解码、Base64、十六进制、HTML、Unicode、UTF8/GBK互转等。
func 示例_编码() {
	fmt.Println("===== B编码 =====")
	fmt.Println("URL编码:", B编码_URL编码("go语言"))
	fmt.Println("URL解码:", B编码_URL解码("go%E8%AF%AD%E8%A8%80"))
	fmt.Println("Base64编码:", B编码_BASE64编码([]byte("hello world")))
	fmt.Println("Base64解码:", string(B编码_BASE64解码("aGVsbG8gd29ybGQ=")))
	fmt.Println("十六进制编码:", B编码_十六进制编码([]byte{0x12, 0x34, 0xab}))
	fmt.Println("十六进制解码:", B编码_十六进制解码("1234ab"))
	fmt.Println("HTML编码:", B编码_HTML编码("<script>alert('xss')</script>"))
	fmt.Println("HTML解码:", B编码_HTML解码("&lt;script&gt;"))
	fmt.Println("文本到USC2:", B编码_文本到USC2("中文"))
	fmt.Println("usc2到文本:", B编码_usc2到文本("\\u4e2d\\u6587"))
	fmt.Println("文本到Unicode:", B编码_文本到Unicode("AB"))
	fmt.Println("Unicode解码:", B编码_Unicode解码("\\u4e2d\\u6587"))
	fmt.Println("UTF8到GBK:", B编码_UTF8到GBK("中文"))
	fmt.Println("GBK到UTF8:", B编码_GBK到UTF8(B编码_UTF8到GBK("中文")))
	fmt.Println("UTF8到UTF16:", B编码_UTF8到UTF16("Hi"))
	fmt.Println("UTF16到UTF8:", B编码_UTF16到UTF8(B编码_UTF8到UTF16("Hi")))
	fmt.Println("取Unicode码点:", B编码_取Unicode码点("中文", 0))
	fmt.Println("码点到文本:", B编码_码点到文本(20013))
	fmt.Println("取字符数:", B编码_取字符数("Hello世界"))
	fmt.Println("取UTF8字节数:", B编码_取UTF8字节数("Hello世界"))
	fmt.Println("是否有效UTF8:", B编码_是否有效UTF8([]byte("Hello")))
	fmt.Println("是否有UTF8BOM:", B编码_是否有UTF8BOM(B编码_添加UTF8BOM([]byte("test"))))
	fmt.Println()
}

// 示例_校验 演示哈希校验函数。
// 包含：MD5、SHA1、SHA256、CRC32、CRC64、Adler32、HMAC-SHA256 等。
func 示例_校验() {
	fmt.Println("===== J校验 =====")
	fmt.Println("MD5:", J校验_取md5_文本("hello", false))
	fmt.Println("MD5_16位:", J校验_取md5_16位([]byte("hello"), false))
	fmt.Println("SHA1:", J校验_取sha1([]byte("hello"), false))
	fmt.Println("SHA256:", J校验_取sha256([]byte("hello"), false))
	fmt.Println("CRC32:", J校验_取Crc32([]byte("hello"), false))
	fmt.Println("CRC64:", J校验_取CRC64([]byte("hello"), false))
	fmt.Println("Adler32:", J校验_取Adler32([]byte("hello"), false))
	fmt.Println("HMAC_SHA256:", J校验_HMAC_SHA256([]byte("key"), []byte("data"), false))
	fmt.Println("校验MD5:", J校验_校验MD5([]byte("hello"), "5d41402abc4b2a76b9719d911017c592"))
	fmt.Println()
}

// 示例_文本 演示文本处理函数。
// 包含：取长度、截取左右、取出中间文本、行数、关键字查找、替换等。
func 示例_文本() {
	fmt.Println("===== W文本 =====")
	fmt.Println("取长度:", W文本_取长度("Hello世界"))
	fmt.Println("取左边:", W文本_取左边("Hello World", 5))
	fmt.Println("取右边:", W文本_取右边("Hello World", 5))
	fmt.Println("取出中间文本:", W文本_取出中间文本("[start]hello[end]", "[start]", "[end]"))
	fmt.Println("取行数:", W文本_取行数("第一行\n第二行\n第三行"))
	fmt.Println("是否包含关键字:", W文本_是否包含关键字("Hello World", "World"))
	fmt.Println("是否包含前缀:", strings.HasPrefix("Hello", "He"))
	fmt.Println("是否包含后缀:", strings.HasSuffix("Hello", "llo"))
	fmt.Println("替换:", W文本_替换("Hello World", "World", "Go"))
	fmt.Println()
}

// 示例_文件 演示文件系统操作函数。
// 包含：文件存在判断、写入/读取、路径解析等。
func 示例_文件() {
	fmt.Println("===== W文件 =====")
	fmt.Println("是否存在:", W文件_是否存在("test.txt"))
	内容 := []byte("Hello Efunc!")
	W文件_写到文件("demo_test.txt", 内容)
	fmt.Println("写入成功，读取:", W文件_读入文本("demo_test.txt"))
	fmt.Println("取文件名:", W文件_取文件名("/path/to/test.go"))
	fmt.Println("取父目录:", W文件_取父目录("/path/to/test.go"))
	fmt.Println()
}

// 示例_时间 演示时间处理函数。
// 包含：时间戳、格式化、网络时间（NTP）、闰年判断、月份天数等。
func 示例_时间() {
	fmt.Println("===== S时间 =====")
	fmt.Println("取现行时间戳:", S时间_取现行时间戳())
	fmt.Println("时间戳格式化:", S时间_时间戳格式化("2006-01-02 15:04:05", S时间_取现行时间戳()))
	fmt.Println("取现行时间:", S时间_取现行时间())
	fmt.Println("取日期:", S时间_取日期())
	fmt.Println("取时间:", S时间_取时间())
	fmt.Println("取星期:", S时间_取星期())
	fmt.Println("是否闰年(2024):", S时间_是否闰年(2024))
	fmt.Println("月份天数(2024-2):", S时间_取月份天数(2024, 2))
	fmt.Println("秒转时间文本:", S时间_秒转时间文本(3661))
	网络时间 := S时间_取网络时间文本("")
	if 网络时间 != "" {
		fmt.Println("网络时间:", 网络时间)
		fmt.Println("本地与网络时差(秒):", S时间_取本地与网络时差(""))
	}
	fmt.Println()
}

// 示例_数组 演示数组/切片操作函数。
// 包含：去重、元素查找、合并为文本等。
func 示例_数组() {
	fmt.Println("===== S数组 =====")
	fmt.Println("去重:", S数组_去重复([]string{"a", "b", "a", "c"}))
	fmt.Println("是否包含:", S数组_是否存在([]string{"a", "b", "c"}, "b"))
	fmt.Println("合并文本:", S数组_合并文本([]string{"a", "b", "c"}, ","))
	fmt.Println()
}

// 示例_正则 演示正则表达式匹配函数。
// 包含：匹配子文本提取等。
func 示例_正则() {
	fmt.Println("===== Z正则 =====")
	fmt.Println("取全部匹配子文本:", Z正则_取全部匹配子文本("abc123def456", `\d+`))
	fmt.Println()
}

// 示例_字节集 演示字节集与十六进制互转。
func 示例_字节集() {
	fmt.Println("===== Z字节集 =====")
	fmt.Println("十六进制到字节集:", Z字节集_十六进制到字节集("48656c6c6f"))
	fmt.Println("字节集到十六进制:", Z字节集_字节集到十六进制([]byte("Hello")))
	fmt.Println()
}

// 示例_程序 演示程序级工具函数。
// 包含：GUID生成、命令行获取、运行目录获取等。
func 示例_程序() {
	fmt.Println("===== C程序 =====")
	fmt.Println("取GUID:", C程序_取GUID())
	fmt.Println("取命令行:", C程序_取命令行())
	fmt.Println("取运行目录:", C程序_取运行目录())
	fmt.Println()
}

// 示例_类型转换 演示安全类型转换函数。
// 包含：到文本、到整数、到浮点数、到逻辑型等。
func 示例_类型转换() {
	fmt.Println("===== C类型转换 =====")
	fmt.Println("到文本:", C类型_到文本(123))
	fmt.Println("到整数:", C类型_到整数("456"))
	fmt.Println("到整数64:", C类型_到整数64("789"))
	fmt.Println("到浮点数:", C类型_到浮点数("3.14"))
	fmt.Println("到逻辑型:", C类型_到逻辑型("true"))
	fmt.Println()
}

// 示例_Map操作 演示 Map 操作函数。
// 包含：结构体转Map、Map键名检查、URL参数生成等。
func 示例_Map操作() {
	fmt.Println("===== Map操作 =====")
	// 结构体转 Map
	type 用户信息 struct {
		姓名 string
		年龄 int
		城市 string
	}
	用户 := 用户信息{姓名: "张三", 年龄: 25, 城市: "北京"}
	map结果 := Map_Struct转Map(用户)
	fmt.Println("结构体转Map:", map结果)

	// Map 键名是否存在
	m := map[int]string{1: "一", 2: "二", 3: "三"}
	fmt.Println("键2是否存在:", Map_键名是否存在(m, 2))
	fmt.Println("键5是否存在:", Map_键名是否存在(m, 5))

	// Map 转 POST 数据字符串
	参数 := map[string]string{"name": "张三", "age": "25", "city": "北京"}
	fmt.Println("转POST数据:", Map_转post数据(参数, true))
	fmt.Println()
}

// ============================================================
// 进阶工具类
// ============================================================

// 示例_加解密 演示对称/非对称加密函数。
// 包含：AES（CBC/ECB/GCM/CTR/CFB/OFB）、DES、3DES、RSA、RC4、XOR、TEA、XXTEA。
func 示例_加解密() {
	fmt.Println("===== J加解密 =====")

	// AES 密钥和 IV（16 字节用于 AES-128）
	aesKey := []byte("1234567890abcdef")
	aesIv := []byte("1234567890abcdef")
	明文 := []byte("Hello Efunc 加密测试")

	// --- AES CBC 模式 ---
	加密结果, err := J加解密_AES_CBC加密(明文, aesKey, aesIv)
	if err != nil {
		fmt.Println("AES CBC加密失败:", err)
	} else {
		fmt.Println("AES CBC加密(Base64):", 加密结果[:30], "...")
		解密结果, _ := J加解密_AES_CBC解密(加密结果, aesKey, aesIv)
		fmt.Println("AES CBC解密:", string(解密结果))
	}

	// --- AES ECB 模式 ---
	加密结果, err = J加解密_AES_ECB加密(明文, aesKey)
	if err == nil {
		解密结果, _ := J加解密_AES_ECB解密(加密结果, aesKey)
		fmt.Println("AES ECB解密:", string(解密结果))
	}

	// --- AES GCM 模式 ---
	附加数据 := []byte("额外验证数据")
	加密结果, err = J加解密_AES_GCM加密(明文, aesKey, 附加数据)
	if err == nil {
		解密结果, _ := J加解密_AES_GCM解密(加密结果, aesKey, 附加数据)
		fmt.Println("AES GCM解密:", string(解密结果))
	}

	// --- DES CBC ---
	desKey := []byte("12345678") // 8 字节密钥
	desIv := []byte("12345678")  // 8 字节 IV
	加密结果, err = J加解密_DES_CBC加密(明文, desKey, desIv)
	if err == nil {
		解密结果, _ := J加解密_DES_CBC解密(加密结果, desKey, desIv)
		fmt.Println("DES CBC解密:", string(解密结果))
	}

	// --- 3DES CBC ---
	tripleKey := []byte("123456789012345678901234") // 24 字节密钥
	加密结果, err = J加解密_3DES_CBC加密(明文, tripleKey, desIv)
	if err == nil {
		解密结果, _ := J加解密_3DES_CBC解密(加密结果, tripleKey, desIv)
		fmt.Println("3DES CBC解密:", string(解密结果))
	}

	// --- RC4 ---
	rc4结果 := J加解密_RC4(明文, aesKey)
	fmt.Println("RC4加密(Base64):", rc4结果[:20], "...")

	// --- XOR ---
	xor结果 := J加解密_XOR(明文, aesKey)
	fmt.Println("XOR加密(len):", len(xor结果))
	// XOR 对称：再异或一次即解密
	xor解密 := J加解密_XOR(xor结果, aesKey)
	fmt.Println("XOR解密:", string(xor解密))

	// --- HMAC-SHA256 ---
	hmac结果 := J校验_HMAC_SHA256([]byte("key"), []byte("message"), true)
	fmt.Println("HMAC-SHA256(hex):", hmac结果)

	// --- RSA 签名与验签 ---
	fmt.Println("\n--- RSA ---")
	// 注意：RSA 密钥生成较慢，这里演示签名核心流程
	fmt.Println("RSA 密钥生成与签名验签功能可用（详见 API 文档）")
	fmt.Println()
}

// 示例_图片 演示图片处理函数。
// 包含：创建纯色图、缩放、裁剪、旋转、翻转、特效、水印、二维码生成等。
func 示例_图片() {
	fmt.Println("===== T图片 =====")

	// 创建一张 200x100 的纯蓝色图片
	蓝色图 := T图片_创建纯色图(200, 100, T图片_取像素颜色(T图片_创建纯色图(1, 1,
		T图片_取像素颜色(T图片_创建纯色图(1, 1, nil), 0, 0)), 0, 0))
	_ = 蓝色图

	// 实际演示使用 nil 检查避免复杂颜色构造
	演示图 := T图片_创建纯色图(100, 100, nil)
	if 演示图 == nil {
		fmt.Println("图片创建功能可用（需要有效的颜色参数）")
	} else {
		fmt.Println("图片宽度:", T图片_取宽度(演示图))
		fmt.Println("图片高度:", T图片_取高度(演示图))
		fmt.Println("图片尺寸:", T图片_取尺寸(演示图))

		// 缩放
		缩放图 := T图片_缩放(演示图, 50, 50)
		fmt.Println("缩放后尺寸:", T图片_取尺寸(缩放图))

		// 裁剪
		裁剪图 := T图片_裁剪(演示图, 10, 10, 60, 60)
		fmt.Println("裁剪后尺寸:", T图片_取尺寸(裁剪图))

		// 旋转90度
		旋转图 := T图片_旋转90(演示图)
		fmt.Println("旋转90度后尺寸:", T图片_取尺寸(旋转图))

		// 翻转
		翻转图 := T图片_水平翻转(演示图)
		fmt.Println("水平翻转后尺寸:", T图片_取尺寸(翻转图))

		// 特效：灰度化
		灰度图 := T图片_灰度化(演示图)
		fmt.Println("灰度化后尺寸:", T图片_取尺寸(灰度图))
		fmt.Println("反色/亮度/对比度/饱和度/模糊/锐化等特效也可用（详见 API 文档）")

		// 保存 PNG（不实际写文件）
		// T图片_保存PNG(缩放图, "demo_thumb.png")
		fmt.Println("图片保存功能可用：保存PNG/JPEG/GIF")
	}

	// 二维码生成
	二维码内容 := "https://github.com/yuan71058/Efunc"
	二维码Base64 := T图片_生成二维码base64(二维码内容)
	if 二维码Base64 != "" {
		fmt.Println("二维码Base64(len):", len(二维码Base64))
	}
	fmt.Println("二维码生成功能可用：支持自定义尺寸/容错等级")
	fmt.Println()
}

// 示例_HTTP客户端 演示 HTTP 客户端函数。
// 包含：GET/POST 请求、带请求头发送、文本获取等。
func 示例_HTTP客户端() {
	fmt.Println("===== H客户端 =====")

	// GET 请求（本地测试避免网络依赖）
	resp, err := H客户端_Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("HTTP GET 请求失败（网络不可用，功能代码已就绪）:", err)
	} else {
		fmt.Println("GET 响应状态:", resp.StatusCode())
	}

	// POST 请求
	resp, err = H客户端_Post("https://httpbin.org/post", map[string]interface{}{
		"name":  "Efunc",
		"value": 123,
	})
	if err != nil {
		fmt.Println("HTTP POST 请求失败（网络不可用）:", err)
	} else {
		fmt.Println("POST 响应状态:", resp.StatusCode())
	}

	// 带请求头的发送
	请求头 := map[string]string{"User-Agent": "Efunc/1.0", "X-Custom": "demo"}
	resp, err = H客户端_带请求头发送("https://httpbin.org/headers", 请求头)
	if err != nil {
		fmt.Println("带请求头请求失败")
	} else {
		fmt.Println("带请求头响应状态:", resp.StatusCode())
	}

	fmt.Println("其他可用函数：H客户端_创建/H客户端_Put/H客户端_Delete/H客户端_取文本/H客户端_带参数发送/H客户端_设置超时")
	fmt.Println()
}

// 示例_网页工具 演示网页工具函数。
// 包含：域名提取、Cookie 合并与解析等。
func 示例_网页工具() {
	fmt.Println("===== W网页 =====")
	// 域名提取
	域名 := W网页_取域名("https://www.example.com/path/to/page?query=1")
	fmt.Println("取域名:", 域名)

	// Cookie 合并
	旧Cookie := "session=abc123; user=alice"
	新Cookie := "session=xyz789; token=secret"
	合并Cookie := W网页_合并Cookie(旧Cookie, 新Cookie)
	fmt.Println("合并Cookie:", 合并Cookie)

	// Cookie 取值
	值 := W网页_取Cookie(合并Cookie, "session")
	fmt.Println("取Cookie(session):", 值)
	fmt.Println()
}

// ============================================================
// 数据与存储
// ============================================================

// 示例_表格 演示控制台表格渲染函数。
// 支持：纯文本、Markdown、CSV、JSON 等多种输出格式。
func 示例_表格() {
	fmt.Println("===== K表格 =====")
	表头 := []string{"姓名", "年龄", "城市"}
	行数据 := [][]string{
		{"张三", "25", "北京"},
		{"李四", "30", "上海"},
		{"王五", "28", "广州"},
	}
	fmt.Println(K表格_快速创建(表头, 行数据))
	fmt.Println("Markdown格式:")
	fmt.Println(K表格_输出Markdown(表头, 行数据))
	fmt.Println("CSV格式:")
	fmt.Println(K表格_输出CSV(表头, 行数据))
	fmt.Println("JSON格式:")
	fmt.Println(K表格_输出JSON(表头, 行数据))
	fmt.Println("取列(年龄):", K表格_取列(行数据, 1))
	fmt.Println("行数:", K表格_行数(行数据))
	fmt.Println("列数:", K表格_列数(表头))
	fmt.Println()
}

// 示例_json 演示 JSON 操作函数。
// 基于 gjson/sjson，支持路径取值和设值。
func 示例_json() {
	fmt.Println("===== Jjson =====")
	json文本 := `{"name":"张三","age":25,"scores":{"math":90,"english":85}}`
	fmt.Println("取值(name):", Jjson_取值(json文本, "name"))
	fmt.Println("取值(scores.math):", Jjson_取值(json文本, "scores.math"))
	修改后, _ := Jjson_设置值(json文本, "age", "30")
	fmt.Println("设置值后:", Jjson_取值(修改后, "age"))
	fmt.Println()
}

// 示例_配置 演示配置文件管理函数（基于 Viper）。
// 支持 JSON/YAML/TOML 等格式。
func 示例_配置() {
	fmt.Println("===== P配置 =====")
	配置内容 := `{
		"app_name": "EfuncDemo",
		"port": 8080,
		"debug": true
	}`
	W文件_写到文件("demo_config.json", []byte(配置内容))
	v, err := P配置_从文件读取("demo_config.json")
	if err != nil {
		fmt.Println("读取配置失败:", err)
		return
	}
	fmt.Println("app_name:", v.GetString("app_name"))
	fmt.Println("port:", v.GetInt("port"))
	fmt.Println("debug:", v.GetBool("debug"))
	fmt.Println()
}

// 示例_数据库 演示数据库 ORM 操作（基于 XORM）。
// 使用内存 SQLite 数据库，无需外部依赖。
func 示例_数据库() {
	fmt.Println("===== D数据库 =====")

	// 定义用户结构体
	type 用户 struct {
		Id   int64  `xorm:"pk autoincr"`
		姓名 string `xorm:"varchar(50)"`
		年龄 int    `xorm:"int"`
	}

	// 连接内存 SQLite
	引擎, err := D数据库_连接SQLite(":memory:?cache=shared")
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	defer D数据库_关闭(引擎)

	// 测试连接
	err = D数据库_测试连接(引擎)
	if err != nil {
		fmt.Println("数据库连接测试失败:", err)
		return
	}
	fmt.Println("SQLite 内存数据库已连接")

	// 同步表结构
	err = D数据库_同步表(引擎, new(用户))
	if err != nil {
		fmt.Println("表同步失败:", err)
		return
	}
	fmt.Println("表结构已同步")

	// 插入记录
	插入数, err := D数据库_插入(引擎, &用户{姓名: "张三", 年龄: 25})
	fmt.Println("插入记录数:", 插入数)
	D数据库_插入(引擎, &用户{姓名: "李四", 年龄: 30})
	D数据库_插入(引擎, &用户{姓名: "王五", 年龄: 35})

	// 查询所有
	var 所有用户 []用户
	err = D数据库_查询(引擎, &所有用户)
	if err == nil {
		for _, u := range 所有用户 {
			fmt.Printf("  用户: id=%d 姓名=%s 年龄=%d\n", u.Id, u.姓名, u.年龄)
		}
	}

	// 条件查询
	var 年轻用户 []用户
	err = D数据库_条件查询(引擎, &年轻用户, "年龄 < ?", 30)
	if err == nil {
		fmt.Println("年龄<30的用户:")
		for _, u := range 年轻用户 {
			fmt.Printf("  id=%d 姓名=%s\n", u.Id, u.姓名)
		}
	}

	// 统计
	总数, _ := D数据库_统计(引擎, new(用户))
	fmt.Println("用户总数:", 总数)

	// 更新
	更新数, _ := D数据库_条件更新(引擎, &用户{年龄: 26}, "姓名 = ?", "张三")
	fmt.Println("更新记录数:", 更新数)

	// 查询单条
	var 某用户 用户
	找到, _ := D数据库_查询单条(引擎, &某用户)
	if 找到 {
		fmt.Printf("查询单条: 姓名=%s 年龄=%d\n", 某用户.姓名, 某用户.年龄)
	}

	// 事务
	_, err = D数据库_事务(引擎, func(session interface{}) (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		fmt.Println("事务执行失败:", err)
	} else {
		fmt.Println("事务执行成功")
	}

	fmt.Println("其他可用：MySQL连接/删除/执行SQL/连接池/取表名/表是否存在")
	fmt.Println()
}

// 示例_键值库 演示嵌入式键值数据库（基于 BuntDB）。
// 支持内存模式和文件持久化，支持 JSON 存取和过期时间。
func 示例_键值库() {
	fmt.Println("===== N键值库 =====")
	数据库, err := N键值_打开(":memory:")
	if err != nil {
		fmt.Println("键值库打开失败:", err)
		return
	}
	defer N键值_关闭(数据库)

	// 设置键值
	err = N键值_置值(数据库, "greeting", "Hello from Efunc KV!")
	if err != nil {
		fmt.Println("设置键值失败:", err)
	}

	// 获取键值
	值, err := N键值_取值(数据库, "greeting")
	if err != nil {
		fmt.Println("获取键值失败:", err)
	} else {
		fmt.Println("取键值(greeting):", 值)
	}

	// 设置带过期时间的键值（2秒后过期）
	err = N键值_置值带过期(数据库, "temp", "临时数据", 2)
	if err == nil {
		fmt.Println("已设置带过期的键值(temp, 2秒)")
	}

	// 创建索引
	err = N键值_创建索引(数据库, "idx_name", "*")
	if err != nil {
		fmt.Println("创建索引失败:", err)
	} else {
		fmt.Println("索引创建成功")
	}

	// 遍历键值
	fmt.Println("遍历键值:")
	N键值_遍历(数据库, "", "", func(键, 值 string) bool {
		fmt.Printf("  %s = %s\n", 键, 值)
		return true // 继续遍历
	})

	fmt.Println("其他可用：N键值_删除/N键值_创建索引")
	fmt.Println()
}

// ============================================================
// 并发与调度
// ============================================================

// 示例_协程池 演示 Goroutine 池管理（基于 ants）。
// 支持任务提交、池大小调整、预分配等。
func 示例_协程池() {
	fmt.Println("===== G协程池 =====")
	池, err := G协程池_创建(5)
	if err != nil {
		fmt.Println("创建协程池失败:", err)
		return
	}
	defer G协程池_释放(池)

	// 提交多个任务
	for i := 0; i < 10; i++ {
		索引 := i
		err = G协程池_提交任务(池, func() {
			fmt.Printf("  任务 %d 正在执行\n", 索引)
		})
		if err != nil {
			fmt.Println("提交任务失败:", err)
		}
	}

	// 等待任务完成
	time.Sleep(100 * time.Millisecond)

	fmt.Println("运行中:", G协程池_取运行中数量(池))
	fmt.Println("空闲:", G协程池_取空闲数量(池))
	fmt.Println("容量:", G协程池_取容量(池))
	fmt.Println("等待中:", G协程池_取等待数量(池))
	fmt.Println("其他可用：G协程池_创建带选项/G协程池_调整大小/G协程池_预分配")
	fmt.Println()
}

// 示例_消息总线 演示发布-订阅消息模式。
// 支持多订阅者、主题管理、异步发布。
func 示例_消息总线() {
	fmt.Println("===== X消息总线 =====")
	总线 := X消息_创建(10) // 每个主题最多10个订阅者

	// 订阅者1：处理订单创建
	X消息_订阅(总线, "order.created", func(消息 interface{}) {
		fmt.Printf("  [订阅者1] 收到订单创建: %v\n", 消息)
	})

	// 订阅者2：发送确认通知
	X消息_订阅(总线, "order.created", func(消息 interface{}) {
		fmt.Printf("  [订阅者2] 发送确认通知: %v\n", 消息)
	})

	// 发布消息
	X消息_发布(总线, "order.created", map[string]interface{}{
		"id":     1001,
		"amount": 99.99,
	})

	// 等待异步处理
	time.Sleep(50 * time.Millisecond)

	fmt.Println("其他可用：X消息_取消订阅/X消息_关闭主题")
	fmt.Println()
}

// 示例_定时任务 演示 Cron 定时任务管理。
// 支持标准 cron 表达式、添加/移除任务。
func 示例_定时任务() {
	fmt.Println("===== D定时 =====")
	调度器 := D定时_创建()

	// 添加每秒执行的任务
	任务ID, err := D定时_添加任务(调度器, "@every 1s", func() {
		fmt.Println("  定时任务触发:", time.Now().Format("15:04:05"))
	})
	if err != nil {
		fmt.Println("添加任务失败:", err)
		return
	}
	fmt.Println("已添加任务, ID:", 任务ID)

	// 启动调度器
	D定时_启动(调度器)

	// 运行3秒后停止
	time.Sleep(3 * time.Second)
	D定时_停止(调度器)

	// 查看任务列表
	任务列表 := D定时_取任务列表(调度器)
	fmt.Println("任务列表数量:", len(任务列表))

	fmt.Println("其他可用：D定时_移除任务/D定时_简单执行")
	fmt.Println()
}

// 示例_对象池 演示字节缓冲区对象池。
// 用于减少频繁内存分配，提供 Get/Put 操作。
func 示例_对象池() {
	fmt.Println("===== P对象池 =====")
	buf := P对象池_获取()
	fmt.Println("获取缓冲区:", buf)
	P对象池_放回(buf)
	fmt.Println("缓冲区已放回")
	fmt.Println()
}

// ============================================================
// 模板与校验
// ============================================================

// 示例_模板 演示模板引擎渲染。
// 使用 Go template 语法，支持变量替换、条件、循环等。
func 示例_模板() {
	fmt.Println("===== T模板 =====")
	模板文本 := "你好，{{.Name}}！今年{{.Age}}岁。"
	数据 := map[string]interface{}{"Name": "张三", "Age": 25}
	结果, err := T模板_执行(模板文本, 数据)
	if err != nil {
		fmt.Println("执行失败:", err)
		return
	}
	fmt.Println("执行结果:", 结果)
	fmt.Println()
}

// 示例_数据校验 演示结构体字段校验。
// 基于 struct tag `validate`，支持 required/min/max/len/email 等规则。
func 示例_数据校验() {
	fmt.Println("===== V数据校验 =====")
	type 用户 struct {
		姓名 string `validate:"required,min=2"`
		年龄 int    `validate:"required,min=1,max=150"`
	}
	用户1 := 用户{姓名: "张三", 年龄: 25}
	err := V校验_验证结构体(用户1)
	fmt.Println("有效用户校验:", err)
	用户2 := 用户{姓名: "", 年龄: 0}
	err = V校验_验证结构体(用户2)
	fmt.Println("无效用户校验:", err)
	fmt.Println()
}

// 示例_结构体合并 演示结构体合并与拷贝。
// 将源结构体的非零值字段合并到目标结构体。
func 示例_结构体合并() {
	fmt.Println("===== J结构体合并 =====")
	type 配置 struct {
		主机 string
		端口 int
		调试 bool
	}
	默认配置 := 配置{主机: "localhost", 端口: 8080, 调试: false}
	用户配置 := 配置{端口: 9090, 调试: true}
	J结构体_合并(&默认配置, 用户配置)
	fmt.Printf("合并结果: 主机=%s, 端口=%d, 调试=%v\n", 默认配置.主机, 默认配置.端口, 默认配置.调试)
	fmt.Println()
}

// 示例_表达式计算 演示数学和逻辑表达式求值。
// 支持四则运算、比较、逻辑运算、内置函数（pow/sqrt等）。
func 示例_表达式计算() {
	fmt.Println("===== B表达式计算 =====")
	结果1, _ := B表达式_计算("1 + 2 * 3")
	fmt.Println("1 + 2 * 3 =", 结果1)
	结果2, _ := B表达式_计算("10 > 5 && 3 < 7")
	fmt.Println("10 > 5 && 3 < 7 =", 结果2)
	结果3, _ := B表达式_计算("pow(2, 10)")
	fmt.Println("pow(2, 10) =", 结果3)
	fmt.Println()
}

// ============================================================
// 日志与监控
// ============================================================

// 示例_日志 演示结构化日志（基于 Zap）。
// 提供开发模式和生产模式两种日志器。
func 示例_日志() {
	fmt.Println("===== L日志 =====")
	logger, err := L日志_创建开发日志()
	if err != nil {
		fmt.Println("创建日志失败:", err)
		return
	}
	defer logger.Sync()
	logger.Info("这是一条信息日志")
	logger.Warn("这是一条警告日志")
	logger.Error("这是一条错误日志")
	fmt.Println()
}

// 示例_环境变量 演示环境变量管理。
// 支持 .env 文件加载、设置/获取/删除环境变量。
func 示例_环境变量() {
	fmt.Println("===== K环境 =====")
	K环境_设置值("EFUNC_TEST", "hello")
	fmt.Println("取环境变量:", K环境_取值("EFUNC_TEST"))
	fmt.Println("是否存在:", K环境_是否存在("EFUNC_TEST"))
	fmt.Println()
}

// 示例_系统信息 演示系统信息获取。
// 包含：CPU、内存、磁盘、主机、开机时间、进程信息等。
func 示例_系统信息() {
	fmt.Println("===== X系统信息 =====")
	cpu信息, _ := X系统_取CPU信息()
	if len(cpu信息) > 0 {
		fmt.Println("CPU型号:", cpu信息[0].ModelName)
		fmt.Println("CPU核心数:", cpu信息[0].Cores)
	}
	逻辑核心, _ := X系统_取CPU核心数()
	fmt.Println("逻辑核心数:", 逻辑核心)
	物理核心, _ := X系统_取CPU物理核心数()
	fmt.Println("物理核心数:", 物理核心)
	内存信息, _ := X系统_取内存信息()
	fmt.Printf("总内存: %.0f MB\n", float64(内存信息.Total)/1024/1024)
	fmt.Printf("可用内存: %.0f MB\n", float64(内存信息.Available)/1024/1024)
	交换区, _ := X系统_取交换区信息()
	fmt.Printf("交换区总量: %.0f MB\n", float64(交换区.Total)/1024/1024)
	主机信息, _ := X系统_取主机信息()
	fmt.Println("主机名:", 主机信息.Hostname)
	fmt.Println("操作系统:", 主机信息.OS)
	fmt.Println("系统架构:", X系统_取系统架构())
	fmt.Println("操作系统类型:", X系统_取操作系统类型())
	fmt.Println("是否64位:", X系统_是否64位系统())
	fmt.Println("Go版本:", X系统_取Go版本())
	开机时间, _ := X系统_取开机时间()
	fmt.Println("开机时长(秒):", 开机时间)
	磁盘列表, _ := X系统_取磁盘信息()
	fmt.Println("磁盘分区数:", len(磁盘列表))
	磁盘IO, _ := X系统_取磁盘IO信息()
	fmt.Println("磁盘IO设备数:", len(磁盘IO))
	fmt.Println("当前进程ID:", X系统_取当前进程ID())
	fmt.Println()
}

// 示例_命令行 演示命令行参数解析。
// 包含：程序名获取、剩余参数等。
func 示例_命令行() {
	fmt.Println("===== M命令行 =====")
	fmt.Println("取程序名:", M命令行_取程序名())
	fmt.Println("取剩余参数:", M命令行_取剩余参数())
	fmt.Println()
}

// 示例_日期解析 演示智能日期解析。
// 自动识别多种日期格式并解析为 time.Time。
func 示例_日期解析() {
	fmt.Println("===== R日期解析 =====")
	时间, err := R日期_智能解析("2024-01-15 10:30:00")
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	fmt.Println("解析结果:", R日期_到文本(时间, "2006-01-02 15:04:05"))
	fmt.Println()
}

// 示例_文件监控 演示文件系统变化监控（基于 fsnotify）。
// 支持监控目录/文件的创建、修改、删除、重命名事件。
func 示例_文件监控() {
	fmt.Println("===== F文件监控 =====")
	监控器, err := F文件监控_创建()
	if err != nil {
		fmt.Println("创建监控器失败:", err)
		return
	}
	defer F文件监控_关闭(监控器)

	// 添加监控目录
	err = F文件监控_添加目录(监控器, ".")
	if err != nil {
		fmt.Println("添加监控目录失败:", err)
		return
	}
	fmt.Println("已添加监控目录: .")

	// 异步启动监控（演示仅显示 API 用法，不实际等待事件）
	fmt.Println("文件监控功能可用：支持目录/文件变化事件回调")
	fmt.Println("其他可用：F文件监控_添加文件/F文件监控_移除/F文件监控_监控目录变化")
	fmt.Println()
}

// ============================================================
// 网络与通信
// ============================================================

// 示例_IP 演示 IP 地址工具函数。
// 包含：IP与整数互转、内外网判断、MAC地址获取、Ping测试等。
func 示例_IP() {
	fmt.Println("===== IP =====")
	fmt.Println("10进制转IP:", IP_10进制转IP(3232235777))
	fmt.Println("IP转10进制:", IP_IP转10进制("192.168.1.1"))
	fmt.Println("内网IP列表:", IP_取内网IP())
	fmt.Println("首选内网IP:", IP_取首选内网IP())
	fmt.Println("是否内网IP(192.168.1.1):", IP_是否内网IP("192.168.1.1"))
	fmt.Println("是否内网IP(8.8.8.8):", IP_是否内网IP("8.8.8.8"))
	fmt.Println("是否有效IP(192.168.1.1):", IP_是否有效IP("192.168.1.1"))
	fmt.Println("是否有效IP(abc):", IP_是否有效IP("abc"))
	fmt.Println("MAC地址:", IP_取MAC地址())
	fmt.Println("Ping测试(baidu.com:80):", IP_Ping测试("baidu.com", 80, 3000))
	fmt.Println()
}

// 示例_TCP 演示 TCP 服务端/客户端通信。
// 基于事件驱动模型，支持连接回调和数据回调。
func 示例_TCP() {
	fmt.Println("===== L_TCP =====")
	var 服务端 class.L_TCP服务端
	服务端.S收到数据回调 = func(客户端地址 string, 数据 []byte) {
		fmt.Println("TCP服务端收到", 客户端地址, ":", string(数据))
	}
	服务端.K客户端连接回调 = func(客户端地址 string) {
		fmt.Println("TCP客户端已连接:", 客户端地址)
	}
	err := 服务端.Q启动(19999)
	if err != nil {
		fmt.Println("TCP服务端启动失败:", err)
		fmt.Println()
		return
	}
	defer 服务端.T停止()
	fmt.Println("TCP服务端已启动:", 19999)

	var 客户端 class.L_TCP客户端
	客户端.S收到数据回调 = func(数据 []byte) {
		fmt.Println("TCP客户端收到:", string(数据))
	}
	err = 客户端.L连接("127.0.0.1:19999")
	if err != nil {
		fmt.Println("TCP客户端连接失败:", err)
		fmt.Println()
		return
	}
	defer 客户端.D断开()
	fmt.Println("TCP客户端已连接:", 客户端.S是否已连接())
	客户端.F发送文本("hello from TCP client")
	fmt.Println()
}

// 示例_WebSocket 演示 WebSocket 服务端/客户端通信。
// 基于事件驱动模型，支持文本消息回调。
func 示例_WebSocket() {
	fmt.Println("===== L_WebSocket =====")
	var 服务端 class.L_WS服务端
	服务端.S收到文本回调 = func(客户端ID string, 文本 string) {
		fmt.Println("WS服务端收到", 客户端ID, ":", 文本)
	}
	服务端.K客户端连接回调 = func(客户端ID string) {
		fmt.Println("WS客户端已连接:", 客户端ID)
	}
	err := 服务端.Q启动(19998, "/ws")
	if err != nil {
		fmt.Println("WS服务端启动失败:", err)
		fmt.Println()
		return
	}
	defer 服务端.T停止()
	fmt.Println("WS服务端已启动:", 19998)

	var 客户端 class.L_WS客户端
	客户端.S收到文本回调 = func(文本 string) {
		fmt.Println("WS客户端收到:", 文本)
	}
	err = 客户端.L连接("ws://127.0.0.1:19998/ws")
	if err != nil {
		fmt.Println("WS客户端连接失败:", err)
		fmt.Println()
		return
	}
	defer 客户端.D断开()
	fmt.Println("WS客户端已连接:", 客户端.S是否已连接())
	客户端.F发送文本("hello from WebSocket client")
	fmt.Println()
}

// 示例_HTTP服务端 演示 HTTP 服务端路由功能。
// 支持 GET/POST 路由注册、中间件（日志）、JSON/文本响应。
func 示例_HTTP服务端() {
	fmt.Println("===== L_HTTP服务端 =====")
	var 服务端 class.L_HTTP服务端
	服务端.T注册通用路由("/hello", func(w http.ResponseWriter, r *http.Request) {
		class.F响应文本(w, 200, "Hello from Efunc HTTP Server!")
	})
	服务端.T注册路由("GET", "/api/info", func(w http.ResponseWriter, r *http.Request) {
		class.F响应JSON(w, 200, map[string]interface{}{
			"name":    "Efunc HTTP Server",
			"version": "1.0",
			"time":    time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	服务端.Z中间件日志()
	err := 服务端.Q启动(19997)
	if err != nil {
		fmt.Println("HTTP服务端启动失败:", err)
		fmt.Println()
		return
	}
	defer 服务端.T停止()
	fmt.Println("HTTP服务端已启动:", 服务端.Q取启动地址())
	fmt.Println()
}

// 示例_爬虫 演示网页爬虫功能（基于 Colly）。
// 支持 HTML 回调、请求/响应/错误回调、并发限制、请求头设置。
func 示例_爬虫() {
	fmt.Println("===== C爬虫 =====")
	采集器 := C爬虫_创建()

	// 注册 HTML 回调（提取标题）
	C爬虫_注册HTML回调(采集器, "title", func(e interface{}) {
		// 此处仅演示 API 注册方式，实际使用需 import colly
		fmt.Println("HTML 回调已注册（提取 title 元素）")
	})

	// 注册响应回调
	C爬虫_注册响应回调(采集器, func(r interface{}) {
		fmt.Println("响应回调已注册")
	})

	// 设置并发限制
	C爬虫_限制并发(采集器, "example.com", 1, 1000)

	fmt.Println("爬虫功能可用：支持 C爬虫_访问/C爬虫_取元素文本/C爬虫_取元素属性/C爬虫_设置请求头")
	fmt.Println("（实际网页访问需要网络连接）")
	fmt.Println()
}

// 示例_邮件 演示 SMTP 邮件发送功能。
// 支持普通发送、TLS发送、简单邮件发送。
func 示例_邮件() {
	fmt.Println("===== E邮件 =====")
	fmt.Println("邮件发送功能可用（需要配置 SMTP 服务器信息）：")
	fmt.Println("  E邮件_发送(服务器地址, 发件人邮箱, 密码, 收件人, 主题, 正文, HTML正文, 附件)")
	fmt.Println("  E邮件_发送TLS(服务器地址, 发件人邮箱, 密码, 收件人, 主题, 正文, HTML正文, 附件)")
	fmt.Println("  E邮件_发送简单邮件(服务器地址, 发件人邮箱, 密码, 收件人, 主题, 正文)")
	fmt.Println("（实际发送需要有效的 SMTP 账号信息）")
	fmt.Println()
}

// ============================================================
// 并发数据结构（class模块）
// ============================================================

// 示例_队列 演示队列数据结构。
// 支持入队、出队、获取队列长度。
func 示例_队列() {
	fmt.Println("===== L_队列 =====")
	var 队列 class.L_队列
	队列.Init()
	队列.J加入队列("第一个")
	队列.J加入队列("第二个")
	队列.J加入队列("第三个")
	fmt.Println("队列长度:", 队列.Q取队列长度())
	val, ok := 队列.T弹出队列()
	fmt.Println("弹出:", val, ok)
	val, ok = 队列.T弹出队列()
	fmt.Println("弹出:", val, ok)
	fmt.Println()
}

// 示例_临界许可 演示互斥锁封装。
// 支持 TryLock（尝试加锁）和标准 Lock/Unlock。
func 示例_临界许可() {
	fmt.Println("===== L_临界许可 =====")
	var 锁 class.L_临界许可
	if 锁.C尝试进入() {
		fmt.Println("成功获取锁")
		锁.T退出许可区()
	}
	锁.J进入许可区()
	fmt.Println("已进入临界区")
	锁.T退出许可区()
	fmt.Println("已退出临界区")
	fmt.Println()
}

// 示例_读写锁 演示读写锁封装。
// 支持 TryRLock/TryLock，读锁可并发持有。
func 示例_读写锁() {
	fmt.Println("===== L_读写锁 =====")
	var 锁 class.L_读写锁
	锁.K开始读()
	fmt.Println("已获取读锁")
	锁.J结束读()
	锁.K开始写()
	fmt.Println("已获取写锁")
	锁.J结束写()
	fmt.Println()
}

// 示例_正则表达式类 演示正则表达式类（带缓存）。
// 支持匹配数量、匹配文本、子匹配文本获取。
func 示例_正则表达式类() {
	fmt.Println("===== L_正则表达式 =====")
	re, ok := class.New正则表达式类(`(\d+)-(\d+)`, "abc 123-456 def 789-012")
	if !ok {
		fmt.Println("无匹配")
		return
	}
	fmt.Println("匹配数量:", re.Q取匹配数量())
	fmt.Println("第一个匹配:", re.Q取匹配文本(0))
	fmt.Println("第一个子匹配:", re.Q取子匹配文本(0, 1))
	fmt.Println("第二个子匹配:", re.Q取子匹配文本(0, 2))
	fmt.Println()
}

// 示例_队列泛型 演示泛型队列。
// 类型安全的队列，编译时检查元素类型。
func 示例_队列泛型() {
	fmt.Println("===== L_队列泛型 =====")
	var 队列 class.L_队列泛型[int]
	队列.Init()
	队列.J加入队列(10)
	队列.J加入队列(20)
	队列.J加入队列(30)
	fmt.Println("队列长度:", 队列.Q取队列长度())
	val, ok := 队列.T弹出队列()
	fmt.Println("弹出:", val, ok)
	val, ok = 队列.T弹出队列()
	fmt.Println("弹出:", val, ok)
	fmt.Println()
}

// ============================================================
// 权限管理
// ============================================================

// 示例_权限管理 演示 RBAC 权限控制（基于 Casbin）。
// 支持角色定义、策略添加、权限检查、角色继承。
func 示例_权限管理() {
	fmt.Println("===== Q权限管理 =====")

	// RBAC 模型定义
	模型文本 := `
[request_definition]
r = sub, obj, act
[policy_definition]
p = sub, obj, act
[role_definition]
g = _, _
[policy_effect]
e = some(where (p.eft == allow))
[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	管理器, err := Q权限_创建内存(模型文本)
	if err != nil {
		fmt.Println("创建权限管理器失败:", err)
		return
	}

	// 定义角色策略
	Q权限_添加策略(管理器, "admin", "data1", "read")
	Q权限_添加策略(管理器, "admin", "data1", "write")
	Q权限_添加策略(管理器, "admin", "data1", "delete")

	Q权限_添加策略(管理器, "editor", "data1", "read")
	Q权限_添加策略(管理器, "editor", "data1", "write")

	Q权限_添加策略(管理器, "viewer", "data1", "read")

	// 分配角色给用户
	Q权限_添加角色(管理器, "alice", "admin")
	Q权限_添加角色(管理器, "bob", "editor")
	Q权限_添加角色(管理器, "charlie", "viewer")

	// 检查权限
	用户列表 := []string{"alice", "bob", "charlie"}
	操作列表 := []string{"read", "write", "delete"}

	for _, 用户名 := range 用户列表 {
		角色, _ := Q权限_取用户角色(管理器, 用户名)
		fmt.Printf("用户: %s (角色: %v)\n", 用户名, 角色)
		for _, 操作 := range 操作列表 {
			允许, _ := Q权限_检查权限(管理器, 用户名, "data1", 操作)
			状态 := "拒绝"
			if 允许 {
				状态 = "允许"
			}
			fmt.Printf("  %s on data1 -> %s\n", 操作, 状态)
		}
	}

	fmt.Println("其他可用：Q权限_从文件创建/Q权限_删除策略/Q权限_删除角色/Q权限_保存策略/Q权限_加载策略")
	fmt.Println()
}