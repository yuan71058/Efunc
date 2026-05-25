package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// File_Exists 判断文件或目录是否存在。
//
// 参数:
//   - path: 文件或目录的路径
//
// 返回:
//   - bool: true 表示存在
func File_Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// File_Write 将字节数据写入文件。
// 如果父目录不存在，会自动创建。文件权限为 0644。
//
// 参数:
//   - filePath: 目标文件路径
//   - data: 待写入的字节数据
//
// 返回:
//   - error: 写入失败时返回错误信息
func File_Write(filePath string, data []byte) error {
	parentDir := File_GetParentDir(filePath)
	if !File_Exists(parentDir) {
		Dir_Create(parentDir)
	}

	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// File_List 枚举指定目录下的文件，支持按扩展名过滤和递归遍历。
// 结果通过指针参数返回，会追加到现有数组中。
//
// 参数:
//   - targetDir: 要枚举的根目录路径
//   - filterExts: 文件扩展名过滤，多个用 | 分隔（如 ".txt|.jpg"），空表示所有文件
//   - files: 用于接收结果的字符串数组指针
//   - includePath: true 返回完整路径，false 仅返回文件名
//   - recursive: true 递归遍历子目录
//
// 返回:
//   - error: 枚举过程中遇到的错误
func File_List(targetDir string, filterExts string, files *[]string, includePath bool, recursive bool) error {
	var ok bool
	filterExtsArr := strings.Split(filterExts, "|")
	l, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return err
	}

	separator := "/"

	for _, f := range l {
		tmp := targetDir + separator + f.Name()

		if f.IsDir() {
			if recursive {
				err = File_List(tmp, filterExts, files, includePath, recursive)
				if err != nil {
					return err
				}
			}
		} else {
			ok = false
			if len(filterExtsArr) > 0 && strings.TrimSpace(filterExts) != "" {
				if isInSuffixList(filterExtsArr, f.Name()) {
					ok = true
				}
			} else {
				ok = true
			}
			if ok {
				if includePath {
					*files = append(*files, tmp)
				} else {
					*files = append(*files, f.Name())
				}
			}
		}
	}
	return err
}

// isInSuffixList 判断目标字符串的末尾是否含有数组中指定的后缀。
func isInSuffixList(list []string, s string) (isIn bool) {
	isIn = false
	for _, f := range list {
		if strings.TrimSpace(f) != "" && strings.HasSuffix(s, f) {
			isIn = true
			break
		}
	}
	return isIn
}

// File_GetName 从完整路径中提取文件名部分。
//
// 参数:
//   - filePath: 文件完整路径
//
// 返回:
//   - string: 文件名（含扩展名）
func File_GetName(filePath string) string {
	return filepath.Base(filePath)
}

// File_PathJoin 将多个路径元素合并为一个路径。
// 使用 path.Join 处理路径分隔符。
//
// 参数:
//   - elem: 路径元素列表
//
// 返回:
//   - string: 合并后的路径
func File_PathJoin(elem ...string) string {
	return path.Join(elem...)
}

// File_GetParentDir 获取文件路径的父目录。
//
// 参数:
//   - filePath: 文件路径
//
// 返回:
//   - string: 父目录路径
func File_GetParentDir(filePath string) string {
	return path.Dir(filePath)
}

// File_Remove 删除指定文件。
//
// 参数:
//   - filePath: 要删除的文件路径
//
// 返回:
//   - error: 删除失败时返回错误信息
func File_Remove(filePath string) error {
	return os.Remove(filePath)
}

// File_Rename 重命名文件或目录。
//
// 参数:
//   - oldName: 原名称路径
//   - newName: 新名称路径
//
// 返回:
//   - error: 重命名失败时返回错误信息
func File_Rename(oldName string, newName string) error {
	return os.Rename(oldName, newName)
}

// File_WriteAny 将数据写入文件，父目录不存在时自动创建。
// 数据通过 ToBytes 转换为字节集后写入，文件权限为可执行权限。
//
// 参数:
//   - filePath: 目标文件路径
//   - data: 待写入的数据（支持字符串、字节集等类型）
//
// 返回:
//   - error: 写入失败时返回错误信息
func File_WriteAny(filePath string, data interface{}) error {
	parentDir := File_GetParentDir(filePath)
	if !File_Exists(parentDir) {
		Dir_Create(parentDir)
	}
	return ioutil.WriteFile(filePath, ToBytes(data), os.ModePerm)
}

// File_AppendText 向文件末尾追加文本内容，自动添加换行符。
// 文件不存在时自动创建，父目录不存在时自动创建。
//
// 参数:
//   - filePath: 目标文件路径
//   - text: 要追加的文本内容
//
// 返回:
//   - error: 追加失败时返回错误信息
func File_AppendText(filePath string, text string) error {
	parentDir := File_GetParentDir(filePath)
	if !File_Exists(parentDir) {
		Dir_Create(parentDir)
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(ToBytes(text + "\r\n"))
	return err
}

// File_ReadText 从文件中读取全部内容并返回字符串。
// 读取失败时返回空串。
//
// 参数:
//   - filePath: 源文件路径
//
// 返回:
//   - string: 文件内容文本
func File_ReadText(filePath string) string {
	data, _ := ioutil.ReadFile(filePath)
	return ToString(data)
}

// File_ReadBytes 从文件中读取全部内容并返回字节集。
// 读取失败时返回空字节集。
//
// 参数:
//   - filePath: 源文件路径
//
// 返回:
//   - []byte: 文件内容字节集
func File_ReadBytes(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte{}
	}
	return data
}

// File_Save 智能保存文件，仅在内容不同时才写入。
// 先检查文件是否存在，如果存在则比较内容是否一致，不一致才写出。
// 避免不必要的磁盘写入操作。
//
// 参数:
//   - filePath: 目标文件路径
//   - data: 待写入的数据
//
// 返回:
//   - error: 写入失败时返回错误信息；内容相同时返回 nil
func File_Save(filePath string, data interface{}) error {
	if File_Exists(filePath) {
		existingData := File_ReadBytes(filePath)
		newData := ToBytes(data)
		if !bytes.Equal(existingData, newData) {
			return File_WriteAny(filePath, newData)
		}
	} else {
		return File_WriteAny(filePath, data)
	}
	return nil
}

// File_GetTempFile 在指定目录中创建一个临时文件，返回文件对象和路径。
// 如果目录名为空，使用系统默认临时目录。
//
// 参数:
//   - dir: 临时文件所在目录，可为空
//
// 返回:
//   - f: 临时文件对象
//   - filePath: 临时文件的完整路径
//   - err: 创建失败时的错误信息
func File_GetTempFile(dir string) (f *os.File, filePath string, err error) {
	prefix := ""
	f, err = ioutil.TempFile(dir, prefix)
	filePath = dir + f.Name()
	return f, filePath, err
}

// File_GetSize 获取文件的大小（字节数）。
//
// 参数:
//   - filePath: 文件路径
//
// 返回:
//   - int64: 文件大小（字节）；文件不存在返回 -1
func File_GetSize(filePath string) int64 {
	f, err := os.Stat(filePath)
	if err == nil {
		return f.Size()
	}
	return -1
}