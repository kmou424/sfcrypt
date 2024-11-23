package main

import (
	"github.com/kmou424/sfcrypt/app/version"
	"github.com/kmou424/sfcrypt/cmd/v2/cli"
)

func init() {
	version.InitVersion(1, 9, 9)
}

func main() {
	cli.Run()
}
