package cmd

import (
	"rune/http"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP server",
}

var httpStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		server := http.NewServer()
		server.Start()
	},
}

func init() {
	httpCmd.AddCommand(httpStartCmd)
	rootCmd.AddCommand(httpCmd)
}