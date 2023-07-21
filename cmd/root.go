package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "ks",
	Short: "a lightweight tool for checking pod status",
}

func Execute() error {
	return rootCmd.Execute()
}
