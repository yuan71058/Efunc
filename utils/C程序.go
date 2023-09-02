package utils

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func C程序_延时(秒数 int) bool {
	for i := 0; i < 秒数; {
		time.Sleep(1 * time.Second)
	}
	return true
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
