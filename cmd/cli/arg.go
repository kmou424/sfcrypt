package cli

import (
	"flag"
	"fmt"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
	"github.com/kmou424/sfcrypt/internal/consts"
	"github.com/kmou424/sfcrypt/v1/sfcrypt"
	"github.com/kmou424/sfcrypt/v1/sfcrypt/kit"
	"os"
)

var (
	Input     string
	Output    string
	Threads   int
	Overwrite bool
	Password  string
	Salt      string
	version   bool
)

func Parse() {
	flag.StringVar(&Input, "i", "", "specify input file")
	flag.StringVar(&Input, "input", "", "specify input file (same as -i)")

	flag.StringVar(&Output, "o", "", "specify output file")
	flag.StringVar(&Output, "output", "", "specify output file (same as -o)")

	flag.IntVar(&Threads, "t", sfcrypt.DefaultThreads, "specify threads to encrypt/decrypt file")
	flag.IntVar(&Threads, "threads", sfcrypt.DefaultThreads, "specify threads to encrypt/decrypt file (same as -t)")

	flag.StringVar(&Password, "p", "", "set password")
	flag.StringVar(&Password, "password", "", "set password (same as -p)")

	flag.StringVar(&Salt, "s", "", "set extra salt to increase security (suggested, optional)")
	flag.StringVar(&Salt, "salt", "", "set extra salt to increase security (suggested, optional, same as -s)")

	flag.BoolVar(&Overwrite, "overwrite", false, "overwrite input file (will ignore if output file is specified)")

	flag.BoolVar(&version, "v", false, "show version")
	flag.BoolVar(&version, "version", false, "show version (same as -v)")

	flag.Parse()

	argCheck()
}

func argCheck() {
	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Printf("sfcrypt %s by kmou424\n", consts.GetVersion())
		os.Exit(0)
	}

	if strutil.IsEmpty(Input) {
		kit.Panic("input file is required")
	}

	if strutil.IsEmpty(Output) && !Overwrite {
		kit.Panic("output file is required")
	}

	if Overwrite {
		fmt.Println("NOTICE: overwrite is enabled, sfcrypt will write to input file and ignore output file if specified")
	}

	if strutil.IsEmpty(Password) {
		kit.Panic("password is required")
	}

	if strutil.IsEmpty(Salt) {
		fmt.Println("WARN: we suggest you to use a salt to increase the security of your password")
	}

	if !fsutil.FileExists(Input) {
		kit.Panic("input file %s not found", Input)
	}

	if fsutil.FileExists(Output) {
		kit.Panic("output file %s is already existed", Output)
	}
}
