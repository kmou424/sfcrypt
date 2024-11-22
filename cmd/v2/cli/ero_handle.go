package cli

import (
	"fmt"
	"github.com/kmou424/ero"
	"github.com/kmou424/sfcrypt/app/buildinfo"
	"os"
)

func HandleEro() {
	if err := recover(); err != nil {
		err, ok := err.(error)
		if !ok {
			fmt.Println(err)
			return
		}
		if !ero.IsEro(err) {
			fmt.Println(err)
			return
		}
		//goland:noinspection GoBoolExpressions
		fmt.Println(ero.AllTrace(err, buildinfo.Debug))
		os.Exit(1)
	}
}
