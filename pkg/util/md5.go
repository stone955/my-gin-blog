package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(src string) string {
	m := md5.New()
	_, _ = m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}
