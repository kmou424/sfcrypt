package sfcrypt

import (
	"github.com/kmou424/sfcrypt/internal/consts"
	"github.com/kmou424/sfcrypt/internal/generator"
	"github.com/kmou424/sfcrypt/internal/kit"
)

type SFCrypt struct {
	password   string
	threads    int
	blockRatio int
}

func NewSFCrypt(password string, salt string) *SFCrypt {
	return &SFCrypt{
		password:   generator.NewPasswordFactory(password, salt).GenerateHash(),
		threads:    consts.DefaultThreads,
		blockRatio: consts.BlockRatio,
	}
}

func (s *SFCrypt) SetThreads(threads int) *SFCrypt {
	if threads > consts.MaxThreads {
		threads = consts.MaxThreads
	} else if threads < 0 {
		threads = consts.DefaultThreads
	}
	s.threads = threads
	return s
}

func (s *SFCrypt) SetBlockRatio(ratio int) {
	s.blockRatio = ratio
}

func (s *SFCrypt) CryptBytes(in []byte, length int) []byte {
	return kit.XORCryptBytes(in, length, s.password)
}

func (s *SFCrypt) CryptFile(input string, output string) {
	var blockSize = consts.BufferSize * s.blockRatio
	doSFCrypt(input, output, blockSize, s.password, s.threads)
}
