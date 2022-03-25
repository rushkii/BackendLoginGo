package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

var sum [32]byte

func EncPwd(pwd string) string {
	sum = sha256.Sum256([]byte(pwd))
	return "SHA256:" + base64.StdEncoding.EncodeToString(sum[:])
}
