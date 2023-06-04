package utils

import (
	"bytes"
	"encoding/json"
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

func JsonPretty(elem interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(elem); err != nil {
		return ""
	}
	return buffer.String()
}
