package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func generateSha256(text string) string {
	var f = sha256.New()
	f.Write([]byte(text))
	return fmt.Sprintf("%x", f.Sum(nil))
}

func timestampToSha256() string {
	var sha256Key = time.Now().UnixNano() + time.Now().UnixMicro()
	var toText = fmt.Sprintf("%d", sha256Key)
	return generateSha256(toText)
}
