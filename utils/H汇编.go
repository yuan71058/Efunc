package utils

import (
	"math/rand"
	"time"
)

func H汇编_取随机数(起始数, 结束数 int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(结束数-起始数+1) + 起始数
}
