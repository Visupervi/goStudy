package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	tempStr := h.Sum(nil)

	//fmt.Println(hex.EncodeToString(tempStr))
	return hex.EncodeToString(tempStr)
}

// MD5Encode 大写
func MD5Encode(str string) string {
	return strings.ToUpper(Md5Encode(str))
}

// MakeRandomNum 密码=MD5+随机数
func MakeRandomNum(pwd, salt string) string {
	return Md5Encode(pwd + salt)
}

// ValidPassword 密码解密
func ValidPassword(pwd, salt string, password string) bool {
	return Md5Encode(pwd+salt) == password
}
