package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func MD5Cncode(str string) string {
	return strings.ToUpper(Md5Encode(str))
}

// 密码=MD5+随机数
func MakeRandomNum(pwd, salt string) string {
	return Md5Encode(pwd + salt)
}

// 密码解密
func ValidPassword(pwd, salt string, password string) bool {
	return Md5Encode(pwd+salt) == password
}
