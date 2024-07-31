package v1

func (s *SFCrypt) CryptBytes(in []byte) {
	xorCryptBytes(in, s.password)
}
