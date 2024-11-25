package cli

import (
	"github.com/kmou424/ero"
	"github.com/kmou424/sfcrypt/core/cipher"
	"github.com/kmou424/sfcrypt/core/keygen"
	v2 "github.com/kmou424/sfcrypt/core/v2"
	"github.com/spf13/cobra"
)

var cipherCmdOpt struct {
	key    string
	salt   string
	input  string
	output string
	force  bool
}

func bindCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&cipherCmdOpt.input, "input", "i", "", "")
	_ = cmd.MarkFlagRequired("input")
	cmd.Flags().StringVarP(&cipherCmdOpt.output, "output", "o", "", "")
	cmd.Flags().StringVarP(&cipherCmdOpt.key, "key", "k", "", "Secret key for encryption/decryption")
	_ = cmd.MarkFlagRequired("key")
	cmd.Flags().StringVarP(&cipherCmdOpt.salt, "salt", "s", "", "Salt to use for encryption/decryption")
	cmd.Flags().BoolVarP(&cipherCmdOpt.force, "force", "f", false, "Force continue even if any issues are encountered")
}

func getSFCipher() *v2.SFCipher {
	kg, err := v2.NewPBKDF2KeyGen(&keygen.Options{
		Password: []byte(cipherCmdOpt.key),
		Salt:     []byte(cipherCmdOpt.salt),
	})
	if err != nil {
		panic(ero.Wrap(err, "failed to generate key"))
	}

	bytesCipher := &v2.BytesCipher{}
	bytesCipher.Init(kg)

	sfCipher := &v2.SFCipher{}
	sfCipher.Init(&cipher.FileCipherOptions{
		Input:  cipherCmdOpt.input,
		Output: cipherCmdOpt.output,
		Force:  cipherCmdOpt.force,
		Cipher: bytesCipher,
	})

	return sfCipher
}
