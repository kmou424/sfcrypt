package keygen

type Options struct {
	Password []byte
	Salt     []byte
}

type KeyGen interface {
	Generate(opt *Options) error
	GetKey() []byte
}
