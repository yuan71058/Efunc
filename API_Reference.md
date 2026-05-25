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
  - [1.6 L_TCP服务端 / L_TCP客户端](#16-l_tcp服务端--l_tcp客户端)
  - [1.7 L_WS服务端 / L_WS客户端](#17-l_ws服务端--l_ws客户端)
  - [1.8 L_HTTP服务端](#18-l_http服务端)
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
  - [2.23 Jjson](#223-jjsonjson-操作)
  - [2.24 C类型转换](#224-c类型转换安全类型转换)
  - [2.25 P配置](#225-p配置配置文件管理)
  - [2.26 E邮件](#226-e邮件邮件发送)
  - [2.27 X系统信息](#227-x系统信息系统指标获取)
  - [2.28 D定时](#228-d定时时任务管理)
  - [2.29 G协程池](#229-g协程池goroutine-池管理)
  - [2.30 L日志](#230-l日志高性能结构化日志)
  - [2.31 K环境变量](#231-k环境变量环境变量管理)
  - [2.32 M命令行](#232-m命令行命令行参数解析)
  - [2.33 R日期解析](#233-r日期解析智能日期解析)
  - [2.34 P对象池](#234-p对象池字节缓冲区对象池)
  - [2.35 B表达式计算](#235-b表达式计算数学逻辑表达式)
  - [2.36 T模板](#236-t模板高性能模板引擎)
  - [2.37 V数据校验](#237-v数据校验结构体校验)
  - [2.38 J结构体合并](#238-j结构体合并结构体合并与拷贝)
  - [2.39 K表格](#239-k表格控制台表格渲染)
  - [2.40 F文件监控](#240-f文件监控文件系统监控)
  - [2.41 X消息总线](#241-x消息总线发布订阅消息)
  - [2.42 H客户端](#242-h客户端http-客户端)
  - [2.43 N键值库](#243-n键值库嵌入式键值数据库)
  - [2.44 C爬虫](#244-c爬虫网页爬虫框架)
  - [2.45 Q权限管理](#245-q权限管理rbacabac权限)
  - [2.46 D数据库](#246-d数据库orm-数据库操作)
  - [2.47 C窗口](#247-c窗口windows-窗口操作)
  - [2.48 C进程](#248-c进程windows-进程管理)

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

## 1.6 L_TCP服务端 / L_TCP客户端

> 源文件：`class/类_TCP.go` | 依赖：标准库 | 用途：TCP 服务端与客户端，支持多连接并发、消息收发与广播

基于 Go 标准库 `net` 实现，使用 `bufio` 以换行符（`\n`）作为消息分隔符。每个客户端连接在独立 goroutine 中处理，通过回调函数通知上层。30 秒无数据交互自动断开。

**工作原理**：服务端 `Q启动` 后监听端口，客户端 `L连接` 后双方通过回调接收数据，发送数据时自动追加 `\n` 分隔符。

---

### L_TCP服务端

```go
type L_TCP服务端 struct {
    S收到数据回调  func(客户端地址 string, 数据 []byte)
    K客户端连接回调 func(客户端地址 string)
    K客户端断开回调 func(客户端地址 string)
}
```

| 方法 | 签名 | 说明 |
|------|------|------|
| `Q启动` | `func (s *L_TCP服务端) Q启动(端口 int) error` | 启动 TCP 服务端，监听指定端口（如 8888），内部自动启动 goroutine 处理 accept 循环 |
| `T停止` | `func (s *L_TCP服务端) T停止()` | 停止服务端，关闭 listener 和所有客户端连接 |
| `F发送数据` | `func (s *L_TCP服务端) F发送数据(客户端地址 string, 数据 []byte) error` | 向指定客户端发送字节数据，自动追加 `\n` 分隔符。客户端地址为连接时的 `conn.RemoteAddr().String()` |
| `F发送文本` | `func (s *L_TCP服务端) F发送文本(客户端地址 string, 文本 string) error` | 向指定客户端发送文本，等价于 `F发送数据(地址, []byte(文本))` |
| `G广播数据` | `func (s *L_TCP服务端) G广播数据(数据 []byte)` | 向所有已连接客户端广播字节数据 |
| `G广播文本` | `func (s *L_TCP服务端) G广播文本(文本 string)` | 向所有已连接客户端广播文本 |
| `Q取客户端数量` | `func (s *L_TCP服务端) Q取客户端数量() int` | 获取当前已连接的客户端数量 |
| `Q取客户端列表` | `func (s *L_TCP服务端) Q取客户端列表() []string` | 获取所有已连接客户端的地址列表（如 `["192.168.1.1:12345"]`） |

**回调字段**：所有回调均在独立 goroutine 中执行，内部可安全调用发送方法。

| 字段 | 类型 | 说明 |
|------|------|------|
| `S收到数据回调` | `func(客户端地址 string, 数据 []byte)` | 收到客户端数据时触发，数据已去除 `\n` 分隔符 |
| `K客户端连接回调` | `func(客户端地址 string)` | 新客户端连接建立时触发 |
| `K客户端断开回调` | `func(客户端地址 string)` | 客户端连接断开时触发 |

**示例**：

```go
// 创建 TCP 服务端
服务端 := &class.L_TCP服务端{}

// 注册回调
服务端.S收到数据回调 = func(客户端地址 string, 数据 []byte) {
    fmt.Println("收到", 客户端地址, ":", string(数据))
    服务端.F发送文本(客户端地址, "服务端已收到: "+string(数据))
}
服务端.K客户端连接回调 = func(客户端地址 string) {
    fmt.Println("新连接:", 客户端地址)
    服务端.F发送文本(客户端地址, "欢迎连接!")
}
服务端.K客户端断开回调 = func(客户端地址 string) {
    fmt.Println("断开:", 客户端地址)
}

// 启动服务端
if err := 服务端.Q启动(8888); err != nil {
    panic(err)
}
defer 服务端.T停止()

// 广播给所有客户端
服务端.G广播文本("服务器公告: xxx")
```

---

### L_TCP客户端

```go
type L_TCP客户端 struct {
    S收到数据回调  func(数据 []byte)
    D断开回调      func()
}
```

| 方法 | 签名 | 说明 |
|------|------|------|
| `L连接` | `func (c *L_TCP客户端) L连接(地址 string) error` | 连接 TCP 服务端，地址格式 `IP:端口`（如 `127.0.0.1:8888`），超时 10 秒 |
| `D断开` | `func (c *L_TCP客户端) D断开()` | 断开与服务端的连接 |
| `F发送数据` | `func (c *L_TCP客户端) F发送数据(数据 []byte) error` | 发送字节数据，自动追加 `\n` 分隔符 |
| `F发送文本` | `func (c *L_TCP客户端) F发送文本(文本 string) error` | 发送文本数据，等价于 `F发送数据([]byte(文本))` |
| `S是否已连接` | `func (c *L_TCP客户端) S是否已连接() bool` | 检查当前是否与服务端保持连接 |
| `Q取本地地址` | `func (c *L_TCP客户端) Q取本地地址() string` | 获取本端 socket 地址（通过 `LocalAddr()`） |

**回调字段**：

| 字段 | 类型 | 说明 |
|------|------|------|
| `S收到数据回调` | `func(数据 []byte)` | 收到服务端数据时触发，数据已去除 `\n` 分隔符 |
| `D断开回调` | `func()` | 连接断开（主动或被动）时触发 |

**示例**：

```go
// 创建客户端
客户端 := &class.L_TCP客户端{}

// 注册回调
客户端.S收到数据回调 = func(数据 []byte) {
    fmt.Println("收到:", string(数据))
}
客户端.D断开回调 = func() {
    fmt.Println("连接已断开")
}

// 连接服务端
if err := 客户端.L连接("127.0.0.1:8888"); err != nil {
    panic(err)
}
defer 客户端.D断开()

// 发送数据
客户端.F发送文本("hello server")
客户端.F发送数据([]byte{0x01, 0x02, 0x03})
```

---

## 1.7 L_WS服务端 / L_WS客户端

> 源文件：`class/类_WebSocket.go` | 依赖：`github.com/gorilla/websocket` | 用途：WebSocket 服务端与客户端，支持文本/二进制消息收发与广播

基于 `gorilla/websocket` 库实现的全双工 WebSocket 通信。服务端通过 HTTP 升级协议建立 WebSocket 连接，支持同时处理文本消息（`TextMessage`）和二进制消息（`BinaryMessage`）。

**与 TCP 的区别**：WebSocket 基于 HTTP 握手后升级为全双工通道，适合浏览器与服务器通信，支持文本/二进制/JSON 等多种消息格式，天然穿透防火墙和代理。

---

### L_WS服务端

```go
type L_WS服务端 struct {
    S收到文本回调  func(客户端ID string, 文本 string)
    S收到字节回调  func(客户端ID string, 数据 []byte)
    K客户端连接回调 func(客户端ID string)
    K客户端断开回调 func(客户端ID string)
}
```

| 方法 | 签名 | 说明 |
|------|------|------|
| `Q启动` | `func (s *L_WS服务端) Q启动(端口 int, 路径 string) error` | 启动 WebSocket 服务端。客户端通过 `ws://host:端口/路径` 连接。监听地址自动设为 `:端口` |
| `Q启动带地址` | `func (s *L_WS服务端) Q启动带地址(地址 string, 路径 string) error` | 启动服务端，可指定完整监听地址（如 `0.0.0.0:8080` 或 `127.0.0.1:8080`） |
| `T停止` | `func (s *L_WS服务端) T停止()` | 停止服务端，关闭 HTTP 服务器和所有 WebSocket 连接 |
| `F发送文本` | `func (s *L_WS服务端) F发送文本(客户端ID string, 文本 string) error` | 向指定客户端发送文本消息 |
| `F发送字节` | `func (s *L_WS服务端) F发送字节(客户端ID string, 数据 []byte) error` | 向指定客户端发送二进制消息 |
| `G广播文本` | `func (s *L_WS服务端) G广播文本(文本 string)` | 向所有已连接客户端广播文本消息 |
| `G广播字节` | `func (s *L_WS服务端) G广播字节(数据 []byte)` | 向所有已连接客户端广播二进制消息 |
| `Q取客户端数量` | `func (s *L_WS服务端) Q取客户端数量() int` | 获取当前已连接的客户端数量 |
| `Q取客户端列表` | `func (s *L_WS服务端) Q取客户端列表() []string` | 获取所有已连接客户端的 ID 列表（ID 为 `conn.RemoteAddr().String()`） |

**回调字段**：

| 字段 | 类型 | 说明 |
|------|------|------|
| `S收到文本回调` | `func(客户端ID string, 文本 string)` | 收到客户端文本消息时触发 |
| `S收到字节回调` | `func(客户端ID string, 数据 []byte)` | 收到客户端二进制消息时触发 |
| `K客户端连接回调` | `func(客户端ID string)` | 新 WebSocket 连接建立后触发 |
| `K客户端断开回调` | `func(客户端ID string)` | WebSocket 连接断开时触发 |

**示例**：

```go
// 创建 WebSocket 服务端
ws := &class.L_WS服务端{}

// 注册文本消息回调
ws.S收到文本回调 = func(客户端ID string, 文本 string) {
    fmt.Printf("收到文本 [%s]: %s\n", 客户端ID, 文本)
    ws.F发送文本(客户端ID, "回复: "+文本)
}

// 注册二进制消息回调
ws.S收到字节回调 = func(客户端ID string, 数据 []byte) {
    fmt.Printf("收到字节 [%s]: %d bytes\n", 客户端ID, len(数据))
}

// 注册连接/断开回调
ws.K客户端连接回调 = func(客户端ID string) {
    fmt.Println("WS连接:", 客户端ID)
    ws.G广播文本("用户 " + 客户端ID + " 加入了聊天室")
}
ws.K客户端断开回调 = func(客户端ID string) {
    fmt.Println("WS断开:", 客户端ID)
}

// 启动 WebSocket 服务端
if err := ws.Q启动(8080, "/ws"); err != nil {
    panic(err)
}
defer ws.T停止()
fmt.Println("WS服务端: ws://localhost:8080/ws")

// 广播消息
ws.G广播文本("系统通知: 服务已启动")
```

---

### L_WS客户端

```go
type L_WS客户端 struct {
    S收到文本回调  func(文本 string)
    S收到字节回调  func(数据 []byte)
    D断开回调      func()
}
```

| 方法 | 签名 | 说明 |
|------|------|------|
| `L连接` | `func (c *L_WS客户端) L连接(URL string) error` | 连接 WebSocket 服务端，URL 格式 `ws://host:port/path`（如 `ws://127.0.0.1:8080/ws`） |
| `D断开` | `func (c *L_WS客户端) D断开()` | 断开 WebSocket 连接 |
| `F发送文本` | `func (c *L_WS客户端) F发送文本(文本 string) error` | 发送文本消息 |
| `F发送字节` | `func (c *L_WS客户端) F发送字节(数据 []byte) error` | 发送二进制消息 |
| `F发送JSON` | `func (c *L_WS客户端) F发送JSON(数据 interface{}) error` | 发送 JSON 消息，使用 `conn.WriteJSON` 自动序列化。适合发送结构体数据 |
| `S是否已连接` | `func (c *L_WS客户端) S是否已连接() bool` | 检查当前是否与服务端保持连接 |
| `Q取远程地址` | `func (c *L_WS客户端) Q取远程地址() string` | 获取远程服务端地址（通过 `RemoteAddr()`） |

**回调字段**：

| 字段 | 类型 | 说明 |
|------|------|------|
| `S收到文本回调` | `func(文本 string)` | 收到服务端文本消息时触发 |
| `S收到字节回调` | `func(数据 []byte)` | 收到服务端二进制消息时触发 |
| `D断开回调` | `func()` | 连接断开时触发 |

**示例**：

```go
// 创建 WebSocket 客户端
客户端 := &class.L_WS客户端{}

// 注册回调
客户端.S收到文本回调 = func(文本 string) {
    fmt.Println("WS收到文本:", 文本)
}
客户端.S收到字节回调 = func(数据 []byte) {
    fmt.Println("WS收到字节:", len(数据), "bytes")
}
客户端.D断开回调 = func() {
    fmt.Println("WebSocket连接已断开")
}

// 连接服务端
if err := 客户端.L连接("ws://127.0.0.1:8080/ws"); err != nil {
    panic(err)
}
defer 客户端.D断开()

// 发送消息
客户端.F发送文本("hello ws")
客户端.F发送字节([]byte{0x01, 0x02, 0x03})
客户端.F发送JSON(map[string]string{"type": "ping"})
```

---

## 1.8 L_HTTP服务端

> 源文件：`class/类_HTTP.go` | 依赖：标准库 | 用途：HTTP 服务端，基于 `net/http` 封装，支持路由注册、中间件、静态文件服务和便捷响应函数

基于 Go 标准库 `net/http` 的 HTTP 服务端封装。提供中文 API 风格的服务器构建体验，支持按 HTTP Method 注册路由、全局中间件链、静态文件服务，以及一组便捷的响应和请求解析辅助函数。

**与 WebSocket 的区别**：HTTP 是短连接（请求-响应模式），每次请求独立；WebSocket 是长连接，适合实时双向通信。本模块两者互补使用。

---

### L_HTTP服务端

```go
type L_HTTP服务端 struct {}
```

| 方法 | 签名 | 说明 |
|------|------|------|
| `Q启动` | `func (s *L_HTTP服务端) Q启动(端口 int) error` | 启动 HTTP 服务端。监听地址自动设为 `:端口`（如 8080 → `:8080`）。无路由注册时返回 404 |
| `Q启动带地址` | `func (s *L_HTTP服务端) Q启动带地址(地址 string) error` | 启动服务端，指定完整监听地址。如 `0.0.0.0:8080`、`127.0.0.1:3000`。中间件在启动时应用 |
| `T停止` | `func (s *L_HTTP服务端) T停止() error` | 优雅停止服务端，等待正在处理的请求完成（最多 5 秒超时）。基于 `http.Server.Shutdown` |
| `T注册路由` | `func (s *L_HTTP服务端) T注册路由(方法 string, 路径 string, 处理函数 http.HandlerFunc)` | 注册 METHOD+路径路由。如 `T注册路由("GET", "/api/users", handler)`。方法区分大小写 |
| `T注册通用路由` | `func (s *L_HTTP服务端) T注册通用路由(路径 string, 处理函数 http.HandlerFunc)` | 注册不区分 HTTP 方法的通配路由，所有 Method 匹配该路径均触发。注意：与 Method 路由冲突时 Go 默认使用更精确的匹配 |
| `J静态文件服务` | `func (s *L_HTTP服务端) J静态文件服务(URL前缀 string, 本地目录 string)` | 将本地目录映射为静态文件服务。如 `J静态文件服务("/static/", "./public")` 将 `/static/index.html` 映射到 `./public/index.html` |
| `Z中间件` | `func (s *L_HTTP服务端) Z中间件(中间件函数 func(http.HandlerFunc) http.HandlerFunc)` | 添加全局中间件，按添加顺序包裹执行。中间件在 `Q启动` / `Q启动带地址` 时生效 |
| `Z中间件CORS` | `func (s *L_HTTP服务端) Z中间件CORS()` | 快捷添加 CORS 跨域中间件，允许所有来源访问。支持 GET/POST/PUT/DELETE/OPTIONS/PATCH 方法 |
| `Z中间件日志` | `func (s *L_HTTP服务端) Z中间件日志()` | 快捷添加请求日志中间件，打印方法、路径和处理耗时到控制台 |
| `Q取启动地址` | `func (s *L_HTTP服务端) Q取启动地址() string` | 获取服务端监听地址（如 `:8080`） |
| `S是否运行中` | `func (s *L_HTTP服务端) S是否运行中() bool` | 检查服务端是否正在运行 |

**中间件模式**：

```go
// 自定义中间件示例
func myAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            F响应JSON(w, 401, map[string]string{"error": "unauthorized"})
            return
        }
        next(w, r)
    }
}

服务端.Z中间件(myAuthMiddleware)
服务端.Z中间件CORS()
服务端.Z中间件日志()
```

---

### 响应辅助函数（class 包级别）

这些函数不从属于 `L_HTTP服务端` 实例，可直接作为 `class.函数名` 在路由处理器中使用。

| 函数 | 签名 | 说明 |
|------|------|------|
| `F响应JSON` | `func F响应JSON(w http.ResponseWriter, 状态码 int, 数据 interface{})` | 将数据 JSON 序列化后写入响应。自动设置 `Content-Type: application/json; charset=utf-8` |
| `F响应文本` | `func F响应文本(w http.ResponseWriter, 状态码 int, 文本 string)` | 写入纯文本响应。自动设置 `Content-Type: text/plain; charset=utf-8` |
| `F响应HTML` | `func F响应HTML(w http.ResponseWriter, 状态码 int, html string)` | 写入 HTML 响应。自动设置 `Content-Type: text/html; charset=utf-8` |
| `F取查询参数` | `func F取查询参数(r *http.Request, 参数名 string) string` | 从 URL Query String 获取参数值。如 `/api?name=张三` → `F取查询参数(r, "name")` 返回 `"张三"` |
| `F取POST参数` | `func F取POST参数(r *http.Request, 参数名 string) string` | 从 POST 表单获取参数值。内部自动调用 `r.ParseForm()` |
| `F解析JSON请求体` | `func F解析JSON请求体(r *http.Request, 目标 interface{}) error` | 从请求 Body 解析 JSON 到目标对象。使用 `json.NewDecoder` 流式解析 |

**完整示例**：

```go
// 创建 HTTP 服务端
server := &class.L_HTTP服务端{}

// 注册 GET 路由
server.T注册路由("GET", "/api/users", func(w http.ResponseWriter, r *http.Request) {
    users := []map[string]interface{}{
        {"id": 1, "name": "张三"},
        {"id": 2, "name": "李四"},
    }
    class.F响应JSON(w, 200, users)
})

// 注册 POST 路由
server.T注册路由("POST", "/api/echo", func(w http.ResponseWriter, r *http.Request) {
    var body map[string]interface{}
    if err := class.F解析JSON请求体(r, &body); err != nil {
        class.F响应JSON(w, 400, map[string]string{"error": err.Error()})
        return
    }
    class.F响应JSON(w, 200, body)
})

// 注册通用路由
server.T注册通用路由("/hello", func(w http.ResponseWriter, r *http.Request) {
    name := class.F取查询参数(r, "name")
    if name == "" {
        name = "World"
    }
    class.F响应HTML(w, 200, "<h1>Hello, "+name+"!</h1>")
})

// 静态文件服务
server.J静态文件服务("/static/", "./public")

// 添加中间件
server.Z中间件CORS()
server.Z中间件日志()

// 启动
if err := server.Q启动(8080); err != nil {
    panic(err)
}
defer server.T停止()
fmt.Println("HTTP服务端:", "http://localhost"+server.Q取启动地址())
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

> 源文件：`utils/F_helper.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `F_选择` | `func F_选择(逻辑 bool, 真返回参数, 假返回参数 interface{}) interface{}` | 三元选择（非泛型版） |
| `F_取随机数` | `func F_取随机数(min, max int) int` | [min, max] 范围随机整数 |
| `F_取文本右边` | `func F_取文本右边(text string, n int) string` | 取右侧 N 个字节字符；超过长度返回原文本 |
| `F_取文本左边` | `func F_取文本左边(text string, n int) string` | 取左侧 N 个字节字符；超过长度返回原文本 |
| `F_加入成员` | `func F_加入成员(数组 []string, 成员 string) []string` | 向数组追加成员 |
| `F_删首尾空` | `func F_删首尾空(text string) string` | TrimSpace 封装 |
| `F_取文本长度` | `func F_取文本长度(text string) int` | 取字节长度（中文占 3 字节） |
| `F_分割文本` | `func F_分割文本(原文本 string, 分割符 string) []string` | 按分隔符分割文本 |

**注意**：`F_取文本右边`/`F_取文本左边` 按字节截取，中文可能截断。中文安全截取请用 `W文本_取右边`/`W文本_取左边`。

---

## 2.3 B编码

> 源文件：`utils/B编码.go` | 依赖：`golang.org/x/net/idna`

### URL 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_URL编码` | `func B编码_URL编码(欲编码的文本 string) string` | URL 编码（百分号编码） |
| `B编码_URL解码` | `func B编码_URL解码(URL string) string` | URL 解码；失败返回空串 |
| `B编码_URL路径编码` | `func B编码_URL路径编码(路径 string) string` | URL 路径编码（保留 / & =） |
| `B编码_URL路径解码` | `func B编码_URL路径解码(路径 string) string` | URL 路径解码 |
| `B编码_URL组件编码` | `func B编码_URL组件编码(网址 string) string` | URL 组件编码 |

### Base64 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_BASE64编码` | `func B编码_BASE64编码(字节集 []byte) string` | 标准 Base64 编码 |
| `B编码_BASE64解码` | `func B编码_BASE64解码(文本 string) []byte` | Base64 解码；失败返回空字节集 |
| `B编码_BASE64URL编码` | `func B编码_BASE64URL编码(字节集 []byte) string` | URL 安全 Base64 编码（无填充） |
| `B编码_BASE64URL解码` | `func B编码_BASE64URL解码(文本 string) []byte` | URL 安全 Base64 解码 |
| `B编码_BASE64无填充编码` | `func B编码_BASE64无填充编码(字节集 []byte) string` | 无填充 Base64 编码 |
| `B编码_BASE64无填充解码` | `func B编码_BASE64无填充解码(文本 string) []byte` | 无填充 Base64 解码 |

### Base32 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_BASE32编码` | `func B编码_BASE32编码(字节集 []byte) string` | Base32 编码 |
| `B编码_BASE32解码` | `func B编码_BASE32解码(文本 string) []byte` | Base32 解码 |
| `B编码_BASE32HEX编码` | `func B编码_BASE32HEX编码(字节集 []byte) string` | Base32 Hex 编码 |
| `B编码_BASE32HEX解码` | `func B编码_BASE32HEX解码(文本 string) []byte` | Base32 Hex 解码 |

### Hex 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_十六进制编码` | `func B编码_十六进制编码(字节集 []byte) string` | 十六进制编码（小写） |
| `B编码_十六进制解码` | `func B编码_十六进制解码(文本 string) []byte` | 十六进制解码 |
| `B编码_十六进制大写` | `func B编码_十六进制大写(字节集 []byte) string` | 十六进制编码（大写） |

### Unicode / USC2 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_usc2到文本` | `func B编码_usc2到文本(字符串 string) string` | USC2 转义转中文文本 |
| `B编码_文本到USC2` | `func B编码_文本到USC2(文本 string) string` | 中文文本转 USC2 转义 |
| `B编码_文本到Unicode` | `func B编码_文本到Unicode(文本 string) string` | 文本转全 Unicode 转义 |

### HTML 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_HTML编码` | `func B编码_HTML编码(文本 string) string` | HTML 特殊字符转义 |
| `B编码_HTML解码` | `func B编码_HTML解码(文本 string) string` | HTML 实体解码 |

### Quoted-Printable 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_QP编码` | `func B编码_QP编码(字节集 []byte) string` | QP 编码（邮件传输） |
| `B编码_QP解码` | `func B编码_QP解码(文本 string) []byte` | QP 解码 |

### JSON 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_JSON编码` | `func B编码_JSON编码(值 interface{}) string` | JSON 编码 |
| `B编码_JSON编码缩进` | `func B编码_JSON编码缩进(值 interface{}, 前缀, 缩进 string) string` | JSON 编码（带缩进） |
| `B编码_JSON解码` | `func B编码_JSON解码(文本 string, 目标 interface{}) error` | JSON 解码 |

### MIME 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_MIME编码` | `func B编码_MIME编码(文本 string) string` | MIME QP 编码（邮件头部） |
| `B编码_MIMEB64编码` | `func B编码_MIMEB64编码(文本 string) string` | MIME Base64 编码 |

### 字节序编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_整数到大端` | `func B编码_整数到大端(值 interface{}) []byte` | 整数转大端字节序 |
| `B编码_整数到小端` | `func B编码_整数到小端(值 interface{}) []byte` | 整数转小端字节序 |
| `B编码_大端到整数` | `func B编码_大端到整数(字节集 []byte) uint64` | 大端字节序转整数 |
| `B编码_小端到整数` | `func B编码_小端到整数(字节集 []byte) uint64` | 小端字节序转整数 |

### Punycode 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_Punycode编码` | `func B编码_Punycode编码(域名 string) string` | 国际化域名 Punycode 编码 |
| `B编码_Punycode解码` | `func B编码_Punycode解码(域名 string) string` | Punycode 解码 |

### ANSI/GBK 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_UTF8到GBK` | `func B编码_UTF8到GBK(文本 string) []byte` | UTF-8 转 GBK（ANSI）编码 |
| `B编码_GBK到UTF8` | `func B编码_GBK到UTF8(数据 []byte) string` | GBK 转 UTF-8 |
| `B编码_UTF8到GB18030` | `func B编码_UTF8到GB18030(文本 string) []byte` | UTF-8 转 GB18030 编码 |
| `B编码_GB18030到UTF8` | `func B编码_GB18030到UTF8(数据 []byte) string` | GB18030 转 UTF-8 |

### UTF-16 编码

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_UTF8到UTF16` | `func B编码_UTF8到UTF16(文本 string) []byte` | UTF-8 转 UTF-16LE（含 BOM） |
| `B编码_UTF16到UTF8` | `func B编码_UTF16到UTF8(数据 []byte) string` | UTF-16 转 UTF-8（自动识别 BOM） |
| `B编码_UTF8到UTF16大端` | `func B编码_UTF8到UTF16大端(文本 string) []byte` | UTF-8 转 UTF-16BE（含 BOM） |

### UTF-8 BOM 处理

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_添加UTF8BOM` | `func B编码_添加UTF8BOM(数据 []byte) []byte` | 添加 UTF-8 BOM 头（EF BB BF） |
| `B编码_移除UTF8BOM` | `func B编码_移除UTF8BOM(数据 []byte) []byte` | 移除 UTF-8 BOM 头 |
| `B编码_是否有UTF8BOM` | `func B编码_是否有UTF8BOM(数据 []byte) bool` | 检查是否包含 UTF-8 BOM |

### Unicode 码点操作

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_Unicode解码` | `func B编码_Unicode解码(文本 string) string` | \uXXXX 格式解码为文本 |
| `B编码_取Unicode码点` | `func B编码_取Unicode码点(文本 string, 位置 int) int` | 获取指定位置字符的 Unicode 码点 |
| `B编码_码点到文本` | `func B编码_码点到文本(码点 int) string` | Unicode 码点转字符 |
| `B编码_取UTF8字节数` | `func B编码_取UTF8字节数(文本 string) int` | 获取 UTF-8 编码字节数 |
| `B编码_取字符数` | `func B编码_取字符数(文本 string) int` | 获取 Unicode 字符数（按 rune） |
| `B编码_是否有效UTF8` | `func B编码_是否有效UTF8(数据 []byte) bool` | 检查是否为有效 UTF-8 |

**示例**：

```go
B编码_URL编码("go语言")           // "go%E8%AF%AD%E8%A8%80"
B编码_usc2到文本("\\u4e2d\\u6587") // "中文"
B编码_文本到USC2("中文")           // "\\u4e2d\\u6587"
B编码_Unicode解码("\\u4e2d\\u6587") // "中文"
B编码_UTF8到GBK("中文")           // GBK 字节集
B编码_GBK到UTF8(gbk数据)          // "中文"
B编码_UTF8到UTF16("Hi")           // UTF-16LE 字节集（含 BOM）
B编码_取Unicode码点("中文", 0)     // 20013
B编码_码点到文本(20013)            // "中"
B编码_取字符数("Hello世界")        // 7
B编码_BASE64编码([]byte("hello"))  // "aGVsbG8="
B编码_Punycode编码("中文.com")     // "xn--fiq228c.com"
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

### IP 地址转换

| 函数 | 签名 | 说明 |
|------|------|------|
| `IP_10进制转IP` | `func IP_10进制转IP(decimal int) string` | 10 进制整数转点分十进制 IP |
| `IP_IP转10进制` | `func IP_IP转10进制(ip string) int` | 点分十进制 IP 转 10 进制整数 |

### 内网/外网 IP

| 函数 | 签名 | 说明 |
|------|------|------|
| `IP_取内网IP` | `func IP_取内网IP() []string` | 获取本机所有内网 IPv4 地址 |
| `IP_取首选内网IP` | `func IP_取首选内网IP() string` | 获取首选内网 IPv4 地址（跳过链路本地） |
| `IP_取外网IP` | `func IP_取外网IP() string` | 通过公共 API 获取公网 IPv4 地址 |
| `IP_取外网IP详细信息` | `func IP_取外网IP详细信息() string` | 获取外网 IP 的地理位置/ISP 等 JSON 信息 |

### IP 验证与判断

| 函数 | 签名 | 说明 |
|------|------|------|
| `IP_是否内网IP` | `func IP_是否内网IP(ip string) bool` | 判断是否为私有地址（10/172.16/192.168） |
| `IP_是否有效IP` | `func IP_是否有效IP(ip string) bool` | 判断字符串是否为有效 IPv4/IPv6 |

### MAC 地址与连通性

| 函数 | 签名 | 说明 |
|------|------|------|
| `IP_取MAC地址` | `func IP_取MAC地址() map[string]string` | 获取本机所有网络接口的 MAC 地址 |
| `IP_Ping测试` | `func IP_Ping测试(主机 string, 端口 int, 超时毫秒 int) bool` | TCP 连通性测试（非 ICMP） |

**示例**：

```go
IP_10进制转IP(3232235777)                    // "192.168.1.1"
IP_IP转10进制("192.168.1.1")                 // 3232235777
IP_取内网IP()                                 // ["192.168.1.100"]
IP_取首选内网IP()                             // "192.168.1.100"
IP_取外网IP()                                 // "123.45.67.89"
IP_是否内网IP("192.168.1.1")                  // true
IP_是否有效IP("abc")                          // false
IP_取MAC地址()                                // map["以太网":"4c:ed:fb:6a:ed:ae"]
IP_Ping测试("baidu.com", 80, 3000)            // true
```

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

### 哈希计算

| 函数 | 签名 | 说明 |
|------|------|------|
| `J校验_取md5` | `func J校验_取md5(字节集数据 []byte, 返回值转成大写 bool) string` | 计算 MD5（32 位 16 进制） |
| `J校验_取md5_文本` | `func J校验_取md5_文本(文本数据 string, 返回值转成大写 bool) string` | 文本 MD5 |
| `J校验_取md5_16位` | `func J校验_取md5_16位(字节集数据 []byte, 返回值转成大写 bool) string` | 16 位 MD5（取 32 位中间 16 位） |
| `J校验_取md5_文件` | `func J校验_取md5_文件(文件路径 string, 返回值转成大写 bool) (string, error)` | 文件 MD5（流式读取，支持大文件） |
| `J校验_取Crc32` | `func J校验_取Crc32(数据 []byte, 返回值转成大写 bool) string` | 计算 CRC32（8 位 16 进制） |
| `J校验_取Crc32_文件` | `func J校验_取Crc32_文件(文件路径 string, 返回值转成大写 bool) (string, error)` | 文件 CRC32 |
| `J校验_取CRC64` | `func J校验_取CRC64(数据 []byte, 返回值转成大写 bool) string` | 计算 CRC64（16 位 16 进制，ECMA 多项式） |
| `J校验_取Adler32` | `func J校验_取Adler32(数据 []byte, 返回值转成大写 bool) string` | 计算 Adler-32（8 位 16 进制） |
| `J校验_取sha1` | `func J校验_取sha1(数据 []byte, 返回值转成大写 bool) string` | 计算 SHA1（40 位 16 进制） |
| `J校验_取sha1_文件` | `func J校验_取sha1_文件(文件路径 string, 返回值转成大写 bool) (string, error)` | 文件 SHA1 |
| `J校验_取sha256` | `func J校验_取sha256(数据 []byte, 返回值转成大写 bool) string` | 计算 SHA256（64 位 16 进制） |
| `J校验_取sha256_文件` | `func J校验_取sha256_文件(文件路径 string, 返回值转成大写 bool) (string, error)` | 文件 SHA256 |
| `J校验_取sha512` | `func J校验_取sha512(数据 []byte, 返回值转成大写 bool) string` | 计算 SHA512（128 位 16 进制） |
| `J校验_取sha512_文件` | `func J校验_取sha512_文件(文件路径 string, 返回值转成大写 bool) (string, error)` | 文件 SHA512 |

### HMAC 消息认证

| 函数 | 签名 | 说明 |
|------|------|------|
| `J校验_HMAC_MD5` | `func J校验_HMAC_MD5(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string` | HMAC-MD5 消息认证码 |
| `J校验_HMAC_SHA1` | `func J校验_HMAC_SHA1(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string` | HMAC-SHA1 消息认证码 |
| `J校验_HMAC_SHA256` | `func J校验_HMAC_SHA256(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string` | HMAC-SHA256 消息认证码（API 签名常用） |
| `J校验_HMAC_SHA512` | `func J校验_HMAC_SHA512(密钥 []byte, 数据 []byte, 返回值转成大写 bool) string` | HMAC-SHA512 消息认证码 |

### 校验比对

| 函数 | 签名 | 说明 |
|------|------|------|
| `J校验_校验MD5` | `func J校验_校验MD5(数据 []byte, 预期值 string) bool` | 校验 MD5 是否与预期一致（不区分大小写） |
| `J校验_校验SHA256` | `func J校验_校验SHA256(数据 []byte, 预期值 string) bool` | 校验 SHA256 是否与预期一致 |
| `J校验_校验文件MD5` | `func J校验_校验文件MD5(文件路径 string, 预期值 string) (bool, error)` | 校验文件 MD5 |
| `J校验_校验文件SHA256` | `func J校验_校验文件SHA256(文件路径 string, 预期值 string) (bool, error)` | 校验文件 SHA256 |
| `J校验_校验HMAC` | `func J校验_校验HMAC(密钥 []byte, 数据 []byte, 预期值 string) bool` | 校验 HMAC-SHA256（防时序攻击） |

**示例**：

```go
J校验_取md5_文本("hello", false) // "5d41402abc4b2a76b9719d911017c592"
J校验_取md5_16位([]byte("hello"), false) // "4b2a76b9719d9110"
J校验_取md5_文件("test.txt", false) // 流式计算大文件 MD5
J校验_HMAC_SHA256([]byte("key"), []byte("data"), false) // API 签名
J校验_校验MD5([]byte("hello"), "5d41402abc4b2a76b9719d911017c592") // true
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

### 网络时间

> 支持 NTP 协议和 HTTP API 两种方式获取网络标准时间，用于校准本地时钟。

| 函数 | 签名 | 说明 |
|------|------|------|
| `S时间_取网络时间` | `func S时间_取网络时间(服务器 string) time.Time` | 通过 NTP 获取网络时间（默认 ntp.aliyun.com） |
| `S时间_取网络时间戳` | `func S时间_取网络时间戳(服务器 string) int64` | 通过 NTP 获取 10 位网络时间戳 |
| `S时间_取网络时间文本` | `func S时间_取网络时间文本(服务器 string) string` | 通过 NTP 获取格式化的网络时间文本 |
| `S时间_取HTTP网络时间` | `func S时间_取HTTP网络时间() time.Time` | 通过 HTTP API 获取网络时间（worldtimeapi.org） |

**示例**：

```go
nt := S时间_取网络时间("")                     // 从 ntp.aliyun.com 获取
S时间_取网络时间戳("ntp.ntsc.ac.cn")           // 从国家授时中心获取
S时间_取网络时间文本("time.windows.com")        // "2026-05-24 21:12:05"
httpt := S时间_取HTTP网络时间()                 // HTTP 方式（NTP 被封时备用）
```

---

## 2.16 T图片

> 源文件：`utils/T图片.go` | 依赖：`github.com/disintegration/imaging` + `github.com/skip2/go-qrcode`

### 读取与保存

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_读取` | `func T图片_读取(文件路径 string) (image.Image, error)` | 从文件读取图片（自动识别格式） |
| `T图片_读取Base64` | `func T图片_读取Base64(base64文本 string) (image.Image, error)` | 从 Base64 字符串读取图片 |
| `T图片_从字节读取` | `func T图片_从字节读取(数据 []byte) (image.Image, error)` | 从字节切片读取图片 |
| `T图片_从读取器读取` | `func T图片_从读取器读取(读取器 io.Reader) (image.Image, string, error)` | 从 io.Reader 读取图片 |
| `T图片_保存` | `func T图片_保存(图片 image.Image, 文件路径 string) error` | 保存图片（按扩展名自动选格式） |
| `T图片_保存PNG` | `func T图片_保存PNG(图片 image.Image, 文件路径 string) error` | 保存为 PNG 格式 |
| `T图片_保存JPEG` | `func T图片_保存JPEG(图片 image.Image, 文件路径 string, 质量 int) error` | 保存为 JPEG 格式（可设质量） |
| `T图片_保存GIF` | `func T图片_保存GIF(图片 image.Image, 文件路径 string) error` | 保存为 GIF 格式 |

### 信息获取

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_取宽度` | `func T图片_取宽度(图片 image.Image) int` | 获取图片宽度（像素） |
| `T图片_取高度` | `func T图片_取高度(图片 image.Image) int` | 获取图片高度（像素） |
| `T图片_取尺寸` | `func T图片_取尺寸(图片 image.Image) (int, int)` | 获取图片尺寸（宽, 高） |
| `T图片_取边界` | `func T图片_取边界(图片 image.Image) image.Rectangle` | 获取图片边界矩形 |
| `T图片_取像素颜色` | `func T图片_取像素颜色(图片 image.Image, x int, y int) color.Color` | 获取指定坐标像素颜色 |
| `T图片_取像素RGBA` | `func T图片_取像素RGBA(图片 image.Image, x int, y int) (uint32, uint32, uint32, uint32)` | 获取指定坐标 RGBA 分量 |
| `T图片_转Base64` | `func T图片_转Base64(图片 image.Image, 格式 string) (string, error)` | 转换为 Base64 编码字符串 |
| `T图片_转DataURI` | `func T图片_转DataURI(图片 image.Image, 格式 string) (string, error)` | 转换为 Data URI 格式 |
| `T图片_转字节` | `func T图片_转字节(图片 image.Image, 格式 string) ([]byte, error)` | 编码为字节切片 |

### 变换（缩放/裁剪/旋转/翻转）

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_缩放` | `func T图片_缩放(图片 image.Image, 宽度 int, 高度 int) image.Image` | 缩放到指定尺寸（支持等比） |
| `T图片_缩放到宽度` | `func T图片_缩放到宽度(图片 image.Image, 宽度 int) image.Image` | 等比缩放到指定宽度 |
| `T图片_缩放到高度` | `func T图片_缩放到高度(图片 image.Image, 高度 int) image.Image` | 等比缩放到指定高度 |
| `T图片_缩略图` | `func T图片_缩略图(图片 image.Image, 宽度 int, 高度 int) image.Image` | 生成缩略图（裁剪填充） |
| `T图片_裁剪` | `func T图片_裁剪(图片 image.Image, 左 int, 上 int, 右 int, 下 int) image.Image` | 裁剪到指定矩形区域 |
| `T图片_居中裁剪` | `func T图片_居中裁剪(图片 image.Image, 宽度 int, 高度 int) image.Image` | 从中心裁剪出指定尺寸 |
| `T图片_旋转` | `func T图片_旋转(图片 image.Image, 角度 float64) image.Image` | 顺时针旋转指定角度 |
| `T图片_旋转90` | `func T图片_旋转90(图片 image.Image) image.Image` | 顺时针旋转 90° |
| `T图片_旋转180` | `func T图片_旋转180(图片 image.Image) image.Image` | 旋转 180° |
| `T图片_旋转270` | `func T图片_旋转270(图片 image.Image) image.Image` | 顺时针旋转 270° |
| `T图片_水平翻转` | `func T图片_水平翻转(图片 image.Image) image.Image` | 左右镜像翻转 |
| `T图片_垂直翻转` | `func T图片_垂直翻转(图片 image.Image) image.Image` | 上下镜像翻转 |

### 效果（灰度/亮度/对比度/模糊/锐化）

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_灰度化` | `func T图片_灰度化(图片 image.Image) image.Image` | 转换为灰度图 |
| `T图片_反色` | `func T图片_反色(图片 image.Image) image.Image` | 颜色取反（底片效果） |
| `T图片_调整亮度` | `func T图片_调整亮度(图片 image.Image, 亮度 int) image.Image` | 调整亮度（-100~100） |
| `T图片_调整对比度` | `func T图片_调整对比度(图片 image.Image, 对比度 int) image.Image` | 调整对比度（-100~100） |
| `T图片_调整饱和度` | `func T图片_调整饱和度(图片 image.Image, 饱和度 int) image.Image` | 调整饱和度（-100~100） |
| `T图片_调整色相` | `func T图片_调整色相(图片 image.Image, 色相 int) image.Image` | 调整色相（-180~180） |
| `T图片_模糊` | `func T图片_模糊(图片 image.Image, 半径 float64) image.Image` | 高斯模糊 |
| `T图片_锐化` | `func T图片_锐化(图片 image.Image, 强度 float64) image.Image` | 锐化处理 |
| `T图片_伽马校正` | `func T图片_伽马校正(图片 image.Image, 伽马值 float64) image.Image` | 伽马校正 |

### 合成与水印

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_添加水印` | `func T图片_添加水印(底图, 水印图 image.Image, 位置 string, 偏移X, 偏移Y int, 透明度 float64) image.Image` | 添加水印（5 个位置可选） |
| `T图片_叠加` | `func T图片_叠加(底图, 上层图 image.Image, x, y int) image.Image` | 叠加图片 |
| `T图片_叠加带透明度` | `func T图片_叠加带透明度(底图, 上层图 image.Image, x, y int, 透明度 float64) image.Image` | 带透明度叠加 |
| `T图片_拼接水平` | `func T图片_拼接水平(图片列表 []image.Image) image.Image` | 水平拼接多张图片 |
| `T图片_拼接垂直` | `func T图片_拼接垂直(图片列表 []image.Image) image.Image` | 垂直拼接多张图片 |

### 创建

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_创建纯色图` | `func T图片_创建纯色图(宽度, 高度 int, 颜色值 color.Color) image.Image` | 创建纯色图片 |
| `T图片_创建透明图` | `func T图片_创建透明图(宽度, 高度 int) image.Image` | 创建透明图片 |
| `T图片_设置透明度` | `func T图片_设置透明度(图片 image.Image, 透明度 float64) image.Image` | 设置整体透明度 |

### 二维码

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_生成二维码` | `func T图片_生成二维码(内容 string, 文件路径 string, 尺寸 int) error` | 生成二维码并保存到文件 |
| `T图片_生成二维码base64` | `func T图片_生成二维码base64(内容 string) string` | 生成二维码并返回 Base64 |
| `T图片_生成二维码自定义` | `func T图片_生成二维码自定义(内容 string, 容错等级 int, 尺寸 int) (image.Image, error)` | 自定义参数生成二维码 |
| `T图片_生成二维码到写入器` | `func T图片_生成二维码到写入器(内容 string, 写入器 io.Writer, 尺寸 int) error` | 生成二维码写入 io.Writer |

**示例**：

```go
// 读取并缩放
img, _ := T图片_读取("photo.jpg")
thumb := T图片_缩略图(img, 200, 200)
T图片_保存JPEG(thumb, "thumb.jpg", 85)

// 添加水印
底图, _ := T图片_读取("bg.png")
水印, _ := T图片_读取("logo.png")
结果 := T图片_添加水印(底图, 水印, "右下", 10, 10, 0.5)
T图片_保存(结果, "output.png")

// 灰度化 + 旋转
灰度图 := T图片_灰度化(img)
旋转图 := T图片_旋转90(灰度图)

// 生成二维码
T图片_生成二维码("https://example.com", "qr.png", 256)
base64 := T图片_生成二维码base64("https://example.com")
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

## 2.23 Jjson（JSON 操作）

> 源文件：`utils/Jjson.go` | 依赖：`github.com/tidwall/gjson`、`github.com/tidwall/sjson` | 用途：JSON 路径取值、设置、删除

### 取值类

| 函数 | 签名 | 说明 |
|------|------|------|
| `Jjson_取值` | `func Jjson_取值(json文本 string, 路径 string) string` | 按路径获取字符串值 |
| `Jjson_取整数` | `func Jjson_取整数(json文本 string, 路径 string) int64` | 按路径获取整数值 |
| `Jjson_取浮点数` | `func Jjson_取浮点数(json文本 string, 路径 string) float64` | 按路径获取浮点数值 |
| `Jjson_取逻辑型` | `func Jjson_取逻辑型(json文本 string, 路径 string) bool` | 按路径获取布尔值 |
| `Jjson_取数组` | `func Jjson_取数组(json文本 string, 路径 string) []gjson.Result` | 按路径获取数组 |
| `Jjson_取对象` | `func Jjson_取对象(json文本 string, 路径 string) map[string]gjson.Result` | 按路径获取对象键值对 |
| `Jjson_是否存在` | `func Jjson_是否存在(json文本 string, 路径 string) bool` | 判断路径是否存在 |
| `Jjson_取数组长度` | `func Jjson_取数组长度(json文本 string, 路径 string) int` | 获取数组长度 |
| `Jjson_取所有路径` | `func Jjson_取所有路径(json文本 string) []string` | 获取 JSON 中所有叶节点路径 |

### 设置/删除类

| 函数 | 签名 | 说明 |
|------|------|------|
| `Jjson_设置值` | `func Jjson_设置值(json文本 string, 路径 string, 值 interface{}) (string, error)` | 设置指定路径的值 |
| `Jjson_设置文本值` | `func Jjson_设置文本值(json文本 string, 路径 string, 值 string) (string, error)` | 设置字符串值 |
| `Jjson_设置整数值` | `func Jjson_设置整数值(json文本 string, 路径 string, 值 int64) (string, error)` | 设置整数值 |
| `Jjson_删除值` | `func Jjson_删除值(json文本 string, 路径 string) (string, error)` | 删除指定路径 |

### 序列化/反序列化类

| 函数 | 签名 | 说明 |
|------|------|------|
| `Jjson_到文本` | `func Jjson_到文本(值 interface{}) (string, error)` | Go 值序列化为 JSON |
| `Jjson_到格式化文本` | `func Jjson_到格式化文本(值 interface{}, 缩进 string) (string, error)` | 序列化为带缩进的 JSON |
| `Jjson_解析` | `func Jjson_解析(json文本 string, 目标 interface{}) error` | JSON 反序列化到目标变量 |

**示例**：

```go
json := `{"user":{"name":"张三","age":25},"tags":["go","dev"]}`
Jjson_取值(json, "user.name")           // "张三"
Jjson_取整数(json, "user.age")           // 25
Jjson_取数组(json, "tags")               // [{go} {dev}]
Jjson_是否存在(json, "user.email")        // false

newJson, _ := Jjson_设置值(json, "user.email", "z@test.com")
newJson, _ = Jjson_删除值(newJson, "tags")
```

---

## 2.24 C类型转换（安全类型转换）

> 源文件：`utils/C类型转换.go` | 依赖：`github.com/spf13/cast` | 用途：安全的类型转换，支持默认值

### 基础转换

| 函数 | 签名 | 说明 |
|------|------|------|
| `C类型_到文本` | `func C类型_到文本(值 interface{}) string` | 任意类型转字符串 |
| `C类型_到整数` | `func C类型_到整数(值 interface{}) int` | 任意类型转 int |
| `C类型_到整数64` | `func C类型_到整数64(值 interface{}) int64` | 任意类型转 int64 |
| `C类型_到浮点数` | `func C类型_到浮点数(值 interface{}) float64` | 任意类型转 float64 |
| `C类型_到逻辑型` | `func C类型_到逻辑型(值 interface{}) bool` | 任意类型转 bool |
| `C类型_到文本切片` | `func C类型_到文本切片(值 interface{}) []string` | 任意类型转 []string |
| `C类型_到整数切片` | `func C类型_到整数切片(值 interface{}) []int` | 任意类型转 []int |
| `C类型_到时间` | `func C类型_到时间(值 interface{}) interface{}` | 任意类型转 time.Time |
| `C类型_到Duration` | `func C类型_到Duration(值 interface{}) interface{}` | 任意类型转 time.Duration |

### 安全转换（带默认值）

| 函数 | 签名 | 说明 |
|------|------|------|
| `C类型_安全到文本` | `func C类型_安全到文本(值 interface{}, 默认值 string) string` | 安全转字符串，失败返回默认值 |
| `C类型_安全到整数` | `func C类型_安全到整数(值 interface{}, 默认值 int) int` | 安全转 int，失败返回默认值 |
| `C类型_安全到浮点数` | `func C类型_安全到浮点数(值 interface{}, 默认值 float64) float64` | 安全转 float64，失败返回默认值 |
| `C类型_安全到逻辑型` | `func C类型_安全到逻辑型(值 interface{}, 默认值 bool) bool` | 安全转 bool，失败返回默认值 |

### 进制转换

| 函数 | 签名 | 说明 |
|------|------|------|
| `C类型_进制转换` | `func C类型_进制转换(文本 string, 进制 int) (int64, error)` | 按进制转换字符串为 int64（2/8/10/16） |

**示例**：

```go
C类型_到整数("123")              // 123
C类型_到逻辑型("true")           // true
C类型_安全到整数("abc", 0)       // 0（转换失败返回默认值）
C类型_进制转换("ff", 16)         // 255
C类型_进制转换("1010", 2)        // 10
```

---

## 2.25 P配置（配置文件管理）

> 源文件：`utils/P配置.go` | 依赖：`github.com/spf13/viper`、`github.com/fsnotify/fsnotify` | 用途：多格式配置文件读写与热更新

### 读取类

| 函数 | 签名 | 说明 |
|------|------|------|
| `P配置_从文件读取` | `func P配置_从文件读取(文件路径 string) (*viper.Viper, error)` | 从配置文件创建 viper 实例 |
| `P配置_从文件读取指定项` | `func P配置_从文件读取指定项(文件路径 string, 键名 string) (string, error)` | 从文件读取指定键值 |

### 取值类

| 函数 | 签名 | 说明 |
|------|------|------|
| `P配置_取值` | `func P配置_取值(v *viper.Viper, 键名 string) string` | 获取字符串值 |
| `P配置_取整数值` | `func P配置_取整数值(v *viper.Viper, 键名 string) int` | 获取整数值 |
| `P配置_取浮点数值` | `func P配置_取浮点数值(v *viper.Viper, 键名 string) float64` | 获取浮点数值 |
| `P配置_取逻辑值` | `func P配置_取逻辑值(v *viper.Viper, 键名 string) bool` | 获取布尔值 |
| `P配置_取字符串切片` | `func P配置_取字符串切片(v *viper.Viper, 键名 string) []string` | 获取字符串切片 |
| `P配置_取整数切片` | `func P配置_取整数切片(v *viper.Viper, 键名 string) []int` | 获取整数切片 |

### 写入/监听类

| 函数 | 签名 | 说明 |
|------|------|------|
| `P配置_设置值` | `func P配置_设置值(v *viper.Viper, 键名 string, 值 interface{})` | 设置配置值（需写回文件持久化） |
| `P配置_写回文件` | `func P配置_写回文件(v *viper.Viper) error` | 将配置写回文件 |
| `P配置_监听变更` | `func P配置_监听变更(v *viper.Viper)` | 监听配置文件变更 |
| `P配置_变更回调` | `func P配置_变更回调(v *viper.Viper, 回调 func())` | 注册变更回调函数 |

### 环境变量绑定类

| 函数 | 签名 | 说明 |
|------|------|------|
| `P配置_绑定环境变量` | `func P配置_绑定环境变量(v *viper.Viper, 键名 string, 环境变量名 string)` | 绑定键到环境变量 |
| `P配置_自动环境变量` | `func P配置_自动环境变量(v *viper.Viper, 前缀 string)` | 开启自动环境变量绑定 |

**示例**：

```go
v, _ := P配置_从文件读取("config.yaml")
P配置_取值(v, "database.host")           // "localhost"
P配置_取整数值(v, "database.port")        // 3306
P配置_监听变更(v)
P配置_变更回调(v, func() {
    fmt.Println("配置已变更")
})
```

---

## 2.26 E邮件（邮件发送）

> 源文件：`utils/E邮件.go` | 依赖：`github.com/jordan-wright/email` | 用途：SMTP 邮件发送，支持 HTML 正文、附件和 TLS

| 函数 | 签名 | 说明 |
|------|------|------|
| `E邮件_发送` | `func E邮件_发送(服务器地址, 发件人邮箱, 密码, 收件人, 主题, 正文, HTML正文 string, 附件路径 []string) error` | SMTP 发送邮件 |
| `E邮件_发送TLS` | `func E邮件_发送TLS(服务器地址, 发件人邮箱, 密码, 收件人, 主题, 正文, HTML正文 string, 附件路径 []string) error` | TLS 加密发送邮件（465 端口） |
| `E邮件_发送简单邮件` | `func E邮件_发送简单邮件(服务器地址, 发件人邮箱, 密码, 收件人, 主题, 正文 string) error` | 发送纯文本邮件（无附件） |

**示例**：

```go
E邮件_发送简单邮件(
    "smtp.qq.com:587",
    "sender@qq.com", "授权码",
    "receiver@example.com",
    "测试邮件", "这是一封测试邮件",
)

E邮件_发送TLS(
    "smtp.qq.com:465",
    "sender@qq.com", "授权码",
    "a@test.com,b@test.com",
    "HTML邮件", "",
    "<h1>Hello</h1>", []string{"./report.pdf"},
)
```

---

## 2.27 X系统信息（系统指标获取）

> 源文件：`utils/X系统信息.go` | 依赖：`github.com/shirou/gopsutil/v3` | 用途：获取 CPU、内存、磁盘、网络等系统指标

### CPU

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_取CPU信息` | `func X系统_取CPU信息() ([]cpu.InfoStat, error)` | 获取 CPU 信息（型号、核心数等） |
| `X系统_取CPU核心数` | `func X系统_取CPU核心数() (int, error)` | 获取逻辑核心数 |
| `X系统_取CPU物理核心数` | `func X系统_取CPU物理核心数() (int, error)` | 获取物理核心数 |
| `X系统_取CPU使用率` | `func X系统_取CPU使用率(间隔 int) ([]float64, error)` | 获取各核心 CPU 使用率（间隔单位：秒） |
| `X系统_取总CPU使用率` | `func X系统_取总CPU使用率(间隔 int) (float64, error)` | 获取总体 CPU 使用率 |

### 内存

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_取内存信息` | `func X系统_取内存信息() (*mem.VirtualMemoryStat, error)` | 获取内存使用情况（总/已用/可用） |
| `X系统_取交换区信息` | `func X系统_取交换区信息() (*mem.SwapMemoryStat, error)` | 获取 Swap 使用情况 |

### 磁盘

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_取磁盘信息` | `func X系统_取磁盘信息() ([]disk.PartitionStat, error)` | 获取磁盘分区列表 |
| `X系统_取磁盘使用量` | `func X系统_取磁盘使用量(路径 string) (*disk.UsageStat, error)` | 获取指定路径磁盘使用量 |
| `X系统_取磁盘IO信息` | `func X系统_取磁盘IO信息() (map[string]disk.IOCountersStat, error)` | 获取磁盘 IO 统计信息 |

### 网络

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_取网络接口信息` | `func X系统_取网络接口信息() ([]net.InterfaceStat, error)` | 获取网络接口列表 |
| `X系统_取网络连接信息` | `func X系统_取网络连接信息() ([]net.ConnectionStat, error)` | 获取活动网络连接 |
| `X系统_取网络IO信息` | `func X系统_取网络IO信息() ([]net.IOCountersStat, error)` | 获取网络 IO 统计信息 |

### 主机

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_取主机信息` | `func X系统_取主机信息() (*host.InfoStat, error)` | 获取主机名、OS、内核版本等 |
| `X系统_取开机时间` | `func X系统_取开机时间() (uint64, error)` | 获取系统开机时长（秒） |
| `X系统_取系统负载` | `func X系统_取系统负载() (*load.AvgStat, error)` | 获取系统负载（Load1/Load5/Load15） |

### 进程

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_取进程列表` | `func X系统_取进程列表() ([]*process.Process, error)` | 获取所有进程列表 |
| `X系统_取进程信息` | `func X系统_取进程信息(pid int32) (*process.Process, error)` | 获取指定 PID 进程信息 |
| `X系统_取当前进程ID` | `func X系统_取当前进程ID() int32` | 获取当前进程 PID |
| `X系统_取当前进程信息` | `func X系统_取当前进程信息() (*process.Process, error)` | 获取当前进程详细信息 |
| `X系统_取进程名` | `func X系统_取进程名(pid int32) string` | 获取指定 PID 的进程名 |
| `X系统_取进程内存占用` | `func X系统_取进程内存占用(pid int32) float64` | 获取进程内存占用（MB） |
| `X系统_取进程CPU占用` | `func X系统_取进程CPU占用(pid int32) float64` | 获取进程 CPU 使用率（%） |

### 系统信息

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_是否64位系统` | `func X系统_是否64位系统() bool` | 判断是否为 64 位系统 |
| `X系统_取系统架构` | `func X系统_取系统架构() string` | 获取系统架构（如 amd64） |
| `X系统_取操作系统类型` | `func X系统_取操作系统类型() string` | 获取操作系统类型（如 windows） |
| `X系统_取逻辑处理器数` | `func X系统_取逻辑处理器数() int` | 获取逻辑处理器数量 |
| `X系统_取Go版本` | `func X系统_取Go版本() string` | 获取 Go 版本号 |

**示例**：

```go
cpuInfo, _ := X系统_取CPU信息()              // CPU 型号和核心数
cores, _ := X系统_取CPU核心数()               // 12
physCores, _ := X系统_取CPU物理核心数()        // 6
memInfo, _ := X系统_取内存信息()               // &{Total:40873MB ...}
usage, _ := X系统_取磁盘使用量("C:\\")         // &{Total:500G Used:200G ...}
hostInfo, _ := X系统_取主机信息()              // &{Hostname:MY-PC OS:windows ...}
uptime, _ := X系统_取开机时间()                // 26817（秒）
X系统_是否64位系统()                           // true
X系统_取系统架构()                              // "amd64"
X系统_取操作系统类型()                          // "windows"
X系统_取Go版本()                               // "go1.25.5"
X系统_取当前进程ID()                           // 7692
```

---

## 2.28 D定时（定时任务管理）

> 源文件：`utils/D定时.go` | 依赖：`github.com/robfig/cron/v3` | 用途：秒级 cron 定时任务调度

| 函数 | 签名 | 说明 |
|------|------|------|
| `D定时_创建` | `func D定时_创建() *cron.Cron` | 创建秒级 cron 调度器 |
| `D定时_添加任务` | `func D定时_添加任务(调度器 *cron.Cron, 表达式 string, 任务 func()) (cron.EntryID, error)` | 添加定时任务 |
| `D定时_移除任务` | `func D定时_移除任务(调度器 *cron.Cron, 任务ID cron.EntryID)` | 移除指定任务 |
| `D定时_启动` | `func D定时_启动(调度器 *cron.Cron)` | 启动调度器 |
| `D定时_停止` | `func D定时_停止(调度器 *cron.Cron)` | 停止调度器 |
| `D定时_取任务列表` | `func D定时_取任务列表(调度器 *cron.Cron) []cron.Entry` | 获取所有已注册任务 |
| `D定时_简单执行` | `func D定时_简单执行(表达式 string, 任务 func()) (*cron.Cron, error)` | 一条龙创建+添加+启动 |

**cron 表达式格式**（秒级）：`秒 分 时 日 月 周`

| 表达式 | 说明 |
|--------|------|
| `*/5 * * * * *` | 每 5 秒执行 |
| `0 30 9 * * *` | 每天 9:30 执行 |
| `0 0 12 * * 1-5` | 周一至周五 12:00 执行 |

**示例**：

```go
调度器, _ := D定时_简单执行("*/10 * * * * *", func() {
    fmt.Println("每10秒执行一次")
})
defer D定时_停止(调度器)
```

---

## 2.29 G协程池（goroutine 池管理）

> 源文件：`utils/G协程池.go` | 依赖：`github.com/panjf2000/ants/v2` | 用途：高性能 goroutine 池，控制并发数

| 函数 | 签名 | 说明 |
|------|------|------|
| `G协程池_创建` | `func G协程池_创建(池大小 int) (*ants.Pool, error)` | 创建指定大小的协程池 |
| `G协程池_创建带选项` | `func G协程池_创建带选项(池大小 int, 选项 ...ants.Option) (*ants.Pool, error)` | 创建带自定义选项的协程池 |
| `G协程池_提交任务` | `func G协程池_提交任务(池 *ants.Pool, 任务 func()) error` | 提交任务到协程池 |
| `G协程池_取运行中数量` | `func G协程池_取运行中数量(池 *ants.Pool) int` | 获取运行中 worker 数量 |
| `G协程池_取空闲数量` | `func G协程池_取空闲数量(池 *ants.Pool) int` | 获取空闲 worker 数量 |
| `G协程池_取容量` | `func G协程池_取容量(池 *ants.Pool) int` | 获取池容量 |
| `G协程池_取等待数量` | `func G协程池_取等待数量(池 *ants.Pool) int` | 获取等待中任务数量 |
| `G协程池_释放` | `func G协程池_释放(池 *ants.Pool)` | 释放协程池资源 |
| `G协程池_调整大小` | `func G协程池_调整大小(池 *ants.Pool, 新大小 int)` | 动态调整池容量 |
| `G协程池_预分配` | `func G协程池_预分配(池大小 int) (*ants.Pool, error)` | 创建预分配 worker 的协程池 |

**示例**：

```go
池, _ := G协程池_创建(100)
defer G协程池_释放(池)

for i := 0; i < 1000; i++ {
    i := i
    G协程池_提交任务(池, func() {
        fmt.Println("任务", i)
    })
}
```

---

## 2.30 L日志（高性能结构化日志）

> 源文件：`utils/L日志.go` | 依赖：`go.uber.org/zap` | 用途：高性能结构化日志，支持 JSON/Console 格式

### 创建类

| 函数 | 签名 | 说明 |
|------|------|------|
| `L日志_创建开发日志` | `func L日志_创建开发日志() (*zap.Logger, error)` | 创建开发环境日志（Console 格式，Debug 级别） |
| `L日志_创建生产日志` | `func L日志_创建生产日志() (*zap.Logger, error)` | 创建生产环境日志（JSON 格式，Info 级别） |
| `L日志_创建自定义日志` | `func L日志_创建自定义日志(日志级别 int, 输出路径 []string, JSON格式 bool) (*zap.Logger, error)` | 自定义配置创建日志实例 |

### 输出类

| 函数 | 签名 | 说明 |
|------|------|------|
| `L日志_调试` | `func L日志_调试(日志 *zap.Logger, 消息 string, 字段 ...zap.Field)` | 输出 Debug 级别日志 |
| `L日志_信息` | `func L日志_信息(日志 *zap.Logger, 消息 string, 字段 ...zap.Field)` | 输出 Info 级别日志 |
| `L日志_警告` | `func L日志_警告(日志 *zap.Logger, 消息 string, 字段 ...zap.Field)` | 输出 Warn 级别日志 |
| `L日志_错误` | `func L日志_错误(日志 *zap.Logger, 消息 string, 字段 ...zap.Field)` | 输出 Error 级别日志 |

### 字段类

| 函数 | 签名 | 说明 |
|------|------|------|
| `L日志_字符串` | `func L日志_字符串(键 string, 值 string) zap.Field` | 创建字符串字段 |
| `L日志_整数` | `func L日志_整数(键 string, 值 int) zap.Field` | 创建整数字段 |
| `L日志_错误类型` | `func L日志_错误类型(键 string, 值 error) zap.Field` | 创建 error 字段 |

### 管理类

| 函数 | 签名 | 说明 |
|------|------|------|
| `L日志_同步` | `func L日志_同步(日志 *zap.Logger) error` | 刷新日志缓冲区（程序退出前调用） |

**示例**：

```go
日志, _ := L日志_创建开发日志()
defer L日志_同步(日志)

L日志_信息(日志, "服务启动",
    L日志_字符串("host", "localhost"),
    L日志_整数("port", 8080),
)
```

---

## 2.31 K环境变量（环境变量管理）

> 源文件：`utils/K环境变量.go` | 依赖：`github.com/joho/godotenv` | 用途：.env 文件加载与环境变量操作

| 函数 | 签名 | 说明 |
|------|------|------|
| `K环境_加载` | `func K环境_加载(文件路径 string) error` | 从 .env 文件加载环境变量（不覆盖已有） |
| `K环境_加载并覆盖` | `func K环境_加载并覆盖(文件路径 string) error` | 从 .env 文件加载环境变量（覆盖已有） |
| `K环境_取值` | `func K环境_取值(名称 string) string` | 获取环境变量值 |
| `K环境_取值带默认值` | `func K环境_取值带默认值(名称 string, 默认值 string) string` | 获取环境变量值，不存在返回默认值 |
| `K环境_设置值` | `func K环境_设置值(名称 string, 值 string) error` | 设置环境变量 |
| `K环境_删除值` | `func K环境_删除值(名称 string) error` | 删除环境变量 |
| `K环境_是否存在` | `func K环境_是否存在(名称 string) bool` | 判断环境变量是否存在 |
| `K环境_取所有` | `func K环境_取所有() []string` | 获取所有环境变量 |
| `K环境_从文本读取` | `func K环境_从文本读取(文本 string) (map[string]string, error)` | 从文本解析环境变量（无需文件） |

**示例**：

```go
K环境_加载(".env")
dbHost := K环境_取值带默认值("DB_HOST", "localhost")
dbPort := K环境_安全到整数(K环境_取值("DB_PORT"), 3306)
K环境_设置值("APP_MODE", "production")
```

---

## 2.32 M命令行（命令行参数解析）

> 基于 Go 标准库 `flag` | 文件：`utils/M命令行.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `M命令行_取字符串参数` | `func M命令行_取字符串参数(短名称 string, 长名称 string, 默认值 string, 说明 string) *string` | 定义并获取字符串类型命令行参数 |
| `M命令行_取整数参数` | `func M命令行_取整数参数(短名称 string, 长名称 string, 默认值 int, 说明 string) *int` | 定义并获取整数类型命令行参数 |
| `M命令行_取布尔参数` | `func M命令行_取布尔参数(短名称 string, 长名称 string, 默认值 bool, 说明 string) *bool` | 定义并获取布尔类型命令行参数 |
| `M命令行_取小数参数` | `func M命令行_取小数参数(短名称 string, 长名称 string, 默认值 float64, 说明 string) *float64` | 定义并获取小数类型命令行参数 |
| `M命令行_解析` | `func M命令行_解析()` | 解析命令行参数（定义参数后必须调用） |
| `M命令行_取所有参数` | `func M命令行_取所有参数() []string` | 获取所有非标志参数 |
| `M命令行_取参数数量` | `func M命令行_取参数数量() int` | 获取非标志参数数量 |
| `M命令行_取用法` | `func M命令行_取用法() string` | 获取参数使用说明 |
| `M命令行_设置用法` | `func M命令行_设置用法(用法函数 func())` | 自定义参数使用说明输出 |

**示例**：

```go
name := M命令行_取字符串参数("n", "name", "world", "名称")
port := M命令行_取整数参数("p", "port", 8080, "端口")
M命令行_解析()
fmt.Println(*name, *port)
```

---

## 2.33 R日期解析（智能日期解析）

> 基于 `github.com/araddon/dateparse` | 文件：`utils/R日期解析.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `R日期_智能解析` | `func R日期_智能解析(日期文本 string) (time.Time, error)` | 自动识别日期格式并解析 |
| `R日期_解析本地` | `func R日期_解析本地(日期文本 string) (time.Time, error)` | 以本地时区解析日期 |
| `R日期_解析带格式` | `func R日期_解析带格式(日期文本 string, 格式 string) (time.Time, error)` | 按指定格式解析日期 |
| `R日期_取可能格式` | `func R日期_取可能格式(日期文本 string) ([]string, error)` | 获取日期文本可能的格式列表 |
| `R日期_解析任意` | `func R日期_解析任意(日期文本 string, 优先本地时区 ...bool) (time.Time, error)` | 智能解析日期（可指定时区优先） |
| `R日期_取时间戳` | `func R日期_取时间戳(日期文本 string) (int64, error)` | 解析日期并返回 Unix 时间戳 |
| `R日期_取日期部分` | `func R日期_取日期部分(日期文本 string) (string, error)` | 解析日期并返回日期部分（YYYY-MM-DD） |
| `R日期_取时间部分` | `func R日期_取时间部分(日期文本 string) (string, error)` | 解析日期并返回时间部分（HH:MM:SS） |
| `R日期_是否合法` | `func R日期_是否合法(日期文本 string) bool` | 判断日期字符串是否可被解析 |

---

## 2.34 P对象池（字节缓冲区对象池）

> 基于 `github.com/valyala/bytebufferpool` | 文件：`utils/P对象池.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `P对象池_获取` | `func P对象池_获取() *bytebufferpool.ByteBuffer` | 从池中获取字节缓冲区 |
| `P对象池_放回` | `func P对象池_放回(buf *bytebufferpool.ByteBuffer)` | 归还字节缓冲区到池 |
| `P对象池_写入` | `func P对象池_写入(数据 []byte) *bytebufferpool.ByteBuffer` | 获取缓冲区并写入数据 |
| `P对象池_写字符串` | `func P对象池_写字符串(文本 string) *bytebufferpool.ByteBuffer` | 获取缓冲区并写入字符串 |

---

## 2.35 B表达式计算（数学/逻辑表达式）

> 基于 `github.com/Knetic/govaluate` | 文件：`utils/B表达式计算.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `B表达式_计算` | `func B表达式_计算(表达式 string) (interface{}, error)` | 计算简单表达式 |
| `B表达式_计算带参数` | `func B表达式_计算带参数(表达式 string, 参数 map[string]interface{}) (interface{}, error)` | 计算带参数的表达式 |
| `B表达式_新建` | `func B表达式_新建(表达式 string) (*govaluate.EvaluableExpression, error)` | 创建可复用的表达式对象 |
| `B表达式_取变量` | `func B表达式_取变量(表达式 string) ([]string, error)` | 获取表达式中的变量名列表 |
| `B表达式_是否合法` | `func B表达式_是否合法(表达式 string) bool` | 检查表达式语法是否正确 |
| `B表达式_取运算符` | `func B表达式_取运算符(表达式 string) ([]string, error)` | 获取表达式中的运算符列表 |

---

## 2.36 T模板（高性能模板引擎）

> 基于 `github.com/valyala/fasttemplate` | 文件：`utils/T模板.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `T模板_执行` | `func T模板_执行(模板文本 string, 标签 map[string]interface{}) (string, error)` | 执行模板替换（`{key}` 占位符） |
| `T模板_执行自定义` | `func T模板_执行自定义(模板文本 string, 左标签 string, 右标签 string, 标签 map[string]interface{}) (string, error)` | 自定义分隔符的模板替换 |
| `T模板_新建` | `func T模板_新建(模板文本 string, 左标签 string, 右标签 string) (*fasttemplate.Template, error)` | 创建可复用的模板对象 |
| `T模板_执行到写入器` | `func T模板_执行到写入器(模板 *fasttemplate.Template, 标签 map[string]interface{}, 输出 io.Writer) (int, error)` | 将模板结果写入 io.Writer |
| `T模板_执行到字节` | `func T模板_执行到字节(模板 *fasttemplate.Template, 标签 map[string]interface{}) ([]byte, error)` | 执行模板并返回字节切片 |

---

## 2.37 V数据校验（结构体校验）

> 基于 `github.com/go-playground/validator/v10` | 文件：`utils/V数据校验.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `V校验_验证结构体` | `func V校验_验证结构体(结构体 interface{}) error` | 验证结构体标签约束 |
| `V校验_验证变量` | `func V校验_验证变量(值 interface{}, 标签 string) error` | 验证单个变量 |
| `V校验_验证字段` | `func V校验_验证字段(值 interface{}, 标签 string) error` | 验证字段值 |
| `V校验_添加校验规则` | `func V校验_添加校验规则(标签名 string, 规则函数 validator.Func) error` | 添加自定义校验规则 |
| `V校验_取错误详情` | `func V校验_取错误详情(err error) string` | 获取校验错误的详细信息 |
| `V校验_取错误字段` | `func V校验_取错误字段(err error) string` | 获取校验失败的字段名 |
| `V校验_是否校验错误` | `func V校验_是否校验错误(err error) bool` | 判断是否为校验错误 |
| `V校验_注册别名` | `func V校验_注册别名(别名 string, 标签 string)` | 为校验标签注册别名 |
| `V校验_排除零值` | `func V校验_排除零值(值 interface{}) bool` | 判断值是否为零值 |
| `V校验_重置校验器` | `func V校验_重置校验器()` | 重置全局校验器实例 |

---

## 2.38 J结构体合并（结构体合并与拷贝）

> 基于 `dario.cat/mergo` 和 `github.com/jinzhu/copier` | 文件：`utils/J结构体合并.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `J结构体_合并` | `func J结构体_合并(目标 interface{}, 源 interface{}) error` | 合并源到目标（仅覆盖零值） |
| `J结构体_合并覆盖` | `func J结构体_合并覆盖(目标 interface{}, 源 interface{}) error` | 合并源到目标（覆盖所有字段） |
| `J结构体_合并带切片` | `func J结构体_合并带切片(目标 interface{}, 源 interface{}) error` | 合并时保留切片类型 |
| `J结构体_深拷贝` | `func J结构体_深拷贝(目标 interface{}, 源 interface{}) error` | 深拷贝源到目标 |
| `J结构体_拷贝字段` | `func J结构体_拷贝字段(目标 interface{}, 源 interface{}) error` | 按字段名拷贝（不同结构体间） |
| `J结构体_拷贝带选项` | `func J结构体_拷贝带选项(目标 interface{}, 源 interface{}, 忽略空值 bool) error` | 带选项拷贝（可忽略空值） |
| `J结构体_合并Map` | `func J结构体_合并Map(目标 interface{}, 源 map[string]interface{}) error` | 将 Map 合并到结构体 |

---

## 2.39 K表格（控制台表格渲染与数据处理）

> 基于 `github.com/scylladb/termtables` | 文件：`utils/K表格.go`

### 表格创建与渲染

| 函数 | 签名 | 说明 |
|------|------|------|
| `K表格_创建` | `func K表格_创建() *termtables.Table` | 创建空表格对象 |
| `K表格_添加表头` | `func K表格_添加表头(表格 *termtables.Table, 表头 ...string)` | 添加表头 |
| `K表格_添加行` | `func K表格_添加行(表格 *termtables.Table, 行数据 ...interface{})` | 添加一行数据 |
| `K表格_添加分隔线` | `func K表格_添加分隔线(表格 *termtables.Table)` | 添加水平分隔线 |
| `K表格_输出` | `func K表格_输出(表格 *termtables.Table) string` | 渲染表格为字符串 |
| `K表格_快速创建` | `func K表格_快速创建(表头 []string, 行数据 [][]string) string` | 快速创建并渲染表格 |

### 多格式输出

| 函数 | 签名 | 说明 |
|------|------|------|
| `K表格_输出Markdown` | `func K表格_输出Markdown(表头 []string, 行数据 [][]string) string` | 输出 Markdown 格式表格 |
| `K表格_输出CSV` | `func K表格_输出CSV(表头 []string, 行数据 [][]string) string` | 输出 CSV 格式 |
| `K表格_输出TSV` | `func K表格_输出TSV(表头 []string, 行数据 [][]string) string` | 输出 TSV 格式（制表符分隔） |
| `K表格_输出JSON` | `func K表格_输出JSON(表头 []string, 行数据 [][]string) string` | 输出 JSON 数组格式 |
| `K表格_输出HTML` | `func K表格_输出HTML(表头 []string, 行数据 [][]string) string` | 输出 HTML 表格格式 |

### 数据导入

| 函数 | 签名 | 说明 |
|------|------|------|
| `K表格_从CSV读取` | `func K表格_从CSV读取(csv文本 string) ([]string, [][]string, error)` | 从 CSV 文本解析表格 |
| `K表格_从TSV读取` | `func K表格_从TSV读取(tsv文本 string) ([]string, [][]string)` | 从 TSV 文本解析表格 |

### 数据操作

| 函数 | 签名 | 说明 |
|------|------|------|
| `K表格_转置` | `func K表格_转置(表头 []string, 行数据 [][]string) ([]string, [][]string)` | 行列互换 |
| `K表格_过滤行` | `func K表格_过滤行(行数据 [][]string, 条件 func(int, []string) bool) [][]string` | 按条件过滤行 |
| `K表格_排序列` | `func K表格_排序列(行数据 [][]string, 列索引 int, 升序 bool) [][]string` | 按列排序 |
| `K表格_取列` | `func K表格_取列(行数据 [][]string, 列索引 int) []string` | 提取指定列 |
| `K表格_取行` | `func K表格_取行(行数据 [][]string, 行索引 int) []string` | 提取指定行 |
| `K表格_行数` | `func K表格_行数(行数据 [][]string) int` | 获取行数 |
| `K表格_列数` | `func K表格_列数(表头 []string) int` | 获取列数 |
| `K表格_合并` | `func K表格_合并(行数据1 [][]string, 行数据2 [][]string) [][]string` | 合并两个表格 |
| `K表格_去重` | `func K表格_去重(行数据 [][]string) [][]string` | 行数据去重 |

---

## 2.40 F文件监控（文件系统监控）

> 基于 `github.com/fsnotify/fsnotify` | 文件：`utils/F文件监控.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `F文件监控_新建` | `func F文件监控_新建() (*fsnotify.Watcher, error)` | 创建文件监控器 |
| `F文件监控_添加目录` | `func F文件监控_添加目录(监控器 *fsnotify.Watcher, 目录路径 string) error` | 添加监控目录 |
| `F文件监控_移除目录` | `func F文件监控_移除目录(监控器 *fsnotify.Watcher, 目录路径 string) error` | 移除监控目录 |
| `F文件监控_关闭` | `func F文件监控_关闭(监控器 *fsnotify.Watcher) error` | 关闭监控器 |
| `F文件监控_监控目录变化` | `func F文件监控_监控目录变化(目录路径 string, 回调函数 func(event fsnotify.Event)) (func(), error)` | 便捷监控目录变化 |
| `F文件监控_监控多个目录` | `func F文件监控_监控多个目录(目录列表 []string, 回调函数 func(event fsnotify.Event)) (func(), error)` | 监控多个目录 |
| `F文件监控_取事件类型` | `func F文件监控_取事件类型(事件 fsnotify.Event) string` | 获取事件类型描述 |
| `F文件监控_是否创建` | `func F文件监控_是否创建(事件 fsnotify.Event) bool` | 判断是否为创建事件 |
| `F文件监控_是否修改` | `func F文件监控_是否修改(事件 fsnotify.Event) bool` | 判断是否为修改事件 |

---

## 2.41 X消息总线（发布-订阅消息）

> 基于 `github.com/vardius/message-bus` | 文件：`utils/X消息总线.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `X消息_新建总线` | `func X消息_新建总线(缓冲大小 int) messagebus.MessageBus` | 创建消息总线 |
| `X消息_发布` | `func X消息_发布(总线 messagebus.MessageBus, 主题 string, 参数 ...interface{})` | 发布消息到主题 |
| `X消息_订阅` | `func X消息_订阅(总线 messagebus.MessageBus, 主题 string, 回调函数 interface{}) error` | 订阅主题 |
| `X消息_取消订阅` | `func X消息_取消订阅(总线 messagebus.MessageBus, 主题 string, 回调函数 interface{}) error` | 取消订阅 |
| `X消息_关闭总线` | `func X消息_关闭总线(总线 messagebus.MessageBus)` | 关闭消息总线 |

---

## 2.42 H客户端（HTTP 客户端）

> 基于 `github.com/go-resty/resty/v2` | 文件：`utils/H客户端.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `H客户端_Get` | `func H客户端_Get(网址 string) (*resty.Response, error)` | 发送 GET 请求 |
| `H客户端_Post` | `func H客户端_Post(网址 string, 数据 interface{}) (*resty.Response, error)` | 发送 POST 请求（JSON） |
| `H客户端_Put` | `func H客户端_Put(网址 string, 数据 interface{}) (*resty.Response, error)` | 发送 PUT 请求（JSON） |
| `H客户端_Delete` | `func H客户端_Delete(网址 string) (*resty.Response, error)` | 发送 DELETE 请求 |
| `H客户端_带头Get` | `func H客户端_带头Get(网址 string, 请求头 map[string]string) (*resty.Response, error)` | 带请求头的 GET |
| `H客户端_带头Post` | `func H客户端_带头Post(网址 string, 数据 interface{}, 请求头 map[string]string) (*resty.Response, error)` | 带请求头的 POST |
| `H客户端_设置超时` | `func H客户端_设置超时(秒数 int) *resty.Client` | 设置全局超时时间 |
| `H客户端_设置代理` | `func H客户端_设置代理(代理地址 string) *resty.Client` | 设置 HTTP 代理 |
| `H客户端_设置Cookie` | `func H客户端_设置Cookie(cookies []*http.Cookie) *resty.Client` | 设置全局 Cookie |
| `H客户端_下载文件` | `func H客户端_下载文件(网址 string, 保存路径 string) error` | 下载文件到本地 |
| `H客户端_提交表单` | `func H客户端_提交表单(网址 string, 表单数据 map[string]string) (*resty.Response, error)` | 提交表单数据 |
| `H客户端_新建客户端` | `func H客户端_新建客户端() *resty.Client` | 创建自定义客户端实例 |

---

## 2.43 N键值库（嵌入式键值数据库）

> 基于 `github.com/tidwall/buntdb` | 文件：`utils/N键值库.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `N键值_打开` | `func N键值_打开(文件路径 string) (*buntdb.DB, error)` | 打开/创建键值数据库 |
| `N键值_关闭` | `func N键值_关闭(数据库 *buntdb.DB) error` | 关闭数据库 |
| `N键值_置值` | `func N键值_置值(数据库 *buntdb.DB, 键 string, 值 string) error` | 设置键值 |
| `N键值_取值` | `func N键值_取值(数据库 *buntdb.DB, 键 string) (string, error)` | 获取键值 |
| `N键值_置值带过期` | `func N键值_置值带过期(数据库 *buntdb.DB, 键 string, 值 string, 过期秒数 float64) error` | 设置键值（带过期时间） |
| `N键值_删除` | `func N键值_删除(数据库 *buntdb.DB, 键 string) error` | 删除键 |
| `N键值_遍历` | `func N键值_遍历(数据库 *buntdb.DB, 回调函数 func(键, 值 string) bool) error` | 遍历所有键值 |
| `N键值_创建索引` | `func N键值_创建索引(数据库 *buntdb.DB, 索引名 string, 模式 string) error` | 创建索引 |

---

## 2.44 C爬虫（网页爬虫框架）

> 基于 `github.com/gocolly/colly/v2` | 文件：`utils/C爬虫.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `C爬虫_新建` | `func C爬虫_新建(选项 ...colly.CollectorOption) *colly.Collector` | 创建采集器 |
| `C爬虫_访问` | `func C爬虫_访问(采集器 *colly.Collector, 网址 string) error` | 访问指定 URL |
| `C爬虫_注册HTML回调` | `func C爬虫_注册HTML回调(采集器 *colly.Collector, 选择器 string, 回调函数 func(e *colly.HTMLElement))` | 注册 HTML 元素回调 |
| `C爬虫_注册请求回调` | `func C爬虫_注册请求回调(采集器 *colly.Collector, 回调函数 func(r *colly.Request))` | 注册请求回调 |
| `C爬虫_注册响应回调` | `func C爬虫_注册响应回调(采集器 *colly.Collector, 回调函数 func(r *colly.Response))` | 注册响应回调 |
| `C爬虫_注册错误回调` | `func C爬虫_注册错误回调(采集器 *colly.Collector, 回调函数 func(r *colly.Response, err error))` | 注册错误回调 |
| `C爬虫_限制并发` | `func C爬虫_限制并发(采集器 *colly.Collector, 并发数 int, 延时毫秒 int64, 域名模式 ...string)` | 限制并发和延时 |
| `C爬虫_设置代理` | `func C爬虫_设置代理(采集器 *colly.Collector, 代理地址 string)` | 设置代理 |
| `C爬虫_设置UserAgent` | `func C爬虫_设置UserAgent(采集器 *colly.Collector, ua string)` | 设置 User-Agent |
| `C爬虫_设置Cookie` | `func C爬虫_设置Cookie(采集器 *colly.Collector, 网址 string, cookies []*http.Cookie)` | 设置 Cookie |

---

## 2.45 Q权限管理（RBAC/ABAC 权限）

> 基于 `github.com/casbin/casbin/v2` | 文件：`utils/Q权限管理.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `Q权限_新建管理器` | `func Q权限_新建管理器(模型路径 string, 策略路径 string) (*casbin.Enforcer, error)` | 从配置文件创建权限管理器 |
| `Q权限_新建管理器从文本` | `func Q权限_新建管理器从文本(模型文本 string, 策略文本 string) (*casbin.Enforcer, error)` | 从文本创建权限管理器 |
| `Q权限_检查权限` | `func Q权限_检查权限(管理器 *casbin.Enforcer, 主体 string, 资源 string, 操作 string) (bool, error)` | 检查权限 |
| `Q权限_添加策略` | `func Q权限_添加策略(管理器 *casbin.Enforcer, 主体 string, 资源 string, 操作 string) (bool, error)` | 添加策略 |
| `Q权限_删除策略` | `func Q权限_删除策略(管理器 *casbin.Enforcer, 主体 string, 资源 string, 操作 string) (bool, error)` | 删除策略 |
| `Q权限_添加角色` | `func Q权限_添加角色(管理器 *casbin.Enforcer, 用户 string, 角色 string) (bool, error)` | 为用户添加角色 |
| `Q权限_删除角色` | `func Q权限_删除角色(管理器 *casbin.Enforcer, 用户 string, 角色 string) (bool, error)` | 删除用户角色 |
| `Q权限_获取角色` | `func Q权限_获取角色(管理器 *casbin.Enforcer, 用户 string) ([]string, error)` | 获取用户角色列表 |
| `Q权限_获取权限` | `func Q权限_获取权限(管理器 *casbin.Enforcer, 主体 string) ([][]string, error)` | 获取主体权限列表 |
| `Q权限_保存策略` | `func Q权限_保存策略(管理器 *casbin.Enforcer) error` | 保存策略到文件 |
| `Q权限_加载策略` | `func Q权限_加载策略(管理器 *casbin.Enforcer) error` | 从文件加载策略 |
| `Q权限_获取所有策略` | `func Q权限_获取所有策略(管理器 *casbin.Enforcer) [][]string` | 获取所有策略 |

---

## 2.46 D数据库（ORM 数据库操作）

> 基于 `xorm.io/xorm` | 文件：`utils/D数据库.go`

| 函数 | 签名 | 说明 |
|------|------|------|
| `D数据库_连接MySQL` | `func D数据库_连接MySQL(用户名 string, 密码 string, 主机地址 string, 数据库名 string) (*xorm.Engine, error)` | 连接 MySQL |
| `D数据库_连接SQLite` | `func D数据库_连接SQLite(文件路径 string) (*xorm.Engine, error)` | 连接 SQLite |
| `D数据库_测试连接` | `func D数据库_测试连接(引擎 *xorm.Engine) error` | 测试数据库连接 |
| `D数据库_同步表` | `func D数据库_同步表(引擎 *xorm.Engine, 结构体 ...interface{}) error` | 同步结构体到表 |
| `D数据库_插入` | `func D数据库_插入(引擎 *xorm.Engine, 记录 interface{}) (int64, error)` | 插入记录 |
| `D数据库_查询` | `func D数据库_查询(引擎 *xorm.Engine, 结果切片 interface{}) error` | 查询记录 |
| `D数据库_查询单条` | `func D数据库_查询单条(引擎 *xorm.Engine, 记录 interface{}) (bool, error)` | 查询单条记录 |
| `D数据库_更新` | `func D数据库_更新(引擎 *xorm.Engine, 记录 interface{}) (int64, error)` | 更新记录 |
| `D数据库_条件更新` | `func D数据库_条件更新(引擎 *xorm.Engine, 记录 interface{}, 条件 string, 参数 ...interface{}) (int64, error)` | 带条件更新 |
| `D数据库_删除` | `func D数据库_删除(引擎 *xorm.Engine, 条件结构体 interface{}) (int64, error)` | 删除记录 |
| `D数据库_条件查询` | `func D数据库_条件查询(引擎 *xorm.Engine, 结果切片 interface{}, 条件 string, 参数 ...interface{}) error` | 带条件查询 |
| `D数据库_统计` | `func D数据库_统计(引擎 *xorm.Engine, 条件结构体 interface{}) (int64, error)` | 统计记录数 |
| `D数据库_执行SQL` | `func D数据库_执行SQL(引擎 *xorm.Engine, sql语句 string, 参数 ...interface{}) (sql.Result, error)` | 执行原生 SQL |
| `D数据库_事务` | `func D数据库_事务(引擎 *xorm.Engine, 事务函数 func(session *xorm.Session) (interface{}, error)) (interface{}, error)` | 执行事务 |
| `D数据库_设置连接池` | `func D数据库_设置连接池(引擎 *xorm.Engine, 最大空闲数 int, 最大连接数 int, 最大生存时间 int)` | 设置连接池参数 |
| `D数据库_关闭` | `func D数据库_关闭(引擎 *xorm.Engine) error` | 关闭数据库连接 |
| `D数据库_取表名` | `func D数据库_取表名(引擎 *xorm.Engine, 结构体 interface{}) string` | 获取表名 |
| `D数据库_表是否存在` | `func D数据库_表是否存在(引擎 *xorm.Engine, 表名 string) (bool, error)` | 检查表是否存在 |
| `D数据库_取反射类型` | `func D数据库_取反射类型(值 interface{}) reflect.Type` | 获取反射类型 |

---

## 2.47 C窗口（Windows 窗口操作）

> 基于 Win32 API `user32.dll` | 文件：`utils/C窗口.go` | 仅 Windows 平台

| 函数 | 签名 | 说明 |
|------|------|------|
| `C窗口_查找` | `func C窗口_查找(类名 string, 标题 string) syscall.Handle` | 按类名/标题查找顶层窗口 |
| `C窗口_查找子窗口` | `func C窗口_查找子窗口(父窗口 syscall.Handle, 子窗口后 syscall.Handle, 类名 string, 标题 string) syscall.Handle` | 查找子窗口 |
| `C窗口_取标题` | `func C窗口_取标题(窗口句柄 syscall.Handle) string` | 获取窗口标题 |
| `C窗口_置标题` | `func C窗口_置标题(窗口句柄 syscall.Handle, 标题 string) bool` | 设置窗口标题 |
| `C窗口_取类名` | `func C窗口_取类名(窗口句柄 syscall.Handle) string` | 获取窗口类名 |
| `C窗口_取矩形` | `func C窗口_取矩形(窗口句柄 syscall.Handle) (RECT, bool)` | 获取窗口位置和大小 |
| `C窗口_移动` | `func C窗口_移动(窗口句柄 syscall.Handle, 左边 int32, 顶边 int32, 宽度 int32, 高度 int32, 重绘 bool) bool` | 移动/调整窗口 |
| `C窗口_显示` | `func C窗口_显示(窗口句柄 syscall.Handle, 命令 int) bool` | 控制窗口显示状态 |
| `C窗口_发送消息` | `func C窗口_发送消息(窗口句柄 syscall.Handle, 消息 uint32, 参数1 uintptr, 参数2 uintptr) uintptr` | 发送同步消息 |
| `C窗口_投递消息` | `func C窗口_投递消息(窗口句柄 syscall.Handle, 消息 uint32, 参数1 uintptr, 参数2 uintptr) bool` | 投递异步消息 |
| `C窗口_关闭` | `func C窗口_关闭(窗口句柄 syscall.Handle) bool` | 关闭窗口 |
| `C窗口_点击按钮` | `func C窗口_点击按钮(按钮句柄 syscall.Handle) uintptr` | 点击按钮控件 |
| `C窗口_取前台窗口` | `func C窗口_取前台窗口() syscall.Handle` | 获取前台窗口 |
| `C窗口_置前台窗口` | `func C窗口_置前台窗口(窗口句柄 syscall.Handle) bool` | 设置前台窗口 |
| `C窗口_是否可见` | `func C窗口_是否可见(窗口句柄 syscall.Handle) bool` | 检查窗口是否可见 |
| `C窗口_是否有效` | `func C窗口_是否有效(窗口句柄 syscall.Handle) bool` | 检查窗口句柄是否有效 |
| `C窗口_取进程ID` | `func C窗口_取进程ID(窗口句柄 syscall.Handle) (uint32, uint32)` | 获取窗口所属进程/线程 ID |
| `C窗口_取父窗口` | `func C窗口_取父窗口(窗口句柄 syscall.Handle) syscall.Handle` | 获取父窗口 |
| `C窗口_取桌面窗口` | `func C窗口_取桌面窗口() syscall.Handle` | 获取桌面窗口 |
| `C窗口_启用` | `func C窗口_启用(窗口句柄 syscall.Handle, 启用 bool) bool` | 启用/禁用窗口 |
| `C窗口_取下一个` | `func C窗口_取下一个(窗口句柄 syscall.Handle) syscall.Handle` | 获取 Z 顺序下一个窗口 |
| `C窗口_取所有子窗口` | `func C窗口_取所有子窗口(父窗口 syscall.Handle) []syscall.Handle` | 获取所有子窗口 |
| `C窗口_置文本` | `func C窗口_置文本(窗口句柄 syscall.Handle, 文本 string) uintptr` | 设置编辑框文本 |
| `C窗口_取文本` | `func C窗口_取文本(窗口句柄 syscall.Handle) string` | 获取编辑框文本 |

**常量**：

| 常量 | 值 | 说明 |
|------|-----|------|
| `SW_HIDE` | 0 | 隐藏窗口 |
| `SW_SHOW` | 5 | 显示窗口 |
| `SW_MINIMIZE` | 6 | 最小化窗口 |
| `SW_MAXIMIZE` | 3 | 最大化窗口 |
| `SW_RESTORE` | 9 | 还原窗口 |
| `WM_CLOSE` | 0x0010 | 关闭消息 |
| `WM_SETTEXT` | 0x000C | 设置文本消息 |
| `WM_GETTEXT` | 0x000D | 获取文本消息 |
| `BM_CLICK` | 0x00F5 | 按钮点击消息 |

---

## 2.48 C进程（Windows 进程管理）

> 基于 Win32 API `kernel32.dll` | 文件：`utils/C进程.go` | 仅 Windows 平台

| 函数 | 签名 | 说明 |
|------|------|------|
| `C进程_创建` | `func C进程_创建(程序路径 string, 命令行 string, 工作目录 string) (*PROCESS_INFORMATION, error)` | 创建新进程 |
| `C进程_打开` | `func C进程_打开(进程ID uint32, 访问权限 uint32) (syscall.Handle, error)` | 打开已存在进程 |
| `C进程_终止` | `func C进程_终止(进程ID uint32, 退出码 uint32) error` | 终止进程 |
| `C进程_是否存活` | `func C进程_是否存活(进程ID uint32) bool` | 检查进程是否运行 |
| `C进程_等待` | `func C进程_等待(进程句柄 syscall.Handle, 超时毫秒 uint32) uint32` | 等待进程退出 |
| `C进程_关闭句柄` | `func C进程_关闭句柄(句柄 syscall.Handle) bool` | 关闭句柄 |
| `C进程_取当前ID` | `func C进程_取当前ID() uint32` | 获取当前进程 ID |
| `C进程_取ID` | `func C进程_取ID(进程句柄 syscall.Handle) uint32` | 通过句柄获取进程 ID |
| `C进程_取退出码` | `func C进程_取退出码(进程句柄 syscall.Handle) (uint32, bool)` | 获取进程退出码 |
| `C进程_设置优先级` | `func C进程_设置优先级(进程句柄 syscall.Handle, 优先级 uint32) bool` | 设置进程优先级 |
| `C进程_取优先级` | `func C进程_取优先级(进程句柄 syscall.Handle) (uint32, bool)` | 获取进程优先级 |
| `C进程_枚举` | `func C进程_枚举() ([]PROCESSENTRY32W, error)` | 枚举所有进程 |
| `C进程_按名查找` | `func C进程_按名查找(进程名 string) ([]uint32, error)` | 按进程名查找 |
| `C进程_取模块路径` | `func C进程_取模块路径(进程ID uint32) (string, error)` | 获取进程可执行文件路径 |
| `C进程_取父进程ID` | `func C进程_取父进程ID(进程ID uint32) (uint32, error)` | 获取父进程 ID |

**常量**：

| 常量 | 值 | 说明 |
|------|-----|------|
| `PROCESS_TERMINATE` | 0x0001 | 终止进程权限 |
| `PROCESS_QUERY_INFORMATION` | 0x0400 | 查询进程信息权限 |
| `STILL_ACTIVE` | 259 | 进程仍在运行 |
| `INFINITE` | 0xFFFFFFFF | 无限等待 |
| `IDLE_PRIORITY_CLASS` | 0x0040 | 空闲优先级 |
| `NORMAL_PRIORITY_CLASS` | 0x0020 | 正常优先级 |
| `HIGH_PRIORITY_CLASS` | 0x0080 | 高优先级 |
| `REALTIME_PRIORITY_CLASS` | 0x0100 | 实时优先级 |

**结构体**：

| 结构体 | 说明 |
|--------|------|
| `RECT` | 窗口矩形区域（Left, Top, Right, Bottom） |
| `PROCESSENTRY32W` | 进程信息（ProcessID, ParentProcessID, ExeFile 等） |
| `STARTUPINFOW` | 进程启动信息 |
| `PROCESS_INFORMATION` | 进程信息（Process, Thread, ProcessID, ThreadID） |

---

## 2.49 OCV视觉（OpenCV 计算机视觉）

> 源文件：`utils/OCV视觉.go` | 依赖：`gocv.io/x/gocv`（需要安装 OpenCV 4.x C++ 库）

### 核心操作

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_取版本` | `func OCV_取版本() string` | 获取 OpenCV 版本号 |
| `OCV_取CUDA设备数` | `func OCV_取CUDA设备数() int` | 获取 CUDA 设备数量 |

### 图像读取与保存

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_读取图片` | `func OCV_读取图片(文件路径 string) (gocv.Mat, error)` | 读取彩色图片 |
| `OCV_读取图片灰度` | `func OCV_读取图片灰度(文件路径 string) (gocv.Mat, error)` | 读取灰度图片 |
| `OCV_读取图片原色` | `func OCV_读取图片原色(文件路径 string) (gocv.Mat, error)` | 读取原始通道图片 |
| `OCV_保存图片` | `func OCV_保存图片(矩阵 gocv.Mat, 文件路径 string) bool` | 保存图片到文件 |
| `OCV_从字节读取` | `func OCV_从字节读取(数据 []byte) (gocv.Mat, error)` | 从字节读取彩色图片 |
| `OCV_从字节读取灰度` | `func OCV_从字节读取灰度(数据 []byte) (gocv.Mat, error)` | 从字节读取灰度图片 |
| `OCV_到字节` | `func OCV_到字节(矩阵 gocv.Mat, 扩展名 string) ([]byte, error)` | 编码为字节切片 |

### Mat 信息与属性

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_取宽度` | `func OCV_取宽度(矩阵 gocv.Mat) int` | 获取宽度 |
| `OCV_取高度` | `func OCV_取高度(矩阵 gocv.Mat) int` | 获取高度 |
| `OCV_取通道数` | `func OCV_取通道数(矩阵 gocv.Mat) int` | 获取通道数 |
| `OCV_取类型` | `func OCV_取类型(矩阵 gocv.Mat) gocv.MatType` | 获取数据类型 |
| `OCV_取像素数` | `func OCV_取像素数(矩阵 gocv.Mat) int` | 获取总像素数 |
| `OCV_是否为空` | `func OCV_是否为空(矩阵 gocv.Mat) bool` | 判断是否为空 |
| `OCV_取尺寸` | `func OCV_取尺寸(矩阵 gocv.Mat) (int, int)` | 获取尺寸（宽, 高） |

### 颜色空间转换

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_BGR转灰度` | `func OCV_BGR转灰度(矩阵 gocv.Mat) gocv.Mat` | BGR → 灰度 |
| `OCV_灰度转BGR` | `func OCV_灰度转BGR(矩阵 gocv.Mat) gocv.Mat` | 灰度 → BGR |
| `OCV_BGR转HSV` | `func OCV_BGR转HSV(矩阵 gocv.Mat) gocv.Mat` | BGR → HSV |
| `OCV_HSV转BGR` | `func OCV_HSV转BGR(矩阵 gocv.Mat) gocv.Mat` | HSV → BGR |
| `OCV_BGR转RGB` | `func OCV_BGR转RGB(矩阵 gocv.Mat) gocv.Mat` | BGR → RGB |
| `OCV_RGB转BGR` | `func OCV_RGB转BGR(矩阵 gocv.Mat) gocv.Mat` | RGB → BGR |
| `OCV_BGR转Lab` | `func OCV_BGR转Lab(矩阵 gocv.Mat) gocv.Mat` | BGR → Lab |
| `OCV_BGR转YUV` | `func OCV_BGR转YUV(矩阵 gocv.Mat) gocv.Mat` | BGR → YUV |

### 图像变换

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_缩放` | `func OCV_缩放(矩阵 gocv.Mat, 宽度, 高度 int) gocv.Mat` | 缩放到指定尺寸 |
| `OCV_缩放按比例` | `func OCV_缩放按比例(矩阵 gocv.Mat, 水平比例, 垂直比例 float64) gocv.Mat` | 按比例缩放 |
| `OCV_裁剪` | `func OCV_裁剪(矩阵 gocv.Mat, 左, 上, 宽度, 高度 int) gocv.Mat` | 裁剪区域 |
| `OCV_旋转` | `func OCV_旋转(矩阵 gocv.Mat, 角度 int) gocv.Mat` | 旋转（0=90°, 1=180°, 2=270°） |
| `OCV_水平翻转` | `func OCV_水平翻转(矩阵 gocv.Mat) gocv.Mat` | 左右镜像 |
| `OCV_垂直翻转` | `func OCV_垂直翻转(矩阵 gocv.Mat) gocv.Mat` | 上下镜像 |
| `OCV_仿射变换` | `func OCV_仿射变换(矩阵 gocv.Mat, 源点, 目标点 []image.Point) gocv.Mat` | 三点仿射变换 |
| `OCV_透视变换` | `func OCV_透视变换(矩阵 gocv.Mat, 源点, 目标点 []image.Point) gocv.Mat` | 四点透视变换 |

### 图像滤波

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_高斯模糊` | `func OCV_高斯模糊(矩阵 gocv.Mat, 核大小 int, 标准差 float64) gocv.Mat` | 高斯模糊 |
| `OCV_中值滤波` | `func OCV_中值滤波(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 中值滤波（去椒盐噪声） |
| `OCV_双边滤波` | `func OCV_双边滤波(矩阵 gocv.Mat, 直径 int, 颜色空间, 坐标空间 float64) gocv.Mat` | 双边滤波（保边去噪） |
| `OCV_方框滤波` | `func OCV_方框滤波(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 方框滤波 |

### 形态学操作

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_腐蚀` | `func OCV_腐蚀(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 腐蚀（膨胀暗区域） |
| `OCV_膨胀` | `func OCV_膨胀(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 膨胀（膨胀亮区域） |
| `OCV_开运算` | `func OCV_开运算(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 开运算（去小噪点） |
| `OCV_闭运算` | `func OCV_闭运算(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 闭运算（填小孔洞） |
| `OCV_形态学梯度` | `func OCV_形态学梯度(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 形态学梯度（提取边缘） |
| `OCV_顶帽变换` | `func OCV_顶帽变换(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 顶帽变换（提取小亮区域） |
| `OCV_黑帽变换` | `func OCV_黑帽变换(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | 黑帽变换（提取小暗区域） |

### 边缘检测

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_Canny边缘检测` | `func OCV_Canny边缘检测(矩阵 gocv.Mat, 低阈值, 高阈值 float64) gocv.Mat` | Canny 边缘检测 |
| `OCV_Sobel边缘检测` | `func OCV_Sobel边缘检测(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | Sobel 边缘检测 |
| `OCV_Laplacian边缘检测` | `func OCV_Laplacian边缘检测(矩阵 gocv.Mat, 核大小 int) gocv.Mat` | Laplacian 边缘检测 |
| `OCV_Scharr边缘检测` | `func OCV_Scharr边缘检测(矩阵 gocv.Mat) gocv.Mat` | Scharr 边缘检测 |

### 阈值处理

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_二值化` | `func OCV_二值化(矩阵 gocv.Mat, 阈值, 最大值 float64) gocv.Mat` | 固定阈值二值化 |
| `OCV_自适应二值化` | `func OCV_自适应二值化(矩阵 gocv.Mat, 最大值 float64, 核大小 int, 常数 float64) gocv.Mat` | 自适应阈值二值化 |
| `OCV_反二值化` | `func OCV_反二值化(矩阵 gocv.Mat, 阈值, 最大值 float64) gocv.Mat` | 反二值化 |
| `OCV_OTSU二值化` | `func OCV_OTSU二值化(矩阵 gocv.Mat, 最大值 float64) gocv.Mat` | OTSU 自动阈值二值化 |

### 轮廓检测

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_查找轮廓` | `func OCV_查找轮廓(矩阵 gocv.Mat) [][]image.Point` | 查找外轮廓 |
| `OCV_查找全部轮廓` | `func OCV_查找全部轮廓(矩阵 gocv.Mat) [][]image.Point` | 查找所有层级轮廓 |
| `OCV_绘制轮廓` | `func OCV_绘制轮廓(矩阵 gocv.Mat, 轮廓 [][]image.Point, 颜色 color.RGBA, 线宽 int) gocv.Mat` | 绘制轮廓 |
| `OCV_轮廓面积` | `func OCV_轮廓面积(轮廓 []image.Point) float64` | 计算轮廓面积 |
| `OCV_轮廓周长` | `func OCV_轮廓周长(轮廓 []image.Point, 闭合 bool) float64` | 计算轮廓周长 |
| `OCV_轮廓外接矩形` | `func OCV_轮廓外接矩形(轮廓 []image.Point) image.Rectangle` | 获取外接矩形 |
| `OCV_轮廓最小外接圆` | `func OCV_轮廓最小外接圆(轮廓 []image.Point) (image.Point, float64)` | 获取最小外接圆 |

### 特征检测

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_Harris角点检测` | `func OCV_Harris角点检测(矩阵 gocv.Mat, 核大小 int, 标准差, K值 float64) gocv.Mat` | Harris 角点检测 |
| `OCV_良好角点` | `func OCV_良好角点(矩阵 gocv.Mat, 最大数量 int, 质量, 最小距离 float64) []image.Point` | Shi-Tomasi 角点 |
| `OCV_FAST角点` | `func OCV_FAST角点(矩阵 gocv.Mat, 阈值 int) []image.Point` | FAST 角点检测 |

### 模板匹配（找图）

> 本节函数在源图中查找模板小图，返回匹配位置和相似度。适用于桌面自动化、图像识别等场景。

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_找图` | `func OCV_找图(源图 gocv.Mat, 模板 gocv.Mat, 相似度 float64) (image.Point, float64, error)` | 归一化相关系数匹配（默认方法，抗光照变化），回左上角坐标+置信度 |
| `OCV_找图带方法` | `func OCV_找图带方法(源图, 模板 gocv.Mat, 相似度 float64, 方法 gocv.TemplateMatchMode) (image.Point, float64, error)` | 指定匹配方法：`TmSqdiffNormed`(平方差)/`TmCcorrNormed`(相关)/`TmCcoeffNormed`(相关系数) |
| `OCV_找图中心` | `func OCV_找图中心(源图, 模板 gocv.Mat, 相似度 float64) (image.Point, float64, error)` | 返回匹配区域中心点（适合直接用于点击操作） |
| `OCV_找图全部` | `func OCV_找图全部(源图, 模板 gocv.Mat, 相似度 float64, 最小间距 int) ([]image.Point, []float64, error)` | 查找所有匹配位置（去重），返回位置列表和置信度列表 |
| `OCV_找图区域` | `func OCV_找图区域(源图, 模板 gocv.Mat, 左, 上, 宽度, 高度 int, 相似度 float64) (image.Point, float64, error)` | 在指定矩形区域内查找，返回源图绝对坐标 |
| `OCV_找图掩码` | `func OCV_找图掩码(源图, 模板, 掩码 gocv.Mat, 相似度 float64) (image.Point, float64, error)` | 带掩码匹配（支持透明图，掩码白色区域参与匹配） |

**示例**：

```go
源图, _ := OCV_读取图片("screenshot.png")
defer 源图.Close()
模板, _ := OCV_读取图片("icon.png")
defer 模板.Close()

位置, 置信度, err := OCV_找图(源图, 模板, 0.8)
if err == nil && 置信度 >= 0.8 {
    fmt.Printf("找到: (%d, %d) 置信度: %.2f\n", 位置.X, 位置.Y, 置信度)
    // 标记匹配位置
    结果 := OCV_画矩形(源图, image.Rect(位置.X, 位置.Y, 位置.X+模板.Cols(), 位置.Y+模板.Rows()), color.RGBA{0, 255, 0, 255}, 2)
    OCV_保存图片(结果, "result.png")
}
```

### 特征匹配（高级找图）

> 本节函数使用 SIFT/ORB/AKAZE 特征匹配算法，**抗缩放、旋转、视角和光照变化**，适合复杂场景下的图像定位。注意：SIFT 有专利限制。

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_找图SIFT` | `func OCV_找图SIFT(源图, 模板 gocv.Mat, 最小匹配点数 int) (image.Rectangle, []gocv.DMatch, error)` | SIFT 特征匹配（抗旋转/缩放/视角，精度最高） |
| `OCV_找图ORB` | `func OCV_找图ORB(源图, 模板 gocv.Mat, 最小匹配点数 int) (image.Rectangle, []gocv.DMatch, error)` | ORB 特征匹配（免费、快速，适合实时场景） |
| `OCV_找图AKAZE` | `func OCV_找图AKAZE(源图, 模板 gocv.Mat, 最小匹配点数 int) (image.Rectangle, []gocv.DMatch, error)` | AKAZE 特征匹配（非线性尺度空间，适合模糊/压缩图片） |
| `OCV_特征匹配` | `func OCV_特征匹配(检测器 gocv.Feature2D, 源图, 模板 gocv.Mat, 最小匹配点数 int, 范数类型 gocv.NormType, 距离比例 float64) (image.Rectangle, []gocv.DMatch, error)` | 通用引擎（支持任意 Feature2D 检测器+Lowe's ratio test） |
| `OCV_找图多尺度` | `func OCV_找图多尺度(源图, 模板 gocv.Mat, 最小比例, 最大比例, 步长, 相似度 float64) (image.Point, float64, float64, error)` | 多尺度模板匹配（缩放模板后逐一匹配，返回最佳比例） |
| `OCV_找图边缘` | `func OCV_找图边缘(源图, 模板 gocv.Mat, 低阈值, 高阈值, 相似度 float64) (image.Point, float64, error)` | 边缘匹配（Canny 提取边缘后模板匹配，对颜色/光照不敏感） |

**示例**：

```go
源图, _ := OCV_读取图片("scene.jpg")
defer 源图.Close()
模板, _ := OCV_读取图片("object.jpg")
defer 模板.Close()

// ORB 特征匹配（推荐，快速 + 抗缩放旋转）
矩形, 匹配点, _ := OCV_找图ORB(源图, 模板, 10)
if len(匹配点) >= 10 {
    fmt.Println("ORB找到:", 矩形)
    结果 := OCV_画矩形(源图, 矩形, color.RGBA{0, 255, 0, 255}, 2)
    OCV_保存图片(结果, "orb_match.jpg")
}

// 多尺度匹配（模板可能被缩放时使用）
位置, 置信度, 比例, _ := OCV_找图多尺度(源图, 模板, 0.3, 3.0, 0.1, 0.8)
fmt.Printf("多尺度: (%d,%d) 比例:%.1fx 置信度:%.2f\n", 位置.X, 位置.Y, 比例, 置信度)

// 边缘匹配（跨主题/颜色不同时使用）
位置, 置信度, _ = OCV_找图边缘(源图, 模板, 50, 150, 0.5)
fmt.Println("边缘匹配:", 位置, 置信度)
```

**方法选择指南**：

| 场景 | 推荐方法 | 原因 |
|------|----------|------|
| 模板与源图尺寸一致、无旋转 | `OCV_找图` | 速度最快，最常用 |
| 大小可能不同、无旋转 | `OCV_找图多尺度` | 多比例逐一匹配 |
| 存在旋转/缩放/视角变化 | `OCV_找图ORB` | ORB 免费、快速 |
| 精度要求极高 | `OCV_找图SIFT` | SIFT 精度最高 |
| 颜色/主题不同（如按钮换肤） | `OCV_找图边缘` | 仅匹配形状轮廓 |
| 模板有透明背景 | `OCV_找图掩码` | 忽略透明区域 |

### 直方图

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_计算直方图` | `func OCV_计算直方图(矩阵 gocv.Mat, 区间数 int) gocv.Mat` | 计算灰度直方图 |
| `OCV_直方图均衡化` | `func OCV_直方图均衡化(矩阵 gocv.Mat) gocv.Mat` | 直方图均衡化 |
| `OCV_CLAHE均衡化` | `func OCV_CLAHE均衡化(矩阵 gocv.Mat, 限幅 float64, 网格大小 int) gocv.Mat` | CLAHE 自适应均衡化 |

### 绘图

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_画线` | `func OCV_画线(矩阵 gocv.Mat, 起点, 终点 image.Point, 颜色 color.RGBA, 线宽 int) gocv.Mat` | 画直线 |
| `OCV_画矩形` | `func OCV_画矩形(矩阵 gocv.Mat, 矩形 image.Rectangle, 颜色 color.RGBA, 线宽 int) gocv.Mat` | 画矩形 |
| `OCV_画圆` | `func OCV_画圆(矩阵 gocv.Mat, 圆心 image.Point, 半径 int, 颜色 color.RGBA, 线宽 int) gocv.Mat` | 画圆 |
| `OCV_画文字` | `func OCV_画文字(矩阵 gocv.Mat, 文字 string, 位置 image.Point, 字体 gocv.HersheyFont, 大小 float64, 颜色 color.RGBA, 线宽 int) gocv.Mat` | 画文字 |
| `OCV_画椭圆` | `func OCV_画椭圆(矩阵 gocv.Mat, 圆心 image.Point, 长轴, 短轴 int, 旋转角, 起始角, 终止角 float64, 颜色 color.RGBA, 线宽 int) gocv.Mat` | 画椭圆 |

### 图像运算

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_加法` | `func OCV_加法(矩阵1, 矩阵2 gocv.Mat) gocv.Mat` | 逐像素相加 |
| `OCV_加权加法` | `func OCV_加权加法(矩阵1 gocv.Mat, 权重1 float64, 矩阵2 gocv.Mat, 权重2, 伽马值 float64) gocv.Mat` | 按权重相加 |
| `OCV_减法` | `func OCV_减法(矩阵1, 矩阵2 gocv.Mat) gocv.Mat` | 逐像素相减 |
| `OCV_按位与` | `func OCV_按位与(矩阵1, 矩阵2 gocv.Mat) gocv.Mat` | 按位与 |
| `OCV_按位或` | `func OCV_按位或(矩阵1, 矩阵2 gocv.Mat) gocv.Mat` | 按位或 |
| `OCV_按位异或` | `func OCV_按位异或(矩阵1, 矩阵2 gocv.Mat) gocv.Mat` | 按位异或 |
| `OCV_按位取反` | `func OCV_按位取反(矩阵 gocv.Mat) gocv.Mat` | 按位取反 |

### 视频与摄像头

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_打开摄像头` | `func OCV_打开摄像头(设备ID int) (*gocv.VideoCapture, error)` | 打开摄像头 |
| `OCV_打开视频文件` | `func OCV_打开视频文件(文件路径 string) (*gocv.VideoCapture, error)` | 打开视频文件 |
| `OCV_读取帧` | `func OCV_读取帧(捕获器 *gocv.VideoCapture) (gocv.Mat, bool)` | 读取一帧 |
| `OCV_创建视频写入器` | `func OCV_创建视频写入器(文件路径, 编码器 string, 帧率 float64, 宽度, 高度 int) (*gocv.VideoWriter, error)` | 创建视频写入器 |
| `OCV_写入帧` | `func OCV_写入帧(写入器 *gocv.VideoWriter, 帧 gocv.Mat)` | 写入一帧 |

### 图像转换

| 函数 | 签名 | 说明 |
|------|------|------|
| `OCV_Mat转Image` | `func OCV_Mat转Image(矩阵 gocv.Mat) image.Image` | Mat → image.Image |
| `OCV_Image转Mat` | `func OCV_Image转Mat(图片 image.Image) (gocv.Mat, error)` | image.Image → Mat |

**示例**：

```go
// 读取图片并处理
img, _ := OCV_读取图片("photo.jpg")
defer img.Close()

gray := OCV_BGR转灰度(img)
defer gray.Close()

edges := OCV_Canny边缘检测(gray, 50, 150)
defer edges.Close()

OCV_保存图片(edges, "edges.jpg")

// 缩放
resized := OCV_缩放(img, 640, 480)
defer resized.Close()

// 查找轮廓
contours := OCV_查找轮廓(edges)
for _, c := range contours {
    area := OCV_轮廓面积(c)
    if area > 1000 {
        rect := OCV_轮廓外接矩形(c)
        fmt.Println("检测到区域:", rect, "面积:", area)
    }
}
```

> ⚠️ **注意**：OCV视觉模块依赖 OpenCV 4.x C++ 库，使用前需安装 OpenCV。详见 [gocv 安装指南](https://gocv.io/getting-started/)

| 模块 | 文件数 | 函数/方法数 |
|------|--------|------------|
| class | 8 | 75 |
| utils/核心库 | 1 | 16 |
| utils/辅助 | 1 | 8 |
| utils/B编码 | 1 | 49 |
| utils/C程序 | 1 | 12 |
| utils/Float64转换 | 1 | 11 |
| utils/H汇编 | 1 | 1 |
| utils/IP | 1 | 10 |
| utils/Int转换 | 1 | 2 |
| utils/J校验 | 1 | 24 |
| utils/Post数据类 | 1 | 11 |
| utils/Map | 1 | 4 |
| utils/M目录 | 1 | 6 |
| utils/Rsa | 1 | 7 |
| utils/S数组 | 1 | 16 |
| utils/S时间 | 1 | 13 |
| utils/T图片 | 1 | 42 |
| utils/W文件 | 1 | 15 |
| utils/W文本 | 1 | 42 |
| utils/W网页 | 1 | 6 |
| utils/Y原子 | 1 | 2 |
| utils/Z字节集 | 1 | 4 |
| utils/Z正则 | 1 | 11 |
| utils/Jjson | 1 | 16 |
| utils/C类型转换 | 1 | 14 |
| utils/P配置 | 1 | 15 |
| utils/E邮件 | 1 | 3 |
| utils/X系统信息 | 1 | 28 |
| utils/D定时 | 1 | 7 |
| utils/G协程池 | 1 | 10 |
| utils/L日志 | 1 | 11 |
| utils/K环境变量 | 1 | 9 |
| utils/M命令行 | 1 | 9 |
| utils/R日期解析 | 1 | 9 |
| utils/P对象池 | 1 | 4 |
| utils/B表达式计算 | 1 | 6 |
| utils/T模板 | 1 | 5 |
| utils/V数据校验 | 1 | 10 |
| utils/J结构体合并 | 1 | 7 |
| utils/K表格 | 1 | 20 |
| utils/F文件监控 | 1 | 9 |
| utils/X消息总线 | 1 | 5 |
| utils/H客户端 | 1 | 12 |
| utils/N键值库 | 1 | 8 |
| utils/C爬虫 | 1 | 10 |
| utils/Q权限管理 | 1 | 12 |
| utils/D数据库 | 1 | 19 |
| utils/C窗口 | 1 | 24 |
| utils/C进程 | 1 | 15 |
| utils/OCV视觉 | 1 | 79 |
| **合计** | **57** | **712+** |
