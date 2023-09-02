package utils

import (
	"io/ioutil"
	"os"
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

func W文件_写到文件(文件名 string, 欲写入文件的数据 []byte) error {

	err := ioutil.WriteFile(文件名, 欲写入文件的数据, 0644)
	if err != nil {
		return err
	}
	return nil
}
