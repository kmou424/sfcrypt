package cli

import (
	"github.com/kmou424/ero"
	. "github.com/kmou424/sfcrypt/app/common"
	"github.com/spf13/cobra"
	"strings"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file using a secret key and optional salt",
	Long:  `Decrypt a file using a secret key and optional salt`,
	Run:   decryptCmdFunc,
}

func init() {
	bindCommonFlags(decryptCmd)
}

var decryptCmdFunc = func(cmd *cobra.Command, args []string) {
	beforeDecrypt()
	sfCipher := getSFCipher()
	err := sfCipher.Decrypt()
	if err != nil {
		panic(ero.Wrap(err, "failed to decrypt file"))
	}
}

func beforeDecrypt() {
	removeFileExtMark := false
	if cipherCmdOpt.output == "" {
		cipherCmdOpt.output = cipherCmdOpt.input
		removeFileExtMark = true
	}
	if cipherCmdOpt.input == cipherCmdOpt.output {
		removeFileExtMark = true
	}
	if removeFileExtMark {
		if strings.HasSuffix(cipherCmdOpt.output, SFCryptFileExt) {
			cipherCmdOpt.output = cipherCmdOpt.output[:len(cipherCmdOpt.output)-len(SFCryptFileExt)]
		}
	}
}
