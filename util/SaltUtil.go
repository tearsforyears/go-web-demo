package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func Salt(str string) string {
	hash := Hash(str)
	//md5 := Hash(str)
	if s, err := strconv.ParseInt(hash[:4], 16, 32); err == nil {
		salt := s % 1024
		return Md5(str + strconv.FormatInt(salt, 10))
	} else {
		panic(err)
	}
}

func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

