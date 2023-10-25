package utils

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(src string) string {
	hash := sha256.New()
	hash.Write([]byte(src))

	result := hash.Sum(nil)
	return fmt.Sprintf("%x", result)
}
