package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// HashKey - Hash a key to MD5
func HashKey(key string) string {
	sum := md5.Sum([]byte(key))
	return hex.EncodeToString(sum[:])
}
