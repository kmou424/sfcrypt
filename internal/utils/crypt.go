package utils

import (
	"math"
)

func XORCryptBytes(src []byte, actualLen int, password string) []byte {
	buf := make([]byte, actualLen)
	pwLen := len(password)
	pwBlockNum := int(math.Ceil(float64(actualLen) / float64(pwLen)))
	for i := 0; i < pwBlockNum; i++ {
		for j := 0; j < pwLen; j++ {
			idx := i*pwLen + j
			if idx >= actualLen {
				break
			}
			buf[idx] = src[idx] ^ password[j]
		}
	}
	return buf
}
