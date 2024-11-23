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
}

func Run() {
	defer HandleEro()
	_ = rootCmd.Execute()
}
