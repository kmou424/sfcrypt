package cipher

import "github.com/kmou424/sfcrypt/core/keygen"

// Cipher is a cipher used to encrypt/decrypt bytes
type Cipher interface {
	Init(gen keygen.KeyGen)
	Encrypt(data []byte) error
	Decrypt(data []byte) error
}

type FileCipherOptions struct {
	Input  string
	Output string
	Force  bool
	Cipher Cipher
	// todo: implement encryption/decryption iterations
}

// FileCipher is a file cipher used to encrypt/decrypt files
type FileCipher interface {
	Init(config *FileCipherOptions)
	Encrypt() error
	Decrypt() error
}
