package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
)

func GenerateId() string {
	hash := md5.New()
	io.WriteString(hash, time.Now().String())
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum)
}
