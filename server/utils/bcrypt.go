package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashAndSalt 加密成哈希字符串
func HashAndSalt(str []byte) string {
	hash, err := bcrypt.GenerateFromPassword(str, bcrypt.MinCost)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return string(hash)
}

// ValidatePasswords 验证哈希密码
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}