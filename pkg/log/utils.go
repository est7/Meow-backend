package log

import (
	"bytes"
)

// GetLogFile get log file absolute path
func GetLogFile(filename string, suffix string) string {
	return ConcatString(logDir, "/", hostname, "/", filename, suffix)
}

// ConcatString 连接字符串
// NOTE: 性能比fmt.Sprintf和+号要好
func ConcatString(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}
