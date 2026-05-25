# Efunc 英文版 API 参考文档

> 生成时间：2026-05-25 北京时间  
> Go 版本：1.25+  
> 包路径：`github.com/yuan71058/Efunc/en`

---

## 目录

- [1. 核心工具 (core.go)](#1-核心工具-corego)
- [2. 类型转换 (cast.go)](#2-类型转换-castgo)
- [3. 数组与映射 (arraygo-map_utilsgo)](#3-数组与映射-arraygo-map_utilsgo)
- [4. 文本与字符串 (text.go)](#4-文本与字符串-textgo)
- [5. 编码 (encoding.go)](#5-编码-encodinggo)
- [6. 字节数据 (bytedata.go)](#6-字节数据-bytedatago)
- [7. 整数与浮点转换 (intconvgo-float64convgo)](#7-整数与浮点转换-intconvgo-float64convgo)
- [8. IP 地址 (ip.go)](#8-ip-地址-ipgo)
- [9. 时间 (time.go)](#9-时间-timego)
- [10. 加解密与校验 (cryptogo-checksumgo-rsago)](#10-加解密与校验-cryptogo-checksumgo-rsago)
- [11. 文件与目录 (filego-directorygo)](#11-文件与目录-filego-directorygo)
- [12. HTTP 客户端 (http_utils.go)](#12-http-客户端-http_utilsgo)
- [13. 正则表达式 (regex_utils.go)](#13-正则表达式-regex_utilsgo)
- [14. 环境变量与日志 (envgo-loggo)](#14-环境变量与日志-envgo-loggo)
- [15. 程序信息 (programgo-helpergo)](#15-程序信息-programgo-helpergo)
- [16. POST 数据 (postdata.go)](#16-post-数据-postdatago)
- [17. 原子操作 (atomic.go)](#17-原子操作-atomicgo)
- [18. 数据校验 (validation.go)](#18-数据校验-validationgo)
- [19. 命令行参数 (cli.go)](#19-命令行参数-cligo)
- [20. 配置管理 (config.go)](#20-配置管理-configgo)
- [21. JSON 操作 (json_utils.go)](#21-json-操作-json_utilsgo)
- [22. 协程池 (goroutine_pool.go)](#22-协程池-goroutine_poolgo)
- [23. 定时任务 (cron.go)](#23-定时任务-crongo)
- [24. 消息总线 (message_bus.go)](#24-消息总线-message_busgo)
- [25. 字节缓冲池 (bytebuffer_pool.go)](#25-字节缓冲池-bytebuffer_poolgo)
- [26. 结构体合并 (struct_merge.go)](#26-结构体合并-struct_mergego)
- [27. 表达式求值 (expression.go)](#27-表达式求值-expressiongo)
- [28. 日期解析 (date_parse.go)](#28-日期解析-date_parsego)
- [29. 文件监控 (file_watcher.go)](#29-文件监控-file_watchergo)
- [30. 模板处理 (template.go)](#30-模板处理-templatego)
- [31. 系统信息 (system_info.go)](#31-系统信息-system_infogo)
- [32. 控制台表格 (table.go)](#32-控制台表格-tablego)
- [33. 键值数据库 (buntdb.go)](#33-键值数据库-buntdbgo)
- [34. 邮件发送 (email.go)](#34-邮件发送-emailgo)
- [35. 图片处理 (image.go)](#35-图片处理-imagego)
- [36. 网页工具 (web_utils.go)](#36-网页工具-web_utilsgo)
- [37. HTTP 高级客户端 (http_client.go)](#37-http-高级客户端-http_clientgo)
- [38. 网页爬虫 (spider.go)](#38-网页爬虫-spidergo)
- [39. 权限管理 (permission.go)](#39-权限管理-permissiongo)
- [40. 数据库 (database.go)](#40-数据库-databasego)
- [41. Windows：进程 (process_windows.go)](#41-windows进程-process_windowsgo)
- [42. Windows：系统命令 (system_command_windows.go)](#42-windows系统命令-system_command_windowsgo)
- [43. Windows：线程 (thread_windows.go)](#43-windows线程-thread_windowsgo)
- [44. Windows：键盘鼠标 (input_windows.go)](#44-windows键盘鼠标-input_windowsgo)
- [45. Windows：窗口 (window_windows.go)](#45-windows窗口-window_windowsgo)
- [46. Windows：内存操作 (memory_windows.go)](#46-windows内存操作-memory_windowsgo)
- [47. Windows：API Hook (hook.go)](#47-windowsapi-hook-hookgo)
- [48. Windows：汇编引擎 (asm_assembler.go)](#48-windows汇编引擎-asm_assemblergo)
- [49. Windows：汇编指令 (asm_instructions.go)](#49-windows汇编指令-asm_instructionsgo)
- [50. OpenCV 视觉 (opencv.go)](#50-opencv-视觉-opencvgo)
- [51. class 模块](#51-class-模块)

---

## 1. 核心工具 (core.go)

通用类型转换、文本格式化与编码工具。

| 函数 | 签名 | 说明 |
|------|------|------|
| `ToBytes` | `func ToBytes(value interface{}) []byte` | 通用类型转 `[]byte` |
| `ToByte` | `func ToByte(value interface{}) byte` | 通用类型转 `byte` |
| `ToInt` | `func ToInt(value interface{}) int` | 通用类型转 `int` |
| `ToInt64` | `func ToInt64(value interface{}) int64` | 通用类型转 `int64` |
| `ToFloat64` | `func ToFloat64(value interface{}) float64` | 通用类型转 `float64` |
| `ToString` | `func ToString(value interface{}) string` | 通用类型转 `string` |
| `ToStruct` | `func ToStruct(source, target interface{}) error` | 结构体/Map 互转（基于 gconv） |
| `Ternary` | `func Ternary[T any](cond bool, t, f T) T` | 泛型三元运算符 |
| `MultiSelect` | `func MultiSelect[T any](idx int, arr []T, def T) T` | 安全索引选取，带默认值 |
| `FormatText` | `func FormatText(format string, args ...interface{}) string` | `fmt.Sprintf` 封装 |
| `FormatJSON` | `func FormatJSON(data string) string` | JSON 字符串格式化美化 |
| `ToStringArray` | `func ToStringArray(value interface{}) []string` | 转为字符串切片 |
| `IsArray` | `func IsArray(value interface{}) bool` | 判断是否为数组/切片 |
| `GBKToUTF8` | `func GBKToUTF8(src string) string` | GBK 转 UTF-8 |
| `UTF8ToGBK` | `func UTF8ToGBK(src string) string` | UTF-8 转 GBK |
| `RandomInt` | `func RandomInt(min, max int) int` | 生成 `[min, max]` 范围内随机整数 |

---

## 2. 类型转换 (cast.go)

基于 `spf13/cast` 的安全类型转换，含默认值回退版本。

| 函数 | 签名 | 说明 |
|------|------|------|
| `Cast_ToString` | `func Cast_ToString(v interface{}) string` | interface{} 转 string |
| `Cast_ToInt` | `func Cast_ToInt(v interface{}) int` | interface{} 转 int |
| `Cast_ToInt64` | `func Cast_ToInt64(v interface{}) int64` | interface{} 转 int64 |
| `Cast_ToFloat64` | `func Cast_ToFloat64(v interface{}) float64` | interface{} 转 float64 |
| `Cast_ToBool` | `func Cast_ToBool(v interface{}) bool` | interface{} 转 bool |
| `Cast_ToStringSlice` | `func Cast_ToStringSlice(v interface{}) []string` | interface{} 转 []string |
| `Cast_ToIntSlice` | `func Cast_ToIntSlice(v interface{}) []int` | interface{} 转 []int |
| `Cast_ToTime` | `func Cast_ToTime(v interface{}) interface{}` | interface{} 转 time.Time |
| `Cast_ToDuration` | `func Cast_ToDuration(v interface{}) interface{}` | interface{} 转 time.Duration |
| `Cast_SafeToString` | `func Cast_SafeToString(v interface{}, def string) string` | 安全转换，带默认值 |
| `Cast_SafeToInt` | `func Cast_SafeToInt(v interface{}, def int) int` | 安全转换，带默认值 |
| `Cast_SafeToFloat64` | `func Cast_SafeToFloat64(v interface{}, def float64) float64` | 安全转换，带默认值 |
| `Cast_SafeToBool` | `func Cast_SafeToBool(v interface{}, def bool) bool` | 安全转换，带默认值 |
| `Cast_ParseInt` | `func Cast_ParseInt(s string, base int) (int64, error)` | 按进制解析字符串为整数 |

---

## 3. 数组与映射 (array.go, map_utils.go)

| 函数 | 签名 | 说明 |
|------|------|------|
| `Array_RemoveDuplicates` | `func Array_RemoveDuplicates(arr []string) []string` | 移除字符串数组重复项 |
| `Array_InArray` | `func Array_InArray(val string, arr []string) bool` | 判断值是否在数组中 |
| `Array_Join` | `func Array_Join(arr []string, sep string) string` | 用分隔符连接字符串数组 |
| `Array_Reverse` | `func Array_Reverse(arr []string) []string` | 反转字符串数组顺序 |
| `Array_Subset` | `func Array_Subset(arr []string, start, length int) []string` | 提取子数组 |
| `Map_GetIntKeys` | `func Map_GetIntKeys(m map[int]string) []int` | 获取 map[int]string 的所有键 |
| `Map_StructToMap` | `func Map_StructToMap(obj interface{}) map[string]interface{}` | 结构体转 map |
| `Map_ToPostData` | `func Map_ToPostData(params map[string]string, urlEncode bool) string` | map 转 POST 参数字符串 |
| `Map_KeyExists` | `func Map_KeyExists(m map[int]string, key int) bool` | 判断键是否存在 |

---

## 4. 文本与字符串 (text.go)

| 函数 | 签名 | 说明 |
|------|------|------|
| `Text_Between` | `func Text_Between(s, start, end string) string` | 取出两个标记之间的文本 |
| `Text_Before` | `func Text_Before(s, sep string) string` | 取分隔符之前的文本 |
| `Text_After` | `func Text_After(s, sep string) string` | 取分隔符之后的文本 |
| `Text_Contains` | `func Text_Contains(s, substr string) bool` | 判断是否包含子串 |
| `Text_Replace` | `func Text_Replace(s, old, new string, n int) string` | 替换子串（-1 表示全部） |
| `Text_TrimSpace` | `func Text_TrimSpace(s string) string` | 去除首尾空白 |
| `Text_Count` | `func Text_Count(s, substr string) int` | 统计子串出现次数 |
| `Text_Split` | `func Text_Split(s, sep string) []string` | 按分隔符拆分字符串 |
| `Text_StartsWith` | `func Text_StartsWith(s, prefix string) bool` | 判断是否以指定前缀开头 |
| `Text_EndsWith` | `func Text_EndsWith(s, suffix string) bool` | 判断是否以指定后缀结尾 |
| `Text_ToUpper` | `func Text_ToUpper(s string) string` | 转为大写（前 N 字符） |
| `Text_ToLower` | `func Text_ToLower(s string) string` | 转为小写（前 N 字符） |
| `Text_Reverse` | `func Text_Reverse(s string) string` | 反转字符串 |
| `Text_Left` | `func Text_Left(s string, n int) string` | 取左边 N 个字符 |
| `Text_Right` | `func Text_Right(s string, n int) string` | 取右边 N 个字符 |
| `Text_Mid` | `func Text_Mid(s string, start, length int) string` | 取中间指定长度子串 |
| `Text_Len` | `func Text_Len(s string) int` | 获取字符串长度 |
| `Text_PadLeft` | `func Text_PadLeft(s string, pad string, totalLen int) string` | 左侧填充至指定长度 |
| `Text_PadRight` | `func Text_PadRight(s string, pad string, totalLen int) string` | 右侧填充至指定长度 |

---

## 5. 编码 (encoding.go)

URL 与 Base64 编解码工具。

| 函数 | 签名 | 说明 |
|------|------|------|
| `Encoding_URLEncode` | `func Encoding_URLEncode(s string) string` | URL 编码 |
| `Encoding_URLDecode` | `func Encoding_URLDecode(s string) (string, error)` | URL 解码 |
| `Encoding_Base64Encode` | `func Encoding_Base64Encode(data []byte) string` | Base64 标准编码 |
| `Encoding_Base64Decode` | `func Encoding_Base64Decode(s string) ([]byte, error)` | Base64 标准解码 |
| `Encoding_Base64URLEncode` | `func Encoding_Base64URLEncode(data []byte) string` | Base64 URL 安全编码 |
| `Encoding_Base64URLDecode` | `func Encoding_Base64URLDecode(s string) ([]byte, error)` | Base64 URL 安全解码 |
| `Encoding_UnicodeToUTF8` | `func Encoding_UnicodeToUTF8(src string) string` | \\uXXXX 转义 Unicode 转 UTF-8 |
| `Encoding_UTF8ToUnicode` | `func Encoding_UTF8ToUnicode(src string) string` | UTF-8 转 \\uXXXX 转义 Unicode |

---

## 6. 字节数据 (bytedata.go)

| 函数 | 签名 | 说明 |
|------|------|------|
| `ByteData_BinaryToHex` | `func ByteData_BinaryToHex(data []byte) string` | 字节集转十六进制字符串 |
| `ByteData_HexToBinary` | `func ByteData_HexToBinary(hexStr string) ([]byte, error)` | 十六进制字符串转字节集 |
| `ByteData_BinaryToString` | `func ByteData_BinaryToString(data []byte) string` | 字节集按 GBK 解释后转 UTF-8 |
| `ByteData_StringToBinary` | `func ByteData_StringToBinary(s string) []byte` | 字符串按 GBK 编码转字节集 |
| `ByteData_SubBinary` | `func ByteData_SubBinary(data []byte, start, length int) []byte` | 安全截取子字节集 |

---

## 7. 整数与浮点转换 (intconv.go, float64conv.go)

| 函数 | 签名 | 说明 |
|------|------|------|
| `IntToBytes` | `func IntToBytes(n interface{}) []byte` | int/uint 转 []byte |
| `BytesToInt` | `func BytesToInt(b []byte) int64` | []byte 转 int |
| `Float64ToByte` | `func Float64ToByte(v float64) byte` | float64 转 byte |
| `ByteToFloat64` | `func ByteToFloat64(v byte) float64` | byte 转 float64 |

---

## 8. IP 地址 (ip.go)

| 函数 | 签名 | 说明 |
|------|------|------|
| `IP_IsValid` | `func IP_IsValid(ipStr string) bool` | 判断是否为合法 IP 地址 |
| `IP_ToInt` | `func IP_ToInt(ipStr string) (uint32, error)` | IPv4 转 uint32 |
| `IP_FromInt` | `func IP_FromInt(ipInt uint32) string` | uint32 转 IPv4 |
| `IP_IsPrivate` | `func IP_IsPrivate(ipStr string) bool` | 判断是否为私有 IP |
| `IP_GetLocalIP` | `func IP_GetLocalIP() (string, error)` | 获取本机 IP |
| `IP_GetNetwork` | `func IP_GetNetwork(cidr string) (string,string,error)` | 解析 CIDR 标记 |
| `IP_Subtract` | `func IP_Subtract(ipStr string, offset int) (string, error)` | IP 地址减去偏移量 |
| `IP_Add` | `func IP_Add(ipStr string, offset int) (string, error)` | IP 地址加上偏移量 |
| `IP_Distance` | `func IP_Distance(ip1, ip2 string) (int, error)` | 计算两个 IP 之间的距离 |

---

## 9. 时间 (time.go)

| 函数 | 签名 | 说明 |
|------|------|------|
| `Time_CurrentTimestamp` | `func Time_CurrentTimestamp() int64` | 当前 Unix 时间戳（秒） |
| `Time_CurrentTimestampMs` | `func Time_CurrentTimestampMs() int64` | 当前 Unix 时间戳（毫秒） |
| `Time_Format` | `func Time_Format(t time.Time, layout string) string` | 格式化时间 |
| `Time_FormatNow` | `func Time_FormatNow(layout string) string` | 格式化当前时间 |
| `Time_Parse` | `func Time_Parse(layout, value string) (time.Time, error)` | 解析时间字符串 |
| `Time_UnixToTime` | `func Time_UnixToTime(sec int64) time.Time` | Unix 时间戳转 time.Time |
| `Time_AddDays` | `func Time_AddDays(t time.Time, days int) time.Time` | 增加/减少天数 |
| `Time_DiffDays` | `func Time_DiffDays(t1, t2 time.Time) int` | 计算日期差（天） |
| `Time_NTPTime` | `func Time_NTPTime(server string) (time.Time, error)` | 获取 NTP 时间 |

---

## 10. 加解密与校验 (crypto.go, checksum.go, rsa.go)

### 10.1 对称加密

| 函数 | 说明 |
|------|------|
| `Crypto_AESEncrypt(plaintext, key []byte) ([]byte, error)` | AES-256-CBC 加密（PKCS7 填充） |
| `Crypto_AESDecrypt(ciphertext, key []byte) ([]byte, error)` | AES-256-CBC 解密 |
| `Crypto_DESEncrypt(plaintext, key []byte) ([]byte, error)` | DES-CBC 加密 |
| `Crypto_DESDecrypt(ciphertext, key []byte) ([]byte, error)` | DES-CBC 解密 |
| `Crypto_3DESEncrypt(plaintext, key []byte) ([]byte, error)` | 3DES-CBC 加密 |
| `Crypto_3DESDecrypt(ciphertext, key []byte) ([]byte, error)` | 3DES-CBC 解密 |
| `Crypto_HmacMD5(data, key []byte) string` | HMAC-MD5 十六进制 |
| `Crypto_HmacSHA1(data, key []byte) string` | HMAC-SHA1 十六进制 |
| `Crypto_HmacSHA256(data, key []byte) string` | HMAC-SHA256 十六进制 |
| `Crypto_HmacSHA512(data, key []byte) string` | HMAC-SHA512 十六进制 |

### 10.2 哈希校验

| 函数 | 说明 |
|------|------|
| `Checksum_MD5(data []byte) string` | MD5 哈希 |
| `Checksum_MD5File(filePath string) (string, error)` | 文件 MD5 哈希 |
| `Checksum_SHA1(data []byte) string` | SHA-1 哈希 |
| `Checksum_SHA256(data []byte) string` | SHA-256 哈希 |
| `Checksum_SHA512(data []byte) string` | SHA-512 哈希 |
| `Checksum_CRC32(data []byte) uint32` | CRC32-IEEE |

### 10.3 RSA

| 函数 | 说明 |
|------|------|
| `RSA_GenerateKey() (error, string, string)` | 生成密钥对（PEM 格式） |
| `RSA_SignWithPrivateKey(plaintext, privPEM string) string` | 签名，返回 Base64 |
| `RSA_EncryptWithPublicKey(pubPEM string, plaintext []byte) string` | 公钥加密 |
| `RSA_DecryptWithPrivateKey(privBytes, encData []byte) string` | 私钥解密 |
| `RSA_EncryptWithPrivateKey(privBytes, plaintext []byte) string` | 私钥加密 |
| `RSA_DecryptWithPublicKey(pubPEM string, cipher []byte) []byte` | 公钥解密 |

---

## 11. 文件与目录 (file.go, directory.go)

### 11.1 文件

| 函数 | 说明 |
|------|------|
| `File_Exists(path string) bool` | 判断文件是否存在 |
| `File_ReadAll(path string) ([]byte, error)` | 读取全部文件内容 |
| `File_WriteAll(path string, data []byte) error` | 写入字节到文件 |
| `File_ReadString(path string) (string, error)` | 读取文件为字符串 |
| `File_WriteString(path, content string) error` | 写入字符串到文件 |
| `File_AppendString(path, content string) error` | 追加字符串 |
| `File_AppendBytes(path string, data []byte) error` | 追加字节 |
| `File_Copy(src, dst string) error` | 复制文件 |
| `File_Move(src, dst string) error` | 移动/重命名文件 |
| `File_Delete(path string) error` | 删除文件 |
| `File_Size(path string) (int64, error)` | 获取文件大小 |
| `File_ReadLines(path string) ([]string, error)` | 读取文件所有行 |
| `File_GetExt(path string) string` | 获取文件扩展名 |
| `File_GetName(path string) string` | 获取文件名 |
| `File_GetNameNoExt(path string) string` | 获取不含扩展名的文件名 |
| `File_IsDir(path string) bool` | 判断是否为目录 |

### 11.2 目录

| 函数 | 说明 |
|------|------|
| `Dir_Create(path string) error` | 创建目录 |
| `Dir_Exists(path string) bool` | 判断目录是否存在 |
| `Dir_Delete(path string) error` | 删除目录及内容 |
| `Dir_ListFiles(path string) ([]string, error)` | 列出文件 |
| `Dir_ListDirs(path string) ([]string, error)` | 列出子目录 |
| `Dir_ListAll(path string) ([]string, error)` | 列出所有条目 |
| `Dir_Walk(root string, fn filepath.WalkFunc) error` | 遍历目录树 |
| `Dir_GetCurrent() (string, error)` | 获取当前工作目录 |

---

## 12. HTTP 客户端 (http_utils.go)

基于 `resty/v2`。

| 函数 | 说明 |
|------|------|
| `HTTP_Get(url string, headers map[string]string, timeout time.Duration) (string, error)` | GET 请求 |
| `HTTP_Post(url, contentType, body string, headers map[string]string, timeout time.Duration) (string, error)` | POST 请求 |
| `HTTP_PostJSON(url, jsonBody string, headers map[string]string, timeout time.Duration) (string, error)` | POST JSON |
| `HTTP_Put(url, contentType, body string, headers map[string]string, timeout time.Duration) (string, error)` | PUT 请求 |
| `HTTP_Delete(url string, headers map[string]string, timeout time.Duration) (string, error)` | DELETE 请求 |
| `HTTP_Download(url, savePath string, headers map[string]string, timeout time.Duration) error` | 下载文件 |
| `HTTP_GetResponseCode(url string, timeout time.Duration) (int, error)` | 获取状态码 |
| `HTTP_NewRequest() *resty.Request` | 创建高级请求对象 |

---

## 13. 正则表达式 (regex_utils.go)

| 函数 | 说明 |
|------|------|
| `Regex_Match(pattern, text string) bool` | 测试是否匹配 |
| `Regex_Find(pattern, text string) string` | 查找第一个匹配 |
| `Regex_FindAll(pattern, text string) []string` | 查找所有匹配 |
| `Regex_Replace(pattern, text, repl string) string` | 替换所有匹配 |
| `Regex_MatchGroups(pattern, text string) []string` | 获取捕获组 |
| `Regex_IsValidPattern(pattern string) bool` | 验证正则表达式是否合法 |

---

## 14. 环境变量与日志 (env.go, log.go)

### 14.1 环境变量

| 函数 | 说明 |
|------|------|
| `Env_Load(filenames ...string) error` | 加载 .env 文件 |
| `Env_Get(key string, defaultVal ...string) string` | 获取环境变量 |
| `Env_Set(key, value string) error` | 设置环境变量 |
| `Env_GetInt(key string, defaultVal int) int` | 获取为整数 |
| `Env_GetBool(key string, defaultVal bool) bool` | 获取为布尔值 |

### 14.2 日志

| 函数 | 说明 |
|------|------|
| `Log_Info/Debug/Warn/Error/Fatal(format string, v ...interface{})` | 各级别日志输出 |
| `Log_SetLevel(level string)` | 设置日志级别 |

---

## 15. 程序信息 (program.go, helper.go)

| 函数 | 说明 |
|------|------|
| `Program_GetProcessID() int` | 当前进程 ID |
| `Program_GetName() string` | 可执行文件名 |
| `Program_GetPath() string` | 可执行文件路径 |
| `Program_GetDir() string` | 可执行文件所在目录 |
| `Program_Sleep(milliseconds int)` | 休眠 |

---

## 16. POST 数据 (postdata.go)

| 函数 | 说明 |
|------|------|
| `PostData_New() url.Values` | 创建 POST 数据 |
| `PostData_Add(data url.Values, key, value string)` | 添加键值对 |
| `PostData_Encode(data url.Values) string` | 编码为字符串 |

---

## 17. 原子操作 (atomic.go)

| 函数 | 说明 |
|------|------|
| `Atomic_Increment(p *int64) int64` | 原子自增 |

---

## 18. 数据校验 (validation.go)

| 函数 | 说明 |
|------|------|
| `Validation_Validate(s interface{}) error` | 校验结构体标签 |
| `Validation_New() *validator.Validate` | 创建校验器实例 |
| `Validation_Var(field interface{}, tag string) error` | 校验单个变量 |

---

## 19. 命令行参数 (cli.go)

| 函数 | 说明 |
|------|------|
| `CLI_Parse()` | 解析命令行标志 |
| `CLI_GetString(short, long, def, usage string) *string` | 字符串标志 |
| `CLI_GetInt(short, long string, def int, usage string) *int` | 整数标志 |
| `CLI_GetBool(short, long string, def bool, usage string) *bool` | 布尔标志 |
| `CLI_GetFloat64(short, long string, def float64, usage string) *float64` | 浮点标志 |
| `CLI_GetArgs() []string` | 获取参数列表 |
| `CLI_GetProgramName() string` | 获取程序名称 |

---

## 20. 配置管理 (config.go)

基于 `spf13/viper`。

| 函数 | 说明 |
|------|------|
| `Config_New() *viper.Viper` | 创建配置管理器 |
| `Config_LoadFile(v *viper.Viper, path string) error` | 加载配置文件 |
| `Config_Get/GetString/GetInt/GetBool/GetFloat64` | 读取配置值 |
| `Config_Set(v *viper.Viper, key string, value interface{})` | 设置配置值 |
| `Config_SetDefault(v *viper.Viper, key string, value interface{})` | 设置默认值 |
| `Config_BindEnv(v *viper.Viper, key, envName string)` | 绑定环境变量 |
| `Config_AutoEnv(v *viper.Viper, prefix string)` | 自动绑定环境变量 |

---

## 21. JSON 操作 (json_utils.go)

基于 `tidwall/gjson` + `tidwall/sjson`。

| 函数 | 说明 |
|------|------|
| `JSON_Get(jsonText, path string) string` | 按路径获取值 |
| `JSON_GetInt/GetFloat/GetBool(jsonText, path string)` | 按路径获取类型化值 |
| `JSON_GetArray(jsonText, path string) []gjson.Result` | 获取数组 |
| `JSON_GetMap(jsonText, path string) map[string]gjson.Result` | 获取映射 |
| `JSON_Exists(jsonText, path string) bool` | 判断路径是否存在 |
| `JSON_ArrayLen(jsonText, path string) int` | 获取数组长度 |
| `JSON_Set(jsonText, path string, value interface{}) (string, error)` | 设置值 |
| `JSON_Delete(jsonText, path string) (string, error)` | 删除值 |
| `JSON_ToString/JSON_ToIndent(v interface{}) (string, error)` | 序列化 |

---

## 22. 协程池 (goroutine_pool.go)

基于 `panjf2000/ants/v2`。

| 函数 | 说明 |
|------|------|
| `GoroutinePool_New(size int) (*ants.Pool, error)` | 创建协程池 |
| `GoroutinePool_Submit(pool *ants.Pool, task func()) error` | 提交任务 |
| `GoroutinePool_Running/Free/Cap/Waiting(pool *ants.Pool) int` | 协程池状态 |
| `GoroutinePool_Release(pool *ants.Pool)` | 释放协程池 |
| `GoroutinePool_Tune(pool *ants.Pool, newSize int)` | 调整池大小 |
| `GoroutinePool_NewPreAlloc(size int) (*ants.Pool, error)` | 创建预分配池 |

---

## 23. 定时任务 (cron.go)

基于 `robfig/cron/v3`。

| 函数 | 说明 |
|------|------|
| `Cron_New() *cron.Cron` | 创建定时器 |
| `Cron_AddFunc(c *cron.Cron, spec string, cmd func()) (cron.EntryID, error)` | 添加任务 |
| `Cron_Remove(c *cron.Cron, id cron.EntryID)` | 移除任务 |
| `Cron_Start/Stop(c *cron.Cron)` | 启动/停止 |
| `Cron_Entries(c *cron.Cron) []cron.Entry` | 列出所有任务 |
| `Cron_Run(spec string, cmd func()) (*cron.Cron, error)` | 快速启动 |

---

## 24. 消息总线 (message_bus.go)

基于 `vardius/message-bus`。

| 函数 | 说明 |
|------|------|
| `MessageBus_New(maxSubscribers int) messagebus.MessageBus` | 创建消息总线 |
| `MessageBus_Publish(bus, topic string, msg ...interface{})` | 发布消息 |
| `MessageBus_Subscribe(bus, topic string, fn interface{}) error` | 订阅消息 |
| `MessageBus_Unsubscribe(bus, topic string, fn interface{}) error` | 取消订阅 |
| `MessageBus_CloseTopic(bus, topic string)` | 关闭主题 |

---

## 25. 字节缓冲池 (bytebuffer_pool.go)

基于 `valyala/bytebufferpool`。

| 函数 | 说明 |
|------|------|
| `ByteBuffer_Get() *bytebufferpool.ByteBuffer` | 获取缓冲区 |
| `ByteBuffer_Put(buf *bytebufferpool.ByteBuffer)` | 归还缓冲区 |
| `ByteBuffer_GetBytes(data []byte) []byte` | 池化字节复制 |
| `ByteBuffer_GetString(s string) string` | 池化字符串复制 |

---

## 26. 结构体合并 (struct_merge.go)

基于 `mergo` + `copier`。

| 函数 | 说明 |
|------|------|
| `Struct_Merge(dst, src interface{}) error` | 合并（跳过空值） |
| `Struct_MergeOverride(dst, src interface{}) error` | 合并（覆盖） |
| `Struct_MergeMap/MergeMapOverride(dst, src interface{}) error` | map 到结构体合并 |
| `Struct_Copy(dst, src interface{}) error` | 深拷贝 |
| `Struct_CopyWithOption(dst, src interface{}, ignoreEmpty bool) error` | 带选项拷贝 |

---

## 27. 表达式求值 (expression.go)

基于 `Knetic/govaluate`。

| 函数 | 说明 |
|------|------|
| `Expression_Eval(expr string) (interface{}, error)` | 求值 |
| `Expression_EvalWithParams(expr string, params map[string]interface{}) (interface{}, error)` | 带参数求值 |
| `Expression_Compile(expr string) (*govaluate.EvaluableExpression, error)` | 编译表达式 |
| `Expression_CompileAndEval(expr string) (interface{}, error)` | 编译并求值 |
| `Expression_Run(e *govaluate.EvaluableExpression) (interface{}, error)` | 运行已编译表达式 |
| `Expression_RunWithParams(e *govaluate.EvaluableExpression, params map[string]interface{}) (interface{}, error)` | 带参数运行 |

---

## 28. 日期解析 (date_parse.go)

基于 `araddon/dateparse`。

| 函数 | 说明 |
|------|------|
| `DateParse_Any(dateText string) (time.Time, error)` | 解析任意格式日期 |
| `DateParse_Local(dateText string) (time.Time, error)` | 解析为本地时间 |
| `DateParse_Strict(dateText string) (time.Time, error)` | 严格解析 |
| `DateParse_Format(dateText string) string` | 解析并格式化为 "2006-01-02 15:04:05" |
| `DateParse_ToString/ToTimestamp/ToMillis/GetDate/GetTime` | 格式化辅助方法 |

---

## 29. 文件监控 (file_watcher.go)

基于 `fsnotify`。

| 函数 | 说明 |
|------|------|
| `FileWatcher_New() (*fsnotify.Watcher, error)` | 创建监控器 |
| `FileWatcher_AddDir(w, dirPath string) error` | 监控目录 |
| `FileWatcher_AddFile(w, filePath string) error` | 监控文件 |
| `FileWatcher_Remove(w, path string) error` | 取消监控 |
| `FileWatcher_Events/Errors(w) <-chan` | 事件/错误通道 |
| `FileWatcher_Start(w, fn func(event fsnotify.Event))` | 启动并设置回调 |
| `FileWatcher_WatchDir(dirPath string, fn func(event fsnotify.Event)) (func(), error)` | 快速监控目录 |

---

## 30. 模板处理 (template.go)

基于 `valyala/fasttemplate`。

| 函数 | 说明 |
|------|------|
| `Template_Execute(tpl string, tags map[string]interface{}) (string, error)` | 执行 {{key}} 模板 |
| `Template_ExecuteDelim(tpl, startTag, endTag string, tags map[string]interface{}) (string, error)` | 自定义分隔符 |
| `Template_Write(tpl string, w io.Writer, tags map[string]interface{}) (int64, error)` | 写入 writer |
| `Template_New(tpl, startTag, endTag string) (*fasttemplate.Template, error)` | 创建可复用模板 |
| `Template_Replace(tpl string, tags map[string]interface{}) string` | 简单字符串替换 |

---

## 31. 系统信息 (system_info.go)

基于 `shirou/gopsutil/v3`。

| 函数 | 说明 |
|------|------|
| `SystemInfo_CPUInfo() ([]cpu.InfoStat, error)` | CPU 详细信息 |
| `SystemInfo_CPULogicalCount() (int, error)` | 逻辑核心数 |
| `SystemInfo_CPUPhysCount() (int, error)` | 物理核心数 |
| `SystemInfo_CPUUsage(interval int) ([]float64, error)` | 每核心使用率 % |
| `SystemInfo_CPUTotalUsage(interval int) (float64, error)` | 总 CPU 使用率 % |
| `SystemInfo_MemInfo() (*mem.VirtualMemoryStat, error)` | 内存信息 |
| `SystemInfo_SwapInfo() (*mem.SwapMemoryStat, error)` | 交换内存信息 |
| `SystemInfo_DiskPartitions() ([]disk.PartitionStat, error)` | 磁盘分区列表 |
| `SystemInfo_DiskUsage(path string) (*disk.UsageStat, error)` | 磁盘使用情况 |
| `SystemInfo_DiskIO() (map[string]disk.IOCountersStat, error)` | 磁盘 I/O |
| `SystemInfo_HostInfo() (*host.InfoStat, error)` | 主机信息 |
| `SystemInfo_NetInterfaces() ([]net.InterfaceStat, error)` | 网络接口列表 |
| `SystemInfo_NetConnections() ([]net.ConnectionStat, error)` | 网络连接列表 |
| `SystemInfo_NetIO() ([]net.IOCountersStat, error)` | 网络 I/O |
| `SystemInfo_Uptime() (uint64, error)` | 运行时长（秒） |
| `SystemInfo_LoadAvg() (*load.AvgStat, error)` | 系统负载 |
| `SystemInfo_Processes() ([]*process.Process, error)` | 所有进程 |
| `SystemInfo_ProcessByPID(pid int32) (*process.Process, error)` | 按 PID 获取进程 |

---

## 32. 控制台表格 (table.go)

基于 `termtables`。

| 函数 | 说明 |
|------|------|
| `Table_New() *termtables.Table` | 创建表格 |
| `Table_AddHeaders/AddRow/AddSeparator` | 构建表格 |
| `Table_Render(t) string` | 渲染表格 |
| `Table_Quick(headers []string, rows [][]string) string` | 快速渲染 |
| `Table_ToMarkdown/ToCSV/ToTSV/ToJSON/ToHTML` | 导出各种格式 |

---

## 33. 键值数据库 (buntdb.go)

基于 `tidwall/buntdb`。

| 函数 | 说明 |
|------|------|
| `BuntDB_Open(filePath string) (*buntdb.DB, error)` | 打开/创建数据库 |
| `BuntDB_Get(db, key string) (string, error)` | 获取值 |
| `BuntDB_Set(db, key, value string) error` | 设置值 |
| `BuntDB_SetWithTTL(db, key, value string, expireSeconds float64) error` | 设置值并设置过期时间 |
| `BuntDB_Delete(db, key string) error` | 删除键 |
| `BuntDB_Scan(db, startKey, endKey string, fn func(k,v string) bool) error` | 范围扫描 |
| `BuntDB_CreateIndex(db, indexName, pattern string) error` | 创建空间索引 |
| `BuntDB_Close(db *buntdb.DB) error` | 关闭数据库 |

---

## 34. 邮件发送 (email.go)

基于 `jordan-wright/email`。

| 函数 | 说明 |
|------|------|
| `Email_Send(server, from, password, to, subject, textBody, htmlBody string, attachments []string) error` | SMTP 发送邮件 |
| `Email_SendTLS(...)` | TLS 方式发送 |
| `Email_SendSimple(server, from, password, to, subject, textBody string) error` | 简单文本邮件 |

---

## 35. 图片处理 (image.go)

基于 `disintegration/imaging` + `skip2/go-qrcode`。

### 35.1 输入输出

`Image_Read`、`Image_ReadBase64`、`Image_FromBytes`、`Image_ReadFromReader`、`Image_Save`、`Image_SavePNG`、`Image_SaveJPEG`、`Image_SaveGIF`

### 35.2 信息

`Image_Width`、`Image_Height`、`Image_Size`、`Image_Bounds`、`Image_GetPixel`、`Image_GetPixelRGBA`

### 35.3 编码

`Image_ToBase64`、`Image_ToDataURI`、`Image_ToBytes`

### 35.4 变换

`Image_Resize`、`Image_ResizeToWidth`、`Image_ResizeToHeight`、`Image_Thumbnail`、`Image_Crop`、`Image_CropCenter`、`Image_Rotate`、`Image_Rotate90/180/270`、`Image_FlipH`、`Image_FlipV`

### 35.5 滤镜

`Image_Grayscale`、`Image_Invert`、`Image_AdjustBrightness`、`Image_AdjustContrast`、`Image_AdjustSaturation`、`Image_Blur`、`Image_Sharpen`、`Image_AdjustGamma`、`Image_AdjustHue`

### 35.6 合成

`Image_Watermark(bg, wm image.Image, pos string, ox, oy int, opacity float64) image.Image`、`Image_Overlay`、`Image_OverlayWithOpacity`、`Image_ConcatHorizontal`、`Image_ConcatVertical`、`Image_NewSolid`、`Image_NewTransparent`、`Image_SetOpacity`

### 35.7 二维码

`Image_QRCodeWrite`、`Image_QRCodeBase64`、`Image_QRCodeCustom`、`Image_QRCodeToWriter`

---

## 36. 网页工具 (web_utils.go)

综合网页访问工具，支持 HTTP 请求、Cookie 管理、域名提取等。

| 函数 | 签名 | 说明 |
|------|------|------|
| `WebUtils_GetDomain` | `func WebUtils_GetDomain(rawUrl string) string` | 从 URL 中提取域名 |
| `WebUtils_Request` | `func WebUtils_Request(url string, method int, postData string, cookies string, respCookies *string, headers string, respContentType *string, respStatusCode *int, noRedirect bool, postBytes []byte, proxy string, timeout int, proxyUser string, proxyPass string, proxyType int, inherit interface{}, autoMergeCookie bool, autoCompleteHeaders bool, normalizeHeaders bool) []byte` | 发送 HTTP 请求并返回响应数据 |
| `WebUtils_MergeCookies` | `func WebUtils_MergeCookies(oldCookie, newCookie string) string` | 合并 Cookie 字符串 |
| `WebUtils_GetCookie` | `func WebUtils_GetCookie(cookies, name string) string` | 从 Cookie 字符串中取指定值 |
| `WebUtils_SetCookie` | `func WebUtils_SetCookie(cookies, name, value string) string` | 设置 Cookie 字符串中的指定值 |

---

## 37. HTTP 高级客户端 (http_client.go)

基于 `go-resty/resty/v2`，提供更简洁的 HTTP 客户端封装。

| 函数 | 签名 | 说明 |
|------|------|------|
| `HTTPClient_Get` | `func HTTPClient_Get(url string) (*resty.Response, error)` | 发送 GET 请求 |
| `HTTPClient_GetWithHeaders` | `func HTTPClient_GetWithHeaders(url string, headers map[string]string) (*resty.Response, error)` | 带请求头 GET 请求 |
| `HTTPClient_Post` | `func HTTPClient_Post(url string, data interface{}) (*resty.Response, error)` | 发送 POST 请求（JSON） |
| `HTTPClient_Put` | `func HTTPClient_Put(url string, data interface{}) (*resty.Response, error)` | 发送 PUT 请求 |
| `HTTPClient_Delete` | `func HTTPClient_Delete(url string) (*resty.Response, error)` | 发送 DELETE 请求 |
| `HTTPClient_DownloadFile` | `func HTTPClient_DownloadFile(url string, savePath string) error` | 下载文件到指定路径 |
| `HTTPClient_SetProxy` | `func HTTPClient_SetProxy(client *resty.Client, proxyURL string)` | 设置代理 |
| `HTTPClient_SetTimeout` | `func HTTPClient_SetTimeout(client *resty.Client, timeout time.Duration)` | 设置超时 |
| `HTTPClient_SetHeader` | `func HTTPClient_SetHeader(client *resty.Client, key, value string)` | 设置通用请求头 |

---

## 38. 网页爬虫 (spider.go)

基于 `gocolly/colly/v2`，提供网页爬虫框架封装。

| 函数 | 签名 | 说明 |
|------|------|------|
| `Spider_New` | `func Spider_New() *colly.Collector` | 创建新的采集器实例 |
| `Spider_OnHTML` | `func Spider_OnHTML(collector *colly.Collector, selector string, callback func(e *colly.HTMLElement))` | 注册 HTML 元素选择器回调 |
| `Spider_OnError` | `func Spider_OnError(collector *colly.Collector, callback func(r *colly.Response, err error))` | 注册错误处理回调 |
| `Spider_OnResponse` | `func Spider_OnResponse(collector *colly.Collector, callback func(r *colly.Response))` | 注册响应处理回调 |
| `Spider_Visit` | `func Spider_Visit(collector *colly.Collector, url string) error` | 访问指定 URL |
| `Spider_SetProxy` | `func Spider_SetProxy(collector *colly.Collector, proxyURL string) error` | 设置代理 |
| `Spider_SetTimeout` | `func Spider_SetTimeout(collector *colly.Collector, timeout time.Duration)` | 设置超时 |
| `Spider_Limit` | `func Spider_Limit(collector *colly.Collector, parallelism int, delay time.Duration)` | 设置并发与延迟限制 |

---

## 39. 权限管理 (permission.go)

基于 `casbin/casbin/v2`，支持 RBAC/ABAC 权限模型。

| 函数 | 签名 | 说明 |
|------|------|------|
| `Permission_CreateMemory` | `func Permission_CreateMemory(modelText string) (*casbin.Enforcer, error)` | 基于内存创建权限管理器 |
| `Permission_CreateWithAdapter` | `func Permission_CreateWithAdapter(modelPath string, adapter persist.Adapter) (*casbin.Enforcer, error)` | 使用适配器创建权限管理器 |
| `Permission_Enforce` | `func Permission_Enforce(enforcer *casbin.Enforcer, sub string, obj string, act string) (bool, error)` | 检查权限 |
| `Permission_AddPolicy` | `func Permission_AddPolicy(enforcer *casbin.Enforcer, params ...string) (bool, error)` | 添加策略 |
| `Permission_RemovePolicy` | `func Permission_RemovePolicy(enforcer *casbin.Enforcer, params ...string) (bool, error)` | 移除策略 |
| `Permission_AddRoleForUser` | `func Permission_AddRoleForUser(enforcer *casbin.Enforcer, user string, role string) (bool, error)` | 为用户分配角色 |
| `Permission_GetRolesForUser` | `func Permission_GetRolesForUser(enforcer *casbin.Enforcer, user string) ([]string, error)` | 获取用户所有角色 |
| `Permission_SavePolicy` | `func Permission_SavePolicy(enforcer *casbin.Enforcer) error` | 保存策略到持久层 |

---

## 40. 数据库 (database.go)

基于 `xorm.io/xorm`，提供 ORM 数据库操作封装。

| 函数 | 签名 | 说明 |
|------|------|------|
| `Database_ConnectMySQL` | `func Database_ConnectMySQL(user string, password string, host string, dbName string) (*xorm.Engine, error)` | 连接 MySQL 数据库 |
| `Database_ConnectSQLite` | `func Database_ConnectSQLite(dbPath string) (*xorm.Engine, error)` | 连接 SQLite 数据库 |
| `Database_ConnectPostgres` | `func Database_ConnectPostgres(connStr string) (*xorm.Engine, error)` | 连接 PostgreSQL 数据库 |
| `Database_Close` | `func Database_Close(engine *xorm.Engine)` | 关闭数据库连接 |
| `Database_Ping` | `func Database_Ping(engine *xorm.Engine) error` | 测试数据库连接 |
| `Database_SyncTables` | `func Database_SyncTables(engine *xorm.Engine, beans ...interface{}) error` | 同步表结构 |
| `Database_Insert` | `func Database_Insert(engine *xorm.Engine, beans ...interface{}) (int64, error)` | 插入记录 |
| `Database_Find` | `func Database_Find(engine *xorm.Engine, beans interface{}, condiBeans ...interface{}) error` | 查询记录 |
| `Database_Update` | `func Database_Update(engine *xorm.Engine, bean interface{}, condiBeans ...interface{}) (int64, error)` | 更新记录 |
| `Database_Delete` | `func Database_Delete(engine *xorm.Engine, bean interface{}) (int64, error)` | 删除记录 |
| `Database_Transaction` | `func Database_Transaction(engine *xorm.Engine, fn func(session *xorm.Session) (interface{}, error)) (interface{}, error)` | 执行事务 |

---

## 41. Windows：进程 (process_windows.go)

> **编译标签**：`windows` | 使用 `kernel32.dll`

| 函数 | 说明 |
|------|------|
| `Process_Create(programPath, cmdLine, workDir string) (*PROCESS_INFORMATION, error)` | 创建进程 |
| `Process_Open(processID, access uint32) (syscall.Handle, error)` | 按 PID 打开进程 |
| `Process_Terminate(processID, exitCode uint32) error` | 终止进程 |
| `Process_IsAlive(processID uint32) bool` | 判断进程是否存活 |
| `Process_Wait/CloseHandle/GetID/GetExitCode/SetPriority/GetPriority` | 句柄操作 |
| `Process_Enum() ([]PROCESSENTRY32W, error)` | 枚举所有进程 |
| `Process_FindByName(name string) ([]uint32, error)` | 按名称查找进程 |
| `Process_GetModulePath(processID uint32) (string, error)` | 获取可执行文件路径 |
| `Process_GetParentID(processID uint32) (uint32, error)` | 获取父进程 PID |

---

## 42. Windows：系统命令 (system_command_windows.go)

> **编译标签**：`windows` | 使用 `user32.dll`、`kernel32.dll`、`advapi32.dll`

| 函数 | 说明 |
|------|------|
| `SysCmd_Shutdown(force bool) error` | 关机 |
| `SysCmd_Reboot(force bool) error` | 重启 |
| `SysCmd_Logoff(force bool) error` | 注销 |
| `SysCmd_LockWorkstation() error` | 锁定工作站 |
| `SysCmd_RemoteShutdown(computerName, message string, timeoutSec uint32, force bool) error` | 远程关机 |
| `SysCmd_SetClipboard(text string) error` | 设置剪贴板 |
| `SysCmd_GetClipboard() (string, error)` | 获取剪贴板内容 |
| `SysCmd_ClearClipboard() error` | 清空剪贴板 |
| `SysCmd_MessageBox(title, text string)` | 弹出消息框（确定） |
| `SysCmd_MessageBoxConfirm(title, text string) int` | 弹出确认框（确定/取消） |
| `SysCmd_MessageBoxYesNo(title, text string) int` | 弹出询问框（是/否） |
| `SysCmd_Exec(command string, args ...string) (string, error)` | 执行命令行 |
| `SysCmd_ExecHidden(command string, args ...string) error` | 隐藏执行命令行 |
| `SysCmd_GetComputerName() (string, error)` | 获取计算机名 |
| `SysCmd_GetUserName() (string, error)` | 获取用户名 |
| `SysCmd_DisableScreenSaver(disable bool) error` | 禁用/启用屏保 |
| `SysCmd_SetScreenSaverTimeout(seconds int) error` | 设置屏保超时 |
| `SysCmd_SetEnv/GetEnv/UnsetEnv(name, value string)` | 环境变量操作 |

---

## 43. Windows：线程 (thread_windows.go)

> **编译标签**：`windows` | 使用 `kernel32.dll`

| 分类 | 函数 |
|------|------|
| 线程 | `Thread_Create`、`Thread_CurrentID`、`Thread_CurrentPseudoHandle`、`Thread_Suspend`、`Thread_Resume`、`Thread_Terminate`、`Thread_CloseHandle`、`Thread_Sleep`、`Thread_WaitSingle`、`Thread_WaitMultiple` |
| 临界区 | `Thread_CriticalSection_Create`、`_Delete`、`_Enter`、`_Leave` |
| 事件 | `Thread_Event_Create`、`_Set`、`_Reset`、`_Pulse` |
| 互斥体 | `Thread_Mutex_Create`、`_Open`、`_Release` |
| 信号量 | `Thread_Semaphore_Create`、`_Open`、`_Release` |

---

## 44. Windows：键盘鼠标 (input_windows.go)

> **编译标签**：`windows` | 使用 `user32.dll`

| 函数 | 说明 |
|------|------|
| `Input_Key(vk int, down bool)` | 按下/释放按键 |
| `Input_KeyCombo(vk1, vk2 int)` | 组合键（如 Ctrl+C） |
| `Input_KeyPress(vk int)` | 完整按下+释放 |
| `Input_GetKeyState(vk int) int16` | 异步按键状态 |
| `Input_IsKeyDown(vk int) bool` | 判断按键是否按下 |
| `Input_GetKeyToggle(vk int) int16` | 切换键状态 |
| `Input_TypeText(text string)` | 输入文本 |
| `Input_MoveMouse(x, y int)` | 移动鼠标 |
| `Input_GetMousePos() (int, int)` | 获取鼠标位置 |
| `Input_LeftClick/RightClick/MiddleClick()` | 鼠标点击 |
| `Input_LeftDown/LeftUp()` | 鼠标按下/释放 |
| `Input_MouseWheel(delta int)` | 滚轮滚动 |
| `Input_ScreenWidth/ScreenHeight() int` | 屏幕尺寸 |
| `Input_Block(block bool)` | 锁定/解锁输入 |

**虚拟键常量**：`VK_A`..`VK_Z`、`VK_0`..`VK_9`、`VK_F1`..`VK_F12`、`VK_SHIFT`、`VK_CONTROL`、`VK_MENU`、`VK_RETURN`、`VK_SPACE`、`VK_ESCAPE`、`VK_UP/DOWN/LEFT/RIGHT`、`VK_HOME/VK_END`、`VK_LBUTTON/VK_RBUTTON` 等。

---

## 45. Windows：窗口 (window_windows.go)

> **编译标签**：`windows` | 使用 `user32.dll`

| 函数 | 签名 | 说明 |
|------|------|------|
| `Window_Find` | `func Window_Find(className string, title string) syscall.Handle` | 按类名和标题查找顶层窗口 |
| `Window_FindEx` | `func Window_FindEx(parent syscall.Handle, className string, title string) syscall.Handle` | 查找子窗口 |
| `Window_GetTitle` | `func Window_GetTitle(hWnd syscall.Handle) string` | 获取窗口标题 |
| `Window_SetTitle` | `func Window_SetTitle(hWnd syscall.Handle, title string) bool` | 设置窗口标题 |
| `Window_GetClassName` | `func Window_GetClassName(hWnd syscall.Handle) string` | 获取窗口类名 |
| `Window_GetRect` | `func Window_GetRect(hWnd syscall.Handle) (RECT, bool)` | 获取窗口矩形区域 |
| `Window_Show` | `func Window_Show(hWnd syscall.Handle, cmdShow int) bool` | 显示/隐藏窗口 |
| `Window_Close` | `func Window_Close(hWnd syscall.Handle) bool` | 关闭窗口 |
| `Window_SendMessage` | `func Window_SendMessage(hWnd syscall.Handle, msg uint32, wParam, lParam uintptr) uintptr` | 发送消息到窗口 |
| `Window_GetDesktop` | `func Window_GetDesktop() syscall.Handle` | 获取桌面窗口句柄 |
| `Window_GetForeground` | `func Window_GetForeground() syscall.Handle` | 获取前台窗口句柄 |
| `Window_SetForeground` | `func Window_SetForeground(hWnd syscall.Handle) bool` | 设置前台窗口 |

---

## 46. Windows：内存操作 (memory_windows.go)

> **编译标签**：`windows` | 使用 `kernel32.dll`

| 函数 | 签名 | 说明 |
|------|------|------|
| `Memory_OpenProcess` | `func Memory_OpenProcess(pid uint32, access uint32) (syscall.Handle, error)` | 打开进程以进行内存操作 |
| `Memory_CloseHandle` | `func Memory_CloseHandle(h syscall.Handle) error` | 关闭句柄 |
| `Memory_ReadBytes` | `func Memory_ReadBytes(hProcess syscall.Handle, address uintptr, size uintptr) ([]byte, error)` | 读取进程内存 |
| `Memory_WriteBytes` | `func Memory_WriteBytes(hProcess syscall.Handle, address uintptr, data []byte) error` | 写入进程内存 |
| `Memory_ReadInt32` | `func Memory_ReadInt32(hProcess syscall.Handle, address uintptr) (int32, error)` | 读取 32 位整数 |
| `Memory_ReadInt64` | `func Memory_ReadInt64(hProcess syscall.Handle, address uintptr) (int64, error)` | 读取 64 位整数 |
| `Memory_ReadFloat32` | `func Memory_ReadFloat32(hProcess syscall.Handle, address uintptr) (float32, error)` | 读取 32 位浮点 |
| `Memory_ReadFloat64` | `func Memory_ReadFloat64(hProcess syscall.Handle, address uintptr) (float64, error)` | 读取 64 位浮点 |
| `Memory_ReadString` | `func Memory_ReadString(hProcess syscall.Handle, address uintptr, maxLen int) (string, error)` | 读取字符串 |
| `Memory_WriteInt32` | `func Memory_WriteInt32(hProcess syscall.Handle, address uintptr, value int32) error` | 写入 32 位整数 |
| `Memory_WriteFloat64` | `func Memory_WriteFloat64(hProcess syscall.Handle, address uintptr, value float64) error` | 写入 64 位浮点 |
| `Memory_SearchBytes` | `func Memory_SearchBytes(hProcess syscall.Handle, pattern []byte, startAddr uintptr, length uintptr) ([]uintptr, error)` | 搜索字节特征码（AOB 搜索） |
| `Memory_GetModuleBase` | `func Memory_GetModuleBase(pid uint32, moduleName string) (uintptr, error)` | 获取模块基址 |
| `Memory_GetProcessID` | `func Memory_GetProcessID(processName string) (uint32, error)` | 按进程名获取 PID |
| `Memory_EnumProcesses` | `func Memory_EnumProcesses() ([]PROCESSENTRY32W, error)` | 枚举所有进程（用于内存操作） |

---

## 47. Windows：API Hook (hook.go)

> **编译标签**：`windows` | 基于 MinHook 库

| 函数 | 签名 | 说明 |
|------|------|------|
| `Hook_Init` | `func Hook_Init() error` | 初始化 Hook 引擎 |
| `Hook_Uninit` | `func Hook_Uninit() error` | 反初始化 Hook 引擎 |
| `Hook_Install` | `func Hook_Install(target, callback uintptr) (uintptr, error)` | 安装钩子（通过地址） |
| `Hook_InstallApi` | `func Hook_InstallApi(moduleName, funcName string, callbackAddr uintptr) (uintptr, error)` | 通过模块名和函数名安装钩子 |
| `Hook_Remove` | `func Hook_Remove(target uintptr) error` | 移除钩子 |
| `Hook_Enable` | `func Hook_Enable(target uintptr) error` | 启用指定钩子 |
| `Hook_Disable` | `func Hook_Disable(target uintptr) error` | 禁用指定钩子 |
| `Hook_EnableAll` | `func Hook_EnableAll() error` | 启用所有钩子 |
| `Hook_DisableAll` | `func Hook_DisableAll() error` | 禁用所有钩子 |
| `Hook_QueueEnable` | `func Hook_QueueEnable(target uintptr) error` | 排队启用钩子 |
| `Hook_ApplyQueued` | `func Hook_ApplyQueued() error` | 应用排队操作 |

---

## 48. Windows：汇编引擎 (asm_assembler.go)

> **编译标签**：`windows` | 运行时机器码生成与执行

| 函数 | 签名 | 说明 |
|------|------|------|
| `ASM_Clear` | `func ASM_Clear()` | 清空汇编代码缓冲区 |
| `ASM_Execute` | `func ASM_Execute() (uintptr, error)` | 在当前进程中执行汇编代码 |
| `ASM_ExecuteRemote` | `func ASM_ExecuteRemote(hProcess syscall.Handle) (uintptr, error)` | 在远程进程中执行汇编代码 |
| `ASM_SetRegister` | `func ASM_SetRegister(reg string, value uintptr)` | 设置寄存器值（用于调用约定） |

---

## 49. Windows：汇编指令 (asm_instructions.go)

> **编译标签**：`windows` | x86/x64 汇编指令集

### 数据传输

| 函数 | 说明 |
|------|------|
| `ASM_MOV_EAX(v int32)` | MOV EAX, imm32 |
| `ASM_MOV_ECX(v int32)` | MOV ECX, imm32 |
| `ASM_MOV_EDX(v int32)` | MOV EDX, imm32 |
| `ASM_MOV_EBX(v int32)` | MOV EBX, imm32 |
| `ASM_PUSH(v int32)` | PUSH imm32 |
| `ASM_PUSH_EAX()` | PUSH EAX |
| `ASM_PUSH_ECX()` | PUSH ECX |
| `ASM_POP_EAX()` | POP EAX |

### 算术运算

| 函数 | 说明 |
|------|------|
| `ASM_ADD_EAX(v int32)` | ADD EAX, imm32 |
| `ASM_SUB_EAX(v int32)` | SUB EAX, imm32 |

### 控制流

| 函数 | 说明 |
|------|------|
| `ASM_CALL(addr int32)` | CALL rel32 |
| `ASM_JMP(addr int32)` | JMP rel32 |
| `ASM_RET()` | RET |
| `ASM_NOP()` | NOP |

### 比较与条件跳转

| 函数 | 说明 |
|------|------|
| `ASM_CMP_EAX(v int32)` | CMP EAX, imm32 |
| `ASM_JE(addr int32)` | JE rel32 |
| `ASM_JNE(addr int32)` | JNE rel32 |

---

## 50. OpenCV 视觉 (opencv.go)

> **编译标签**：`opencv` | 基于 `gocv.io/x/gocv`

### 50.1 核心

`OCV_Version() string`、`OCV_CUDADeviceCount() int`

### 50.2 图像 I/O

`OCV_IMRead`、`OCV_IMReadGray`、`OCV_IMReadUnchanged`、`OCV_IMWrite`、`OCV_IMDecode`、`OCV_IMDecodeGray`、`OCV_IMEncode`

### 50.3 Mat 信息

`OCV_Width`、`OCV_Height`、`OCV_Channels`、`OCV_Type`、`OCV_Total`、`OCV_IsEmpty`、`OCV_Size`

### 50.4 颜色空间

`OCV_BGRToGray`、`OCV_GrayToBGR`、`OCV_BGRToHSV`、`OCV_HSVToBGR`、`OCV_BGRToRGB`、`OCV_RGBToBGR`、`OCV_BGRToLab`、`OCV_BGRToYUV`

### 50.5 变换

`OCV_Resize`、`OCV_ResizeByRatio`、`OCV_Crop`、`OCV_Rotate`、`OCV_FlipH`、`OCV_FlipV`、`OCV_AffineTransform`、`OCV_PerspectiveTransform`

### 50.6 滤波

`OCV_GaussianBlur`、`OCV_MedianBlur`、`OCV_BilateralFilter`、`OCV_Blur`

### 50.7 形态学

`OCV_Erode`、`OCV_Dilate`、`OCV_MorphOpen`、`OCV_MorphClose`、`OCV_MorphGradient`、`OCV_MorphTopHat`、`OCV_MorphBlackHat`

### 50.8 边缘检测

`OCV_Canny`、`OCV_Sobel`、`OCV_Laplacian`、`OCV_Scharr`

### 50.9 阈值

`OCV_Threshold`、`OCV_AdaptiveThreshold`、`OCV_ThresholdInv`、`OCV_ThresholdOtsu`

### 50.10 轮廓

`OCV_FindContours`、`OCV_FindAllContours`、`OCV_DrawContours`、`OCV_ContourArea`、`OCV_ArcLength`、`OCV_BoundingRect`、`OCV_MinEnclosingCircle`

### 50.11 模板匹配

`OCV_MatchTemplate`、`OCV_MatchTemplateAll`

### 50.12 人脸检测

`OCV_FaceDetect`、`OCV_FaceDetectWithParams`