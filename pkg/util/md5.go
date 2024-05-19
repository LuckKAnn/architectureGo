package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	// 十六进制
	return hex.EncodeToString(m.Sum(nil))
}
