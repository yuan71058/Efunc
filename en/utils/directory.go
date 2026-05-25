package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Dir_Exists 检查指定路径的目录是否存在。
//
// 参数:
//   - path: 目录路径
//
// 返回:
//   - bool: true 表示目录存在
//   - error: 路径存在但不是目录时返回错误；不存在时返回 nil
func Dir_Exists(path string) (bool, error) {
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

// Dir_Create 递归创建目录（含所有必要的父目录）。
// 目录已存在时不做任何操作，直接返回 nil。
//
// 参数:
//   - path: 要创建的目录路径
//
// 返回:
//   - error: 创建失败时返回错误信息
func Dir_Create(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Dir_ListSubdirs 枚举指定目录下的子目录，结果通过指针参数返回。
// 可选择是否带完整路径、是否递归枚举子目录的子目录。
//
// 参数:
//   - parentPath: 要枚举的根目录路径
//   - subdirs: 用于接收结果的字符串数组指针（会追加到现有内容）
//   - includePath: true 返回完整路径，false 仅返回目录名
//   - recursive: true 递归枚举所有层级的子目录
//
// 返回:
//   - error: 枚举过程中遇到的错误
func Dir_ListSubdirs(parentPath string, subdirs *[]string, includePath bool, recursive bool) error {
	l, err := ioutil.ReadDir(parentPath)
	if err != nil {
		return err
	}
	separator := "/"
	for _, f := range l {
		tmp := parentPath + separator + f.Name()

		if f.IsDir() {
			if includePath {
				*subdirs = append(*subdirs, tmp)
			} else {
				*subdirs = append(*subdirs, f.Name())
			}
			if recursive {
				err = Dir_ListSubdirs(tmp, subdirs, includePath, recursive)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

// Dir_GetRunDir 获取当前可执行文件所在的目录。
// 路径中的反斜杠会被替换为正斜杠。
//
// 返回:
//   - string: 可执行文件所在目录的绝对路径
func Dir_GetRunDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// Dir_GetCurrent 获取当前工作目录。
//
// 返回:
//   - string: 当前工作目录的绝对路径
func Dir_GetCurrent() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

// Dir_Remove 递归删除指定目录及其所有内容（文件和子目录）。
// 等同于 rm -rf，请谨慎使用。
//
// 参数:
//   - path: 要删除的目录路径
//
// 返回:
//   - error: 删除失败时返回错误信息
func Dir_Remove(path string) error {
	return os.RemoveAll(path)
}