package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sfcrypt",
	Short: "Simple file cryptor",
	Long:  `Simple file cryptor`,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
}

func Run() {
	defer HandleEro()
	_ = rootCmd.Execute()
}
