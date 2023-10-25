package sfcrypt

import (
	"github.com/kmou424/sfcrypt/internal/consts"
	"github.com/kmou424/sfcrypt/internal/core"
	"github.com/kmou424/sfcrypt/internal/factory"
	"github.com/kmou424/sfcrypt/internal/utils"
)

type SFCrypt struct {
	password   string
	threads    int
	blockRatio int
}

func NewSFCrypt(password string, salt string) *SFCrypt {
	return &SFCrypt{
		password:   factory.NewPasswordFactory(password, salt).GenerateHash(),
		threads:    consts.DefaultThreads,
		blockRatio: consts.DefaultBlockRatio,
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
	return utils.XORCryptBytes(in, length, s.password)
}

func (s *SFCrypt) CryptFile(input string, output string) {
	var blockSize = consts.BufferSize * s.blockRatio
	core.SFCryptFile(input, output, blockSize, s.password, s.threads)
}
