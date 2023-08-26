package utils

import (
	"log"
	"os"
	"path/filepath"
)

// W文件_是否存在 判断一个文件或文件夹是否存在
// 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在
func W文件_是否存在(路径 string) bool {
	_, err := os.Stat(路径)
	if err == nil {
		return true
	}

	return false
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
