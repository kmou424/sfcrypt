package main

import (
	"fmt"
	"github.com/kmou424/sfcrypt/cmd/cli"
	"github.com/kmou424/sfcrypt/v1/sfcrypt"
	"os"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}()

	cli.Parse()
	start()
}

func start() {
	if cli.Threads > sfcrypt.MaxThreads {
		cli.Threads = sfcrypt.MaxThreads
	} else if cli.Threads < 0 {
		cli.Threads = sfcrypt.DefaultThreads
	}

	if cli.Overwrite {
		cli.Output = cli.Input
	}

	start := time.Now()
	sfCrypt := sfcrypt.NewSFCrypt(cli.Password, cli.Salt)
	sfCrypt.SetThreads(cli.Threads)
	err := sfCrypt.CryptFile(cli.Input, cli.Output)
	if err != nil {
		panic(err)
	}
	end := time.Now()

	fmt.Println("Encrypt/Decrypt took: ", end.Sub(start))
}
