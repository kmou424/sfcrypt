package sfcrypt

func (s *SFCrypt) CryptBytes(in []byte) {
	xorCryptBytes(in, s.password)
}
