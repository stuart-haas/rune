package cmd

import (
	"log"
	"rune/ssh"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH commands",
}

var sshStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start ssh command on node.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ssh.NewSSHClient(ssh.SSHClientConfig{
			Hostname: "localhost",
			User:     "stuart",
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
	Use:   "run",
	Short: "Run ssh command on node.",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ssh.NewSSHClient(ssh.SSHClientConfig{
			Hostname: "localhost",
			User:     "stuart",
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
	sshCmd.AddCommand(sshStartCmd)
	sshCmd.AddCommand(sshRunCmd)
	rootCmd.AddCommand(sshCmd)
}