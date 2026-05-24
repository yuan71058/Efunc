<div align="center">

# ⚡ Efunc

**Go 语言版的精易模块**

中文命名的工具函数库，涵盖编码转换、加密校验、文本处理、文件操作、网络请求、并发安全数据结构、数据库操作、权限管理等常用功能

[![Version](https://img.shields.io/badge/Version-v2.1.0-blue?style=flat-square)](https://github.com/yuan71058/Efunc/releases)
[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.18-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Functions](https://img.shields.io/badge/Functions-650%2B-orange?style=flat-square)](API_Reference.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/yuan71058/Efunc.svg)](https://pkg.go.dev/github.com/yuan71058/Efunc)

[快速开始](#-快速开始) · [模块一览](#-模块一览) · [命名规则](#-命名规则) · [更新日志](#-更新日志)

</div>

---

## ✨ 特性

- 🇨🇳 **中文命名** — 函数名直观易读，降低学习门槛
- 📦 **开箱即用** — `go get` 一键安装，点导入直接调用
- 🔒 **并发安全** — 内置互斥锁、读写锁、线程安全队列等数据结构
- 🔐 **加密校验** — MD5 / SHA / CRC32 / RSA 签名验签一应俱全
- 🌐 **网络请求** — HTTP 客户端、Cookie 管理、代理支持、网页爬虫
- 📝 **文本处理** — 查找、截取、替换、编码转换、正则校验、模板引擎
- 📂 **文件操作** — 读写、枚举、目录管理、路径处理、文件监控
- ⏱ **时间工具** — 时间戳转换、格式化、时间计算、智能日期解析
- 🗄 **数据库** — MySQL / SQLite ORM 操作、键值数据库
- 🛡 **权限管理** — RBAC / ABAC 权限控制
- 🧮 **表达式计算** — 数学/逻辑表达式求值
- ✅ **数据校验** — 结构体字段验证

## 📥 安装

```bash
go get -u github.com/yuan71058/Efunc@v2.1.0
```

## 🚀 快速开始

```go
package main

import (
	"fmt"

	. "github.com/yuan71058/Efunc/utils"
)

func main() {
	// 编码转换
	fmt.Println(B编码_URL编码("go语言版的精易模块"))

	// 文本截取
	fmt.Println(W文本_取出中间文本("<div>Hello</div>", "<div>", "</div>"))

	// 时间获取
	fmt.Println(S时间_取现行时间())

	// 哈希校验
	fmt.Println(J校验_取md5_文本("hello", false))

	// 表达式计算
	result, _ := B表达式_计算带参数("a + b * c", map[string]interface{}{"a": 1, "b": 2, "c": 3})
	fmt.Println(result) // 7

	// 数据校验
	err := V校验_验证结构体(user)
	fmt.Println(err)
}
```

## 📦 导入方式

```go
import (
	"github.com/yuan71058/Efunc/class"       // 类模块（并发安全数据结构）
	. "github.com/yuan71058/Efunc/utils"      // 工具函数（点导入可直接调用）
)
```

## 🧩 模块一览

### 🏗 class — 并发安全数据结构

| 类 | 说明 |
|:---|:-----|
| `L_临界许可` | 互斥锁 |
| `L_读写锁` | 读写锁 |
| `L_正则表达式` | 正则匹配封装 |
| `L_队列` | 线程安全队列 |
| `L_队列泛型` | 泛型线程安全队列 |

### 🔧 utils — 工具函数（650+）

#### 基础模块

| 模块 | 文件 | 功能 |
|:----:|:-----|:-----|
| 🏠 | `核心库` | 类型转换、三元运算、格式化 |
| 🛠 | `辅助` | 简易文本/数组操作 |
| 🔤 | `B编码` | URL / Base64 / Base32 / Hex / HTML / JSON / QP / Punycode 编解码 |
| ⚙️ | `C程序` | 延时、GUID、日志、命令执行 |
| 🔢 | `Float64转换` | 高精度浮点运算 |
| 🎲 | `H汇编` | 随机数生成 |
| 🌐 | `IP` | IP 地址转换 |
| 🔢 | `Int转换` | 整数转换 |
| 🔐 | `J校验` | MD5 / CRC32 / CRC64 / SHA / HMAC / 文件哈希 / 校验比对 |
| 📮 | `L类_post数据类` | POST 数据构造 |
| 🗺 | `Map` | Map 操作、结构体互转 |
| 📂 | `M目录` | 目录创建、枚举、删除 |
| 🔑 | `Rsa` | RSA 加解密 / 签名验签 |
| 📊 | `S数组` | 数组去重、排序、差集 |
| ⏱ | `S时间` | 时间戳转换、格式化、计算 |
| 🖼 | `T图片` | 图片处理（读写/缩放/裁剪/旋转/效果/水印/二维码） |
| 📄 | `W文件` | 文件读写、枚举、路径处理 |
| ✏️ | `W文本` | 文本查找、截取、替换（最大模块） |
| 🌍 | `W网页` | HTTP 请求、Cookie 管理 |
| ⚛️ | `Y原子` | 原子递增/递减计数器 |
| 📦 | `Z字节集` | 字节集与十六进制互转、Gzip 解压 |
| 🔍 | `Z正则` | 密码校验、邮箱验证、IP 提取 |

#### 扩展模块（v2.0.0 新增）

| 模块 | 文件 | 基于库 | 功能 |
|:----:|:-----|:------|:-----|
| 🖥 | `M命令行` | `flag` | 命令行参数解析（字符串/整数/布尔/小数） |
| 📅 | `R日期解析` | `dateparse` | 智能日期时间格式自动识别 |
| ♻️ | `P对象池` | `bytebufferpool` | 字节缓冲区对象池，减少 GC 压力 |
| 🧮 | `B表达式计算` | `govaluate` | 数学/逻辑表达式求值 |
| 📝 | `T模板` | `fasttemplate` | 高性能模板替换引擎 |
| ✅ | `V数据校验` | `validator/v10` | 结构体字段校验 |
| 🔀 | `J结构体合并` | `mergo` + `copier` | 结构体合并与深拷贝 |
| 📊 | `K表格` | `termtables` | 控制台表格渲染 + 多格式输出 + 数据操作 |
| 👁 | `F文件监控` | `fsnotify` | 文件系统变化监控 |
| 📡 | `X消息总线` | `message-bus` | 发布-订阅消息总线 |
| 🌐 | `H客户端` | `resty/v2` | HTTP 客户端（GET/POST/PUT/DELETE） |
| 🗃 | `N键值库` | `buntdb` | 嵌入式键值数据库 |
| 🕷 | `C爬虫` | `colly/v2` | 网页爬虫框架 |
| 🛡 | `Q权限管理` | `casbin/v2` | RBAC/ABAC 权限管理 |
| 🗄 | `D数据库` | `xorm` | ORM 数据库操作（MySQL/SQLite） |
| 🪟 | `C窗口` | `user32.dll` | Windows 窗口操作（查找/枚举/消息/移动） |
| ⚙️ | `C进程` | `kernel32.dll` | Windows 进程管理（创建/枚举/终止/优先级） |
| 👁 | `OCV视觉` | `gocv` | OpenCV 计算机视觉（滤波/边缘/轮廓/特征/视频） |

#### 企业级模块（v2.0.0 新增）

| 模块 | 文件 | 基于库 | 功能 |
|:----:|:-----|:------|:-----|
| 📋 | `P配置` | `viper` | 配置文件管理（JSON/YAML/TOML/ENV） |
| 📊 | `L日志` | `zap` | 高性能结构化日志 |
| 🔄 | `G协程池` | `ants` | 协程池管理 |
| ⏰ | `D定时` | `cron/v3` | 定时任务调度 |
| 📧 | `E邮件` | `gomail/v2` | 邮件发送 |
| 💻 | `X系统信息` | `gopsutil` | 系统信息获取（CPU/内存/磁盘） |
| 🔧 | `K环境变量` | `godotenv` | 环境变量管理 |
| 🔄 | `C类型转换` | `cast` | 类型安全转换 |
| 📦 | `Jjson` | `gjson` / `sjson` | JSON 高效读写 |

## 📐 命名规则

所有函数采用 **拼音首字母大写 + 分类名 + 下划线 + 功能名** 命名，直观易读：

| 前缀 | 分类 | 示例 |
|:----:|:----:|:-----|
| `B` | 编码/表达式 | `B编码_URL编码` / `B表达式_计算` |
| `C` | 程序/爬虫/类型转换/窗口/进程 | `C程序_延时` / `C爬虫_访问` / `C窗口_查找` / `C进程_终止` |
| `O` | OpenCV 视觉 | `OCV_读取图片` / `OCV_Canny边缘检测` |
| `D` | 数据库/定时 | `D数据库_连接MySQL` / `D定时_添加任务` |
| `E` | 邮件 | `E邮件_发送` |
| `F` | 文件监控 | `F文件监控_监控目录变化` |
| `G` | 协程池 | `G协程池_提交任务` |
| `H` | 汇编/客户端 | `H客户端_Get` |
| `J` | 校验/JSON/结构体 | `J校验_取md5` / `Jjson_取值` / `J结构体_合并` |
| `K` | 表格/环境变量 | `K表格_快速创建` / `K环境变量_加载` |
| `L` | 日志/POST数据 | `L日志_信息` |
| `M` | 目录/命令行 | `M目录_创建` / `M命令行_取字符串参数` |
| `N` | 键值库 | `N键值_置值` |
| `P` | 配置/对象池 | `P配置_读取` / `P对象池_获取` |
| `Q` | 权限管理 | `Q权限_检查权限` |
| `R` | RSA/日期解析 | `Rsa_加密` / `R日期_智能解析` |
| `S` | 数组/时间 | `S数组_去重复` / `S时间_取现行时间` |
| `T` | 图片/模板 | `T模板_执行` |
| `V` | 数据校验 | `V校验_验证结构体` |
| `W` | 文件/文本/网页 | `W文本_取出中间文本` |
| `X` | 系统信息/消息总线 | `X系统信息_取CPU使用率` / `X消息_发布` |
| `Y` | 原子 | `Y原子_递增` |
| `Z` | 字节集/正则 | `Z字节集_寻找` |

## 📚 依赖

### 核心依赖

| 库 | 用途 |
|:---|:-----|
| [GoFrame](https://github.com/gogf/gf) | 工具库 |
| [decimal](https://github.com/shopspring/decimal) | 高精度十进制运算 |
| [mahonia](https://github.com/axgle/mahonia) | 字符编码转换 |
| [go-qrcode](https://github.com/skip2/go-qrcode) | 二维码生成 |

### 扩展依赖（v2.0.0 新增）

| 库 | 用途 |
|:---|:-----|
| [viper](https://github.com/spf13/viper) | 配置文件管理 |
| [zap](https://go.uber.org/zap) | 高性能结构化日志 |
| [ants](https://github.com/panjf2000/ants) | 协程池 |
| [cron](https://github.com/robfig/cron/v3) | 定时任务调度 |
| [gomail](https://gopkg.in/gomail.v2) | 邮件发送 |
| [gopsutil](https://github.com/shirou/gopsutil) | 系统信息获取 |
| [godotenv](https://github.com/joho/godotenv) | 环境变量管理 |
| [cast](https://github.com/spf13/cast) | 类型安全转换 |
| [gjson](https://github.com/tidwall/gjson) | JSON 高效读取 |
| [sjson](https://github.com/tidwall/sjson) | JSON 高效写入 |
| [flag](https://pkg.go.dev/flag) | 命令行参数解析 |
| [dateparse](https://github.com/araddon/dateparse) | 智能日期解析 |
| [bytebufferpool](https://github.com/valyala/bytebufferpool) | 字节缓冲区对象池 |
| [govaluate](https://github.com/Knetic/govaluate) | 表达式计算 |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 高性能模板引擎 |
| [validator](https://github.com/go-playground/validator/v10) | 数据校验 |
| [mergo](https://dario.cat/mergo) | 结构体合并 |
| [copier](https://github.com/jinzhu/copier) | 结构体深拷贝 |
| [termtables](https://github.com/scylladb/termtables) | 控制台表格 |
| [fsnotify](https://github.com/fsnotify/fsnotify) | 文件系统监控 |
| [message-bus](https://github.com/vardius/message-bus) | 消息总线 |
| [resty](https://github.com/go-resty/resty/v2) | HTTP 客户端 |
| [buntdb](https://github.com/tidwall/buntdb) | 嵌入式键值数据库 |
| [colly](https://github.com/gocolly/colly/v2) | 网页爬虫 |
| [casbin](https://github.com/casbin/casbin/v2) | 权限管理 |
| [xorm](https://xorm.io/xorm) | ORM 数据库操作 |
| [imaging](https://github.com/disintegration/imaging) | 图片处理（缩放/裁剪/旋转/效果） |
| [gocv](https://gocv.io/) | OpenCV 4.x 计算机视觉 |

## 📖 文档

| 文档 | 说明 |
|:-----|:-----|
| [API 参考文档](API_Reference.md) | 全部 500+ 函数的详细说明 |
| [Code Wiki](Code_Wiki.md) | 项目架构与模块详解 |

## 📋 更新日志

### v2.1.0 (2026-05-24)

**🌐 IP 模块全面增强**
- 新增 `IP_取内网IP` — 获取本机所有内网 IPv4 地址
- 新增 `IP_取首选内网IP` — 获取首选局域网 IP（跳过链路本地地址）
- 新增 `IP_取外网IP` — 通过公共 API 获取公网 IPv4 地址
- 新增 `IP_取外网IP详细信息` — 获取外网 IP 地理位置/ISP JSON 信息
- 新增 `IP_IP转10进制` — 点分十进制 IP 转整数
- 新增 `IP_是否内网IP` — 判断是否为私有地址
- 新增 `IP_是否有效IP` — 验证 IPv4/IPv6 格式
- 新增 `IP_取MAC地址` — 获取本机全部网络接口 MAC 地址
- 新增 `IP_Ping测试` — TCP 连通性测试

**⏱ 网络时间功能**
- 新增 `S时间_取网络时间` — 通过 NTP 协议获取网络标准时间
- 新增 `S时间_取网络时间戳` — 通过 NTP 获取 10 位时间戳
- 新增 `S时间_取网络时间文本` — 通过 NTP 获取格式化时间文本
- 新增 `S时间_取HTTP网络时间` — HTTP 方式获取网络时间（NTP 备选）

**💻 系统信息模块扩展**
- 新增 `X系统_取CPU物理核心数` / `X系统_取总CPU使用率`
- 新增 `X系统_取磁盘IO信息` / `X系统_取网络IO信息`
- 新增 `X系统_取系统负载` — Load1/Load5/Load15
- 新增 `X系统_取进程列表/信息/CPU占用/内存占用/进程名`
- 新增 `X系统_取当前进程ID/信息`
- 新增 `X系统_是否64位系统` / `X系统_取系统架构` / `X系统_取操作系统类型` / `X系统_取逻辑处理器数` / `X系统_取Go版本`

**📝 其他改进**
- 辅助函数模块重命名为 `F_helper.go`，函数名统一加 `F_` 前缀以符合 Go 导出规则
- 函数总数增长至 650+

### v2.0.0 (2026-05-24)

**🚀 新增 24 个扩展模块**

企业级模块：
- `P配置` — 配置文件管理（JSON/YAML/TOML/ENV）
- `L日志` — 高性能结构化日志
- `G协程池` — 协程池管理
- `D定时` — 定时任务调度
- `E邮件` — 邮件发送
- `X系统信息` — 系统信息获取
- `K环境变量` — 环境变量管理
- `C类型转换` — 类型安全转换
- `Jjson` — JSON 高效读写

扩展工具模块：
- `M命令行` — 命令行参数解析
- `R日期解析` — 智能日期时间解析
- `P对象池` — 字节缓冲区对象池
- `B表达式计算` — 数学/逻辑表达式计算
- `T模板` — 高性能模板替换
- `V数据校验` — 结构体数据校验
- `J结构体合并` — 结构体合并与深拷贝
- `K表格` — 控制台表格渲染 + 多格式输出（Markdown/CSV/TSV/JSON/HTML）+ 数据操作
- `F文件监控` — 文件系统监控
- `X消息总线` — 发布-订阅消息总线
- `H客户端` — HTTP 客户端
- `N键值库` — 嵌入式键值数据库
- `C爬虫` — 网页爬虫框架
- `Q权限管理` — RBAC/ABAC 权限管理
- `D数据库` — ORM 数据库操作
- `C窗口` — Windows 窗口操作（查找/枚举/消息/移动/关闭）
- `C进程` — Windows 进程管理（创建/枚举/终止/优先级）

**📝 其他改进**
- 全部代码添加详细中文注释
- 修复已知编译错误
- 函数总数从 225+ 增长至 500+

### v1.0.0

- 初始版本，包含 225+ 工具函数
- 中文命名风格
- 并发安全数据结构

## 🙏 致谢

- 原库地址：[https://gitee.com/anyueyinluo/Efunc](https://gitee.com/anyueyinluo/Efunc)

## 📄 License

[MIT](LICENSE)
