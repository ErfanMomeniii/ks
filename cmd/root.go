package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "ks",
	Short: "a lightweight tool for checking pod status",
}

var (
	ConfigPath string
)

func init() {
	homeDir, _ := os.UserHomeDir()

	defaultConfigPath := fmt.Sprintf("%s%s%s%s%s", homeDir, string(os.PathSeparator), ".kube", string(os.PathSeparator), "config")

	rootCmd.PersistentFlags().StringVarP(
		&ConfigPath, "config", "c", defaultConfigPath, "kube config path",
	)

	rootCmd.AddCommand(runCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
