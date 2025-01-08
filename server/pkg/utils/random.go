package utils

import (
	"math/rand"
)

// 生成指定长度的随机字符串
func RandString(n int, letters string) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(letters[rand.Intn(len(letters))])
	}
	return string(b)
}

// 根据性别随机生成头像
func RandAvatar(gender uint) string {
	idx := RandString(1, "123456")
	if gender == 1 {
		return "/image/avatar/male_" + idx + ".svg"
	} else if gender == 2 {
		return "/image/avatar/female_" + idx + ".svg"
	}
	return "/image/avatar/default.png"
}
