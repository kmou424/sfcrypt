package cli

import (
	"flag"
	. "github.com/kmou424/sfcrypt/app/common"
	"github.com/kmou424/sfcrypt/app/version"
	. "github.com/kmou424/sfcrypt/core/v1"
	"os"
)

var (
	Input     string
	Output    string
	Routines  int
	Overwrite bool
	Password  string
	Salt      string
	Version   bool
)

func Parse() {
	flag.StringVar(&Input, "i", "", "specify input file")
	flag.StringVar(&Input, "input", "", "specify input file (same as -i)")

	flag.StringVar(&Output, "o", "", "specify output file")
	flag.StringVar(&Output, "output", "", "specify output file (same as -o)")

	flag.IntVar(&Routines, "r", DefaultRoutines, "specify goroutines to encrypt/decrypt file")
	flag.IntVar(&Routines, "routines", DefaultRoutines, "specify goroutines to encrypt/decrypt file (same as -t)")

	flag.StringVar(&Password, "p", "", "set password")
	flag.StringVar(&Password, "password", "", "set password (same as -p)")

	flag.StringVar(&Salt, "s", "", "set extra salt to increase security (suggested, optional)")
	flag.StringVar(&Salt, "salt", "", "set extra salt to increase security (suggested, optional, same as -s)")

	flag.BoolVar(&Overwrite, "overwrite", false, "overwrite input file (will ignore if output file is specified)")

	flag.BoolVar(&Version, "v", false, "show version")
	flag.BoolVar(&Version, "version", false, "show version (same as -v)")

	flag.Parse()

	argCheck()
}

func argCheck() {
	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if Version {
		Logger.Info("sfcrypt %s by kmou424\n", version.GetVersion())
		os.Exit(0)
	}

	if len(Input) == 0 {
		Panic("input file is required")
	}

	if len(Output) == 0 && !Overwrite {
		Panic("output file is required")
	}

	if Overwrite {
		Logger.Info("NOTICE: overwrite is enabled, sfcrypt will write to input file and ignore output file if specified")
	}

	if len(Password) == 0 {
		Panic("password is required")
	}

	if len(Salt) == 0 {
		Logger.Info("WARN: we suggest you to use a salt to increase the security of your password")
	}

	if !FileExists(Input) {
		Panic("input file %s not found", Input)
	}

	if FileExists(Output) {
		Panic("output file %s is already existed", Output)
	}
}
