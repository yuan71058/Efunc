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

func C程序_延时2(毫秒 int) bool {
	for i := 0; i < 毫秒; i++ {
		time.Sleep(1 * time.Millisecond)
	}
	return true
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

// 取运行目录
func C程序_取运行目录() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res

}

//调用格式： 〈逻辑型〉 运行 （文本型 欲运行的命令行，逻辑型 是否等待程序运行完毕，［整数型 被运行程序窗口显示方式］） - 系统核心支持库->系统处理
//英文名称：run
//本命令运行指定的可执行文件或者外部命令。如果成功，返回真，否则返回假。本命令为初级命令。
//参数<1>的名称为“欲运行的命令行”，类型为“文本型（text）”。
//参数<2>的名称为“是否等待程序运行完毕”，类型为“逻辑型（bool）”，初始值为“假”。
//参数<3>的名称为“被运行程序窗口显示方式”，类型为“整数型（int）”，可以被省略。参数值可以为以下常量之一：1、#隐藏窗口； 2、#普通激活； 3、#最小化激活； 4、#最大化激活； 5、#普通不激活； 6、#最小化不激活。如果省略本参数，默认为“普通激活”方式。
//
//操作系统需求： Windows、Linux

func C程序_运行Win(欲运行的命令行 string) string {
	var err error

	//cmd := exec.Command("cmd")
	cmd := exec.Command("powershell")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	go func(欲运行的命令行 string) {
		// start stop restart
		in.WriteString(欲运行的命令行) //写入你的命令，可以有多行，"\n"表示回车
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
	//fmt.Println(rt)

	return rt
}

func C程序_延时(毫秒数 int64) bool {
	time.Sleep(time.Duration(毫秒数) * time.Millisecond)
	return true
}
