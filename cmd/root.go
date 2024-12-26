package cmd

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "rune",
	Short: "Welcome to the Rune CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		figure.NewFigure("RUNE", "small", true).Print()
		fmt.Println()
	},
}

func init() {
	viper.AutomaticEnv()
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
