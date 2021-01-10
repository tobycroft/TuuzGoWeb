package Input

import "strings"

func Fliter_Ascii(json_string string) string {
	originStr := json_string

	// 将字符串转换为rune数组
	srcRunes := []rune(originStr)

	// 创建一个新的rune数组，用来存放过滤后的数据
	dstRunes := make([]rune, 0, len(srcRunes))

	// 过滤不可见字符，根据上面的表的0-32和127都是不可见的字符
	for _, c := range srcRunes {
		if c >= 0 && c <= 31 {
			continue
		}
		if c == 127 {
			continue
		}
		dstRunes = append(dstRunes, c)
	}

	return string(dstRunes)
}

func Fliter_error_encode(json_string string) string {
	return strings.ReplaceAll(json_string, "\\'", "'")
}
