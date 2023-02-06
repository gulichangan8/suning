package tool

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func Hmac(key, date string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(date))
	return hex.EncodeToString(hash.Sum(nil))
}
