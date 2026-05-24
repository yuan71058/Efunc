package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// W文件_是否存在 判断文件或目录是否存在。
//
// 参数:
//   - 路径: 文件或目录的路径
//
// 返回:
//   - bool: true 表示存在
func W文件_是否存在(路径 string) bool {
	_, err := os.Stat(路径)
	if err == nil {
		return true
	}

	return false
}

// W文件_写到文件 将字节数据写入文件。
// 如果父目录不存在，会自动创建。文件权限为 0644。
//
// 参数:
//   - 文件名: 目标文件路径
//   - 欲写入文件的数据: 待写入的字节数据
//
// 返回:
//   - error: 写入失败时返回错误信息
func W文件_写到文件(文件名 string, 欲写入文件的数据 []byte) error {
	if !W文件_是否存在(W文件_取父目录(文件名)) {
		M目录_创建(W文件_取父目录(文件名))
	}

	err := ioutil.WriteFile(文件名, 欲写入文件的数据, 0644)
	if err != nil {
		return err
	}
	return nil
}

// W文件_枚举 枚举指定目录下的文件，支持按扩展名过滤和递归遍历。
// 结果通过指针参数返回，会追加到现有数组中。
//
// 参数:
//   - 欲寻找的目录: 要枚举的根目录路径
//   - 欲寻找的文件名: 文件扩展名过滤，多个用 | 分隔（如 ".txt|.jpg"），空表示所有文件
//   - files: 用于接收结果的字符串数组指针
//   - 是否带路径: true 返回完整路径，false 仅返回文件名
//   - 是否遍历子目录: true 递归遍历子目录
//
// 返回:
//   - error: 枚举过程中遇到的错误
func W文件_枚举(欲寻找的目录 string, 欲寻找的文件名 string, files *[]string, 是否带路径 bool, 是否遍历子目录 bool) error {
	var ok bool
	欲寻找的文件名arr := strings.Split(欲寻找的文件名, "|")
	l, err := ioutil.ReadDir(欲寻找的目录)
	if err != nil {
		return err
	}

	separator := "/"

	for _, f := range l {
		tmp := 欲寻找的目录 + separator + f.Name()

		if f.IsDir() {
			if 是否遍历子目录 {
				err = W文件_枚举(tmp, 欲寻找的文件名, files, 是否带路径, 是否遍历子目录)
				if err != nil {
					return err
				}
			}
		} else {
			ok = false
			if !S数组_是否为空(欲寻找的文件名arr) {
				if isInSuffix(欲寻找的文件名arr, f.Name()) {
					ok = true

				}
			} else {
				ok = true
			}
			if ok {
				if 是否带路径 {
					*files = append(*files, tmp)
				} else {
					*files = append(*files, f.Name())
				}
			}
		}
	}
	return err
}

// isInSuffix 判断目标字符串的末尾是否含有数组中指定的后缀。
func isInSuffix(list []string, s string) (isIn bool) {

	isIn = false
	for _, f := range list {

		if strings.TrimSpace(f) != "" && strings.HasSuffix(s, f) {
			isIn = true
			break
		}
	}

	return isIn
}

// W文件_取文件名 从完整路径中提取文件名部分。
//
// 参数:
//   - 路径: 文件完整路径
//
// 返回:
//   - string: 文件名（含扩展名）
func W文件_取文件名(路径 string) string {
	return filepath.Base(路径)
}

// W文件_路径合并处理 将多个路径元素合并为一个路径。
// 使用 path.Join 处理路径分隔符。
//
// 参数:
//   - elem: 路径元素列表
//
// 返回:
//   - string: 合并后的路径
func W文件_路径合并处理(elem ...string) string {
	return path.Join(elem...)
}

// W文件_取父目录 获取文件路径的父目录。
//
// 参数:
//   - dirpath: 文件路径
//
// 返回:
//   - string: 父目录路径
func W文件_取父目录(dirpath string) string {
	return path.Dir(dirpath)
}

// W文件_删除 删除指定文件。
//
// 参数:
//   - 欲删除的文件名: 要删除的文件路径
//
// 返回:
//   - error: 删除失败时返回错误信息
func W文件_删除(欲删除的文件名 string) error {
	return os.Remove(欲删除的文件名)

}

// W文件_更名 重命名文件或目录。
//
// 参数:
//   - 欲更名的原文件或目录名: 原名称路径
//   - 欲更改为的现文件或目录名: 新名称路径
//
// 返回:
//   - error: 重命名失败时返回错误信息
func W文件_更名(欲更名的原文件或目录名 string, 欲更改为的现文件或目录名 string) error {
	return os.Rename(欲更名的原文件或目录名, 欲更改为的现文件或目录名)
}

// W文件_写出 将数据写入文件，父目录不存在时自动创建。
// 数据通过 D到字节集 转换为字节集后写入，文件权限为可执行权限。
//
// 参数:
//   - 文件名: 目标文件路径
//   - 欲写入文件的数据: 待写入的数据（支持字符串、字节集等类型）
//
// 返回:
//   - error: 写入失败时返回错误信息
func W文件_写出(文件名 string, 欲写入文件的数据 interface{}) error {
	fpath := W文件_取父目录(文件名)
	if !W文件_是否存在(fpath) {
		M目录_创建(fpath)
	}
	return ioutil.WriteFile(文件名, D到字节集(欲写入文件的数据), os.ModePerm)
}

// W文件_写出文件 W文件_写出 的别名，为使用习惯添加。
// 功能与 W文件_写出 完全相同。
//
// 参数:
//   - 文件名: 目标文件路径
//   - 欲写入文件的数据: 待写入的数据
//
// 返回:
//   - error: 写入失败时返回错误信息
func W文件_写出文件(文件名 string, 欲写入文件的数据 interface{}) error {
	return W文件_写出(文件名, 欲写入文件的数据)
}

// W文件_追加文本 向文件末尾追加文本内容，自动添加换行符。
// 文件不存在时自动创建，父目录不存在时自动创建。
//
// 参数:
//   - 文件名: 目标文件路径
//   - 欲追加的文本: 要追加的文本内容
//
// 返回:
//   - error: 追加失败时返回错误信息
func W文件_追加文本(文件名 string, 欲追加的文本 string) error {
	fpath := W文件_取父目录(文件名)
	if !W文件_是否存在(fpath) {
		M目录_创建(fpath)
	}
	file, err := os.OpenFile(文件名, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer file.Close()

	_, err = file.Write(D到字节集(欲追加的文本 + "\r\n"))
	return err
}

// W文件_读入文本 从文件中读取全部内容并返回字符串。
// 读取失败时返回空串。
//
// 参数:
//   - 文件名: 源文件路径
//
// 返回:
//   - string: 文件内容文本
func W文件_读入文本(文件名 string) string {
	var data []byte
	data, _ = ioutil.ReadFile(文件名)
	return D到文本(data)
}

// W文件_读入文件 从文件中读取全部内容并返回字节集。
// 读取失败时返回空字节集。
//
// 参数:
//   - 文件名: 源文件路径
//
// 返回:
//   - []byte: 文件内容字节集
func W文件_读入文件(文件名 string) []byte {
	data, err := ioutil.ReadFile(文件名)
	if err != nil {
		return []byte{}
	}
	return data
}

// W文件_保存 智能保存文件，仅在内容不同时才写入。
// 先检查文件是否存在，如果存在则比较内容是否一致，不一致才写出。
// 避免不必要的磁盘写入操作。
//
// 参数:
//   - 文件名: 目标文件路径
//   - 欲写入文件的数据: 待写入的数据
//
// 返回:
//   - error: 写入失败时返回错误信息；内容相同时返回 nil
func W文件_保存(文件名 string, 欲写入文件的数据 interface{}) error {
	if W文件_是否存在(文件名) {
		data := W文件_读入文件(文件名)
		wdata := D到字节集(欲写入文件的数据)
		if !bytes.Equal(data, wdata) {
			return W文件_写出(文件名, wdata)
		}
	} else {
		return W文件_写出(文件名, 欲写入文件的数据)
	}
	return nil
}

// W文件_取临时文件名 在指定目录中创建一个临时文件，返回文件对象和路径。
// 如果目录名为空，使用系统默认临时目录。
//
// 参数:
//   - 目录名: 临时文件所在目录，可为空
//
// 返回:
//   - f: 临时文件对象
//   - filepath: 临时文件的完整路径
//   - err: 创建失败时的错误信息
func W文件_取临时文件名(目录名 string) (f *os.File, filepath string, err error) {
	prefix := ""
	f, err = ioutil.TempFile(目录名, prefix)
	filepath = 目录名 + f.Name()
	return f, filepath, err
}

// W文件_取大小 获取文件的大小（字节数）。
//
// 参数:
//   - 文件名: 文件路径
//
// 返回:
//   - int64: 文件大小（字节）；文件不存在返回 -1
func W文件_取大小(文件名 string) int64 {
	f, err := os.Stat(文件名)
	if err == nil {
		return f.Size()
	} else {
		return -1
	}
}
