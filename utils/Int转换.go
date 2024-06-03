package utils

// 防止精度丢失
func Int取绝对值(值 int) int {
	if 值 >= 0 {
		return 值
	}
	return -值
}
