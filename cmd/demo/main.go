package main

import (
	"fmt"

	"github.com/yuan71058/Efunc/class"
	. "github.com/yuan71058/Efunc/utils"
)

func main() {
	fmt.Printf(B编码_URL编码("go语言版的精易模块"))
}

func 测试队列() {
	var 队列 = class.L_队列{}
	for i := 0; i < 1000; i++ {
		队列.J加入队列(fmt.Sprintf("%d", i))
	}
	局_临时文本 := ""
	局_临时文本2 := ""
	for 队列.Q取队列长度() > 0 {
		go func() {
			if 队列.T弹出队列文本(&局_临时文本) {
				局_临时文本2 += (局_临时文本 + "\r\n")
			}
		}()
	}
	C程序_延时(5000)
	fmt.Println("最终" + fmt.Sprintf("%d", W文本_取行数(局_临时文本2)))
	fmt.Println(局_临时文本2)
}
