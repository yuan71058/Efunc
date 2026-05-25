package utils

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Program_SleepMS 逐毫秒循环延时，精度为 1ms。
// 相比 Program_Sleep，此函数可被中断（如通过 channel），
// 但循环开销较大，一般场景推荐使用 Program_Sleep。
//
// 参数:
//   - ms: 延时的毫秒数
//
// 返回:
//   - bool: 始终返回 true
func Program_SleepMS(ms int) bool {
	for i := 0; i < ms; i++ {
		time.Sleep(1 * time.Millisecond)
	}
	return true
}

// Program_GetCMDPath 尝试获取系统 cmd.exe 的路径。
// 注意：当前实现使用 ReadLink 读取符号链接，可能在某些系统上失败。
//
// 返回:
//   - string: cmd.exe 的路径；失败时程序终止（log.Fatal）
func Program_GetCMDPath() string {
	file, err := os.Readlink("." + "cmd/cmd.exe")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// Program_GetGUID 生成标准的 Version 4 UUID（随机 UUID）。
// 格式如：635897f8-2a48-4882-b3e1-823b8e5b6df8（小写十六进制）。
// 符合 RFC 4122 第 4.4 节规范。
//
// 返回:
//   - string: 8-4-4-4-12 格式的 UUID 字符串
func Program_GetGUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// Program_DeleteSelf 删除当前运行的可执行文件。
// 注意：在 Windows 上，正在运行的程序文件被锁定，删除操作通常会失败。
// 此函数存在自身被占用的问题，建议仅在程序即将退出时调用。
//
// 返回:
//   - error: 成功返回 nil，失败返回错误信息
func Program_DeleteSelf() error {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("获取可执行文件路径失败:", err)
		return errors.New("获取可执行文件路径失败:" + err.Error())
	}

	err = os.Remove(exePath)
	if err != nil {
		return errors.New("删除可执行文件失败:" + err.Error())
	}
	return nil
}

// Program_IsDebugged 检测当前程序是否正在被调试器附加。
// 注意：当前实现仅检查进程 ID 是否非零，实际无法可靠检测调试状态。
// 如需真正的反调试检测，需要使用平台特定 API。
//
// 返回:
//   - bool: 当前实现始终返回 true（pid != 0）
func Program_IsDebugged() bool {
	return os.Getpid() != 0
}

// Program_PreventRerun 通过环境变量防止程序重复运行。
// 如果环境变量 GO_NO_RERUN 已设置，则调用 os.Exit(0) 退出当前进程。
//
// 返回:
//   - bool: 如果环境变量已设置则退出程序；否则返回 false
func Program_PreventRerun() bool {
	if os.Getenv("GO_NO_RERUN") != "" {
		os.Exit(0)
		return true
	}
	return false
}

// Program_WriteLog 将日志内容追加写入到指定路径的日志文件。
// 每条日志自动添加时间戳前缀（格式：2006-01-02 15:04:05）。
// 如果日志文件不存在，会自动创建。日志路径为空时默认写入程序同目录下的"运行日志.txt"。
//
// 参数:
//   - content: 要记录的日志文本
//   - logPath: 日志文件的完整路径，可为空（默认使用程序目录下的"运行日志.txt"）
func Program_WriteLog(content string, logPath string) {
	if logPath == "" {
		exePath, _ := os.Executable()
		logPath = filepath.Dir(exePath) + "/运行日志.txt"
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		file, _ := os.Create(logPath)
		_ = file.Close()
	}

	file, _ := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()

	logLine := fmt.Sprintf("%s   %s\n", time.Now().Format("2006-01-02 15:04:05"), content)
	file.WriteString(logLine)
}

// Program_GetCommandLine 获取启动程序时的命令行参数。
// 返回 os.Args 的完整切片，args[0] 为程序路径，args[1:] 为附加参数。
//
// 返回:
//   - []string: 命令行参数数组
func Program_GetCommandLine() []string {
	args := os.Args
	return args
}

// Program_GetRunDir 获取当前可执行文件所在的目录。
// 会解析符号链接，返回真实的目录路径。
//
// 返回:
//   - string: 可执行文件所在目录的绝对路径；失败时程序终止
func Program_GetRunDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// Program_GetTempDir 获取操作系统的临时目录路径。
// 依次尝试 TEMP、TMP 环境变量，最后回退到 SYSTEMROOT\Temp。
//
// 返回:
//   - string: 临时目录的绝对路径
func Program_GetTempDir() string {
	tempDir := os.Getenv("TEMP")
	if tempDir == "" {
		tempDir = os.Getenv("TMP")
	}

	if tempDir == "" {
		tempDir = filepath.Join(os.Getenv("SYSTEMROOT"), "Temp")
	}

	return tempDir
}

// Program_RunWin 通过 PowerShell 执行命令行指令，并返回输出结果。
// 输出内容会自动从 GBK 编码转换为 UTF-8。
// 注意：此函数会阻塞等待命令执行完毕。
//
// 参数:
//   - command: 要执行的 PowerShell 命令
//
// 返回:
//   - string: 命令执行后的输出文本（UTF-8 编码）
func Program_RunWin(command string) string {
	var err error

	cmd := exec.Command("powershell")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out
	go func(command string) {
		in.WriteString(command)
	}(command)
	err = cmd.Start()

	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	rt := GBKToUTF8(out.String())

	return rt
}

// Program_Sleep 暂停当前 goroutine 指定的毫秒数。
// 使用 time.Sleep 实现，精度取决于操作系统调度器。
//
// 参数:
//   - ms: 延时的毫秒数
//
// 返回:
//   - bool: 始终返回 true
func Program_Sleep(ms int64) bool {
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return true
}