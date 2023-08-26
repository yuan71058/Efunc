package utils

import (
	"fmt"
	"testing"
)

func Test_启动子程序(t *testing.T) {
	fmt.Println("执行完毕")

	fmt.Println(B编码_URL编码("https%3A%2F%2Fwww.example.com%2F%3Fq%3D%E4%B8%AD%E6%96%87"))
	fmt.Println(B编码_URL解码("https%3A%2F%2Fwww.example.com%2F%3Fq%3D%E4%B8%AD%E6%96%87"))
}
