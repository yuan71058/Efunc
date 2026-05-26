# Efunc DLL 导出文档

> 版本：1.0.0 | 更新时间：2026-05-26

## 一、概述

Efunc DLL 将 Go 工具库的全部函数通过 Windows DLL 形式导出，供 C/C++、Python、C#、易语言、Delphi 等任何支持调用原生 DLL 的编程语言使用。

**核心设计**：采用 **JSON 通用调用模式**，通过单一入口 `Efunc_Call` 传递函数名和 JSON 参数，返回 JSON 结果。无需为每个函数编写 C 导出声明，且能覆盖 `interface{}`、`[]byte`、`map` 等复杂 Go 类型。

## 二、DLL 导出函数

DLL 共导出 4 个 C 函数：

| 函数 | 签名 | 说明 |
|------|------|------|
| `Efunc_Call` | `char* Efunc_Call(char* name, char* params)` | 通用调用入口 |
| `Efunc_Free` | `void Efunc_Free(void* ptr)` | 释放 DLL 分配的内存 |
| `Efunc_List` | `char* Efunc_List()` | 列出所有可用函数（JSON） |
| `Efunc_Version` | `char* Efunc_Version()` | 返回版本号字符串 |

### 2.1 Efunc_Call

**参数**：
- `name`：函数名（与 Go 源码中的函数名一致，支持中文）
- `params`：JSON 格式的参数对象，键名与 Go 函数参数名一致

**返回值**：JSON 字符串，格式如下：

```json
// 成功
{"ok": true, "data": <结果值>}

// 失败
{"ok": false, "error": "错误信息"}
```

**内存管理**：返回的 `char*` 由 DLL 内部通过 `C.CString` 分配，调用方**必须**在使用完毕后调用 `Efunc_Free` 释放，否则会造成内存泄漏。

### 2.2 Efunc_Free

释放 `Efunc_Call` 或 `Efunc_List` 返回的字符串指针。每次调用 `Efunc_Call` 后都必须配对调用。

### 2.3 Efunc_List

返回所有已注册函数的信息列表（JSON 数组），格式：

```json
[
  {"name": "W文本_取长度", "params": ["value"], "desc": "获取文本字符数"},
  {"name": "J校验_取md5_文本", "params": ["文本数据", "返回值转成大写"], "desc": "MD5哈希(文本输入)"}
]
```

### 2.4 Efunc_Version

返回版本号字符串，如 `"1.0.0"`。

## 三、类型映射规则

Go 与 C/JSON 之间的类型映射：

| Go 类型 | JSON 传递方式 | 示例 | 说明 |
|---------|-------------|------|------|
| `string` | JSON 字符串 | `"hello"` | 直接传递 |
| `int` | JSON 整数 | `42` | 直接传递 |
| `int64` | JSON 整数 | `1700000000` | 直接传递 |
| `int32` | JSON 整数 | `1234` | 直接传递 |
| `uint32` | JSON 整数 | `5000` | 直接传递 |
| `float64` | JSON 浮点数 | `3.14` | 直接传递 |
| `bool` | JSON 布尔 | `true` | 直接传递 |
| `int8` | JSON 整数 | `65` | 范围 -128~127 |
| `[]byte` | **Base64 字符串** | `"SGVsbG8="` | 编码后传递 |
| `[]string` | JSON 字符串数组 | `["a","b","c"]` | 直接传递 |
| `[]int` | JSON 整数数组 | `[1,2,3]` | 直接传递 |
| `interface{}` | 任意 JSON 值 | `"text"`, `123`, `{...}` | 自动推断类型 |
| `map[string]string` | JSON 对象 | `{"k":"v"}` | 直接传递 |
| `error` | 响应的 `error` 字段 | — | 出错时填入 `error` 字段 |

### 3.1 []byte 的 Base64 传递规则

Go 的 `[]byte` 类型无法直接映射到 C 类型，因此统一采用 **Base64 编码** 传递：

- **输入方向**（调用方 → DLL）：将原始字节集进行 Base64 编码后作为 JSON 字符串传入
- **输出方向**（DLL → 调用方）：DLL 将 `[]byte` 返回值进行 Base64 编码后放入 `data` 字段

示例：要传递 `[]byte{0x48, 0x65, 0x6C, 0x6C, 0x6F}`（即 "Hello"），编码为 `"SGVsbG8="` 传入。

### 3.2 无参数函数

对于无参数的函数（如 `S时间_取现行时间戳`），`params` 传空对象 `{}` 即可。

## 四、完整函数列表

### 4.1 文本处理（W文本_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `W文本_取长度` | `value: string` | `int` | 获取文本字符数（非字节数） |
| `W文本_是否包含关键字` | `内容: string, 关键字: string` | `bool` | 检查是否包含关键字 |
| `W文本_是否存在` | `内容: string, 关键字: string` | `bool` | 检查是否包含关键字（别名） |
| `W文本_是否存在_任意` | `内容: string, 关键字: []string` | `bool` | 检查是否包含任意一个关键字 |
| `W文本_是否存在_同时` | `内容: string, 关键字: []string` | `bool` | 检查是否同时包含所有关键字 |
| `W文本_取出中间文本` | `内容: string, 左边文本: string, 右边文本: string` | `string` | 提取左右标记之间的文本 |
| `W文本_取出中间文本_批量正则` | `内容: string, 左边文本: string, 右边文本: string` | `[]string` | 正则批量提取中间文本 |
| `W文本_取左边` | `欲取其部分的文本: string, 欲取出字符的数目: int` | `string` | 从左侧截取指定字符数 |
| `W文本_取右边` | `欲取其部分的文本: string, 欲取出字符的数目: int` | `string` | 从右侧截取指定字符数 |
| `W文本_取文本左边` | `内容: string, 关键字: string` | `string` | 获取关键字左侧文本（不含关键字） |
| `W文本_取文本右边` | `内容: string, 关键字: string` | `string` | 获取关键字右侧文本（不含关键字） |
| `W文本_替换` | `源文本: string, 旧文本: string, 新文本: string` | `string` | 全部替换 |
| `W文本_子文本替换` | `欲被替换的文本: string, 欲被替换的子文本: string, 用作替换的子文本: string` | `string` | 全部替换（别名） |
| `W文本_删首尾空` | `内容: string` | `string` | 去除首尾空白 |
| `W文本_删首空` | `欲删除空格的文本: string` | `string` | 去除左侧空格 |
| `W文本_删尾空` | `欲删除空格的文本: string` | `string` | 去除右侧空格 |
| `W文本_到大写` | `value: string` | `string` | 转换为大写 |
| `W文本_到小写` | `value: string` | `string` | 转换为小写 |
| `W文本_首字母改大写` | `英文文本: string` | `string` | 首字母大写 |
| `W文本_分割文本` | `待分割文本: string, 用作分割的文本: string` | `[]string` | 按分隔符分割 |
| `W文本_逐字分割` | `原文本: string` | `[]string` | 逐字符拆分 |
| `W文本_寻找` | `源文本: string, 要寻找的文本: string` | `int` | 查找文本位置（字节偏移，-1 未找到） |
| `W文本_寻找文本` | `被搜寻的文本: string, 欲寻找的文本: string` | `int` | 查找文本位置（别名） |
| `W文本_倒找文本` | `被搜寻的文本: string, 欲寻找的文本: string` | `int` | 从后查找文本位置 |
| `W文本_取出现次数` | `被搜索文本: string, 欲搜索文本: string` | `int` | 统计出现次数 |
| `W文本_取重复` | `重复次数: int, 待重复文本: string` | `string` | 重复文本 |
| `W文本_取空白` | `重复次数: int` | `string` | 生成空格字符串 |
| `W文本_取行数` | `文本: string` | `int` | 统计行数 |
| `W文本_删除空行` | `要操作的文本: string` | `string` | 删除空行 |
| `W文本_去重复文本` | `原文本: string, 分割符: string` | `string` | 去重 |
| `W文本_是否JSON` | `s: string` | `bool` | 检查是否为 JSON |
| `W文本_是否为英数字母` | `s: string` | `bool` | 检查是否仅含英数字母 |
| `W文本_是否为数字` | `s: string` | `bool` | 检查是否仅含数字 |
| `W文本_是否为字母` | `s: string` | `bool` | 检查是否仅含字母 |
| `W文本_取随机字符串` | `字符串长度: int` | `string` | 生成随机字符串 |
| `W文本_取随机字符串_数字` | `字符串长度: int` | `string` | 生成随机数字字符串 |
| `W文本_去除敏感信息` | `内容: string` | `string` | 文本脱敏 |
| `W文本_颠倒` | `欲转换文本: string, 带有中文: bool` | `string` | 文本反转 |
| `W文本_取指定变量文本行` | `文本: string, 行号: int` | `string` | 获取指定行内容（行号从1开始） |
| `W文本_删除指定文本行` | `源文本: string, 行数: int` | `string` | 删除指定行 |
| `W文本_取文本所在行` | `源文本: string, 欲查找的文本: string, 是否区分大小写: bool` | `int` | 查找文本所在行号 |
| `W文本_字符` | `字节型: int8` | `string` | 字节值转字符 |
| `W文本_gbk到utf8` | `src: string` | `string` | GBK 转 UTF-8 |
| `W文本_utf8到gbk` | `src: string` | `string` | UTF-8 转 GBK |

### 4.2 文件操作（W文件_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `W文件_是否存在` | `路径: string` | `bool` | 判断文件或目录是否存在 |
| `W文件_写到文件` | `文件名: string, 欲写入文件的数据: string(Base64)` | `bool` | 写入字节数据到文件 |
| `W文件_读入文本` | `文件名: string` | `string` | 读取文件全部内容为文本 |
| `W文件_读入文件` | `文件名: string` | `string(Base64)` | 读取文件全部内容为字节集 |
| `W文件_删除` | `欲删除的文件名: string` | `bool` | 删除文件 |
| `W文件_更名` | `欲更名的原文件或目录名: string, 欲更改为的现文件或目录名: string` | `bool` | 重命名 |
| `W文件_取文件名` | `路径: string` | `string` | 从路径提取文件名 |
| `W文件_取父目录` | `dirpath: string` | `string` | 获取父目录 |
| `W文件_取大小` | `文件名: string` | `int64` | 获取文件大小（字节） |
| `W文件_追加文本` | `文件名: string, 欲追加的文本: string` | `bool` | 追加文本到文件末尾 |
| `W文件_写出` | `文件名: string, 欲写入文件的数据: any` | `bool` | 写出数据（支持任意类型） |
| `W文件_保存` | `文件名: string, 欲写入文件的数据: any` | `bool` | 智能保存（内容不同才写入） |

### 4.3 时间处理（S时间_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `S时间_取现行时间戳` | 无 | `int64` | 获取当前 Unix 时间戳（秒） |
| `S时间_取现行时间` | 无 | `string` | 获取当前时间字符串 |
| `S时间_时间戳格式化` | `format: string, 时间戳: int64` | `string` | 时间戳格式化 |
| `S时间_是否闰年` | `年: int` | `bool` | 判断是否闰年 |
| `S时间_取月份天数` | `年: int, 月: int` | `int` | 获取月份天数 |

### 4.4 编码转换（B编码_ / Z字节集_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `B编码_URL编码` | `欲编码的文本: string` | `string` | URL 编码 |
| `B编码_URL解码` | `欲解码的文本: string` | `string` | URL 解码 |
| `B编码_BASE64编码` | `字节集: string(Base64)` | `string` | Base64 编码 |
| `B编码_BASE64解码` | `文本: string` | `string(Base64)` | Base64 解码 |
| `B编码_十六进制编码` | `字节集: string(Base64)` | `string` | 十六进制编码 |
| `B编码_十六进制解码` | `文本: string` | `string(Base64)` | 十六进制解码 |
| `B编码_UTF8到GBK` | `文本: string` | `string(Base64)` | UTF-8 转 GBK |
| `B编码_GBK到UTF8` | `数据: string(Base64)` | `string` | GBK 转 UTF-8 |
| `B编码_JSON编码` | `值: any` | `string` | JSON 编码 |
| `Z字节集_十六进制到字节集` | `原始16进制文本: string` | `string(Base64)` | 十六进制文本转字节集 |
| `Z字节集_字节集到十六进制` | `字节集: string(Base64)` | `string` | 字节集转十六进制 |
| `Z字节集_Gzip解压` | `字节集: string(Base64)` | `string(Base64)` | Gzip 解压 |

### 4.5 加解密（J加解密_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `J加解密_AES_CBC加密` | `明文: string(Base64), 密钥: string(Base64), IV: string(Base64)` | `string` | AES-CBC 加密 |
| `J加解密_AES_CBC解密` | `密文Base64: string, 密钥: string(Base64), IV: string(Base64)` | `string(Base64)` | AES-CBC 解密 |
| `J加解密_AES_GCM加密` | `明文: string(Base64), 密钥: string(Base64), 附加数据: string(Base64)` | `string` | AES-GCM 加密 |
| `J加解密_AES_GCM解密` | `密文Base64: string, 密钥: string(Base64), 附加数据: string(Base64)` | `string(Base64)` | AES-GCM 解密 |
| `J加解密_AES_ECB加密` | `明文: string(Base64), 密钥: string(Base64)` | `string` | AES-ECB 加密 |
| `J加解密_AES_ECB解密` | `密文Base64: string, 密钥: string(Base64)` | `string(Base64)` | AES-ECB 解密 |
| `J加解密_AES_CTR加密` | `明文: string(Base64), 密钥: string(Base64)` | `string` | AES-CTR 加密 |
| `J加解密_AES_CTR解密` | `密文Base64: string, 密钥: string(Base64)` | `string(Base64)` | AES-CTR 解密 |
| `J加解密_RC4` | `数据: string(Base64), 密钥: string(Base64)` | `string` | RC4 加解密 |
| `J加解密_XOR` | `数据: string(Base64), 密钥: string(Base64)` | `string(Base64)` | XOR 加解密 |
| `J加解密_TEA加密` | `明文: string(Base64), 密钥: string(Base64)` | `string` | TEA 加密 |
| `J加解密_TEA解密` | `密文Base64: string, 密钥: string(Base64)` | `string(Base64)` | TEA 解密 |
| `J加解密_XXTEA加密` | `明文: string(Base64), 密钥: string(Base64)` | `string` | XXTEA 加密 |
| `J加解密_XXTEA解密` | `密文Base64: string, 密钥: string(Base64)` | `string(Base64)` | XXTEA 解密 |
| `J加解密_生成AES密钥` | `长度: int` | `string(Base64)` | 生成 AES 密钥（16/24/32） |
| `J加解密_生成IV` | `块大小: int` | `string(Base64)` | 生成随机 IV |

### 4.6 校验哈希（J校验_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `J校验_取md5` | `字节集数据: string(Base64), 返回值转成大写: bool` | `string` | MD5 哈希 |
| `J校验_取md5_文本` | `文本数据: string, 返回值转成大写: bool` | `string` | MD5 哈希（文本输入） |
| `J校验_取md5_文件` | `文件路径: string, 返回值转成大写: bool` | `string` | MD5 哈希（文件） |
| `J校验_取sha256` | `数据: string(Base64), 返回值转成大写: bool` | `string` | SHA256 哈希 |
| `J校验_HMAC_SHA256` | `密钥: string(Base64), 数据: string(Base64), 返回值转成大写: bool` | `string` | HMAC-SHA256 |
| `J校验_取CRC64` | `数据: string(Base64), 返回值转成大写: bool` | `string` | CRC64 校验 |

### 4.7 核心库与类型转换（D_ / C类型_ / Q取随机数）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `D到文本` | `value: any` | `string` | 转换为文本 |
| `D到整数` | `value: any` | `int` | 转换为整数 |
| `D到数值` | `value: any` | `float64` | 转换为浮点数 |
| `D到字节集` | `value: any` | `string(Base64)` | 转换为字节集 |
| `Q取随机数` | `min: int, max: int` | `int` | 获取随机整数 |
| `C类型_到文本` | `值: any` | `string` | 类型转换到文本 |
| `C类型_到整数` | `值: any` | `int` | 类型转换到整数 |
| `C类型_到浮点数` | `值: any` | `float64` | 类型转换到浮点数 |
| `C类型_到逻辑型` | `值: any` | `bool` | 类型转换到布尔 |
| `C类型_安全到文本` | `值: any, 默认值: string` | `string` | 安全转换到文本 |
| `C类型_进制转换` | `文本: string, 进制: int` | `int64` | 进制转换 |

### 4.8 数组操作（S数组_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `S数组_是否为空` | `list: []string` | `bool` | 判断字符串数组是否为空 |
| `S数组_取文本索引` | `文本数组: []string, 文本: string` | `int` | 查找文本索引 |
| `S数组_取文本出现次数` | `参数_数组: []string, 参数_成员: string` | `int` | 统计成员出现次数 |
| `S数组_整数是否存在` | `数组: []int, 整数: int` | `bool` | 检查整数是否存在 |
| `S数组_排序整数` | `arr: []int` | `[]int` | 整数数组升序排序 |
| `S数组_排序文本` | `arr: []string` | `[]string` | 字符串数组排序 |
| `S数组_求平均值` | `参数: []int` | `int` | 整数数组求平均值 |
| `S数组_取随机成员` | `源数组: []string, 数量: int` | `[]string` | 随机选取成员 |
| `S数组_整数取差集` | `int1: []int, int2: []int` | `[]int` | 整数数组差集 |
| `S数组_取差集` | `a: []int, b: []int` | `[]int` | 整数数组差集 |

### 4.9 系统信息（X系统_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `X系统_取CPU核心数` | 无 | `int` | 获取 CPU 逻辑核心数 |
| `X系统_取CPU物理核心数` | 无 | `int` | 获取 CPU 物理核心数 |
| `X系统_取总CPU使用率` | `间隔: int` | `float64` | 获取总 CPU 使用率（0-100） |
| `X系统_取内存信息` | 无 | `object` | 获取内存使用情况 |
| `X系统_取主机信息` | 无 | `object` | 获取主机信息 |
| `X系统_取磁盘使用量` | `路径: string` | `object` | 获取磁盘使用量 |
| `X系统_取开机时间` | 无 | `uint64` | 获取系统开机时长（秒） |
| `X系统_取进程名` | `pid: int32` | `string` | 根据 PID 获取进程名 |
| `X系统_取进程内存占用` | `pid: int32` | `float64` | 获取进程内存占用（字节） |
| `X系统_取进程CPU占用` | `pid: int32` | `float64` | 获取进程 CPU 占用 |
| `X系统_取当前进程ID` | 无 | `int32` | 获取当前进程 PID |
| `X系统_是否64位系统` | 无 | `bool` | 判断是否 64 位系统 |
| `X系统_取系统架构` | 无 | `string` | 获取系统架构 |
| `X系统_取操作系统类型` | 无 | `string` | 获取操作系统类型 |
| `X系统_取逻辑处理器数` | 无 | `int` | 获取逻辑处理器数 |
| `X系统_取Go版本` | 无 | `string` | 获取 Go 版本 |

### 4.10 线程操作（X线程_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `X线程_延时` | `毫秒: uint32` | `bool` | 线程延时（Sleep） |

### 4.11 环境变量（K环境_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `K环境_取值` | `名称: string` | `string` | 获取环境变量 |
| `K环境_取值带默认值` | `名称: string, 默认值: string` | `string` | 获取环境变量（带默认值） |
| `K环境_设置值` | `名称: string, 值: string` | `bool` | 设置环境变量 |
| `K环境_删除值` | `名称: string` | `bool` | 删除环境变量 |
| `K环境_是否存在` | `名称: string` | `bool` | 判断环境变量是否存在 |
| `K环境_取所有` | 无 | `[]string` | 获取所有环境变量 |
| `K环境_加载` | `文件路径: string` | `bool` | 加载 .env 文件 |

### 4.12 HTTP 客户端（H客户端_）

| 函数名 | 参数 | 返回类型 | 说明 |
|--------|------|----------|------|
| `H客户端_取文本` | `网址: string` | `string` | HTTP GET 获取文本 |
| `H客户端_下载文件` | `网址: string, 保存路径: string` | `bool` | 下载文件 |

## 五、编译方法

### 5.1 前置条件

1. **Go 1.25+**：安装 Go 编译器
2. **MinGW-w64**：安装 GCC 编译器（CGO 需要），推荐 [winlibs](https://winlibs.com/) 或 [MSYS2](https://www.msys2.org/)
3. **依赖下载**：在项目根目录执行 `go mod tidy`

### 5.2 编译命令

```bash
cd cmd/export_dll

# 方式一：使用构建脚本
build.bat

# 方式二：手动编译
set CGO_ENABLED=1
set GOOS=windows
go build -buildmode=c-shared -o Efunc.dll .
```

编译成功后生成两个文件：
- `Efunc.dll`：动态链接库
- `Efunc.h`：C 头文件（Go 自动生成）

### 5.3 交叉编译 32 位

```bash
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=386
set CC=i686-w64-mingw32-gcc
go build -buildmode=c-shared -o Efunc32.dll .
```

## 六、各语言调用示例

### 6.1 C 语言

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 声明 DLL 导出函数（或 #include "Efunc.h"）
extern char* Efunc_Call(char* name, char* params);
extern void  Efunc_Free(void* ptr);
extern char* Efunc_List();

int main() {
    // 示例1：获取文本长度
    char* r1 = Efunc_Call("W文本_取长度", "{\"value\": \"你好世界\"}");
    printf("文本长度: %s\n", r1);  // {"ok":true,"data":4}
    Efunc_Free(r1);

    // 示例2：MD5 哈希
    char* r2 = Efunc_Call("J校验_取md5_文本",
        "{\"文本数据\": \"hello\", \"返回值转成大写\": false}");
    printf("MD5: %s\n", r2);
    Efunc_Free(r2);

    // 示例3：获取当前时间戳
    char* r3 = Efunc_Call("S时间_取现行时间戳", "{}");
    printf("时间戳: %s\n", r3);
    Efunc_Free(r3);

    // 示例4：AES 加密（[]byte 用 Base64 传递）
    // "Hello" 的 Base64 = "SGVsbG8="
    // 16字节密钥 "1234567890123456" 的 Base64 = "MTIzNDU2Nzg5MDEyMzQ1Ng=="
    char* r4 = Efunc_Call("J加解密_AES_CBC加密",
        "{\"明文\": \"SGVsbG8=\","
        " \"密钥\": \"MTIzNDU2Nzg5MDEyMzQ1Ng==\","
        " \"IV\": \"MTIzNDU2Nzg5MDEyMzQ1Ng==\"}");
    printf("AES加密: %s\n", r4);
    Efunc_Free(r4);

    return 0;
}
```

编译：`gcc main.c -L. -lEfunc -o main.exe`

### 6.2 Python

```python
import ctypes
import json
import base64

dll = ctypes.CDLL("Efunc.dll")

# 设置函数签名
dll.Efunc_Call.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
dll.Efunc_Call.restype = ctypes.c_char_p
dll.Efunc_Free.argtypes = [ctypes.c_void_p]
dll.Efunc_Free.restype = None

def efunc_call(name, params=None):
    """调用 Efunc DLL 函数"""
    if params is None:
        params = {}
    name_bytes = name.encode("utf-8")
    params_bytes = json.dumps(params, ensure_ascii=False).encode("utf-8")
    result_ptr = dll.Efunc_Call(name_bytes, params_bytes)
    result_str = ctypes.string_at(result_ptr).decode("utf-8")
    dll.Efunc_Free(result_ptr)
    return json.loads(result_str)

def to_b64(data: bytes) -> str:
    """字节集转 Base64 字符串"""
    return base64.b64encode(data).decode("ascii")

def from_b64(s: str) -> bytes:
    """Base64 字符串转字节集"""
    return base64.b64decode(s)

# 示例：获取文本长度
r = efunc_call("W文本_取长度", {"value": "你好世界"})
print(r)  # {'ok': True, 'data': 4}

# 示例：MD5
r = efunc_call("J校验_取md5_文本", {"文本数据": "hello", "返回值转成大写": False})
print(r)  # {'ok': True, 'data': '5d41402abc4b2a76b9719d911017c592'}

# 示例：AES 加密
key = to_b64(b"1234567890123456")
iv = to_b64(b"1234567890123456")
plaintext = to_b64(b"Hello World")
r = efunc_call("J加解密_AES_CBC加密", {"明文": plaintext, "密钥": key, "IV": iv})
print(r)  # {'ok': True, 'data': '...base64密文...'}

# 示例：文件操作
r = efunc_call("W文件_是否存在", {"路径": "C:\\Windows"})
print(r)  # {'ok': True, 'data': True}

# 示例：系统信息
r = efunc_call("X系统_取主机信息", {})
print(r)  # {'ok': True, 'data': {...}}
```

### 6.3 C# (.NET)

```csharp
using System;
using System.Runtime.InteropServices;
using System.Text.Json;

class EfuncDLL
{
    [DllImport("Efunc.dll", CallingConvention = CallingConvention.Cdecl)]
    private static extern IntPtr Efunc_Call(string name, string params);

    [DllImport("Efunc.dll", CallingConvention = CallingConvention.Cdecl)]
    private static extern void Efunc_Free(IntPtr ptr);

    [DllImport("Efunc.dll", CallingConvention = CallingConvention.Cdecl)]
    private static extern IntPtr Efunc_List();

    public static JsonElement Call(string name, object? parameters = null)
    {
        string jsonParams = parameters == null ? "{}" :
            JsonSerializer.Serialize(parameters);
        IntPtr resultPtr = Efunc_Call(name, jsonParams);
        string resultJson = Marshal.PtrToStringUTF8(resultPtr) ?? "";
        Efunc_Free(resultPtr);
        return JsonDocument.Parse(resultJson).RootElement;
    }
}

// 使用示例
var r1 = EfuncDLL.Call("W文本_取长度", new { value = "你好世界" });
Console.WriteLine(r1);  // {"ok":true,"data":4}

var r2 = EfuncDLL.Call("J校验_取md5_文本",
    new { 文本数据 = "hello", 返回值转成大写 = false });
Console.WriteLine(r2.GetProperty("data"));
```

### 6.4 易语言

```
.版本 2

.DLL命令 Efunc_Call, 整数型, "Efunc.dll", "Efunc_Call"
    .参数 name, 文本型
    .参数 params, 文本型

.DLL命令 Efunc_Free, , "Efunc.dll", "Efunc_Free"
    .参数 ptr, 整数型

.子程序 调用Efunc, 文本型
.参数 函数名, 文本型
.参数 JSON参数, 文本型
.局部变量 结果指针, 整数型
.局部变量 结果文本, 文本型

结果指针 = Efunc_Call(函数名, JSON参数)
结果文本 = 指针到文本(结果指针)
Efunc_Free(结果指针)
返回(结果文本)

' 使用示例：
' 结果 = 调用Efunc("W文本_取长度", "{""value"": ""你好世界""}")
' 结果 = 调用Efunc("J校验_取md5_文本", "{""文本数据"": ""hello"", ""返回值转成大写"": false}")
```

## 七、架构设计

### 7.1 文件结构

```
cmd/export_dll/
├── main.go                  # DLL 入口，//export 导出的 C 函数
├── registry.go              # 函数注册表、JSON 分发器、panic 恢复
├── wrap_text.go             # W文本_ 系列包装器
├── wrap_file.go             # W文件_ 系列包装器
├── wrap_time_encoding.go    # S时间_ + B编码_ + Z字节集_ 包装器
├── wrap_crypto.go           # J加解密_ + J校验_ 包装器
├── wrap_core.go             # 核心库 + C类型转换_ + S数组_ 包装器
├── wrap_system.go           # X系统_ + K环境_ + H客户端_ + X线程_ 包装器
└── build.bat                # 一键构建脚本
```

### 7.2 调用流程

```
调用方                    DLL
  │                       │
  │  Efunc_Call(name, json) │
  │──────────────────────>│
  │                       │  1. C.GoString 转换参数
  │                       │  2. Registry.Call 查找函数
  │                       │  3. json.Unmarshal 解析参数
  │                       │  4. 调用 Go 原始函数
  │                       │  5. json.Marshal 序列化结果
  │                       │  6. C.CString 返回（含 panic 恢复）
  │  <返回 JSON 字符串>    │
  │                       │
  │  Efunc_Free(ptr)      │
  │──────────────────────>│
  │                       │  释放 C.CString 内存
```

### 7.3 如何添加新函数

在任何 `wrap_*.go` 文件中，按以下模式添加：

```go
r.Register("函数名", []string{"参数1", "参数2"}, "函数说明",
    func(p json.RawMessage) *CallResult {
        var v struct {
            参数1 string `json:"参数1"`
            参数2 int    `json:"参数2"`
        }
        if err := json.Unmarshal(p, &v); err != nil {
            return errResult(err.Error())
        }
        result, err := utils.函数名(v.参数1, v.参数2)
        if err != nil {
            return errResult(err.Error())
        }
        return okResult(result)
    })
```

对于 `[]byte` 参数，使用 `decodeBase64` 解码输入，使用 `encodeBase64` 编码输出：

```go
r.Register("示例_字节集函数", []string{"数据"}, "处理字节集",
    func(p json.RawMessage) *CallResult {
        var v struct {
            数据 string `json:"数据"`
        }
        if err := json.Unmarshal(p, &v); err != nil {
            return errResult(err.Error())
        }
        data, err := decodeBase64(v.数据)
        if err != nil {
            return errResultf("base64解码失败: %v", err)
        }
        result := utils.某函数(data)
        return okResult(encodeBase64(result))
    })
```

## 八、注意事项

1. **内存管理**：每次 `Efunc_Call` 返回的字符串**必须**调用 `Efunc_Free` 释放，否则内存泄漏
2. **线程安全**：DLL 内部使用 `sync.RWMutex` 保护注册表，支持多线程并发调用
3. **panic 恢复**：所有调用均被 `defer recover` 包裹，Go 侧 panic 不会导致宿主进程崩溃
4. **Go Runtime**：DLL 加载时会初始化 Go runtime，首次调用可能有短暂延迟
5. **中文函数名**：DLL 导出函数名为英文（`Efunc_Call` 等），中文函数名作为参数传入，无兼容性问题
6. **Base64 编码**：所有 `[]byte` 类型参数和返回值均使用标准 Base64 编码
7. **错误处理**：调用失败时 `ok` 字段为 `false`，`error` 字段包含错误描述
8. **无参数函数**：传空 JSON 对象 `{}` 即可

## 九、未覆盖函数说明

以下类型的函数暂未导出，可按需逐步添加：

| 类型 | 原因 | 示例 |
|------|------|------|
| 泛型函数 | Go 泛型无法直接反射调用 | `S三元[T]`, `D多项选择[T]`, `S数组_去重复[T]` |
| 可变参数函数 | JSON 不支持可变参数 | `G格式化文本(str, ...interface{})` |
| 回调函数参数 | C/Go 回调传递复杂 | `X线程_创建(线程函数 uintptr, ...)`, `G协程池_创建` |
| 结构体方法 | 需要对象生命周期管理 | `类_HTTP`, `类_TCP`, `类_WebSocket` 等的方法 |
| `interface{}` 输出参数 | JSON 反序列化无法还原类型 | `B编码_JSON解码(文本, 目标 interface{})` |
| 指针参数 | 跨语言指针不安全 | `W文件_枚举(..., files *[]string, ...)` |

对于泛型函数，可以为每个具体类型创建特化版本，例如：
- `S数组_去重复_文本` → `S数组_去重复[string]`
- `S数组_去重复_整数` → `S数组_去重复[int]`
