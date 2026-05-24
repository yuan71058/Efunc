package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// M目录_是否存在 检查指定路径的目录是否存在。
//
// 参数:
//   - path: 目录路径
//
// 返回:
//   - bool: true 表示目录存在
//   - error: 路径存在但不是目录时返回错误；不存在时返回 nil
func M目录_是否存在(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// M目录_创建 递归创建目录（含所有必要的父目录）。
// 目录已存在时不做任何操作，直接返回 nil。
//
// 参数:
//   - 路径: 要创建的目录路径
//
// 返回:
//   - error: 创建失败时返回错误信息
func M目录_创建(路径 string) error {
	return os.MkdirAll(路径, os.ModePerm)
}

// M目录_枚举子目录 枚举指定目录下的子目录，结果通过指针参数返回。
// 可选择是否带完整路径、是否递归枚举子目录的子目录。
//
// 参数:
//   - 父文件夹路径: 要枚举的根目录路径
//   - 子目录数组: 用于接收结果的字符串数组指针（会追加到现有内容）
//   - 是否带路径: true 返回完整路径，false 仅返回目录名
//   - 是否继续向下枚举: true 递归枚举所有层级的子目录
//
// 返回:
//   - error: 枚举过程中遇到的错误
func M目录_枚举子目录(父文件夹路径 string, 子目录数组 *[]string, 是否带路径 bool, 是否继续向下枚举 bool) error {
	l, err := ioutil.ReadDir(父文件夹路径)
	if err != nil {
		return err
	}
	separator := "/"
	for _, f := range l {
		tmp := string(父文件夹路径 + separator + f.Name())

		if f.IsDir() {
			if 是否带路径 {
				*子目录数组 = append(*子目录数组, tmp)
			} else {
				*子目录数组 = append(*子目录数组, f.Name())
			}
			if 是否继续向下枚举 {
				err = M目录_枚举子目录(tmp, 子目录数组, 是否带路径, 是否继续向下枚举)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

// M目录_取运行目录 获取当前可执行文件所在的目录。
// 路径中的反斜杠会被替换为正斜杠。
//
// 返回:
//   - string: 可执行文件所在目录的绝对路径
func M目录_取运行目录() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// M目录_取当前目录 获取当前工作目录。
//
// 返回:
//   - string: 当前工作目录的绝对路径
func M目录_取当前目录() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

// M目录_删除 递归删除指定目录及其所有内容（文件和子目录）。
// 等同于 rm -rf，请谨慎使用。
//
// 参数:
//   - 欲删除的目录名称: 要删除的目录路径
//
// 返回:
//   - error: 删除失败时返回错误信息
func M目录_删除(欲删除的目录名称 string) error {
	return os.RemoveAll(欲删除的目录名称)
}
