package cmd

import (
	"fmt"
	"github.com/erfanmomeniii/ks/internal/app"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Command for running tool",
	Run:   runFunc,
}

func runFunc(_ *cobra.Command, _ []string) {
	if err := app.Run(ConfigPath); err != nil {
		fmt.Println(err)
	}

	return
}
