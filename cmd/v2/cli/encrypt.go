package cli

import (
	"github.com/kmou424/ero"
	. "github.com/kmou424/sfcrypt/app/common"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a file using a secret key and optional salt",
	Long:  `Encrypt a file using a secret key and optional salt`,
	Run:   encryptCmdFunc,
}

func init() {
	bindCommonFlags(encryptCmd)
}

var encryptCmdFunc = func(cmd *cobra.Command, args []string) {
	beforeEncrypt()
	sfCipher := getSFCipher()
	err := sfCipher.Encrypt()
	if err != nil {
		panic(ero.Wrap(err, "failed to encrypt file"))
	}
}

func beforeEncrypt() {
	appendFileExtMark := false
	if cipherCmdOpt.output == "" {
		cipherCmdOpt.output = cipherCmdOpt.input
		appendFileExtMark = true
	}
	if cipherCmdOpt.input == cipherCmdOpt.output {
		appendFileExtMark = true
	}
	if appendFileExtMark {
		cipherCmdOpt.output += SFCryptFileExt
	}
}
