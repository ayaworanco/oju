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

func MapPut[Key comparable, Value any](old_map map[Key]Value, key Key, value Value) map[Key]Value {
	new_map := make(map[Key]Value)
	for new_key, new_value := range old_map {
		new_map[new_key] = new_value
	}

	new_map[key] = value

	return new_map
}
