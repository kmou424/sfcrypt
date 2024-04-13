package sfcrypt

import (
	"math"
)

func xorCryptBytes(src []byte, password string) {
	srcLen := len(src)
	pwLen := len(password)
	pwBlockNum := int(math.Ceil(float64(srcLen) / float64(pwLen)))
	for i := 0; i < pwBlockNum; i++ {
		for j := 0; j < pwLen; j++ {
			idx := i*pwLen + j
			if idx >= srcLen {
				break
			}
			src[idx] = src[idx] ^ password[j]
		}
	}
}
