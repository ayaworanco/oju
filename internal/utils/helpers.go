package utils

import (
	"crypto/md5"
	"io"
	"time"
)

func GenerateId() string {
	hash := md5.New()
	io.WriteString(hash, time.Now().String())
	return string(hash.Sum(nil))
}
