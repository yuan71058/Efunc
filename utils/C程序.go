package utils

import "time"

func C程序_延时(秒数 int) bool {
	for i := 0; i < 秒数; {
		time.Sleep(1 * time.Second)
	}
	return true
}
