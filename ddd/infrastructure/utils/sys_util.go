package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(encode string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(encode)))
}
