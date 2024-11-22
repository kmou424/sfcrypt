package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sfcrypt",
	Short: "Simple file cryptor",
	Long:  `Simple file cryptor`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version and build information",
	Long:  `Show version and build information`,
	Run:   versionCmdFunc,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Run() {
	defer HandleEro()
	_ = rootCmd.Execute()
}
