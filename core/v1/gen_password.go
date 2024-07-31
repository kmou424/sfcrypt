package v1

import (
	"fmt"
	"gonum.org/v1/gonum/mathext/prng"
)

type PasswordGenerator struct {
	password string
	saltSrc  string
}

func NewPasswordFactory(password string, salt string) *PasswordGenerator {
	return &PasswordGenerator{password, salt}
}

func (p *PasswordGenerator) GenerateHash() string {
	if len(p.saltSrc) == 0 {
		return SHA256(p.password)
	}
	return SHA256(p.password + p.generateSalt())
}

func (p *PasswordGenerator) generateSalt() (salt string) {
	var saltRuneList [][]rune
	var saltCharLength int

	saltHash := SHA256(p.saltSrc)
	runeListLength := len(saltHash)

	for _, ch := range saltHash {
		prngRunes := p.getPRNG(int(ch))
		saltCharLength += len(prngRunes)
		saltRuneList = append(saltRuneList, prngRunes)
	}
	for saltCharLength > 0 {
		for i := 0; i < runeListLength; i++ {
			if len(saltRuneList[i]) > 0 {
				salt += string(saltRuneList[i][0])
				saltRuneList[i] = saltRuneList[i][1:]
				saltCharLength--
			}
		}
	}
	return
}

func (p *PasswordGenerator) getPRNG(seed int) []rune {
	prnger := prng.NewMT19937_64()
	prnger.Seed(uint64(seed))
	var ret []rune
	for _, c := range SHA256(fmt.Sprintf("%d", prnger.Uint64())) {
		ret = append(ret, c)
	}
	return ret
}
