package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func GenerateRandomUsername() string {
	// 使用 UUID 作为基础，确保唯一性
	id := uuid.New()
	// 取 UUID 的前 8 个字符，并添加时间戳，以增加可读性
	return fmt.Sprintf("user_%s_%d", id.String()[:8], time.Now().Unix())
}

// GenerateSecurePassword 生成一个安全的随机密码
func GenerateSecurePassword() (string, error) {
	const length = 16
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
