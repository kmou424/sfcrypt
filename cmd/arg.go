package main

import (
	"flag"
	"fmt"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"github.com/hanakogo/exceptiongo"
	"github.com/kmou424/sfcrypt/internal/consts"
	"github.com/kmou424/sfcrypt/internal/types"
	"os"
)

var input string
var output string
var threads int

var password string
var salt string

var version bool

func init() {
	flag.StringVar(&input, "i", "", "specify input file")
	flag.StringVar(&input, "input", "", "specify input file (same as -i)")

	flag.StringVar(&output, "o", "", "specify output file")
	flag.StringVar(&output, "output", "", "specify output file (same as -o)")

	flag.IntVar(&threads, "t", consts.DefaultThreads, "specify threads to encrypt/decrypt file")
	flag.IntVar(&threads, "threads", consts.DefaultThreads, "specify threads to encrypt/decrypt file (same as -t)")

	flag.StringVar(&password, "p", "", "set password")
	flag.StringVar(&password, "password", "", "set password (same as -p)")

	flag.StringVar(&salt, "s", "", "set extra salt to increase security (suggested, optional)")
	flag.StringVar(&salt, "salt", "", "set extra salt to increase security (suggested, optional, same as -s)")

	flag.BoolVar(&version, "v", false, "show version")
	flag.BoolVar(&version, "version", false, "show version (same as -v)")

	flag.Parse()

	argCheck()
}

func argCheck() {
	defer exceptiongo.NewExceptionHandler(func(exception *exceptiongo.Exception) {
		exception.PrintStackTrace()
		os.Exit(1)
	}).Deploy()

	if version {

		fmt.Printf("sfcrypt %s by kmou424\n", consts.GetVersion())
		os.Exit(0)
	}

	if strutil.IsEmpty(input) {
		exceptiongo.QuickThrowMsg[types.InvalidArgumentException]("input file is required")
	}

	if strutil.IsEmpty(output) {
		exceptiongo.QuickThrowMsg[types.InvalidArgumentException]("output file is required")
	}

	if strutil.IsEmpty(password) {
		exceptiongo.QuickThrowMsg[types.InvalidArgumentException]("password is required")
	}

	if strutil.IsEmpty(salt) {
		fmt.Println("WARN: we suggest you to use a salt to increase the security of your password")
	}

	if !fsutil.FileExists(input) {
		exceptiongo.QuickThrowMsg[types.FileNotFoundException](fmt.Sprintf("input file %s not found", input))
	}

	if fsutil.FileExists(output) {
		exceptiongo.QuickThrowMsg[types.FileNotFoundException](fmt.Sprintf("output file %s is already existed", output))
	}
}
