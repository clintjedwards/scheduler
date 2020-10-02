package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//RootCmd is the base command for the cli
var RootCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
