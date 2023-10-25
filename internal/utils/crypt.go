package utils

func XORCryptBytes(src []byte, actualLen int, password string) []byte {
	buf := make([]byte, actualLen)
	pwLen := len(password)
	for i := 0; i < int(float32(actualLen)/64.0+0.5); i++ {
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
