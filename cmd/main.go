package main

import (
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/sfcrypt/internal/core"
	"os"
)

func main() {
	defer exceptiongo.NewExceptionHandler(func(exception *exceptiongo.Exception) {
		exception.PrintStackTrace()
		os.Exit(1)
	}).Deploy()

	core.SFCryptCli(core.NewSFCryptCliArgs(input, output, password, salt, overwrite, threads))
}
