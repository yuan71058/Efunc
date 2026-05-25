# Efunc 英文版工具库

Efunc 是一个综合性的 Go 语言工具库，涵盖数据类型转换、文本处理、加解密、网络通信、并发编程、图片处理、系统操作等功能。

## 安装

```bash
go get github.com/yuan71058/Efunc/en
```

## 运行 Demo

`demo/` 目录下有 9 个独立示例，分别演示各模块用法：

```bash
# 基础类型转换、文本处理、编码、校验和
go run ./demo/01_basic

# AES/DES/3DES/RSA/HMAC 加解密
go run ./demo/02_crypto

# 协程池、消息总线、定时任务、缓冲区池
go run ./demo/03_concurrency

# 文件读写、目录操作、日志
go run ./demo/04_file_system

# HTTP 客户端、URL 解析、Cookie 管理、IP 工具
go run ./demo/05_web_network

# 数据库 ORM（SQLite）、BuntDB 键值存储
go run ./demo/06_database

# 数据校验、配置管理（Viper）、模板引擎、表达式求值
go run ./demo/07_config_validate

# Windows 平台：窗口管理、进程枚举、剪贴板读写
go run ./demo/08_windows

# RBAC 权限管理（Casbin）
go run ./demo/09_permission
```

## 模块总览

### class/ — 并发数据结构

| 模块 | 源文件 | 说明 |
|------|--------|------|
| Mutex | mutex.go | 互斥锁封装，支持 TryLock |
| RWMutex | rwmutex.go | 读写锁封装，支持 TryRLock/TryLock |
| Queue | queue.go | 队列数据结构 |
| QueueG | queue_generic.go | 泛型队列 |
| Regex | regex.go | 正则表达式封装，带缓存 |
| TCP | tcp.go | TCP 服务端/客户端，事件驱动 |
| WebSocket | ws.go | WebSocket 服务端/客户端，事件驱动 |
| HTTP | http.go | HTTP 服务端，支持路由 |

### utils/ — 工具函数

#### 核心与类型转换
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Core | core.go | `IntToStr`、`FloatToStr`、`BaseToInt`、`StrToInt`、`StrToFloat64` |
| Cast | cast.go | `ToInt`、`ToString`、`ToBool`、`ToFloat64`、`ToTime`、`ToSlice` |
| IntConv | intconv.go | `IntToInt64`、`Int64ToInt`、`UintToStr`、`StrToUint` |
| Float64Conv | float64conv.go | `Float64ToStr`、`Float64ToInt64`、`Float64ToFloat32` |
| ByteData | bytedata.go | `BytesToInt`、`BytesToFloat64`、`IntToBytes`、`Float64ToBytes` |
| Array | array.go | `ArrayContains`、`ArrayRemove`、`ArrayDistinct`、`ArrayJoin` |
| Map | map_utils.go | `MapKeys`、`MapValues`、`MapMerge`、`MapToJSON` |
| Atomic | atomic.go | `AtomicInc`、`AtomicDec`、`AtomicLoad`、`AtomicStore`、`AtomicCAS` |

#### 文本与编码
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Text | text.go | `Text_Between`、`Text_Left`、`Text_Right`、`Text_Replace`、`Text_Split` |
| Encoding | encoding.go | `Encoding_GBKToUTF8`、`Encoding_UTF8ToGBK`、`Encoding_Base64Encode`、`Encoding_URLEncode` |
| Regex | regex_utils.go | `Regex_IsMatch`、`Regex_Find`、`Regex_Replace`、`Regex_FindAll` |

#### 加密与安全
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Crypto | crypto.go | `Crypto_AESEncrypt`、`Crypto_DESEncrypt`、`Crypto_3DESEncrypt`、`Crypto_HMAC` |
| RSA | rsa.go | `RSA_GenerateKey`、`RSA_Encrypt`、`RSA_Decrypt`、`RSA_Sign`、`RSA_Verify` |
| Checksum | checksum.go | `Checksum_MD5`、`Checksum_SHA1`、`Checksum_SHA256`、`Checksum_CRC32` |

#### 网络与 HTTP
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| IP | ip.go | `IP_IsValid`、`IP_LongToString`、`IP_StringToLong`、`IP_IsPrivate` |
| HTTP Utils | http_utils.go | `HTTP_Get`、`HTTP_Post`、`HTTP_PostJSON`、`HTTP_SetTimeout` |
| HTTP Client | http_client.go | `HTTPClient_Get`、`HTTPClient_Post`、`HTTPClient_DownloadFile`、`HTTPClient_SetProxy` |
| Web Utils | web_utils.go | `WebUtils_GetDomain`、`WebUtils_Request`、`WebUtils_MergeCookies`、`WebUtils_GetCookie` |
| Spider | spider.go | `Spider_New`、`Spider_Visit`、`Spider_OnHTML`、`Spider_OnError` |

#### 文件与系统
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| File | file.go | `File_Read`、`File_Write`、`File_Exists`、`File_Copy`、`File_GetSize` |
| FileWatcher | file_watcher.go | `FileWatch_New`、`FileWatch_Add`、`FileWatch_Start`、`FileWatch_Stop` |
| Directory | directory.go | `Dir_Create`、`Dir_Exists`、`Dir_List`、`Dir_RemoveAll` |
| Program | program.go | `Program_GetPath`、`Program_GetDir`、`Program_IsRunning` |

#### 时间
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Time | time.go | `Time_Now`、`Time_Format`、`Time_Parse`、`Time_Unix`、`Time_NTPSync` |
| DateParse | date_parse.go | `DateParse_Any`、`DateParse_Format`、`DateParse_ToTime` |
| Cron | cron.go | `Cron_New`、`Cron_AddJob`、`Cron_Start`、`Cron_Stop` |

#### 数据与配置
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| JSON | json_utils.go | `JSON_Parse`、`JSON_Get`、`JSON_Set`、`JSON_ToStruct` |
| Config | config.go | `Config_Load`、`Config_Get`、`Config_Set`、`Config_Save` |
| Env | env.go | `Env_Get`、`Env_Set`、`Env_Load`、`Env_GetInt` |
| Log | log.go | `Log_Info`、`Log_Error`、`Log_Debug`、`Log_SetLevel`、`Log_ToFile` |
| Validation | validation.go | `Validation_Struct`、`Validation_Var`、`Validation_Register` |
| PostData | postdata.go | `PostData_Parse`、`PostData_Get`、`PostData_Set` |
| Expression | expression.go | `Expression_Eval`、`Expression_EvalWithVars` |
| StructMerge | struct_merge.go | `StructMerge_Merge`、`StructMerge_MergeWithOverride` |

#### 并发
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Goroutine Pool | goroutine_pool.go | `Pool_New`、`Pool_Submit`、`Pool_Release`、`Pool_Running` |
| Message Bus | message_bus.go | `Bus_Publish`、`Bus_Subscribe`、`Bus_Unsubscribe` |
| ByteBuffer Pool | bytebuffer_pool.go | `BufferPool_Get`、`BufferPool_Put` |

#### 媒体与图像
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Image | image.go | `Image_Resize`、`Image_Crop`、`Image_Rotate`、`Image_Watermark`、`Image_QRCode` |
| OpenCV | opencv.go | `OpenCV_ReadImage`、`OpenCV_DetectFaces`、`OpenCV_MatchTemplate` |

#### 数据库
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Database | database.go | `Database_ConnectMySQL`、`Database_Insert`、`Database_Find`、`Database_Update`、`Database_Transaction` |
| BuntDB | buntdb.go | `BuntDB_Open`、`BuntDB_Set`、`BuntDB_Get`、`BuntDB_Delete`、`BuntDB_Iterate` |

#### 系统与平台
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| CLI | cli.go | `CLI_ReadInput`、`CLI_ReadPassword`、`CLI_Confirm` |
| Helper | helper.go | `Helper_Wait`、`Helper_Retry`、`Helper_PanicRecover` |
| Table | table.go | `Table_Print`、`Table_PrintStructs`、`Table_ToJSON` |
| Template | template.go | `Template_Render`、`Template_RenderFile`、`Template_AddFunc` |
| Email | email.go | `Email_Send`、`Email_SendWithAttachment`、`Email_SendHTML` |
| System Info | system_info.go | `SysInfo_CPU`、`SysInfo_Memory`、`SysInfo_Disk`、`SysInfo_Network` |

#### Windows 平台专用（build tag: windows）
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Process | process_windows.go | `Process_Open`、`Process_Enum`、`Process_GetWindow` |
| Thread | thread_windows.go | `Thread_Create`、`Thread_Suspend`、`Thread_Resume`、`Thread_Wait` |
| Input | input_windows.go | `Input_KeyDown`、`Input_KeyUp`、`Input_MouseClick`、`Input_TypeString` |
| System Command | system_command_windows.go | `SysCmd_Shutdown`、`SysCmd_MessageBox`、`SysCmd_GetClipboard`、`SysCmd_SetClipboard` |
| Window | window_windows.go | `Window_Find`、`Window_GetTitle`、`Window_Show`、`Window_SendMessage`、`Window_Close` |
| Memory | memory_windows.go | `Memory_ReadInt32`、`Memory_WriteBytes`、`Memory_SearchBytes`、`Memory_GetModuleBase` |
| Hook | hook.go | `Hook_Install`、`Hook_InstallApi`、`Hook_Enable`、`Hook_Remove` |
| ASM Assembler | asm_assembler.go | `ASM_Execute`、`ASM_ExecuteRemote` |
| ASM Instructions | asm_instructions.go | `ASM_MOV_EAX`、`ASM_ADD_EAX`、`ASM_CALL`、`ASM_JMP`、`ASM_CMP_EAX` |

#### 高级模块
| 模块 | 源文件 | 主要函数 |
|------|--------|----------|
| Permission | permission.go | `Permission_Enforce`、`Permission_AddPolicy`、`Permission_AddRoleForUser` |

## 快速开始

### 类型转换

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    str := utils.IntToStr(12345)           // "12345"
    num := utils.StrToInt("67890")         // 67890
    f := utils.Float64ToStr(3.14159, 2)    // "3.14"

    fmt.Println(str, num, f)
}
```

### HTTP 客户端

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    // GET 请求
    resp, err := utils.HTTPClient_Get("https://api.example.com/data")
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.String())

    // POST JSON
    data := map[string]string{"name": "test", "value": "hello"}
    resp, err = utils.HTTPClient_Post("https://api.example.com/submit", data)
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.StatusCode())
}
```

### 网页爬虫

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    collector := utils.Spider_New()

    utils.Spider_OnHTML(collector, "h1", func(e *colly.HTMLElement) {
        fmt.Println("标题:", e.Text)
    })

    utils.Spider_OnHTML(collector, "a[href]", func(e *colly.HTMLElement) {
        fmt.Println("链接:", e.Attr("href"))
    })

    utils.Spider_Visit(collector, "https://example.com")
}
```

### 数据库操作

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

type User struct {
    Id   int64  `xorm:"pk autoincr"`
    Name string `xorm:"varchar(50)"`
    Age  int    `xorm:"int"`
}

func main() {
    engine, err := utils.Database_ConnectMySQL("root", "password", "127.0.0.1:3306", "test")
    if err != nil {
        panic(err)
    }
    defer utils.Database_Close(engine)

    utils.Database_SyncTables(engine, new(User))

    // 插入
    user := User{Name: "Alice", Age: 25}
    utils.Database_Insert(engine, &user)

    // 查询
    var users []User
    utils.Database_Find(engine, &users)
    fmt.Printf("%+v\n", users)
}
```

### 图片处理

```go
package main

import "github.com/yuan71058/Efunc/en/utils"

func main() {
    // 缩放图片
    utils.Image_Resize("input.jpg", "output.jpg", 800, 600)

    // 生成二维码
    utils.Image_QRCode("https://example.com", "qrcode.png", 256)

    // 添加水印
    utils.Image_Watermark("input.jpg", "watermark.png", "output.jpg")
}
```

### 并发任务

```go
package main

import (
    "fmt"
    "sync"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    pool := utils.Pool_New(10)
    defer utils.Pool_Release(pool)

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        idx := i
        utils.Pool_Submit(pool, func() {
            defer wg.Done()
            fmt.Printf("任务 %d 完成\n", idx)
        })
    }
    wg.Wait()
    fmt.Printf("运行中协程数: %d\n", utils.Pool_Running(pool))
}
```

### 定时任务

```go
package main

import (
    "fmt"
    "time"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    cron := utils.Cron_New()

    utils.Cron_AddJob(cron, "*/5 * * * *", func() {
        fmt.Println("每 5 分钟执行:", time.Now())
    })

    utils.Cron_AddJob(cron, "0 8 * * 1-5", func() {
        fmt.Println("工作日 8 点任务:", time.Now())
    })

    utils.Cron_Start(cron)
    select {}
}
```

### 消息总线

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    utils.Bus_Subscribe("user.created", func(msg interface{}) {
        fmt.Println("用户创建:", msg)
    })

    utils.Bus_Subscribe("user.created", func(msg interface{}) {
        fmt.Println("发送欢迎邮件:", msg)
    })

    utils.Bus_Publish("user.created", "Alice")
}
```

### 加解密

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    key := []byte("1234567890abcdef")
    plaintext := []byte("Hello, World!")

    // AES 加解密
    encrypted, _ := utils.Crypto_AESEncrypt(plaintext, key)
    decrypted, _ := utils.Crypto_AESDecrypt(encrypted, key)
    fmt.Println(string(decrypted)) // "Hello, World!"

    // RSA
    priv, pub, _ := utils.RSA_GenerateKey(2048)
    cipher, _ := utils.RSA_Encrypt(plaintext, pub)
    clear, _ := utils.RSA_Decrypt(cipher, priv)
    fmt.Println(string(clear)) // "Hello, World!"

    // 哈希
    md5 := utils.Checksum_MD5("hello")
    sha256 := utils.Checksum_SHA256("hello")
    fmt.Println(md5, sha256)
}
```

### 权限管理（RBAC）

```go
package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    modelText := `
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
    enforcer, _ := utils.Permission_CreateMemory(modelText)

    utils.Permission_AddPolicy(enforcer, "admin", "data1", "read")
    utils.Permission_AddPolicy(enforcer, "admin", "data1", "write")
    utils.Permission_AddRoleForUser(enforcer, "alice", "admin")

    ok, _ := utils.Permission_Enforce(enforcer, "alice", "data1", "read")
    fmt.Println("Alice 可读 data1:", ok) // true

    ok, _ = utils.Permission_Enforce(enforcer, "bob", "data1", "read")
    fmt.Println("Bob 可读 data1:", ok) // false
}
```

### Windows 平台：进程内存

```go
//go:build windows

package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    pid, _ := utils.Memory_GetProcessID("notepad.exe")
    hProcess, _ := utils.Memory_OpenProcess(pid, utils.PROCESS_VM_READ)
    defer utils.Memory_CloseHandle(hProcess)

    base, _ := utils.Memory_GetModuleBase(pid, "notepad.exe")
    fmt.Printf("基址: 0x%X\n", base)
}
```

### Windows 平台：窗口管理

```go
//go:build windows

package main

import (
    "fmt"
    "github.com/yuan71058/Efunc/en/utils"
)

func main() {
    hwnd := utils.Window_Find("Notepad", "")
    if hwnd != 0 {
        title := utils.Window_GetTitle(hwnd)
        fmt.Println("找到窗口:", title)

        utils.Window_SetTitle(hwnd, "由 Efunc 修改")
        utils.Window_Show(hwnd, utils.SW_MAXIMIZE)
    }
}
```

## 编译约束

部分模块使用 Go build tag 进行条件编译：

- `//go:build windows` — 仅 Windows 平台编译（进程、线程、输入、窗口、内存、Hook、汇编、系统命令）
- `//go:build opencv` — 需要 OpenCV 库支持

非 Windows 平台会自动排除这些文件。

## 环境要求

- Go 1.25+
- Windows API 通过 `golang.org/x/sys/windows`
- 可选：OpenCV 库（供 opencv.go 使用）