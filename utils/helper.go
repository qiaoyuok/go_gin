package utils

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"unicode/utf8"
)

// GetUtf8 获取utf8编码的内容
func GetUtf8(str string) string {
	if !utf8.Valid([]byte(str)) {
		bytes, err := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(str))
		if err != nil {
			return str
		}
		return string(bytes)
	}
	return str
}
