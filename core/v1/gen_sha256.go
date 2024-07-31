package v1

import (
	"crypto/sha256"
)

func SHA256(src string) string {
	hash := sha256.New()
	hash.Write([]byte(src))

	return string(hash.Sum(nil))
}
