package command

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "securerandom",
	Short: "Secure random number generators",
	Long:  "an interface to secure random number generators",
}

func Execute() error {
	return rootCmd.Execute()
}
