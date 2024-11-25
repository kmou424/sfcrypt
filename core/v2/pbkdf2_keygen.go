package v2

import (
	"crypto/sha256"
	"github.com/kmou424/ero"
	"github.com/kmou424/sfcrypt/core/keygen"
	"golang.org/x/crypto/pbkdf2"
)

const (
	pbkdf2Iteration = 200000
	pbkdf2KeySize   = 512
)

type PBKDF2KeyGen struct {
	key []byte
}

func (kg *PBKDF2KeyGen) Generate(opt *keygen.Options) error {
	if opt == nil {
		return ero.Newf("can't initialize PBKDF2KeyGen with nil options")
	}
	kg.key = pbkdf2.Key(opt.Password, opt.Salt, pbkdf2Iteration, pbkdf2KeySize, sha256.New)
	return nil
}

func (kg *PBKDF2KeyGen) GetKey() []byte {
	return kg.key
}

func NewPBKDF2KeyGen(opt *keygen.Options) (*PBKDF2KeyGen, error) {
	kg := &PBKDF2KeyGen{}
	return kg, kg.Generate(opt)
}
