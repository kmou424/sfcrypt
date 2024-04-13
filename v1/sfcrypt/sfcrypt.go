package sfcrypt

type SFCrypt struct {
	password   string
	threads    int
	blockRatio int
}

func NewSFCrypt(password string, salt string) *SFCrypt {
	return &SFCrypt{
		password:   NewPasswordFactory(password, salt).GenerateHash(),
		threads:    DefaultThreads,
		blockRatio: BlockRatio,
	}
}

func (s *SFCrypt) SetThreads(threads int) *SFCrypt {
	if threads > MaxThreads {
		threads = MaxThreads
	} else if threads < 0 {
		threads = DefaultThreads
	}
	s.threads = threads
	return s
}

func (s *SFCrypt) SetBlockRatio(ratio int) {
	s.blockRatio = ratio
}
