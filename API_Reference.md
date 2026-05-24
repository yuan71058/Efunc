# EFunc API 参考文档

> 版本：基于仓库最新源码生成 | Go 版本：1.22+ | 模块路径：`github.com/yuan71058/Efunc`

---

## 目录

- [一、class 模块（类定义）](#一class-模块)
  - [1.1 L_临界许可](#11-l_临界许可)
  - [1.2 L_读写锁](#12-l_读写锁)
  - [1.3 L_正则表达式](#13-l_正则表达式)
  - [1.4 L_队列](#14-l_队列)
  - [1.5 L_队列泛型](#15-l_队列泛型)
- [二、utils 模块（工具函数）](#二utils-模块)
  - [2.1 核心库](#21-核心库)
  - [2.2 辅助](#22-辅助)
  - [2.3 B编码](#23-b编码)
  - [2.4 C程序](#24-c程序)
  - [2.5 Float64转换](#25-float64转换)
  - [2.6 H汇编](#26-h汇编)
  - [2.7 IP](#27-ip)
  - [2.8 Int转换](#28-int转换)
  - [2.9 J校验](#29-j校验)
  - [2.10 Post数据类](#210-post数据类)
  - [2.11 Map](#211-map)
  - [2.12 M目录](#212-m目录)
  - [2.13 Rsa](#213-rsa)
  - [2.14 S数组](#214-s数组)
  - [2.15 S时间](#215-s时间)
  - [2.16 T图片](#216-t图片)
  - [2.17 W文件](#217-w文件)
  - [2.18 W文本](#218-w文本)
  - [2.19 W网页](#219-w网页)
  - [2.20 Y原子](#220-y原子)
  - [2.21 Z字节集](#221-z字节集)
  - [2.22 Z正则](#222-z正则)

---

# 一、class 模块

> 包路径：`github.com/yuan71058/Efunc/class` | 导入方式：`import "github.com/yuan71058/Efunc/class"` | 所有类均以 `L_` 前缀命名

---

## 1.1 L_临界许可

> 源文件：`class/类_临界许可.go` | 底层实现：`sync.Mutex` | 用途：互斥锁封装

### 结构体

```go
type L_临界许可 struct {
    lock sync.Mutex
}
```

### 方法

#### `J进入许可区`

进入许可区（加锁），阻塞直到获取锁。

```go
func (l *L_临界许可) J进入许可区()
```

---

#### `T退出许可区`

退出许可区（解锁）。

```go
func (l *L_临界许可) T退出许可区()
```

**注意**：必须在 `J进入许可区()` 之后调用，否则会引发 panic。

---

#### `C尝试进入`

尝试进入许可区（非阻塞加锁）。

```go
func (l *L_临界许可) C尝试进入() bool
```

**返回值**：`bool` — `true` = 成功获取锁；`false` = 锁已被占用

**示例**：

```go
锁 := class.L_临界许可{}
锁.J进入许可区()
// 临界区代码
锁.T退出许可区()

if 锁.C尝试进入() {
    // 成功进入
    锁.T退出许可区()
} else {
    // 锁被占用
}
```

---

## 1.2 L_读写锁

> 源文件：`class/类_读写锁.go` | 底层实现：`sync.RWMutex` | 用途：读共享、写独享

### 结构体

```go
type L_读写锁 struct {
    lock sync.RWMutex
}
```

### 方法

| 方法 | 签名 | 说明 |
|------|------|------|
| `K开始读` | `func (l *L_读写锁) K开始读()` | 获取读锁，多个读操作可并发 |
| `J结束读` | `func (l *L_读写锁) J结束读()` | 释放读锁 |
| `K开始写` | `func (l *L_读写锁) K开始写()` | 获取写锁，阻塞直到所有读写完成 |
| `J结束写` | `func (l *L_读写锁) J结束写()` | 释放写锁 |

**示例**：

```go
锁 := class.L_读写锁{}
锁.K开始读()
// 读取共享数据
锁.J结束读()

锁.K开始写()
// 修改共享数据
锁.J结束写()
```

---

## 1.3 L_正则表达式

> 源文件：`class/类_正则表达式.go` | 底层实现：`regexp.Regexp`

### 结构体

```go
type L_正则表达式 struct {
    Count          int      // 匹配数量
    SubmatchCount2 int      // 子匹配数量
}
```

### 构造函数

#### `New正则表达式类`

```go
func New正则表达式类(正则表达式文本 string, 被搜索的文本 string) (*L_正则表达式, bool)
```

| 参数 | 类型 | 说明 |
|------|------|------|
| 正则表达式文本 | `string` | 正则表达式字符串 |
| 被搜索的文本 | `string` | 待搜索的目标文本 |

**返回值**：`(*L_正则表达式, bool)` — 实例指针和是否匹配成功

---

### 方法

| 方法 | 签名 | 说明 |
|------|------|------|
| `E创建` | `func (this *L_正则表达式) E创建(正则表达式文本 string, 被搜索的文本 string) bool` | 初始化并执行匹配，返回是否匹配成功 |
| `Q取匹配数量` | `func (this *L_正则表达式) Q取匹配数量() int` | 返回匹配结果数量 |
| `Q取匹配文本` | `func (this *L_正则表达式) Q取匹配文本(匹配索引 int) string` | 取指定索引的完整匹配文本 |
| `Q取子匹配文本` | `func (this *L_正则表达式) Q取子匹配文本(匹配索引 int, 子表达式索引 int) string` | 取子匹配文本，越界返回空串 |
| `Q取子匹配数量` | `func (this *L_正则表达式) Q取子匹配数量() int` | 返回子匹配数量（含完整匹配） |
| `GetResult` | `func (this *L_正则表达式) GetResult() [][]string` | 返回原始匹配结果二维数组 |

**示例**：

```go
正则, ok := class.New正则表达式类(`(\d+)-(\d+)`, "abc 12-34 def 56-78")
if ok {
    fmt.Println(正则.Q取匹配数量())       // 2
    fmt.Println(正则.Q取匹配文本(0))      // "12-34"
    fmt.Println(正则.Q取子匹配文本(0, 1)) // "12"
    fmt.Println(正则.Q取子匹配文本(0, 2)) // "34"
}
```

---

## 1.4 L_队列

> 源文件：`class/类_队列.go` | 底层实现：`container/list` + `sync.Mutex` | 用途：FIFO 线程安全队列

### 结构体

```go
type L_队列 struct {}
```

### 方法

| 方法 | 签名 | 说明 |
|------|------|------|
| `Init` | `func (j *L_队列) Init()` | 初始化队列 |
| `J加入队列` | `func (q *L_队列) J加入队列(v interface{}) int` | 入队，返回队列长度 |
| `T弹出队列` | `func (q *L_队列) T弹出队列() (interface{}, bool)` | 弹出队尾元素；队列为空返回 `(nil, false)` |
| `T弹出队列文本` | `func (q *L_队列) T弹出队列文本(值 *string) bool` | 弹出文本元素；类型不匹配返回 `false` |
| `T弹出队列整数` | `func (q *L_队列) T弹出队列整数(值 *int) bool` | 弹出整数元素；类型不匹配返回 `false` |
| `Q取队列长度` | `func (q *L_队列) Q取队列长度() int` | 获取队列长度 |
| `Q清空队列` | `func (q *L_队列) Q清空队列() interface{}` | 清空队列 |
| `Dump` | `func (q *L_队列) Dump()` | 打印队列内容（调试用） |

**示例**：

```go
队列 := class.L_队列{}
队列.J加入队列("hello")
队列.J加入队列(42)

var 文本 string
if 队列.T弹出队列文本(&文本) {
    fmt.Println(文本) // "hello"
}
```

---

## 1.5 L_队列泛型

> 源文件：`class/类_队列泛型.go` | Go 版本要求：1.18+ | 用途：泛型线程安全队列

### 结构体

```go
type L_队列泛型[T any] struct {}
```

### 方法

| 方法 | 签名 | 说明 |
|------|------|------|
| `Init` | `func (q *L_队列泛型[T]) Init()` | 初始化队列 |
| `J加入队列` | `func (q *L_队列泛型[T]) J加入队列(v T) int` | 入队，返回队列长度 |
| `T弹出队列` | `func (q *L_队列泛型[T]) T弹出队列() (T, bool)` | 弹出队尾元素；空队列返回零值和 `false` |
| `Q取队列长度` | `func (q *L_队列泛型[T]) Q取队列长度() int` | 获取队列长度 |
| `Q清空队列` | `func (q *L_队列泛型[T]) Q清空队列()` | 清空队列 |
| `Dump` | `func (q *L_队列泛型[T]) Dump()` | 打印队列内容（调试用） |

**示例**：

```go
队列 := class.L_队列泛型[string]{}
队列.J加入队列("hello")
值, ok := 队列.T弹出队列() // "hello", true
```

---

# 二、utils 模块

> 包路径：`github.com/yuan71058/Efunc/utils` | 推荐导入：`. "github.com/yuan71058/Efunc/utils"` | 命名规则：拼音首字母 + 分类名 + 下划线 + 功能名

---

## 2.1 核心库

> 源文件：`utils/核心库.go` | 依赖：`github.com/gogf/gf/v2/util/gconv`

| 函数 | 签名 | 说明 |
|------|------|------|
| `D到字节集` | `func D到字节集(value interface{}) []byte` | 任意类型转字节集 |
| `D到字节` | `func D到字节(value interface{}) byte` | 任意类型转 byte |
| `D到整数` | `func D到整数(value interface{}) int` | 任意类型转 int |
| `D到整数64` | `func D到整数64(value interface{}) int64` | 任意类型转 int64 |
| `D到数值` | `func D到数值(value interface{}) float64` | 任意类型转 float64 |
| `D到文本` | `func D到文本(value interface{}) string` | 任意类型转 string |
| `D到结构体` | `func D到结构体(待转换的参数 interface{}, 结构体指针 interface{}) error` | 任意类型转结构体；失败返回 error |
| `S三元` | `func S三元[T any](value bool, string1, string2 T) T` | 泛型三元运算符 |
| `D多项选择` | `func D多项选择[T any](index int, arr []T, 默认值 T) T` | 泛型多项选择；索引越界返回默认值 |
| `G格式化文本` | `func G格式化文本(str string, 参数 ...interface{}) string` | fmt.Sprintf 封装 |
| `G格式化_JSON` | `func G格式化_JSON(data string) string` | JSON 缩进格式化；失败返回原字符串 |
| `D到文本数组` | `func D到文本数组(通用型变量 interface{}) []string` | 通用型变量转文本数组 |
| `S是否为数组` | `func S是否为数组(通用型变量 interface{}) bool` | 判断是否为数组/切片 |
| `W文本到utf8` | `func W文本到utf8(src string) string` | GBK → UTF-8；失败返回原字符串 |
| `Utf8到文本` | `func Utf8到文本(src string) string` | UTF-8 → GBK；失败返回原字符串 |
| `Q取随机数` | `func Q取随机数(min, max int) int` | [min, max] 范围随机整数 |

**示例**：

```go
D到文本(123)           // "123"
S三元(true, "是", "否") // "是"
D多项选择(1, []string{"a","b","c"}, "x") // "b"
```

---

## 2.2 辅助

> 源文件：`utils/辅助.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `选择` | `func 选择(逻辑 bool, 真返回参数, 假返回参数 interface{}) interface{}` | 三元选择（非泛型版） |
| `取随机数` | `func 取随机数(min, max int) int` | [min, max] 范围随机整数 |
| `取文本右边` | `func 取文本右边(text string, n int) string` | 取右侧 N 个字节字符；超过长度返回原文本 |
| `取文本左边` | `func 取文本左边(text string, n int) string` | 取左侧 N 个字节字符；超过长度返回原文本 |
| `加入成员` | `func 加入成员(数组 []string, 成员 string) []string` | 向数组追加成员 |
| `删首尾空` | `func 删首尾空(text string) string` | TrimSpace 封装 |
| `取文本长度` | `func 取文本长度(text string) int` | 取字节长度（中文占 3 字节） |
| `分割文本` | `func 分割文本(原文本 string, 分割符 string) []string` | 按分隔符分割文本 |

**注意**：`取文本右边`/`取文本左边` 按字节截取，中文可能截断。中文安全截取请用 `W文本_取右边`/`W文本_取左边`。

---

## 2.3 B编码

> 源文件：`utils/B编码.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_URL编码` | `func B编码_URL编码(欲编码的文本 string) string` | URL 编码 |
| `B编码_URL解码` | `func B编码_URL解码(URL string) string` | URL 解码；失败返回空串 |
| `B编码_usc2到文本` | `func B编码_usc2到文本(字符串 string) string` | USC2 转义转中文文本；失败返回空串 |
| `B编码_BASE64编码` | `func B编码_BASE64编码(字节集 []byte) string` | Base64 编码 |
| `B编码_BASE64解码` | `func B编码_BASE64解码(文本 string) []byte` | Base64 解码；失败返回空字节集 |

**示例**：

```go
B编码_URL编码("go语言")           // "go%E8%AF%AD%E8%A8%80"
B编码_usc2到文本("\\u4e2d\\u6587") // "中文"
B编码_BASE64编码([]byte("hello"))  // "aGVsbG8="
B编码_BASE64解码("aGVsbG8=")       // []byte("hello")
```

---

## 2.4 C程序

> 源文件：`utils/C程序.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `C程序_延时` | `func C程序_延时(毫秒数 int64) bool` | 毫秒级延时（time.Sleep）；始终返回 true |
| `C程序_延时2` | `func C程序_延时2(毫秒 int) bool` | 逐毫秒循环延时；效率较低，建议用 `C程序_延时` |
| `C程序_取cmd路径` | `func C程序_取cmd路径() string` | 获取 cmd.exe 路径（实现可能有问题） |
| `C程序_取GUID` | `func C程序_取GUID() string` | 生成 V4 GUID（如 `635897f8-2a48-4882-b3e1-823b8e5b6df8`） |
| `C程序_删除自身` | `func C程序_删除自身() error` | 删除当前可执行文件；Windows 下运行时可能失败 |
| `C程序_是否被调试` | `func C程序_是否被调试() bool` | 检测是否被调试（当前实现不完善） |
| `C程序_禁止重复运行` | `func C程序_禁止重复运行() bool` | 通过环境变量防重复运行；检测到重复时调用 `os.Exit(0)` |
| `C程序_写日志` | `func C程序_写日志(日志内容 string, 日志路径 string)` | 写日志文件（自动追加时间戳）；路径为空默认 `运行日志.txt` |
| `C程序_取命令行` | `func C程序_取命令行() []string` | 获取命令行参数（os.Args） |
| `C程序_取运行目录` | `func C程序_取运行目录() string` | 获取可执行文件目录（解析符号链接） |
| `C程序_取临时目录` | `func C程序_取临时目录() string` | 获取系统临时目录（依次检查 TEMP、TMP、%SYSTEMROOT%\Temp） |
| `C程序_运行Win` | `func C程序_运行Win(欲运行的命令行 string) string` | 执行 PowerShell 命令；输出自动 GBK→UTF-8 |

**日志格式**：`2024-01-15 10:30:45   日志内容`

---

## 2.5 Float64转换

> 源文件：`utils/Float64转换.go` | 依赖：`github.com/shopspring/decimal` | 所有运算防止精度丢失

| 函数 | 签名 | 说明 |
|------|------|------|
| `Float64取绝对值` | `func Float64取绝对值(值 float64) float64` | 高精度取绝对值 |
| `Float64乘int64` | `func Float64乘int64(值1 float64, 值2 int64) float64` | 高精度浮点×整数 |
| `Float64乘Float64` | `func Float64乘Float64(值1 float64, 值2 float64) float64` | 高精度浮点×浮点 |
| `Float64除int64` | `func Float64除int64(值1 float64, 值2 int64, 保留长度 int32) float64` | 高精度浮点÷整数，四舍五入 |
| `Float64除float64` | `func Float64除float64(值1 float64, 值2 float64, 保留长度 int32) float64` | 高精度浮点÷浮点，四舍五入 |
| `Float64取负值` | `func Float64取负值(值 float64) float64` | 正数取负，负数不变 |
| `Float64到文本` | `func Float64到文本(值 float64, 保留小数点多少位 int) string` | 浮点数转文本 |
| `Float64从文本` | `func Float64从文本(值 string, 保留小数点多少位 int) float64` | 文本转浮点数；失败返回 0 |
| `Int64到Float64` | `func Int64到Float64(值 int64) float64` | int64 转 float64（保留 2 位小数） |
| `Float64减float64` | `func Float64减float64(值1 float64, 值2 float64, 保留长度 int32) float64` | 高精度减法，四舍五入 |
| `Float64加float64` | `func Float64加float64(值1 float64, 值2 float64, 保留长度 int32) float64` | 高精度加法，四舍五入 |

**示例**：

```go
Float64到文本(3.14159, 2)      // "3.14"
Float64加float64(0.1, 0.2, 2)  // 0.3（不会出现 0.30000000000000004）
Float64除float64(1.0, 3.0, 4)  // 0.3333
```

---

## 2.6 H汇编

> 源文件：`utils/H汇编.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `H汇编_取随机数` | `func H汇编_取随机数(起始数, 结束数 int) int` | [起始数, 结束数] 范围随机整数；每次调用重置种子 |

**注意**：高频调用可能因种子相同导致结果重复。

---

## 2.7 IP

> 源文件：`utils/IP.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `IP_10进制转IP` | `func IP_10进制转IP(decimal int) string` | 10 进制整数转点分十进制 IP |

**示例**：`IP_10进制转IP(3232235777)` → `"192.168.1.1"`

---

## 2.8 Int转换

> 源文件：`utils/Int转换.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `Int取绝对值` | `func Int取绝对值(值 int) int` | 整数取绝对值 |
| `Int32ToBytes` | `func Int32ToBytes(i int32) []byte` | int32 转 4 字节大端字节集；失败返回空字节集 |

---

## 2.9 J校验

> 源文件：`utils/J校验.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `J校验_取md5` | `func J校验_取md5(字节集数据 []byte, 返回值转成大写 bool) string` | 计算 MD5（32 位 16 进制） |
| `J校验_取md5_文本` | `func J校验_取md5_文本(文本数据 string, 返回值转成大写 bool) string` | 文本 MD5 |
| `J校验_取Crc32` | `func J校验_取Crc32(数据 []byte, 返回值转成大写 bool) string` | 计算 CRC32（8 位 16 进制） |
| `J校验_取sha1` | `func J校验_取sha1(数据 []byte, 返回值转成大写 bool) string` | 计算 SHA1（40 位 16 进制） |
| `J校验_取sha256` | `func J校验_取sha256(数据 []byte, 返回值转成大写 bool) string` | 计算 SHA256（64 位 16 进制） |
| `J校验_取sha512` | `func J校验_取sha512(数据 []byte, 返回值转成大写 bool) string` | 计算 SHA512（128 位 16 进制） |

**示例**：

```go
J校验_取md5_文本("hello", false) // "5d41402abc4b2a76b9719d911017c592"
J校验_取md5_文本("hello", true)  // "5D41402ABC4B2A76B9719D911017C592"
```

---

## 2.10 Post数据类

> 源文件：`utils/L类_post数据类.go` | 用途：构造 HTTP POST 请求体和协议头

### 结构体

```go
type Post数据类 struct {}
```

### 方法

| 方法 | 签名 | 说明 |
|------|------|------|
| `T添加` | `func (p *Post数据类) T添加(key, value string, 转码 bool)` | 添加键值对；转码=true 时对 value URL 编码 |
| `T添加_批量` | `func (p *Post数据类) T添加_批量(文本 string, 转码 bool)` | 批量添加（`&` 分隔的 `key=value` 格式） |
| `Q取值` | `func (p *Post数据类) Q取值(key string) string` | 按 key 取值；未找到返回空串 |
| `Z置值` | `func (p *Post数据类) Z置值(key, value string)` | 设置/更新键值；不存在则新增 |
| `H获取Post数据` | `func (p *Post数据类) H获取Post数据(是否URL编码 bool) string` | 生成 POST 请求体 `key1=value1&key2=value2` |
| `H获取协议头数据` | `func (p *Post数据类) H获取协议头数据(是否URL编码 bool) string` | 生成协议头 `key1: value1\r\nkey2: value2` |
| `H获取Key数组` | `func (p *Post数据类) H获取Key数组() []string` | 获取所有键名 |
| `H获取Value数组` | `func (p *Post数据类) H获取Value数组() []string` | 获取所有值 |
| `H获取JSON文本` | `func (p *Post数据类) H获取JSON文本() string` | 生成 JSON `{"key1":"value1","key2":"value2"}` |
| `Q清空` | `func (p *Post数据类) Q清空()` | 清空所有键值对 |
| `S删除` | `func (p *Post数据类) S删除(key string)` | 按 key 删除键值对 |

**示例**：

```go
post := Post数据类{}
post.T添加("name", "张三", false)
post.T添加("age", "20", false)
post.H获取Post数据(false) // "name=张三&age=20"
post.H获取JSON文本()      // {"name":"张三","age":"20"}
```

---

## 2.11 Map

> 源文件：`utils/Map.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `Map_取key整数数组` | `func Map_取key整数数组(m map[int]string) []int` | 提取 int 类型 key 数组 |
| `Map_Struct转Map` | `func Map_Struct转Map(obj interface{}) map[string]interface{}` | 结构体转 Map（反射）；优先使用 `mapstructure` tag |
| `Map_转post数据` | `func Map_转post数据(URL参数 map[string]string, 是否url编码 bool) string` | Map 转 POST 参数字符串 |
| `Map_键名是否存在` | `func Map_键名是否存在(m map[int]string, key int) bool` | 判断 key 是否存在 |

---

## 2.12 M目录

> 源文件：`utils/M目录.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `M目录_是否存在` | `func M目录_是否存在(path string) (bool, error)` | 判断目录是否存在；存在同名文件时返回 error |
| `M目录_创建` | `func M目录_创建(路径 string) error` | 递归创建目录；已存在返回 nil |
| `M目录_枚举子目录` | `func M目录_枚举子目录(父文件夹路径 string, 子目录数组 *[]string, 是否带路径 bool, 是否继续向下枚举 bool) error` | 枚举子目录；可选递归、带路径 |
| `M目录_取运行目录` | `func M目录_取运行目录() string` | 获取可执行文件目录（`\` → `/`） |
| `M目录_取当前目录` | `func M目录_取当前目录() string` | 获取当前工作目录 |
| `M目录_删除` | `func M目录_删除(欲删除的目录名称 string) error` | 递归删除目录 |

---

## 2.13 Rsa

> 源文件：`utils/Rsa.go` | 用途：RSA 加解密与签名

| 函数 | 签名 | 说明 |
|------|------|------|
| `Rsa_私钥签名` | `func Rsa_私钥签名(base64后明文 string, RSA私钥 string) string` | MD5 with RSA 签名；返回 16 进制大写；失败返回空串 |
| `Rsa_GetKey` | `func Rsa_GetKey() (error, string, string)` | 生成 1024 位 RSA 密钥对；返回 (error, PKCS8公钥, PKCS1私钥) |
| `Rsa_私钥解密` | `func Rsa_私钥解密(Rsa私钥 []byte, 加密数据 []byte) string` | PKCS1v15 私钥解密；返回明文字符串 |
| `Rsa_私钥解密2` | `func Rsa_私钥解密2(Rsa私钥 []byte, 加密数据 []byte) []byte` | PKCS1v15 私钥解密；返回字节集 |
| `Rsa_公钥加密` | `func Rsa_公钥加密(公钥 string, 加密内容 []byte) string` | PKCS8 公钥加密；返回 Base64 密文 |
| `RSA_私钥加密` | `func RSA_私钥加密(Rsa私钥 []byte, 明文 []byte) string` | PKCS1 私钥加密（非标准）；返回 Base64 密文 |
| `RSA_公钥解密` | `func RSA_公钥解密(公钥 string, 密文 []byte) []byte` | PKCS1 公钥解密（非标准，对应私钥加密） |

**密钥格式注意**：
- 公钥加密（`Rsa_公钥加密`）使用 **PKCS8** 格式公钥（`ParsePKIXPublicKey`）
- 公钥解密（`RSA_公钥解密`）使用 **PKCS1** 格式公钥（`ParsePKCS1PublicKey`）
- 私钥操作统一使用 **PKCS1** 格式私钥（`ParsePKCS1PrivateKey`）

---

## 2.14 S数组

> 源文件：`utils/S数组.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `S数组_取随机成员` | `func S数组_取随机成员(源数组 []string, 数量 int) []string` | 随机取 N 个不重复成员；数量超过数组长度取全部 |
| `S数组_到文本` | `func S数组_到文本(array []interface{}) string` | `[]interface{}` 转逗号分隔文本 |
| `S数组_反转` | `func S数组_反转(反转的数组切片 []interface{})` | 原地反转数组 |
| `S数组_合并文本` | `func S数组_合并文本[T comparable](数组 []T, 连接字符 string) string` | 泛型数组合并为文本 |
| `S数组_取文本出现次数` | `func S数组_取文本出现次数(参数_数组 []string, 参数_成员 string) int` | 统计成员出现次数 |
| `S数组_取文本索引` | `func S数组_取文本索引(文本数组 []string, 文本 string) int` | 查找索引；未找到返回 -1 |
| `S数组_整数是否存在` | `func S数组_整数是否存在(数组 []int, 整数 int) bool` | 整数是否存在 |
| `S数组_是否存在` | `func S数组_是否存在[T comparable](数组 []T, 元素 T) bool` | 泛型判断元素是否存在 |
| `S数组_求平均值` | `func S数组_求平均值(参数 []int) int` | 整数数组求平均（整数除法） |
| `S数组_是否为空` | `func S数组_是否为空(list []string) bool` | 判断数组是否全为空串/空格 |
| `S数组_排序整数` | `func S数组_排序整数(arr []int) []int` | 升序排序（返回新数组） |
| `S数组_排序文本` | `func S数组_排序文本(arr []string) []string` | 字典序排序（返回新数组） |
| `S数组_去重复` | `func S数组_去重复[T comparable](数组 []T) []T` | 泛型去重（保持原始顺序） |
| `S数组_乱序` | `func S数组_乱序[T comparable](数组 []T) []T` | Fisher-Yates 乱序（返回新数组） |
| `S数组_整数取差集` | `func S数组_整数取差集(int1 []int, int2 []int) []int` | int2 有但 int1 没有的元素 |
| `S数组_取差集` | `func S数组_取差集(a, b []int) []int` | a 有但 b 没有的元素 |

**示例**：

```go
S数组_合并文本([]int{1, 2, 3}, ",")       // "1,2,3"
S数组_去重复([]string{"a","b","a"})         // ["a","b"]
S数组_是否存在([]string{"a","b"}, "b")      // true
```

---

## 2.15 S时间

> 源文件：`utils/S时间.go` | 默认格式：`2006-01-02 15:04:05`

| 函数 | 签名 | 说明 |
|------|------|------|
| `S时间_文本到时间戳` | `func S时间_文本到时间戳(时间文本 string) int` | 时间文本 → 10 位时间戳；失败返回 0 |
| `S时间_取现行时间戳13` | `func S时间_取现行时间戳13() int64` | 13 位毫秒级时间戳 |
| `S时间_取现行时间戳` | `func S时间_取现行时间戳() int64` | 10 位秒级时间戳 |
| `S时间_取现行时间` | `func S时间_取现行时间() string` | 当前时间文本 `2006-01-02 15:04:05` |
| `S时间_时间戳到时间` | `func S时间_时间戳到时间(时间戳 int64) string` | 10 位时间戳 → 时间文本 |
| `S时间_时间戳13到时间` | `func S时间_时间戳13到时间(时间戳 int64) string` | 13 位时间戳 → 时间文本 |
| `S时间_时间到时间戳` | `func S时间_时间到时间戳(时间 string) int64` | 时间文本 → 10 位时间戳；失败返回 0 |
| `S时间_时间戳格式化` | `func S时间_时间戳格式化(format string, 时间戳 int64) string` | 自定义格式化；时间戳=0 表示当前时间 |
| `S时间_秒转时间文本` | `func S时间_秒转时间文本(秒 int64) string` | 秒数转 `X年X月X天X时X分X秒`；仅显示非零项 |

**格式化占位符**（`S时间_时间戳格式化`）：

| 占位符 | 说明 | 占位符 | 说明 |
|--------|------|--------|------|
| `y` / `Y` | 年 | `h` | 时（12小时制） |
| `m` / `M` | 月 | `H` | 时（24小时制） |
| `d` / `D` | 日 | `i` | 分 |
| `s` | 秒 | `t`/`T` | 上下午 |

**示例**：

```go
S时间_秒转时间文本(3661) // "1时1分1秒"
S时间_时间戳格式化("Y-m-d H:i:s", 0) // 当前时间
```

---

## 2.16 T图片

> 源文件：`utils/T图片.go` | 依赖：`github.com/skip2/go-qrcode`

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_生成二维码base64` | `func T图片_生成二维码base64(内容 string) string` | 生成 256px 二维码并返回 Base64 编码；失败返回空串 |

**示例**：

```go
base64 := T图片_生成二维码base64("https://example.com")
// 可直接用于 <img src="data:image/png;base64,xxx">
```

---

## 2.17 W文件

> 源文件：`utils/W文件.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文件_是否存在` | `func W文件_是否存在(路径 string) bool` | 判断文件/目录是否存在 |
| `W文件_写到文件` | `func W文件_写到文件(文件名 string, 欲写入文件的数据 []byte) error` | 写入文件（自动创建父目录） |
| `W文件_枚举` | `func W文件_枚举(欲寻找的目录 string, 欲寻找的文件名 string, files *[]string, 是否带路径 bool, 是否遍历子目录 bool) error` | 枚举指定类型文件；文件名支持 `\|` 分隔多类型如 `.txt\|.jpg` |
| `W文件_取文件名` | `func W文件_取文件名(路径 string) string` | 从路径提取文件名 |
| `W文件_路径合并处理` | `func W文件_路径合并处理(elem ...string) string` | 路径拼接 |
| `W文件_取父目录` | `func W文件_取父目录(dirpath string) string` | 取父目录路径 |
| `W文件_删除` | `func W文件_删除(欲删除的文件名 string) error` | 删除文件 |
| `W文件_更名` | `func W文件_更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) error` | 重命名文件/目录 |
| `W文件_写出` | `func W文件_写出(文件名 string, 欲写入文件的数据 interface{}) error` | 写出文件（自动创建父目录） |
| `W文件_写出文件` | `func W文件_写出文件(文件名 string, 欲写入文件的数据 interface{}) error` | 同 `W文件_写出` |
| `W文件_追加文本` | `func W文件_追加文本(文件名 string, 欲追加的文本 string) error` | 追加文本到文件（自动加 `\r\n`） |
| `W文件_读入文本` | `func W文件_读入文本(文件名 string) string` | 读取文件文本内容 |
| `W文件_读入文件` | `func W文件_读入文件(文件名 string) []byte` | 读取文件字节集；失败返回空字节集 |
| `W文件_保存` | `func W文件_保存(文件名 string, 欲写入文件的数据 interface{}) error` | 智能保存（内容一致则跳过写出） |
| `W文件_取临时文件名` | `func W文件_取临时文件名(目录名 string) (f *os.File, filepath string, err error)` | 获取临时文件 |

**示例**：

```go
W文件_写出("./data/test.txt", "Hello World")
W文件_追加文本("./data/test.txt", "追加内容")
text := W文件_读入文本("./data/test.txt")
W文件_枚举("./data", ".txt|.csv", &files, true, true)
```

---

## 2.18 W文本

> 源文件：`utils/W文本.go` | 最大模块，提供全面的文本处理功能

### 查找与判断

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_是否包含关键字` | `func W文本_是否包含关键字(内容, 关键字 string) bool` | 是否包含关键字 |
| `W文本_是否存在` | `func W文本_是否存在(内容, 关键字 string) bool` | 是否包含文本 |
| `W文本_是否存在_任意` | `func W文本_是否存在_任意(内容 string, 关键字 []string) bool` | 是否包含任一关键字 |
| `W文本_是否存在_同时` | `func W文本_是否存在_同时(内容 string, 关键字 []string) bool` | 是否同时包含所有关键字 |
| `W文本_是否为英数字母` | `func W文本_是否为英数字母(s string) bool` | 是否全为英数 |
| `W文本_是否为字母` | `func W文本_是否为字母(s string) bool` | 是否全为字母 |
| `W文本_是否为数字` | `func W文本_是否为数字(s string) bool` | 是否全为数字 |
| `W文本_是否JSON` | `func W文本_是否JSON(s string) bool` | 是否为合法 JSON |
| `W文本_可能为json` | `func W文本_可能为json(内容 string) bool` | 高性能预判 JSON（仅检查首尾字符） |
| `W文本_寻找文本` | `func W文本_寻找文本(被搜寻的文本 string, 欲寻找的文本 string) int` | 查找位置；未找到返回 -1 |
| `W文本_倒找文本` | `func W文本_倒找文本(被搜寻的文本 string, 欲寻找的文本 string) int` | 从后查找位置；未找到返回 -1 |
| `W文本_寻找` | `func W文本_寻找(源文本, 要寻找的文本 string) int` | 同 `strings.Index` |
| `W文本_取出现次数` | `func W文本_取出现次数(被搜索文本 string, 欲搜索文本 string) int` | 统计文本出现次数 |
| `W文本_取文本所在行` | `func W文本_取文本所在行(源文本 string, 欲查找的文本 string, 是否区分大小写 bool) int` | 查找文本所在行号（0 起始） |

### 取文本

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_取出中间文本` | `func W文本_取出中间文本(内容 string, 左边文本 string, 右边文本 string) string` | 取左右标记之间的文本；未找到返回空串 |
| `W文本_取出中间文本_批量正则` | `func W文本_取出中间文本_批量正则(内容 string, 左边文本 string, 右边文本 string) []string` | 正则批量取中间文本 |
| `W文本_倒取出中间文本` | `func W文本_倒取出中间文本(欲取全文本 string, 右边文本 string, 左边文本 string, 倒数搜寻位置 int, 是否不区分大小写 bool) string` | 从后往前取中间文本 |
| `W文本_取文本左边` | `func W文本_取文本左边(内容 string, 关键字 string) string` | 取关键字左侧文本（不含关键字） |
| `W文本_取文本左边2` | `func W文本_取文本左边2(内容 string, 关键字 string) string` | 取关键字左侧文本（含关键字） |
| `W文本_取文本右边` | `func W文本_取文本右边(内容 string, 关键字 string) string` | 取关键字右侧文本（不含关键字） |
| `W文本_取文本右边_带关键字` | `func W文本_取文本右边_带关键字(内容 string, 关键字 string) string` | 取关键字右侧文本（含关键字） |
| `W文本_取文本右边2` | `func W文本_取文本右边2(被查找的文本 string, 欲寻找的文本 string, 起始寻找位置 int, 是否不区分大小写 bool) string` | 高级取右边文本 |
| `W文本_取左边` | `func W文本_取左边(欲取其部分的文本 string, 欲取出字符的数目 int) string` | 取左侧 N 个字符（中文算 1） |
| `W文本_取右边` | `func W文本_取右边(欲取其部分的文本 string, 欲取出字符的数目 int) string` | 取右侧 N 个字符（中文算 1） |
| `W文本_取指定变量文本行` | `func W文本_取指定变量文本行(文本 string, 行号 int) string` | 取指定行文本（1 起始） |

### 替换与删除

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_替换` | `func W文本_替换(源文本, 旧文本, 新文本 string) string` | 全局替换 |
| `W文本_替换2` | `func W文本_替换2(源文本 string, 替换内容 map[string]string) string` | 批量替换 |
| `W文本_子文本替换` | `func W文本_子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string` | 全局子文本替换 |
| `W文本_删除指定文本行` | `func W文本_删除指定文本行(源文本 string, 行数 int) string` | 删除指定行（0 起始） |
| `W文本_删除空行` | `func W文本_删除空行(要操作的文本 string) string` | 删除空行 |
| `W文本_删首尾空` | `func W文本_删首尾空(内容 string) string` | TrimSpace |
| `W文本_删首空` | `func W文本_删首空(欲删除空格的文本 string) string` | 去除左侧空格 |
| `W文本_删尾空` | `func W文本_删尾空(欲删除空格的文本 string) string` | 去除右侧空格 |
| `W文本_去重复文本` | `func W文本_去重复文本(原文本 string, 分割符 string) string` | 去除重复文本 |

### 转换与生成

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_到大写` | `func W文本_到大写(value string) string` | 转大写 |
| `W文本_到小写` | `func W文本_到小写(value string) string` | 转小写 |
| `W文本_首字母改大写` | `func W文本_首字母改大写(英文文本 string) string` | 首字母大写 |
| `W文本_颠倒` | `func W文本_颠倒(欲转换文本 string, 带有中文 bool) string` | 文本颠倒；中文模式按双字节处理 |
| `W文本_字符` | `func W文本_字符(字节型 int8) string` | 字节码转字符 |
| `W文本_取长度` | `func W文本_取长度(value string) int` | 取字符数（中文算 1） |
| `W文本_取行数` | `func W文本_取行数(文本 string) int` | 统计行数 |
| `W文本_取空白` | `func W文本_取空白(重复次数 int) string` | 生成 N 个空格 |
| `W文本_取重复` | `func W文本_取重复(重复次数 int, 待重复文本 string) string` | 重复文本 N 次 |
| `W文本_分割文本` | `func W文本_分割文本(待分割文本 string, 用作分割的文本 string) []string` | 分割文本 |
| `W文本_逐字分割` | `func W文本_逐字分割(原文本 string) []string` | 逐字分割为数组 |

### 随机生成

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_取随机字符串` | `func W文本_取随机字符串(字符串长度 int) string` | 随机字母数字串（首位不为 0） |
| `W文本_取随机字符串_数字` | `func W文本_取随机字符串_数字(字符串长度 int) string` | 随机纯数字串（首位不为 0） |
| `W文本_取随机数字数组` | `func W文本_取随机数字数组(最小值, 最大值 int, 数量 int) []string` | 不重复随机数字数组 |
| `W文本_取随机范围数字` | `func W文本_取随机范围数字(起始数, 结束数, 单双选择 int) string` | 随机数字（1=奇数, 2=偶数, 0=不限） |
| `W文本_取随机IP` | `func W文本_取随机IP() string` | 生成随机国内 IP |
| `W文本_取随机ip` | `func W文本_取随机ip() string` | 生成随机 IP（简单版） |

### 编码转换

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_gbk到utf8` | `func W文本_gbk到utf8(src string) string` | GBK → UTF-8（mahonia） |
| `W文本_utf8到gbk` | `func W文本_utf8到gbk(src string) string` | UTF-8 → GBK（mahonia） |

### 其他

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_去除敏感信息` | `func W文本_去除敏感信息(内容 string) string` | 中间字符替换为 `*`；长度 ≤ 2 时首字保留 |

**示例**：

```go
W文本_取出中间文本("<div>Hello</div>", "<div>", "</div>") // "Hello"
W文本_取长度("中文ab") // 4
W文本_去除敏感信息("张三丰") // "张*丰"
W文本_取随机字符串(8) // "aB3dE5fG"
```

---

## 2.19 W网页

> 源文件：`utils/W网页.go` | 用途：HTTP 网页访问与 Cookie 管理

### 函数

| 函数 | 签名 | 说明 |
|------|------|------|
| `W网页_取域名` | `func W网页_取域名(Url string) string` | 从 URL 提取域名 |
| `W网页_访问_对象` | `func W网页_访问_对象(网址 string, 访问方式 int, 提交信息 string, 提交Cookies string, 返回Cookies *string, 附加协议头 string, 返回协议头 *string, 返回状态代码 *int, 禁止重定向 bool, 字节集提交 []byte, 代理地址 string, 超时 int, 代理用户名 string, 代理密码 string, 代理标识 int, 对象继承 interface{}, 是否自动合并更新Cookie bool, 是否补全必要协议头 bool, 是否处理协议头大小写 bool) []byte` | 完整 HTTP 请求 |
| `网页_访问_对象` | `func 网页_访问_对象(网址 string, 访问方式 int, 提交信息 string, 提交Cookies string, 返回Cookies *string, 附加协议头 string, 返回协议头 *string, 返回状态代码 *int, 禁止重定向 bool, 字节集提交 []byte, 代理地址 string, 超时 int, 代理用户名 string, 代理密码 string, 代理标识 int, 对象继承 interface{}, 是否自动合并更新Cookie bool, 是否补全必要协议头 bool, 是否处理协议头大小写 bool) []byte` | 同上（无 W 前缀版本） |
| `Q取单条Cookie` | `func Q取单条Cookie(原Cookies, 单条Cookie名称 string) string` | 从 Cookie 字符串提取指定值 |
| `W网页_Cookie合并更新` | `func W网页_Cookie合并更新(旧Cookie, 新Cookie string) string` | 合并新旧 Cookie（去重、删除 `=deleted`） |
| `W网页_处理协议头` | `func W网页_处理协议头(原始协议头 string) string` | 规范化 HTTP 协议头格式（Title 化键名） |

### HTTP 访问方式枚举

| 值 | 方式 |
|----|------|
| 0 | GET |
| 1 | POST |
| 2 | HEAD |
| 3 | PUT |
| 4 | OPTIONS |
| 5 | DELETE |
| 6 | TRACE |
| 7 | CONNECT |

### `W网页_访问_对象` 参数详解

| 参数 | 类型 | 说明 |
|------|------|------|
| 网址 | `string` | 请求 URL |
| 访问方式 | `int` | 0=GET, 1=POST, 2=HEAD, 3=PUT, 4=OPTIONS, 5=DELETE, 6=TRACE, 7=CONNECT |
| 提交信息 | `string` | POST 提交数据 |
| 提交Cookies | `string` | 请求 Cookie |
| 返回Cookies | `*string` | 输出响应 Cookie |
| 附加协议头 | `string` | 附加 HTTP 头（`\n` 分隔） |
| 返回协议头 | `*string` | 输出 Content-Type |
| 返回状态代码 | `*int` | 输出 HTTP 状态码 |
| 禁止重定向 | `bool` | `true` = 不跟随重定向 |
| 字节集提交 | `[]byte` | 字节集形式的请求体 |
| 代理地址 | `string` | HTTP 代理地址 |
| 超时 | `int` | 超时秒数；-1=不限，<1=默认 15 秒 |
| 代理用户名 | `string` | 代理用户名（暂未实现） |
| 代理密码 | `string` | 代理密码（暂未实现） |
| 代理标识 | `int` | 代理标识（暂未实现） |
| 对象继承 | `interface{}` | 对象继承（暂未实现） |
| 是否自动合并更新Cookie | `bool` | 自动合并 Cookie |
| 是否补全必要协议头 | `bool` | 补全协议头 |
| 是否处理协议头大小写 | `bool` | 规范化协议头大小写 |

---

## 2.20 Y原子

> 源文件：`utils/Y原子.go` | 底层实现：`sync/atomic`

| 函数 | 签名 | 说明 |
|------|------|------|
| `Y原子_递增` | `func Y原子_递增(整数变量 *int64) int64` | 原子递增 1，返回新值 |
| `Y原子_递减` | `func Y原子_递减(整数变量 *int64) int64` | 原子递减 1，返回新值 |

**示例**：

```go
var 计数器 int64 = 0
Y原子_递增(&计数器) // 1
Y原子_递增(&计数器) // 2
Y原子_递减(&计数器) // 1
```

---

## 2.21 Z字节集

> 源文件：`utils/Z字节集.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `Z字节集_十六进制到字节集` | `func Z字节集_十六进制到字节集(原始16进制文本 string) []byte` | 16 进制文本转字节集 |
| `Z字节集_字节集到十六进制` | `func Z字节集_字节集到十六进制(字节集 []byte) string` | 字节集转 16 进制文本 |
| `Z字节集_寻找` | `func Z字节集_寻找(被搜寻的字节集 []byte, 欲寻找的字节集 []byte, 起始搜寻位置 int) int` | 字节集中搜索子字节集；1 起始索引；未找到返回 -1 |
| `Z字节集_Gzip解压` | `func Z字节集_Gzip解压(字节集 []byte) ([]byte, error)` | Gzip 解压；空数据返回错误 |

**示例**：

```go
hex := Z字节集_字节集到十六进制([]byte{0x48, 0x65}) // "4865"
data := Z字节集_十六进制到字节集("4865")              // []byte{0x48, 0x65}
pos := Z字节集_寻找([]byte{1,2,3,4}, []byte{3,4}, 1)  // 3
```

---

## 2.22 Z正则

> 源文件：`utils/Z正则.go` | 用途：正则校验与提取

### 校验类

| 函数 | 签名 | 说明 |
|------|------|------|
| `Z正则_校验密码` | `func Z正则_校验密码(s string, msg *string) bool` | 5-17 位非空白字符；失败时 msg 赋值 |
| `Z正则_校验代理用户名` | `func Z正则_校验代理用户名(s string, msg *string) bool` | 英文+数字+中文 |
| `Z正则_校验用户名` | `func Z正则_校验用户名(s string, msg *string) bool` | 5-17 位字母数字下划线 |
| `Z正则_校验email` | `func Z正则_校验email(s string, msg *string) bool` | 邮箱格式 |
| `Z正则_校验纯数字` | `func Z正则_校验纯数字(s string, msg *string) bool` | 纯数字（含小数和负数） |
| `Z正则_校验纯数字指定位数` | `func Z正则_校验纯数字指定位数(s string, msg *string, 位数 int) bool` | 指定位数纯数字 |
| `Z正则_是否英数` | `func Z正则_是否英数(s string, msg *string) bool` | 仅英文+数字 |

**说明**：所有校验函数的 `msg` 参数为输出参数，校验失败时赋值错误提示文本。

### 提取类

| 函数 | 签名 | 说明 |
|------|------|------|
| `Z正则_取Url连接地址` | `func Z正则_取Url连接地址(str string) []string` | 提取所有 URL |
| `Z正则_取全部匹配子文本` | `func Z正则_取全部匹配子文本(str, 正则表达式 string) []string` | 正则提取全部匹配 |
| `Z正则_取ip端口` | `func Z正则_取ip端口(str string) string` | 提取首个 IP:端口；未找到返回空串 |
| `Z正则_取ip端口多个` | `func Z正则_取ip端口多个(str string) []string` | 提取所有 IP:端口；未找到返回空数组 |

**示例**：

```go
var msg string
Z正则_校验email("test@example.com", &msg) // true
Z正则_校验email("invalid", &msg)           // false, msg="非正确email格式"

Z正则_取Url连接地址("访问 https://example.com 查看") // ["https://example.com"]
Z正则_取ip端口("代理 192.168.1.1:8080 连接")         // "192.168.1.1:8080"
```

---

## 附录：函数总数统计

| 模块 | 文件数 | 函数/方法数 |
|------|--------|------------|
| class | 5 | 28 |
| utils/核心库 | 1 | 16 |
| utils/辅助 | 1 | 8 |
| utils/B编码 | 1 | 5 |
| utils/C程序 | 1 | 12 |
| utils/Float64转换 | 1 | 11 |
| utils/H汇编 | 1 | 1 |
| utils/IP | 1 | 1 |
| utils/Int转换 | 1 | 2 |
| utils/J校验 | 1 | 6 |
| utils/Post数据类 | 1 | 11 |
| utils/Map | 1 | 4 |
| utils/M目录 | 1 | 6 |
| utils/Rsa | 1 | 7 |
| utils/S数组 | 1 | 16 |
| utils/S时间 | 1 | 9 |
| utils/T图片 | 1 | 1 |
| utils/W文件 | 1 | 15 |
| utils/W文本 | 1 | 42 |
| utils/W网页 | 1 | 6 |
| utils/Y原子 | 1 | 2 |
| utils/Z字节集 | 1 | 4 |
| utils/Z正则 | 1 | 11 |
| **合计** | **27** | **225** |
