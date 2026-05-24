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

// C程序_延时2 逐毫秒循环延时，精度为 1ms。
// 相比 C程序_延时，此函数可被中断（如通过 channel），
// 但循环开销较大，一般场景推荐使用 C程序_延时。
//
// 参数:
//   - 毫秒: 延时的毫秒数
//
// 返回:
//   - bool: 始终返回 true
func C程序_延时2(毫秒 int) bool {
	for i := 0; i < 毫秒; i++ {
		time.Sleep(1 * time.Millisecond)
	}
	return true
}

// C程序_取cmd路径 尝试获取系统 cmd.exe 的路径。
// 注意：当前实现使用 ReadLink 读取符号链接，可能在某些系统上失败。
//
// 返回:
//   - string: cmd.exe 的路径；失败时程序终止（log.Fatal）
func C程序_取cmd路径() string {
	file, err := os.Readlink("." + "cmd/cmd.exe")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// C程序_取GUID 生成标准的 Version 4 UUID（随机 UUID）。
// 格式如：635897f8-2a48-4882-b3e1-823b8e5b6df8（小写十六进制）。
// 符合 RFC 4122 第 4.4 节规范。
//
// 返回:
//   - string: 8-4-4-4-12 格式的 UUID 字符串
func C程序_取GUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// C程序_删除自身 删除当前运行的可执行文件。
// 注意：在 Windows 上，正在运行的程序文件被锁定，删除操作通常会失败。
// 此函数存在自身被占用的问题，建议仅在程序即将退出时调用。
//
// 返回:
//   - error: 成功返回 nil，失败返回错误信息
func C程序_删除自身() error {
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

// C程序_是否被调试 检测当前程序是否正在被调试器附加。
// 注意：当前实现仅检查进程 ID 是否非零，实际无法可靠检测调试状态。
// 如需真正的反调试检测，需要使用平台特定 API。
//
// 返回:
//   - bool: 当前实现始终返回 true（pid != 0）
func C程序_是否被调试() bool {
	return os.Getpid() != 0
}

// C程序_禁止重复运行 通过环境变量防止程序重复运行。
// 如果环境变量 GO_NO_RERUN 已设置，则调用 os.Exit(0) 退出当前进程。
//
// 返回:
//   - bool: 如果环境变量已设置则退出程序；否则返回 false
func C程序_禁止重复运行() bool {
	if os.Getenv("GO_NO_RERUN") != "" {
		os.Exit(0)
		return true
	}
	return false
}

// C程序_写日志 将日志内容追加写入到指定路径的日志文件。
// 每条日志自动添加时间戳前缀（格式：2006-01-02 15:04:05）。
// 如果日志文件不存在，会自动创建。日志路径为空时默认写入程序同目录下的"运行日志.txt"。
//
// 参数:
//   - 日志内容: 要记录的日志文本
//   - 日志路径: 日志文件的完整路径，可为空（默认使用程序目录下的"运行日志.txt"）
func C程序_写日志(日志内容 string, 日志路径 string) {
	if 日志路径 == "" {
		exePath, _ := os.Executable()
		日志路径 = filepath.Dir(exePath) + "/运行日志.txt"
	}

	if _, err := os.Stat(日志路径); os.IsNotExist(err) {
		file, _ := os.Create(日志路径)
		_ = file.Close()
	}

	file, _ := os.OpenFile(日志路径, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer file.Close()

	logLine := fmt.Sprintf("%s   %s\n", time.Now().Format("2006-01-02 15:04:05"), 日志内容)
	file.WriteString(logLine)
}

// C程序_取命令行 获取启动程序时的命令行参数。
// 返回 os.Args 的完整切片，args[0] 为程序路径，args[1:] 为附加参数。
//
// 返回:
//   - []string: 命令行参数数组
func C程序_取命令行() []string {
	args := os.Args
	return args
}

// C程序_取运行目录 获取当前可执行文件所在的目录。
// 会解析符号链接，返回真实的目录路径。
//
// 返回:
//   - string: 可执行文件所在目录的绝对路径；失败时程序终止
func C程序_取运行目录() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res

}

// C程序_取临时目录 获取操作系统的临时目录路径。
// 依次尝试 TEMP、TMP 环境变量，最后回退到 SYSTEMROOT\Temp。
//
// 返回:
//   - string: 临时目录的绝对路径
func C程序_取临时目录() string {
	tempDir := os.Getenv("TEMP")
	if tempDir == "" {
		tempDir = os.Getenv("TMP")
	}

	if tempDir == "" {
		tempDir = filepath.Join(os.Getenv("SYSTEMROOT"), "Temp")
	}

	return tempDir
}

// C程序_运行Win 通过 PowerShell 执行命令行指令，并返回输出结果。
// 输出内容会自动从 GBK 编码转换为 UTF-8。
// 注意：此函数会阻塞等待命令执行完毕。
//
// 参数:
//   - 欲运行的命令行: 要执行的 PowerShell 命令
//
// 返回:
//   - string: 命令执行后的输出文本（UTF-8 编码）
func C程序_运行Win(欲运行的命令行 string) string {
	var err error

	cmd := exec.Command("powershell")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out
	go func(欲运行的命令行 string) {
		in.WriteString(欲运行的命令行)
	}(欲运行的命令行)
	err = cmd.Start()

	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	rt := W文本_gbk到utf8(out.String())

	return rt
}

// C程序_延时 暂停当前 goroutine 指定的毫秒数。
// 使用 time.Sleep 实现，精度取决于操作系统调度器。
//
// 参数:
//   - 毫秒数: 延时的毫秒数
//
// 返回:
//   - bool: 始终返回 true
func C程序_延时(毫秒数 int64) bool {
	time.Sleep(time.Duration(毫秒数) * time.Millisecond)
	return true
}
