<div align="center">

# ⚡ Efunc

**Go 语言版的精易模块**

中文命名的工具函数库，涵盖编码转换、加密校验、文本处理、文件操作、网络请求、并发安全数据结构等常用功能

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.18-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Functions](https://img.shields.io/badge/Functions-225%2B-orange?style=flat-square)](API_Reference.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/yuan71058/Efunc.svg)](https://pkg.go.dev/github.com/yuan71058/Efunc)

[快速开始](#-快速开始) · [模块一览](#-模块一览) · [命名规则](#-命名规则) · [API 文档](#-文档)

</div>

---

## ✨ 特性

- 🇨🇳 **中文命名** — 函数名直观易读，降低学习门槛
- 📦 **开箱即用** — `go get` 一键安装，点导入直接调用
- 🔒 **并发安全** — 内置互斥锁、读写锁、线程安全队列等数据结构
- 🔐 **加密校验** — MD5 / SHA / CRC32 / RSA 签名验签一应俱全
- 🌐 **网络请求** — HTTP 请求、Cookie 管理、代理支持
- 📝 **文本处理** — 查找、截取、替换、编码转换、正则校验
- 📂 **文件操作** — 读写、枚举、目录管理、路径处理
- ⏱ **时间工具** — 时间戳转换、格式化、时间计算

## 📥 安装

```bash
go get -u github.com/yuan71058/Efunc
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

### 🔧 utils — 工具函数（225+）

| 模块 | 文件 | 功能 |
|:----:|:-----|:-----|
| 🏠 | `核心库` | 类型转换、三元运算、格式化 |
| 🛠 | `辅助` | 简易文本/数组操作 |
| 🔤 | `B编码` | URL / Base64 / USC2 编解码 |
| ⚙️ | `C程序` | 延时、GUID、日志、命令执行 |
| 🔢 | `Float64转换` | 高精度浮点运算 |
| 🎲 | `H汇编` | 随机数生成 |
| 🌐 | `IP` | IP 地址转换 |
| 🔢 | `Int转换` | 整数转换 |
| 🔐 | `J校验` | MD5 / CRC32 / SHA 校验 |
| 📮 | `L类_post数据类` | POST 数据构造 |
| 🗺 | `Map` | Map 操作、结构体互转 |
| 📂 | `M目录` | 目录创建、枚举、删除 |
| 🔑 | `Rsa` | RSA 加解密 / 签名验签 |
| 📊 | `S数组` | 数组去重、排序、差集 |
| ⏱ | `S时间` | 时间戳转换、格式化、计算 |
| 🖼 | `T图片` | 二维码生成 |
| 📄 | `W文件` | 文件读写、枚举、路径处理 |
| ✏️ | `W文本` | 文本查找、截取、替换（最大模块） |
| 🌍 | `W网页` | HTTP 请求、Cookie 管理 |
| ⚛️ | `Y原子` | 原子递增/递减计数器 |
| 📦 | `Z字节集` | 字节集与十六进制互转、Gzip 解压 |
| 🔍 | `Z正则` | 密码校验、邮箱验证、IP 提取 |

## 📐 命名规则

所有函数采用 **拼音首字母大写 + 分类名 + 下划线 + 功能名** 命名，直观易读：

| 前缀 | 分类 | 示例 |
|:----:|:----:|:-----|
| `B` | 编码 | `B编码_URL编码` |
| `C` | 程序 | `C程序_延时` |
| `D` | 转换 | `D到文本` |
| `J` | 校验 | `J校验_取md5` |
| `M` | 目录 | `M目录_创建` |
| `S` | 数组/时间 | `S数组_去重复` / `S时间_取现行时间` |
| `W` | 文件/文本/网页 | `W文本_取出中间文本` |
| `Z` | 字节集/正则 | `Z字节集_寻找` |

## 📚 依赖

| 库 | 用途 |
|:---|:-----|
| [GoFrame](https://github.com/gogf/gf) | 工具库 |
| [decimal](https://github.com/shopspring/decimal) | 高精度十进制运算 |
| [mahonia](https://github.com/axgle/mahonia) | 字符编码转换 |
| [go-qrcode](https://github.com/skip2/go-qrcode) | 二维码生成 |

## 📖 文档

| 文档 | 说明 |
|:-----|:-----|
| [API 参考文档](API_Reference.md) | 全部 225+ 函数的详细说明 |
| [Code Wiki](Code_Wiki.md) | 项目架构与模块详解 |

## 🙏 致谢

- 原库地址：[https://gitee.com/anyueyinluo/Efunc](https://gitee.com/anyueyinluo/Efunc)

## 📄 License

[MIT](LICENSE)
