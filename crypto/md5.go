package cryptoutil

import (
	"crypto/md5"
	"fmt"
)

// Md5 md5加密
func Md5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
