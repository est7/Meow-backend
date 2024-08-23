package validator

import (
	"net"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// IsValidUsername 检查用户名是否有效
func IsValidUsername(username string) bool {
	// 例如：用户名长度在3-20之间，只包含字母、数字和下划线
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]{3,20}$`, username)
	return matched
}

// IsValidEmail 检查邮箱是否有效
func IsValidEmail(email string) bool {
	// 简单的邮箱格式检查
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return matched
}

// IsValidPassword 检查密码是否有效
func IsValidPassword(password string) bool {
	// 例如：密码长度至少8位，包含至少一个大写字母，一个小写字母和一个数字
	return len(password) >= 8 &&
		strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
		strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") &&
		strings.ContainsAny(password, "0123456789")
}

// IsValidPhoneNumber 检查手机号码是否有效（以中国大陆号码为例）
func IsValidPhoneNumber(phone string) bool {
	matched, _ := regexp.MatchString(`^1[3-9]\d{9}$`, phone)
	return matched
}

// IsValidURL 检查URL是否有效
func IsValidURL(url string) bool {
	matched, _ := regexp.MatchString(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(?:/\S*)?$`, url)
	return matched
}

// IsValidIPAddress 检查IP地址是否有效
func IsValidIPAddress(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsValidDate 检查日期格式是否有效（YYYY-MM-DD）
func IsValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// IsValidCreditCard 检查信用卡号是否有效（使用Luhn算法）
func IsValidCreditCard(number string) bool {
	var sum int
	var alternate bool

	for i := len(number) - 1; i >= 0; i-- {
		n := int(number[i] - '0')
		if alternate {
			n *= 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}
		sum += n
		alternate = !alternate
	}

	return (sum%10 == 0) && (len(number) >= 13) && (len(number) <= 19)
}

// IsStrongPassword 检查是否为强密码
func IsStrongPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// IsAlphanumeric 检查字符串是否只包含字母和数字
func IsAlphanumeric(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(s)
}

// IsNumeric 检查字符串是否只包含数字
func IsNumeric(s string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(s)
}

// IsValidHexColor 检查是否为有效的十六进制颜色代码
func IsValidHexColor(color string) bool {
	matched, _ := regexp.MatchString(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`, color)
	return matched
}

// IsValidPostalCode 检查是否为有效的邮政编码（以中国为例）
func IsValidPostalCode(code string) bool {
	matched, _ := regexp.MatchString(`^[1-9]\d{5}$`, code)
	return matched
}
