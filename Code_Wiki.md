# EFunc Code Wiki

> Go 语言版的精易模块 — 一套面向中文开发者的通用工具函数库

---

## 一、项目概览

| 项目 | 说明 |
|------|------|
| 模块名 | `github.com/yuan71058/Efunc` |
| Go 版本 | 1.22+ |
| 仓库地址 | https://github.com/yuan71058/Efunc |
| 设计理念 | 以中文命名函数，降低中文开发者使用门槛；命名规则为 **首字拼音大写 + 功能描述**（如 `W文本_取出中间文本`） |
| 核心依赖 | `github.com/gogf/gf/v2`（GoFrame 工具库）、`github.com/shopspring/decimal`（高精度十进制）、`github.com/axgle/mahonia`（编码转换）、`github.com/skip2/go-qrcode`（二维码生成） |

---

## 二、项目架构

```
EFunc/
├── main.go                  # 入口文件（示例 & 测试）
├── go.mod                   # Go 模块定义
├── go.sum                   # 依赖校验
├── README.md                # 项目说明
│
├── class/                   # 类（结构体）定义模块
│   ├── 类_临界许可.go        # 互斥锁封装
│   ├── 类_读写锁.go          # 读写锁封装
│   ├── 类_正则表达式.go      # 正则表达式封装
│   ├── 类_队列.go            # 线程安全队列（interface{}）
│   └── 类_队列泛型.go        # 泛型线程安全队列
│
└── utils/                   # 工具函数模块（核心）
    ├── utils.go             # 包声明
    ├── 核心库.go             # 类型转换 & 通用工具
    ├── 辅助.go              # 简易辅助函数
    ├── B编码.go              # 编码/解码（URL、Base64、USC2）
    ├── C程序.go              # 程序控制（延时、GUID、日志、命令行等）
    ├── Float64转换.go        # 高精度浮点数运算
 │   ├── H汇编_汇编器.go       # x86 机器码构建器（动态生成执行机器码）
│   ├── H汇编_汇编指令.go      # x86 汇编指令生成（MOV/CALL/JMP/ADD 等链式调用）
│   ├── H内存.go              # 内存操作（进程内存读写、搜索、分配释放、进程枚举）
│   ├── HHook.go              # MinHook API Hook 封装
│   ├── HHook/                # MinHook C 库 & cgo 包装
│   │   ├── HHook.go          # cgo 包装层
│   │   ├── MinHook.h         # MinHook 主头文件
│   │   ├── hook.c            # Hook 核心实现
│   │   └── ...               # 其他 C 源文件
│   ├── J键鼠.go              # 键盘鼠标操作（模拟按键、鼠标移动点击、虚拟键码常量）
│   ├── X线程.go              # 多线程（线程创建/挂起/恢复、临界区、事件、互斥体、信号量）
│   ├── X系统命令.go           # 系统命令（关机/重启/注销、剪辑版读写、消息框、环境变量）
    ├── IP.go                 # IP 地址转换
    ├── Int转换.go            # 整数工具
    ├── J校验.go              # 校验摘要（MD5、CRC32、SHA 系列）
    ├── J加解密.go            # 对称加解密（AES/DES/3DES/RC4/XOR/TEA/XXTEA）
    ├── L类_post数据类.go     # POST 数据构造器
    ├── Map.go                # Map 工具
    ├── M目录.go              # 目录操作
    ├── Rsa.go                # RSA 加解密/签名
    ├── S数组.go              # 数组/切片工具
    ├── S时间.go              # 时间/时间戳工具
    ├── T图片.go              # 图片工具（二维码）
    ├── W文件.go              # 文件操作
    ├── W文本.go              # 文本处理（最大模块）
    ├── W网页.go              # HTTP 网页访问
    ├── Y原子.go              # 原子操作
    ├── Z字节集.go            # 字节集工具
    ├── Z正则.go              # 正则校验与提取
    └── 工具测试_test.go      # 单元测试
```

---

## 三、模块详解

### 3.1 class — 类定义模块

提供面向对象封装的并发安全数据结构，所有类均以 `L_` 前缀命名。

#### 3.1.1 L_临界许可（互斥锁）

| 文件 | [类_临界许可.go](class/类_临界许可.go) |
|------|------|
| 结构体 | `L_临界许可` |
| 底层实现 | `sync.Mutex` |

| 方法 | 说明 |
|------|------|
| `J进入许可区()` | 加锁，进入临界区 |
| `T退出许可区()` | 解锁，退出临界区 |
| `C尝试进入() bool` | 尝试加锁（非阻塞），成功返回 `true` |

---

#### 3.1.2 L_读写锁

| 文件 | [类_读写锁.go](class/类_读写锁.go) |
|------|------|
| 结构体 | `L_读写锁` |
| 底层实现 | `sync.RWMutex` |

| 方法 | 说明 |
|------|------|
| `K开始读()` | 获取读锁 |
| `J结束读()` | 释放读锁 |
| `K开始写()` | 获取写锁 |
| `J结束写()` | 释放写锁 |

---

#### 3.1.3 L_正则表达式

| 文件 | [类_正则表达式.go](class/类_正则表达式.go) |
|------|------|
| 结构体 | `L_正则表达式` |
| 底层实现 | `regexp.Regexp` |

| 字段 | 类型 | 说明 |
|------|------|------|
| `Count` | `int` | 匹配数量 |
| `SubmatchCount2` | `int` | 子匹配数量 |

| 方法 | 说明 |
|------|------|
| `New正则表达式类(正则, 文本) (*L_正则表达式, bool)` | 构造函数 |
| `E创建(正则, 文本) bool` | 初始化正则并执行匹配 |
| `Q取匹配数量() int` | 返回匹配结果数 |
| `Q取匹配文本(索引) string` | 取指定索引的完整匹配文本 |
| `Q取子匹配文本(匹配索引, 子表达式索引) string` | 取子匹配文本，越界返回空串 |
| `Q取子匹配数量() int` | 返回子匹配数量 |
| `GetResult() [][]string` | 返回原始匹配结果二维数组 |

---

#### 3.1.4 L_队列（线程安全队列）

| 文件 | [类_队列.go](class/类_队列.go) |
|------|------|
| 结构体 | `L_队列` |
| 底层实现 | `container/list` + `sync.Mutex` |

| 方法 | 说明 |
|------|------|
| `Init()` | 初始化队列 |
| `J加入队列(v interface{}) int` | 入队，返回队列长度 |
| `T弹出队列() (interface{}, bool)` | 弹出队尾元素 |
| `T弹出队列文本(值 *string) bool` | 弹出文本类型元素 |
| `T弹出队列整数(值 *int) bool` | 弹出整数类型元素 |
| `Q取队列长度() int` | 获取队列长度 |
| `Q清空队列() interface{}` | 清空队列 |
| `Dump()` | 打印队列内容（调试用） |

---

#### 3.1.5 L_队列泛型（泛型线程安全队列）

| 文件 | [类_队列泛型.go](class/类_队列泛型.go) |
|------|------|
| 结构体 | `L_队列泛型[T any]` |
| 底层实现 | `container/list` + `sync.Mutex` |
| Go 版本要求 | 1.18+（泛型支持） |

| 方法 | 说明 |
|------|------|
| `Init()` | 初始化队列 |
| `J加入队列(v T) int` | 入队，返回队列长度 |
| `T弹出队列() (T, bool)` | 弹出队尾元素，失败返回零值 |
| `Q取队列长度() int` | 获取队列长度 |
| `Q清空队列()` | 清空队列 |
| `Dump()` | 打印队列内容（调试用） |

---

#### 3.1.6 L_TCP服务端 / L_TCP客户端

| 文件 | [类_TCP.go](class/类_TCP.go) |
|------|------|

**L_TCP服务端**

| 结构体 | `L_TCP服务端` |
|------|------|
| 底层实现 | `net.ListenTCP` + `bufio.Reader` |
| 协议特点 | 换行符（`\n`）作为消息分隔符，30 秒读超时 |

| 方法 | 说明 |
|------|------|
| `Q启动(端口 int) error` | 启动服务端，监听端口 |
| `T停止()` | 停止服务端并关闭所有连接 |
| `F发送数据/文本(客户端地址, 数据)` | 向指定客户端发送 |
| `G广播数据/文本(数据)` | 向所有客户端广播 |
| `Q取客户端数量() int` | 获取连接数 |
| `Q取客户端列表() []string` | 获取所有客户端地址 |

| 回调字段 | 类型 | 说明 |
|------|------|------|
| `S收到数据回调` | `func(客户端地址, 数据[]byte)` | 收到数据 |
| `K客户端连接回调` | `func(客户端地址)` | 新连接 |
| `K客户端断开回调` | `func(客户端地址)` | 断开 |

**L_TCP客户端**

| 结构体 | `L_TCP客户端` |
|------|------|

| 方法 | 说明 |
|------|------|
| `L连接(地址) error` | 连接服务端（10s 超时） |
| `D断开()` | 断开连接 |
| `F发送数据/文本(数据)` | 发送数据 |
| `S是否已连接() bool` | 连接状态 |
| `Q取本地地址() string` | 本端地址 |

---

#### 3.1.7 L_WS服务端 / L_WS客户端

| 文件 | [类_WebSocket.go](class/类_WebSocket.go) |
|------|------|
| 依赖 | `github.com/gorilla/websocket` |

**L_WS服务端**

| 结构体 | `L_WS服务端` |
|------|------|
| 底层实现 | HTTP 升级协议 → gorilla/websocket |

| 方法 | 说明 |
|------|------|
| `Q启动(端口, 路径)` / `Q启动带地址(地址, 路径)` | 启动服务端 |
| `T停止()` | 停止服务端 |
| `F发送文本/字节(客户端ID, 数据)` | 向指定客户端发送 |
| `G广播文本/字节(数据)` | 广播 |
| `Q取客户端数量/列表()` | 连接管理 |

| 回调字段 | 类型 |
|------|------|
| `S收到文本/字节回调` | `func(客户端ID, 数据)` |
| `K客户端连接/断开回调` | `func(客户端ID)` |

**L_WS客户端**

| 方法 | 说明 |
|------|------|
| `L连接(URL) error` | 连接（如 `ws://host/ws`） |
| `D断开()` | 断开 |
| `F发送文本/字节/JSON(数据)` | 发送消息 |
| `S是否已连接() bool` | 连接状态 |
| `Q取远程地址() string` | 远端地址 |

---

#### 3.1.8 L_HTTP服务端

| 文件 | [类_HTTP.go](class/类_HTTP.go) |
|------|------|
| 底层实现 | `net/http.ServeMux` + 中间件链 |

| 方法 | 说明 |
|------|------|
| `Q启动(端口)` / `Q启动带地址(地址)` | 启动服务端 |
| `T停止() error` | 优雅停止（5s 超时） |
| `T注册路由(Method, Path, Handler)` | 按 Method 注册 |
| `T注册通用路由(Path, Handler)` | 通配路由 |
| `J静态文件服务(URL前缀, 目录)` | 静态文件 |
| `Z中间件(func)` | 自定义中间件 |
| `Z中间件CORS()` | 跨域中间件 |
| `Z中间件日志()` | 请求日志 |
| `Q取启动地址() / S是否运行中()` | 状态查询 |

**包级辅助函数**：

| 函数 | 说明 |
|------|------|
| `F响应JSON(w, code, data)` | JSON 响应 |
| `F响应文本/HTML(w, code, text)` | 文本/HTML 响应 |
| `F取查询参数(r, key)` | URL 查询参数 |
| `F取POST参数(r, key)` | POST 表单参数 |
| `F解析JSON请求体(r, target)` | JSON Body 解析 |

---

### 3.2 utils — 工具函数模块

所有函数以 **拼音首字母大写 + 分类名 + 下划线 + 功能名** 命名，导入后通过 `. "github.com/yuan71058/Efunc/utils"` 直接调用。

#### 3.2.1 核心库（类型转换 & 通用工具）

| 文件 | [核心库.go](utils/核心库.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `D到字节集` | `(interface{}) []byte` | 任意类型转字节集 |
| `D到字节` | `(interface{}) byte` | 任意类型转 byte |
| `D到整数` | `(interface{}) int` | 任意类型转 int |
| `D到整数64` | `(interface{}) int64` | 任意类型转 int64 |
| `D到数值` | `(interface{}) float64` | 任意类型转 float64 |
| `D到文本` | `(interface{}) string` | 任意类型转 string |
| `D到结构体` | `(interface{}, interface{}) error` | 任意类型转结构体 |
| `S三元` | `[T any](bool, T, T) T` | 泛型三元运算 |
| `D多项选择` | `[T any](int, []T, T) T` | 泛型多项选择 |
| `G格式化文本` | `(string, ...interface{}) string` | 格式化文本（Sprintf 封装） |
| `G格式化_JSON` | `(string) string` | JSON 格式化输出 |
| `D到文本数组` | `(interface{}) []string` | 通用型变量转文本数组 |
| `S是否为数组` | `(interface{}) bool` | 判断变量是否为数组/切片 |
| `W文本到utf8` | `(string) string` | GBK 文本转 UTF-8 |
| `Utf8到文本` | `(string) string` | UTF-8 文本转 GBK |
| `Q取随机数` | `(int, int) int` | 取随机整数 |

> 底层依赖 `github.com/gogf/gf/v2/util/gconv` 实现高性能类型转换。

---

#### 3.2.2 辅助（简易函数）

| 文件 | [辅助.go](utils/辅助.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `选择` | `(bool, interface{}, interface{}) interface{}` | 三元选择（非泛型版） |
| `取随机数` | `(int, int) int` | 取随机整数 |
| `取文本右边` | `(string, int) string` | 取文本右侧 N 个字符 |
| `取文本左边` | `(string, int) string` | 取文本左侧 N 个字符 |
| `加入成员` | `([]string, string) []string` | 向数组追加成员 |
| `删首尾空` | `(string) string` | TrimSpace 封装 |
| `取文本长度` | `(string) int` | 取文本字节长度 |
| `分割文本` | `(string, string) []string` | 按分割符分割文本 |

---

#### 3.2.3 B编码（编码/解码）

| 文件 | [B编码.go](utils/B编码.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `B编码_URL编码` | `(string) string` | URL 编码 |
| `B编码_URL解码` | `(string) string` | URL 解码 |
| `B编码_usc2到文本` | `(string) string` | USC2 转中文文本 |
| `B编码_BASE64编码` | `([]byte) string` | Base64 编码 |
| `B编码_BASE64解码` | `(string) []byte` | Base64 解码 |

---

#### 3.2.4 C程序（程序控制）

| 文件 | [C程序.go](utils/C程序.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `C程序_延时` | `(int64) bool` | 毫秒级延时（Sleep） |
| `C程序_延时2` | `(int) bool` | 毫秒级延时（逐毫秒循环） |
| `C程序_取cmd路径` | `() string` | 获取 cmd.exe 路径 |
| `C程序_取GUID` | `() string` | 生成标准 V4 GUID |
| `C程序_删除自身` | `() error` | 删除当前可执行文件 |
| `C程序_是否被调试` | `() bool` | 检测是否被调试 |
| `C程序_禁止重复运行` | `() bool` | 通过环境变量防止重复运行 |
| `C程序_写日志` | `(string, string)` | 写入日志文件（自动追加时间戳） |
| `C程序_取命令行` | `() []string` | 获取命令行参数 |
| `C程序_取运行目录` | `() string` | 获取可执行文件所在目录 |
| `C程序_取临时目录` | `() string` | 获取系统临时目录 |
| `C程序_运行Win` | `(string) string` | 执行 PowerShell 命令（GBK→UTF-8） |

---

#### 3.2.5 Float64转换（高精度浮点运算）

| 文件 | [Float64转换.go](utils/Float64转换.go) |
|------|------|
| 核心依赖 | `github.com/shopspring/decimal` |

| 函数 | 签名 | 说明 |
|------|------|------|
| `Float64取绝对值` | `(float64) float64` | 高精度取绝对值 |
| `Float64乘int64` | `(float64, int64) float64` | 高精度乘法 |
| `Float64乘Float64` | `(float64, float64) float64` | 高精度乘法 |
| `Float64除int64` | `(float64, int64, int32) float64` | 高精度除法（指定保留位数） |
| `Float64除float64` | `(float64, float64, int32) float64` | 高精度除法 |
| `Float64取负值` | `(float64) float64` | 高精度取负值 |
| `Float64到文本` | `(float64, int) string` | 浮点数转文本（指定小数位） |
| `Float64从文本` | `(string, int) float64` | 文本转浮点数（指定小数位） |
| `Int64到Float64` | `(int64) float64` | 整数转浮点数 |
| `Float64减float64` | `(float64, float64, int32) float64` | 高精度减法 |
| `Float64加float64` | `(float64, float64, int32) float64` | 高精度加法 |

> 所有运算均使用 `decimal` 库，避免浮点精度丢失。

---

#### 3.2.6 H汇编（汇编器：机器码构建 & 执行）

| 文件 | [H汇编_汇编器.go](utils/H汇编_汇编器.go)、[H汇编_汇编指令.go](utils/H汇编_汇编指令.go) |
|------|------|

> ⚠️ 汇编器模块为 **Windows 专用**（`//go:build windows`），使用 `VirtualAlloc`/`VirtualAllocEx` + `CreateRemoteThread`。

**随机数函数**（已迁移至 `核心库.go`）：

| 函数 | 签名 | 说明 |
|------|------|------|
| `H汇编_取随机数` | `(起始数, 结束数 int) int` | 生成 [起始数, 结束数] 范围内的随机整数（线程安全） |
| `H汇编_取随机字节` | `(长度 int) ([]byte, error)` | 使用 `crypto/rand` 生成安全随机字节序列 |

##### 3.2.6.1 汇编器核心（机器码构建 & 执行）

| 函数 | 签名 | 说明 |
|------|------|------|
| `H汇编_置代码` | `(代码 []byte)` | 初始化/重置代码缓冲区 |
| `H汇编_取代码` | `() []byte` | 获取当前构建的机器码字节集 |
| `H汇编_运行汇编代码` | `() (uintptr, error)` | 在当前进程分配可执行内存并运行机器码，返回 EAX 值 |
| `H汇编_远程执行汇编代码` | `(进程ID uint32, 代码 []byte) error` | 远程注入目标进程执行机器码（VirtualAllocEx + CreateRemoteThread） |

##### 3.2.6.2 汇编指令构建（x86 机器码生成）

**栈操作**：`H汇编_PUSHAD`(60)、`H汇编_POPAD`(61)、`H汇编_LEAVE`(C9)、`H汇编_NOP`(90)、`H汇编_RET`(C3)

**MOV 立即数→寄存器**：`H汇编_MOV_EAX(v)` ~ `H汇编_MOV_EBP(v)`（8 寄存器）

**MOV [addr] ← EAX**：`H汇编_MOV_DWORD_PTR_EAX(addr)` → `mov [addr], eax`

**MOV 寄存器 ← [addr]**：`H汇编_MOV_EAX_DWORD_PTR(addr)` 等 8 个函数

**MOV 寄存器 ← [寄存器]**：`H汇编_MOV_EAX_DWORD_PTR_EAX()` 等 8 个函数

**ADD/CALL/JMP/CMP/INC/DEC/IDIV/IMUL** 等更多指令详见源码

> 仅支持 Windows x86_64 平台。

---

#### 3.2.7 H内存（内存操作）

| 文件 | [H内存.go](utils/H内存.go) |
|------|------|

> ⚠️ Windows 专用模块，依赖 kernel32.dll API 实现进程内存读写、AOB 搜索、进程枚举等功能。

**进程内存读写**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `H内存_打开进程` | `(进程ID uint32, 权限 uint32) (syscall.Handle, error)` | 打开目标进程并获取句柄 |
| `H内存_读整数` | `(hProcess syscall.Handle, 地址 uintptr) (int32, error)` | 读取 32 位有符号整数 |
| `H内存_读整数64` | `(hProcess syscall.Handle, 地址 uintptr) (int64, error)` | 读取 64 位有符号整数 |
| `H内存_读字节集` | `(hProcess syscall.Handle, 地址 uintptr, 长度 uint32) ([]byte, error)` | 读取原始字节数据 |
| `H内存_读浮点数` | `(hProcess syscall.Handle, 地址 uintptr) (float32, error)` | 读取 32 位浮点数 |
| `H内存_读浮点数64` | `(hProcess syscall.Handle, 地址 uintptr) (float64, error)` | 读取 64 位浮点数 |
| `H内存_写整数` | `(hProcess syscall.Handle, 地址 uintptr, 值 int32) error` | 写入 32 位整数 |
| `H内存_写整数64` | `(hProcess syscall.Handle, 地址 uintptr, 值 int64) error` | 写入 64 位整数 |
| `H内存_写字节集` | `(hProcess syscall.Handle, 地址 uintptr, 数据 []byte) error` | 写入原始字节数据 |
| `H内存_关闭句柄` | `(hProcess syscall.Handle) error` | 关闭进程句柄 |

**内存搜索**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `H内存_搜索字节集` | `(hProcess syscall.Handle, 特征码 []byte, 起始地址, 长度 uintptr) ([]uintptr, error)` | AOB 特征码搜索（0xFF 为通配符） |
| `H内存_搜索文本` | `(hProcess syscall.Handle, 文本 string, 起始地址, 长度 uintptr) ([]uintptr, error)` | 搜索文本字符串 |

**虚拟内存管理**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `H内存_分配内存` | `(hProcess syscall.Handle, 大小 uintptr) (uintptr, error)` | 在目标进程分配可执行读写内存 |
| `H内存_释放内存` | `(hProcess syscall.Handle, 地址 uintptr) error` | 释放远程内存 |
| `H内存_修改保护` | `(hProcess syscall.Handle, 地址, 大小 uintptr, 新保护 uint32) (uint32, error)` | 修改内存保护属性 |
| `H内存_查询内存` | `(hProcess syscall.Handle, 地址 uintptr) (*MEMORY_BASIC_INFORMATION, error)` | 查询内存区域信息 |

**进程枚举与管理**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `H内存_取进程ID` | `(进程名 string) (uint32, error)` | 通过进程名获取 PID |
| `H内存_枚举进程` | `() ([]PROCESSENTRY32W, error)` | 枚举所有运行进程 |
| `H内存_取模块基址` | `(进程ID uint32, 模块名 string) (uintptr, error)` | 获取模块加载基址 |
| `H内存_终止进程` | `(进程ID uint32) error` | 强制终止进程 |
| `H内存_取进程名称` | `(进程ID uint32) (string, []uint16, error)` | 获取进程对应可执行文件名 |
| `H内存_是否存在进程` | `(进程ID uint32) bool` | 检查进程是否运行 |
| `H内存_取进程位数` | `(进程ID uint32) int` | 检测进程是 32 位还是 64 位 |

---

#### 3.2.8 J键鼠（键盘鼠标操作）

| 文件 | [J键鼠.go](utils/J键鼠.go) |
|------|------|

> ⚠️ Windows 专用模块，基于 user32.dll API 实现模拟输入与状态检测。

**键盘操作**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J键鼠_按键` | `(虚拟键码 int, 按下 bool)` | 模拟按键按下/弹起 |
| `J键鼠_按键组合` | `(虚拟键码1, 虚拟键码2 int)` | 模拟组合键（如 Ctrl+C） |
| `J键鼠_模拟按键` | `(虚拟键码 int)` | 模拟完整按下+弹起 |
| `J键鼠_模拟文本输入` | `(text string)` | 逐字符模拟中文/英文输入 |
| `J键鼠_取按键状态` | `(虚拟键码 int) int16` | 异步获取按键状态 |
| `J键鼠_是否按下` | `(虚拟键码 int) bool` | 判断按键是否按下 |

**鼠标操作**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J键鼠_移动鼠标` | `(x, y int)` | 移动光标到指定坐标 |
| `J键鼠_取鼠标位置` | `() (int, int)` | 获取光标坐标 |
| `J键鼠_鼠标左键单击` | `()` | 左键单击 |
| `J键鼠_鼠标左键按下` | `()` | 左键按下 |
| `J键鼠_鼠标左键弹起` | `()` | 左键弹起 |
| `J键鼠_鼠标右键单击` | `()` | 右键单击 |
| `J键鼠_鼠标中键单击` | `()` | 中键单击 |
| `J键鼠_鼠标滚轮` | `(delta int)` | 滚动滚轮 |

**其他**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J键鼠_取屏幕宽度` | `() int` | 主屏幕宽度 |
| `J键鼠_取屏幕高度` | `() int` | 主屏幕高度 |
| `J键鼠_锁定输入` | `(锁定 bool)` | 锁定/解锁键盘鼠标 |

> 文件内置完整的 Win32 虚拟键码常量（VK_*），支持所有标准键（字母 A-Z、F1-F12、方向键、小键盘等）。

---

#### 3.2.9 X线程（多线程）

| 文件 | [X线程.go](utils/X线程.go) |
|------|------|

> ⚠️ Windows 专用模块，基于 kernel32.dll API。

**线程创建与管理**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X线程_创建` | `(线程函数 uintptr, 参数 uintptr) (syscall.Handle, error)` | 创建内核级线程 |
| `X线程_取当前ID` | `() uint32` | 获取当前线程 ID |
| `X线程_挂起` | `(hThread syscall.Handle) (uint32, error)` | 挂起线程 |
| `X线程_恢复` | `(hThread syscall.Handle) (uint32, error)` | 恢复线程 |
| `X线程_终止` | `(hThread syscall.Handle, 退出码 uint32) error` | 强制终止线程 |
| `X线程_关闭句柄` | `(handle syscall.Handle) error` | 关闭内核句柄 |
| `X线程_延时` | `(毫秒 uint32)` | Sleep 延时 |
| `X线程_等待单个` | `(handle syscall.Handle, 超时毫秒 uint32) uint32` | 等待单个对象 |
| `X线程_等待多个` | `(handles []syscall.Handle, 等待全部 bool, 超时毫秒 uint32) uint32` | 等待多个对象 |

**临界区**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X线程_临界区_创建` | `() (*CRITICAL_SECTION, error)` | 创建临界区 |
| `X线程_临界区_销毁` | `(cs *CRITICAL_SECTION)` | 销毁临界区 |
| `X线程_临界区_进入` | `(cs *CRITICAL_SECTION)` | 进入临界区 |
| `X线程_临界区_离开` | `(cs *CRITICAL_SECTION)` | 离开临界区 |

**事件（Event）**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X线程_事件_创建` | `(手动重置, 初始状态 bool, 名称 string) (syscall.Handle, error)` | 创建事件 |
| `X线程_事件_设置` | `(hEvent syscall.Handle) error` | 设置事件有信号 |
| `X线程_事件_重置` | `(hEvent syscall.Handle) error` | 重置事件无信号 |
| `X线程_事件_脉冲` | `(hEvent syscall.Handle) error` | 脉冲触发 |

**互斥体（Mutex）**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X线程_互斥体_创建` | `(名称 string) (syscall.Handle, bool, error)` | 创建互斥体（防多开） |
| `X线程_互斥体_打开` | `(名称 string) (syscall.Handle, error)` | 打开已有互斥体 |
| `X线程_互斥体_释放` | `(hMutex syscall.Handle) error` | 释放互斥体 |

**信号量（Semaphore）**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X线程_信号量_创建` | `(初始计数, 最大计数 int32, 名称 string) (syscall.Handle, error)` | 创建信号量 |
| `X线程_信号量_打开` | `(名称 string) (syscall.Handle, error)` | 打开已有信号量 |
| `X线程_信号量_释放` | `(hSemaphore syscall.Handle, 释放数量 int32) (int32, error)` | 释放信号量 |

**Goroutine 封装**（G线程）：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X线程_协程_创建` | `(执行函数 func()) *G线程` | 创建协作式线程 |
| `启动` / `暂停` / `继续` / `退出` | goroutine 生命周期控制 |

---

#### 3.2.10 X系统命令（系统命令）

| 文件 | [X系统命令.go](utils/X系统命令.go) |
|------|------|

> ⚠️ Windows 专用模块，整合 user32.dll / kernel32.dll / advapi32.dll API。

**关机/重启/注销**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_关机` | `(强制 bool) error` | 关机（支持强制） |
| `X系统_重启` | `(强制 bool) error` | 重启（支持强制） |
| `X系统_注销` | `(强制 bool) error` | 注销当前用户 |
| `X系统_锁定工作站` | `() error` | 锁定屏幕 |
| `X系统_启用关机权限` | `() error` | 获取 SE_SHUTDOWN_NAME 权限 |
| `X系统_远程关机` | `(计算机名, 消息 string, 超时秒 uint32, 强制 bool) error` | 远程关机 |

**剪辑版操作**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_置剪辑版文本` | `(text string) error` | 设置文本到剪贴板 |
| `X系统_取剪辑版文本` | `() (string, error)` | 读取剪贴板文本 |
| `X系统_清空剪辑版` | `() error` | 清空剪贴板 |

**消息框**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_信息框` | `(标题, 内容 string)` | 弹出信息框（MB_OK） |
| `X系统_信息框_确认` | `(标题, 内容 string) int` | 确认框（OK/Cancel） |
| `X系统_信息框_是否` | `(标题, 内容 string) int` | 是否框（Yes/No） |

**系统命令 & 工具**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `X系统_执行命令` | `(命令 string, 参数 ...string) (string, error)` | 执行命令并返回输出 |
| `X系统_执行命令隐藏` | `(命令 string, 参数 ...string) error` | 后台执行（不显示窗口） |
| `X系统_取计算机名` | `() (string, error)` | 获取本机计算机名 |
| `X系统_取用户名` | `() (string, error)` | 获取当前用户名 |
| `X系统_禁止屏幕保护` | `(禁止 bool) error` | 禁用/启用屏保 |
| `X系统_置屏保超时` | `(秒 int) error` | 设置屏保超时 |
| `X系统_置环境变量` | `(名称, 值 string) error` | 设置环境变量 |
| `X系统_取环境变量` | `(名称 string) string` | 获取环境变量 |
| `X系统_删除环境变量` | `(名称 string) error` | 删除环境变量 |

---

#### 3.2.7 IP（IP 地址工具）

| 文件 | [IP.go](utils/IP.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `IP_10进制转IP` | `(int) string` | 10 进制整数转 IP 地址 |

---

#### 3.2.8 Int转换（整数工具）

| 文件 | [Int转换.go](utils/Int转换.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `Int取绝对值` | `(int) int` | 整数取绝对值 |
| `Int32ToBytes` | `(int32) []byte` | int32 转大端字节集 |

---

#### 3.2.9 J校验（校验摘要）

| 文件 | [J校验.go](utils/J校验.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `J校验_取md5` | `([]byte, bool) string` | 计算 MD5（可返回大写） |
| `J校验_取md5_文本` | `(string, bool) string` | 文本 MD5 |
| `J校验_取Crc32` | `([]byte, bool) string` | 计算 CRC32（16 进制） |
| `J校验_取sha1` | `([]byte, bool) string` | 计算 SHA1（40 位） |
| `J校验_取sha256` | `([]byte, bool) string` | 计算 SHA256 |
| `J校验_取sha512` | `([]byte, bool) string` | 计算 SHA512 |

---

#### 3.2.10 Post数据类（POST 数据构造器）

| 文件 | [L类_post数据类.go](utils/L类_post数据类.go) |
|------|------|
| 结构体 | `Post数据类` |

| 方法 | 说明 |
|------|------|
| `T添加(key, value string, 转码 bool)` | 添加键值对（可选 URL 编码） |
| `T添加_批量(文本 string, 转码 bool)` | 批量添加（`&` 分隔的 key=value 格式） |
| `Q取值(key string) string` | 按 key 取值 |
| `Z置值(key, value string)` | 设置/更新键值 |
| `H获取Post数据(是否URL编码 bool) string` | 生成 POST 请求体 |
| `H获取协议头数据(是否URL编码 bool) string` | 生成 HTTP 协议头 |
| `H获取Key数组() []string` | 获取所有 key |
| `H获取Value数组() []string` | 获取所有 value |
| `H获取JSON文本() string` | 生成 JSON 格式文本 |
| `Q清空()` | 清空数据 |
| `S删除(key string)` | 按 key 删除 |

---

#### 3.2.11 Map（Map 工具）

| 文件 | [Map.go](utils/Map.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `Map_取key整数数组` | `(map[int]string) []int` | 提取 int 类型的 key 数组 |
| `Map_Struct转Map` | `(interface{}) map[string]interface{}` | 结构体转 Map（反射） |
| `Map_转post数据` | `(map[string]string, bool) string` | Map 转 POST 请求参数 |
| `Map_键名是否存在` | `(map[int]string, int) bool` | 判断 key 是否存在 |

---

#### 3.2.12 M目录（目录操作）

| 文件 | [M目录.go](utils/M目录.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `M目录_是否存在` | `(string) (bool, error)` | 判断目录是否存在 |
| `M目录_创建` | `(string) error` | 递归创建目录 |
| `M目录_枚举子目录` | `(string, *[]string, bool, bool) error` | 枚举子目录（可选递归、带路径） |
| `M目录_取运行目录` | `() string` | 获取可执行文件目录 |
| `M目录_取当前目录` | `() string` | 获取当前工作目录 |
| `M目录_删除` | `(string) error` | 递归删除目录 |

---

#### 3.2.13 Rsa（RSA 加解密）

| 文件 | [Rsa.go](utils/Rsa.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `Rsa_私钥签名` | `(string, string) string` | RSA 私钥签名（MD5 with RSA） |
| `Rsa_GetKey` | `() (error, string, string)` | 生成 1024 位 RSA 公私钥对 |
| `Rsa_私钥解密` | `([]byte, []byte) string` | RSA 私钥解密（PKCS1v15） |
| `Rsa_私钥解密2` | `([]byte, []byte) []byte` | RSA 私钥解密（返回字节集） |
| `Rsa_公钥加密` | `(string, []byte) string` | RSA 公钥加密（PKCS8 公钥，返回 Base64） |
| `RSA_私钥加密` | `([]byte, []byte) string` | RSA 私钥加密（返回 Base64） |
| `RSA_公钥解密` | `(string, []byte) []byte` | RSA 公钥解密 |

> 注意：公钥加载使用 `x509.ParsePKIXPublicKey`（PKCS8 格式），私钥使用 `x509.ParsePKCS1PrivateKey`（PKCS1 格式）。

---

#### 3.2.14 J加解密（对称加解密算法）

| 文件 | [J加解密.go](utils/J加解密.go) |
|------|------|

> 跨平台模块，使用 Go 标准库 `crypto/aes`、`crypto/des`、`crypto/cipher`、`crypto/rand`，支持 AES（CBC/ECB/GCM/CTR/CFB/OFB）、DES/3DES（CBC/ECB）、RC4、XOR、TEA/XXTEA 等常用对称加解密算法。输出统一为 Base64 编码字符串。

**AES 加解密**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J加解密_AES_CBC加密` | `(明文, 密钥, IV []byte) (string, error)` | AES-CBC 加密，PKCS7 填充，返回 Base64 密文 |
| `J加解密_AES_CBC解密` | `(密文Base64 string, 密钥, IV []byte) ([]byte, error)` | AES-CBC 解密 |
| `J加解密_AES_ECB加密` | `(明文, 密钥 []byte) (string, error)` | AES-ECB 加密（不推荐生产环境） |
| `J加解密_AES_ECB解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | AES-ECB 解密 |
| `J加解密_AES_GCM加密` | `(明文, 密钥, 附加数据 []byte) (string, error)` | AES-GCM 认证加密（推荐），自动生成 nonce |
| `J加解密_AES_GCM解密` | `(密文Base64 string, 密钥, 附加数据 []byte) ([]byte, error)` | AES-GCM 认证解密 |
| `J加解密_AES_CTR加密` | `(明文, 密钥 []byte) (string, error)` | AES-CTR 流加密，自动生成 IV |
| `J加解密_AES_CTR解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | AES-CTR 解密 |
| `J加解密_AES_CFB加密` | `(明文, 密钥 []byte) (string, error)` | AES-CFB 流加密，自动生成 IV |
| `J加解密_AES_CFB解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | AES-CFB 解密 |
| `J加解密_AES_OFB加密` | `(明文, 密钥 []byte) (string, error)` | AES-OFB 流加密，自动生成 IV |
| `J加解密_AES_OFB解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | AES-OFB 解密 |

**DES 加解密**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J加解密_DES_CBC加密` | `(明文, 密钥, IV []byte) (string, error)` | DES-CBC 加密（密钥 8 字节） |
| `J加解密_DES_CBC解密` | `(密文Base64 string, 密钥, IV []byte) ([]byte, error)` | DES-CBC 解密 |
| `J加解密_DES_ECB加密` | `(明文, 密钥 []byte) (string, error)` | DES-ECB 加密 |
| `J加解密_DES_ECB解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | DES-ECB 解密 |

**3DES 加解密**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J加解密_3DES_CBC加密` | `(明文, 密钥, IV []byte) (string, error)` | 3DES-CBC 加密（密钥 24 字节） |
| `J加解密_3DES_CBC解密` | `(密文Base64 string, 密钥, IV []byte) ([]byte, error)` | 3DES-CBC 解密 |
| `J加解密_3DES_ECB加密` | `(明文, 密钥 []byte) (string, error)` | 3DES-ECB 加密 |
| `J加解密_3DES_ECB解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | 3DES-ECB 解密 |

**RC4 / XOR 加解密**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J加解密_RC4` | `(数据, 密钥 []byte) string` | RC4 流加密，返回 Base64 密文 |
| `J加解密_RC4字节集` | `(数据, 密钥 []byte) []byte` | RC4 流加密，返回原始字节集 |
| `J加解密_XOR` | `(数据, 密钥 []byte) []byte` | XOR 异或加密 |
| `J加解密_XOR文本` | `(数据, 密钥 []byte) string` | XOR 加密，返回 Base64 密文 |

**TEA / XXTEA 加解密**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J加解密_TEA加密` | `(明文, 密钥 []byte) (string, error)` | TEA 加密（32 轮 Feistel，密钥 16 字节） |
| `J加解密_TEA解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | TEA 解密 |
| `J加解密_XXTEA加密` | `(明文, 密钥 []byte) (string, error)` | XXTEA 加密（支持任意长度数据） |
| `J加解密_XXTEA解密` | `(密文Base64 string, 密钥 []byte) ([]byte, error)` | XXTEA 解密 |

**密钥/IV 生成工具**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `J加解密_生成AES密钥` | `(长度 int) ([]byte, error)` | 生成随机 AES 密钥（16/24/32 字节） |
| `J加解密_生成IV` | `(块大小 int) ([]byte, error)` | 生成随机初始化向量（AES=16, DES=8） |

**PKCS7 填充辅助**：

| 函数 | 签名 | 说明 |
|------|------|------|
| `pkcs7Padding` | `([]byte, int) []byte` | 加密前 PKCS7 填充 |
| `pkcs7Unpadding` | `([]byte) ([]byte, error)` | 解密后去除 PKCS7 填充 |

**使用示例**：
```go
// AES-CBC 加密示例
key, _ := J加解密_生成AES密钥(32)
iv, _ := J加解密_生成IV(16)
密文, _ := J加解密_AES_CBC加密([]byte("Hello World"), key, iv)
明文, _ := J加解密_AES_CBC解密(密文, key, iv)

// AES-GCM 认证加密（推荐）
密文, _ := J加解密_AES_GCM加密([]byte("sensitive data"), key, nil)
明文, _ := J加解密_AES_GCM解密(密文, key, nil)

// XOR 简单加密
密文 := J加解密_XOR文本([]byte("hello"), []byte("pwd"))

// RC4 加密
密文 := J加解密_RC4([]byte("hello"), []byte("12345"))

// TEA 加密（常用于游戏修改等场景）
密文, _ := J加解密_TEA加密([]byte("data"), []byte("0123456789abcdef"))
```

> 注意：ECB 模式安全性低，不建议生产使用；GCM 模式同时提供机密性和完整性验证，为推荐方案。

---

#### 3.2.15 S数组（数组/切片工具）

| 文件 | [S数组.go](utils/S数组.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `S数组_取随机成员` | `([]string, int) []string` | 随机取 N 个成员（不重复） |
| `S数组_到文本` | `([]interface{}) string` | 数组转逗号分隔文本 |
| `S数组_反转` | `([]interface{})` | 原地反转数组 |
| `S数组_合并文本` | `[T comparable]([]T, string) string` | 泛型数组合并为文本 |
| `S数组_取文本出现次数` | `([]string, string) int` | 统计成员出现次数 |
| `S数组_取文本索引` | `([]string, string) int` | 查找文本索引（失败返回 -1） |
| `S数组_整数是否存在` | `([]int, int) bool` | 整数是否存在 |
| `S数组_是否存在` | `[T comparable]([]T, T) bool` | 泛型判断元素是否存在 |
| `S数组_求平均值` | `([]int) int` | 整数数组求平均 |
| `S数组_是否为空` | `([]string) bool` | 判断数组是否全为空串 |
| `S数组_排序整数` | `([]int) []int` | 整数数组排序 |
| `S数组_排序文本` | `([]string) []string` | 文本数组排序 |
| `S数组_去重复` | `[T comparable]([]T) []T` | 泛型去重 |
| `S数组_乱序` | `[T comparable]([]T) []T` | Fisher-Yates 乱序 |
| `S数组_整数取差集` | `([]int, []int) []int` | 整数数组差集 |
| `S数组_取差集` | `([]int, []int) []int` | 整数数组差集（a 有 b 无） |

---

#### 3.2.15 S时间（时间工具）

| 文件 | [S时间.go](utils/S时间.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `S时间_文本到时间戳` | `(string) int` | `2006-01-02 15:04:05` 格式转 10 位时间戳 |
| `S时间_取现行时间戳13` | `() int64` | 获取 13 位毫秒时间戳 |
| `S时间_取现行时间戳` | `() int64` | 获取 10 位秒时间戳 |
| `S时间_取现行时间` | `() string` | 获取当前时间文本 |
| `S时间_时间戳到时间` | `(int64) string` | 10 位时间戳转时间文本 |
| `S时间_时间戳13到时间` | `(int64) string` | 13 位时间戳转时间文本 |
| `S时间_时间到时间戳` | `(string) int64` | 时间文本转时间戳 |
| `S时间_时间戳格式化` | `(string, int64) string` | 自定义格式化时间戳（支持 y/m/d/H/i/s 占位符） |
| `S时间_秒转时间文本` | `(int64) string` | 秒数转 `X年X月X天X时X分X秒` 文本 |

---

#### 3.2.16 T图片（图片工具）

| 文件 | [T图片.go](utils/T图片.go) |
|------|------|
| 依赖 | `github.com/skip2/go-qrcode` |

| 函数 | 签名 | 说明 |
|------|------|------|
| `T图片_生成二维码base64` | `(string) string` | 生成二维码并返回 Base64 编码 |

---

#### 3.2.17 W文件（文件操作）

| 文件 | [W文件.go](utils/W文件.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文件_是否存在` | `(string) bool` | 判断文件/目录是否存在 |
| `W文件_写到文件` | `(string, []byte) error` | 写入文件（自动创建目录） |
| `W文件_枚举` | `(string, string, *[]string, bool, bool) error` | 枚举目录下指定类型文件 |
| `W文件_取文件名` | `(string) string` | 从路径提取文件名 |
| `W文件_路径合并处理` | `(...string) string` | 路径拼接 |
| `W文件_取父目录` | `(string) string` | 取父目录路径 |
| `W文件_删除` | `(string) error` | 删除文件 |
| `W文件_更名` | `(string, string) error` | 重命名文件/目录 |
| `W文件_写出` | `(string, interface{}) error` | 写出文件（自动创建目录） |
| `W文件_写出文件` | `(string, interface{}) error` | 同 `W文件_写出` |
| `W文件_追加文本` | `(string, string) error` | 追加文本到文件 |
| `W文件_读入文本` | `(string) string` | 读取文件文本内容 |
| `W文件_读入文件` | `(string) []byte` | 读取文件字节集 |
| `W文件_保存` | `(string, interface{}) error` | 智能保存（内容一致则跳过写出） |
| `W文件_取临时文件名` | `(string) (*os.File, string, error)` | 获取临时文件 |

---

#### 3.2.18 W文本（文本处理 — 最大模块）

| 文件 | [W文本.go](utils/W文本.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `W文本_是否包含关键字` | `(string, string) bool` | 是否包含关键字 |
| `W文本_是否存在` | `(string, string) bool` | 是否包含文本 |
| `W文本_是否存在_任意` | `(string, []string) bool` | 是否包含任一关键字 |
| `W文本_是否存在_同时` | `(string, []string) bool` | 是否同时包含所有关键字 |
| `W文本_是否为英数字母` | `(string) bool` | 是否为英数 |
| `W文本_是否为字母` | `(string) bool` | 是否全为字母 |
| `W文本_是否为数字` | `(string) bool` | 是否全为数字 |
| `W文本_倒取出中间文本` | `(string, string, string, int, bool) string` | 从后往前取中间文本 |
| `W文本_取文本所在行` | `(string, string, bool) int` | 查找文本所在行号 |
| `W文本_删除指定文本行` | `(string, int) string` | 删除指定行 |
| `W文本_取随机范围数字` | `(int, int, int) string` | 取随机数字（可选单/双） |
| `W文本_取指定变量文本行` | `(string, int) string` | 取指定行文本 |
| `W文本_颠倒` | `(string, bool) string` | 文本颠倒（支持中文） |
| `W文本_取出现次数` | `(string, string) int` | 统计文本出现次数 |
| `W文本_首字母改大写` | `(string) string` | 首字母大写 |
| `W文本_替换` | `(string, string, string) string` | 全局替换 |
| `W文本_替换2` | `(string, map[string]string) string` | 批量替换 |
| `W文本_寻找` | `(string, string) int` | 查找文本位置 |
| `W文本_取随机IP` | `() string` | 生成随机国内 IP |
| `W文本_取行数` | `(string) int` | 统计文本行数 |
| `W文本_取文本右边2` | `(string, string, int, bool) string` | 高级取右边文本 |
| `W文本_删除空行` | `(string) string` | 删除空行 |
| `W文本_逐字分割` | `(string) []string` | 逐字分割为数组 |
| `W文本_去重复文本` | `(string, string) string` | 去除重复文本 |
| `W文本_取出中间文本_批量正则` | `(string, string, string) []string` | 正则批量取中间文本 |
| `W文本_取出中间文本` | `(string, string, string) string` | 取左右标记之间的文本 |
| `W文本_取文本左边2` | `(string, string) string` | 取关键字左边文本（含关键字） |
| `W文本_取文本左边` | `(string, string) string` | 取关键字左边文本 |
| `W文本_取文本右边` | `(string, string) string` | 取关键字右边文本 |
| `W文本_取随机字符串` | `(int) string` | 生成随机字母数字串 |
| `W文本_取随机字符串_数字` | `(int) string` | 生成随机纯数字串 |
| `W文本_分割文本` | `(string, string) []string` | 分割文本 |
| `W文本_gbk到utf8` | `(string) string` | GBK 转 UTF-8 |
| `W文本_utf8到gbk` | `(string) string` | UTF-8 转 GBK |
| `W文本_取左边` | `(string, int) string` | 取左侧 N 个字符（中文算 1） |
| `W文本_取右边` | `(string, int) string` | 取右侧 N 个字符（中文算 1） |
| `W文本_删首尾空` | `(string) string` | 去除首尾空格 |
| `W文本_是否JSON` | `(string) bool` | 判断是否为 JSON |
| `W文本_删首空` | `(string) string` | 去除左侧空格 |
| `W文本_删尾空` | `(string) string` | 去除右侧空格 |
| `W文本_子文本替换` | `(string, string, string) string` | 子文本替换 |
| `W文本_取随机ip` | `() string` | 生成随机 IP |
| `W文本_到大写` | `(string) string` | 转大写 |
| `W文本_到小写` | `(string) string` | 转小写 |
| `W文本_取长度` | `(string) int` | 取字符数（中文算 1） |
| `W文本_字符` | `(int8) string` | 字节码转字符 |
| `W文本_寻找文本` | `(string, string) int` | 查找文本位置 |
| `W文本_倒找文本` | `(string, string) int` | 从后查找文本位置 |
| `W文本_取空白` | `(int) string` | 生成 N 个空格 |
| `W文本_取重复` | `(int, string) string` | 重复文本 N 次 |
| `W文本_取随机数字数组` | `(int, int, int) []string` | 生成不重复随机数字数组 |
| `W文本_去除敏感信息` | `(string) string` | 中间字符替换为 `*` |
| `W文本_可能为json` | `(string) bool` | 高性能预判 JSON（首尾字符检测） |

---

#### 3.2.19 W网页（HTTP 网页访问）

| 文件 | [W网页.go](utils/W网页.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `W网页_取域名` | `(string) string` | 从 URL 提取域名 |
| `W网页_访问_对象` | `(网址, 方式, ...) []byte` | 完整 HTTP 请求（支持代理、Cookie、重定向控制等） |
| `网页_访问_对象` | `(网址, 方式, ...) []byte` | 同上（无 W 前缀版本） |
| `Q取单条Cookie` | `(string, string) string` | 从 Cookie 字符串提取指定值 |
| `W网页_Cookie合并更新` | `(string, string) string` | 合并新旧 Cookie |
| `W网页_处理协议头` | `(string) string` | 规范化 HTTP 协议头格式 |

> HTTP 访问方式枚举：0=GET, 1=POST, 2=HEAD, 3=PUT, 4=OPTIONS, 5=DELETE, 6=TRACE, 7=CONNECT

---

#### 3.2.20 Y原子（原子操作）

| 文件 | [Y原子.go](utils/Y原子.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `Y原子_递增` | `(*int64) int64` | 原子递增 1，返回新值 |
| `Y原子_递减` | `(*int64) int64` | 原子递减 1，返回新值 |

---

#### 3.2.21 Z字节集（字节集工具）

| 文件 | [Z字节集.go](utils/Z字节集.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `Z字节集_十六进制到字节集` | `(string) []byte` | 16 进制文本转字节集 |
| `Z字节集_字节集到十六进制` | `([]byte) string` | 字节集转 16 进制文本 |
| `Z字节集_寻找` | `([]byte, []byte, int) int` | 字节集中搜索子字节集（1 起始索引） |
| `Z字节集_Gzip解压` | `([]byte) ([]byte, error)` | Gzip 解压 |

---

#### 3.2.22 Z正则（正则校验与提取）

| 文件 | [Z正则.go](utils/Z正则.go) |
|------|------|

| 函数 | 签名 | 说明 |
|------|------|------|
| `Z正则_校验密码` | `(string, *string) bool` | 校验密码（5-17 位非空白） |
| `Z正则_校验代理用户名` | `(string, *string) bool` | 校验代理用户名 |
| `Z正则_校验用户名` | `(string, *string) bool` | 校验用户名（5-17 位） |
| `Z正则_校验email` | `(string, *string) bool` | 校验邮箱格式 |
| `Z正则_校验纯数字` | `(string, *string) bool` | 校验纯数字 |
| `Z正则_校验纯数字指定位数` | `(string, *string, int) bool` | 校验指定位数纯数字 |
| `Z正则_是否英数` | `(string, *string) bool` | 是否为英文+数字 |
| `Z正则_取Url连接地址` | `(string) []string` | 提取所有 URL |
| `Z正则_取全部匹配子文本` | `(string, string) []string` | 正则提取全部匹配 |
| `Z正则_取ip端口` | `(string) string` | 提取首个 IP:端口 |
| `Z正则_取ip端口多个` | `(string) []string` | 提取所有 IP:端口 |

#### 3.2.23 OCV视觉（OpenCV 图像处理，build tag: opencv）

> ⚠️ 本模块需编译标签 `go build -tags opencv`，依赖 `gocv.io/x/gocv` 及本地 OpenCV 库。

| 文件 | [OCV视觉.go](utils/OCV视觉.go) |
|------|------|

本模块提供完整的 OpenCV 封装，包含图像读写、几何变换、颜色空间转换、滤波、形态学、边缘检测、阈值、轮廓、角点检测、直方图、绘图、视频处理、模板匹配（找图）和特征匹配（SIFT/ORB/AKAZE 高级找图）等 79 个函数。

**找图函数速览**：

| 层级 | 函数 | 适用场景 |
|------|------|----------|
| 基础 | `OCV_找图`, `OCV_找图中心` | 尺寸一致、无旋转（最常用） |
| 基础+ | `OCV_找图全部`, `OCV_找图区域`, `OCV_找图掩码` | 多目标/区域/透明图 |
| 高级 | `OCV_找图多尺度` | 目标可能缩放 |
| 高级 | `OCV_找图边缘` | 颜色/光照不同 |
| 特征 | `OCV_找图ORB` | 有旋转/缩放/视角变化（免费快速） |
| 特征 | `OCV_找图SIFT` | 精度最高（有专利） |
| 特征 | `OCV_找图AKAZE` | 模糊/压缩图片 |

> 完整 API 参考见 [API_Reference.md](API_Reference.md) OCV 视觉模块。

---

#### 3.2.24 HHook（API Hook 引擎）

> ⚠️ 本模块为 Windows 专用（编译标签 `amd64` + cgo），依赖 `gcc`（MinGW-w64 / TDM-GCC）。基于 [MinHook](https://github.com/TsudaKageyu/minhook) 库实现。

| 文件 | [HHook.go](utils/HHook.go)、[HHook/HHook.go](utils/HHook/HHook.go)、[HHook/*.c](utils/HHook/) |
|------|------|

本模块封装了 MinHook 库的核心 API，支持对 Windows API 函数的 Inline Hook（运行时劫持目标函数跳转到回调函数）。

**底层架构**：
- `utils/HHook/` — cgo 子包，内含 MinHook v1.3.4 完整 C 源码及 Go 包装层
- `utils/HHook.go` — 面向用户的公开 API（中文命名）

##### 3.2.24.1 基础函数

| 函数 | 签名 | 说明 |
|------|------|------|
| `HHook_初始化` | `() error` | 初始化 Hook 引擎，分配内部资源。必须在使用其他函数前调用 |
| `HHook_卸载` | `() error` | 卸载所有 Hook 并释放资源。应在程序退出前调用 |

##### 3.2.24.2 Hook 管理

| 函数 | 签名 | 说明 |
|------|------|------|
| `HHook_安装Hook` | `(目标地址, 回调地址 uintptr) (uintptr, error)` | 为目标函数安装 Hook（禁用状态创建），返回原始函数的 trampoline 地址 |
| `HHook_安装ApiHook` | `(模块名, 函数名 string, 回调地址 uintptr) (uintptr, error)` | 按模块名+函数名为导出 API 安装 Hook，如 `"kernel32.dll"` + `"MessageBoxW"` |
| `HHook_安装ApiHookEx` | `(模块名, 函数名 string, 回调地址 uintptr) (原地址, 目标地址 uintptr, err error)` | 扩展版安装 API Hook，同时返回原始函数地址和目标函数地址 |
| `HHook_卸载Hook` | `(目标地址 uintptr) error` | 移除指定 Hook 并恢复原始代码 |
| `HHook_启用Hook` | `(目标地址 uintptr) error` | 启用指定 Hook（或传入 `0` 启用全部） |
| `HHook_禁用Hook` | `(目标地址 uintptr) error` | 禁用指定 Hook（或传入 `0` 禁用全部） |
| `HHook_启用全部Hook` | `() error` | 批量启用所有已创建的 Hook |
| `HHook_禁用全部Hook` | `() error` | 批量禁用所有已创建的 Hook |

##### 3.2.24.3 队列操作

用于线程安全的批量切换：先排队，再一次性应用。

| 函数 | 签名 | 说明 |
|------|------|------|
| `HHook_排队启用Hook` | `(目标地址 uintptr) error` | 将指定 Hook 加入启用队列 |
| `HHook_排队禁用Hook` | `(目标地址 uintptr) error` | 将指定 Hook 加入禁用队列 |
| `HHook_应用排队` | `() error` | 一次性应用所有排队的启用/禁用操作 |

##### 3.2.24.4 错误处理

| 函数 | 签名 | 说明 |
|------|------|------|
| `HHook_取状态文本` | `(状态码 int) string` | 将 MH_STATUS 状态码转换为中文描述文本 |

**状态码对应表**：

| 状态码 | 含义 |
|--------|------|
| 0 | 成功 |
| 1 | 已初始化 |
| 2 | 未初始化 |
| 3 | 已创建 |
| 4 | 未创建 |
| 5 | 已启用 |
| 6 | 已禁用 |
| 7 | 不可执行 |
| 8 | 不支持的函数 |
| 9 | 内存分配失败 |
| 10 | 内存保护失败 |
| 11 | 模块未找到 |
| 12 | 函数未找到 |

##### 3.2.24.5 使用示例

```go
import (
    "syscall"
    "unsafe"
    "golang.org/x/sys/windows"
    "github.com/yuan71058/Efunc/utils"
)

// 1. 初始化
if err := utils.HHook_初始化(); err != nil {
    panic(err)
}
defer utils.HHook_卸载()

// 2. 安装 Hook
// 使用 windows.NewCallback 将 Go 函数转为 C 函数指针
callback := windows.NewCallback(myDetour)
original, err := utils.HHook_安装ApiHook("kernel32.dll", "MessageBoxW", callback)
if err != nil {
    panic(err)
}

// 3. 启用 Hook
utils.HHook_启用Hook(0) // 0 = MH_ALL_HOOKS

// 4. 需要调用原始函数时，通过 trampoline 地址回调
// ...

// 5. 卸载 Hook
utils.HHook_卸载Hook(0)
```

> 编译要求：安装 MinGW-w64 或 TDM-GCC 确保 gcc 可用。cgo 会自动编译 `HHook/` 下的 C 源文件并链接 kernel32。

---

## 四、依赖关系图

```
┌─────────────────────────────────────────────────────┐
│                      main.go                        │
│              (入口 / 示例 / 测试)                    │
└──────────────┬──────────────────┬───────────────────┘
               │                  │
       import  │                  │  import
               ▼                  ▼
┌──────────────────────┐  ┌──────────────────────────┐
│      class/          │  │       utils/             │
│  ┌────────────────┐  │  │  ┌────────────────────┐  │
│  │ L_临界许可     │  │  │  │ 核心库 (gconv)     │  │
│  │ L_读写锁       │  │  │  │ 辅助               │  │
│  │ L_正则表达式   │  │  │  │ B编码 (net/url)    │  │
│  │ L_队列         │  │  │  │ C程序 (os/exec)    │  │
│  │ L_队列泛型     │  │  │  │ Float64 (decimal)  │  │
│  └────────────────┘  │  │  │ H汇编              │  │
│                      │  │  │ IP                  │  │
│  依赖:               │  │  │ Int转换             │  │
│  · sync              │  │  │ J校验 (crypto)      │  │
│  · regexp            │  │  │ Post数据类          │  │
│  · container/list    │  │  │ Map                 │  │
│                      │  │  │ M目录 (os)          │  │
│                      │  │  │ Rsa (crypto/rsa)    │  │
│                      │  │  │ S数组               │  │
│                      │  │  │ S时间 (time)        │  │
│                      │  │  │ T图片 (go-qrcode)   │  │
│                      │  │  │ W文件 (os/io)       │  │
│                      │  │  │ W文本 (strings)     │  │
│                      │  │  │ W网页 (net/http)    │  │
│                      │  │  │ Y原子 (sync/atomic) │  │
│                      │  │  │ Z字节集 (hex/gzip)  │  │
│                      │  │  │ Z正则 (regexp)      │  │
│                      │  │  └────────────────────┘  │
│                      │  │                          │
│                      │  │  内部依赖:               │
│                      │  │  · W文本 ← 辅助, H汇编   │
│                      │  │  · W网页 ← W文本, 辅助   │
│                      │  │  · W文件 ← M目录, 核心库  │
│                      │  │  · S数组 ← 核心库        │
│                      │  └──────────────────────────┘
└──────────────────────┘
```

### 外部第三方依赖

| 依赖包 | 版本 | 用途 |
|--------|------|------|
| `github.com/gogf/gf/v2` | v2.5.7 | GoFrame 核心库，主要用于 `gconv` 类型转换 |
| `github.com/shopspring/decimal` | v1.3.1 | 高精度十进制运算，防止浮点精度丢失 |
| `github.com/axgle/mahonia` | v0.0.0-20180208002826 | 字符编码转换（GBK ↔ UTF-8） |
| `github.com/skip2/go-qrcode` | — | 二维码生成 |
| `golang.org/x/text` | — | 文本编码处理（间接依赖） |
| `go.opentelemetry.io/otel` | v1.14.0 | OpenTelemetry（间接依赖，GoFrame 引入） |

---

## 五、命名规范

本项目的核心设计特色是 **中文命名**，遵循以下规则：

| 规则 | 示例 | 说明 |
|------|------|------|
| 首字拼音大写 + 分类名 | `W文本_取出中间文本` | W=文本, S=时间/数组, C=程序, B=编码 |
| 方法名前缀语义化 | `Q取`/`T添加`/`J加入`/`H获取`/`Z置`/`S删除` | Q=取, T=弹/添, J=加, H=获, Z=设, S=删 |
| 类名前缀 `L_` | `L_队列`, `L_读写锁` | L=类 |
| 函数名前缀分类 | `W文件_`, `W文本_`, `W网页_` | 按功能域分类 |

### 分类前缀速查

| 前缀 | 分类 | 文件 |
|------|------|------|
| B | 编码 | B编码.go |
| C | 程序控制 | C程序.go |
| F | Float64 浮点 | Float64转换.go |
| H | 汇编/底层 | H汇编.go |
| I | IP 地址 | IP.go |
| I | Int 整数 | Int转换.go |
| J | 校验摘要 | J校验.go |
| L | Post 数据类 | L类_post数据类.go |
| M | Map/目录 | Map.go, M目录.go |
| R | RSA 加密 | Rsa.go |
| J | 对称加解密 | J加解密.go |
| S | 数组/时间 | S数组.go, S时间.go |
| T | 图片 | T图片.go |
| W | 文件/文本/网页 | W文件.go, W文本.go, W网页.go |
| Y | 原子操作 | Y原子.go |
| Z | 字节集/正则 | Z字节集.go, Z正则.go |

---

## 六、项目运行方式

### 6.1 安装

```bash
go get -u gitee.com/anyueyinluo/Efunc
```

### 6.2 使用示例

```go
package main

import (
    . "github.com/yuan71058/Efunc/utils"
    "fmt"
)

func main() {
    fmt.Println(B编码_URL编码("go语言版的精易模块"))
    fmt.Println(S时间_取现行时间())
    fmt.Println(W文本_取出中间文本("<div>Hello</div>", "<div>", "</div>"))
    fmt.Println(J校验_取md5_文本("hello", false))
}
```

> 使用 `. "github.com/yuan71058/Efunc/utils"` 点导入后，所有工具函数可直接调用，无需包名前缀。

### 6.3 运行测试

```bash
go test ./utils/ -v
```

### 6.4 运行主程序

```bash
go run main.go
```

---

## 七、注意事项

1. **编码格式**：项目大量使用中文命名，确保源文件以 UTF-8 编码保存
2. **Go 版本**：需要 Go 1.22+（`go.mod` 声明），泛型功能需要 Go 1.18+
3. **RSA 密钥格式**：公钥使用 PKCS8 格式（`ParsePKIXPublicKey`），私钥使用 PKCS1 格式（`ParsePKCS1PrivateKey`），密钥格式不匹配会导致加载失败
4. **浮点精度**：所有 `Float64` 系列函数使用 `decimal` 库，避免精度丢失，适用于金融计算场景
5. **线程安全**：`class/` 下的队列、锁等结构均为线程安全设计，可在并发环境直接使用
6. **已弃用函数**：`C程序_延时2` 为逐毫秒循环实现，建议使用 `C程序_延时`（基于 `time.Sleep`）
7. **待修复项**：`C程序_删除自身` 在 Windows 下可能因文件占用而失败；`C程序_是否被调试` 实现简单，仅检查 PID
