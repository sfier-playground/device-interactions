package cmd

import (
	"github.com/sifer169966/device-interactions/pkg/flags"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "user task",
	Short:   "task manages user task",
	Version: flags.Version,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
