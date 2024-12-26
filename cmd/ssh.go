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
			Hostname: cmd.Flag("hostname").Value.String(),
			User:     cmd.Flag("user").Value.String(),
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
			Hostname: cmd.Flag("hostname").Value.String(),
			User:     cmd.Flag("user").Value.String(),
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
	sshCmd.PersistentFlags().StringP("hostname", "h", "", "Hostname of the node")
	sshCmd.PersistentFlags().StringP("user", "u", "", "User of the node")
	sshCmd.AddCommand(sshStartCmd)
	sshCmd.AddCommand(sshRunCmd)
	rootCmd.AddCommand(sshCmd)
}