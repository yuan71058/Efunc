package utils

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func C程序_延时(秒数 int) bool {
	for i := 0; i < 秒数; {
		time.Sleep(1 * time.Second)
	}
	return true
}

func C程序_延时2(毫秒 int) bool {
	for i := 0; i < 毫秒; {
		time.Sleep(1 * time.Millisecond)
	}
	return true
}

// 取运行目录
func C程序_取运行目录() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

func C程序_取cmd路径() string {
	file, err := os.Readlink("." + "cmd/cmd.exe")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// 生成标准的GUID格式：635897F8-2A48-4882-B3E1-823B8E5B6DF8
func C程序_取GUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// 设置GUID版本和变体
	b[6] = (b[6] & 0x0f) | 0x40 // Version 4
	b[8] = (b[8] & 0x3f) | 0x80 // Variant is 10
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// 待修复 感觉有问题,自身被占用的问题
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

func C程序_是否被调试() bool {
	return os.Getpid() != 0
}

func C程序_禁止重复运行() bool {
	if os.Getenv("GO_NO_RERUN") != "" {
		os.Exit(0)
		return true
	}
	return false
}

// 在程序根目录建立一个txt文件用于记录相关日志内容
func C程序_写日志(日志内容 string, 日志路径 /*可为空*/ string) {
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

// 本命令可以取出在启动易程序时附加在其可执行文件名后面的所有以空格分隔的命令行文本段
func C程序_取命令行() []string {
	args := os.Args
	/*	// 打印程序名称
		fmt.Println("程序名称:", args[0])
		// 打印命令行参数
		fmt.Println("命令行参数:")
		for i := 1; i < len(args); i++ {
			fmt.Println(i, args[i])
		}*/
	return args
}
