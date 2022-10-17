package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(MD5("test"))
	fmt.Println(Sha256("test"))
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))

	// h := md5.New()
	// h.Write([]byte(s))
	// return hex.EncodeToString(h.Sum(nil))
}
