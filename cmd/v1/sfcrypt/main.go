package main

import (
	. "github.com/kmou424/sfcrypt/app/common"
	"github.com/kmou424/sfcrypt/app/version"
	"github.com/kmou424/sfcrypt/cmd/v1/cli"
	. "github.com/kmou424/sfcrypt/core/v1"
	"os"
	"time"
)

func init() {
	version.InitVersion(1, 1, 0)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			Logger.Info("error:", err)
			os.Exit(1)
		}
	}()

	cli.Parse()
	start()
}

func start() {
	if cli.Routines > MaxRoutines {
		cli.Routines = MaxRoutines
	} else if cli.Routines < 0 {
		cli.Routines = DefaultRoutines
	}

	if cli.Overwrite {
		cli.Output = cli.Input
	}

	start := time.Now()
	sfCrypt := NewSFCrypt(cli.Password, cli.Salt)
	sfCrypt.SetThreads(cli.Routines)
	err := sfCrypt.CryptFile(cli.Input, cli.Output)
	if err != nil {
		panic(err)
	}
	end := time.Now()

	Logger.Info("Encrypt/Decrypt took: ", end.Sub(start))
}
