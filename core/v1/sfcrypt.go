package v1

import (
	"github.com/kmou424/sfcrypt/app/common"
)

type SFCrypt struct {
	password   string
	threads    int
	blockRatio int
}

func NewSFCrypt(password string, salt string) *SFCrypt {
	return &SFCrypt{
		password:   NewPasswordFactory(password, salt).GenerateHash(),
		threads:    common.DefaultRoutines,
		blockRatio: common.BlockRatio,
	}
}

func (s *SFCrypt) SetThreads(threads int) *SFCrypt {
	if threads > common.MaxRoutines {
		threads = common.MaxRoutines
	} else if threads < 0 {
		threads = common.DefaultRoutines
	}
	s.threads = threads
	return s
}

func (s *SFCrypt) SetBlockRatio(ratio int) {
	s.blockRatio = ratio
}
