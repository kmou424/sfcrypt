package main

import (
	"github.com/kmou424/sfcrypt/app/version"
	"github.com/kmou424/sfcrypt/cmd/v2/cli"
	v2 "github.com/kmou424/sfcrypt/core/v2"
)

func init() {
	version.InitVersion(1, 9, 9)
	v2.InitSFHeader()
}

func main() {
	cli.Run()
}
