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

	示例_核心库()
	示例_辅助()
	示例_编码()
	示例_校验()
	示例_文本()
	示例_文件()
	示例_时间()
	示例_数组()
	示例_正则()
	示例_字节集()
	示例_程序()
	示例_类型转换()
	示例_表格()
	示例_表达式计算()
	示例_json()
	示例_配置()
	示例_日志()
	示例_环境变量()
	示例_系统信息()
	示例_IP()
	示例_命令行()
	示例_对象池()
	示例_模板()
	示例_数据校验()
	示例_结构体合并()
	示例_日期解析()
	示例_队列()
	示例_临界许可()
	示例_读写锁()
	示例_正则表达式类()
	示例_队列泛型()
	示例_TCP()
	示例_WebSocket()
	示例_HTTP服务端()
}

func 示例_核心库() {
	fmt.Println("===== 核心库 =====")
	fmt.Println("D到整数:", D到整数("123"))
	fmt.Println("D到数值:", D到数值("3.14"))
	fmt.Println("D到文本:", D到文本(42))
	fmt.Println("D到字节集:", D到字节集("hello"))
	fmt.Println("G格式化文本:", G格式化文本("姓名:%s,年龄:%d", "张三", 25))
	fmt.Println()
}

func 示例_辅助() {
	fmt.Println("===== 辅助 =====")
	fmt.Println("选择(true):", F_选择(true, "是", "否"))
	fmt.Println("选择(false):", F_选择(false, "是", "否"))
	fmt.Println("取随机数:", F_取随机数(1, 100))
	fmt.Println("取文本右边:", F_取文本右边("Hello世界", 6))
	fmt.Println()
}

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

func 示例_数组() {
	fmt.Println("===== S数组 =====")
	fmt.Println("去重:", S数组_去重复([]string{"a", "b", "a", "c"}))
	fmt.Println("是否包含:", S数组_是否存在([]string{"a", "b", "c"}, "b"))
	fmt.Println("合并文本:", S数组_合并文本([]string{"a", "b", "c"}, ","))
	fmt.Println()
}

func 示例_正则() {
	fmt.Println("===== Z正则 =====")
	fmt.Println("取全部匹配子文本:", Z正则_取全部匹配子文本("abc123def456", `\d+`))
	fmt.Println()
}

func 示例_字节集() {
	fmt.Println("===== Z字节集 =====")
	fmt.Println("十六进制到字节集:", Z字节集_十六进制到字节集("48656c6c6f"))
	fmt.Println("字节集到十六进制:", Z字节集_字节集到十六进制([]byte("Hello")))
	fmt.Println()
}

func 示例_程序() {
	fmt.Println("===== C程序 =====")
	fmt.Println("取GUID:", C程序_取GUID())
	fmt.Println("取命令行:", C程序_取命令行())
	fmt.Println("取运行目录:", C程序_取运行目录())
	fmt.Println()
}

func 示例_类型转换() {
	fmt.Println("===== C类型转换 =====")
	fmt.Println("到文本:", C类型_到文本(123))
	fmt.Println("到整数:", C类型_到整数("456"))
	fmt.Println("到整数64:", C类型_到整数64("789"))
	fmt.Println("到浮点数:", C类型_到浮点数("3.14"))
	fmt.Println("到逻辑型:", C类型_到逻辑型("true"))
	fmt.Println()
}

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

func 示例_json() {
	fmt.Println("===== Jjson =====")
	json文本 := `{"name":"张三","age":25,"scores":{"math":90,"english":85}}`
	fmt.Println("取值(name):", Jjson_取值(json文本, "name"))
	fmt.Println("取值(scores.math):", Jjson_取值(json文本, "scores.math"))
	修改后, _ := Jjson_设置值(json文本, "age", "30")
	fmt.Println("设置值后:", Jjson_取值(修改后, "age"))
	fmt.Println()
}

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

func 示例_环境变量() {
	fmt.Println("===== K环境 =====")
	K环境_设置值("EFUNC_TEST", "hello")
	fmt.Println("取环境变量:", K环境_取值("EFUNC_TEST"))
	fmt.Println("是否存在:", K环境_是否存在("EFUNC_TEST"))
	fmt.Println()
}

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

func 示例_命令行() {
	fmt.Println("===== M命令行 =====")
	fmt.Println("取程序名:", M命令行_取程序名())
	fmt.Println("取剩余参数:", M命令行_取剩余参数())
	fmt.Println()
}

func 示例_对象池() {
	fmt.Println("===== P对象池 =====")
	buf := P对象池_获取()
	fmt.Println("获取缓冲区:", buf)
	P对象池_放回(buf)
	fmt.Println("缓冲区已放回")
	fmt.Println()
}

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
