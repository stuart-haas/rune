package cmd

import (
	"fmt"
	"log"
	"rune/services"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "rune",
	Short: "Welcome to the Rune CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		figure.NewFigure("RUNE", "small", true).Print()
		fmt.Println()
	},
}

var sshStartCmd = &cobra.Command{
	Use: "ssh-start",
	Short: "Start ssh command on node.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := services.NewSSHClient(services.SSHClientConfig{
			Hostname: "localhost",
			User: "stuart",
		})
		if err != nil {
			log.Fatal(err)
		}
		if err := client.Start(); err != nil {
			log.Fatal(err)
		}
	},
}

var sshRunCmd = &cobra.Command{
	Use: "ssh-run",
	Short: "Run ssh command on node.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := services.NewSSHClient(services.SSHClientConfig{
			Hostname: "localhost",
			User: "stuart",
		})
		if err != nil {
			log.Fatal(err)
		}
		if err := client.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	viper.AutomaticEnv()

	rootCmd.AddCommand(sshStartCmd)
	rootCmd.AddCommand(sshRunCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
