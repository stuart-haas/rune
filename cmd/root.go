package cmd

import (
	"fmt"
	"log"
	"rune/node"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "rune",
	Short: "Welcome to the Rune CLI",
	Run: func(cmd *cobra.Command, args []string) {
		figure.NewFigure("RUNE", "small", true).Print()
		fmt.Println()
	},
}

var sshExecCmd = &cobra.Command{
	Use: "ssh-exec",
	Short: "Run ssh command on Node and display result once finished",
	Run: func(cmd *cobra.Command, args []string) {
		host := node.Host{
			Hostname: "localhost",
			User: "stuart",
		}
		if err := node.ExecuteSSH(host); err != nil {
			log.Fatal(err)
		}
	},
}

var sshStreamCmd = &cobra.Command{
	Use: "ssh-stream",
	Short: "Run ssh command on Node and stream result",
	Run: func(cmd *cobra.Command, args []string) {
		host := node.Host{
			Hostname: "localhost",
			User: "stuart",
		}
		if err := node.StreamSSH(host); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	viper.AutomaticEnv()

	rootCmd.AddCommand(sshExecCmd)
	rootCmd.AddCommand(sshStreamCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
