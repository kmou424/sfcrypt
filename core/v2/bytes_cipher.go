package v2

import (
	"github.com/kmou424/sfcrypt/core/keygen"
	"math"
)

type BytesCipher struct {
	gen keygen.KeyGen
}

func (c *BytesCipher) Init(gen keygen.KeyGen) {
	c.gen = gen
}

func (c *BytesCipher) Encrypt(data []byte) ([]byte, error) {
	return c.byte2byteWithKey(data, func(srcByte byte, keyByte byte) byte {
		return srcByte + keyByte
	})
}

func (c *BytesCipher) Decrypt(data []byte) ([]byte, error) {
	return c.byte2byteWithKey(data, func(srcByte byte, keyByte byte) byte {
		return srcByte - keyByte
	})
}

func (c *BytesCipher) byte2byteWithKey(src []byte, fun func(srcByte byte, keyByte byte) (dstByte byte)) (dst []byte, err error) {
	key := c.gen.GetKey()
	dataLen, keyLen := len(src), len(key)
	dst = make([]byte, dataLen)

	blkProcessNum := int(math.Ceil(float64(dataLen) / float64(keyLen)))
	for round := 0; round < blkProcessNum; round++ {
		for keyOffset := 0; keyOffset < keyLen; keyOffset++ {
			idx := round*keyLen + keyOffset
			if idx >= dataLen {
				break
			}
			dst[idx] = fun(src[idx], key[keyOffset])
		}
	}

	return dst, nil
}
