package main

import (
	"fmt"
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/sfcrypt/cmd/cli"
	"github.com/kmou424/sfcrypt/internal/consts"
	"github.com/kmou424/sfcrypt/v1/sfcrypt"
	"os"
	"time"
)

func main() {
	defer exceptiongo.NewExceptionHandler(func(exception *exceptiongo.Exception) {
		exception.PrintStackTrace()
		os.Exit(1)
	}).Deploy()

	start()
}

func start() {
	if cli.Threads > consts.MaxThreads {
		cli.Threads = consts.MaxThreads
	} else if cli.Threads < 0 {
		cli.Threads = consts.DefaultThreads
	}

	if cli.Overwrite {
		cli.Output = cli.Input
	}

	start := time.Now()
	sfCrypt := sfcrypt.NewSFCrypt(cli.Password, cli.Salt)
	sfCrypt.SetThreads(cli.Threads)
	sfCrypt.CryptFile(cli.Input, cli.Output)
	end := time.Now()

	fmt.Println("Encrypt/Decrypt took: ", end.Sub(start))
}
