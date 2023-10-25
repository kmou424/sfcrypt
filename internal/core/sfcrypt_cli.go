package core

import (
	"fmt"
	"github.com/kmou424/sfcrypt/internal/consts"
	"github.com/kmou424/sfcrypt/internal/factory"
	"time"
)

type SFCryptCliArgs struct {
	input    string
	output   string
	password string
	salt     string
	threads  int
}

func NewSFCryptCliArgs(input string, output string, password string, salt string, threads int) *SFCryptCliArgs {
	return &SFCryptCliArgs{
		input,
		output,
		password,
		salt,
		threads,
	}
}

var passwordHash string

func SFCryptCli(args *SFCryptCliArgs) {
	passwordHash = factory.NewPasswordFactory(args.password, args.salt).GenerateHash()

	if args.threads > consts.MaxThreads {
		args.threads = consts.MaxThreads
	} else if args.threads < 0 {
		args.threads = consts.DefaultThreads
	}

	start := time.Now()
	doSFCrypt(args)
	end := time.Now()

	fmt.Println("Process task took: ", end.Sub(start))
}

var blockSize = consts.BufferSize * consts.DefaultBlockRatio

func doSFCrypt(args *SFCryptCliArgs) {
	SFCryptFile(args.input, args.output, blockSize, passwordHash, args.threads)
}
